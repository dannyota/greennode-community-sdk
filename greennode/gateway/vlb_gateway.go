package gateway

import (
	"github.com/dannyota/greennode-community-sdk/v2/greennode/client"
	"github.com/dannyota/greennode-community-sdk/v2/greennode/services/loadbalancer/inter"
	lbv2 "github.com/dannyota/greennode-community-sdk/v2/greennode/services/loadbalancer/v2"
)

type VLBGatewayV2 struct {
	lbService *lbv2.LoadBalancerServiceV2
}

type VLBGatewayInternal struct {
	lbService *inter.LoadBalancerServiceInternal
}

func NewVLBGatewayV2(lbSvcClient, serverSvcClient client.ServiceClient) *VLBGatewayV2 {
	return &VLBGatewayV2{
		lbService: &lbv2.LoadBalancerServiceV2{VLBClient: lbSvcClient, VServerClient: serverSvcClient},
	}
}

func NewVLBGatewayInternal(svcClient client.ServiceClient) *VLBGatewayInternal {
	return &VLBGatewayInternal{
		lbService: &inter.LoadBalancerServiceInternal{VLBClient: svcClient},
	}
}

func (g *VLBGatewayInternal) LoadBalancerService() *inter.LoadBalancerServiceInternal {
	return g.lbService
}

func (g *VLBGatewayV2) LoadBalancerService() *lbv2.LoadBalancerServiceV2 {
	return g.lbService
}
