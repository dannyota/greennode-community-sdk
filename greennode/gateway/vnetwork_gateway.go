package gateway

import (
	"github.com/dannyota/greennode-community-sdk/v2/greennode/client"
	networkv1 "github.com/dannyota/greennode-community-sdk/v2/greennode/services/network/v1"
	networkv2 "github.com/dannyota/greennode-community-sdk/v2/greennode/services/network/v2"
)

type VNetworkGatewayV1 struct {
	networkService *networkv1.NetworkServiceV1
}

type VNetworkGatewayV2 struct {
	networkService *networkv2.NetworkServiceV2
}

type VNetworkGatewayInternalV1 struct {
	networkService *networkv1.NetworkServiceInternalV1
}

func (g *VNetworkGatewayV1) NetworkService() *networkv1.NetworkServiceV1 {
	return g.networkService
}

func (g *VNetworkGatewayV2) NetworkService() *networkv2.NetworkServiceV2 {
	return g.networkService
}

func (g *VNetworkGatewayInternalV1) NetworkService() *networkv1.NetworkServiceInternalV1 {
	return g.networkService
}

func NewVNetworkGatewayV1(svcClient *client.ServiceClient) *VNetworkGatewayV1 {
	return &VNetworkGatewayV1{
		networkService: &networkv1.NetworkServiceV1{VNetworkClient: svcClient},
	}
}

func NewVNetworkGatewayV2(svcClient *client.ServiceClient) *VNetworkGatewayV2 {
	return &VNetworkGatewayV2{
		networkService: &networkv2.NetworkServiceV2{VServerClient: svcClient},
	}
}

func NewVNetworkGatewayInternalV1(svcClient *client.ServiceClient) *VNetworkGatewayInternalV1 {
	return &VNetworkGatewayInternalV1{
		networkService: &networkv1.NetworkServiceInternalV1{VNetworkClient: svcClient},
	}
}
