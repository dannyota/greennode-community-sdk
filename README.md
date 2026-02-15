# GreenNode Community SDK

A community Go SDK for GreenNode cloud services.

## Installation

```bash
go get github.com/dannyota/greennode-community-sdk/v2
```

## Quick Start

```go
package main

import (
  "context"
  "fmt"

  "github.com/dannyota/greennode-community-sdk/v2/client"
  lbv2 "github.com/dannyota/greennode-community-sdk/v2/greennode/services/loadbalancer/v2"
)

func main() {
  sdkConfig := client.NewSdkConfigure().
    WithClientID("__YOUR_CLIENT_ID__").
    WithClientSecret("__YOUR_CLIENT_SECRET__").
    WithProjectID("__YOUR_PROJECT_ID__").
    WithIAMEndpoint("https://iamapis.vngcloud.vn/accounts-api").
    WithVLBEndpoint("https://hcm-3.api.vngcloud.vn/vserver/vlb-gateway")

  c := client.NewClient().WithRetryCount(1).WithSleep(10).Configure(sdkConfig)

  packages, err := c.VLBGateway().V2().LoadBalancerService().
    ListLoadBalancerPackages(context.Background(), lbv2.NewListLoadBalancerPackagesRequest())
  if err != nil {
    panic(err)
  }

  for _, pkg := range packages.Items {
    fmt.Printf("Package: %+v\n", pkg)
  }
}
```

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
- [Test Coverage](docs/test-coverage.md) — coverage map and improvement plan

## Contributing

Contributions are welcome. Please open an issue or submit a pull request.
