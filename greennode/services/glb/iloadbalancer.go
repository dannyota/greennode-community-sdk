package glb

import (
	"github.com/dannyota/greennode-community-sdk/v2/greennode/entity"
	sdkerror "github.com/dannyota/greennode-community-sdk/v2/greennode/sdkerror"
	v1 "github.com/dannyota/greennode-community-sdk/v2/greennode/services/glb/v1"
)

type GLBServiceV1 interface {
	ListGlobalPools(opts v1.IListGlobalPoolsRequest) (*entity.ListGlobalPools, sdkerror.Error)
	CreateGlobalPool(opts v1.ICreateGlobalPoolRequest) (*entity.GlobalPool, sdkerror.Error)
	UpdateGlobalPool(opts v1.IUpdateGlobalPoolRequest) (*entity.GlobalPool, sdkerror.Error)
	DeleteGlobalPool(opts v1.IDeleteGlobalPoolRequest) sdkerror.Error

	ListGlobalPoolMembers(opts v1.IListGlobalPoolMembersRequest) (*entity.ListGlobalPoolMembers, sdkerror.Error)
	GetGlobalPoolMember(opts v1.IGetGlobalPoolMemberRequest) (*entity.GlobalPoolMember, sdkerror.Error)
	UpdateGlobalPoolMember(opts v1.IUpdateGlobalPoolMemberRequest) (*entity.GlobalPoolMember, sdkerror.Error)
	DeleteGlobalPoolMember(opts v1.IDeleteGlobalPoolMemberRequest) sdkerror.Error
	PatchGlobalPoolMembers(opts v1.IPatchGlobalPoolMembersRequest) sdkerror.Error

	ListGlobalListeners(opts v1.IListGlobalListenersRequest) (*entity.ListGlobalListeners, sdkerror.Error)
	CreateGlobalListener(opts v1.ICreateGlobalListenerRequest) (*entity.GlobalListener, sdkerror.Error)
	UpdateGlobalListener(opts v1.IUpdateGlobalListenerRequest) (*entity.GlobalListener, sdkerror.Error)
	DeleteGlobalListener(opts v1.IDeleteGlobalListenerRequest) sdkerror.Error
	GetGlobalListener(opts v1.IGetGlobalListenerRequest) (*entity.GlobalListener, sdkerror.Error)

	ListGlobalLoadBalancers(opts v1.IListGlobalLoadBalancersRequest) (*entity.ListGlobalLoadBalancers, sdkerror.Error)
	CreateGlobalLoadBalancer(opts v1.ICreateGlobalLoadBalancerRequest) (*entity.GlobalLoadBalancer, sdkerror.Error)
	DeleteGlobalLoadBalancer(opts v1.IDeleteGlobalLoadBalancerRequest) sdkerror.Error
	GetGlobalLoadBalancerByID(opts v1.IGetGlobalLoadBalancerByIDRequest) (*entity.GlobalLoadBalancer, sdkerror.Error)

	ListGlobalPackages(opts v1.IListGlobalPackagesRequest) (*entity.ListGlobalPackages, sdkerror.Error)
	ListGlobalRegions(opts v1.IListGlobalRegionsRequest) (*entity.ListGlobalRegions, sdkerror.Error)
	GetGlobalLoadBalancerUsageHistories(opts v1.IGetGlobalLoadBalancerUsageHistoriesRequest) (*entity.ListGlobalLoadBalancerUsageHistories, sdkerror.Error)
}
