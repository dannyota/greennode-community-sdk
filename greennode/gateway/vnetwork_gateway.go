package gateway

import (
	"github.com/dannyota/greennode-community-sdk/v2/greennode/client"
	"github.com/dannyota/greennode-community-sdk/v2/greennode/services/network"
)

type vnetworkGatewayV1 struct {
	networkService network.NetworkServiceV1
}

type vnetworkGatewayInternalV1 struct {
	networkService network.NetworkServiceInternalV1
}

func (s *vnetworkGatewayV1) NetworkService() network.NetworkServiceV1 {
	return s.networkService
}

func (s *vnetworkGatewayInternalV1) NetworkService() network.NetworkServiceInternalV1 {
	return s.networkService
}

func NewVNetworkGatewayV1(svcClient client.ServiceClient) VNetworkGatewayV1 {
	return &vnetworkGatewayV1{
		networkService: network.NewNetworkServiceV1(svcClient),
	}
}

func NewVNetworkGatewayInternalV1(svcClient client.ServiceClient) VNetworkGatewayInternalV1 {
	return &vnetworkGatewayInternalV1{
		networkService: network.NewNetworkServiceInternalV1(svcClient),
	}
}
