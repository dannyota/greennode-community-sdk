package gateway

import (
	"github.com/dannyota/greennode-community-sdk/v2/greennode/client"
	"github.com/dannyota/greennode-community-sdk/v2/greennode/services/loadbalancer"
)

type VLBGatewayV2 struct {
	lbService loadbalancer.LoadBalancerServiceV2
}

type VLBGatewayInternal struct {
	lbService loadbalancer.LoadBalancerServiceInternal
}

func NewVLBGatewayV2(lbSvcClient, serverSvcClient client.ServiceClient) *VLBGatewayV2 {
	return &VLBGatewayV2{
		lbService: loadbalancer.NewLoadBalancerServiceV2(lbSvcClient, serverSvcClient),
	}
}

func NewVLBGatewayInternal(svcClient client.ServiceClient) *VLBGatewayInternal {
	return &VLBGatewayInternal{
		lbService: loadbalancer.NewLoadBalancerServiceInternal(svcClient),
	}
}

func (g *VLBGatewayInternal) LoadBalancerService() loadbalancer.LoadBalancerServiceInternal {
	return g.lbService
}

func (g *VLBGatewayV2) LoadBalancerService() loadbalancer.LoadBalancerServiceV2 {
	return g.lbService
}
