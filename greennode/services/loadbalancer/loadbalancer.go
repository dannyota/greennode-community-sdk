package loadbalancer

import (
	"github.com/dannyota/greennode-community-sdk/v2/greennode/client"
	"github.com/dannyota/greennode-community-sdk/v2/greennode/entity"
	sdkerror "github.com/dannyota/greennode-community-sdk/v2/greennode/sdkerror"
	"github.com/dannyota/greennode-community-sdk/v2/greennode/services/loadbalancer/inter"
	lbv2 "github.com/dannyota/greennode-community-sdk/v2/greennode/services/loadbalancer/v2"
)

type LoadBalancerServiceV2 interface {
	CreateLoadBalancer(opts lbv2.ICreateLoadBalancerRequest) (*entity.LoadBalancer, sdkerror.Error)
	ResizeLoadBalancer(opts lbv2.IResizeLoadBalancerRequest) (*entity.LoadBalancer, sdkerror.Error)
	ListLoadBalancerPackages(opts lbv2.IListLoadBalancerPackagesRequest) (*entity.ListLoadBalancerPackages, sdkerror.Error)
	GetLoadBalancerByID(opts lbv2.IGetLoadBalancerByIDRequest) (*entity.LoadBalancer, sdkerror.Error)
	ListLoadBalancers(opts lbv2.IListLoadBalancersRequest) (*entity.ListLoadBalancers, sdkerror.Error)
	GetPoolHealthMonitorByID(opts lbv2.IGetPoolHealthMonitorByIDRequest) (*entity.HealthMonitor, sdkerror.Error)
	CreatePool(opts lbv2.ICreatePoolRequest) (*entity.Pool, sdkerror.Error)
	UpdatePool(opts lbv2.IUpdatePoolRequest) sdkerror.Error
	CreateListener(opts lbv2.ICreateListenerRequest) (*entity.Listener, sdkerror.Error)
	UpdateListener(opts lbv2.IUpdateListenerRequest) sdkerror.Error
	ListListenersByLoadBalancerID(opts lbv2.IListListenersByLoadBalancerIDRequest) (*entity.ListListeners, sdkerror.Error)
	ListPoolsByLoadBalancerID(opts lbv2.IListPoolsByLoadBalancerIDRequest) (*entity.ListPools, sdkerror.Error)
	UpdatePoolMembers(opts lbv2.IUpdatePoolMembersRequest) sdkerror.Error
	ListPoolMembers(opts lbv2.IListPoolMembersRequest) (*entity.ListMembers, sdkerror.Error)
	DeletePoolByID(opt lbv2.IDeletePoolByIDRequest) sdkerror.Error
	DeleteListenerByID(opts lbv2.IDeleteListenerByIDRequest) sdkerror.Error
	DeleteLoadBalancerByID(opts lbv2.IDeleteLoadBalancerByIDRequest) sdkerror.Error
	ListTags(opts lbv2.IListTagsRequest) (*entity.ListTags, sdkerror.Error)
	CreateTags(opts lbv2.ICreateTagsRequest) sdkerror.Error
	UpdateTags(opts lbv2.IUpdateTagsRequest) sdkerror.Error

	ListPolicies(opts lbv2.IListPoliciesRequest) (*entity.ListPolicies, sdkerror.Error)
	CreatePolicy(opts lbv2.ICreatePolicyRequest) (*entity.Policy, sdkerror.Error)
	GetPolicyByID(opts lbv2.IGetPolicyByIDRequest) (*entity.Policy, sdkerror.Error)
	UpdatePolicy(opts lbv2.IUpdatePolicyRequest) sdkerror.Error
	DeletePolicyByID(opts lbv2.IDeletePolicyByIDRequest) sdkerror.Error
	GetPoolByID(opts lbv2.IGetPoolByIDRequest) (*entity.Pool, sdkerror.Error)
	GetListenerByID(opts lbv2.IGetListenerByIDRequest) (*entity.Listener, sdkerror.Error)
	ResizeLoadBalancerByID(opts lbv2.IResizeLoadBalancerByIDRequest) sdkerror.Error
	ScaleLoadBalancer(opts lbv2.IScaleLoadBalancerRequest) (*entity.LoadBalancer, sdkerror.Error)
	ReorderPolicies(opts lbv2.IReorderPoliciesRequest) sdkerror.Error

	ListCertificates(opts lbv2.IListCertificatesRequest) (*entity.ListCertificates, sdkerror.Error)
	GetCertificateByID(opts lbv2.IGetCertificateByIDRequest) (*entity.Certificate, sdkerror.Error)
	CreateCertificate(opts lbv2.ICreateCertificateRequest) (*entity.Certificate, sdkerror.Error)
	DeleteCertificateByID(opts lbv2.IDeleteCertificateByIDRequest) sdkerror.Error
}

type LoadBalancerServiceInternal interface {
	CreateLoadBalancer(opts inter.ICreateLoadBalancerRequest) (*entity.LoadBalancer, sdkerror.Error)
}

func NewLoadBalancerServiceV2(lbSvcClient, serverSvcClient client.ServiceClient) LoadBalancerServiceV2 {
	return &lbv2.LoadBalancerServiceV2{
		VLBClient:     lbSvcClient,
		VServerClient: serverSvcClient,
	}
}

func NewLoadBalancerServiceInternal(svcClient client.ServiceClient) LoadBalancerServiceInternal {
	return &inter.LoadBalancerServiceInternal{
		VLBClient: svcClient,
	}
}
