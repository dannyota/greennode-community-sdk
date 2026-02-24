# SDK Architecture

## 1. Overview

greennode-community-sdk is a multi-layered Go SDK for VNG Cloud services, built on
[`imroc/req/v3`](https://github.com/imroc/req) as its single direct dependency. The SDK
provides typed clients for compute, networking, storage, load balancing, DNS, and identity
services exposed by VNG Cloud's REST APIs.

- **Module path:** `danny.vn/greennode`
- **Go version:** 1.24
- **Source files:** 195 `.go` files, ~19.5 k LOC

## 2. Package Map

```
greennode-community-sdk/
├── greennode/
│   ├── greennode.go               SDK entry point — Config, Client, NewClient()
│   ├── endpoints.go               Default endpoint resolution from Region + auth method
│   ├── auth/                      IAM user auth (PKCE + TOTP) and TOTP providers
│   ├── client/                    Low-level HTTP client, ServiceClient, request builder
│   ├── services/                  Per-service business logic (9 services, versioned)
│   │   ├── common/                Shared helpers (StructToMap, Paging)
│   │   ├── compute/               Server lifecycle, floating IPs, server groups
│   │   ├── dns/                   Hosted zones, DNS records
│   │   ├── glb/                   Global load balancer pools, listeners, health checks
│   │   ├── identity/              OAuth2 token acquisition
│   │   ├── loadbalancer/          Load balancers, listeners, pools, policies, certs
│   │   ├── network/               VPCs, subnets, security groups, endpoints
│   │   ├── portal/                Portal info, project listing
│   │   ├── server/                Internal server system tags
│   │   └── volume/                Block volumes, snapshots, volume types
│   ├── entity/                    Domain model structs (~79 types)
│   └── sdkerror/                  Error codes, categories, handler chain
└── test/                          Integration tests
```

## 3. Layered Architecture

```
┌─────────────────────────────────────────────────────────────┐
│  User Code                                                  │
│  client.Compute.CreateServer(ctx, req)                      │
└──────────────────────────┬──────────────────────────────────┘
                           │
┌──────────────────────────▼──────────────────────────────────┐
│  1. Client               greennode/greennode.go              │
│     Config, flat service fields, auth wiring                │
└──────────────────────────┬──────────────────────────────────┘
                           │
┌──────────────────────────▼──────────────────────────────────┐
│  2. Service              greennode/services/*/v*/            │
│     Request building, URL construction, error enrichment    │
└──────────────────────────┬──────────────────────────────────┘
                           │
┌──────────────────────────▼──────────────────────────────────┐
│  3. HTTP Client          greennode/client/                   │
│     Token management, retry, reauth, request execution      │
└──────────────────────────┬──────────────────────────────────┘
                           │
┌──────────────────────────▼──────────────────────────────────┐
│  4. API                  *.vngcloud.vn REST endpoints        │
└─────────────────────────────────────────────────────────────┘
```

**Data flows down** through these layers on every SDK call. Responses flow back
up, converted from raw JSON into domain entities at the service layer.

## 4. Authentication

The SDK supports two authentication methods, selected by `Config` fields.

### 4a. Client Credentials (service accounts)

Set `ClientID` + `ClientSecret` in Config. Uses `IAMOauth2` auth option.

1. `POST {iam_endpoint}/v2/auth/token` with HTTP Basic auth
   (`base64(clientId:clientSecret)`)
2. Response contains `access_token` and `expires_in` (seconds)
3. Token stored in `httpClient` as `SdkAuthentication{accessToken, expiresAt}`

Implementation:
- Token request: `greennode/services/identity/v2/identity.go`
- URL builder: `greennode/services/identity/v2/url.go`
- Response conversion: `greennode/services/identity/v2/identity_response.go`
- Authentication type: `greennode/client/auth.go`

### 4b. IAM User Auth (username/password + optional TOTP)

Set `IAMAuth` in Config with an `*auth.IAMUserAuth`. Uses `IAMUserOauth2` auth option.

1. Generate PKCE verifier/challenge
2. GET login page at `signin.vngcloud.vn`, extract CSRF token
3. POST username/password credentials
4. If 2FA required (redirect to `/ap/auth/iam/google`):
   - GET 2FA page, extract CSRF
   - Call `TOTPProvider.GetCode()` for the code
   - POST TOTP code
5. Extract authorization code from final redirect URL
6. POST token exchange to dashboard API with Basic auth (dashboard client ID)
7. Return `accessToken` and computed `expiresAt`

Implementation:
- Auth flow: `greennode/auth/iam_user.go`
- TOTP providers: `greennode/auth/totp.go` (`TOTPProvider` interface, `TOTPFunc` adapter, `SecretTOTP`)
- Wiring: `iamUserReauthFunc()` in `greennode/greennode.go`

### Default Endpoints from Region

`resolveEndpoints()` in `greennode/endpoints.go` fills empty endpoint fields
based on `Config.Region` and the auth method. IAM user auth uses different
gateway URLs (e.g. `iam-vserver-gateway` vs `vserver-gateway`) for VServer and
VLB. Explicit endpoint fields always override the defaults.

### Token Refresh

Two refresh strategies run automatically:

| Strategy | Trigger | Mechanism |
|----------|---------|-----------|
| **Proactive** | Token expires within 5 minutes | `NeedReauth()` in `greennode/client/auth.go` checks `time.Until(expiresAt) < 5*time.Minute` before every request |
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

End-to-end trace of `client.Compute.CreateServer(ctx, req)`:

```
Step  Layer        Code path
────  ───────────  ──────────────────────────────────────────────
 1    Client       greennode/greennode.go — client.Compute is *ComputeServiceV2
 2    Service      greennode/services/compute/v2/server.go — CreateServer(ctx, opts)
 2a     URL        greennode/services/compute/v2/url.go — builds {endpoint}/v2/{projectId}/servers
 2b     Request    greennode/client/request.go — NewRequest().WithJSONBody(opts).WithOkCodes(202)...
 2c     Dispatch   service calls s.Client.Post(ctx, url, req)
 3    HTTP Client  greennode/client/service_client.go — Post() delegates to httpClient.DoRequest(ctx, ...)
 4    HTTP Client  greennode/client/http.go — DoRequest(ctx, ...):
 4a     Prepare      prepareRequest(ctx, ...) — set headers, marshal body
 4b     Auth         needReauth() → reauthenticate(ctx) if token is nil/expiring
 4c     Execute      executeHTTPMethod(...) — req/v3 HTTP call
 4d     Handle       handleResponse() → handleStatusCode()
 4e     Retry        401 → handleUnauthorized(ctx, ...) → reauthenticate(ctx) → retry
 5    Response     JSON unmarshaled into entity struct; error enriched via SdkErrorHandler
```

## 6. Error Handling

### SdkError

`*SdkError` (`greennode/sdkerror/sdk_error.go`) provides:
- Error code queries: `IsError(code)`, `IsCategory(cat)`
- Builder methods: `WithErrorCode()`, `WithMessage()`, `WithParameters()`
- Getters: `ErrorCode()`, `GetMessage()`, `Parameters()`
- Standard `error` interface: `Error()`, `Unwrap()`, `Is()`

### Classifier Registry

Error classification uses a table-driven registry (`greennode/sdkerror/classifier.go`).
Each `ErrorCode` registers a `classifier` struct with a match function (string
contains, regex, error-code comparison) and optional category/message format.

`SdkErrorHandler` (`greennode/sdkerror/common.go`) takes error codes as variadic
arguments and iterates them against the registry:

```go
sdkerror.SdkErrorHandler(sdkErr, errResp,
    sdkerror.EcVServerSubnetNotFound,
    sdkerror.EcVServerImageNotFound,
    sdkerror.EcVServerServerExceedQuota,
    // ...more error codes
)
```

The first matching classifier sets the error code, message, and optional category.

### Error Enrichment

At the service layer, errors are enriched with context:

```go
return nil, sdkerror.SdkErrorHandler(sdkErr, errResp,
    sdkerror.EcVServerServerNotFound,
    sdkerror.EcVServerVolumeInProcess).
    WithParameters(common.StructToMap(opts)).
    WithKVparameters("projectId", s.getProjectID()).
    AppendCategories(sdkerror.ErrCatVServer)
```

### Counts

- **100 error codes** defined in `greennode/sdkerror/error_codes.go`
- **10 error categories** in `greennode/sdkerror/categories.go`
- Classifier registration files: `server.go`, `volume.go`, `loadbalancer.go`,
  `network.go`, `secgroup.go`, `secgroup_rule.go`, `snapshot.go`, `endpoint.go`,
  `identity.go`, `virtualaddress.go`, `volumetype.go`, `common.go`

## 7. Key Design Patterns

### Builder Pattern
Request objects use fluent builders (`NewRequest().WithJSONBody().WithOkCodes()`).
SDK configuration follows the same style (`NewClient().WithAuthOption().Configure()`).

### Version Multiplexing
`NewClient()` creates per-version `ServiceClient` instances with versioned endpoint
suffixes (`endpoint + "v1"`, `endpoint + "v2"`, `endpoint + "internal"`). Each
service struct receives the `ServiceClient` for its API version. The `Client`
struct exposes primary services (latest version) as top-level fields and
legacy/internal versions as secondary fields.

### Entity Conversion
API responses are unmarshaled into `*Response` structs (in service packages), then
converted to domain entities via `ToEntity*()` methods (e.g.,
`CreateServerResponse.ToEntityServer()`). This keeps API wire format separate from
the public domain model.

### URL Helpers
Each service version has a `url.go` file with functions that build endpoint paths
from a `ServiceClient` (e.g., `createServerURL(s.Client)` →
`{endpoint}/v2/{projectId}/servers`).

### Shared Commons
`greennode/services/common/` provides `StructToMap()` (generic struct-to-map
conversion for error parameter enrichment) and the `Paging` embedded type for
paginated request structs.

### Table-Driven Error Classification
Error codes are registered in `init()` functions with match functions (string
contains, regex, error-code comparison). `SdkErrorHandler` iterates the
caller-provided error codes against the classifier registry, stopping at the
first match.

## 8. Service Catalog

| Service | Package | Versions | Purpose |
|---------|---------|----------|---------|
| **Compute** | `services/compute` | v2 | Server lifecycle, floating IPs, server groups |
| **Volume** | `services/volume` | v1, v2 | Block volumes, snapshots, volume types |
| **Network** | `services/network` | v1, v2 | VPCs, subnets, security groups, endpoints, virtual addresses |
| **Load Balancer** | `services/loadbalancer` | v2, inter | Load balancers, listeners, pools, policies, certificates |
| **GLB** | `services/glb` | v1 | Global load balancer pools, listeners, health checks |
| **DNS** | `services/dns` | v1, internal_system | Hosted zones, DNS records |
| **Identity** | `services/identity` | v2 | OAuth2 token acquisition |
| **Portal** | `services/portal` | v1, v2 | Portal info, project listing |
| **Server** | `services/server` | v1 | Internal server system tags |
| **Auth** | `auth` | — | IAM user auth (PKCE + TOTP), TOTP providers |

### Client Service Wiring

`NewClient()` in `greennode/greennode.go` wires each API endpoint to its services:

| Endpoint | Version suffixes | Services |
|----------|-----------------|----------|
| **IAM** | v2 | Identity |
| **VServer** | v1, v2, internal | Compute, Volume, Portal, Network, Server |
| **VLB** | v2, internal | Load Balancer |
| **VNetwork** | vnetwork/v1, vnetwork/az/v1, internal/v1 | Network (V1, AZ, Internal) |
| **GLB** | v1 | GLB |
| **DNS** | v1, internal/v1 | DNS |

