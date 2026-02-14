package v2

import "github.com/dannyota/greennode-community-sdk/v2/greennode/services/common"

type ICreatePoolRequest interface {
	ToRequestBody() any
	WithHealthMonitor(monitor IHealthMonitorRequest) ICreatePoolRequest
	WithMembers(members ...IMemberRequest) ICreatePoolRequest
	WithLoadBalancerID(lbID string) ICreatePoolRequest
	WithAlgorithm(algorithm PoolAlgorithm) ICreatePoolRequest
	ToMap() map[string]any
	GetLoadBalancerID() string
	ParseUserAgent() string
	AddUserAgent(agent ...string) ICreatePoolRequest
}

type IUpdatePoolRequest interface {
	GetPoolID() string
	ToRequestBody() any
	WithHealthMonitor(monitor IHealthMonitorRequest) IUpdatePoolRequest
	WithLoadBalancerID(lbID string) IUpdatePoolRequest
	WithAlgorithm(algorithm PoolAlgorithm) IUpdatePoolRequest
	WithStickiness(v *bool) IUpdatePoolRequest
	WithTLSEncryption(v *bool) IUpdatePoolRequest
	ToMap() map[string]any
	GetLoadBalancerID() string
	ParseUserAgent() string
	AddUserAgent(agent ...string) IUpdatePoolRequest
}

type IGetPoolHealthMonitorByIDRequest interface {
	GetLoadBalancerID() string
	GetPoolID() string
	ParseUserAgent() string
	AddUserAgent(agent ...string) IGetPoolHealthMonitorByIDRequest
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

type IGetPoolByIDRequest interface {
	GetLoadBalancerID() string
	GetPoolID() string
	ParseUserAgent() string
	AddUserAgent(agent ...string) IGetPoolByIDRequest
}

type IListPoolsByLoadBalancerIDRequest interface {
	GetLoadBalancerID() string
	ParseUserAgent() string
	AddUserAgent(agent ...string) IListPoolsByLoadBalancerIDRequest
}

type IUpdatePoolMembersRequest interface {
	WithMembers(members ...IMemberRequest) IUpdatePoolMembersRequest
	ToRequestBody() any
	GetLoadBalancerID() string
	GetPoolID() string
	ParseUserAgent() string
	AddUserAgent(agent ...string) IUpdatePoolMembersRequest
}

type IHealthMonitorRequest interface {
	ToRequestBody() any
	ToMap() map[string]any
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
	ToRequestBody() any
	ToMap() map[string]any
}

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
	HealthCheckProtocolHTTPs   HealthCheckProtocol = "HTTPS"
	HealthCheckProtocolPINGUDP HealthCheckProtocol = "PING-UDP"
)

const (
	HealthCheckMethodGET  HealthCheckMethod = "GET"
	HealthCheckMethodPUT  HealthCheckMethod = "PUT"
	HealthCheckMethodPOST HealthCheckMethod = "POST"
)

const (
	HealthCheckHTTPVersionHttp1       HealthCheckHTTPVersion = "1.0"
	HealthCheckHTTPVersionHttp1Minor1 HealthCheckHTTPVersion = "1.1"
)

const (
	defaultFakeDomainName = "nip.io"
)

func NewCreatePoolRequest(name string, protocol PoolProtocol) *CreatePoolRequest {
	opts := new(CreatePoolRequest)
	opts.PoolName = name
	opts.Algorithm = PoolAlgorithmRoundRobin
	opts.PoolProtocol = protocol
	opts.Members = make([]IMemberRequest, 0)

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

func (r *GetPoolHealthMonitorByIDRequest) AddUserAgent(agent ...string) IGetPoolHealthMonitorByIDRequest {
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
	opts.Members = make([]IMemberRequest, 0)

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
		IpAddress:   ipAddress,
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
	Algorithm     PoolAlgorithm         `json:"algorithm"`
	PoolName      string                `json:"poolName"`
	PoolProtocol  PoolProtocol          `json:"poolProtocol"`
	Stickiness    *bool                 `json:"stickiness,omitempty"`    // only for l7, l4 doesn't have this field => nil
	TLSEncryption *bool                 `json:"tlsEncryption,omitempty"` // only for l7, l4 doesn't have this field => nil
	HealthMonitor IHealthMonitorRequest `json:"healthMonitor"`
	Members       []IMemberRequest      `json:"members"`

	common.LoadBalancerCommon
	common.UserAgent
}

func (r *CreatePoolRequest) AddUserAgent(agent ...string) ICreatePoolRequest {
	r.UserAgent.AddUserAgent(agent...)
	return r
}

type UpdatePoolRequest struct {
	Algorithm     PoolAlgorithm         `json:"algorithm"`
	Stickiness    *bool                 `json:"stickiness,omitempty"`    // only for l7, l4 doesn't have this field => nil
	TLSEncryption *bool                 `json:"tlsEncryption,omitempty"` // only for l7, l4 doesn't have this field => nil
	HealthMonitor IHealthMonitorRequest `json:"healthMonitor"`

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

func (r *ListPoolMembersRequest) AddUserAgent(agent ...string) IListPoolMembersRequest {
	r.UserAgent.AddUserAgent(agent...)
	return r
}

type DeletePoolByIDRequest struct {
	common.UserAgent
	common.LoadBalancerCommon
	common.PoolCommon
}

func (r *DeletePoolByIDRequest) AddUserAgent(agent ...string) IDeletePoolByIDRequest {
	r.UserAgent.AddUserAgent(agent...)
	return r
}

type GetPoolByIDRequest struct {
	common.UserAgent
	common.LoadBalancerCommon
	common.PoolCommon
}

func (r *GetPoolByIDRequest) AddUserAgent(agent ...string) IGetPoolByIDRequest {
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
	IpAddress   string `json:"ipAddress"`
	MonitorPort int    `json:"monitorPort"`
	Name        string `json:"name"`
	Port        int    `json:"port"`
	Weight      int    `json:"weight"`
}

type ListPoolsByLoadBalancerIDRequest struct {
	common.LoadBalancerCommon
	common.UserAgent
}

func (r *ListPoolsByLoadBalancerIDRequest) AddUserAgent(agent ...string) IListPoolsByLoadBalancerIDRequest {
	r.UserAgent.AddUserAgent(agent...)
	return r
}

type UpdatePoolMembersRequest struct {
	Members []IMemberRequest `json:"members"`

	common.UserAgent
	common.LoadBalancerCommon
	common.PoolCommon
}

func (r *UpdatePoolMembersRequest) AddUserAgent(agent ...string) IUpdatePoolMembersRequest {
	r.UserAgent.AddUserAgent(agent...)
	return r
}

func (r *CreatePoolRequest) ToRequestBody() any {
	r.HealthMonitor = r.HealthMonitor.(*HealthMonitor).toRequestBody()
	return r
}

func (h *HealthMonitor) toRequestBody() IHealthMonitorRequest {
	switch h.HealthCheckProtocol {
	case HealthCheckProtocolPINGUDP, HealthCheckProtocolTCP:
		h.HealthCheckPath = nil
		h.HTTPVersion = nil
		h.SuccessCode = nil
		h.HealthCheckMethod = nil
		h.DomainName = nil

	case HealthCheckProtocolHTTP, HealthCheckProtocolHTTPs:
		if h.HTTPVersion != nil {
			switch opt := *h.HTTPVersion; opt {
			case HealthCheckHTTPVersionHttp1:
				h.DomainName = nil
			case HealthCheckHTTPVersionHttp1Minor1:
				if h.DomainName == nil ||
					(h.DomainName != nil && len(*h.DomainName) < 1) {

					fakeDomainName := defaultFakeDomainName
					h.DomainName = &fakeDomainName
				}
			}
		}
	}

	return h
}

func (h *HealthMonitor) AddUserAgent(agent ...string) IHealthMonitorRequest {
	h.UserAgent.AddUserAgent(agent...)
	return h
}

func (h *HealthMonitor) WithHealthCheckProtocol(protocol HealthCheckProtocol) IHealthMonitorRequest {
	h.HealthCheckProtocol = protocol
	return h
}

func (r *CreatePoolRequest) WithHealthMonitor(monitor IHealthMonitorRequest) ICreatePoolRequest {
	r.HealthMonitor = monitor
	return r
}

func (r *CreatePoolRequest) WithMembers(members ...IMemberRequest) ICreatePoolRequest {
	r.Members = append(r.Members, members...)
	return r
}

func (r *CreatePoolRequest) WithLoadBalancerID(lbID string) ICreatePoolRequest {
	r.LoadBalancerID = lbID
	return r
}

func (r *CreatePoolRequest) ToMap() map[string]any {
	return map[string]any{
		"algorithm":     r.Algorithm,
		"poolName":      r.PoolName,
		"poolProtocol":  r.PoolProtocol,
		"stickiness":    r.Stickiness,
		"tlsEncryption": r.TLSEncryption,
		"healthMonitor": r.HealthMonitor.ToMap(),
		"members": func() []map[string]any {
			members := make([]map[string]any, 0, len(r.Members))
			for _, member := range r.Members {
				members = append(members, member.ToMap())
			}
			return members
		}(),
	}
}

func (r *CreatePoolRequest) WithAlgorithm(algorithm PoolAlgorithm) ICreatePoolRequest {
	r.Algorithm = algorithm
	return r
}

func (r *UpdatePoolRequest) ToMap() map[string]any {
	return map[string]any{
		"algorithm":     r.Algorithm,
		"stickiness":    r.Stickiness,
		"tlsEncryption": r.TLSEncryption,
		"healthMonitor": r.HealthMonitor.ToMap(),
	}
}

func (r *UpdatePoolRequest) ToRequestBody() any {
	r.HealthMonitor = r.HealthMonitor.(*HealthMonitor).toRequestBody()
	return r
}

func (r *UpdatePoolRequest) WithAlgorithm(algorithm PoolAlgorithm) IUpdatePoolRequest {
	r.Algorithm = algorithm
	return r
}

func (r *UpdatePoolRequest) WithHealthMonitor(monitor IHealthMonitorRequest) IUpdatePoolRequest {
	r.HealthMonitor = monitor
	return r
}

func (r *UpdatePoolRequest) WithLoadBalancerID(lbID string) IUpdatePoolRequest {
	r.LoadBalancerID = lbID
	return r
}

func (r *UpdatePoolRequest) WithTLSEncryption(v *bool) IUpdatePoolRequest {
	r.TLSEncryption = v
	return r
}

func (r *UpdatePoolRequest) WithStickiness(v *bool) IUpdatePoolRequest {
	r.Stickiness = v
	return r
}

func (r *UpdatePoolRequest) AddUserAgent(agent ...string) IUpdatePoolRequest {
	r.UserAgent.AddUserAgent(agent...)
	return r
}

func (h *HealthMonitor) ToRequestBody() any {
	return h
}

func (h *HealthMonitor) WithHealthyThreshold(ht int) IHealthMonitorRequest {
	if ht < 1 {
		ht = 3
	}

	h.HealthyThreshold = ht
	return h
}

func (h *HealthMonitor) WithUnhealthyThreshold(uht int) IHealthMonitorRequest {
	if uht < 1 {
		uht = 3
	}

	h.UnhealthyThreshold = uht
	return h
}

func (h *HealthMonitor) WithInterval(interval int) IHealthMonitorRequest {
	if interval < 1 {
		interval = 30
	}

	h.Interval = interval
	return h
}

func (h *HealthMonitor) WithTimeout(to int) IHealthMonitorRequest {
	if to < 1 {
		to = 5
	}

	h.Timeout = to
	return h
}

func (h *HealthMonitor) WithHealthCheckMethod(method *HealthCheckMethod) IHealthMonitorRequest {
	h.HealthCheckMethod = method
	return h
}

func (h *HealthMonitor) WithHTTPVersion(version *HealthCheckHTTPVersion) IHealthMonitorRequest {
	h.HTTPVersion = version
	return h
}

func (h *HealthMonitor) WithHealthCheckPath(path *string) IHealthMonitorRequest {
	h.HealthCheckPath = path
	return h
}

func (h *HealthMonitor) WithDomainName(domain *string) IHealthMonitorRequest {
	h.DomainName = domain
	return h
}

func (h *HealthMonitor) WithSuccessCode(code *string) IHealthMonitorRequest {
	h.SuccessCode = code
	return h
}

func (h *HealthMonitor) ToMap() map[string]any {
	return map[string]any{
		"healthCheckProtocol": h.HealthCheckProtocol,
		"healthyThreshold":    h.HealthyThreshold,
		"unhealthyThreshold":  h.UnhealthyThreshold,
		"interval":            h.Interval,
		"timeout":             h.Timeout,
		"healthCheckMethod":   h.HealthCheckMethod,
		"httpVersion":         h.HTTPVersion,
		"healthCheckPath":     h.HealthCheckPath,
		"domainName":          h.DomainName,
		"successCode":         h.SuccessCode,
	}
}

func (m *Member) ToRequestBody() any {
	return m
}

func (m *Member) ToMap() map[string]any {
	return map[string]any{
		"backup":      m.Backup,
		"ipAddress":   m.IpAddress,
		"monitorPort": m.MonitorPort,
		"name":        m.Name,
		"port":        m.Port,
		"weight":      m.Weight,
	}
}

func (r *UpdatePoolMembersRequest) WithMembers(members ...IMemberRequest) IUpdatePoolMembersRequest {
	r.Members = append(r.Members, members...)
	return r
}

func (r *UpdatePoolMembersRequest) ToRequestBody() any {
	return r
}
