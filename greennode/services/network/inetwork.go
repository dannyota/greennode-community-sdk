package network

import (
	"github.com/dannyota/greennode-community-sdk/v2/greennode/entity"
	sdkerror "github.com/dannyota/greennode-community-sdk/v2/greennode/sdk_error"
	networkv1 "github.com/dannyota/greennode-community-sdk/v2/greennode/services/network/v1"
	networkv2 "github.com/dannyota/greennode-community-sdk/v2/greennode/services/network/v2"
)

type INetworkServiceV1 interface {
	GetEndpointById(pops networkv1.IGetEndpointByIdRequest) (*entity.Endpoint, sdkerror.IError)
	CreateEndpoint(popts networkv1.ICreateEndpointRequest) (*entity.Endpoint, sdkerror.IError)
	DeleteEndpointById(popts networkv1.IDeleteEndpointByIdRequest) sdkerror.IError
	ListEndpoints(popts networkv1.IListEndpointsRequest) (*entity.ListEndpoints, sdkerror.IError)
}

type INetworkServiceInternalV1 interface {
	ListTagsByEndpointId(popts networkv1.IListTagsByEndpointIdRequest) (*entity.ListTags, sdkerror.IError)
	CreateTagsWithEndpointId(popts networkv1.ICreateTagsWithEndpointIdRequest) sdkerror.IError
	DeleteTagOfEndpoint(popts networkv1.IDeleteTagOfEndpointRequest) sdkerror.IError
	UpdateTagValueOfEndpoint(popts networkv1.IUpdateTagValueOfEndpointRequest) sdkerror.IError
	CreateEndpoint(popts networkv1.ICreateEndpointRequest) (*entity.Endpoint, sdkerror.IError)
}

type INetworkServiceV2 interface {
	// The group of Network APIs
	GetNetworkById(popts networkv2.IGetNetworkByIdRequest) (*entity.Network, sdkerror.IError)

	// The group of Secgroup APIs

	GetSecgroupById(popts networkv2.IGetSecgroupByIdRequest) (*entity.Secgroup, sdkerror.IError)
	CreateSecgroup(popts networkv2.ICreateSecgroupRequest) (*entity.Secgroup, sdkerror.IError)
	ListSecgroup(popts networkv2.IListSecgroupRequest) (*entity.ListSecgroups, sdkerror.IError)
	DeleteSecgroupById(popts networkv2.IDeleteSecgroupByIdRequest) sdkerror.IError

	// The group of SecgroupRule APIs

	CreateSecgroupRule(popts networkv2.ICreateSecgroupRuleRequest) (*entity.SecgroupRule, sdkerror.IError)
	DeleteSecgroupRuleById(popts networkv2.IDeleteSecgroupRuleByIdRequest) sdkerror.IError
	ListSecgroupRulesBySecgroupId(popts networkv2.IListSecgroupRulesBySecgroupIdRequest) (*entity.ListSecgroupRules, sdkerror.IError)

	// Subnet
	GetSubnetById(popts networkv2.IGetSubnetByIdRequest) (*entity.Subnet, sdkerror.IError)
	UpdateSubnetById(popts networkv2.IUpdateSubnetByIdRequest) (*entity.Subnet, sdkerror.IError)

	// Address Pair
	GetAllAddressPairByVirtualSubnetId(popts networkv2.IGetAllAddressPairByVirtualSubnetIdRequest) ([]*entity.AddressPair, sdkerror.IError)
	SetAddressPairInVirtualSubnet(popts networkv2.ISetAddressPairInVirtualSubnetRequest) (*entity.AddressPair, sdkerror.IError)
	DeleteAddressPair(popts networkv2.IDeleteAddressPairRequest) sdkerror.IError
	CreateAddressPair(popts networkv2.ICreateAddressPairRequest) (*entity.AddressPair, sdkerror.IError)

	// Servers
	ListAllServersBySecgroupId(popts networkv2.IListAllServersBySecgroupIdRequest) (*entity.ListServers, sdkerror.IError)

	// Virtual Address API group
	CreateVirtualAddressCrossProject(popts networkv2.ICreateVirtualAddressCrossProjectRequest) (*entity.VirtualAddress, sdkerror.IError)
	DeleteVirtualAddressById(popts networkv2.IDeleteVirtualAddressByIdRequest) sdkerror.IError
	GetVirtualAddressById(popts networkv2.IGetVirtualAddressByIdRequest) (*entity.VirtualAddress, sdkerror.IError)
	ListAddressPairsByVirtualAddressId(popts networkv2.IListAddressPairsByVirtualAddressIdRequest) (*entity.ListAddressPairs, sdkerror.IError)
}
