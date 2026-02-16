# Test Coverage

This document catalogs the current state of test coverage across the
`greennode-community-sdk` project. It identifies gaps, remaining issues, and
serves as the roadmap for improving test quality.

---

## 1. Coverage Map

### 1.1 Integration tests

Integration tests live in `test/` and call real VNGCloud APIs (181 tests
across 22 files).

| Service                | Test File(s)                                                                                                                       | Tests |
| ---------------------- | ---------------------------------------------------------------------------------------------------------------------------------- | ----- |
| identity/v2            | `identity_test.go`                                                                                                                 | 4     |
| loadbalancer/v2        | `lb_test.go`                                                                                                                       | 39    |
| loadbalancer/inter     | `lb_test.go`                                                                                                                       | (included above) |
| loadbalancer/v2 certs  | `certificate_test.go`                                                                                                              | 4     |
| glb/v1                 | `glb_test.go`, `lb_global_test.go`                                                                                                 | 22    |
| compute/v2             | `server_test.go`                                                                                                                   | 24    |
| volume/v2              | `volume_test.go`, `snapshot_test.go`                                                                                               | 17    |
| volume/v1              | `volumetype_test.go`                                                                                                               | 5     |
| network/v2             | `network_test.go`, `secgroup_test.go`, `secgroup_rule_test.go`, `subnet_test.go`, `virtualaddress_test.go`, `address_pair_test.go` | 26    |
| network/v1             | `endpoint_test.go`                                                                                                                 | 10    |
| portal/v1, v2          | `portal_test.go`                                                                                                                   | 10    |
| dns/v1                 | `dns_test.go`                                                                                                                      | 12    |
| dns/internal_system/v1 | `dns_internal_test.go`                                                                                                             | 6     |
| server/v1              | `server_test.go`                                                                                                                   | (included in compute above) |

### 1.2 Unit tests

Unit tests live alongside the packages they test (139 tests across 14 files).

| Package                    | Test File(s)                                          | Tests |
| -------------------------- | ----------------------------------------------------- | ----- |
| `greennode/sdkerror/`      | `sdk_error_test.go`, `classifier_test.go`, `errors_test.go` | 61    |
| `greennode/client/`        | `auth_test.go`, `request_test.go`, `service_client_test.go` | 26    |
| `greennode/entity/`        | `entity_test.go`                                      | 16    |
| `greennode/services/volume/v2/`   | `blockvolume_request_test.go`, `blockvolume_response_test.go` | 16    |
| `greennode/services/common/`      | `common_test.go`                                      | 7     |
| `greennode/services/network/v2/`  | `secgroup_request_test.go`, `secgroup_response_test.go` | 8     |
| `greennode/services/identity/v2/` | `identity_request_test.go`, `identity_response_test.go` | 5     |

### 1.3 Packages with zero tests

| Package              | Description                     |
| -------------------- | ------------------------------- |
| `greennode/auth/`     | IAM user auth flow and TOTP providers |
| `greennode/gateway/`  | Gateway routing layer (consolidated into `greennode.go`) |

### 1.4 Utility test files

| File            | Purpose                                              |
| --------------- | ---------------------------------------------------- |
| `data_test.go`  | Fake certs/keys constants for test use               |
| `error_test.go` | 1 test for error message format                      |
| `other_test.go` | `TestUnixNano` — unrelated to SDK, should be removed |

---

## 2. Remaining Issues

### 2.1 Hardcoded resource IDs — **Open**

Every integration test uses hardcoded UUIDs for cloud resources (e.g.,
`"lb-8f54cbd4-b8ee-4b86-aa9b-d365c468a902"`), making tests non-reproducible
across environments and fragile if resources are deleted.

### 2.2 No build tags on integration tests — **Open**

All 22 test files run against live APIs but lack `//go:build integration` tags.
Running `go test ./test/...` always attempts live API calls. Unit tests cannot
be run in isolation from integration tests.

### 2.3 `TestGetLoadBalancerFailure` asserts success — **Open**

Named as a failure test but asserts success conditions (error must be nil,
result must not be nil).

### 2.4 `TestUnixNano` noise — **Open**

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

### 3.6 No unit tests — **RESOLVED**

141 unit tests added across `sdkerror/`, `client/`, `entity/`,
`services/common/`, `services/volume/v2/`, `services/network/v2/`, and
`services/identity/v2/`.

---

## 4. Improvement Plan

### Phase 2: Expand unit test coverage — **Open**

Existing unit tests cover core packages. Remaining gaps:
- `auth/`: TOTP computation (SecretTOTP against RFC 6238 test vectors), endpoint resolution
- `services/compute/v2/`: request building, response parsing
- `services/loadbalancer/v2/`: request building, response parsing
- `services/dns/v1/`: request building, response parsing
- `services/glb/v1/`: request building, response parsing
- `services/network/v1/`: request building, response parsing
- `services/portal/`: request building, response parsing

### Phase 3: Integration test infrastructure — **Open**

- Tag all integration tests with `//go:build integration`
- Add HTTP mocking with `httptest.Server` or fixture files
