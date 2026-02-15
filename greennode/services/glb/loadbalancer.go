package glb

import (
	"github.com/dannyota/greennode-community-sdk/v2/greennode/client"
	"github.com/dannyota/greennode-community-sdk/v2/greennode/entity"
	v1 "github.com/dannyota/greennode-community-sdk/v2/greennode/services/glb/v1"
)

type GlobalPoolOps interface {
	ListGlobalPools(opts *v1.ListGlobalPoolsRequest) (*entity.ListGlobalPools, error)
	CreateGlobalPool(opts *v1.CreateGlobalPoolRequest) (*entity.GlobalPool, error)
	UpdateGlobalPool(opts *v1.UpdateGlobalPoolRequest) (*entity.GlobalPool, error)
	DeleteGlobalPool(opts *v1.DeleteGlobalPoolRequest) error
	ListGlobalPoolMembers(opts *v1.ListGlobalPoolMembersRequest) (*entity.ListGlobalPoolMembers, error)
	GetGlobalPoolMember(opts *v1.GetGlobalPoolMemberRequest) (*entity.GlobalPoolMember, error)
	UpdateGlobalPoolMember(opts *v1.UpdateGlobalPoolMemberRequest) (*entity.GlobalPoolMember, error)
	DeleteGlobalPoolMember(opts *v1.DeleteGlobalPoolMemberRequest) error
	PatchGlobalPoolMembers(opts *v1.PatchGlobalPoolMembersRequest) error
}

type GlobalListenerOps interface {
	ListGlobalListeners(opts *v1.ListGlobalListenersRequest) (*entity.ListGlobalListeners, error)
	CreateGlobalListener(opts *v1.CreateGlobalListenerRequest) (*entity.GlobalListener, error)
	UpdateGlobalListener(opts *v1.UpdateGlobalListenerRequest) (*entity.GlobalListener, error)
	DeleteGlobalListener(opts *v1.DeleteGlobalListenerRequest) error
	GetGlobalListener(opts *v1.GetGlobalListenerRequest) (*entity.GlobalListener, error)
}

type GlobalLoadBalancerOps interface {
	ListGlobalLoadBalancers(opts *v1.ListGlobalLoadBalancersRequest) (*entity.ListGlobalLoadBalancers, error)
	CreateGlobalLoadBalancer(opts *v1.CreateGlobalLoadBalancerRequest) (*entity.GlobalLoadBalancer, error)
	DeleteGlobalLoadBalancer(opts *v1.DeleteGlobalLoadBalancerRequest) error
	GetGlobalLoadBalancerByID(opts *v1.GetGlobalLoadBalancerByIDRequest) (*entity.GlobalLoadBalancer, error)
	ListGlobalPackages(opts *v1.ListGlobalPackagesRequest) (*entity.ListGlobalPackages, error)
	ListGlobalRegions(opts *v1.ListGlobalRegionsRequest) (*entity.ListGlobalRegions, error)
	GetGlobalLoadBalancerUsageHistories(opts *v1.GetGlobalLoadBalancerUsageHistoriesRequest) (*entity.ListGlobalLoadBalancerUsageHistories, error)
}

type GLBServiceV1 interface {
	GlobalPoolOps
	GlobalListenerOps
	GlobalLoadBalancerOps
}

func NewGLBServiceV1(svcClient client.ServiceClient) *v1.GLBServiceV1 {
	return &v1.GLBServiceV1{
		VLBClient: svcClient,
	}
}
