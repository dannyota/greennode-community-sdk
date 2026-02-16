# Legacy Patterns Review: Java/C# Style Code

**Date**: 2026-02-16
**Status**: Initial audit — tracking for incremental improvement

---

## Executive Summary

The codebase (forked from `vngcloud-go-sdk`) carries significant Java/C#-style patterns that are not idiomatic Go. The most pervasive issues are:

1. **346 `With*()` builder methods** — mutating setter chains instead of struct literals or functional options
2. **150+ `Get*()` getter methods** — on exported fields that can be accessed directly
3. **17 single-field "Common" wrapper structs** — over-engineered ID holders
4. **71 Request/Response wrapper files** — excessive DTO layer with `ToEntity*()` converters
5. **100+ `New*()` factory functions** — many just do `new(T)` + field assignment

What's actually fine:
- Enum patterns (const + typed strings) are idiomatic Go
- No `I`-prefix interfaces
- No `service/impl` separation — architecture is pragmatic
- Clean dependency graph with no circular imports
- Flat client API (after recent refactor)

---

## Pattern Inventory

### 1. `With*()` Builder Methods — 346 instances

**Severity**: HIGH | **Effort**: LARGE

Every request struct uses Java-style fluent builder:

```go
// Current (Java-style)
req := dns.NewCreateHostedZoneRequest("example.com").
    WithDescription("my zone").
    WithType(dns.HostedZoneTypePrivate).
    WithAssocVpcIDs([]string{"net-abc123"})
```

**Breakdown by area:**

| Area | Count | Location |
|------|-------|----------|
| Service request builders | 318 | `greennode/services/*/` |
| Client/ServiceClient config | 24 | `greennode/client/` |
| SdkError construction | 6 | `greennode/sdkerror/` |

**Top files by density:**

| File | Count |
|------|-------|
| `glb/v1/glb_pool_requests.go` | 57 |
| `dns/v1/dns_record_request.go` | 21 |
| `loadbalancer/v2/pool_requests.go` | 20 |
| `loadbalancer/v2/listener_request.go` | 19 |
| `glb/v1/glb_listener_request.go` | 24 |
| `glb/v1/glb_request.go` | 15 |

**Idiomatic alternatives:**

```go
// Option A: Struct literals (simplest, recommended for most cases)
req := &dns.CreateHostedZoneRequest{
    Name:        "example.com",
    Description: "my zone",
    Type:        dns.HostedZoneTypePrivate,
    AssocVpcIDs: []string{"net-abc123"},
}

// Option B: Functional options (for complex config with many optional fields)
req := dns.NewCreateHostedZoneRequest("example.com",
    dns.WithDescription("my zone"),
    dns.WithType(dns.HostedZoneTypePrivate),
)

// Option C: Constructor + direct field assignment
req := dns.NewCreateHostedZoneRequest("example.com")
req.Description = "my zone"
req.Type = dns.HostedZoneTypePrivate
```

---

### 2. `Get*()` Getter Methods — 150+ instances

**Severity**: HIGH | **Effort**: MEDIUM (breaking API change)

Getter methods on exported fields add no value:

```go
// Current — entity/endpoint.go
func (e Endpoint) GetID() string       { return e.ID }
func (e Endpoint) GetName() string     { return e.Name }
func (e Endpoint) GetIPv4Address() string { return e.IPv4Address }

// Current — entity/virtualaddress.go
func (v VirtualAddress) GetID() string          { return v.ID }
func (v VirtualAddress) GetName() string        { return v.Name }
func (v VirtualAddress) GetVpcID() string       { return v.VpcID }
func (v VirtualAddress) GetSubnetID() string    { return v.SubnetID }
func (v VirtualAddress) GetDescription() string { return v.Description }
// ... and many more

// Current — services/common/loadbalancer.go
func (l *LoadBalancerCommon) GetLoadBalancerID() string { return l.LoadBalancerID }
```

**Locations:**
- `greennode/entity/` — 25+ getters across endpoint, virtualaddress, loadbalancer, listener, pool, server
- `greennode/services/common/` — 26+ getters on Common wrapper structs
- `greennode/client/service_client.go` — 3 getters
- `greennode/sdkerror/` — GetMessage()

**Idiomatic Go**: Just access the exported field directly (`endpoint.ID`, `endpoint.Name`).

---

### 3. Common Wrapper Structs — 17 types

**Severity**: MEDIUM | **Effort**: MEDIUM

Single-field structs in `greennode/services/common/` exist only to hold an ID:

```go
// Current — services/common/loadbalancer.go
type LoadBalancerCommon struct {
    LoadBalancerID string
}
func (l *LoadBalancerCommon) GetLoadBalancerID() string {
    return l.LoadBalancerID
}

// Embedded in request types
type GetLoadBalancerByIDRequest struct {
    common.LoadBalancerCommon
}
```

**All Common wrappers:**
- `LoadBalancerCommon`, `ListenerCommon`, `PoolCommon`, `PolicyCommon`
- `L7RuleCommon`, `CertificateCommon`, `MemberCommon`
- `BlockVolumeCommon`, `SnapshotCommon`
- `NetworkCommon`, `SubnetCommon`, `SecgroupCommon`, `SecgroupRuleCommon`
- `VirtualAddressCommon`, `EndpointCommon`
- `Project`, `PortalUser`

**Idiomatic Go**: Just put the ID field directly on the request struct.

---

### 4. Request/Response DTO Layer — 71 files

**Severity**: MEDIUM | **Effort**: LARGE (architectural)

Every API operation has a separate Request type, Response type, and `ToEntity*()` converter:

```
services/<service>/<version>/
    *_request.go    (32 files, ~145 request types)
    *_response.go   (31 files, ~102 response types)
    service.go      (API method implementations)
    url.go          (URL builders)
```

The response types exist mainly to call `ToEntity*()`:

```go
// Response type just maps JSON to entity
type GetBlockVolumeByIDResponse struct { /* fields */ }

func (r *GetBlockVolumeByIDResponse) ToEntityVolume() *entity.Volume {
    return &entity.Volume{
        ID:   r.ID,
        Name: r.Name,
        // ... field mapping
    }
}
```

**109 `ToEntity*()` methods** across the codebase.

**Consideration**: Some mapping is genuinely needed (API shape != domain shape), but many are 1:1 copies. Could use JSON tags on entity types directly for simple cases.

---

### 5. `New*()` Factory Functions — 100+ instances

**Severity**: LOW-MEDIUM | **Effort**: SMALL

Many factories just do `new(T)` + field assignment:

```go
// Current
func NewCreateBlockVolumeRequest(name, volType string, size int64) *CreateBlockVolumeRequest {
    opt := new(CreateBlockVolumeRequest)
    opt.VolumeTypeID = volType
    opt.CreatedFrom = CreateFromNew
    opt.Name = name
    opt.Size = size
    return opt
}

// Idiomatic Go
func NewCreateBlockVolumeRequest(name, volType string, size int64) *CreateBlockVolumeRequest {
    return &CreateBlockVolumeRequest{
        Name:        name,
        VolumeTypeID: volType,
        CreatedFrom: CreateFromNew,
        Size:        size,
    }
}
```

Not a big deal functionally — mostly a style issue. The `new(T)` + assignment style is verbose but works.

---

### 6. Hidden `prepare()` Methods — 20+ instances

**Severity**: MEDIUM | **Effort**: SMALL

Request types have unexported `prepare()` methods called before sending, which silently mutate fields:

```go
// In service method
func (s *LoadBalancerServiceV2) CreateLoadBalancer(ctx context.Context, opts *CreateLoadBalancerRequest) (*entity.LoadBalancer, error) {
    opts.prepare()  // Hidden mutation!
    // ...
}

// In request type
func (h *HealthMonitor) prepare() {
    if h.DomainName == nil {
        h.DomainName = common.Ptr(defaultFakeDomainName)
    }
}
```

**Problem**: Callers don't know their request objects are being mutated. Could cause subtle bugs if the same request is reused.

**Better pattern**: Apply defaults in the constructor, or use a copy before mutation.

---

### 7. `Set*()` Setter Methods — 7 instances

**Severity**: LOW | **Effort**: SMALL

```go
// greennode/client/request.go
func (r *Request) SetJSONResponse(jsonResponse any) { r.jsonResponse = jsonResponse }
func (r *Request) SetJSONError(jsonError any)       { r.jsonError = jsonError }

// greennode/services/common/common.go
func (pr *Project) SetProjectID(id string)       { pr.ID = id }
func (pu *PortalUser) SetPortalUserID(id string) { pu.PortalUserID = id }
func (p *Paging) SetPage(page int) *Paging       { p.Page = page; return p }
func (p *Paging) SetSize(size int) *Paging       { p.Size = size; return p }
```

Minor issue — small count.

---

## What's Already Good

| Pattern | Assessment |
|---------|-----------|
| Enum types (`const` + typed strings) | Idiomatic Go |
| No `I`-prefix interfaces | Clean |
| No `service/impl` separation | Pragmatic |
| No circular dependencies | Clean dependency graph |
| Flat client API | Recently refactored, good |
| Entity package is pure data | No business logic, minimal deps |
| Service isolation | No cross-service imports |
| Error code system | Well-organized, domain-specific |

---

## Refactoring Priority

### Phase 1: Quick wins (non-breaking internal changes) — DONE
- [x] Replace `new(T)` + assignment with struct literals in `New*()` functions (20 files, ~70 instances)
- [x] Rename hidden `prepare()` to `normalizeForAPI()` with doc comments (v2 + inter packages)
- [x] Remove `Set*()` methods — replaced callers with direct field access or struct literals (7 removed)

### Phase 2: Export field access (breaking change) — DONE
- [x] Remove `Get*()` getters on entity types — users access fields directly (22 removed)
- [x] Remove Common wrapper structs — inline the ID fields into request types (17 inlined, 8 files deleted)
- [x] Remove `Get*()` getters on Common types (Paging getters removed)

### Phase 3: Remove With*() builder methods from service requests — DONE
- [x] Replace `With*()` builder pattern with direct struct field access (318 removed from `greennode/services/`)
- [x] Move validation logic from `With*()` into `normalizeForAPI()` methods
- [x] Add `NewTags()` / `NewServerTags()` / `NewVolumeTags()` helpers for variadic kv-pair convenience
- [x] Export `PolicyPosition` type and `PolicyPositions` field in loadbalancer/v2
- [x] Extract anonymous tag struct to named `SystemTag` in server/v1
- [x] Remove remnant `Get*()` getters on request types (missed in phase 2)
- [x] Update all internal callers and tests

### Phase 4: Reduce DTO overhead (breaking change, optional)
- [ ] Evaluate which `ToEntity*()` conversions are 1:1 copies vs genuine mappings
- [ ] For 1:1 cases, consider using entity types directly as response targets
- [ ] Keep separate response types only where API shape genuinely differs

---

## File Quick Reference

| Area | Path | Key Issue |
|------|------|-----------|
| Entity getters | `greennode/entity/*.go` | 25+ unnecessary `Get*()` |
| Common wrappers | `greennode/services/common/*.go` | 17 wrapper structs + 26 getters |
| Client builders | `greennode/client/service_client.go` | 10 `With*()` methods |
| Client builders | `greennode/client/request.go` | 9 `With*()` + 2 `Set*()` |
| Client builders | `greennode/client/http.go` | 5 `With*()` methods |
| Error builders | `greennode/sdkerror/sdk_error.go` | 6 `With*()` methods |
| GLB requests | `greennode/services/glb/v1/*_requests.go` | 96 `With*()` methods |
| LB requests | `greennode/services/loadbalancer/v2/*_request*.go` | 74 `With*()` methods |
| DNS requests | `greennode/services/dns/v1/*_request.go` | 31 `With*()` methods |
| Network requests | `greennode/services/network/v1/*_request.go` | 23 `With*()` methods |
| Volume requests | `greennode/services/volume/v2/*_request.go` | 17 `With*()` methods |
| Compute requests | `greennode/services/compute/v2/*_request.go` | 16 `With*()` methods |
