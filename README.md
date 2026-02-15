# GreenNode Community SDK

A community Go SDK for GreenNode cloud services.

<hr>

###### Table of contents

- [GreenNode Community SDK](#greennode-community-sdk)
          - [Table of contents](#table-of-contents)
- [Introduction](#introduction)
- [Usage](#usage)
- [Contributing](#contributing)

<hr>

# Introduction

- `greennode-community-sdk` is a Go SDK for GreenNode cloud services. It helps you interact with cloud services easily.

# Usage
- You can install the SDK by running the following command:
  ```bash
  go get github.com/dannyota/greennode-community-sdk/v2
  ```

- Now for example, imagine you want to list all available load-balancer packages. You can implement this code in your Go Application:
  ```go
  package main

  import (
    "context"
    "fmt"

    "github.com/dannyota/greennode-community-sdk/v2/client"
    lbv2 "github.com/dannyota/greennode-community-sdk/v2/greennode/services/loadbalancer/v2"
  )

  func main() {
    c := validSdkConfig()
    opt := lbv2.NewListLoadBalancerPackagesRequest()
    packages, sdkerr := c.VLBGateway().V2().LoadBalancerService().ListLoadBalancerPackages(opt)
    if sdkerr != nil {
      fmt.Printf("Expect nil but got %+v", sdkerr)
    }

    if packages == nil {
      fmt.Printf("Expect not nil but got nil")
    }

    for _, pkg := range packages.Items {
      fmt.Printf("Package: %+v", pkg)
    }
  }

  func validSdkConfig() *client.Client {
    clientID, clientSecret := "__PUT_YOUR_CLIENT_ID__", "__PUT_YOUR_CLIENT_SECRET__"
    sdkConfig := client.NewSdkConfigure().
      WithClientID(clientID).
      WithClientSecret(clientSecret).
      WithProjectID("__PUT_YOUR_PROJECT_ID__").
      WithZoneID("65e12ffcb6d82cd39f8cf023").
      WithIAMEndpoint("https://iamapis.vngcloud.vn/accounts-api").
      WithVServerEndpoint("https://hcm-3.api.vngcloud.vn/vserver/vserver-gateway").
      WithVLBEndpoint("https://hcm-3.api.vngcloud.vn/vserver/vlb-gateway").
      WithVNetworkEndpoint("https://vnetwork-hcm03.vngcloud.vn/vnetwork-gateway/vnetwork")

    return client.NewClient(context.TODO()).WithRetryCount(1).WithSleep(10).Configure(sdkConfig)
  }
  ```

# Contributing

- To release a new version of the SDK, create a new tag with the format `vX.Y.Z` and push it to the repository:

  ```bash
  git tag -am "release vX.Y.Z" vX.Y.Z
  git push --tags
  ```

- To get the latest version of the SDK:
  ```bash
  git tag -l --sort=-creatordate | head -n 1
  ```
