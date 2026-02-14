package glb

import (
	"github.com/dannyota/greennode-community-sdk/v2/greennode/entity"
	sdkerror "github.com/dannyota/greennode-community-sdk/v2/greennode/sdkerror"
	v1 "github.com/dannyota/greennode-community-sdk/v2/greennode/services/glb/v1"
)

type GLBServiceV1 interface {
	ListGlobalPools(popts v1.IListGlobalPoolsRequest) (*entity.ListGlobalPools, sdkerror.Error)
	CreateGlobalPool(popts v1.ICreateGlobalPoolRequest) (*entity.GlobalPool, sdkerror.Error)
	UpdateGlobalPool(popts v1.IUpdateGlobalPoolRequest) (*entity.GlobalPool, sdkerror.Error)
	DeleteGlobalPool(popts v1.IDeleteGlobalPoolRequest) sdkerror.Error

	ListGlobalPoolMembers(popts v1.IListGlobalPoolMembersRequest) (*entity.ListGlobalPoolMembers, sdkerror.Error)
	GetGlobalPoolMember(popts v1.IGetGlobalPoolMemberRequest) (*entity.GlobalPoolMember, sdkerror.Error)
	UpdateGlobalPoolMember(popts v1.IUpdateGlobalPoolMemberRequest) (*entity.GlobalPoolMember, sdkerror.Error)
	DeleteGlobalPoolMember(popts v1.IDeleteGlobalPoolMemberRequest) sdkerror.Error
	PatchGlobalPoolMembers(popts v1.IPatchGlobalPoolMembersRequest) sdkerror.Error

	ListGlobalListeners(popts v1.IListGlobalListenersRequest) (*entity.ListGlobalListeners, sdkerror.Error)
	CreateGlobalListener(popts v1.ICreateGlobalListenerRequest) (*entity.GlobalListener, sdkerror.Error)
	UpdateGlobalListener(popts v1.IUpdateGlobalListenerRequest) (*entity.GlobalListener, sdkerror.Error)
	DeleteGlobalListener(popts v1.IDeleteGlobalListenerRequest) sdkerror.Error
	GetGlobalListener(popts v1.IGetGlobalListenerRequest) (*entity.GlobalListener, sdkerror.Error)

	ListGlobalLoadBalancers(popts v1.IListGlobalLoadBalancersRequest) (*entity.ListGlobalLoadBalancers, sdkerror.Error)
	CreateGlobalLoadBalancer(popts v1.ICreateGlobalLoadBalancerRequest) (*entity.GlobalLoadBalancer, sdkerror.Error)
	DeleteGlobalLoadBalancer(popts v1.IDeleteGlobalLoadBalancerRequest) sdkerror.Error
	GetGlobalLoadBalancerById(popts v1.IGetGlobalLoadBalancerByIdRequest) (*entity.GlobalLoadBalancer, sdkerror.Error)

	ListGlobalPackages(popts v1.IListGlobalPackagesRequest) (*entity.ListGlobalPackages, sdkerror.Error)
	ListGlobalRegions(popts v1.IListGlobalRegionsRequest) (*entity.ListGlobalRegions, sdkerror.Error)
	GetGlobalLoadBalancerUsageHistories(popts v1.IGetGlobalLoadBalancerUsageHistoriesRequest) (*entity.ListGlobalLoadBalancerUsageHistories, sdkerror.Error)
}
