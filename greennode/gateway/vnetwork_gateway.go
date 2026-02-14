package gateway

import (
	"github.com/dannyota/greennode-community-sdk/v2/greennode/client"
	"github.com/dannyota/greennode-community-sdk/v2/greennode/services/network"
)

type vnetworkGatewayV1 struct {
	networkService network.INetworkServiceV1
}

type vnetworkGatewayInternalV1 struct {
	networkService network.INetworkServiceInternalV1
}

func (s *vnetworkGatewayV1) NetworkService() network.INetworkServiceV1 {
	return s.networkService
}

func (s *vnetworkGatewayInternalV1) NetworkService() network.INetworkServiceInternalV1 {
	return s.networkService
}

func NewVNetworkGatewayV1(psvcClient client.IServiceClient) IVNetworkGatewayV1 {
	return &vnetworkGatewayV1{
		networkService: network.NewNetworkServiceV1(psvcClient),
	}
}

func NewVNetworkGatewayInternalV1(psvcClient client.IServiceClient) IVNetworkGatewayInternalV1 {
	return &vnetworkGatewayInternalV1{
		networkService: network.NewNetworkServiceInternalV1(psvcClient),
	}
}
