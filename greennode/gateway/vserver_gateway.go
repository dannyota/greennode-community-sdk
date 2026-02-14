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
	portalService portal.IPortalServiceV1
	volumeService volume.IVolumeServiceV1
}

type vServerGatewayInternalV1 struct {
	serverService server.IServerServiceInternalV1
}

type vserverGatewayV2 struct {
	networkService network.INetworkServiceV2
	computeService compute.IComputeServiceV2
	portalService  portal.IPortalServiceV2
	volumeService  volume.IVolumeServiceV2
}

func NewVServerGatewayV1(psvcClient client.IServiceClient) IVServerGatewayV1 {
	return &vserverGatewayV1{
		portalService: portal.NewPortalServiceV1(psvcClient),
		volumeService: volume.NewVolumeServiceV1(psvcClient),
	}
}

func NewVServerGatewayV2(psvcClient client.IServiceClient) IVServerGatewayV2 {
	return &vserverGatewayV2{
		networkService: network.NewNetworkServiceV2(psvcClient),
		computeService: compute.NewComputeServiceV2(psvcClient),
		portalService:  portal.NewPortalServiceV2(psvcClient),
		volumeService:  volume.NewVolumeServiceV2(psvcClient),
	}
}

func NewVServerGatewayInternalV1(psvcClient client.IServiceClient) IVServerGatewayInternalV1 {
	return &vServerGatewayInternalV1{
		serverService: server.NewServerServiceInternalV1(psvcClient),
	}
}

func (s *vserverGatewayV1) PortalService() portal.IPortalServiceV1 {
	return s.portalService
}

func (s *vserverGatewayV1) VolumeService() volume.IVolumeServiceV1 {
	return s.volumeService
}

func (s *vserverGatewayV2) NetworkService() network.INetworkServiceV2 {
	return s.networkService
}

func (s *vserverGatewayV2) ComputeService() compute.IComputeServiceV2 {
	return s.computeService
}

func (s *vserverGatewayV2) PortalService() portal.IPortalServiceV2 {
	return s.portalService
}

func (s *vserverGatewayV2) VolumeService() volume.IVolumeServiceV2 {
	return s.volumeService
}

func (s *vServerGatewayInternalV1) ServerService() server.IServerServiceInternalV1 {
	return s.serverService
}
