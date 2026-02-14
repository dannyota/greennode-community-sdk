package inter

import "github.com/dannyota/greennode-community-sdk/v2/greennode/services/common"

type ICreateLoadBalancerRequest interface {
	ToRequestBody() interface{}
	AddUserAgent(agent ...string) ICreateLoadBalancerRequest
	WithListener(listener ICreateListenerRequest) ICreateLoadBalancerRequest
	WithPool(pool ICreatePoolRequest) ICreateLoadBalancerRequest
	WithProjectID(projectID string) ICreateLoadBalancerRequest
	WithTags(tags ...string) ICreateLoadBalancerRequest
	WithZoneID(zoneID common.Zone) ICreateLoadBalancerRequest
	GetMapHeaders() map[string]string
	ParseUserAgent() string
	ToMap() map[string]interface{}
}

type ICreateListenerRequest interface {
	ToRequestBody() interface{}
	WithAllowedCidrs(cidrs ...string) ICreateListenerRequest
	WithLoadBalancerID(lbid string) ICreateListenerRequest
	WithDefaultPoolID(poolID string) ICreateListenerRequest
	WithTimeoutClient(toc int) ICreateListenerRequest
	WithTimeoutConnection(toc int) ICreateListenerRequest
	WithTimeoutMember(tom int) ICreateListenerRequest
	AddCidrs(cidrs ...string) ICreateListenerRequest
	ParseUserAgent() string
	GetLoadBalancerID() string
	ToMap() map[string]interface{}
}

type ICreatePoolRequest interface {
	ToRequestBody() interface{}
	WithHealthMonitor(monitor IHealthMonitorRequest) ICreatePoolRequest
	WithMembers(members ...IMemberRequest) ICreatePoolRequest
	WithAlgorithm(algorithm PoolAlgorithm) ICreatePoolRequest
}

type IHealthMonitorRequest interface {
	ToRequestBody() interface{}
	ToMap() map[string]interface{}
	WithHealthyThreshold(ht int) IHealthMonitorRequest
	WithUnhealthyThreshold(uht int) IHealthMonitorRequest
	WithInterval(interval int) IHealthMonitorRequest
	WithTimeout(to int) IHealthMonitorRequest
	WithHealthCheckMethod(method HealthCheckMethod) IHealthMonitorRequest
	WithHTTPVersion(version HealthCheckHTTPVersion) IHealthMonitorRequest
	WithHealthCheckPath(path string) IHealthMonitorRequest
	WithSuccessCode(code string) IHealthMonitorRequest
	WithDomainName(domain string) IHealthMonitorRequest
}

type IMemberRequest interface {
	ToRequestBody() interface{}
	ToMap() map[string]interface{}
}
