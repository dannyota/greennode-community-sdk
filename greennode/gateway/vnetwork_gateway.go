package gateway

import (
	"github.com/dannyota/greennode-community-sdk/v2/greennode/client"
	"github.com/dannyota/greennode-community-sdk/v2/greennode/services/network"
)

type vnetworkGatewayV1 struct {
	networkService network.NetworkServiceV1
}

type vnetworkGatewayV2 struct {
	networkService network.NetworkServiceV2
}

type vnetworkGatewayInternalV1 struct {
	networkService network.NetworkServiceInternalV1
}

func (g *vnetworkGatewayV1) NetworkService() network.NetworkServiceV1 {
	return g.networkService
}

func (g *vnetworkGatewayV2) NetworkService() network.NetworkServiceV2 {
	return g.networkService
}

func (g *vnetworkGatewayInternalV1) NetworkService() network.NetworkServiceInternalV1 {
	return g.networkService
}

func NewVNetworkGatewayV1(svcClient client.ServiceClient) VNetworkGatewayV1 {
	return &vnetworkGatewayV1{
		networkService: network.NewNetworkServiceV1(svcClient),
	}
}

func NewVNetworkGatewayV2(svcClient client.ServiceClient) VNetworkGatewayV2 {
	return &vnetworkGatewayV2{
		networkService: network.NewNetworkServiceV2(svcClient),
	}
}

func NewVNetworkGatewayInternalV1(svcClient client.ServiceClient) VNetworkGatewayInternalV1 {
	return &vnetworkGatewayInternalV1{
		networkService: network.NewNetworkServiceInternalV1(svcClient),
	}
}
