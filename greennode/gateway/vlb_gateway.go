package gateway

import (
	"github.com/dannyota/greennode-community-sdk/v2/greennode/client"
	"github.com/dannyota/greennode-community-sdk/v2/greennode/services/loadbalancer"
)

type vlbGatewayV2 struct {
	lbService loadbalancer.LoadBalancerServiceV2
}

type vlbGatewayInternal struct {
	lbService loadbalancer.LoadBalancerServiceInternal
}

func NewVLBGatewayV2(lbSvcClient, serverSvcClient client.ServiceClient) VLBGatewayV2 {
	return &vlbGatewayV2{
		lbService: loadbalancer.NewLoadBalancerServiceV2(lbSvcClient, serverSvcClient),
	}
}

func NewVLBGatewayInternal(svcClient client.ServiceClient) VLBGatewayInternal {
	return &vlbGatewayInternal{
		lbService: loadbalancer.NewLoadBalancerServiceInternal(svcClient),
	}
}

func (g *vlbGatewayInternal) LoadBalancerService() loadbalancer.LoadBalancerServiceInternal {
	return g.lbService
}

func (g *vlbGatewayV2) LoadBalancerService() loadbalancer.LoadBalancerServiceV2 {
	return g.lbService
}
