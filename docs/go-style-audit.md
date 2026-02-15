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
One exception: `IBulkActionRequest` (3 implementations, genuine polymorphism).

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

### 1.5 Java-style getters — **PARTIALLY RESOLVED**

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
- `GetProjectID()`, `GetZoneID()`, `GetUserID()`, `GetClientID()`, `GetClientSecret()`
- `GetMessage()` (collides with `Message` field on error response structs)

### 1.6 Package names with underscores — **RESOLVED**

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
Method return types (`With*()`, `AddUserAgent()`) now return `*ConcreteType`
instead of `I*Interface`. URL builders, service implementations, and parent
service interfaces all accept concrete pointer types directly.

**Exception kept:** `IBulkActionRequest` in `glb/v1/glb_pool_requests.go` has
three implementations (`PatchGlobalPoolCreateBulkActionRequest`,
`PatchGlobalPoolDeleteBulkActionRequest`,
`PatchGlobalPoolUpdateBulkActionRequest`) — genuine polymorphism.

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

### 4.1 Custom error framework — **PARTIALLY RESOLVED (phases 1 & 2)**

The `sdkerror` package implements a bespoke error system:

- `Error` interface with `ErrorCode()`, `GetMessage()`, `Err()`, etc.
- `SdkErrorHandler` with functional-option error handlers
- Named error codes (`EcVServerWanIDNotFound`, `EcInternalServerError`, ...)
- Error categories for classification

**Phase 1 done:** `SdkError` now bridges to the stdlib `error` interface:

- `Error() string` added to the `Error` interface and `*SdkError` (delegates to `ErrorMessages()`)
- `Unwrap() error` added to `*SdkError` (returns wrapped error for `errors.Is()`/`errors.As()` traversal)
- `Is(error) bool` added to `*SdkError` (matches by `ErrorCode` when comparing two `SdkError` values)

**Phase 2 done:** All public API surfaces now return `error` instead of `sdkerror.Error`:

- `SdkErrorHandler` widened to accept `error` (backward-compatible)
- 137 service method signatures across 9 interface files migrated
- 134 implementations across 25 files updated
- HTTP client layer (`ServiceClient`, `HTTPClient`, `DoRequest`, `WithReauthFunc`) returns `error`
- Tests updated to use `errors.As(err, &sdkErr)` for rich error inspection

Callers can now use standard Go error handling: `errors.Is()`/`errors.As()`,
`fmt.Errorf("...: %w", err)`, and type-switch on `*sdkerror.SdkError`.

**Remaining:** The functional-option error handler pattern (§4.2) is deferred to a
future phase.

**Go convention:** Implement the `error` interface. Use sentinel errors or typed
errors with `errors.Is()` / `errors.As()`. Wrap context with `%w`. This lets
callers use the standard toolchain to inspect errors.

### 4.2 Functional-option error handlers

```go
WithErrorServerNotFound(errResp ErrorResponse) func(sdkError Error) // and many similar
```

Error matching is done by registering handler functions, adding indirection that
makes control flow hard to follow.

**Go convention:** Return errors directly. Let the caller switch on error type
or use `errors.Is()`.

---

## 5. Structural Patterns

### 5.1 Builder pattern returning interfaces — **RESOLVED**

All request-type constructors (~141), service-factory constructors (14),
gateway constructors (7 + 9 sub-gateways), and client constructors
(`NewClient`, `NewSdkConfigure`) now return `*ConcreteType`. Gateway and
client interfaces have been deleted; their concrete structs are exported.

**Deferred:** `ServiceClient`, `HTTPClient`, `Request`, `SdkAuthentication`
(internal types used as parameters/fields in every service and gateway file —
changing them requires touching 50+ files for minimal user-facing benefit).

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
(§5.1). Only 3 `IBulkActionRequest` assertions remain in
`glb_pool_requests.go` — these are legitimate multi-implementation cases.

### 5.5 Missing godoc

| Metric | Value |
|--------|-------|
| Exported symbols | ~1,046 |
| Documented | ~30 (~3%) |

97% of exported types and functions have no documentation comments. Public API
surface should have at minimum a one-line summary for each exported type,
function, and method.

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
| `I`-prefix interfaces | 184 interfaces | 34 files | **Done** (1 exception: `IBulkActionRequest`) |
| `p`-prefix parameters | ~1,456 occurrences | 140 files | **Done** |
| `s` receiver name | ~967 methods | 86 files | **Done** |
| Acronym casing (`Id`, `Json`, `Http`) | ~284 identifiers | codebase-wide | **Done** |
| Java-style `Get*()` accessors | ~162 methods | codebase-wide | **Partial** (6 kept due to collisions) |
| Underscore package names | 1 package | `sdkerror` | **Done** |
| Producer-side interfaces | all interfaces | codebase-wide | **Done** |
| Interface-per-type | all request types | codebase-wide | **Done** (1 exception: `IBulkActionRequest`) |
| God interfaces (>10 methods) | 5 interfaces | 3 packages | **Done** (deleted with §2.1) |
| `i`-prefixed filenames | 26 files | codebase-wide | **Done** |
| Horizontal separators | ~113 occurrences | 24 files | **Done** |
| Custom error framework | 1 package | `sdkerror` | **Partial** (phases 1–2: `error` bridge + return types) |
| Constructors returning interfaces | ~33 functions | gateways, clients | **Done** (4 deferred: `ServiceClient`, `HTTPClient`, `Request`, `SdkAuthentication`) |
| Pointer receivers on read-only types | 20 types, ~54 methods | entity, sdkerror | **Partial** (entity + error types done) |
| `interface{}` → `any` | ~411 occurrences | ~47 files | **Done** |
| `var _` assertions | ~45 | codebase-wide | **Done** (3 remain: `IBulkActionRequest`) |
| Missing godoc | ~97% of exports | codebase-wide | Open |
| V2/V1 mismatch bug | 1 | `gateway/gateway.go` | **Done** |
