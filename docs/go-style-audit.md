# Go Style Audit

This document catalogs every non-idiomatic Go pattern inherited from the upstream
Java/C#-style codebase. It serves as the roadmap for incremental refactoring toward
idiomatic Go.

Counts were gathered programmatically and may shift as the codebase evolves; treat
them as order-of-magnitude guides, not exact totals.

---

## 1. Naming Conventions

### 1.1 `I` prefix on interfaces — **RESOLVED (non-request interfaces)**

Non-request interfaces (~73) have been renamed: `IClient` → `Client`,
`IHttpClient` → `HTTPClient`, `IError` → `Error`, `IIamGateway` → `IamGateway`, etc.

The ~106 `I*Request` interfaces (e.g., `ICreateLoadBalancerRequest`) that had
name collisions with their concrete structs have been removed entirely per §2.2.
`IBulkActionRequest` (previously the sole exception) was also deleted when
`ToMap()` was removed (§5.7).

### 1.2 `p` prefix on parameters — **RESOLVED**

All function parameters have been stripped of the `p` prefix (~1,456 occurrences
across 140 files). Natural words starting with `p` (e.g., `portalUserID`,
`poolID`, `path`, `policies`) were correctly excluded.

Residual `psdkErr` in `sdkerror/common.go` and `perrResp` in `test/error_test.go`
(function-type parameters missed by the AST tool) have also been fixed.

### 1.3 `s` receiver name on all types — **RESOLVED**

All methods now use type-appropriate receiver names (~967 methods across 86 files):

| Type | Receiver |
|------|----------|
| `client` (top-level) | `c` |
| `httpClient` | `hc` |
| `serviceClient` | `sc` |
| `request` | `r` |
| `sdkConfigure` | `sc` |
| `SdkError` | `e` |
| `*Gateway` types | `g` |
| `*ServiceV2` types | `s` (kept — already appropriate for "service") |
| Entity types | type-specific (`e` for Endpoint, `lb` for LoadBalancer, etc.) |
| Request/Response types | `r` |

### 1.4 Acronym casing — **RESOLVED**

All acronyms in exported identifiers now use correct ALL CAPS casing:

| Pattern | Status |
|---------|--------|
| `Id` → `ID` | Done (~284 occurrences). `WithClientID()`, `GetProjectID()`, `WithZoneID()`, etc. |
| `Json` → `JSON` | Done. `WithJSONBody()`, `WithJSONResponse()`, `WithJSONError()`, etc. |
| `Http` → `HTTP` | Done. `HTTPClient`, `executeHTTPMethod()` |
| `Url` → `URL` | Done where applicable. `createServerURL()`, `ServiceURL()`, etc. |
| `Iam` → `IAM` | Done. `IAMGateway`, `IAMEndpoint()`, `IAMErrorResponse`, etc. |

**Note:** `SetBodyJsonMarshal` is an external method from `imroc/req/v3` — not renamed.

### 1.5 Abbreviation casing (secondary pass) — **RESOLVED**

Remaining abbreviation casing issues in constant names, struct fields, and a
struct field name inconsistency:

| Pattern | Status |
|---------|--------|
| `Http1`/`Http1Minor1` → `HTTP1`/`HTTP1Minor1` | Done. `HealthCheckHTTPVersionHTTP1`, `GlobalPoolHealthCheckHTTPVersionHTTP1Minor1`, etc. (loadbalancer/v2, inter, glb/v1) |
| `HealthCheckProtocolHTTPs` → `HTTPS` | Done. Trailing lowercase `s` fixed in all 3 packages + tests. |
| `IpAddress` → `IPAddress` | Done. Struct field in `Member` (loadbalancer/v2, inter) and `CreateVirtualAddressCrossProjectRequest` (network/v2). JSON tags unchanged. |
| `VirtualIpAddressID` → `VirtualIPAddressID` | Done. Entity field + response structs (entity/address_pair, network/v2). |
| `NetworkInterfaceIp` → `NetworkInterfaceIP` | Done. Entity field + response structs. |
| `VserverClient` → `VServerClient` | Done. Struct field in `NetworkServiceV2` + all usages (9 files in network/v2, 2 gateway files). |

Also removed: commented-out struct fields (`address_pair_response.go`),
commented-out error handling code (`address_pair.go`), and a lone `/** */`
block comment (`virtualaddress_request.go`).

### 1.6 Java-style getters — **PARTIALLY RESOLVED**

(Renumbered from §1.5)

Most `Get*()` accessors have been simplified:

| Change | Status |
|--------|--------|
| `GetErrorCode()` → `ErrorCode()` | Done |
| `GetParameters()` → `Parameters()` | Done |
| `GetErrorCategories()` → `ErrorCategories()` | Done |
| `GetAccessToken()` → `AccessToken()` | Done |
| `GetExpiresAt()` → `ExpiresAt()` | Done |

**Kept with `Get` prefix** (name collisions with exported struct fields used for
JSON serialization on request types):
- `GetProjectID()`, `GetZoneID()`, `GetUserID()`
- `GetMessage()` (collides with `Message` field on error response structs)

### 1.7 Package names with underscores — **RESOLVED**

The `sdk_error` package has been renamed to `sdkerror` (`greennode/sdkerror/`,
18 files). All 42 importing files updated.

---

## 2. Interface Design

### 2.1 Producer-side interfaces — **RESOLVED**

All producer-side interfaces have been deleted. Gateway and client interfaces
were removed earlier; service-level interfaces (38 total: 7 composition +
31 sub-interfaces across 9 parent-package files) have now been deleted as well.
Gateway accessor methods return concrete pointers (e.g.,
`*identityv2.IdentityServiceV2`, `*lbv2.LoadBalancerServiceV2`), and gateway
files import sub-packages directly instead of going through parent packages.

**Go convention:** Define interfaces at the call site (the "consumer" side).
The producer exports concrete types; consumers define the small interfaces they
actually need. This yields naturally small interfaces and decouples packages.

**Reference:** [Go wiki: Accept interfaces, return structs](https://go.dev/wiki/CodeReviewComments#interfaces).

### 2.2 Interface-per-type (one-to-one interface/struct pairs) — **RESOLVED**

All one-to-one `I*Request` interface/struct pairs (~142) have been removed.
Method return types (`With*()`) now return `*ConcreteType`
instead of `I*Interface`. URL builders, service implementations, and parent
service interfaces all accept concrete pointer types directly.

The `IBulkActionRequest` interface (previously the sole exception) was also
deleted when its only method `ToMap()` was removed (§5.7). The `BulkActions`
field now uses `[]any`.

### 2.3 God interfaces (>10 methods) — **RESOLVED**

All five god interfaces were decomposed into focused sub-interfaces using
interface embedding. These interfaces (along with their sub-interfaces) have
since been deleted entirely as part of §2.1 (producer-side interface removal).

### 2.4 Double-I naming — **RESOLVED**

The `I` prefix has been removed, resolving the stutter:

- `IIamGateway` → `IAMGateway`
- `IIamGatewayV2` → `IAMGatewayV2`

All `Iam` identifiers have been renamed to `IAM` (e.g., `IamEndpoint` → `IAMEndpoint`,
`IamErrorResponse` → `IAMErrorResponse`, `ErrCatIam` → `ErrCatIAM`). String values
in error codes (e.g., `"VngCloudIamAuthenticationFailed"`) are unchanged.

### 2.5 Empty interface declaration — **RESOLVED**

`VBackUpGateway any` stub deleted from `gateway.go`, along with the corresponding
`Client` interface method, struct field, and getter in `client/client.go`.

---

## 3. File Organization

### 3.1 `i`-prefixed filenames — **RESOLVED**

All 26 `i`-prefixed files have been eliminated. Interfaces were merged into their
corresponding implementation files:
- 5 core infrastructure files (`iclient.go`, `iservice_client.go`, `igateway.go`, `isdk_error.go`)
- 9 service-level files (`icompute.go`, `idns.go`, `iloadbalancer.go`, etc.)
- 12 version-specific `irequest.go` files (distributed to files containing their implementation structs)

### 3.2 Horizontal separator comments — **RESOLVED**

All ~113 decorative `// ---...` separator lines have been removed across 24 files.

---

## 4. Error Handling

### 4.1 Custom error framework — **RESOLVED (phases 1–3)**

The `sdkerror` package implements a bespoke error system with named error codes
(`EcVServerWanIDNotFound`, `EcInternalServerError`, ...) and error categories.

**Phase 1:** `SdkError` bridges to the stdlib `error` interface:

- `Error() string` on `*SdkError` (delegates to `ErrorMessages()`)
- `Unwrap() error` on `*SdkError` (returns wrapped error for `errors.Is()`/`errors.As()` traversal)
- `Is(error) bool` on `*SdkError` (matches by `ErrorCode` when comparing two `SdkError` values)

**Phase 2:** All public API surfaces return `error` instead of `sdkerror.Error`:

- `SdkErrorHandler` widened to accept `error` (backward-compatible)
- 137 service method signatures across 9 interface files migrated
- 134 implementations across 25 files updated
- HTTP client layer (`ServiceClient`, `HTTPClient`, `DoRequest`, `WithReauthFunc`) returns `error`
- Tests updated to use `errors.As(err, &sdkErr)` for rich error inspection

**Phase 3:** Functional-option error handlers replaced with table-driven classifier:

- Deleted the `Error` interface (23 methods, single implementation) — `*SdkError`
  is returned directly from all chaining methods and `SdkErrorHandler`
- Deleted ~96 `WithError*` closure-factory functions
- Replaced with a classifier registry: each `ErrorCode` registers a `classifier`
  struct with a match function, optional category, and optional message format
- `SdkErrorHandler` now takes `...ErrorCode` instead of `...func(Error)` —
  call sites pass error code constants directly
- 5 `New*` constructors added for non-errResp error cases (`NewUnexpectedError`,
  `NewReauthFuncNotSet`, `NewInternalServerError`, `NewServiceMaintenance`,
  `NewPermissionDenied`, `NewQuotaNotFound`)

Callers use standard Go error handling: `errors.Is()`/`errors.As()`,
`fmt.Errorf("...: %w", err)`, and type-switch on `*sdkerror.SdkError`.

**Go convention:** Implement the `error` interface. Use sentinel errors or typed
errors with `errors.Is()` / `errors.As()`. Wrap context with `%w`.

### 4.2 Functional-option error handlers — **RESOLVED**

Replaced by table-driven classifier registry (see §4.1 phase 3).

---

## 5. Structural Patterns

### 5.1 Builder pattern returning interfaces — **RESOLVED**

All request-type constructors (~141), service-factory constructors (14),
gateway constructors (7 + 9 sub-gateways), and client constructors
(`NewClient`, `NewSdkConfigure`) now return `*ConcreteType`. Gateway and
client interfaces have been deleted; their concrete structs are exported.

The four remaining internal interfaces (`ServiceClient`, `HTTPClient`,
`Request`, `SdkAuthentication`) have now been concretized as well (§5.8).
All constructors and builder methods return `*ConcreteStruct`.

### 5.2 100% pointer receivers — **PARTIALLY RESOLVED**

Read-only entity types (16 types, ~46 methods) and error response types (4
types, 8 methods) now use value receivers: `Endpoint`, `ListEndpoints`,
`AddressPair`, `ListAddressPairs`, `VirtualAddress`, `Listener`, `Pool`,
`HealthMonitor`, `LoadBalancer`, `Server`, `Volume`, `ListVolumes`,
`ListQuotas`, `ListPortals`, `ListSecgroupRules`, `AccessToken`,
`IAMErrorResponse`, `NormalErrorResponse`, `NetworkGatewayErrorResponse`,
`GlobalLoadBalancerErrorResponse`.

**Remaining:** Types with mutating methods (`ListServerGroups`,
`ListListeners`, `ListLoadBalancers`, etc.), all client/SDK types (builder
pattern, sync primitives), and request types (builder pattern with `With*`
mutations) correctly keep pointer receivers.

### 5.3 `interface{}` instead of `any` — **RESOLVED**

All ~411 `interface{}` occurrences across ~47 files have been replaced with `any`
using `gofmt -r 'interface{} -> any'`.

### 5.4 Overuse of `var _` compile-time assertions — **RESOLVED**

~33 request-type assertions were removed along with their one-to-one interfaces
(§2.2). 9 gateway assertions were removed when gateway interfaces were deleted
(§5.1). The 3 `IBulkActionRequest` assertions in `glb_pool_requests.go` were
removed when `IBulkActionRequest` was deleted (§5.7). No assertions remain.

### 5.5 Missing godoc

| Metric | Value |
|--------|-------|
| Exported symbols | ~535 |
| Documented | ~54 (~10%) |

90% of exported types and functions have no documentation comments. Public API
surface should have at minimum a one-line summary for each exported type,
function, and method.

### 5.6 `ToRequestBody()` boilerplate — **RESOLVED**

54 request types implemented `ToRequestBody() any` as identity returns
(`return r`). These were deleted and call sites changed from
`WithJSONBody(opts.ToRequestBody())` to `WithJSONBody(opts)`.

10 methods with actual cleanup logic (health monitor field clearing, listener
certificate clearing, nested delegation in loadbalancer/glb) were refactored to
unexported `prepare()` methods called from the service method before
`WithJSONBody(opts)`.

### 5.7 `ToMap()` boilerplate — **RESOLVED**

72 hand-written `ToMap() map[string]any` methods (manually listing every field)
were replaced with a generic helper in `greennode/services/common/`:

```go
func StructToMap(v any) map[string]any {
    b, _ := json.Marshal(v)
    var m map[string]any
    _ = json.Unmarshal(b, &m)
    return m
}
```

Call sites changed from `WithParameters(opts.ToMap())` to
`WithParameters(common.StructToMap(opts))`. All 72 `ToMap()` methods deleted.

The `IBulkActionRequest` interface (which only declared `ToMap()`) was also
deleted; its `BulkActions` field now uses `[]any`.

### 5.8 `UserAgent` per-request embedding — **RESOLVED**

Every request struct embedded `common.UserAgent` and every service method
manually called `WithHeader("User-Agent", opts.ParseUserAgent())`. Replaced
with a single `WithKvDefaultHeaders("User-Agent", c.userAgent)` call in
`client/client.go` during `Configure()`.

Deleted: `common.UserAgent` struct and `ParseUserAgent()` method, ~121
`common.UserAgent` embeddings in request structs, ~97 `AddUserAgent()`
forwarding methods, ~113 `WithHeader("User-Agent", ...)` lines in service
methods.

### 5.9 Missing `context.Context` — **RESOLVED**

All ~137 service methods and the client infrastructure now accept
`ctx context.Context` as their first parameter. The stored `context` field was
removed from both `Client` and `HTTPClient` structs.

Context is threaded through all internal HTTP client methods: `DoRequest`,
`prepareRequest`, `executeHTTPMethod`, `handleReauthBeforeRequest`,
`handleResponse`, `handleStatusCode`, `handleUnauthorized`, `reauthenticate`.

The `reauthFunc` signature changed from `func() (*SdkAuthentication, error)` to
`func(ctx context.Context) (*SdkAuthentication, error)`.

### 5.10 Single-implementation interfaces (client package) — **RESOLVED**

Four interfaces in `greennode/client/` had exactly one concrete implementation
each. All four were deleted and replaced with exported concrete structs:

| Interface | Concrete struct | Files changed |
|-----------|----------------|---------------|
| `Request` | `*Request` | 3 (client pkg) |
| `SdkAuthentication` | `*SdkAuthentication` | 4 (client + entity) |
| `HTTPClient` | `*HTTPClient` | 4 (client + gateway) |
| `ServiceClient` | `*ServiceClient` | 43 (client, gateway, all service base.go + url.go) |

All constructors and builder methods now return `*ConcreteStruct`.

---

## 6. Potential Bugs

### 6.1 V2/V1 return-type mismatch — **RESOLVED**

Added `vnetworkGatewayV2` struct and `NewVNetworkGatewayV2` constructor.
`VNetworkGateway.V2()` now correctly returns `VNetworkGatewayV2` (backed by
`NetworkServiceV2`) instead of `VNetworkGatewayV1`.

---

## Summary

| Category | Items | Scope | Status |
|----------|-------|-------|--------|
| `I`-prefix interfaces | 184 interfaces | 34 files | **Done** |
| `p`-prefix parameters | ~1,456 occurrences | 140 files | **Done** |
| `s` receiver name | ~967 methods | 86 files | **Done** |
| Acronym casing (`Id`, `Json`, `Http`) | ~284 identifiers | codebase-wide | **Done** |
| Abbreviation casing (secondary: `Ip`, `Http1`, `Vserver`) | 19 files | loadbalancer, glb, network, entity | **Done** |
| Java-style `Get*()` accessors | ~162 methods | codebase-wide | **Partial** (4 kept due to collisions) |
| Underscore package names | 1 package | `sdkerror` | **Done** |
| Producer-side interfaces | all interfaces | codebase-wide | **Done** |
| Interface-per-type | all request types | codebase-wide | **Done** |
| God interfaces (>10 methods) | 5 interfaces | 3 packages | **Done** (deleted with §2.1) |
| `i`-prefixed filenames | 26 files | codebase-wide | **Done** |
| Horizontal separators | ~113 occurrences | 24 files | **Done** |
| Custom error framework | 1 package | `sdkerror` | **Done** (phases 1–3: `error` bridge, return types, classifier) |
| Constructors returning interfaces | ~33 functions | gateways, clients | **Done** |
| Pointer receivers on read-only types | 20 types, ~54 methods | entity, sdkerror | **Partial** (entity + error types done) |
| `interface{}` → `any` | ~411 occurrences | ~47 files | **Done** |
| `var _` assertions | ~45 | codebase-wide | **Done** (0 remain; `IBulkActionRequest` deleted) |
| Missing godoc | ~90% of exports | codebase-wide | Open |
| V2/V1 mismatch bug | 1 | `gateway/gateway.go` | **Done** |
| `ToRequestBody()` boilerplate | 54 methods | request + service files | **Done** |
| `ToMap()` boilerplate | 72 methods | request + service files | **Done** |
| `UserAgent` per-request embedding | ~218 embed+methods | request + service files | **Done** |
| Missing `context.Context` | ~137 methods | service + client files | **Done** |
| Single-impl interfaces (client pkg) | 4 interfaces | 43 files | **Done** |
