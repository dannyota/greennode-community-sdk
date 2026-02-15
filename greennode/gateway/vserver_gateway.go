package gateway

import (
	"github.com/dannyota/greennode-community-sdk/v2/greennode/client"
	"github.com/dannyota/greennode-community-sdk/v2/greennode/services/compute"
	"github.com/dannyota/greennode-community-sdk/v2/greennode/services/network"
	"github.com/dannyota/greennode-community-sdk/v2/greennode/services/portal"
	"github.com/dannyota/greennode-community-sdk/v2/greennode/services/server"
	"github.com/dannyota/greennode-community-sdk/v2/greennode/services/volume"
)

type VServerGatewayV1 struct {
	portalService portal.PortalServiceV1
	volumeService volume.VolumeServiceV1
}

type VServerGatewayInternalV1 struct {
	serverService server.ServerServiceInternalV1
}

type VServerGatewayV2 struct {
	networkService network.NetworkServiceV2
	computeService compute.ComputeServiceV2
	portalService  portal.PortalServiceV2
	volumeService  volume.VolumeServiceV2
}

func NewVServerGatewayV1(svcClient client.ServiceClient) *VServerGatewayV1 {
	return &VServerGatewayV1{
		portalService: portal.NewPortalServiceV1(svcClient),
		volumeService: volume.NewVolumeServiceV1(svcClient),
	}
}

func NewVServerGatewayV2(svcClient client.ServiceClient) *VServerGatewayV2 {
	return &VServerGatewayV2{
		networkService: network.NewNetworkServiceV2(svcClient),
		computeService: compute.NewComputeServiceV2(svcClient),
		portalService:  portal.NewPortalServiceV2(svcClient),
		volumeService:  volume.NewVolumeServiceV2(svcClient),
	}
}

func NewVServerGatewayInternalV1(svcClient client.ServiceClient) *VServerGatewayInternalV1 {
	return &VServerGatewayInternalV1{
		serverService: server.NewServerServiceInternalV1(svcClient),
	}
}

func (g *VServerGatewayV1) PortalService() portal.PortalServiceV1 {
	return g.portalService
}

func (g *VServerGatewayV1) VolumeService() volume.VolumeServiceV1 {
	return g.volumeService
}

func (g *VServerGatewayV2) NetworkService() network.NetworkServiceV2 {
	return g.networkService
}

func (g *VServerGatewayV2) ComputeService() compute.ComputeServiceV2 {
	return g.computeService
}

func (g *VServerGatewayV2) PortalService() portal.PortalServiceV2 {
	return g.portalService
}

func (g *VServerGatewayV2) VolumeService() volume.VolumeServiceV2 {
	return g.volumeService
}

func (g *VServerGatewayInternalV1) ServerService() server.ServerServiceInternalV1 {
	return g.serverService
}
