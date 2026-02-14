package network

import (
	"github.com/dannyota/greennode-community-sdk/v2/greennode/client"
	networkv1 "github.com/dannyota/greennode-community-sdk/v2/greennode/services/network/v1"
	networkv2 "github.com/dannyota/greennode-community-sdk/v2/greennode/services/network/v2"
)

func NewNetworkServiceV2(psvcClient client.ServiceClient) NetworkServiceV2 {
	return &networkv2.NetworkServiceV2{
		VserverClient: psvcClient,
	}
}

func NewNetworkServiceV1(psvcClient client.ServiceClient) NetworkServiceV1 {
	return &networkv1.NetworkServiceV1{
		VNetworkClient: psvcClient,
	}
}

func NewNetworkServiceInternalV1(psvcClient client.ServiceClient) NetworkServiceInternalV1 {
	return &networkv1.NetworkServiceInternalV1{
		VNetworkClient: psvcClient,
	}
}
