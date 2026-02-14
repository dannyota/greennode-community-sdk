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

type IIamGateway interface {
	V2() IIamGatewayV2
}

type IIamGatewayV2 interface {
	IdentityService() identity.IIdentityServiceV2
}

type IVServerGateway interface {
	V1() IVServerGatewayV1
	InternalV1() IVServerGatewayInternalV1
	V2() IVServerGatewayV2
	// GetEndpoint returns the endpoint of the vServer service
	GetEndpoint() string
}

type IVNetworkGateway interface {
	V1() IVNetworkGatewayV1
	V2() IVNetworkGatewayV1
	InternalV1() IVNetworkGatewayInternalV1
	GetEndpoint() string
}

type IVServerGatewayV1 interface {
	PortalService() portal.IPortalServiceV1
	VolumeService() volume.IVolumeServiceV1
}

type IVServerGatewayInternalV1 interface {
	ServerService() server.IServerServiceInternalV1
}

type IVServerGatewayV2 interface {
	NetworkService() network.INetworkServiceV2
	ComputeService() compute.IComputeServiceV2
	PortalService() portal.IPortalServiceV2
	VolumeService() volume.IVolumeServiceV2
}

type IVLBGatewayV2 interface {
	LoadBalancerService() loadbalancer.ILoadBalancerServiceV2
}

type IVNetworkGatewayV1 interface {
	NetworkService() network.INetworkServiceV1
}

type IVNetworkGatewayV2 interface {
	NetworkService() network.INetworkServiceV2
}

type IVNetworkGatewayInternalV1 interface {
	NetworkService() network.INetworkServiceInternalV1
}

type IVLBGatewayInternal interface {
	LoadBalancerService() loadbalancer.ILoadBalancerServiceInternal
}

type IVLBGateway interface {
	Internal() IVLBGatewayInternal
	V2() IVLBGatewayV2
	GetEndpoint() string
}

type IVBackUpGateway interface{}

type IGLBGateway interface {
	V1() IGLBGatewayV1
}

type IGLBGatewayV1 interface {
	GLBService() glb.IGLBServiceV1
}

type IVDnsGateway interface {
	V1() IVDnsGatewayV1
	Internal() IVDnsGatewayInternal
	GetEndpoint() string
}

type IVDnsGatewayV1 interface {
	DnsService() dns.IVDnsServiceV1
}

type IVDnsGatewayInternal interface {
	DnsService() dns.IVDnsServiceInternal
}
