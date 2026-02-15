package gateway

import (
	"github.com/dannyota/greennode-community-sdk/v2/greennode/client"
	computev2 "github.com/dannyota/greennode-community-sdk/v2/greennode/services/compute/v2"
	networkv2 "github.com/dannyota/greennode-community-sdk/v2/greennode/services/network/v2"
	portalv1 "github.com/dannyota/greennode-community-sdk/v2/greennode/services/portal/v1"
	portalv2 "github.com/dannyota/greennode-community-sdk/v2/greennode/services/portal/v2"
	serverv1 "github.com/dannyota/greennode-community-sdk/v2/greennode/services/server/v1"
	volumev1 "github.com/dannyota/greennode-community-sdk/v2/greennode/services/volume/v1"
	volumev2 "github.com/dannyota/greennode-community-sdk/v2/greennode/services/volume/v2"
)

type VServerGatewayV1 struct {
	portalService *portalv1.PortalServiceV1
	volumeService *volumev1.VolumeServiceV1
}

type VServerGatewayInternalV1 struct {
	serverService *serverv1.ServerServiceInternalV1
}

type VServerGatewayV2 struct {
	networkService *networkv2.NetworkServiceV2
	computeService *computev2.ComputeServiceV2
	portalService  *portalv2.PortalServiceV2
	volumeService  *volumev2.VolumeServiceV2
}

func NewVServerGatewayV1(svcClient client.ServiceClient) *VServerGatewayV1 {
	return &VServerGatewayV1{
		portalService: &portalv1.PortalServiceV1{PortalClient: svcClient},
		volumeService: &volumev1.VolumeServiceV1{VServerClient: svcClient},
	}
}

func NewVServerGatewayV2(svcClient client.ServiceClient) *VServerGatewayV2 {
	return &VServerGatewayV2{
		networkService: &networkv2.NetworkServiceV2{VServerClient: svcClient},
		computeService: &computev2.ComputeServiceV2{VServerClient: svcClient},
		portalService:  &portalv2.PortalServiceV2{PortalClient: svcClient},
		volumeService:  &volumev2.VolumeServiceV2{VServerClient: svcClient},
	}
}

func NewVServerGatewayInternalV1(svcClient client.ServiceClient) *VServerGatewayInternalV1 {
	return &VServerGatewayInternalV1{
		serverService: &serverv1.ServerServiceInternalV1{VServerClient: svcClient},
	}
}

func (g *VServerGatewayV1) PortalService() *portalv1.PortalServiceV1 {
	return g.portalService
}

func (g *VServerGatewayV1) VolumeService() *volumev1.VolumeServiceV1 {
	return g.volumeService
}

func (g *VServerGatewayV2) NetworkService() *networkv2.NetworkServiceV2 {
	return g.networkService
}

func (g *VServerGatewayV2) ComputeService() *computev2.ComputeServiceV2 {
	return g.computeService
}

func (g *VServerGatewayV2) PortalService() *portalv2.PortalServiceV2 {
	return g.portalService
}

func (g *VServerGatewayV2) VolumeService() *volumev2.VolumeServiceV2 {
	return g.volumeService
}

func (g *VServerGatewayInternalV1) ServerService() *serverv1.ServerServiceInternalV1 {
	return g.serverService
}
