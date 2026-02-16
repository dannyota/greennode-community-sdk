# Test Coverage

Current state of testing in the `greennode-community-sdk` project, organized
into unit tests and integration tests.

---

## 1. Unit Tests

Unit tests live alongside the packages they test under `greennode/`. They
require no credentials, no network, and run in milliseconds.

**Run:** `go test ./greennode/...`

### 1.1 Coverage by package

| Package | Coverage | Test files |
| ------- | -------- | ---------- |
| `greennode/sdkerror/` | **94.5%** | `sdk_error_test.go`, `classifier_test.go`, `errors_test.go` |
| `greennode/client/` | **69.7%** | `auth_test.go`, `http_test.go`, `request_test.go`, `service_client_test.go` |
| `greennode/services/common/` | **57.1%** | `common_test.go` |
| `greennode/entity/` | **42.6%** | `entity_test.go` |
| `greennode/services/identity/v2/` | **20.0%** | `identity_request_test.go`, `identity_response_test.go` |
| `greennode/services/volume/v2/` | **18.1%** | `blockvolume_request_test.go`, `blockvolume_response_test.go` |
| `greennode/services/network/v2/` | **4.8%** | `secgroup_request_test.go`, `secgroup_response_test.go` |

### 1.2 Packages with zero unit test coverage

| Package | What it does |
| ------- | ------------ |
| `greennode/` (root) | Client constructor, config, top-level wiring |
| `greennode/auth/` | IAM user auth flow, TOTP provider |
| `greennode/services/compute/v2/` | Server request/response types |
| `greennode/services/loadbalancer/v2/` | Load balancer request/response types |
| `greennode/services/loadbalancer/inter/` | Inter-region LB request/response types |
| `greennode/services/dns/v1/` | DNS request/response types |
| `greennode/services/dns/internal_system/v1/` | DNS internal request/response types |
| `greennode/services/glb/v1/` | Global LB request/response types |
| `greennode/services/network/v1/` | Network v1 request/response types |
| `greennode/services/portal/v1/` | Portal v1 request/response types |
| `greennode/services/portal/v2/` | Portal v2 request/response types |
| `greennode/services/server/v1/` | Server v1 request/response types |
| `greennode/services/volume/v1/` | Volume v1 request/response types |

### 1.3 Future work

Priority order for expanding unit test coverage:

1. **`auth/`** — TOTP computation (test against RFC 6238 vectors), endpoint
   resolution logic
2. **`client/`** — increase from 69.7% to 85%+, cover remaining retry/auth
   edge cases
3. **`entity/`** — increase from 42.6%, cover remaining entity validation
   methods
4. **Service request/response packages** — each has a builder pattern that
   can be tested without any network calls:
   - `services/compute/v2/`
   - `services/loadbalancer/v2/`
   - `services/dns/v1/`
   - `services/glb/v1/`
   - `services/network/v1/`
   - `services/portal/v1/`, `services/portal/v2/`

---

## 2. Integration Tests

Integration tests live in `test/` and make real API calls to GreenNode/VNGCloud
services. They validate end-to-end behavior against live infrastructure.

**Run:** `go test -tags=integration ./test/...`

### 2.1 Test inventory by service

| Service | Test file(s) | Tests |
| ------- | ------------ | ----- |
| Identity / Auth | `identity_test.go` | 4 |
| Load Balancer | `lb_test.go` | 39 |
| LB Certificates | `certificate_test.go` | 4 |
| Global LB | `glb_test.go` | 22 |
| Compute (Servers) | `server_test.go` | 24 |
| Block Volume | `volume_test.go`, `snapshot_test.go` | 17 |
| Volume Types | `volumetype_test.go` | 5 |
| Network / VPC | `network_test.go`, `subnet_test.go` | 10 |
| Security Groups | `secgroup_test.go`, `secgroup_rule_test.go` | 12 |
| Virtual Addresses | `virtualaddress_test.go`, `address_pair_test.go` | 6 |
| Endpoints | `endpoint_test.go` | 10 |
| Portal | `portal_test.go` | 10 |
| DNS | `dns_test.go` | 12 |
| DNS Internal | `dns_internal_test.go` | 6 |

**Support files** (not tests themselves):

| File | Purpose |
| ---- | ------- |
| `helpers_test.go` | Client constructors, env reading, shared utilities |
| `data_test.go` | Fake certs/keys/tokens for test inputs |

### 2.2 Current credential model (env.yaml)

Tests currently read credentials from `test/env.yaml` (gitignored). The file
uses `KEY=VALUE` format:

```yaml
# Service account
VNGCLOUD_CLIENT_ID=<service-account-client-id>
VNGCLOUD_CLIENT_SECRET=<service-account-secret>

# IAM user credentials
USER_CLIENT_ID=<user-client-id>
USER_CLIENT_SECRET=<user-secret>
ALT_USER_CLIENT_ID=<alt-user-client-id>
ALT_USER_CLIENT_SECRET=<alt-user-secret>

# Resource identifiers
VNGCLOUD_USER_ID=<portal-user-id>
VNGCLOUD_ZONE_ID=<zone-id>
VNGCLOUD_PROJECT_ID=<project-id>
USER_PROJECT=<user-project>
ALT_USER_PROJECT_ID=<alt-user-project>
HAN01_PROJECT_ID=<han-region-project>
HCM3B_PROJECT_ID=<hcm3b-region-project>
VNGCLOUD_PORTAL_USER_ID=<portal-user-id>
SECONDARY_USER_ID=<secondary-user-id>
SERVER_ID=<existing-server-id>

# Regional
HCM3B_CLIENT_ID=<hcm3b-client-id>
HCM3B_CLIENT_SECRET=<hcm3b-secret>
```

### 2.3 Target credential model (HashiCorp Vault)

Replace `env.yaml` with HashiCorp Vault for secrets management. This enables:
- No secrets on disk — credentials are fetched at test runtime
- Centralized rotation — update secrets in one place
- Audit trail — Vault logs all secret access
- CI/CD friendly — authenticate via Vault token or AppRole

**Vault secret layout** (proposed):

```
secret/greennode-sdk/test/
  service-account      → client_id, client_secret
  user                 → client_id, client_secret
  alt-user             → client_id, client_secret
  hcm3b                → client_id, client_secret
  resources            → user_id, zone_id, project_id, server_id, ...
  regional             → han01_project_id, hcm3b_project_id, ...
```

**How tests will read credentials:**

1. Tests check for `VAULT_ADDR` and `VAULT_TOKEN` environment variables
2. If set, read credentials from Vault using the Vault Go client
3. If not set, fall back to `env.yaml` for local development
4. A `testconfig` helper package handles the abstraction

```
VAULT_ADDR=https://vault.example.com VAULT_TOKEN=s.xxx go test -tags=integration ./test/...
```

### 2.4 File restructuring

Current state has test helpers mixed into `identity_test.go`. Target structure:

```
test/
  helpers_test.go          # client constructors, config reading, Vault integration
  data_test.go             # fake certs/keys/tokens (no changes)
  identity_test.go         # identity/auth tests only (helpers extracted)
  lb_test.go               # load balancer tests
  lb_global_test.go        # global LB tests
  glb_test.go              # GLB tests
  server_test.go           # compute tests
  volume_test.go           # volume tests
  volumetype_test.go       # volume type tests
  snapshot_test.go         # snapshot tests
  network_test.go          # network/VPC tests
  subnet_test.go           # subnet tests
  secgroup_test.go         # security group tests
  secgroup_rule_test.go    # security group rule tests
  address_pair_test.go     # address pair tests
  virtualaddress_test.go   # virtual address tests
  endpoint_test.go         # endpoint tests
  certificate_test.go      # certificate tests
  portal_test.go           # portal tests
  dns_test.go              # DNS tests
  dns_internal_test.go     # DNS internal tests
```

**Changes completed:**

1. **Extracted helpers from `identity_test.go`** into `helpers_test.go` —
   `readEnvFile()`, `getEnv()`, `getValueOfEnv()`, `newClientFromEnvKeys()`,
   and all `validXxxSdkConfig()` functions
2. **Deleted `other_test.go`** — `TestUnixNano` was unrelated to the SDK
3. **Deleted `error_test.go`** — `TestDeleteListener` was a misplaced unit test
4. **Added `//go:build integration`** tag to all 21 test files

### 2.5 Running tests

```bash
# Unit tests only (default — no credentials needed)
go test ./greennode/...

# Integration tests (requires env.yaml or Vault credentials)
go test -tags=integration ./test/...

# Both
go test -tags=integration ./...
```

### 2.6 Known issues

| Issue | Status | Description |
| ----- | ------ | ----------- |
| Hardcoded resource IDs | Open | Tests use hardcoded UUIDs (e.g. `lb-8f54cbd4-...`) — fragile if resources are deleted |
| `TestGetLoadBalancerFailure` | Open | Named as failure test but asserts success conditions |

---

## 3. Implementation roadmap

### Phase 1: Restructure integration tests — **DONE**

- [x] Extract helpers from `identity_test.go` into `helpers_test.go`
- [x] Delete `other_test.go`
- [x] Delete `error_test.go`
- [x] Add `//go:build integration` to all `test/*_test.go` files
- [ ] Move hardcoded resource IDs into `helpers_test.go` as named constants

### Phase 2: Vault integration

- [ ] Add `github.com/hashicorp/vault/api` dependency
- [ ] Implement Vault reader in `helpers_test.go` with `env.yaml` fallback
- [ ] Document Vault setup in this file
- [ ] Set up CI/CD pipeline with Vault-authenticated test runs

### Phase 3: Expand unit test coverage

- [ ] `auth/` — TOTP against RFC 6238, endpoint resolution
- [ ] `client/` — increase to 85%+
- [ ] `entity/` — increase to 70%+
- [ ] Service request/response packages (compute, LB, DNS, GLB, network, portal)

### Phase 4: CI/CD integration

- [ ] GitHub Actions workflow for unit tests (on every PR)
- [ ] GitHub Actions workflow for integration tests (nightly or manual trigger, Vault-authenticated)
- [ ] Coverage reporting and badges

---

## 4. Resolved issues (historical)

<details>
<summary>Previously fixed issues</summary>

- **Developer names in test code** — personal identifiers replaced with generic names
- **False-passing tests** — 12 tests changed from `t.Log` to `t.Fatalf`
- **Nil pointer panic in TestAuthenPass** — added nil guard
- **Config helper bloat** — 9 functions consolidated into `newClientFromEnvKeys`
- **Test name typo** — `TestASuperuthenPass` renamed to `TestSuperAdminAuthenPass`
- **No unit tests** — 139+ unit tests added across core packages

</details>
