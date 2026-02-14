# Go Style Audit

This document catalogs every non-idiomatic Go pattern inherited from the upstream
Java/C#-style codebase. It serves as the roadmap for incremental refactoring toward
idiomatic Go.

Counts were gathered programmatically and may shift as the codebase evolves; treat
them as order-of-magnitude guides, not exact totals.

---

## 1. Naming Conventions

### 1.1 `I` prefix on interfaces

| Metric | Value |
|--------|-------|
| Interfaces affected | ~184 |
| Files | 34 `i`-prefixed files plus inline declarations |

Every interface in the codebase uses a Java/C#-style `I` prefix:
`IClient`, `IHttpClient`, `IServiceClient`, `ILoadBalancerServiceV2`, `IError`, etc.

**Go convention:** Name interfaces after what they do — `Reader`, `Writer`,
`Closer` — or after the concept they represent. Drop the `I` prefix entirely.

**Examples:**

| Current | Idiomatic Go |
|---------|-------------|
| `IClient` | `Client` |
| `IHttpClient` | `HTTPClient` |
| `ILoadBalancerServiceV2` | `LoadBalancerService` (or decompose — see §2.3) |
| `IError` | `Error` |
| `ICreateLoadBalancerRequest` | remove entirely (see §2.2) |

### 1.2 `p` prefix on parameters

| Metric | Value |
|--------|-------|
| Files affected | ~120 |
| Total occurrences | ~1,009 |

Every function parameter is prefixed with `p`: `pOpts`, `pId`, `pProjectId`,
`pClientId`, `pEndpoint`, etc.

**Go convention:** Use short, descriptive names without prefixes. Single-letter
names are fine for small scopes.

```go
// Before
func (s *client) WithHttpClient(pclient svcclient.IHttpClient) IClient {

// After
func (c *Client) WithHTTPClient(hc HTTPClient) *Client {
```

### 1.3 `s` receiver name on all types

| Metric | Value |
|--------|-------|
| Methods affected | ~1,009 |

Every method in the codebase uses `s` as the receiver regardless of type:

```go
func (s *sdkConfigure) GetClientId() string { ... }
func (s *Endpoint) GetId() string { ... }
func (s *request) WithOkCodes(pokCodes ...int) IRequest { ... }
```

**Go convention:** Use a one- or two-letter abbreviation of the type name,
consistently within each type.

| Type | Receiver |
|------|----------|
| `Client` | `c` |
| `LoadBalancer` | `lb` |
| `Endpoint` | `e` |
| `Request` | `r` |
| `SdkConfigure` | `sc` |

### 1.4 Acronym casing

Go treats common acronyms as single words in ALL CAPS
([Go wiki: Initialisms](https://go.dev/wiki/CodeReviewComments#initialisms)).

| Pattern | Instances | Fix |
|---------|-----------|-----|
| `Id` → `ID` | ~239 | `GetClientId()` → `ClientID()` |
| `Json` → `JSON` | pervasive | `WithJsonBody()` → `WithJSONBody()` |
| `Http` → `HTTP` | ~2 types + constants | `IHttpClient` → `HTTPClient` |
| `Url` → `URL` | 0 | (already correct or not used) |
| `Api` → `API` | 0 in identifiers | (only in string constants) |

### 1.5 Java-style getters

| Metric | Value |
|--------|-------|
| `Get*()` accessor methods | ~162 |

Go does not prefix simple accessors with `Get`. An exported field or a method
named after the thing it returns is preferred.

```go
// Before
func (s *Endpoint) GetId() string     { return s.id }
func (s *Endpoint) GetName() string   { return s.name }
func (s *Endpoint) GetStatus() string { return s.status }

// After — exported fields (if no invariants to protect)
type Endpoint struct {
    ID     string
    Name   string
    Status string
}

// Or, if accessor methods are needed:
func (e *Endpoint) ID() string     { return e.id }
func (e *Endpoint) Name() string   { return e.name }
func (e *Endpoint) Status() string { return e.status }
```

### 1.6 Package names with underscores

| Package | Location |
|---------|----------|
| `sdk_error` | `greennode/sdk_error/` (18 files) |

**Go convention:** Package names are lowercase, single-word, no underscores.
`sdk_error` → `sdkerror`.

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
- `IGetPoolByIdRequest` ↔ `GetPoolByIdRequest`
- `IDeleteLoadBalancerByIdRequest` ↔ `DeleteLoadBalancerByIdRequest`

Each pair also has a `var _` compile-time assertion (~45 total).

**Go convention:** Interfaces with a single implementation add indirection
without value. Export the concrete struct directly and let consumers define
interfaces if they need to mock it.

### 2.3 God interfaces (>10 methods)

| Interface | Methods | File |
|-----------|---------|------|
| `ILoadBalancerServiceV2` | 35 | `greennode/services/loadbalancer/iloadbalancer.go` |
| `INetworkServiceV2` | 28 | `greennode/services/network/inetwork.go` |
| `IGLBServiceV1` | 21 | `greennode/services/glb/iloadbalancer.go` |
| `IVDnsServiceInternal` | 11 | `greennode/services/dns/idns.go` |
| `IVDnsServiceV1` | 11 | `greennode/services/dns/idns.go` |

**Go convention:** Keep interfaces small and composable. A 35-method interface
is impossible to mock, hard to implement, and signals that the type is doing too
much. Break into focused interfaces (e.g., `PoolCreator`, `ListenerManager`)
or — better — let consumers define the subset they need (see §2.1).

### 2.4 Double-I naming

The `I` prefix collides with the `IAM` acronym, producing stutter names:

- `IIamGateway` (`greennode/gateway/igateway.go`)
- `IIamGatewayV2` (`greennode/gateway/igateway.go`)

Removing the `I` prefix resolves this: `IAMGateway`, `IAMGatewayV2`.

### 2.5 Empty interface declaration

```go
// greennode/gateway/igateway.go:80
type IVBackUpGateway interface{}
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
greennode/sdk_error/isdk_error.go
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

The `sdk_error` package implements a bespoke error system:

- `IError` interface with `GetErrorCode()`, `GetMessage()`, `GetErrors()`, etc.
- `SdkErrorHandler` with functional-option error handlers
- Named error codes (`EcVServerWanIdNotFound`, `EcInternalServerError`, ...)
- Error categories for classification

None of this integrates with the standard library's `errors.Is()`, `errors.As()`,
or `fmt.Errorf("...: %w", err)` wrapping.

**Go convention:** Implement the `error` interface. Use sentinel errors or typed
errors with `errors.Is()` / `errors.As()`. Wrap context with `%w`. This lets
callers use the standard toolchain to inspect errors.

### 4.2 Functional-option error handlers

```go
WithErrorServerNotFound(perrResp *sdkerr.SdkError) // and many similar
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
func NewIamGateway(...) IIamGateway          { return &iamGateway{...} }
func NewVServerGateway(...) IVServerGateway  { return &vserverGateway{...} }
func NewVLBGateway(...) IVLBGateway          { return &vlbGateway{...} }
func NewVNetworkGateway(...) IVNetworkGateway { return &vnetworkGateway{...} }

// client/client.go
func NewClient(...) IClient                  { return &client{...} }
func NewSdkConfigure() ISdkConfigure         { return &sdkConfigure{...} }

// greennode/client/request.go
func NewRequest() IRequest                   { return &request{...} }
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
func (s *vnetworkGateway) V2() IVNetworkGatewayV1 {
    return s.vnetworkGatewayV2
}
```

The `V2()` method declares return type `IVNetworkGatewayV1`. The interface
definition in `igateway.go:37` mirrors this:

```go
V2() IVNetworkGatewayV1
```

And the field is initialized from `NewVNetworkGatewayV1(vnetworkSvcV2)` at
`gateway.go:124`.

This likely means callers of `V2()` get the V1 API surface. Either the method
should return `IVNetworkGatewayV2`, or the naming is misleading.

---

## Summary

| Category | Items | Scope |
|----------|-------|-------|
| `I`-prefix interfaces | 184 interfaces | 34 files |
| `p`-prefix parameters | ~1,009 occurrences | ~120 files |
| `s` receiver name | ~1,009 methods | codebase-wide |
| Acronym casing (`Id`, `Json`, `Http`) | ~241+ identifiers | codebase-wide |
| Java-style `Get*()` accessors | ~162 methods | codebase-wide |
| Underscore package names | 1 package | `sdk_error` |
| Producer-side interfaces | all interfaces | codebase-wide |
| Interface-per-type | all request types | codebase-wide |
| God interfaces (>10 methods) | 5 interfaces | 3 packages |
| `i`-prefixed filenames | 34 files | codebase-wide |
| Horizontal separators | ~119 occurrences | ~23 files |
| Custom error framework | 1 package | `sdk_error` |
| Constructors returning interfaces | ~30 functions | gateways, clients |
| `interface{}` → `any` | ~411 occurrences | ~47 files |
| `var _` assertions | ~45 | codebase-wide |
| Missing godoc | ~97% of exports | codebase-wide |
| V2/V1 mismatch bug | 1 | `gateway/gateway.go` |
