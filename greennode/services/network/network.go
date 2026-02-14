package network

import (
	"github.com/dannyota/greennode-community-sdk/v2/greennode/client"
	"github.com/dannyota/greennode-community-sdk/v2/greennode/entity"
	sdkerror "github.com/dannyota/greennode-community-sdk/v2/greennode/sdkerror"
	networkv1 "github.com/dannyota/greennode-community-sdk/v2/greennode/services/network/v1"
	networkv2 "github.com/dannyota/greennode-community-sdk/v2/greennode/services/network/v2"
)

type NetworkServiceV1 interface {
	GetEndpointByID(ops *networkv1.GetEndpointByIDRequest) (*entity.Endpoint, sdkerror.Error)
	CreateEndpoint(opts *networkv1.CreateEndpointRequest) (*entity.Endpoint, sdkerror.Error)
	DeleteEndpointByID(opts *networkv1.DeleteEndpointByIDRequest) sdkerror.Error
	ListEndpoints(opts *networkv1.ListEndpointsRequest) (*entity.ListEndpoints, sdkerror.Error)
}

type NetworkServiceInternalV1 interface {
	ListTagsByEndpointID(opts *networkv1.ListTagsByEndpointIDRequest) (*entity.ListTags, sdkerror.Error)
	CreateTagsWithEndpointID(opts *networkv1.CreateTagsWithEndpointIDRequest) sdkerror.Error
	DeleteTagOfEndpoint(opts *networkv1.DeleteTagOfEndpointRequest) sdkerror.Error
	UpdateTagValueOfEndpoint(opts *networkv1.UpdateTagValueOfEndpointRequest) sdkerror.Error
	CreateEndpoint(opts *networkv1.CreateEndpointRequest) (*entity.Endpoint, sdkerror.Error)
}

type NetworkServiceV2 interface {
	// The group of Network APIs
	GetNetworkByID(opts *networkv2.GetNetworkByIDRequest) (*entity.Network, sdkerror.Error)

	// The group of Secgroup APIs

	GetSecgroupByID(opts *networkv2.GetSecgroupByIDRequest) (*entity.Secgroup, sdkerror.Error)
	CreateSecgroup(opts *networkv2.CreateSecgroupRequest) (*entity.Secgroup, sdkerror.Error)
	ListSecgroup(opts *networkv2.ListSecgroupRequest) (*entity.ListSecgroups, sdkerror.Error)
	DeleteSecgroupByID(opts *networkv2.DeleteSecgroupByIDRequest) sdkerror.Error

	// The group of SecgroupRule APIs

	CreateSecgroupRule(opts *networkv2.CreateSecgroupRuleRequest) (*entity.SecgroupRule, sdkerror.Error)
	DeleteSecgroupRuleByID(opts *networkv2.DeleteSecgroupRuleByIDRequest) sdkerror.Error
	ListSecgroupRulesBySecgroupID(opts *networkv2.ListSecgroupRulesBySecgroupIDRequest) (*entity.ListSecgroupRules, sdkerror.Error)

	// Subnet
	GetSubnetByID(opts *networkv2.GetSubnetByIDRequest) (*entity.Subnet, sdkerror.Error)
	UpdateSubnetByID(opts *networkv2.UpdateSubnetByIDRequest) (*entity.Subnet, sdkerror.Error)

	// Address Pair
	GetAllAddressPairByVirtualSubnetID(opts *networkv2.GetAllAddressPairByVirtualSubnetIDRequest) ([]*entity.AddressPair, sdkerror.Error)
	SetAddressPairInVirtualSubnet(opts *networkv2.SetAddressPairInVirtualSubnetRequest) (*entity.AddressPair, sdkerror.Error)
	DeleteAddressPair(opts *networkv2.DeleteAddressPairRequest) sdkerror.Error
	CreateAddressPair(opts *networkv2.CreateAddressPairRequest) (*entity.AddressPair, sdkerror.Error)

	// Servers
	ListAllServersBySecgroupID(opts *networkv2.ListAllServersBySecgroupIDRequest) (*entity.ListServers, sdkerror.Error)

	// Virtual Address API group
	CreateVirtualAddressCrossProject(opts *networkv2.CreateVirtualAddressCrossProjectRequest) (*entity.VirtualAddress, sdkerror.Error)
	DeleteVirtualAddressByID(opts *networkv2.DeleteVirtualAddressByIDRequest) sdkerror.Error
	GetVirtualAddressByID(opts *networkv2.GetVirtualAddressByIDRequest) (*entity.VirtualAddress, sdkerror.Error)
	ListAddressPairsByVirtualAddressID(opts *networkv2.ListAddressPairsByVirtualAddressIDRequest) (*entity.ListAddressPairs, sdkerror.Error)
}

func NewNetworkServiceV2(svcClient client.ServiceClient) NetworkServiceV2 {
	return &networkv2.NetworkServiceV2{
		VserverClient: svcClient,
	}
}

func NewNetworkServiceV1(svcClient client.ServiceClient) NetworkServiceV1 {
	return &networkv1.NetworkServiceV1{
		VNetworkClient: svcClient,
	}
}

func NewNetworkServiceInternalV1(svcClient client.ServiceClient) NetworkServiceInternalV1 {
	return &networkv1.NetworkServiceInternalV1{
		VNetworkClient: svcClient,
	}
}
