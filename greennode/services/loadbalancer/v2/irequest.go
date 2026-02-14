package v2

import (
	"github.com/dannyota/greennode-community-sdk/v2/greennode/entity"
	"github.com/dannyota/greennode-community-sdk/v2/greennode/services/common"
)

type ICreateLoadBalancerRequest interface {
	ToRequestBody() interface{}
	AddUserAgent(agent ...string) ICreateLoadBalancerRequest
	WithListener(listener ICreateListenerRequest) ICreateLoadBalancerRequest
	WithPool(pool ICreatePoolRequest) ICreateLoadBalancerRequest
	WithTags(tags ...string) ICreateLoadBalancerRequest
	WithScheme(scheme LoadBalancerScheme) ICreateLoadBalancerRequest
	WithAutoScalable(autoScalable bool) ICreateLoadBalancerRequest
	WithPackageId(packageId string) ICreateLoadBalancerRequest
	WithSubnetId(subnetId string) ICreateLoadBalancerRequest
	WithType(typeVal LoadBalancerType) ICreateLoadBalancerRequest
	WithPoc(poc bool) ICreateLoadBalancerRequest
	WithZoneId(zoneId common.Zone) ICreateLoadBalancerRequest
	ParseUserAgent() string
	ToMap() map[string]interface{}
}

type IResizeLoadBalancerRequest interface {
	ToRequestBody() interface{}
	AddUserAgent(agent ...string) IResizeLoadBalancerRequest
	WithPackageId(packageId string) IResizeLoadBalancerRequest
	ParseUserAgent() string

	GetLoadBalancerId() string
}

type IListLoadBalancerPackagesRequest interface {
	WithZoneId(zoneId common.Zone) IListLoadBalancerPackagesRequest
	GetZoneId() string
	AddUserAgent(agent ...string) IListLoadBalancerPackagesRequest
	ParseUserAgent() string
	ToMap() map[string]interface{}
}

type IGetLoadBalancerByIdRequest interface {
	AddUserAgent(agent ...string) IGetLoadBalancerByIdRequest
	ParseUserAgent() string
	GetLoadBalancerId() string
}

type IListLoadBalancersRequest interface {
	WithName(name string) IListLoadBalancersRequest
	WithTags(tags ...string) IListLoadBalancersRequest
	ToListQuery() (string, error)
	ParseUserAgent() string
	GetDefaultQuery() string
	AddUserAgent(agent ...string) IListLoadBalancersRequest
}

type ICreateListenerRequest interface {
	ToRequestBody() interface{}
	WithAllowedCidrs(cidrs ...string) ICreateListenerRequest
	WithLoadBalancerId(lbid string) ICreateListenerRequest
	WithDefaultPoolId(poolId string) ICreateListenerRequest
	WithTimeoutClient(toc int) ICreateListenerRequest
	WithTimeoutConnection(toc int) ICreateListenerRequest
	WithTimeoutMember(tom int) ICreateListenerRequest
	WithInsertHeaders(headers ...string) ICreateListenerRequest
	WithDefaultCertificateAuthority(defaultCA *string) ICreateListenerRequest
	WithCertificateAuthorities(ca *[]string) ICreateListenerRequest
	WithClientCertificate(clientCert *string) ICreateListenerRequest
	AddCidrs(cidrs ...string) ICreateListenerRequest
	ParseUserAgent() string
	GetLoadBalancerId() string
	ToMap() map[string]interface{}
	AddUserAgent(agent ...string) ICreateListenerRequest
}

type IUpdateListenerRequest interface {
	GetLoadBalancerId() string
	GetListenerId() string
	ToRequestBody() interface{}
	WithCidrs(cidrs ...string) IUpdateListenerRequest
	WithTimeoutClient(toc int) IUpdateListenerRequest
	WithTimeoutConnection(toc int) IUpdateListenerRequest
	WithTimeoutMember(tom int) IUpdateListenerRequest
	WithDefaultPoolId(poolId string) IUpdateListenerRequest
	WithInsertHeaders(headers ...string) IUpdateListenerRequest
	ParseUserAgent() string
	AddUserAgent(agent ...string) IUpdateListenerRequest

	WithCertificateAuthorities(ca ...string) IUpdateListenerRequest
	WithClientCertificate(clientCert string) IUpdateListenerRequest
	WithDefaultCertificateAuthority(defaultCA string) IUpdateListenerRequest
}

type IGetPoolHealthMonitorByIdRequest interface {
	GetLoadBalancerId() string
	GetPoolId() string
	ParseUserAgent() string
	AddUserAgent(agent ...string) IGetPoolHealthMonitorByIdRequest
}

type ICreatePoolRequest interface {
	ToRequestBody() interface{}
	WithHealthMonitor(monitor IHealthMonitorRequest) ICreatePoolRequest
	WithMembers(members ...IMemberRequest) ICreatePoolRequest
	WithLoadBalancerId(lbId string) ICreatePoolRequest
	WithAlgorithm(algorithm PoolAlgorithm) ICreatePoolRequest
	ToMap() map[string]interface{}
	GetLoadBalancerId() string
	ParseUserAgent() string
	AddUserAgent(agent ...string) ICreatePoolRequest
}

type IUpdatePoolRequest interface {
	GetPoolId() string
	ToRequestBody() interface{}
	WithHealthMonitor(monitor IHealthMonitorRequest) IUpdatePoolRequest
	WithLoadBalancerId(lbId string) IUpdatePoolRequest
	WithAlgorithm(algorithm PoolAlgorithm) IUpdatePoolRequest
	WithStickiness(v *bool) IUpdatePoolRequest
	WithTLSEncryption(v *bool) IUpdatePoolRequest
	ToMap() map[string]interface{}
	GetLoadBalancerId() string
	ParseUserAgent() string
	AddUserAgent(agent ...string) IUpdatePoolRequest
}

type IListListenersByLoadBalancerIdRequest interface {
	GetLoadBalancerId() string
	ParseUserAgent() string
	AddUserAgent(agent ...string) IListListenersByLoadBalancerIdRequest
}

type IListPoolsByLoadBalancerIdRequest interface {
	GetLoadBalancerId() string
	ParseUserAgent() string
	AddUserAgent(agent ...string) IListPoolsByLoadBalancerIdRequest
}

type IUpdatePoolMembersRequest interface {
	WithMembers(members ...IMemberRequest) IUpdatePoolMembersRequest
	ToRequestBody() interface{}
	GetLoadBalancerId() string
	GetPoolId() string
	ParseUserAgent() string
	AddUserAgent(agent ...string) IUpdatePoolMembersRequest
}

type IListPoolMembersRequest interface {
	GetLoadBalancerId() string
	GetPoolId() string
	ParseUserAgent() string
	AddUserAgent(agent ...string) IListPoolMembersRequest
}

type IDeletePoolByIdRequest interface {
	GetLoadBalancerId() string
	GetPoolId() string
	ParseUserAgent() string
	AddUserAgent(agent ...string) IDeletePoolByIdRequest
}

type IDeleteListenerByIdRequest interface {
	GetLoadBalancerId() string
	GetListenerId() string
	ParseUserAgent() string
	AddUserAgent(agent ...string) IDeleteListenerByIdRequest
}

type IDeleteLoadBalancerByIdRequest interface {
	GetLoadBalancerId() string
	ParseUserAgent() string
	AddUserAgent(agent ...string) IDeleteLoadBalancerByIdRequest
}

type IHealthMonitorRequest interface {
	ToRequestBody() interface{}
	ToMap() map[string]interface{}
	WithHealthCheckProtocol(protocol HealthCheckProtocol) IHealthMonitorRequest
	WithHealthyThreshold(ht int) IHealthMonitorRequest
	WithUnhealthyThreshold(uht int) IHealthMonitorRequest
	WithInterval(interval int) IHealthMonitorRequest
	WithTimeout(to int) IHealthMonitorRequest
	WithHealthCheckMethod(method *HealthCheckMethod) IHealthMonitorRequest
	WithHttpVersion(version *HealthCheckHttpVersion) IHealthMonitorRequest
	WithHealthCheckPath(path *string) IHealthMonitorRequest
	WithSuccessCode(code *string) IHealthMonitorRequest
	WithDomainName(domain *string) IHealthMonitorRequest
	AddUserAgent(agent ...string) IHealthMonitorRequest
}

type IMemberRequest interface {
	ToRequestBody() interface{}
	ToMap() map[string]interface{}
}

type IListTagsRequest interface {
	GetLoadBalancerId() string
	ParseUserAgent() string
	AddUserAgent(agent ...string) IListTagsRequest
}

type ICreateTagsRequest interface {
	GetLoadBalancerId() string
	ToRequestBody() interface{}
	ParseUserAgent() string
	WithTags(tags ...string) ICreateTagsRequest
	AddUserAgent(agent ...string) ICreateTagsRequest
}

type IUpdateTagsRequest interface {
	GetLoadBalancerId() string
	ToRequestBody(lstTags *entity.ListTags) interface{}
	ParseUserAgent() string
	WithTags(tags ...string) IUpdateTagsRequest
	ToMap() map[string]interface{}
	AddUserAgent(agent ...string) IUpdateTagsRequest
}

// --------------------------------------------------------

type IListPoliciesRequest interface {
	ParseUserAgent() string
	GetLoadBalancerId() string
	GetListenerId() string
	AddUserAgent(agent ...string) IListPoliciesRequest
}

type ICreatePolicyRequest interface {
	ToRequestBody() interface{}
	ParseUserAgent() string
	GetLoadBalancerId() string
	GetListenerId() string
	ToMap() map[string]interface{}
	AddUserAgent(agent ...string) ICreatePolicyRequest

	WithName(name string) ICreatePolicyRequest
	WithRules(rules ...L7RuleRequest) ICreatePolicyRequest
	WithAction(action PolicyAction) ICreatePolicyRequest

	// only for action redirect to pool
	WithRedirectPoolId(redirectPoolId string) ICreatePolicyRequest

	// only for action redirect to url
	WithRedirectURL(redirectURL string) ICreatePolicyRequest
	// only for action redirect to url
	WithRedirectHTTPCode(redirectHTTPCode int) ICreatePolicyRequest
	// only for action redirect to url
	WithKeepQueryString(keepQueryString bool) ICreatePolicyRequest
}

type IGetPolicyByIdRequest interface {
	ParseUserAgent() string
	GetLoadBalancerId() string
	GetListenerId() string
	GetPolicyId() string
	AddUserAgent(agent ...string) IGetPolicyByIdRequest
}

type IUpdatePolicyRequest interface {
	ToRequestBody() interface{}
	ParseUserAgent() string
	GetLoadBalancerId() string
	GetListenerId() string
	GetPolicyId() string
	AddUserAgent(agent ...string) IUpdatePolicyRequest

	WithAction(action PolicyAction) IUpdatePolicyRequest
	WithRules(rules ...L7RuleRequest) IUpdatePolicyRequest
	WithRedirectPoolID(redirectPoolId string) IUpdatePolicyRequest
	WithRedirectURL(redirectURL string) IUpdatePolicyRequest
	WithRedirectHTTPCode(redirectHTTPCode int) IUpdatePolicyRequest
	WithKeepQueryString(keepQueryString bool) IUpdatePolicyRequest
}

type IDeletePolicyByIdRequest interface {
	ParseUserAgent() string
	GetLoadBalancerId() string
	GetListenerId() string
	GetPolicyId() string
	AddUserAgent(agent ...string) IDeletePolicyByIdRequest
}

type IReorderPoliciesRequest interface {
	ToRequestBody() interface{}
	ParseUserAgent() string
	GetLoadBalancerId() string
	GetListenerId() string
	AddUserAgent(agent ...string) IReorderPoliciesRequest

	WithPoliciesOrder(lstPolicies []string) IReorderPoliciesRequest
}

type IGetPoolByIdRequest interface {
	GetLoadBalancerId() string
	GetPoolId() string
	ParseUserAgent() string
	AddUserAgent(agent ...string) IGetPoolByIdRequest
}

type IGetListenerByIdRequest interface {
	GetLoadBalancerId() string
	GetListenerId() string
	ParseUserAgent() string
	AddUserAgent(agent ...string) IGetListenerByIdRequest
}

type IResizeLoadBalancerByIdRequest interface {
	GetLoadBalancerId() string
	ToMap() map[string]interface{}
	ParseUserAgent() string
	ToRequestBody() interface{}
	AddUserAgent(agent ...string) IResizeLoadBalancerByIdRequest
}

type IScaleLoadBalancerRequest interface {
	GetLoadBalancerId() string
	ToMap() map[string]interface{}
	ParseUserAgent() string
	ToRequestBody() interface{}
	AddUserAgent(agent ...string) IScaleLoadBalancerRequest
	WithScaling(scaling *ScalingConfig) IScaleLoadBalancerRequest
	WithNetworking(networking *NetworkingConfig) IScaleLoadBalancerRequest
}

// --------------------------------------------------------

type IListCertificatesRequest interface {
	ParseUserAgent() string
	AddUserAgent(agent ...string) IListCertificatesRequest
}

type IGetCertificateByIdRequest interface {
	GetCertificateId() string
	ParseUserAgent() string
	AddUserAgent(agent ...string) IGetCertificateByIdRequest
}

type ICreateCertificateRequest interface {
	ToRequestBody() interface{}
	ParseUserAgent() string
	ToMap() map[string]interface{}
	AddUserAgent(agent ...string) ICreateCertificateRequest

	WithCertificateChain(chain string) ICreateCertificateRequest
	WithPassphrase(passphrase string) ICreateCertificateRequest
	WithPrivateKey(privateKey string) ICreateCertificateRequest
}

type IDeleteCertificateByIdRequest interface {
	GetCertificateId() string
	ParseUserAgent() string
	AddUserAgent(agent ...string) IDeleteCertificateByIdRequest
}
