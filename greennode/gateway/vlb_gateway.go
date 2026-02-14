package gateway

import (
	"github.com/dannyota/greennode-community-sdk/v2/greennode/client"
	"github.com/dannyota/greennode-community-sdk/v2/greennode/services/loadbalancer"
)

type vlbGatewayV2 struct {
	lbService loadbalancer.ILoadBalancerServiceV2
}

type vlbGatewayInternal struct {
	lbService loadbalancer.ILoadBalancerServiceInternal
}

func NewVLBGatewayV2(plbSvcClient, pserverSvcClient client.IServiceClient) IVLBGatewayV2 {
	return &vlbGatewayV2{
		lbService: loadbalancer.NewLoadBalancerServiceV2(plbSvcClient, pserverSvcClient),
	}
}

func NewVLBGatewayInternal(psvcClient client.IServiceClient) IVLBGatewayInternal {
	return &vlbGatewayInternal{
		lbService: loadbalancer.NewLoadBalancerServiceInternal(psvcClient),
	}
}

func (s *vlbGatewayInternal) LoadBalancerService() loadbalancer.ILoadBalancerServiceInternal {
	return s.lbService
}

func (s *vlbGatewayV2) LoadBalancerService() loadbalancer.ILoadBalancerServiceV2 {
	return s.lbService
}
