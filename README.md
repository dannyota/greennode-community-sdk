# GreenNode Community SDK

A community Go SDK for GreenNode cloud services.

## Installation

```bash
go get github.com/dannyota/greennode-community-sdk/v2
```

## Quick Start

### Service account (client credentials)

```go
package main

import (
  "context"
  "fmt"

  "github.com/dannyota/greennode-community-sdk/v2/greennode"
  lbv2 "github.com/dannyota/greennode-community-sdk/v2/greennode/services/loadbalancer/v2"
)

func main() {
  c, err := greennode.NewClient(context.Background(), greennode.Config{
    Region:       "hcm-3",
    ClientID:     "__YOUR_CLIENT_ID__",
    ClientSecret: "__YOUR_CLIENT_SECRET__",
    ProjectID:    "__YOUR_PROJECT_ID__",
  })
  if err != nil {
    panic(err)
  }

  packages, err := c.LoadBalancer.ListLoadBalancerPackages(
    context.Background(), lbv2.NewListLoadBalancerPackagesRequest())
  if err != nil {
    panic(err)
  }

  for _, pkg := range packages.Items {
    fmt.Printf("Package: %+v\n", pkg)
  }
}
```

### IAM user (username/password + optional TOTP)

```go
package main

import (
  "context"

  "github.com/dannyota/greennode-community-sdk/v2/greennode"
  "github.com/dannyota/greennode-community-sdk/v2/greennode/auth"
)

func main() {
  c, err := greennode.NewClient(context.Background(), greennode.Config{
    Region:    "hcm-3",
    ProjectID: "__YOUR_PROJECT_ID__",
    IAMAuth: &auth.IAMUserAuth{
      RootEmail: "root@company.com",
      Username:  "your-username",
      Password:  "your-password",
      TOTP:      &auth.SecretTOTP{Secret: "YOUR_BASE32_SECRET"}, // omit if no 2FA
    },
  })
  if err != nil {
    panic(err)
  }

  // Use c.Compute, c.LoadBalancer, etc. as normal
  _ = c
}
```

The `Region` field (e.g. `"hcm-3"`, `"han-1"`) derives all endpoint URLs automatically.
Explicit endpoint fields (e.g. `VServerEndpoint`) override the defaults if set.

#### TOTP providers

| Provider | Usage |
|----------|-------|
| `&auth.SecretTOTP{Secret: "..."}` | Compute TOTP from a base32 shared secret |
| `auth.TOTPFunc(func(ctx) (string, error) { ... })` | Bring your own source (Vault, env var, CLI prompt) |
| `nil` | No 2FA required |

## Services

| Service       | Description                                    |
| ------------- | ---------------------------------------------- |
| Compute       | Server lifecycle, floating IPs, server groups  |
| Volume        | Block volumes, snapshots, volume types         |
| Network       | VPCs, subnets, security groups, endpoints      |
| Load Balancer | Load balancers, listeners, pools, certificates |
| GLB           | Global load balancer pools, listeners          |
| DNS           | Hosted zones, DNS records                      |
| Identity      | OAuth2 token acquisition                       |
| Portal        | Portal info, project listing                   |

See [docs/architecture.md](docs/architecture.md) for the full architecture overview.

## Documentation

- [Architecture](docs/architecture.md) — layered design, request lifecycle, error handling
- [Go Style Audit](docs/go-style-audit.md) — refactoring roadmap and status
- [Legacy Patterns Review](docs/legacy-patterns-review.md) — Java/C# pattern cleanup tracker
- [Test Coverage](docs/test-coverage.md) — coverage map and improvement plan

## Contributing

Contributions are welcome. Please open an issue or submit a pull request.
