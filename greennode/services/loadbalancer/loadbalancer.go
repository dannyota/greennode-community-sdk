package loadbalancer

import (
	"github.com/dannyota/greennode-community-sdk/v2/greennode/client"
	"github.com/dannyota/greennode-community-sdk/v2/greennode/entity"
	sdkerror "github.com/dannyota/greennode-community-sdk/v2/greennode/sdkerror"
	"github.com/dannyota/greennode-community-sdk/v2/greennode/services/loadbalancer/inter"
	lbv2 "github.com/dannyota/greennode-community-sdk/v2/greennode/services/loadbalancer/v2"
)

type LoadBalancerServiceV2 interface {
	CreateLoadBalancer(opts *lbv2.CreateLoadBalancerRequest) (*entity.LoadBalancer, sdkerror.Error)
	ResizeLoadBalancer(opts *lbv2.ResizeLoadBalancerRequest) (*entity.LoadBalancer, sdkerror.Error)
	ListLoadBalancerPackages(opts *lbv2.ListLoadBalancerPackagesRequest) (*entity.ListLoadBalancerPackages, sdkerror.Error)
	GetLoadBalancerByID(opts *lbv2.GetLoadBalancerByIDRequest) (*entity.LoadBalancer, sdkerror.Error)
	ListLoadBalancers(opts *lbv2.ListLoadBalancersRequest) (*entity.ListLoadBalancers, sdkerror.Error)
	GetPoolHealthMonitorByID(opts *lbv2.GetPoolHealthMonitorByIDRequest) (*entity.HealthMonitor, sdkerror.Error)
	CreatePool(opts *lbv2.CreatePoolRequest) (*entity.Pool, sdkerror.Error)
	UpdatePool(opts *lbv2.UpdatePoolRequest) sdkerror.Error
	CreateListener(opts *lbv2.CreateListenerRequest) (*entity.Listener, sdkerror.Error)
	UpdateListener(opts *lbv2.UpdateListenerRequest) sdkerror.Error
	ListListenersByLoadBalancerID(opts *lbv2.ListListenersByLoadBalancerIDRequest) (*entity.ListListeners, sdkerror.Error)
	ListPoolsByLoadBalancerID(opts *lbv2.ListPoolsByLoadBalancerIDRequest) (*entity.ListPools, sdkerror.Error)
	UpdatePoolMembers(opts *lbv2.UpdatePoolMembersRequest) sdkerror.Error
	ListPoolMembers(opts *lbv2.ListPoolMembersRequest) (*entity.ListMembers, sdkerror.Error)
	DeletePoolByID(opt *lbv2.DeletePoolByIDRequest) sdkerror.Error
	DeleteListenerByID(opts *lbv2.DeleteListenerByIDRequest) sdkerror.Error
	DeleteLoadBalancerByID(opts *lbv2.DeleteLoadBalancerByIDRequest) sdkerror.Error
	ListTags(opts *lbv2.ListTagsRequest) (*entity.ListTags, sdkerror.Error)
	CreateTags(opts *lbv2.CreateTagsRequest) sdkerror.Error
	UpdateTags(opts *lbv2.UpdateTagsRequest) sdkerror.Error

	ListPolicies(opts *lbv2.ListPoliciesRequest) (*entity.ListPolicies, sdkerror.Error)
	CreatePolicy(opts *lbv2.CreatePolicyRequest) (*entity.Policy, sdkerror.Error)
	GetPolicyByID(opts *lbv2.GetPolicyByIDRequest) (*entity.Policy, sdkerror.Error)
	UpdatePolicy(opts *lbv2.UpdatePolicyRequest) sdkerror.Error
	DeletePolicyByID(opts *lbv2.DeletePolicyByIDRequest) sdkerror.Error
	GetPoolByID(opts *lbv2.GetPoolByIDRequest) (*entity.Pool, sdkerror.Error)
	GetListenerByID(opts *lbv2.GetListenerByIDRequest) (*entity.Listener, sdkerror.Error)
	ResizeLoadBalancerByID(opts *lbv2.ResizeLoadBalancerByIDRequest) sdkerror.Error
	ScaleLoadBalancer(opts *lbv2.ScaleLoadBalancerRequest) (*entity.LoadBalancer, sdkerror.Error)
	ReorderPolicies(opts *lbv2.ReorderPoliciesRequest) sdkerror.Error

	ListCertificates(opts *lbv2.ListCertificatesRequest) (*entity.ListCertificates, sdkerror.Error)
	GetCertificateByID(opts *lbv2.GetCertificateByIDRequest) (*entity.Certificate, sdkerror.Error)
	CreateCertificate(opts *lbv2.CreateCertificateRequest) (*entity.Certificate, sdkerror.Error)
	DeleteCertificateByID(opts *lbv2.DeleteCertificateByIDRequest) sdkerror.Error
}

type LoadBalancerServiceInternal interface {
	CreateLoadBalancer(opts *inter.CreateLoadBalancerRequest) (*entity.LoadBalancer, sdkerror.Error)
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
