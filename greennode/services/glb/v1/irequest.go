package v1

type IListGlobalPoolsRequest interface {
	WithLoadBalancerID(lbID string) IListGlobalPoolsRequest
	GetLoadBalancerID() string

	AddUserAgent(agent ...string) IListGlobalPoolsRequest
	ParseUserAgent() string
}


type ICreateGlobalPoolRequest interface {
	WithAlgorithm(algorithm GlobalPoolAlgorithm) ICreateGlobalPoolRequest
	WithDescription(desc string) ICreateGlobalPoolRequest
	WithName(name string) ICreateGlobalPoolRequest
	WithProtocol(protocol GlobalPoolProtocol) ICreateGlobalPoolRequest
	WithHealthMonitor(monitor IGlobalHealthMonitorRequest) ICreateGlobalPoolRequest
	WithMembers(members ...ICreateGlobalPoolMemberRequest) ICreateGlobalPoolRequest

	WithLoadBalancerID(lbID string) ICreateGlobalPoolRequest
	GetLoadBalancerID() string // to use in request url

	AddUserAgent(agent ...string) ICreateGlobalPoolRequest
	ParseUserAgent() string
	ToRequestBody() any
	ToMap() map[string]any
}

type IGlobalHealthMonitorRequest interface {
	WithHealthyThreshold(ht int) IGlobalHealthMonitorRequest
	WithInterval(interval int) IGlobalHealthMonitorRequest
	WithProtocol(protocol GlobalPoolHealthCheckProtocol) IGlobalHealthMonitorRequest
	WithTimeout(to int) IGlobalHealthMonitorRequest
	WithUnhealthyThreshold(uht int) IGlobalHealthMonitorRequest

	// http, https
	WithHealthCheckMethod(method *GlobalPoolHealthCheckMethod) IGlobalHealthMonitorRequest
	WithHTTPVersion(version *GlobalPoolHealthCheckHTTPVersion) IGlobalHealthMonitorRequest
	WithPath(path *string) IGlobalHealthMonitorRequest
	WithSuccessCode(code *string) IGlobalHealthMonitorRequest
	WithDomainName(domain *string) IGlobalHealthMonitorRequest

	AddUserAgent(agent ...string) IGlobalHealthMonitorRequest
	ParseUserAgent() string
	ToRequestBody() any
	ToMap() map[string]any
}
type ICreateGlobalPoolMemberRequest interface {
	WithName(name string) ICreateGlobalPoolMemberRequest
	WithDescription(desc string) ICreateGlobalPoolMemberRequest
	WithRegion(region string) ICreateGlobalPoolMemberRequest
	WithVPCID(vpcID string) ICreateGlobalPoolMemberRequest
	WithTrafficDial(dial int) ICreateGlobalPoolMemberRequest
	WithMembers(members ...IGlobalMemberRequest) ICreateGlobalPoolMemberRequest
	WithType(typeVal GlobalPoolMemberType) ICreateGlobalPoolMemberRequest

	WithLoadBalancerID(lbID string) ICreateGlobalPoolMemberRequest
	WithPoolID(poolID string) ICreateGlobalPoolMemberRequest
	GetLoadBalancerID() string // to use in request url
	GetPoolID() string         // to use in request url

	AddUserAgent(agent ...string) ICreateGlobalPoolMemberRequest
	ParseUserAgent() string
	ToRequestBody() any
	ToMap() map[string]any
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
	ToRequestBody() any
	ToMap() map[string]any
}


type IUpdateGlobalPoolRequest interface {
	WithAlgorithm(algorithm GlobalPoolAlgorithm) IUpdateGlobalPoolRequest
	WithHealthMonitor(monitor IGlobalHealthMonitorRequest) IUpdateGlobalPoolRequest

	WithLoadBalancerID(lbID string) IUpdateGlobalPoolRequest
	WithPoolID(poolID string) IUpdateGlobalPoolRequest
	GetLoadBalancerID() string // to use in request url
	GetPoolID() string

	AddUserAgent(agent ...string) IUpdateGlobalPoolRequest
	ParseUserAgent() string
	ToRequestBody() any
	ToMap() map[string]any
}


type IDeleteGlobalPoolRequest interface {
	WithLoadBalancerID(lbID string) IDeleteGlobalPoolRequest
	WithPoolID(poolID string) IDeleteGlobalPoolRequest
	GetLoadBalancerID() string // to use in request url
	GetPoolID() string

	AddUserAgent(agent ...string) IDeleteGlobalPoolRequest
	ParseUserAgent() string
}


type IListGlobalPoolMembersRequest interface {
	WithLoadBalancerID(lbID string) IListGlobalPoolMembersRequest
	WithPoolID(poolID string) IListGlobalPoolMembersRequest
	GetLoadBalancerID() string // to use in request url
	GetPoolID() string

	AddUserAgent(agent ...string) IListGlobalPoolMembersRequest
	ParseUserAgent() string
}


type IGetGlobalPoolMemberRequest interface {
	WithLoadBalancerID(lbID string) IGetGlobalPoolMemberRequest
	WithPoolID(poolID string) IGetGlobalPoolMemberRequest
	WithPoolMemberID(poolMemberID string) IGetGlobalPoolMemberRequest
	GetLoadBalancerID() string // to use in request url
	GetPoolID() string
	GetPoolMemberID() string

	AddUserAgent(agent ...string) IGetGlobalPoolMemberRequest
	ParseUserAgent() string
}


type IDeleteGlobalPoolMemberRequest interface {
	WithLoadBalancerID(lbID string) IDeleteGlobalPoolMemberRequest
	WithPoolID(poolID string) IDeleteGlobalPoolMemberRequest
	WithPoolMemberID(poolMemberID string) IDeleteGlobalPoolMemberRequest
	GetLoadBalancerID() string // to use in request url
	GetPoolID() string
	GetPoolMemberID() string

	AddUserAgent(agent ...string) IDeleteGlobalPoolMemberRequest
	ParseUserAgent() string
}


type IPatchGlobalPoolMembersRequest interface {
	WithBulkAction(action ...IBulkActionRequest) IPatchGlobalPoolMembersRequest

	WithLoadBalancerID(lbID string) IPatchGlobalPoolMembersRequest
	WithPoolID(poolID string) IPatchGlobalPoolMembersRequest
	GetLoadBalancerID() string // to use in request url
	GetPoolID() string

	AddUserAgent(agent ...string) IPatchGlobalPoolMembersRequest
	ParseUserAgent() string
	ToRequestBody() any
	ToMap() map[string]any
}

type IBulkActionRequest interface {
	ToRequestBody() any
	ToMap() map[string]any
}


type IListGlobalListenersRequest interface {
	WithLoadBalancerID(lbID string) IListGlobalListenersRequest
	GetLoadBalancerID() string // to use in request url

	AddUserAgent(agent ...string) IListGlobalListenersRequest
	ParseUserAgent() string
}


type IGetGlobalListenerRequest interface {
	WithLoadBalancerID(lbID string) IGetGlobalListenerRequest
	WithListenerID(listenerID string) IGetGlobalListenerRequest
	GetLoadBalancerID() string // to use in request url
	GetListenerID() string

	AddUserAgent(agent ...string) IGetGlobalListenerRequest
	ParseUserAgent() string
}


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
	WithGlobalPoolID(poolID string) ICreateGlobalListenerRequest

	WithLoadBalancerID(lbid string) ICreateGlobalListenerRequest
	GetLoadBalancerID() string
	// AddCidrs(pcidrs ...string) ICreateGlobalListenerRequest

	AddUserAgent(agent ...string) ICreateGlobalListenerRequest
	ParseUserAgent() string
	ToRequestBody() any
	ToMap() map[string]any
}


type IUpdateGlobalListenerRequest interface {
	WithAllowedCidrs(cidrs ...string) IUpdateGlobalListenerRequest
	WithTimeoutClient(toc int) IUpdateGlobalListenerRequest
	WithTimeoutMember(tom int) IUpdateGlobalListenerRequest
	WithTimeoutConnection(toc int) IUpdateGlobalListenerRequest
	WithHeaders(headers ...string) IUpdateGlobalListenerRequest
	WithGlobalPoolID(poolID string) IUpdateGlobalListenerRequest

	WithLoadBalancerID(lbID string) IUpdateGlobalListenerRequest
	WithListenerID(listenerID string) IUpdateGlobalListenerRequest
	GetLoadBalancerID() string // to use in request url
	GetListenerID() string

	AddUserAgent(agent ...string) IUpdateGlobalListenerRequest
	ParseUserAgent() string
	ToRequestBody() any
	ToMap() map[string]any
}


type IDeleteGlobalListenerRequest interface {
	WithLoadBalancerID(lbID string) IDeleteGlobalListenerRequest
	WithListenerID(listenerID string) IDeleteGlobalListenerRequest
	GetLoadBalancerID() string // to use in request url
	GetListenerID() string

	AddUserAgent(agent ...string) IDeleteGlobalListenerRequest
	ParseUserAgent() string
}


type IListGlobalPackagesRequest interface {
	AddUserAgent(agent ...string) IListGlobalPackagesRequest
	ParseUserAgent() string
}


type IListGlobalRegionsRequest interface {
	AddUserAgent(agent ...string) IListGlobalRegionsRequest
	ParseUserAgent() string
}


type IGetGlobalLoadBalancerUsageHistoriesRequest interface {
	WithLoadBalancerID(lbID string) IGetGlobalLoadBalancerUsageHistoriesRequest
	WithFrom(from string) IGetGlobalLoadBalancerUsageHistoriesRequest
	WithTo(to string) IGetGlobalLoadBalancerUsageHistoriesRequest
	WithType(typeVal string) IGetGlobalLoadBalancerUsageHistoriesRequest
	GetLoadBalancerID() string

	AddUserAgent(agent ...string) IGetGlobalLoadBalancerUsageHistoriesRequest
	ParseUserAgent() string
	ToListQuery() (string, error)
	GetDefaultQuery() string
}


type IListGlobalLoadBalancersRequest interface {
	WithName(name string) IListGlobalLoadBalancersRequest
	WithTags(tags ...string) IListGlobalLoadBalancersRequest
	ToListQuery() (string, error)
	GetDefaultQuery() string

	AddUserAgent(agent ...string) IListGlobalLoadBalancersRequest
	ParseUserAgent() string
}


type ICreateGlobalLoadBalancerRequest interface {
	WithDescription(desc string) ICreateGlobalLoadBalancerRequest
	WithName(name string) ICreateGlobalLoadBalancerRequest
	WithType(typeVal GlobalLoadBalancerType) ICreateGlobalLoadBalancerRequest
	WithGlobalListener(listener ICreateGlobalListenerRequest) ICreateGlobalLoadBalancerRequest
	WithGlobalPool(pool ICreateGlobalPoolRequest) ICreateGlobalLoadBalancerRequest
	WithPackage(packageID string) ICreateGlobalLoadBalancerRequest
	WithPaymentFlow(paymentFlow GlobalLoadBalancerPaymentFlow) ICreateGlobalLoadBalancerRequest

	// WithTags(ptags ...string) ICreateGlobalLoadBalancerRequest
	// WithScheme(pscheme LoadBalancerScheme) ICreateGlobalLoadBalancerRequest
	// WithAutoScalable(pautoScalable bool) ICreateGlobalLoadBalancerRequest
	// WithPackageId(ppackageId string) ICreateGlobalLoadBalancerRequest
	// WithSubnetId(psubnetId string) ICreateGlobalLoadBalancerRequest
	// WithPoc(poc bool) ICreateGlobalLoadBalancerRequest

	AddUserAgent(agent ...string) ICreateGlobalLoadBalancerRequest
	ParseUserAgent() string
	ToRequestBody() any
	ToMap() map[string]any
}


type IDeleteGlobalLoadBalancerRequest interface {
	WithLoadBalancerID(lbID string) IDeleteGlobalLoadBalancerRequest
	GetLoadBalancerID() string // to use in request url

	AddUserAgent(agent ...string) IDeleteGlobalLoadBalancerRequest
	ParseUserAgent() string
}


type IGetGlobalLoadBalancerByIDRequest interface {
	WithLoadBalancerID(lbID string) IGetGlobalLoadBalancerByIDRequest
	GetLoadBalancerID() string // to use in request url

	AddUserAgent(agent ...string) IGetGlobalLoadBalancerByIDRequest
	ParseUserAgent() string
}
