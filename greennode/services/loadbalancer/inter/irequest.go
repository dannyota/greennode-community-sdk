package inter

import "github.com/dannyota/greennode-community-sdk/v2/greennode/services/common"

type ICreateLoadBalancerRequest interface {
	ToRequestBody() interface{}
	AddUserAgent(agent ...string) ICreateLoadBalancerRequest
	WithListener(listener ICreateListenerRequest) ICreateLoadBalancerRequest
	WithPool(pool ICreatePoolRequest) ICreateLoadBalancerRequest
	WithProjectId(projectId string) ICreateLoadBalancerRequest
	WithTags(tags ...string) ICreateLoadBalancerRequest
	WithZoneId(zoneID common.Zone) ICreateLoadBalancerRequest
	GetMapHeaders() map[string]string
	ParseUserAgent() string
	ToMap() map[string]interface{}
}

type ICreateListenerRequest interface {
	ToRequestBody() interface{}
	WithAllowedCidrs(cidrs ...string) ICreateListenerRequest
	WithLoadBalancerId(lbid string) ICreateListenerRequest
	WithDefaultPoolId(poolId string) ICreateListenerRequest
	WithTimeoutClient(toc int) ICreateListenerRequest
	WithTimeoutConnection(toc int) ICreateListenerRequest
	WithTimeoutMember(tom int) ICreateListenerRequest
	AddCidrs(cidrs ...string) ICreateListenerRequest
	ParseUserAgent() string
	GetLoadBalancerId() string
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
	WithHttpVersion(version HealthCheckHttpVersion) IHealthMonitorRequest
	WithHealthCheckPath(path string) IHealthMonitorRequest
	WithSuccessCode(code string) IHealthMonitorRequest
	WithDomainName(domain string) IHealthMonitorRequest
}

type IMemberRequest interface {
	ToRequestBody() interface{}
	ToMap() map[string]interface{}
}
