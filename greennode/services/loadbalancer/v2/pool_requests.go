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

func NewCreatePoolRequest(name string, protocol PoolProtocol) ICreatePoolRequest {
	opts := new(CreatePoolRequest)
	opts.PoolName = name
	opts.Algorithm = PoolAlgorithmRoundRobin
	opts.PoolProtocol = protocol
	opts.Members = make([]IMemberRequest, 0)

	return opts
}

func NewUpdatePoolRequest(lbID, poolID string) IUpdatePoolRequest {
	opts := new(UpdatePoolRequest)
	opts.LoadBalancerID = lbID
	opts.PoolID = poolID

	return opts
}

func NewGetPoolHealthMonitorByIDRequest(lbID, poolID string) IGetPoolHealthMonitorByIDRequest {
	opts := new(GetPoolHealthMonitorByIDRequest)
	opts.LoadBalancerID = lbID
	opts.PoolID = poolID

	return opts
}

func (s *GetPoolHealthMonitorByIDRequest) AddUserAgent(agent ...string) IGetPoolHealthMonitorByIDRequest {
	s.UserAgent.AddUserAgent(agent...)
	return s
}

func NewListPoolsByLoadBalancerIDRequest(lbID string) IListPoolsByLoadBalancerIDRequest {
	opts := new(ListPoolsByLoadBalancerIDRequest)
	opts.LoadBalancerID = lbID

	return opts
}

func NewUpdatePoolMembersRequest(lbID, poolID string) IUpdatePoolMembersRequest {
	opts := new(UpdatePoolMembersRequest)
	opts.LoadBalancerID = lbID
	opts.PoolID = poolID
	opts.Members = make([]IMemberRequest, 0)

	return opts
}

func NewListPoolMembersRequest(lbID, poolID string) IListPoolMembersRequest {
	opts := new(ListPoolMembersRequest)
	opts.LoadBalancerID = lbID
	opts.PoolID = poolID

	return opts
}

func NewDeletePoolByIDRequest(lbID, poolID string) IDeletePoolByIDRequest {
	opts := new(DeletePoolByIDRequest)
	opts.LoadBalancerID = lbID
	opts.PoolID = poolID

	return opts
}

func NewHealthMonitor(checkProtocol HealthCheckProtocol) IHealthMonitorRequest {
	opts := new(HealthMonitor)
	opts.HealthCheckProtocol = checkProtocol
	opts.HealthyThreshold = 3
	opts.UnhealthyThreshold = 3
	opts.Interval = 30
	opts.Timeout = 5

	return opts
}

func NewMember(name, ipAddress string, port int, monitorPort int) IMemberRequest {
	return &Member{
		Backup:      false,
		IpAddress:   ipAddress,
		MonitorPort: monitorPort,
		Name:        name,
		Port:        port,
		Weight:      1,
	}
}

func NewGetPoolByIDRequest(lbID, poolID string) IGetPoolByIDRequest {
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

func (s *CreatePoolRequest) AddUserAgent(agent ...string) ICreatePoolRequest {
	s.UserAgent.AddUserAgent(agent...)
	return s
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

func (s *ListPoolMembersRequest) AddUserAgent(agent ...string) IListPoolMembersRequest {
	s.UserAgent.AddUserAgent(agent...)
	return s
}

type DeletePoolByIDRequest struct {
	common.UserAgent
	common.LoadBalancerCommon
	common.PoolCommon
}

func (s *DeletePoolByIDRequest) AddUserAgent(agent ...string) IDeletePoolByIDRequest {
	s.UserAgent.AddUserAgent(agent...)
	return s
}

type GetPoolByIDRequest struct {
	common.UserAgent
	common.LoadBalancerCommon
	common.PoolCommon
}

func (s *GetPoolByIDRequest) AddUserAgent(agent ...string) IGetPoolByIDRequest {
	s.UserAgent.AddUserAgent(agent...)
	return s
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

func (s *ListPoolsByLoadBalancerIDRequest) AddUserAgent(agent ...string) IListPoolsByLoadBalancerIDRequest {
	s.UserAgent.AddUserAgent(agent...)
	return s
}

type UpdatePoolMembersRequest struct {
	Members []IMemberRequest `json:"members"`

	common.UserAgent
	common.LoadBalancerCommon
	common.PoolCommon
}

func (s *UpdatePoolMembersRequest) AddUserAgent(agent ...string) IUpdatePoolMembersRequest {
	s.UserAgent.AddUserAgent(agent...)
	return s
}

func (s *CreatePoolRequest) ToRequestBody() interface{} {
	s.HealthMonitor = s.HealthMonitor.(*HealthMonitor).toRequestBody()
	return s
}

func (s *HealthMonitor) toRequestBody() IHealthMonitorRequest {
	switch s.HealthCheckProtocol {
	case HealthCheckProtocolPINGUDP, HealthCheckProtocolTCP:
		s.HealthCheckPath = nil
		s.HTTPVersion = nil
		s.SuccessCode = nil
		s.HealthCheckMethod = nil
		s.DomainName = nil

	case HealthCheckProtocolHTTP, HealthCheckProtocolHTTPs:
		if s.HTTPVersion != nil {
			switch opt := *s.HTTPVersion; opt {
			case HealthCheckHTTPVersionHttp1:
				s.DomainName = nil
			case HealthCheckHTTPVersionHttp1Minor1:
				if s.DomainName == nil ||
					(s.DomainName != nil && len(*s.DomainName) < 1) {

					fakeDomainName := defaultFakeDomainName
					s.DomainName = &fakeDomainName
				}
			}
		}
	}

	return s
}

func (s *HealthMonitor) AddUserAgent(agent ...string) IHealthMonitorRequest {
	s.UserAgent.AddUserAgent(agent...)
	return s
}

func (s *HealthMonitor) WithHealthCheckProtocol(protocol HealthCheckProtocol) IHealthMonitorRequest {
	s.HealthCheckProtocol = protocol
	return s
}

func (s *CreatePoolRequest) WithHealthMonitor(monitor IHealthMonitorRequest) ICreatePoolRequest {
	s.HealthMonitor = monitor
	return s
}

func (s *CreatePoolRequest) WithMembers(members ...IMemberRequest) ICreatePoolRequest {
	s.Members = append(s.Members, members...)
	return s
}

func (s *CreatePoolRequest) WithLoadBalancerID(lbID string) ICreatePoolRequest {
	s.LoadBalancerID = lbID
	return s
}

func (s *CreatePoolRequest) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"algorithm":     s.Algorithm,
		"poolName":      s.PoolName,
		"poolProtocol":  s.PoolProtocol,
		"stickiness":    s.Stickiness,
		"tlsEncryption": s.TLSEncryption,
		"healthMonitor": s.HealthMonitor.ToMap(),
		"members": func() []map[string]interface{} {
			members := make([]map[string]interface{}, 0, len(s.Members))
			for _, member := range s.Members {
				members = append(members, member.ToMap())
			}
			return members
		}(),
	}
}

func (s *CreatePoolRequest) WithAlgorithm(algorithm PoolAlgorithm) ICreatePoolRequest {
	s.Algorithm = algorithm
	return s
}

func (s *UpdatePoolRequest) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"algorithm":     s.Algorithm,
		"stickiness":    s.Stickiness,
		"tlsEncryption": s.TLSEncryption,
		"healthMonitor": s.HealthMonitor.ToMap(),
	}
}

func (s *UpdatePoolRequest) ToRequestBody() interface{} {
	s.HealthMonitor = s.HealthMonitor.(*HealthMonitor).toRequestBody()
	return s
}

func (s *UpdatePoolRequest) WithAlgorithm(algorithm PoolAlgorithm) IUpdatePoolRequest {
	s.Algorithm = algorithm
	return s
}

func (s *UpdatePoolRequest) WithHealthMonitor(monitor IHealthMonitorRequest) IUpdatePoolRequest {
	s.HealthMonitor = monitor
	return s
}

func (s *UpdatePoolRequest) WithLoadBalancerID(lbID string) IUpdatePoolRequest {
	s.LoadBalancerID = lbID
	return s
}

func (s *UpdatePoolRequest) WithTLSEncryption(v *bool) IUpdatePoolRequest {
	s.TLSEncryption = v
	return s
}

func (s *UpdatePoolRequest) WithStickiness(v *bool) IUpdatePoolRequest {
	s.Stickiness = v
	return s
}

func (s *UpdatePoolRequest) AddUserAgent(agent ...string) IUpdatePoolRequest {
	s.UserAgent.AddUserAgent(agent...)
	return s
}

func (s *HealthMonitor) ToRequestBody() interface{} {
	return s
}

func (s *HealthMonitor) WithHealthyThreshold(ht int) IHealthMonitorRequest {
	if ht < 1 {
		ht = 3
	}

	s.HealthyThreshold = ht
	return s
}

func (s *HealthMonitor) WithUnhealthyThreshold(uht int) IHealthMonitorRequest {
	if uht < 1 {
		uht = 3
	}

	s.UnhealthyThreshold = uht
	return s
}

func (s *HealthMonitor) WithInterval(interval int) IHealthMonitorRequest {
	if interval < 1 {
		interval = 30
	}

	s.Interval = interval
	return s
}

func (s *HealthMonitor) WithTimeout(to int) IHealthMonitorRequest {
	if to < 1 {
		to = 5
	}

	s.Timeout = to
	return s
}

func (s *HealthMonitor) WithHealthCheckMethod(method *HealthCheckMethod) IHealthMonitorRequest {
	s.HealthCheckMethod = method
	return s
}

func (s *HealthMonitor) WithHTTPVersion(version *HealthCheckHTTPVersion) IHealthMonitorRequest {
	s.HTTPVersion = version
	return s
}

func (s *HealthMonitor) WithHealthCheckPath(path *string) IHealthMonitorRequest {
	s.HealthCheckPath = path
	return s
}

func (s *HealthMonitor) WithDomainName(domain *string) IHealthMonitorRequest {
	s.DomainName = domain
	return s
}

func (s *HealthMonitor) WithSuccessCode(code *string) IHealthMonitorRequest {
	s.SuccessCode = code
	return s
}

func (s *HealthMonitor) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"healthCheckProtocol": s.HealthCheckProtocol,
		"healthyThreshold":    s.HealthyThreshold,
		"unhealthyThreshold":  s.UnhealthyThreshold,
		"interval":            s.Interval,
		"timeout":             s.Timeout,
		"healthCheckMethod":   s.HealthCheckMethod,
		"httpVersion":         s.HTTPVersion,
		"healthCheckPath":     s.HealthCheckPath,
		"domainName":          s.DomainName,
		"successCode":         s.SuccessCode,
	}
}

func (s *Member) ToRequestBody() interface{} {
	return s
}

func (s *Member) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"backup":      s.Backup,
		"ipAddress":   s.IpAddress,
		"monitorPort": s.MonitorPort,
		"name":        s.Name,
		"port":        s.Port,
		"weight":      s.Weight,
	}
}

func (s *UpdatePoolMembersRequest) WithMembers(members ...IMemberRequest) IUpdatePoolMembersRequest {
	s.Members = append(s.Members, members...)
	return s
}

func (s *UpdatePoolMembersRequest) ToRequestBody() interface{} {
	return s
}
