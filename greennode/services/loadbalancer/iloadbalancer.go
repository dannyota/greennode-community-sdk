package loadbalancer

import (
	"github.com/dannyota/greennode-community-sdk/v2/greennode/entity"
	sdkerror "github.com/dannyota/greennode-community-sdk/v2/greennode/sdkerror"
	"github.com/dannyota/greennode-community-sdk/v2/greennode/services/loadbalancer/inter"
	lbv2 "github.com/dannyota/greennode-community-sdk/v2/greennode/services/loadbalancer/v2"
)

type LoadBalancerServiceV2 interface {
	CreateLoadBalancer(popts lbv2.ICreateLoadBalancerRequest) (*entity.LoadBalancer, sdkerror.Error)
	ResizeLoadBalancer(popts lbv2.IResizeLoadBalancerRequest) (*entity.LoadBalancer, sdkerror.Error)
	ListLoadBalancerPackages(popts lbv2.IListLoadBalancerPackagesRequest) (*entity.ListLoadBalancerPackages, sdkerror.Error)
	GetLoadBalancerById(popts lbv2.IGetLoadBalancerByIdRequest) (*entity.LoadBalancer, sdkerror.Error)
	ListLoadBalancers(popts lbv2.IListLoadBalancersRequest) (*entity.ListLoadBalancers, sdkerror.Error)
	GetPoolHealthMonitorById(popts lbv2.IGetPoolHealthMonitorByIdRequest) (*entity.HealthMonitor, sdkerror.Error)
	CreatePool(popts lbv2.ICreatePoolRequest) (*entity.Pool, sdkerror.Error)
	UpdatePool(popts lbv2.IUpdatePoolRequest) sdkerror.Error
	CreateListener(popts lbv2.ICreateListenerRequest) (*entity.Listener, sdkerror.Error)
	UpdateListener(popts lbv2.IUpdateListenerRequest) sdkerror.Error
	ListListenersByLoadBalancerId(popts lbv2.IListListenersByLoadBalancerIdRequest) (*entity.ListListeners, sdkerror.Error)
	ListPoolsByLoadBalancerId(popts lbv2.IListPoolsByLoadBalancerIdRequest) (*entity.ListPools, sdkerror.Error)
	UpdatePoolMembers(popts lbv2.IUpdatePoolMembersRequest) sdkerror.Error
	ListPoolMembers(popts lbv2.IListPoolMembersRequest) (*entity.ListMembers, sdkerror.Error)
	DeletePoolById(popt lbv2.IDeletePoolByIdRequest) sdkerror.Error
	DeleteListenerById(popts lbv2.IDeleteListenerByIdRequest) sdkerror.Error
	DeleteLoadBalancerById(popts lbv2.IDeleteLoadBalancerByIdRequest) sdkerror.Error
	ListTags(popts lbv2.IListTagsRequest) (*entity.ListTags, sdkerror.Error)
	CreateTags(popts lbv2.ICreateTagsRequest) sdkerror.Error
	UpdateTags(popts lbv2.IUpdateTagsRequest) sdkerror.Error

	ListPolicies(popts lbv2.IListPoliciesRequest) (*entity.ListPolicies, sdkerror.Error)
	CreatePolicy(popts lbv2.ICreatePolicyRequest) (*entity.Policy, sdkerror.Error)
	GetPolicyById(popts lbv2.IGetPolicyByIdRequest) (*entity.Policy, sdkerror.Error)
	UpdatePolicy(popts lbv2.IUpdatePolicyRequest) sdkerror.Error
	DeletePolicyById(popts lbv2.IDeletePolicyByIdRequest) sdkerror.Error
	GetPoolById(popts lbv2.IGetPoolByIdRequest) (*entity.Pool, sdkerror.Error)
	GetListenerById(popts lbv2.IGetListenerByIdRequest) (*entity.Listener, sdkerror.Error)
	ResizeLoadBalancerById(popts lbv2.IResizeLoadBalancerByIdRequest) sdkerror.Error
	ScaleLoadBalancer(popts lbv2.IScaleLoadBalancerRequest) (*entity.LoadBalancer, sdkerror.Error)
	ReorderPolicies(popts lbv2.IReorderPoliciesRequest) sdkerror.Error

	ListCertificates(popts lbv2.IListCertificatesRequest) (*entity.ListCertificates, sdkerror.Error)
	GetCertificateById(popts lbv2.IGetCertificateByIdRequest) (*entity.Certificate, sdkerror.Error)
	CreateCertificate(popts lbv2.ICreateCertificateRequest) (*entity.Certificate, sdkerror.Error)
	DeleteCertificateById(popts lbv2.IDeleteCertificateByIdRequest) sdkerror.Error
}

type LoadBalancerServiceInternal interface {
	CreateLoadBalancer(popts inter.ICreateLoadBalancerRequest) (*entity.LoadBalancer, sdkerror.Error)
}
