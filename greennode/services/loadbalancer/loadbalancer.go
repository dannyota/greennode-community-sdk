package loadbalancer

import (
	"github.com/dannyota/greennode-community-sdk/v2/greennode/client"
	"github.com/dannyota/greennode-community-sdk/v2/greennode/entity"
	"github.com/dannyota/greennode-community-sdk/v2/greennode/services/loadbalancer/inter"
	lbv2 "github.com/dannyota/greennode-community-sdk/v2/greennode/services/loadbalancer/v2"
)

type LoadBalancerOps interface {
	CreateLoadBalancer(opts *lbv2.CreateLoadBalancerRequest) (*entity.LoadBalancer, error)
	ResizeLoadBalancer(opts *lbv2.ResizeLoadBalancerRequest) (*entity.LoadBalancer, error)
	ResizeLoadBalancerByID(opts *lbv2.ResizeLoadBalancerByIDRequest) error
	ScaleLoadBalancer(opts *lbv2.ScaleLoadBalancerRequest) (*entity.LoadBalancer, error)
	ListLoadBalancers(opts *lbv2.ListLoadBalancersRequest) (*entity.ListLoadBalancers, error)
	GetLoadBalancerByID(opts *lbv2.GetLoadBalancerByIDRequest) (*entity.LoadBalancer, error)
	DeleteLoadBalancerByID(opts *lbv2.DeleteLoadBalancerByIDRequest) error
	ListLoadBalancerPackages(opts *lbv2.ListLoadBalancerPackagesRequest) (*entity.ListLoadBalancerPackages, error)
}

type PoolOps interface {
	CreatePool(opts *lbv2.CreatePoolRequest) (*entity.Pool, error)
	UpdatePool(opts *lbv2.UpdatePoolRequest) error
	GetPoolByID(opts *lbv2.GetPoolByIDRequest) (*entity.Pool, error)
	DeletePoolByID(opt *lbv2.DeletePoolByIDRequest) error
	ListPoolsByLoadBalancerID(opts *lbv2.ListPoolsByLoadBalancerIDRequest) (*entity.ListPools, error)
	GetPoolHealthMonitorByID(opts *lbv2.GetPoolHealthMonitorByIDRequest) (*entity.HealthMonitor, error)
	UpdatePoolMembers(opts *lbv2.UpdatePoolMembersRequest) error
	ListPoolMembers(opts *lbv2.ListPoolMembersRequest) (*entity.ListMembers, error)
}

type ListenerOps interface {
	CreateListener(opts *lbv2.CreateListenerRequest) (*entity.Listener, error)
	UpdateListener(opts *lbv2.UpdateListenerRequest) error
	GetListenerByID(opts *lbv2.GetListenerByIDRequest) (*entity.Listener, error)
	DeleteListenerByID(opts *lbv2.DeleteListenerByIDRequest) error
	ListListenersByLoadBalancerID(opts *lbv2.ListListenersByLoadBalancerIDRequest) (*entity.ListListeners, error)
}

type PolicyOps interface {
	CreatePolicy(opts *lbv2.CreatePolicyRequest) (*entity.Policy, error)
	UpdatePolicy(opts *lbv2.UpdatePolicyRequest) error
	GetPolicyByID(opts *lbv2.GetPolicyByIDRequest) (*entity.Policy, error)
	DeletePolicyByID(opts *lbv2.DeletePolicyByIDRequest) error
	ListPolicies(opts *lbv2.ListPoliciesRequest) (*entity.ListPolicies, error)
	ReorderPolicies(opts *lbv2.ReorderPoliciesRequest) error
}

type CertificateOps interface {
	CreateCertificate(opts *lbv2.CreateCertificateRequest) (*entity.Certificate, error)
	GetCertificateByID(opts *lbv2.GetCertificateByIDRequest) (*entity.Certificate, error)
	DeleteCertificateByID(opts *lbv2.DeleteCertificateByIDRequest) error
	ListCertificates(opts *lbv2.ListCertificatesRequest) (*entity.ListCertificates, error)
}

type TagOps interface {
	CreateTags(opts *lbv2.CreateTagsRequest) error
	UpdateTags(opts *lbv2.UpdateTagsRequest) error
	ListTags(opts *lbv2.ListTagsRequest) (*entity.ListTags, error)
}

type LoadBalancerServiceV2 interface {
	LoadBalancerOps
	PoolOps
	ListenerOps
	PolicyOps
	CertificateOps
	TagOps
}

type LoadBalancerServiceInternal interface {
	CreateLoadBalancer(opts *inter.CreateLoadBalancerRequest) (*entity.LoadBalancer, error)
}

func NewLoadBalancerServiceV2(lbSvcClient, serverSvcClient client.ServiceClient) *lbv2.LoadBalancerServiceV2 {
	return &lbv2.LoadBalancerServiceV2{
		VLBClient:     lbSvcClient,
		VServerClient: serverSvcClient,
	}
}

func NewLoadBalancerServiceInternal(svcClient client.ServiceClient) *inter.LoadBalancerServiceInternal {
	return &inter.LoadBalancerServiceInternal{
		VLBClient: svcClient,
	}
}
