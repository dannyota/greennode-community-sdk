package network

import (
	"github.com/dannyota/greennode-community-sdk/v2/greennode/entity"
	sdkerror "github.com/dannyota/greennode-community-sdk/v2/greennode/sdkerror"
	networkv1 "github.com/dannyota/greennode-community-sdk/v2/greennode/services/network/v1"
	networkv2 "github.com/dannyota/greennode-community-sdk/v2/greennode/services/network/v2"
)

type NetworkServiceV1 interface {
	GetEndpointById(pops networkv1.IGetEndpointByIdRequest) (*entity.Endpoint, sdkerror.Error)
	CreateEndpoint(popts networkv1.ICreateEndpointRequest) (*entity.Endpoint, sdkerror.Error)
	DeleteEndpointById(popts networkv1.IDeleteEndpointByIdRequest) sdkerror.Error
	ListEndpoints(popts networkv1.IListEndpointsRequest) (*entity.ListEndpoints, sdkerror.Error)
}

type NetworkServiceInternalV1 interface {
	ListTagsByEndpointId(popts networkv1.IListTagsByEndpointIdRequest) (*entity.ListTags, sdkerror.Error)
	CreateTagsWithEndpointId(popts networkv1.ICreateTagsWithEndpointIdRequest) sdkerror.Error
	DeleteTagOfEndpoint(popts networkv1.IDeleteTagOfEndpointRequest) sdkerror.Error
	UpdateTagValueOfEndpoint(popts networkv1.IUpdateTagValueOfEndpointRequest) sdkerror.Error
	CreateEndpoint(popts networkv1.ICreateEndpointRequest) (*entity.Endpoint, sdkerror.Error)
}

type NetworkServiceV2 interface {
	// The group of Network APIs
	GetNetworkById(popts networkv2.IGetNetworkByIdRequest) (*entity.Network, sdkerror.Error)

	// The group of Secgroup APIs

	GetSecgroupById(popts networkv2.IGetSecgroupByIdRequest) (*entity.Secgroup, sdkerror.Error)
	CreateSecgroup(popts networkv2.ICreateSecgroupRequest) (*entity.Secgroup, sdkerror.Error)
	ListSecgroup(popts networkv2.IListSecgroupRequest) (*entity.ListSecgroups, sdkerror.Error)
	DeleteSecgroupById(popts networkv2.IDeleteSecgroupByIdRequest) sdkerror.Error

	// The group of SecgroupRule APIs

	CreateSecgroupRule(popts networkv2.ICreateSecgroupRuleRequest) (*entity.SecgroupRule, sdkerror.Error)
	DeleteSecgroupRuleById(popts networkv2.IDeleteSecgroupRuleByIdRequest) sdkerror.Error
	ListSecgroupRulesBySecgroupId(popts networkv2.IListSecgroupRulesBySecgroupIdRequest) (*entity.ListSecgroupRules, sdkerror.Error)

	// Subnet
	GetSubnetById(popts networkv2.IGetSubnetByIdRequest) (*entity.Subnet, sdkerror.Error)
	UpdateSubnetById(popts networkv2.IUpdateSubnetByIdRequest) (*entity.Subnet, sdkerror.Error)

	// Address Pair
	GetAllAddressPairByVirtualSubnetId(popts networkv2.IGetAllAddressPairByVirtualSubnetIdRequest) ([]*entity.AddressPair, sdkerror.Error)
	SetAddressPairInVirtualSubnet(popts networkv2.ISetAddressPairInVirtualSubnetRequest) (*entity.AddressPair, sdkerror.Error)
	DeleteAddressPair(popts networkv2.IDeleteAddressPairRequest) sdkerror.Error
	CreateAddressPair(popts networkv2.ICreateAddressPairRequest) (*entity.AddressPair, sdkerror.Error)

	// Servers
	ListAllServersBySecgroupId(popts networkv2.IListAllServersBySecgroupIdRequest) (*entity.ListServers, sdkerror.Error)

	// Virtual Address API group
	CreateVirtualAddressCrossProject(popts networkv2.ICreateVirtualAddressCrossProjectRequest) (*entity.VirtualAddress, sdkerror.Error)
	DeleteVirtualAddressById(popts networkv2.IDeleteVirtualAddressByIdRequest) sdkerror.Error
	GetVirtualAddressById(popts networkv2.IGetVirtualAddressByIdRequest) (*entity.VirtualAddress, sdkerror.Error)
	ListAddressPairsByVirtualAddressId(popts networkv2.IListAddressPairsByVirtualAddressIdRequest) (*entity.ListAddressPairs, sdkerror.Error)
}
