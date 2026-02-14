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

**Remaining:** ~106 `I*Request` interfaces (e.g., `ICreateLoadBalancerRequest`)
are kept — they have name collisions with their concrete structs. These should be
removed entirely per §2.2 in a future pass.

### 1.2 `p` prefix on parameters — **RESOLVED**

All function parameters have been stripped of the `p` prefix (~1,456 occurrences
across 140 files). Natural words starting with `p` (e.g., `portalUserID`,
`poolID`, `path`, `policies`) were correctly excluded.

**Known residuals:** `psdkErr` in function-type parameters within
`sdkerror/common.go` and `perrResp` in `test/error_test.go` were not caught by
the AST tool (they appear in function types, not direct function signatures).

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

### 2.1 Producer-side interfaces

Every interface in this codebase is defined next to its implementation (the
"producer" side), Java style. Files like `iclient.go`, `igateway.go`,
`irequest.go` sit alongside the concrete types they describe.

**Go convention:** Define interfaces at the call site (the "consumer" side).
The producer exports concrete types; consumers define the small interfaces they
actually need. This yields naturally small interfaces and decouples packages.

**Reference:** [Go wiki: Accept interfaces, return structs](https://go.dev/wiki/CodeReviewComments#interfaces).

### 2.2 Interface-per-type (one-to-one interface/struct pairs)

Every request struct has a parallel `I*Request` interface, even when there is
exactly one implementation. Examples:

- `ICreateLoadBalancerRequest` ↔ `CreateLoadBalancerRequest`
- `IGetPoolByIDRequest` ↔ `GetPoolByIDRequest`
- `IDeleteLoadBalancerByIDRequest` ↔ `DeleteLoadBalancerByIDRequest`

Each pair also has a `var _` compile-time assertion (~45 total).

**Go convention:** Interfaces with a single implementation add indirection
without value. Export the concrete struct directly and let consumers define
interfaces if they need to mock it.

### 2.3 God interfaces (>10 methods)

| Interface | Methods | File |
|-----------|---------|------|
| `LoadBalancerServiceV2` | 35 | `greennode/services/loadbalancer/iloadbalancer.go` |
| `NetworkServiceV2` | 28 | `greennode/services/network/inetwork.go` |
| `GLBServiceV1` | 21 | `greennode/services/glb/iloadbalancer.go` |
| `VDnsServiceInternal` | 11 | `greennode/services/dns/idns.go` |
| `VDnsServiceV1` | 11 | `greennode/services/dns/idns.go` |

**Go convention:** Keep interfaces small and composable. A 35-method interface
is impossible to mock, hard to implement, and signals that the type is doing too
much. Break into focused interfaces (e.g., `PoolCreator`, `ListenerManager`)
or — better — let consumers define the subset they need (see §2.1).

### 2.4 Double-I naming — **RESOLVED**

The `I` prefix has been removed, resolving the stutter:

- `IIamGateway` → `IamGateway`
- `IIamGatewayV2` → `IamGatewayV2`

Note: `Iam` was not further changed to `IAM` (that would be a separate acronym
casing fix beyond scope).

### 2.5 Empty interface declaration

```go
// greennode/gateway/igateway.go:80
type VBackUpGateway interface{}
```

This is an unused stub. It should be deleted or, if backup support is planned,
replaced with a concrete TODO tracked in an issue.

---

## 3. File Organization

### 3.1 `i`-prefixed filenames

34 files use an `i` prefix to indicate "interface definition":

```
client/iclient.go
greennode/client/iclient.go
greennode/client/iservice_client.go
greennode/gateway/igateway.go
greennode/sdkerror/isdk_error.go
greennode/services/loadbalancer/iloadbalancer.go
greennode/services/network/inetwork.go
greennode/services/glb/iloadbalancer.go
greennode/services/dns/idns.go
greennode/services/portal/iportal.go
greennode/services/server/iserver.go
greennode/services/volume/ivolume.go
greennode/services/compute/icompute.go
greennode/services/identity/iidentity.go
  ... plus ~19 irequest.go files in v1/v2/inter subdirectories
```

**Go convention:** Go does not conventionally separate interfaces into dedicated
files. When producer-side interfaces are eliminated (§2.1), these files can be
removed entirely. Any types worth keeping move into the main source file for
their package.

### 3.2 Horizontal separator comments

~119 occurrences of decorative separator lines:

```go
// -----------------------------------------------------------------------
```

Found across ~23+ files, concentrated in:

- `greennode/services/glb/v1/` (56 instances across 5 files)
- `greennode/services/loadbalancer/v2/` (multiple files)
- `test/lb_global_test.go`

**Go convention:** Use godoc section comments or blank lines for visual
grouping. Horizontal rules add noise and no semantic value.

---

## 4. Error Handling

### 4.1 Custom error framework

The `sdkerror` package implements a bespoke error system:

- `Error` interface with `ErrorCode()`, `GetMessage()`, `Err()`, etc.
- `SdkErrorHandler` with functional-option error handlers
- Named error codes (`EcVServerWanIDNotFound`, `EcInternalServerError`, ...)
- Error categories for classification

None of this integrates with the standard library's `errors.Is()`, `errors.As()`,
or `fmt.Errorf("...: %w", err)` wrapping.

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

### 5.1 Builder pattern returning interfaces

~30 `New*` constructors return interface types instead of concrete types:

```go
// greennode/gateway/gateway.go
func NewIamGateway(...) IamGateway          { return &iamGateway{...} }
func NewVServerGateway(...) VServerGateway  { return &vserverGateway{...} }
func NewVLBGateway(...) VLBGateway          { return &vlbGateway{...} }
func NewVNetworkGateway(...) VNetworkGateway { return &vnetworkGateway{...} }

// client/client.go
func NewClient(...) Client                  { return &client{...} }
func NewSdkConfigure() SdkConfigure         { return &sdkConfigure{...} }

// greennode/client/request.go
func NewRequest() Request                   { return &request{...} }
```

**Go convention:** "Accept interfaces, return structs." Constructors should
return the concrete `*T` type. This preserves full type information, allows
consumers to define their own interfaces, and avoids unnecessary heap
allocations through interface boxing.

### 5.2 100% pointer receivers

Every method in the codebase uses a pointer receiver, including simple read-only
accessors on small types (`Endpoint`, `Pool`, `VirtualAddress`).

**Go convention:** Value receivers are appropriate for small, immutable types
and simple accessors. Pointer receivers are for methods that mutate state or
when the struct is large. Mixing receiver types on a single type is discouraged,
so this is a judgment call per type, but the blanket use of pointer receivers
everywhere is not idiomatic.

### 5.3 `interface{}` instead of `any`

| Metric | Value |
|--------|-------|
| `interface{}` occurrences | ~411 |
| Files affected | ~47 |

Since Go 1.18, `any` is a built-in alias for `interface{}`. All 411 occurrences
can be mechanically replaced:

```go
// Before
jsonBody interface{}

// After
jsonBody any
```

### 5.4 Overuse of `var _` compile-time assertions

~45 assertions like:

```go
var _ IListCertificatesRequest = &ListCertificatesRequest{}
```

These are valid when you need to guarantee interface satisfaction at compile time,
but most of these guard one-to-one interface/struct pairs (§2.2). When the
unnecessary interfaces are removed, the assertions go with them.

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

### 6.1 V2/V1 return-type mismatch

```go
// greennode/gateway/gateway.go:164
func (g *vnetworkGateway) V2() VNetworkGatewayV1 {
    return g.vnetworkGatewayV2
}
```

The `V2()` method declares return type `VNetworkGatewayV1`. The interface
definition in `igateway.go:33` mirrors this:

```go
V2() VNetworkGatewayV1
```

This likely means callers of `V2()` get the V1 API surface. Either the method
should return `VNetworkGatewayV2`, or the naming is misleading.

---

## Summary

| Category | Items | Scope | Status |
|----------|-------|-------|--------|
| `I`-prefix interfaces | 184 interfaces | 34 files | **Done** (non-request); ~106 `I*Request` remain |
| `p`-prefix parameters | ~1,456 occurrences | 140 files | **Done** (2 residuals in func types) |
| `s` receiver name | ~967 methods | 86 files | **Done** |
| Acronym casing (`Id`, `Json`, `Http`) | ~284 identifiers | codebase-wide | **Done** |
| Java-style `Get*()` accessors | ~162 methods | codebase-wide | **Partial** (6 kept due to collisions) |
| Underscore package names | 1 package | `sdkerror` | **Done** |
| Producer-side interfaces | all interfaces | codebase-wide | Open |
| Interface-per-type | all request types | codebase-wide | Open |
| God interfaces (>10 methods) | 5 interfaces | 3 packages | Open |
| `i`-prefixed filenames | 34 files | codebase-wide | Open |
| Horizontal separators | ~119 occurrences | ~23 files | Open |
| Custom error framework | 1 package | `sdkerror` | Open |
| Constructors returning interfaces | ~30 functions | gateways, clients | Open |
| `interface{}` → `any` | ~411 occurrences | ~47 files | Open |
| `var _` assertions | ~45 | codebase-wide | Open |
| Missing godoc | ~97% of exports | codebase-wide | Open |
| V2/V1 mismatch bug | 1 | `gateway/gateway.go` | Open |
