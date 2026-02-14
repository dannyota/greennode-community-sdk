package v1

type IListGlobalPoolsRequest interface {
	WithLoadBalancerId(lbId string) IListGlobalPoolsRequest
	GetLoadBalancerId() string

	AddUserAgent(agent ...string) IListGlobalPoolsRequest
	ParseUserAgent() string
}

// --------------------------------------------------------

type ICreateGlobalPoolRequest interface {
	WithAlgorithm(algorithm GlobalPoolAlgorithm) ICreateGlobalPoolRequest
	WithDescription(desc string) ICreateGlobalPoolRequest
	WithName(name string) ICreateGlobalPoolRequest
	WithProtocol(protocol GlobalPoolProtocol) ICreateGlobalPoolRequest
	WithHealthMonitor(monitor IGlobalHealthMonitorRequest) ICreateGlobalPoolRequest
	WithMembers(members ...ICreateGlobalPoolMemberRequest) ICreateGlobalPoolRequest

	WithLoadBalancerId(lbId string) ICreateGlobalPoolRequest
	GetLoadBalancerId() string // to use in request url

	AddUserAgent(agent ...string) ICreateGlobalPoolRequest
	ParseUserAgent() string
	ToRequestBody() interface{}
	ToMap() map[string]interface{}
}

type IGlobalHealthMonitorRequest interface {
	WithHealthyThreshold(ht int) IGlobalHealthMonitorRequest
	WithInterval(interval int) IGlobalHealthMonitorRequest
	WithProtocol(protocol GlobalPoolHealthCheckProtocol) IGlobalHealthMonitorRequest
	WithTimeout(to int) IGlobalHealthMonitorRequest
	WithUnhealthyThreshold(uht int) IGlobalHealthMonitorRequest

	// http, https
	WithHealthCheckMethod(method *GlobalPoolHealthCheckMethod) IGlobalHealthMonitorRequest
	WithHttpVersion(version *GlobalPoolHealthCheckHttpVersion) IGlobalHealthMonitorRequest
	WithPath(path *string) IGlobalHealthMonitorRequest
	WithSuccessCode(code *string) IGlobalHealthMonitorRequest
	WithDomainName(domain *string) IGlobalHealthMonitorRequest

	AddUserAgent(agent ...string) IGlobalHealthMonitorRequest
	ParseUserAgent() string
	ToRequestBody() interface{}
	ToMap() map[string]interface{}
}
type ICreateGlobalPoolMemberRequest interface {
	WithName(name string) ICreateGlobalPoolMemberRequest
	WithDescription(desc string) ICreateGlobalPoolMemberRequest
	WithRegion(region string) ICreateGlobalPoolMemberRequest
	WithVPCID(vpcID string) ICreateGlobalPoolMemberRequest
	WithTrafficDial(dial int) ICreateGlobalPoolMemberRequest
	WithMembers(members ...IGlobalMemberRequest) ICreateGlobalPoolMemberRequest
	WithType(typeVal GlobalPoolMemberType) ICreateGlobalPoolMemberRequest

	WithLoadBalancerId(lbId string) ICreateGlobalPoolMemberRequest
	WithPoolId(poolId string) ICreateGlobalPoolMemberRequest
	GetLoadBalancerId() string // to use in request url
	GetPoolId() string         // to use in request url

	AddUserAgent(agent ...string) ICreateGlobalPoolMemberRequest
	ParseUserAgent() string
	ToRequestBody() interface{}
	ToMap() map[string]interface{}
}

type IGlobalMemberRequest interface {
	WithAddress(addr string) IGlobalMemberRequest
	WithBackupRole(backup bool) IGlobalMemberRequest
	WithDescription(desc string) IGlobalMemberRequest
	WithMonitorPort(port int) IGlobalMemberRequest
	WithName(name string) IGlobalMemberRequest
	WithPort(port int) IGlobalMemberRequest
	WithSubnetID(subnetID string) IGlobalMemberRequest
	WithWeight(weight int) IGlobalMemberRequest

	AddUserAgent(agent ...string) IGlobalMemberRequest
	ParseUserAgent() string
	ToRequestBody() interface{}
	ToMap() map[string]interface{}
}

// --------------------------------------------------------

type IUpdateGlobalPoolRequest interface {
	WithAlgorithm(algorithm GlobalPoolAlgorithm) IUpdateGlobalPoolRequest
	WithHealthMonitor(monitor IGlobalHealthMonitorRequest) IUpdateGlobalPoolRequest

	WithLoadBalancerId(lbId string) IUpdateGlobalPoolRequest
	WithPoolId(poolId string) IUpdateGlobalPoolRequest
	GetLoadBalancerId() string // to use in request url
	GetPoolId() string

	AddUserAgent(agent ...string) IUpdateGlobalPoolRequest
	ParseUserAgent() string
	ToRequestBody() interface{}
	ToMap() map[string]interface{}
}

// --------------------------------------------------------

type IDeleteGlobalPoolRequest interface {
	WithLoadBalancerId(lbId string) IDeleteGlobalPoolRequest
	WithPoolId(poolId string) IDeleteGlobalPoolRequest
	GetLoadBalancerId() string // to use in request url
	GetPoolId() string

	AddUserAgent(agent ...string) IDeleteGlobalPoolRequest
	ParseUserAgent() string
}

// --------------------------------------------------------

type IListGlobalPoolMembersRequest interface {
	WithLoadBalancerId(lbId string) IListGlobalPoolMembersRequest
	WithPoolId(poolId string) IListGlobalPoolMembersRequest
	GetLoadBalancerId() string // to use in request url
	GetPoolId() string

	AddUserAgent(agent ...string) IListGlobalPoolMembersRequest
	ParseUserAgent() string
}

// --------------------------------------------------------

type IGetGlobalPoolMemberRequest interface {
	WithLoadBalancerId(lbId string) IGetGlobalPoolMemberRequest
	WithPoolId(poolId string) IGetGlobalPoolMemberRequest
	WithPoolMemberId(poolMemberId string) IGetGlobalPoolMemberRequest
	GetLoadBalancerId() string // to use in request url
	GetPoolId() string
	GetPoolMemberId() string

	AddUserAgent(agent ...string) IGetGlobalPoolMemberRequest
	ParseUserAgent() string
}

// --------------------------------------------------------

type IDeleteGlobalPoolMemberRequest interface {
	WithLoadBalancerId(lbId string) IDeleteGlobalPoolMemberRequest
	WithPoolId(poolId string) IDeleteGlobalPoolMemberRequest
	WithPoolMemberId(poolMemberId string) IDeleteGlobalPoolMemberRequest
	GetLoadBalancerId() string // to use in request url
	GetPoolId() string
	GetPoolMemberId() string

	AddUserAgent(agent ...string) IDeleteGlobalPoolMemberRequest
	ParseUserAgent() string
}

// --------------------------------------------------------

type IPatchGlobalPoolMembersRequest interface {
	WithBulkAction(action ...IBulkActionRequest) IPatchGlobalPoolMembersRequest

	WithLoadBalancerId(lbId string) IPatchGlobalPoolMembersRequest
	WithPoolId(poolId string) IPatchGlobalPoolMembersRequest
	GetLoadBalancerId() string // to use in request url
	GetPoolId() string

	AddUserAgent(agent ...string) IPatchGlobalPoolMembersRequest
	ParseUserAgent() string
	ToRequestBody() interface{}
	ToMap() map[string]interface{}
}

type IBulkActionRequest interface {
	ToRequestBody() interface{}
	ToMap() map[string]interface{}
}

// --------------------------------------------------------

type IListGlobalListenersRequest interface {
	WithLoadBalancerId(lbId string) IListGlobalListenersRequest
	GetLoadBalancerId() string // to use in request url

	AddUserAgent(agent ...string) IListGlobalListenersRequest
	ParseUserAgent() string
}

// --------------------------------------------------------

type IGetGlobalListenerRequest interface {
	WithLoadBalancerId(lbId string) IGetGlobalListenerRequest
	WithListenerId(listenerId string) IGetGlobalListenerRequest
	GetLoadBalancerId() string // to use in request url
	GetListenerId() string

	AddUserAgent(agent ...string) IGetGlobalListenerRequest
	ParseUserAgent() string
}

// --------------------------------------------------------

type ICreateGlobalListenerRequest interface {
	WithAllowedCidrs(cidrs ...string) ICreateGlobalListenerRequest
	WithDescription(desc string) ICreateGlobalListenerRequest
	WithHeaders(headers ...string) ICreateGlobalListenerRequest
	WithName(name string) ICreateGlobalListenerRequest
	WithPort(port int) ICreateGlobalListenerRequest
	WithProtocol(protocol GlobalListenerProtocol) ICreateGlobalListenerRequest
	WithTimeoutClient(toc int) ICreateGlobalListenerRequest
	WithTimeoutConnection(toc int) ICreateGlobalListenerRequest
	WithTimeoutMember(tom int) ICreateGlobalListenerRequest
	WithGlobalPoolId(poolId string) ICreateGlobalListenerRequest

	WithLoadBalancerId(lbid string) ICreateGlobalListenerRequest
	GetLoadBalancerId() string
	// AddCidrs(pcidrs ...string) ICreateGlobalListenerRequest

	AddUserAgent(agent ...string) ICreateGlobalListenerRequest
	ParseUserAgent() string
	ToRequestBody() interface{}
	ToMap() map[string]interface{}
}

// --------------------------------------------------------

type IUpdateGlobalListenerRequest interface {
	WithAllowedCidrs(cidrs ...string) IUpdateGlobalListenerRequest
	WithTimeoutClient(toc int) IUpdateGlobalListenerRequest
	WithTimeoutMember(tom int) IUpdateGlobalListenerRequest
	WithTimeoutConnection(toc int) IUpdateGlobalListenerRequest
	WithHeaders(headers ...string) IUpdateGlobalListenerRequest
	WithGlobalPoolId(poolId string) IUpdateGlobalListenerRequest

	WithLoadBalancerId(lbId string) IUpdateGlobalListenerRequest
	WithListenerId(listenerId string) IUpdateGlobalListenerRequest
	GetLoadBalancerId() string // to use in request url
	GetListenerId() string

	AddUserAgent(agent ...string) IUpdateGlobalListenerRequest
	ParseUserAgent() string
	ToRequestBody() interface{}
	ToMap() map[string]interface{}
}

// --------------------------------------------------------

type IDeleteGlobalListenerRequest interface {
	WithLoadBalancerId(lbId string) IDeleteGlobalListenerRequest
	WithListenerId(listenerId string) IDeleteGlobalListenerRequest
	GetLoadBalancerId() string // to use in request url
	GetListenerId() string

	AddUserAgent(agent ...string) IDeleteGlobalListenerRequest
	ParseUserAgent() string
}

// --------------------------------------------------------

type IListGlobalPackagesRequest interface {
	AddUserAgent(agent ...string) IListGlobalPackagesRequest
	ParseUserAgent() string
}

// --------------------------------------------------------

type IListGlobalRegionsRequest interface {
	AddUserAgent(agent ...string) IListGlobalRegionsRequest
	ParseUserAgent() string
}

// --------------------------------------------------------

type IGetGlobalLoadBalancerUsageHistoriesRequest interface {
	WithLoadBalancerId(lbId string) IGetGlobalLoadBalancerUsageHistoriesRequest
	WithFrom(from string) IGetGlobalLoadBalancerUsageHistoriesRequest
	WithTo(to string) IGetGlobalLoadBalancerUsageHistoriesRequest
	WithType(typeVal string) IGetGlobalLoadBalancerUsageHistoriesRequest
	GetLoadBalancerId() string

	AddUserAgent(agent ...string) IGetGlobalLoadBalancerUsageHistoriesRequest
	ParseUserAgent() string
	ToListQuery() (string, error)
	GetDefaultQuery() string
}

// --------------------------------------------------------

type IListGlobalLoadBalancersRequest interface {
	WithName(name string) IListGlobalLoadBalancersRequest
	WithTags(tags ...string) IListGlobalLoadBalancersRequest
	ToListQuery() (string, error)
	GetDefaultQuery() string

	AddUserAgent(agent ...string) IListGlobalLoadBalancersRequest
	ParseUserAgent() string
}

// --------------------------------------------------------

type ICreateGlobalLoadBalancerRequest interface {
	WithDescription(desc string) ICreateGlobalLoadBalancerRequest
	WithName(name string) ICreateGlobalLoadBalancerRequest
	WithType(typeVal GlobalLoadBalancerType) ICreateGlobalLoadBalancerRequest
	WithGlobalListener(listener ICreateGlobalListenerRequest) ICreateGlobalLoadBalancerRequest
	WithGlobalPool(pool ICreateGlobalPoolRequest) ICreateGlobalLoadBalancerRequest
	WithPackage(packageId string) ICreateGlobalLoadBalancerRequest
	WithPaymentFlow(paymentFlow GlobalLoadBalancerPaymentFlow) ICreateGlobalLoadBalancerRequest

	// WithTags(ptags ...string) ICreateGlobalLoadBalancerRequest
	// WithScheme(pscheme LoadBalancerScheme) ICreateGlobalLoadBalancerRequest
	// WithAutoScalable(pautoScalable bool) ICreateGlobalLoadBalancerRequest
	// WithPackageId(ppackageId string) ICreateGlobalLoadBalancerRequest
	// WithSubnetId(psubnetId string) ICreateGlobalLoadBalancerRequest
	// WithPoc(poc bool) ICreateGlobalLoadBalancerRequest

	AddUserAgent(agent ...string) ICreateGlobalLoadBalancerRequest
	ParseUserAgent() string
	ToRequestBody() interface{}
	ToMap() map[string]interface{}
}

// --------------------------------------------------------

type IDeleteGlobalLoadBalancerRequest interface {
	WithLoadBalancerId(lbId string) IDeleteGlobalLoadBalancerRequest
	GetLoadBalancerId() string // to use in request url

	AddUserAgent(agent ...string) IDeleteGlobalLoadBalancerRequest
	ParseUserAgent() string
}

// --------------------------------------------------------

type IGetGlobalLoadBalancerByIdRequest interface {
	WithLoadBalancerId(lbId string) IGetGlobalLoadBalancerByIdRequest
	GetLoadBalancerId() string // to use in request url

	AddUserAgent(agent ...string) IGetGlobalLoadBalancerByIdRequest
	ParseUserAgent() string
}
