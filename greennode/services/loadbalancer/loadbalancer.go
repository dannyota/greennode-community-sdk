package loadbalancer

import (
	"github.com/dannyota/greennode-community-sdk/v2/greennode/client"
	"github.com/dannyota/greennode-community-sdk/v2/greennode/services/loadbalancer/inter"
	lbv2 "github.com/dannyota/greennode-community-sdk/v2/greennode/services/loadbalancer/v2"
)

func NewLoadBalancerServiceV2(plbSvcClient, pserverSvcClient client.IServiceClient) ILoadBalancerServiceV2 {
	return &lbv2.LoadBalancerServiceV2{
		VLBClient:     plbSvcClient,
		VServerClient: pserverSvcClient,
	}
}

func NewLoadBalancerServiceInternal(psvcClient client.IServiceClient) ILoadBalancerServiceInternal {
	return &inter.LoadBalancerServiceInternal{
		VLBClient: psvcClient,
	}
}
