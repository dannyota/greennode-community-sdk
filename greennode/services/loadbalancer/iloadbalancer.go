package loadbalancer

import (
	"github.com/dannyota/greennode-community-sdk/v2/greennode/entity"
	sdkerror "github.com/dannyota/greennode-community-sdk/v2/greennode/sdkerror"
	"github.com/dannyota/greennode-community-sdk/v2/greennode/services/loadbalancer/inter"
	lbv2 "github.com/dannyota/greennode-community-sdk/v2/greennode/services/loadbalancer/v2"
)

type LoadBalancerServiceV2 interface {
	CreateLoadBalancer(opts lbv2.ICreateLoadBalancerRequest) (*entity.LoadBalancer, sdkerror.Error)
	ResizeLoadBalancer(opts lbv2.IResizeLoadBalancerRequest) (*entity.LoadBalancer, sdkerror.Error)
	ListLoadBalancerPackages(opts lbv2.IListLoadBalancerPackagesRequest) (*entity.ListLoadBalancerPackages, sdkerror.Error)
	GetLoadBalancerById(opts lbv2.IGetLoadBalancerByIdRequest) (*entity.LoadBalancer, sdkerror.Error)
	ListLoadBalancers(opts lbv2.IListLoadBalancersRequest) (*entity.ListLoadBalancers, sdkerror.Error)
	GetPoolHealthMonitorById(opts lbv2.IGetPoolHealthMonitorByIdRequest) (*entity.HealthMonitor, sdkerror.Error)
	CreatePool(opts lbv2.ICreatePoolRequest) (*entity.Pool, sdkerror.Error)
	UpdatePool(opts lbv2.IUpdatePoolRequest) sdkerror.Error
	CreateListener(opts lbv2.ICreateListenerRequest) (*entity.Listener, sdkerror.Error)
	UpdateListener(opts lbv2.IUpdateListenerRequest) sdkerror.Error
	ListListenersByLoadBalancerId(opts lbv2.IListListenersByLoadBalancerIdRequest) (*entity.ListListeners, sdkerror.Error)
	ListPoolsByLoadBalancerId(opts lbv2.IListPoolsByLoadBalancerIdRequest) (*entity.ListPools, sdkerror.Error)
	UpdatePoolMembers(opts lbv2.IUpdatePoolMembersRequest) sdkerror.Error
	ListPoolMembers(opts lbv2.IListPoolMembersRequest) (*entity.ListMembers, sdkerror.Error)
	DeletePoolById(opt lbv2.IDeletePoolByIdRequest) sdkerror.Error
	DeleteListenerById(opts lbv2.IDeleteListenerByIdRequest) sdkerror.Error
	DeleteLoadBalancerById(opts lbv2.IDeleteLoadBalancerByIdRequest) sdkerror.Error
	ListTags(opts lbv2.IListTagsRequest) (*entity.ListTags, sdkerror.Error)
	CreateTags(opts lbv2.ICreateTagsRequest) sdkerror.Error
	UpdateTags(opts lbv2.IUpdateTagsRequest) sdkerror.Error

	ListPolicies(opts lbv2.IListPoliciesRequest) (*entity.ListPolicies, sdkerror.Error)
	CreatePolicy(opts lbv2.ICreatePolicyRequest) (*entity.Policy, sdkerror.Error)
	GetPolicyById(opts lbv2.IGetPolicyByIdRequest) (*entity.Policy, sdkerror.Error)
	UpdatePolicy(opts lbv2.IUpdatePolicyRequest) sdkerror.Error
	DeletePolicyById(opts lbv2.IDeletePolicyByIdRequest) sdkerror.Error
	GetPoolById(opts lbv2.IGetPoolByIdRequest) (*entity.Pool, sdkerror.Error)
	GetListenerById(opts lbv2.IGetListenerByIdRequest) (*entity.Listener, sdkerror.Error)
	ResizeLoadBalancerById(opts lbv2.IResizeLoadBalancerByIdRequest) sdkerror.Error
	ScaleLoadBalancer(opts lbv2.IScaleLoadBalancerRequest) (*entity.LoadBalancer, sdkerror.Error)
	ReorderPolicies(opts lbv2.IReorderPoliciesRequest) sdkerror.Error

	ListCertificates(opts lbv2.IListCertificatesRequest) (*entity.ListCertificates, sdkerror.Error)
	GetCertificateById(opts lbv2.IGetCertificateByIdRequest) (*entity.Certificate, sdkerror.Error)
	CreateCertificate(opts lbv2.ICreateCertificateRequest) (*entity.Certificate, sdkerror.Error)
	DeleteCertificateById(opts lbv2.IDeleteCertificateByIdRequest) sdkerror.Error
}

type LoadBalancerServiceInternal interface {
	CreateLoadBalancer(opts inter.ICreateLoadBalancerRequest) (*entity.LoadBalancer, sdkerror.Error)
}
