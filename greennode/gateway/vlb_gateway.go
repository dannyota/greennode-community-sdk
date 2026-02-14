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

func NewVLBGatewayV2(plbSvcClient, pserverSvcClient client.ServiceClient) VLBGatewayV2 {
	return &vlbGatewayV2{
		lbService: loadbalancer.NewLoadBalancerServiceV2(plbSvcClient, pserverSvcClient),
	}
}

func NewVLBGatewayInternal(psvcClient client.ServiceClient) VLBGatewayInternal {
	return &vlbGatewayInternal{
		lbService: loadbalancer.NewLoadBalancerServiceInternal(psvcClient),
	}
}

func (s *vlbGatewayInternal) LoadBalancerService() loadbalancer.LoadBalancerServiceInternal {
	return s.lbService
}

func (s *vlbGatewayV2) LoadBalancerService() loadbalancer.LoadBalancerServiceV2 {
	return s.lbService
}
