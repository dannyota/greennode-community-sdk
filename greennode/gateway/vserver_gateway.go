package gateway

import (
	"github.com/dannyota/greennode-community-sdk/v2/greennode/client"
	"github.com/dannyota/greennode-community-sdk/v2/greennode/services/compute"
	"github.com/dannyota/greennode-community-sdk/v2/greennode/services/network"
	"github.com/dannyota/greennode-community-sdk/v2/greennode/services/portal"
	"github.com/dannyota/greennode-community-sdk/v2/greennode/services/server"
	"github.com/dannyota/greennode-community-sdk/v2/greennode/services/volume"
)

type vserverGatewayV1 struct {
	portalService portal.PortalServiceV1
	volumeService volume.VolumeServiceV1
}

type vServerGatewayInternalV1 struct {
	serverService server.ServerServiceInternalV1
}

type vserverGatewayV2 struct {
	networkService network.NetworkServiceV2
	computeService compute.ComputeServiceV2
	portalService  portal.PortalServiceV2
	volumeService  volume.VolumeServiceV2
}

func NewVServerGatewayV1(svcClient client.ServiceClient) VServerGatewayV1 {
	return &vserverGatewayV1{
		portalService: portal.NewPortalServiceV1(svcClient),
		volumeService: volume.NewVolumeServiceV1(svcClient),
	}
}

func NewVServerGatewayV2(svcClient client.ServiceClient) VServerGatewayV2 {
	return &vserverGatewayV2{
		networkService: network.NewNetworkServiceV2(svcClient),
		computeService: compute.NewComputeServiceV2(svcClient),
		portalService:  portal.NewPortalServiceV2(svcClient),
		volumeService:  volume.NewVolumeServiceV2(svcClient),
	}
}

func NewVServerGatewayInternalV1(svcClient client.ServiceClient) VServerGatewayInternalV1 {
	return &vServerGatewayInternalV1{
		serverService: server.NewServerServiceInternalV1(svcClient),
	}
}

func (s *vserverGatewayV1) PortalService() portal.PortalServiceV1 {
	return s.portalService
}

func (s *vserverGatewayV1) VolumeService() volume.VolumeServiceV1 {
	return s.volumeService
}

func (s *vserverGatewayV2) NetworkService() network.NetworkServiceV2 {
	return s.networkService
}

func (s *vserverGatewayV2) ComputeService() compute.ComputeServiceV2 {
	return s.computeService
}

func (s *vserverGatewayV2) PortalService() portal.PortalServiceV2 {
	return s.portalService
}

func (s *vserverGatewayV2) VolumeService() volume.VolumeServiceV2 {
	return s.volumeService
}

func (s *vServerGatewayInternalV1) ServerService() server.ServerServiceInternalV1 {
	return s.serverService
}
