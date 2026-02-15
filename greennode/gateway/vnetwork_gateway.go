package gateway

import (
	"github.com/dannyota/greennode-community-sdk/v2/greennode/client"
	"github.com/dannyota/greennode-community-sdk/v2/greennode/services/network"
)

type VNetworkGatewayV1 struct {
	networkService network.NetworkServiceV1
}

type VNetworkGatewayV2 struct {
	networkService network.NetworkServiceV2
}

type VNetworkGatewayInternalV1 struct {
	networkService network.NetworkServiceInternalV1
}

func (g *VNetworkGatewayV1) NetworkService() network.NetworkServiceV1 {
	return g.networkService
}

func (g *VNetworkGatewayV2) NetworkService() network.NetworkServiceV2 {
	return g.networkService
}

func (g *VNetworkGatewayInternalV1) NetworkService() network.NetworkServiceInternalV1 {
	return g.networkService
}

func NewVNetworkGatewayV1(svcClient client.ServiceClient) *VNetworkGatewayV1 {
	return &VNetworkGatewayV1{
		networkService: network.NewNetworkServiceV1(svcClient),
	}
}

func NewVNetworkGatewayV2(svcClient client.ServiceClient) *VNetworkGatewayV2 {
	return &VNetworkGatewayV2{
		networkService: network.NewNetworkServiceV2(svcClient),
	}
}

func NewVNetworkGatewayInternalV1(svcClient client.ServiceClient) *VNetworkGatewayInternalV1 {
	return &VNetworkGatewayInternalV1{
		networkService: network.NewNetworkServiceInternalV1(svcClient),
	}
}
