package v2

import "github.com/dannyota/greennode-community-sdk/v2/greennode/services/common"

const (
	PoolAlgorithmRoundRobin PoolAlgorithm = "ROUND_ROBIN"
	PoolAlgorithmLeastConn  PoolAlgorithm = "LEAST_CONNECTIONS"
	PoolAlgorithmSourceIP   PoolAlgorithm = "SOURCE_IP"
)

const (
	PoolProtocolTCP   PoolProtocol = "TCP"
	PoolProtocolUDP   PoolProtocol = "UDP"
	PoolProtocolHTTP  PoolProtocol = "HTTP"
	PoolProtocolProxy PoolProtocol = "PROXY"
)

const (
	HealthCheckProtocolTCP     HealthCheckProtocol = "TCP"
	HealthCheckProtocolHTTP    HealthCheckProtocol = "HTTP"
	HealthCheckProtocolHTTPS   HealthCheckProtocol = "HTTPS"
	HealthCheckProtocolPINGUDP HealthCheckProtocol = "PING-UDP"
)

const (
	HealthCheckMethodGET  HealthCheckMethod = "GET"
	HealthCheckMethodPUT  HealthCheckMethod = "PUT"
	HealthCheckMethodPOST HealthCheckMethod = "POST"
)

const (
	HealthCheckHTTPVersionHTTP1       HealthCheckHTTPVersion = "1.0"
	HealthCheckHTTPVersionHTTP1Minor1 HealthCheckHTTPVersion = "1.1"
)

const (
	defaultFakeDomainName = "nip.io"
)

func NewCreatePoolRequest(name string, protocol PoolProtocol) *CreatePoolRequest {
	opts := new(CreatePoolRequest)
	opts.PoolName = name
	opts.Algorithm = PoolAlgorithmRoundRobin
	opts.PoolProtocol = protocol
	opts.Members = make([]*Member, 0)

	return opts
}

func NewUpdatePoolRequest(lbID, poolID string) *UpdatePoolRequest {
	opts := new(UpdatePoolRequest)
	opts.LoadBalancerID = lbID
	opts.PoolID = poolID

	return opts
}

func NewGetPoolHealthMonitorByIDRequest(lbID, poolID string) *GetPoolHealthMonitorByIDRequest {
	opts := new(GetPoolHealthMonitorByIDRequest)
	opts.LoadBalancerID = lbID
	opts.PoolID = poolID

	return opts
}

func (r *GetPoolHealthMonitorByIDRequest) AddUserAgent(agent ...string) *GetPoolHealthMonitorByIDRequest {
	r.UserAgent.AddUserAgent(agent...)
	return r
}

func NewListPoolsByLoadBalancerIDRequest(lbID string) *ListPoolsByLoadBalancerIDRequest {
	opts := new(ListPoolsByLoadBalancerIDRequest)
	opts.LoadBalancerID = lbID

	return opts
}

func NewUpdatePoolMembersRequest(lbID, poolID string) *UpdatePoolMembersRequest {
	opts := new(UpdatePoolMembersRequest)
	opts.LoadBalancerID = lbID
	opts.PoolID = poolID
	opts.Members = make([]*Member, 0)

	return opts
}

func NewListPoolMembersRequest(lbID, poolID string) *ListPoolMembersRequest {
	opts := new(ListPoolMembersRequest)
	opts.LoadBalancerID = lbID
	opts.PoolID = poolID

	return opts
}

func NewDeletePoolByIDRequest(lbID, poolID string) *DeletePoolByIDRequest {
	opts := new(DeletePoolByIDRequest)
	opts.LoadBalancerID = lbID
	opts.PoolID = poolID

	return opts
}

func NewHealthMonitor(checkProtocol HealthCheckProtocol) *HealthMonitor {
	opts := new(HealthMonitor)
	opts.HealthCheckProtocol = checkProtocol
	opts.HealthyThreshold = 3
	opts.UnhealthyThreshold = 3
	opts.Interval = 30
	opts.Timeout = 5

	return opts
}

func NewMember(name, ipAddress string, port int, monitorPort int) *Member {
	return &Member{
		Backup:      false,
		IPAddress:   ipAddress,
		MonitorPort: monitorPort,
		Name:        name,
		Port:        port,
		Weight:      1,
	}
}

func NewGetPoolByIDRequest(lbID, poolID string) *GetPoolByIDRequest {
	opts := new(GetPoolByIDRequest)
	opts.LoadBalancerID = lbID
	opts.PoolID = poolID

	return opts
}

type (
	PoolAlgorithm          string
	PoolProtocol           string
	HealthCheckProtocol    string
	HealthCheckMethod      string
	HealthCheckHTTPVersion string
)

type CreatePoolRequest struct {
	Algorithm     PoolAlgorithm  `json:"algorithm"`
	PoolName      string         `json:"poolName"`
	PoolProtocol  PoolProtocol   `json:"poolProtocol"`
	Stickiness    *bool          `json:"stickiness,omitempty"`    // only for l7, l4 doesn't have this field => nil
	TLSEncryption *bool          `json:"tlsEncryption,omitempty"` // only for l7, l4 doesn't have this field => nil
	HealthMonitor *HealthMonitor `json:"healthMonitor"`
	Members       []*Member      `json:"members"`

	common.LoadBalancerCommon
	common.UserAgent
}

func (r *CreatePoolRequest) AddUserAgent(agent ...string) *CreatePoolRequest {
	r.UserAgent.AddUserAgent(agent...)
	return r
}

type UpdatePoolRequest struct {
	Algorithm     PoolAlgorithm  `json:"algorithm"`
	Stickiness    *bool          `json:"stickiness,omitempty"`    // only for l7, l4 doesn't have this field => nil
	TLSEncryption *bool          `json:"tlsEncryption,omitempty"` // only for l7, l4 doesn't have this field => nil
	HealthMonitor *HealthMonitor `json:"healthMonitor"`

	common.LoadBalancerCommon
	common.PoolCommon
	common.UserAgent
}

type GetPoolHealthMonitorByIDRequest struct {
	common.LoadBalancerCommon
	common.PoolCommon
	common.UserAgent
}

type ListPoolMembersRequest struct {
	common.UserAgent
	common.LoadBalancerCommon
	common.PoolCommon
}

func (r *ListPoolMembersRequest) AddUserAgent(agent ...string) *ListPoolMembersRequest {
	r.UserAgent.AddUserAgent(agent...)
	return r
}

type DeletePoolByIDRequest struct {
	common.UserAgent
	common.LoadBalancerCommon
	common.PoolCommon
}

func (r *DeletePoolByIDRequest) AddUserAgent(agent ...string) *DeletePoolByIDRequest {
	r.UserAgent.AddUserAgent(agent...)
	return r
}

type GetPoolByIDRequest struct {
	common.UserAgent
	common.LoadBalancerCommon
	common.PoolCommon
}

func (r *GetPoolByIDRequest) AddUserAgent(agent ...string) *GetPoolByIDRequest {
	r.UserAgent.AddUserAgent(agent...)
	return r
}

type HealthMonitor struct {
	HealthCheckProtocol HealthCheckProtocol     `json:"healthCheckProtocol"`
	HealthyThreshold    int                     `json:"healthyThreshold"`
	UnhealthyThreshold  int                     `json:"unhealthyThreshold"`
	Interval            int                     `json:"interval"`
	Timeout             int                     `json:"timeout"`
	HealthCheckMethod   *HealthCheckMethod      `json:"healthCheckMethod,omitempty"`
	HTTPVersion         *HealthCheckHTTPVersion `json:"httpVersion,omitempty"`
	HealthCheckPath     *string                 `json:"healthCheckPath,omitempty"`
	DomainName          *string                 `json:"domainName,omitempty"`
	SuccessCode         *string                 `json:"successCode,omitempty"`

	common.UserAgent
}

type Member struct {
	Backup      bool   `json:"backup"`
	IPAddress   string `json:"ipAddress"`
	MonitorPort int    `json:"monitorPort"`
	Name        string `json:"name"`
	Port        int    `json:"port"`
	Weight      int    `json:"weight"`
}

type ListPoolsByLoadBalancerIDRequest struct {
	common.LoadBalancerCommon
	common.UserAgent
}

func (r *ListPoolsByLoadBalancerIDRequest) AddUserAgent(agent ...string) *ListPoolsByLoadBalancerIDRequest {
	r.UserAgent.AddUserAgent(agent...)
	return r
}

type UpdatePoolMembersRequest struct {
	Members []*Member `json:"members"`

	common.UserAgent
	common.LoadBalancerCommon
	common.PoolCommon
}

func (r *UpdatePoolMembersRequest) AddUserAgent(agent ...string) *UpdatePoolMembersRequest {
	r.UserAgent.AddUserAgent(agent...)
	return r
}

func (r *CreatePoolRequest) prepare() {
	r.HealthMonitor.prepare()
}

func (h *HealthMonitor) prepare() {
	switch h.HealthCheckProtocol {
	case HealthCheckProtocolPINGUDP, HealthCheckProtocolTCP:
		h.HealthCheckPath = nil
		h.HTTPVersion = nil
		h.SuccessCode = nil
		h.HealthCheckMethod = nil
		h.DomainName = nil

	case HealthCheckProtocolHTTP, HealthCheckProtocolHTTPS:
		if h.HTTPVersion != nil {
			switch opt := *h.HTTPVersion; opt {
			case HealthCheckHTTPVersionHTTP1:
				h.DomainName = nil
			case HealthCheckHTTPVersionHTTP1Minor1:
				if h.DomainName == nil ||
					(h.DomainName != nil && len(*h.DomainName) < 1) {

					fakeDomainName := defaultFakeDomainName
					h.DomainName = &fakeDomainName
				}
			}
		}
	}
}

func (h *HealthMonitor) AddUserAgent(agent ...string) *HealthMonitor {
	h.UserAgent.AddUserAgent(agent...)
	return h
}

func (h *HealthMonitor) WithHealthCheckProtocol(protocol HealthCheckProtocol) *HealthMonitor {
	h.HealthCheckProtocol = protocol
	return h
}

func (r *CreatePoolRequest) WithHealthMonitor(monitor *HealthMonitor) *CreatePoolRequest {
	r.HealthMonitor = monitor
	return r
}

func (r *CreatePoolRequest) WithMembers(members ...*Member) *CreatePoolRequest {
	r.Members = append(r.Members, members...)
	return r
}

func (r *CreatePoolRequest) WithLoadBalancerID(lbID string) *CreatePoolRequest {
	r.LoadBalancerID = lbID
	return r
}

func (r *CreatePoolRequest) WithAlgorithm(algorithm PoolAlgorithm) *CreatePoolRequest {
	r.Algorithm = algorithm
	return r
}

func (r *UpdatePoolRequest) prepare() {
	r.HealthMonitor.prepare()
}

func (r *UpdatePoolRequest) WithAlgorithm(algorithm PoolAlgorithm) *UpdatePoolRequest {
	r.Algorithm = algorithm
	return r
}

func (r *UpdatePoolRequest) WithHealthMonitor(monitor *HealthMonitor) *UpdatePoolRequest {
	r.HealthMonitor = monitor
	return r
}

func (r *UpdatePoolRequest) WithLoadBalancerID(lbID string) *UpdatePoolRequest {
	r.LoadBalancerID = lbID
	return r
}

func (r *UpdatePoolRequest) WithTLSEncryption(v *bool) *UpdatePoolRequest {
	r.TLSEncryption = v
	return r
}

func (r *UpdatePoolRequest) WithStickiness(v *bool) *UpdatePoolRequest {
	r.Stickiness = v
	return r
}

func (r *UpdatePoolRequest) AddUserAgent(agent ...string) *UpdatePoolRequest {
	r.UserAgent.AddUserAgent(agent...)
	return r
}

func (h *HealthMonitor) WithHealthyThreshold(ht int) *HealthMonitor {
	if ht < 1 {
		ht = 3
	}

	h.HealthyThreshold = ht
	return h
}

func (h *HealthMonitor) WithUnhealthyThreshold(uht int) *HealthMonitor {
	if uht < 1 {
		uht = 3
	}

	h.UnhealthyThreshold = uht
	return h
}

func (h *HealthMonitor) WithInterval(interval int) *HealthMonitor {
	if interval < 1 {
		interval = 30
	}

	h.Interval = interval
	return h
}

func (h *HealthMonitor) WithTimeout(to int) *HealthMonitor {
	if to < 1 {
		to = 5
	}

	h.Timeout = to
	return h
}

func (h *HealthMonitor) WithHealthCheckMethod(method *HealthCheckMethod) *HealthMonitor {
	h.HealthCheckMethod = method
	return h
}

func (h *HealthMonitor) WithHTTPVersion(version *HealthCheckHTTPVersion) *HealthMonitor {
	h.HTTPVersion = version
	return h
}

func (h *HealthMonitor) WithHealthCheckPath(path *string) *HealthMonitor {
	h.HealthCheckPath = path
	return h
}

func (h *HealthMonitor) WithDomainName(domain *string) *HealthMonitor {
	h.DomainName = domain
	return h
}

func (h *HealthMonitor) WithSuccessCode(code *string) *HealthMonitor {
	h.SuccessCode = code
	return h
}

func (r *UpdatePoolMembersRequest) WithMembers(members ...*Member) *UpdatePoolMembersRequest {
	r.Members = append(r.Members, members...)
	return r
}

