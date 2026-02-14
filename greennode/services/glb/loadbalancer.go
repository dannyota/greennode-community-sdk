package glb

import (
	"github.com/dannyota/greennode-community-sdk/v2/greennode/client"
	"github.com/dannyota/greennode-community-sdk/v2/greennode/entity"
	sdkerror "github.com/dannyota/greennode-community-sdk/v2/greennode/sdkerror"
	v1 "github.com/dannyota/greennode-community-sdk/v2/greennode/services/glb/v1"
)

type GLBServiceV1 interface {
	ListGlobalPools(opts *v1.ListGlobalPoolsRequest) (*entity.ListGlobalPools, sdkerror.Error)
	CreateGlobalPool(opts *v1.CreateGlobalPoolRequest) (*entity.GlobalPool, sdkerror.Error)
	UpdateGlobalPool(opts *v1.UpdateGlobalPoolRequest) (*entity.GlobalPool, sdkerror.Error)
	DeleteGlobalPool(opts *v1.DeleteGlobalPoolRequest) sdkerror.Error

	ListGlobalPoolMembers(opts *v1.ListGlobalPoolMembersRequest) (*entity.ListGlobalPoolMembers, sdkerror.Error)
	GetGlobalPoolMember(opts *v1.GetGlobalPoolMemberRequest) (*entity.GlobalPoolMember, sdkerror.Error)
	UpdateGlobalPoolMember(opts *v1.UpdateGlobalPoolMemberRequest) (*entity.GlobalPoolMember, sdkerror.Error)
	DeleteGlobalPoolMember(opts *v1.DeleteGlobalPoolMemberRequest) sdkerror.Error
	PatchGlobalPoolMembers(opts *v1.PatchGlobalPoolMembersRequest) sdkerror.Error

	ListGlobalListeners(opts *v1.ListGlobalListenersRequest) (*entity.ListGlobalListeners, sdkerror.Error)
	CreateGlobalListener(opts *v1.CreateGlobalListenerRequest) (*entity.GlobalListener, sdkerror.Error)
	UpdateGlobalListener(opts *v1.UpdateGlobalListenerRequest) (*entity.GlobalListener, sdkerror.Error)
	DeleteGlobalListener(opts *v1.DeleteGlobalListenerRequest) sdkerror.Error
	GetGlobalListener(opts *v1.GetGlobalListenerRequest) (*entity.GlobalListener, sdkerror.Error)

	ListGlobalLoadBalancers(opts *v1.ListGlobalLoadBalancersRequest) (*entity.ListGlobalLoadBalancers, sdkerror.Error)
	CreateGlobalLoadBalancer(opts *v1.CreateGlobalLoadBalancerRequest) (*entity.GlobalLoadBalancer, sdkerror.Error)
	DeleteGlobalLoadBalancer(opts *v1.DeleteGlobalLoadBalancerRequest) sdkerror.Error
	GetGlobalLoadBalancerByID(opts *v1.GetGlobalLoadBalancerByIDRequest) (*entity.GlobalLoadBalancer, sdkerror.Error)

	ListGlobalPackages(opts *v1.ListGlobalPackagesRequest) (*entity.ListGlobalPackages, sdkerror.Error)
	ListGlobalRegions(opts *v1.ListGlobalRegionsRequest) (*entity.ListGlobalRegions, sdkerror.Error)
	GetGlobalLoadBalancerUsageHistories(opts *v1.GetGlobalLoadBalancerUsageHistoriesRequest) (*entity.ListGlobalLoadBalancerUsageHistories, sdkerror.Error)
}

func NewGLBServiceV1(svcClient client.ServiceClient) GLBServiceV1 {
	return &v1.GLBServiceV1{
		VLBClient: svcClient,
	}
}
