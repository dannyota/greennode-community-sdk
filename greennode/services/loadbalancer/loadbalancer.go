package loadbalancer

import (
	lsclient "github.com/dannyota/greennode-community-sdk/v2/greennode/client"
	lsloadbalancerInternal "github.com/dannyota/greennode-community-sdk/v2/greennode/services/loadbalancer/inter"
	lsloadbalancerV2 "github.com/dannyota/greennode-community-sdk/v2/greennode/services/loadbalancer/v2"
)

func NewLoadBalancerServiceV2(plbSvcClient, pserverSvcClient lsclient.IServiceClient) ILoadBalancerServiceV2 {
	return &lsloadbalancerV2.LoadBalancerServiceV2{
		VLBClient:     plbSvcClient,
		VServerClient: pserverSvcClient,
	}
}

func NewLoadBalancerServiceInternal(psvcClient lsclient.IServiceClient) ILoadBalancerServiceInternal {
	return &lsloadbalancerInternal.LoadBalancerServiceInternal{
		VLBClient: psvcClient,
	}
}
