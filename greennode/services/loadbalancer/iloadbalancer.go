package loadbalancer

import (
	"github.com/dannyota/greennode-community-sdk/v2/greennode/entity"
	sdkerror "github.com/dannyota/greennode-community-sdk/v2/greennode/sdk_error"
	"github.com/dannyota/greennode-community-sdk/v2/greennode/services/loadbalancer/inter"
	lbv2 "github.com/dannyota/greennode-community-sdk/v2/greennode/services/loadbalancer/v2"
)

type ILoadBalancerServiceV2 interface {
	CreateLoadBalancer(popts lbv2.ICreateLoadBalancerRequest) (*entity.LoadBalancer, sdkerror.IError)
	ResizeLoadBalancer(popts lbv2.IResizeLoadBalancerRequest) (*entity.LoadBalancer, sdkerror.IError)
	ListLoadBalancerPackages(popts lbv2.IListLoadBalancerPackagesRequest) (*entity.ListLoadBalancerPackages, sdkerror.IError)
	GetLoadBalancerById(popts lbv2.IGetLoadBalancerByIdRequest) (*entity.LoadBalancer, sdkerror.IError)
	ListLoadBalancers(popts lbv2.IListLoadBalancersRequest) (*entity.ListLoadBalancers, sdkerror.IError)
	GetPoolHealthMonitorById(popts lbv2.IGetPoolHealthMonitorByIdRequest) (*entity.HealthMonitor, sdkerror.IError)
	CreatePool(popts lbv2.ICreatePoolRequest) (*entity.Pool, sdkerror.IError)
	UpdatePool(popts lbv2.IUpdatePoolRequest) sdkerror.IError
	CreateListener(popts lbv2.ICreateListenerRequest) (*entity.Listener, sdkerror.IError)
	UpdateListener(popts lbv2.IUpdateListenerRequest) sdkerror.IError
	ListListenersByLoadBalancerId(popts lbv2.IListListenersByLoadBalancerIdRequest) (*entity.ListListeners, sdkerror.IError)
	ListPoolsByLoadBalancerId(popts lbv2.IListPoolsByLoadBalancerIdRequest) (*entity.ListPools, sdkerror.IError)
	UpdatePoolMembers(popts lbv2.IUpdatePoolMembersRequest) sdkerror.IError
	ListPoolMembers(popts lbv2.IListPoolMembersRequest) (*entity.ListMembers, sdkerror.IError)
	DeletePoolById(popt lbv2.IDeletePoolByIdRequest) sdkerror.IError
	DeleteListenerById(popts lbv2.IDeleteListenerByIdRequest) sdkerror.IError
	DeleteLoadBalancerById(popts lbv2.IDeleteLoadBalancerByIdRequest) sdkerror.IError
	ListTags(popts lbv2.IListTagsRequest) (*entity.ListTags, sdkerror.IError)
	CreateTags(popts lbv2.ICreateTagsRequest) sdkerror.IError
	UpdateTags(popts lbv2.IUpdateTagsRequest) sdkerror.IError

	ListPolicies(popts lbv2.IListPoliciesRequest) (*entity.ListPolicies, sdkerror.IError)
	CreatePolicy(popts lbv2.ICreatePolicyRequest) (*entity.Policy, sdkerror.IError)
	GetPolicyById(popts lbv2.IGetPolicyByIdRequest) (*entity.Policy, sdkerror.IError)
	UpdatePolicy(popts lbv2.IUpdatePolicyRequest) sdkerror.IError
	DeletePolicyById(popts lbv2.IDeletePolicyByIdRequest) sdkerror.IError
	GetPoolById(popts lbv2.IGetPoolByIdRequest) (*entity.Pool, sdkerror.IError)
	GetListenerById(popts lbv2.IGetListenerByIdRequest) (*entity.Listener, sdkerror.IError)
	ResizeLoadBalancerById(popts lbv2.IResizeLoadBalancerByIdRequest) sdkerror.IError
	ScaleLoadBalancer(popts lbv2.IScaleLoadBalancerRequest) (*entity.LoadBalancer, sdkerror.IError)
	ReorderPolicies(popts lbv2.IReorderPoliciesRequest) sdkerror.IError

	ListCertificates(popts lbv2.IListCertificatesRequest) (*entity.ListCertificates, sdkerror.IError)
	GetCertificateById(popts lbv2.IGetCertificateByIdRequest) (*entity.Certificate, sdkerror.IError)
	CreateCertificate(popts lbv2.ICreateCertificateRequest) (*entity.Certificate, sdkerror.IError)
	DeleteCertificateById(popts lbv2.IDeleteCertificateByIdRequest) sdkerror.IError
}

type ILoadBalancerServiceInternal interface {
	CreateLoadBalancer(popts inter.ICreateLoadBalancerRequest) (*entity.LoadBalancer, sdkerror.IError)
}
