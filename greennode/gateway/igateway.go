package gateway

import (
	"github.com/dannyota/greennode-community-sdk/v2/greennode/services/compute"
	"github.com/dannyota/greennode-community-sdk/v2/greennode/services/dns"
	"github.com/dannyota/greennode-community-sdk/v2/greennode/services/glb"
	"github.com/dannyota/greennode-community-sdk/v2/greennode/services/identity"
	"github.com/dannyota/greennode-community-sdk/v2/greennode/services/loadbalancer"
	"github.com/dannyota/greennode-community-sdk/v2/greennode/services/network"
	"github.com/dannyota/greennode-community-sdk/v2/greennode/services/portal"
	"github.com/dannyota/greennode-community-sdk/v2/greennode/services/server"
	"github.com/dannyota/greennode-community-sdk/v2/greennode/services/volume"
)

type IAMGateway interface {
	V2() IAMGatewayV2
}

type IAMGatewayV2 interface {
	IdentityService() identity.IdentityServiceV2
}

type VServerGateway interface {
	V1() VServerGatewayV1
	InternalV1() VServerGatewayInternalV1
	V2() VServerGatewayV2
	// GetEndpoint returns the endpoint of the vServer service
	GetEndpoint() string
}

type VNetworkGateway interface {
	V1() VNetworkGatewayV1
	V2() VNetworkGatewayV2
	InternalV1() VNetworkGatewayInternalV1
	GetEndpoint() string
}

type VServerGatewayV1 interface {
	PortalService() portal.PortalServiceV1
	VolumeService() volume.VolumeServiceV1
}

type VServerGatewayInternalV1 interface {
	ServerService() server.ServerServiceInternalV1
}

type VServerGatewayV2 interface {
	NetworkService() network.NetworkServiceV2
	ComputeService() compute.ComputeServiceV2
	PortalService() portal.PortalServiceV2
	VolumeService() volume.VolumeServiceV2
}

type VLBGatewayV2 interface {
	LoadBalancerService() loadbalancer.LoadBalancerServiceV2
}

type VNetworkGatewayV1 interface {
	NetworkService() network.NetworkServiceV1
}

type VNetworkGatewayV2 interface {
	NetworkService() network.NetworkServiceV2
}

type VNetworkGatewayInternalV1 interface {
	NetworkService() network.NetworkServiceInternalV1
}

type VLBGatewayInternal interface {
	LoadBalancerService() loadbalancer.LoadBalancerServiceInternal
}

type VLBGateway interface {
	Internal() VLBGatewayInternal
	V2() VLBGatewayV2
	GetEndpoint() string
}

type VBackUpGateway any

type GLBGateway interface {
	V1() GLBGatewayV1
}

type GLBGatewayV1 interface {
	GLBService() glb.GLBServiceV1
}

type VDnsGateway interface {
	V1() VDnsGatewayV1
	Internal() VDnsGatewayInternal
	GetEndpoint() string
}

type VDnsGatewayV1 interface {
	DnsService() dns.VDnsServiceV1
}

type VDnsGatewayInternal interface {
	DnsService() dns.VDnsServiceInternal
}
