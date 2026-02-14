package loadbalancer

import (
	"github.com/dannyota/greennode-community-sdk/v2/greennode/client"
	"github.com/dannyota/greennode-community-sdk/v2/greennode/services/loadbalancer/inter"
	lbv2 "github.com/dannyota/greennode-community-sdk/v2/greennode/services/loadbalancer/v2"
)

func NewLoadBalancerServiceV2(lbSvcClient, serverSvcClient client.ServiceClient) LoadBalancerServiceV2 {
	return &lbv2.LoadBalancerServiceV2{
		VLBClient:     lbSvcClient,
		VServerClient: serverSvcClient,
	}
}

func NewLoadBalancerServiceInternal(svcClient client.ServiceClient) LoadBalancerServiceInternal {
	return &inter.LoadBalancerServiceInternal{
		VLBClient: svcClient,
	}
}
