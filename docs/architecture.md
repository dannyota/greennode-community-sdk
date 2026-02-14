# SDK Architecture

## 1. Overview

greennode-community-sdk is a multi-layered Go SDK for VNG Cloud services, built on
[`imroc/req/v3`](https://github.com/imroc/req) as its single direct dependency. The SDK
provides typed clients for compute, networking, storage, load balancing, DNS, and identity
services exposed by VNG Cloud's REST APIs.

- **Module path:** `github.com/dannyota/greennode-community-sdk/v2`
- **Go version:** 1.24
- **Source files:** 244 `.go` files, ~26 k LOC

## 2. Package Map

```
greennode-community-sdk/
├── client/                        SDK entry point — top-level Client, SdkConfigure
├── greennode/
│   ├── client/                    Low-level HTTP client, ServiceClient, request builder
│   ├── gateway/                   Versioned gateway multiplexers (6 gateways)
│   ├── services/                  Per-service business logic (9 services, versioned)
│   │   ├── common/                Shared embedded types (UserAgent, Paging)
│   │   ├── compute/               Server lifecycle, floating IPs, server groups
│   │   ├── dns/                   Hosted zones, DNS records
│   │   ├── glb/                   Global load balancer pools, listeners, health checks
│   │   ├── identity/              OAuth2 token acquisition
│   │   ├── loadbalancer/          Load balancers, listeners, pools, policies, certs
│   │   ├── network/               VPCs, subnets, security groups, endpoints
│   │   ├── portal/                Portal info, project listing
│   │   ├── server/                Internal server system tags
│   │   └── volume/                Block volumes, snapshots, volume types
│   ├── entity/                    Domain model structs (26 entities)
│   └── sdkerror/                  Error codes, categories, handler chain
└── test/                          Integration tests
```

## 3. Layered Architecture

```
┌─────────────────────────────────────────────────────────────┐
│  User Code                                                  │
│  client.VServerGateway().V2().ComputeService().CreateServer │
└──────────────────────────┬──────────────────────────────────┘
                           │
┌──────────────────────────▼──────────────────────────────────┐
│  1. Client               client/client.go                   │
│     Orchestration, gateway creation, auth configuration     │
└──────────────────────────┬──────────────────────────────────┘
                           │
┌──────────────────────────▼──────────────────────────────────┐
│  2. Gateway              greennode/gateway/                  │
│     Version multiplexing — V1 / V2 / Internal per service   │
└──────────────────────────┬──────────────────────────────────┘
                           │
┌──────────────────────────▼──────────────────────────────────┐
│  3. Service              greennode/services/*/v*/            │
│     Request building, URL construction, error enrichment    │
└──────────────────────────┬──────────────────────────────────┘
                           │
┌──────────────────────────▼──────────────────────────────────┐
│  4. HTTP Client          greennode/client/                   │
│     Token management, retry, reauth, request execution      │
└──────────────────────────┬──────────────────────────────────┘
                           │
┌──────────────────────────▼──────────────────────────────────┐
│  5. API                  *.vngcloud.vn REST endpoints        │
└─────────────────────────────────────────────────────────────┘
```

**Data flows down** through these layers on every SDK call. Responses flow back
up, converted from raw JSON into domain entities at the service layer.

## 4. Authentication

The SDK uses an IAM OAuth2 client-credentials flow.

### Token Acquisition

1. `POST {iam_endpoint}/v2/auth/token` with HTTP Basic auth
   (`base64(clientId:clientSecret)`)
2. Response contains `access_token` and `expires_in` (seconds)
3. Token stored in `httpClient` as `SdkAuthentication{accessToken, expiresAt}`

Implementation:
- Token request: `greennode/services/identity/v2/identity.go`
- URL builder: `greennode/services/identity/v2/url.go`
- Response conversion: `greennode/services/identity/v2/identity_response.go`

### Token Refresh

Two refresh strategies run automatically:

| Strategy | Trigger | Mechanism |
|----------|---------|-----------|
| **Proactive** | Token expires within 5 minutes | `NeedReauth()` in `greennode/client/service_client.go` checks `time.Until(expiresAt) < 5*time.Minute` before every request |
| **Reactive** | HTTP 401 response | `handleUnauthorized()` in `greennode/client/http.go` triggers reauth then retries the original request |

### Concurrency Safety

Token access is protected by `sync.RWMutex` (`httpClient.mut`). A separate
`reauthlock` struct with a `reauthFuture` (channel-based future) prevents
thundering-herd reauth — the first goroutine to request a refresh creates a
future; concurrent goroutines block on `future.done` and reuse the result.

Key fields in `httpClient` (`greennode/client/http.go`):

```
mut       *sync.RWMutex   — guards accessToken and defaultHeaders
reauthmut *reauthlock      — serializes token refresh
  └─ ongoing *reauthFuture
       ├─ done chan struct{}
       └─ err  Error
```

## 5. Request Lifecycle

End-to-end trace of `client.VServerGateway().V2().ComputeService().CreateServer(req)`:

```
Step  Layer        Code path
────  ───────────  ──────────────────────────────────────────────
 1    Client       client/client.go — VServerGateway() returns gateway
 2    Gateway      greennode/gateway/gateway.go — V2() returns versioned gateway
 3    Gateway      greennode/gateway/vserver_gateway.go — ComputeService() returns service
 4    Service      greennode/services/compute/v2/server.go — CreateServer()
 4a     URL        greennode/services/compute/v2/url.go — builds {endpoint}/v2/{projectId}/servers
 4b     Request    greennode/client/request.go — NewRequest().WithJSONBody().WithOkCodes(202)...
 4c     Dispatch   service calls s.VServerClient.Post(url, req)
 5    HTTP Client  greennode/client/service_client.go — Post() delegates to httpClient.DoRequest()
 6    HTTP Client  greennode/client/http.go — DoRequest():
 6a     Prepare      prepareRequest() — set context, headers, marshal body
 6b     Auth         needReauth() → reauthenticate() if token is nil/expiring
 6c     Execute      executeHTTPMethod() — req/v3 HTTP call
 6d     Handle       handleResponse() → handleStatusCode()
 6e     Retry        401 → handleUnauthorized() → reauthenticate() → retry DoRequest()
 7    Response     JSON unmarshaled into entity struct; error enriched via SdkErrorHandler
```

## 6. Error Handling

### Interface

`Error` (`greennode/sdkerror/isdk_error.go`) provides:
- Error code queries: `IsError(code)`, `IsCategory(cat)`
- Builder methods: `WithErrorCode()`, `WithMessage()`, `WithParameters()`
- Getters: `ErrorCode()`, `GetMessage()`, `Parameters()`

### Handler Chain

`SdkErrorHandler` (`greennode/sdkerror/common.go`) applies a chain of
pattern-matching handler functions to classify errors:

```go
sdkerror.SdkErrorHandler(sdkErr, errResp,
    sdkerror.WithErrorSubnetNotFound(errResp),
    sdkerror.WithErrorImageNotFound(errResp),
    sdkerror.WithErrorServerExceedQuota(errResp),
    // ...more handlers
)
```

Each handler inspects the error response message (string matching) and, if it
matches, calls `WithErrorCode()` on the `Error`. The chain stops at the first
match (error code changes from `EcUnknownError`).

### Error Enrichment

At the service layer, errors are enriched with context:

```go
return nil, sdkerror.SdkErrorHandler(sdkErr, errResp, ...handlers...).
    WithParameters(opts.ToMap()).
    WithKVparameters("projectId", s.getProjectID()).
    WithErrorCategories(sdkerror.ErrCatVServer)
```

### Counts

- **100 error codes** defined in `greennode/sdkerror/error_codes.go`
- **10 error categories** in `greennode/sdkerror/categories.go`
  (`ErrCatQuota`, `ErrCatIam`, `ErrCatInfra`, `ErrCatPurchase`, `ErrCatAll`,
  `ErrCatProductVlb`, `ErrCatProductVNetwork`, `ErrCatProductVdns`,
  `ErrCatVServer`, `ErrCatVirtualAddress`)
- Handler files: `server.go`, `volume.go`, `loadbalancer.go`, `network.go`,
  `secgroup.go`, `secgroup_rule.go`, `snapshot.go`, `quota.go`, `endpoint.go`,
  `identity.go`, `virtualaddress.go`

## 7. Key Design Patterns

### Builder Pattern
Request objects use fluent builders (`NewRequest().WithJSONBody().WithOkCodes()`).
SDK configuration follows the same style (`NewClient().WithAuthOption().Configure()`).

### Version Multiplexing
Each gateway creates per-version `ServiceClient` instances with versioned endpoint
suffixes (`endpoint + "v1"`, `endpoint + "v2"`, `endpoint + "internal"`). User code
selects a version via `gateway.V2()`, which returns a typed interface exposing only
that version's services.

### Entity Conversion
API responses are unmarshaled into `*Response` structs (in service packages), then
converted to domain entities via `ToEntity*()` methods (e.g.,
`CreateServerResponse.ToEntityServer()`). This keeps API wire format separate from
the public domain model.

### URL Helpers
Each service version has a `url.go` file with functions that build endpoint paths
from a `ServiceClient` (e.g., `createServerURL(s.VServerClient)` →
`{endpoint}/v2/{projectId}/servers`).

### Shared Commons
`greennode/services/common/` provides embedded types (`UserAgent`, `Paging`) that
services compose into their request structs for consistent field handling.

### Functional-Option Error Handlers
Error handlers are functions with signature `func(Error)` passed as variadic
arguments to `SdkErrorHandler`. This makes the handler chain composable and lets
each service choose which error patterns to match.

## 8. Service Catalog

| Service | Package | Versions | Purpose |
|---------|---------|----------|---------|
| **Compute** | `services/compute` | v2 | Server lifecycle, floating IPs, server groups |
| **Volume** | `services/volume` | v1, v2 | Block volumes, snapshots, volume types |
| **Network** | `services/network` | v1, v2, internal | VPCs, subnets, security groups, endpoints, virtual addresses |
| **Load Balancer** | `services/loadbalancer` | v2, internal | Load balancers, listeners, pools, policies, certificates |
| **GLB** | `services/glb` | v1 | Global load balancer pools, listeners, health checks |
| **DNS** | `services/dns` | v1, internal | Hosted zones, DNS records |
| **Identity** | `services/identity` | v2 | OAuth2 token acquisition |
| **Portal** | `services/portal` | v1, v2 | Portal info, project listing |
| **Server** | `services/server` | internal | Internal server system tags |

### Gateway Routing

| Gateway | Versions | Services |
|---------|----------|----------|
| **IAM** | v2 | Identity |
| **VServer** | v1, v2, internal | Compute, Volume, Portal, Network, Server |
| **VLB** | v2, internal | Load Balancer |
| **VNetwork** | v1, v2, internal | Network |
| **GLB** | v1 | GLB |
| **VDns** | v1, internal | DNS |
