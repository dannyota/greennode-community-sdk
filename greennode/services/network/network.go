package network

import (
	"github.com/dannyota/greennode-community-sdk/v2/greennode/client"
	"github.com/dannyota/greennode-community-sdk/v2/greennode/entity"
	networkv1 "github.com/dannyota/greennode-community-sdk/v2/greennode/services/network/v1"
	networkv2 "github.com/dannyota/greennode-community-sdk/v2/greennode/services/network/v2"
)

type NetworkServiceV1 interface {
	GetEndpointByID(ops *networkv1.GetEndpointByIDRequest) (*entity.Endpoint, error)
	CreateEndpoint(opts *networkv1.CreateEndpointRequest) (*entity.Endpoint, error)
	DeleteEndpointByID(opts *networkv1.DeleteEndpointByIDRequest) error
	ListEndpoints(opts *networkv1.ListEndpointsRequest) (*entity.ListEndpoints, error)
}

type NetworkServiceInternalV1 interface {
	ListTagsByEndpointID(opts *networkv1.ListTagsByEndpointIDRequest) (*entity.ListTags, error)
	CreateTagsWithEndpointID(opts *networkv1.CreateTagsWithEndpointIDRequest) error
	DeleteTagOfEndpoint(opts *networkv1.DeleteTagOfEndpointRequest) error
	UpdateTagValueOfEndpoint(opts *networkv1.UpdateTagValueOfEndpointRequest) error
	CreateEndpoint(opts *networkv1.CreateEndpointRequest) (*entity.Endpoint, error)
}

type NetworkOps interface {
	GetNetworkByID(opts *networkv2.GetNetworkByIDRequest) (*entity.Network, error)
	GetSubnetByID(opts *networkv2.GetSubnetByIDRequest) (*entity.Subnet, error)
	UpdateSubnetByID(opts *networkv2.UpdateSubnetByIDRequest) (*entity.Subnet, error)
}

type SecgroupOps interface {
	GetSecgroupByID(opts *networkv2.GetSecgroupByIDRequest) (*entity.Secgroup, error)
	CreateSecgroup(opts *networkv2.CreateSecgroupRequest) (*entity.Secgroup, error)
	ListSecgroup(opts *networkv2.ListSecgroupRequest) (*entity.ListSecgroups, error)
	DeleteSecgroupByID(opts *networkv2.DeleteSecgroupByIDRequest) error
	CreateSecgroupRule(opts *networkv2.CreateSecgroupRuleRequest) (*entity.SecgroupRule, error)
	DeleteSecgroupRuleByID(opts *networkv2.DeleteSecgroupRuleByIDRequest) error
	ListSecgroupRulesBySecgroupID(opts *networkv2.ListSecgroupRulesBySecgroupIDRequest) (*entity.ListSecgroupRules, error)
	ListAllServersBySecgroupID(opts *networkv2.ListAllServersBySecgroupIDRequest) (*entity.ListServers, error)
}

type AddressPairOps interface {
	GetAllAddressPairByVirtualSubnetID(opts *networkv2.GetAllAddressPairByVirtualSubnetIDRequest) ([]*entity.AddressPair, error)
	SetAddressPairInVirtualSubnet(opts *networkv2.SetAddressPairInVirtualSubnetRequest) (*entity.AddressPair, error)
	DeleteAddressPair(opts *networkv2.DeleteAddressPairRequest) error
	CreateAddressPair(opts *networkv2.CreateAddressPairRequest) (*entity.AddressPair, error)
}

type VirtualAddressOps interface {
	CreateVirtualAddressCrossProject(opts *networkv2.CreateVirtualAddressCrossProjectRequest) (*entity.VirtualAddress, error)
	DeleteVirtualAddressByID(opts *networkv2.DeleteVirtualAddressByIDRequest) error
	GetVirtualAddressByID(opts *networkv2.GetVirtualAddressByIDRequest) (*entity.VirtualAddress, error)
	ListAddressPairsByVirtualAddressID(opts *networkv2.ListAddressPairsByVirtualAddressIDRequest) (*entity.ListAddressPairs, error)
}

type NetworkServiceV2 interface {
	NetworkOps
	SecgroupOps
	AddressPairOps
	VirtualAddressOps
}

func NewNetworkServiceV2(svcClient client.ServiceClient) *networkv2.NetworkServiceV2 {
	return &networkv2.NetworkServiceV2{
		VserverClient: svcClient,
	}
}

func NewNetworkServiceV1(svcClient client.ServiceClient) *networkv1.NetworkServiceV1 {
	return &networkv1.NetworkServiceV1{
		VNetworkClient: svcClient,
	}
}

func NewNetworkServiceInternalV1(svcClient client.ServiceClient) *networkv1.NetworkServiceInternalV1 {
	return &networkv1.NetworkServiceInternalV1{
		VNetworkClient: svcClient,
	}
}
