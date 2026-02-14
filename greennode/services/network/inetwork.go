package network

import (
	"github.com/dannyota/greennode-community-sdk/v2/greennode/entity"
	sdkerror "github.com/dannyota/greennode-community-sdk/v2/greennode/sdkerror"
	networkv1 "github.com/dannyota/greennode-community-sdk/v2/greennode/services/network/v1"
	networkv2 "github.com/dannyota/greennode-community-sdk/v2/greennode/services/network/v2"
)

type NetworkServiceV1 interface {
	GetEndpointByID(ops networkv1.IGetEndpointByIDRequest) (*entity.Endpoint, sdkerror.Error)
	CreateEndpoint(opts networkv1.ICreateEndpointRequest) (*entity.Endpoint, sdkerror.Error)
	DeleteEndpointByID(opts networkv1.IDeleteEndpointByIDRequest) sdkerror.Error
	ListEndpoints(opts networkv1.IListEndpointsRequest) (*entity.ListEndpoints, sdkerror.Error)
}

type NetworkServiceInternalV1 interface {
	ListTagsByEndpointID(opts networkv1.IListTagsByEndpointIDRequest) (*entity.ListTags, sdkerror.Error)
	CreateTagsWithEndpointID(opts networkv1.ICreateTagsWithEndpointIDRequest) sdkerror.Error
	DeleteTagOfEndpoint(opts networkv1.IDeleteTagOfEndpointRequest) sdkerror.Error
	UpdateTagValueOfEndpoint(opts networkv1.IUpdateTagValueOfEndpointRequest) sdkerror.Error
	CreateEndpoint(opts networkv1.ICreateEndpointRequest) (*entity.Endpoint, sdkerror.Error)
}

type NetworkServiceV2 interface {
	// The group of Network APIs
	GetNetworkByID(opts networkv2.IGetNetworkByIDRequest) (*entity.Network, sdkerror.Error)

	// The group of Secgroup APIs

	GetSecgroupByID(opts networkv2.IGetSecgroupByIDRequest) (*entity.Secgroup, sdkerror.Error)
	CreateSecgroup(opts networkv2.ICreateSecgroupRequest) (*entity.Secgroup, sdkerror.Error)
	ListSecgroup(opts networkv2.IListSecgroupRequest) (*entity.ListSecgroups, sdkerror.Error)
	DeleteSecgroupByID(opts networkv2.IDeleteSecgroupByIDRequest) sdkerror.Error

	// The group of SecgroupRule APIs

	CreateSecgroupRule(opts networkv2.ICreateSecgroupRuleRequest) (*entity.SecgroupRule, sdkerror.Error)
	DeleteSecgroupRuleByID(opts networkv2.IDeleteSecgroupRuleByIDRequest) sdkerror.Error
	ListSecgroupRulesBySecgroupID(opts networkv2.IListSecgroupRulesBySecgroupIDRequest) (*entity.ListSecgroupRules, sdkerror.Error)

	// Subnet
	GetSubnetByID(opts networkv2.IGetSubnetByIDRequest) (*entity.Subnet, sdkerror.Error)
	UpdateSubnetByID(opts networkv2.IUpdateSubnetByIDRequest) (*entity.Subnet, sdkerror.Error)

	// Address Pair
	GetAllAddressPairByVirtualSubnetID(opts networkv2.IGetAllAddressPairByVirtualSubnetIDRequest) ([]*entity.AddressPair, sdkerror.Error)
	SetAddressPairInVirtualSubnet(opts networkv2.ISetAddressPairInVirtualSubnetRequest) (*entity.AddressPair, sdkerror.Error)
	DeleteAddressPair(opts networkv2.IDeleteAddressPairRequest) sdkerror.Error
	CreateAddressPair(opts networkv2.ICreateAddressPairRequest) (*entity.AddressPair, sdkerror.Error)

	// Servers
	ListAllServersBySecgroupID(opts networkv2.IListAllServersBySecgroupIDRequest) (*entity.ListServers, sdkerror.Error)

	// Virtual Address API group
	CreateVirtualAddressCrossProject(opts networkv2.ICreateVirtualAddressCrossProjectRequest) (*entity.VirtualAddress, sdkerror.Error)
	DeleteVirtualAddressByID(opts networkv2.IDeleteVirtualAddressByIDRequest) sdkerror.Error
	GetVirtualAddressByID(opts networkv2.IGetVirtualAddressByIDRequest) (*entity.VirtualAddress, sdkerror.Error)
	ListAddressPairsByVirtualAddressID(opts networkv2.IListAddressPairsByVirtualAddressIDRequest) (*entity.ListAddressPairs, sdkerror.Error)
}
