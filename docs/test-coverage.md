# Test Coverage

This document catalogs the current state of test coverage across the
`greennode-community-sdk` project. It identifies gaps, remaining issues, and
serves as the roadmap for improving test quality.

---

## 1. Coverage Map

### 1.1 Services with tests

All existing tests are **live integration tests** that call real VNGCloud APIs.
There are currently **zero unit tests**.

| Service                | Test File(s)                                                                                                                       | Tests |
| ---------------------- | ---------------------------------------------------------------------------------------------------------------------------------- | ----- |
| identity/v2            | `identity_test.go`                                                                                                                 | 4     |
| loadbalancer/v2        | `lb_test.go`                                                                                                                       | 30    |
| loadbalancer/inter     | `lb_test.go`                                                                                                                       | 6     |
| glb/v1                 | `glb_test.go`, `lb_global_test.go`                                                                                                 | 22    |
| compute/v2             | `server_test.go`                                                                                                                   | 16    |
| volume/v2              | `volume_test.go`, `snapshot_test.go`                                                                                               | 15    |
| volume/v1              | `volumetype_test.go`                                                                                                               | 5     |
| network/v2             | `network_test.go`, `secgroup_test.go`, `secgroup_rule_test.go`, `subnet_test.go`, `virtualaddress_test.go`, `address_pair_test.go` | 20    |
| network/v1             | `endpoint_test.go`                                                                                                                 | 10    |
| portal/v1, v2          | `portal_test.go`                                                                                                                   | 9     |
| dns/v1                 | `dns_test.go`                                                                                                                      | 14    |
| dns/internal_system/v1 | `dns_internal_test.go`                                                                                                             | 5     |
| server/v1              | `server_test.go`                                                                                                                   | 1     |

### 1.2 Packages with zero tests

| Package                      | Description                                                         |
| ---------------------------- | ------------------------------------------------------------------- |
| `greennode/client/`          | HTTP client, request building, service client — core infrastructure |
| `greennode/gateway/`         | All 5 gateway files — routing layer                                 |
| `greennode/sdkerror/`        | 18 error classification files — error handling logic                |
| `greennode/services/common/` | 9 shared utility files                                              |

### 1.3 Utility test files

| File            | Purpose                                              |
| --------------- | ---------------------------------------------------- |
| `data_test.go`  | Fake certs/keys constants for test use               |
| `error_test.go` | 1 test for error message format                      |
| `other_test.go` | `TestUnixNano` — unrelated to SDK, should be removed |

---

## 2. Remaining Issues

### 2.1 No unit tests — **Open**

0% unit test coverage on the `greennode/` package tree. No mocking or stubbing
of HTTP calls. Request building, response parsing, URL construction, and error
classification are all untested in isolation. The `client/`, `gateway/`,
`sdkerror/`, and `services/common/` packages have no tests at all.

### 2.2 Hardcoded resource IDs — **Open**

Every integration test uses hardcoded UUIDs for cloud resources (e.g.,
`"lb-8f54cbd4-b8ee-4b86-aa9b-d365c468a902"`), making tests non-reproducible
across environments and fragile if resources are deleted.

### 2.3 No build tags on integration tests — **Open**

All 22 test files run against live APIs but lack `//go:build integration` tags.
Running `go test ./test/...` always attempts live API calls. Unit tests (once
added) cannot be run in isolation.

### 2.4 `TestGetLoadBalancerFailure` asserts success — **Open**

Named as a failure test but asserts success conditions (error must be nil,
result must not be nil).

### 2.5 `TestUnixNano` noise — **Open**

Unrelated to SDK functionality; should be removed.

---

## 3. Resolved Issues

### 3.1 Developer names in test code — **RESOLVED**

All personal identifiers (cuongdm3, vinhnt8, annd2, hannibal, phongnt10,
tantm3, duynh7, tytv2, User11412) replaced with generic test names across
12 files (~50 replacements). Inappropriate strings removed.

### 3.2 False-passing tests — **RESOLVED**

12 tests used `t.Logf`/`t.Log` instead of `t.Fatalf` for error checks, causing
them to report success when API calls failed. All changed to `t.Fatalf`.

### 3.3 Nil pointer panic in `TestAuthenPass` — **RESOLVED**

`TestAuthenPass` dereferenced `token.Token` without nil check, crashing the
suite. Fixed with proper nil guard and `t.Fatalf`.

### 3.4 Config helper bloat — **RESOLVED**

9 developer-named config functions consolidated into generic helpers. Added
reusable `newClientFromEnvKeys` function.

### 3.5 Test name typo — **RESOLVED**

`TestASuperuthenPass` renamed to `TestSuperAdminAuthenPass`.

---

## 4. Improvement Plan

### Phase 2: Add unit tests — **Open**

- `sdkerror/`: error creation, classification, code matching
- `client/`: request building, URL construction, retry logic
- `services/`: `New*Request()` builders produce correct JSON, `*Response` structs parse correctly
- `entity/`: struct marshal/unmarshal

### Phase 3: Integration test infrastructure — **Open**

- Tag all integration tests with `//go:build integration`
- Add HTTP mocking with `httptest.Server` or fixture files
