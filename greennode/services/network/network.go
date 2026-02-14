package network

import (
	"github.com/dannyota/greennode-community-sdk/v2/greennode/client"
	networkv1 "github.com/dannyota/greennode-community-sdk/v2/greennode/services/network/v1"
	networkv2 "github.com/dannyota/greennode-community-sdk/v2/greennode/services/network/v2"
)

func NewNetworkServiceV2(svcClient client.ServiceClient) NetworkServiceV2 {
	return &networkv2.NetworkServiceV2{
		VserverClient: svcClient,
	}
}

func NewNetworkServiceV1(svcClient client.ServiceClient) NetworkServiceV1 {
	return &networkv1.NetworkServiceV1{
		VNetworkClient: svcClient,
	}
}

func NewNetworkServiceInternalV1(svcClient client.ServiceClient) NetworkServiceInternalV1 {
	return &networkv1.NetworkServiceInternalV1{
		VNetworkClient: svcClient,
	}
}
