package glb

import (
	"github.com/dannyota/greennode-community-sdk/v2/greennode/entity"
	sdkerror "github.com/dannyota/greennode-community-sdk/v2/greennode/sdkerror"
	v1 "github.com/dannyota/greennode-community-sdk/v2/greennode/services/glb/v1"
)

type IGLBServiceV1 interface {
	ListGlobalPools(popts v1.IListGlobalPoolsRequest) (*entity.ListGlobalPools, sdkerror.IError)
	CreateGlobalPool(popts v1.ICreateGlobalPoolRequest) (*entity.GlobalPool, sdkerror.IError)
	UpdateGlobalPool(popts v1.IUpdateGlobalPoolRequest) (*entity.GlobalPool, sdkerror.IError)
	DeleteGlobalPool(popts v1.IDeleteGlobalPoolRequest) sdkerror.IError

	ListGlobalPoolMembers(popts v1.IListGlobalPoolMembersRequest) (*entity.ListGlobalPoolMembers, sdkerror.IError)
	GetGlobalPoolMember(popts v1.IGetGlobalPoolMemberRequest) (*entity.GlobalPoolMember, sdkerror.IError)
	UpdateGlobalPoolMember(popts v1.IUpdateGlobalPoolMemberRequest) (*entity.GlobalPoolMember, sdkerror.IError)
	DeleteGlobalPoolMember(popts v1.IDeleteGlobalPoolMemberRequest) sdkerror.IError
	PatchGlobalPoolMembers(popts v1.IPatchGlobalPoolMembersRequest) sdkerror.IError

	ListGlobalListeners(popts v1.IListGlobalListenersRequest) (*entity.ListGlobalListeners, sdkerror.IError)
	CreateGlobalListener(popts v1.ICreateGlobalListenerRequest) (*entity.GlobalListener, sdkerror.IError)
	UpdateGlobalListener(popts v1.IUpdateGlobalListenerRequest) (*entity.GlobalListener, sdkerror.IError)
	DeleteGlobalListener(popts v1.IDeleteGlobalListenerRequest) sdkerror.IError
	GetGlobalListener(popts v1.IGetGlobalListenerRequest) (*entity.GlobalListener, sdkerror.IError)

	ListGlobalLoadBalancers(popts v1.IListGlobalLoadBalancersRequest) (*entity.ListGlobalLoadBalancers, sdkerror.IError)
	CreateGlobalLoadBalancer(popts v1.ICreateGlobalLoadBalancerRequest) (*entity.GlobalLoadBalancer, sdkerror.IError)
	DeleteGlobalLoadBalancer(popts v1.IDeleteGlobalLoadBalancerRequest) sdkerror.IError
	GetGlobalLoadBalancerById(popts v1.IGetGlobalLoadBalancerByIdRequest) (*entity.GlobalLoadBalancer, sdkerror.IError)

	ListGlobalPackages(popts v1.IListGlobalPackagesRequest) (*entity.ListGlobalPackages, sdkerror.IError)
	ListGlobalRegions(popts v1.IListGlobalRegionsRequest) (*entity.ListGlobalRegions, sdkerror.IError)
	GetGlobalLoadBalancerUsageHistories(popts v1.IGetGlobalLoadBalancerUsageHistoriesRequest) (*entity.ListGlobalLoadBalancerUsageHistories, sdkerror.IError)
}
