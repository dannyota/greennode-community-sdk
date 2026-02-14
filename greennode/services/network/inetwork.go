package network

import (
	"github.com/dannyota/greennode-community-sdk/v2/greennode/entity"
	sdkerror "github.com/dannyota/greennode-community-sdk/v2/greennode/sdkerror"
	networkv1 "github.com/dannyota/greennode-community-sdk/v2/greennode/services/network/v1"
	networkv2 "github.com/dannyota/greennode-community-sdk/v2/greennode/services/network/v2"
)

type NetworkServiceV1 interface {
	GetEndpointById(ops networkv1.IGetEndpointByIdRequest) (*entity.Endpoint, sdkerror.Error)
	CreateEndpoint(opts networkv1.ICreateEndpointRequest) (*entity.Endpoint, sdkerror.Error)
	DeleteEndpointById(opts networkv1.IDeleteEndpointByIdRequest) sdkerror.Error
	ListEndpoints(opts networkv1.IListEndpointsRequest) (*entity.ListEndpoints, sdkerror.Error)
}

type NetworkServiceInternalV1 interface {
	ListTagsByEndpointId(opts networkv1.IListTagsByEndpointIdRequest) (*entity.ListTags, sdkerror.Error)
	CreateTagsWithEndpointId(opts networkv1.ICreateTagsWithEndpointIdRequest) sdkerror.Error
	DeleteTagOfEndpoint(opts networkv1.IDeleteTagOfEndpointRequest) sdkerror.Error
	UpdateTagValueOfEndpoint(opts networkv1.IUpdateTagValueOfEndpointRequest) sdkerror.Error
	CreateEndpoint(opts networkv1.ICreateEndpointRequest) (*entity.Endpoint, sdkerror.Error)
}

type NetworkServiceV2 interface {
	// The group of Network APIs
	GetNetworkById(opts networkv2.IGetNetworkByIdRequest) (*entity.Network, sdkerror.Error)

	// The group of Secgroup APIs

	GetSecgroupById(opts networkv2.IGetSecgroupByIdRequest) (*entity.Secgroup, sdkerror.Error)
	CreateSecgroup(opts networkv2.ICreateSecgroupRequest) (*entity.Secgroup, sdkerror.Error)
	ListSecgroup(opts networkv2.IListSecgroupRequest) (*entity.ListSecgroups, sdkerror.Error)
	DeleteSecgroupById(opts networkv2.IDeleteSecgroupByIdRequest) sdkerror.Error

	// The group of SecgroupRule APIs

	CreateSecgroupRule(opts networkv2.ICreateSecgroupRuleRequest) (*entity.SecgroupRule, sdkerror.Error)
	DeleteSecgroupRuleById(opts networkv2.IDeleteSecgroupRuleByIdRequest) sdkerror.Error
	ListSecgroupRulesBySecgroupId(opts networkv2.IListSecgroupRulesBySecgroupIdRequest) (*entity.ListSecgroupRules, sdkerror.Error)

	// Subnet
	GetSubnetById(opts networkv2.IGetSubnetByIdRequest) (*entity.Subnet, sdkerror.Error)
	UpdateSubnetById(opts networkv2.IUpdateSubnetByIdRequest) (*entity.Subnet, sdkerror.Error)

	// Address Pair
	GetAllAddressPairByVirtualSubnetId(opts networkv2.IGetAllAddressPairByVirtualSubnetIdRequest) ([]*entity.AddressPair, sdkerror.Error)
	SetAddressPairInVirtualSubnet(opts networkv2.ISetAddressPairInVirtualSubnetRequest) (*entity.AddressPair, sdkerror.Error)
	DeleteAddressPair(opts networkv2.IDeleteAddressPairRequest) sdkerror.Error
	CreateAddressPair(opts networkv2.ICreateAddressPairRequest) (*entity.AddressPair, sdkerror.Error)

	// Servers
	ListAllServersBySecgroupId(opts networkv2.IListAllServersBySecgroupIdRequest) (*entity.ListServers, sdkerror.Error)

	// Virtual Address API group
	CreateVirtualAddressCrossProject(opts networkv2.ICreateVirtualAddressCrossProjectRequest) (*entity.VirtualAddress, sdkerror.Error)
	DeleteVirtualAddressById(opts networkv2.IDeleteVirtualAddressByIdRequest) sdkerror.Error
	GetVirtualAddressById(opts networkv2.IGetVirtualAddressByIdRequest) (*entity.VirtualAddress, sdkerror.Error)
	ListAddressPairsByVirtualAddressId(opts networkv2.IListAddressPairsByVirtualAddressIdRequest) (*entity.ListAddressPairs, sdkerror.Error)
}
