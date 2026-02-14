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
	WithPackageID(packageID string) ICreateLoadBalancerRequest
	WithSubnetID(subnetID string) ICreateLoadBalancerRequest
	WithType(typeVal LoadBalancerType) ICreateLoadBalancerRequest
	WithPoc(poc bool) ICreateLoadBalancerRequest
	WithZoneID(zoneID common.Zone) ICreateLoadBalancerRequest
	ParseUserAgent() string
	ToMap() map[string]interface{}
}

type IResizeLoadBalancerRequest interface {
	ToRequestBody() interface{}
	AddUserAgent(agent ...string) IResizeLoadBalancerRequest
	WithPackageID(packageID string) IResizeLoadBalancerRequest
	ParseUserAgent() string

	GetLoadBalancerID() string
}

type IListLoadBalancerPackagesRequest interface {
	WithZoneID(zoneID common.Zone) IListLoadBalancerPackagesRequest
	GetZoneID() string
	AddUserAgent(agent ...string) IListLoadBalancerPackagesRequest
	ParseUserAgent() string
	ToMap() map[string]interface{}
}

type IGetLoadBalancerByIDRequest interface {
	AddUserAgent(agent ...string) IGetLoadBalancerByIDRequest
	ParseUserAgent() string
	GetLoadBalancerID() string
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
	WithLoadBalancerID(lbid string) ICreateListenerRequest
	WithDefaultPoolID(poolID string) ICreateListenerRequest
	WithTimeoutClient(toc int) ICreateListenerRequest
	WithTimeoutConnection(toc int) ICreateListenerRequest
	WithTimeoutMember(tom int) ICreateListenerRequest
	WithInsertHeaders(headers ...string) ICreateListenerRequest
	WithDefaultCertificateAuthority(defaultCA *string) ICreateListenerRequest
	WithCertificateAuthorities(ca *[]string) ICreateListenerRequest
	WithClientCertificate(clientCert *string) ICreateListenerRequest
	AddCidrs(cidrs ...string) ICreateListenerRequest
	ParseUserAgent() string
	GetLoadBalancerID() string
	ToMap() map[string]interface{}
	AddUserAgent(agent ...string) ICreateListenerRequest
}

type IUpdateListenerRequest interface {
	GetLoadBalancerID() string
	GetListenerID() string
	ToRequestBody() interface{}
	WithCidrs(cidrs ...string) IUpdateListenerRequest
	WithTimeoutClient(toc int) IUpdateListenerRequest
	WithTimeoutConnection(toc int) IUpdateListenerRequest
	WithTimeoutMember(tom int) IUpdateListenerRequest
	WithDefaultPoolID(poolID string) IUpdateListenerRequest
	WithInsertHeaders(headers ...string) IUpdateListenerRequest
	ParseUserAgent() string
	AddUserAgent(agent ...string) IUpdateListenerRequest

	WithCertificateAuthorities(ca ...string) IUpdateListenerRequest
	WithClientCertificate(clientCert string) IUpdateListenerRequest
	WithDefaultCertificateAuthority(defaultCA string) IUpdateListenerRequest
}

type IGetPoolHealthMonitorByIDRequest interface {
	GetLoadBalancerID() string
	GetPoolID() string
	ParseUserAgent() string
	AddUserAgent(agent ...string) IGetPoolHealthMonitorByIDRequest
}

type ICreatePoolRequest interface {
	ToRequestBody() interface{}
	WithHealthMonitor(monitor IHealthMonitorRequest) ICreatePoolRequest
	WithMembers(members ...IMemberRequest) ICreatePoolRequest
	WithLoadBalancerID(lbID string) ICreatePoolRequest
	WithAlgorithm(algorithm PoolAlgorithm) ICreatePoolRequest
	ToMap() map[string]interface{}
	GetLoadBalancerID() string
	ParseUserAgent() string
	AddUserAgent(agent ...string) ICreatePoolRequest
}

type IUpdatePoolRequest interface {
	GetPoolID() string
	ToRequestBody() interface{}
	WithHealthMonitor(monitor IHealthMonitorRequest) IUpdatePoolRequest
	WithLoadBalancerID(lbID string) IUpdatePoolRequest
	WithAlgorithm(algorithm PoolAlgorithm) IUpdatePoolRequest
	WithStickiness(v *bool) IUpdatePoolRequest
	WithTLSEncryption(v *bool) IUpdatePoolRequest
	ToMap() map[string]interface{}
	GetLoadBalancerID() string
	ParseUserAgent() string
	AddUserAgent(agent ...string) IUpdatePoolRequest
}

type IListListenersByLoadBalancerIDRequest interface {
	GetLoadBalancerID() string
	ParseUserAgent() string
	AddUserAgent(agent ...string) IListListenersByLoadBalancerIDRequest
}

type IListPoolsByLoadBalancerIDRequest interface {
	GetLoadBalancerID() string
	ParseUserAgent() string
	AddUserAgent(agent ...string) IListPoolsByLoadBalancerIDRequest
}

type IUpdatePoolMembersRequest interface {
	WithMembers(members ...IMemberRequest) IUpdatePoolMembersRequest
	ToRequestBody() interface{}
	GetLoadBalancerID() string
	GetPoolID() string
	ParseUserAgent() string
	AddUserAgent(agent ...string) IUpdatePoolMembersRequest
}

type IListPoolMembersRequest interface {
	GetLoadBalancerID() string
	GetPoolID() string
	ParseUserAgent() string
	AddUserAgent(agent ...string) IListPoolMembersRequest
}

type IDeletePoolByIDRequest interface {
	GetLoadBalancerID() string
	GetPoolID() string
	ParseUserAgent() string
	AddUserAgent(agent ...string) IDeletePoolByIDRequest
}

type IDeleteListenerByIDRequest interface {
	GetLoadBalancerID() string
	GetListenerID() string
	ParseUserAgent() string
	AddUserAgent(agent ...string) IDeleteListenerByIDRequest
}

type IDeleteLoadBalancerByIDRequest interface {
	GetLoadBalancerID() string
	ParseUserAgent() string
	AddUserAgent(agent ...string) IDeleteLoadBalancerByIDRequest
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
	WithHTTPVersion(version *HealthCheckHTTPVersion) IHealthMonitorRequest
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
	GetLoadBalancerID() string
	ParseUserAgent() string
	AddUserAgent(agent ...string) IListTagsRequest
}

type ICreateTagsRequest interface {
	GetLoadBalancerID() string
	ToRequestBody() interface{}
	ParseUserAgent() string
	WithTags(tags ...string) ICreateTagsRequest
	AddUserAgent(agent ...string) ICreateTagsRequest
}

type IUpdateTagsRequest interface {
	GetLoadBalancerID() string
	ToRequestBody(lstTags *entity.ListTags) interface{}
	ParseUserAgent() string
	WithTags(tags ...string) IUpdateTagsRequest
	ToMap() map[string]interface{}
	AddUserAgent(agent ...string) IUpdateTagsRequest
}

// --------------------------------------------------------

type IListPoliciesRequest interface {
	ParseUserAgent() string
	GetLoadBalancerID() string
	GetListenerID() string
	AddUserAgent(agent ...string) IListPoliciesRequest
}

type ICreatePolicyRequest interface {
	ToRequestBody() interface{}
	ParseUserAgent() string
	GetLoadBalancerID() string
	GetListenerID() string
	ToMap() map[string]interface{}
	AddUserAgent(agent ...string) ICreatePolicyRequest

	WithName(name string) ICreatePolicyRequest
	WithRules(rules ...L7RuleRequest) ICreatePolicyRequest
	WithAction(action PolicyAction) ICreatePolicyRequest

	// only for action redirect to pool
	WithRedirectPoolID(redirectPoolID string) ICreatePolicyRequest

	// only for action redirect to url
	WithRedirectURL(redirectURL string) ICreatePolicyRequest
	// only for action redirect to url
	WithRedirectHTTPCode(redirectHTTPCode int) ICreatePolicyRequest
	// only for action redirect to url
	WithKeepQueryString(keepQueryString bool) ICreatePolicyRequest
}

type IGetPolicyByIDRequest interface {
	ParseUserAgent() string
	GetLoadBalancerID() string
	GetListenerID() string
	GetPolicyID() string
	AddUserAgent(agent ...string) IGetPolicyByIDRequest
}

type IUpdatePolicyRequest interface {
	ToRequestBody() interface{}
	ParseUserAgent() string
	GetLoadBalancerID() string
	GetListenerID() string
	GetPolicyID() string
	AddUserAgent(agent ...string) IUpdatePolicyRequest

	WithAction(action PolicyAction) IUpdatePolicyRequest
	WithRules(rules ...L7RuleRequest) IUpdatePolicyRequest
	WithRedirectPoolID(redirectPoolID string) IUpdatePolicyRequest
	WithRedirectURL(redirectURL string) IUpdatePolicyRequest
	WithRedirectHTTPCode(redirectHTTPCode int) IUpdatePolicyRequest
	WithKeepQueryString(keepQueryString bool) IUpdatePolicyRequest
}

type IDeletePolicyByIDRequest interface {
	ParseUserAgent() string
	GetLoadBalancerID() string
	GetListenerID() string
	GetPolicyID() string
	AddUserAgent(agent ...string) IDeletePolicyByIDRequest
}

type IReorderPoliciesRequest interface {
	ToRequestBody() interface{}
	ParseUserAgent() string
	GetLoadBalancerID() string
	GetListenerID() string
	AddUserAgent(agent ...string) IReorderPoliciesRequest

	WithPoliciesOrder(lstPolicies []string) IReorderPoliciesRequest
}

type IGetPoolByIDRequest interface {
	GetLoadBalancerID() string
	GetPoolID() string
	ParseUserAgent() string
	AddUserAgent(agent ...string) IGetPoolByIDRequest
}

type IGetListenerByIDRequest interface {
	GetLoadBalancerID() string
	GetListenerID() string
	ParseUserAgent() string
	AddUserAgent(agent ...string) IGetListenerByIDRequest
}

type IResizeLoadBalancerByIDRequest interface {
	GetLoadBalancerID() string
	ToMap() map[string]interface{}
	ParseUserAgent() string
	ToRequestBody() interface{}
	AddUserAgent(agent ...string) IResizeLoadBalancerByIDRequest
}

type IScaleLoadBalancerRequest interface {
	GetLoadBalancerID() string
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

type IGetCertificateByIDRequest interface {
	GetCertificateID() string
	ParseUserAgent() string
	AddUserAgent(agent ...string) IGetCertificateByIDRequest
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

type IDeleteCertificateByIDRequest interface {
	GetCertificateID() string
	ParseUserAgent() string
	AddUserAgent(agent ...string) IDeleteCertificateByIDRequest
}
