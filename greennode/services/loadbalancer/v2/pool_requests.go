package v2


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
	return &CreatePoolRequest{
		PoolName:     name,
		Algorithm:    PoolAlgorithmRoundRobin,
		PoolProtocol: protocol,
		Members:      make([]*Member, 0),
	}
}

func NewUpdatePoolRequest(lbID, poolID string) *UpdatePoolRequest {
	return &UpdatePoolRequest{
		LoadBalancerID: lbID,
		PoolID:         poolID,
	}
}

func NewGetPoolHealthMonitorByIDRequest(lbID, poolID string) *GetPoolHealthMonitorByIDRequest {
	return &GetPoolHealthMonitorByIDRequest{
		LoadBalancerID: lbID,
		PoolID:         poolID,
	}
}

func NewListPoolsByLoadBalancerIDRequest(lbID string) *ListPoolsByLoadBalancerIDRequest {
	return &ListPoolsByLoadBalancerIDRequest{
		LoadBalancerID: lbID,
	}
}

func NewUpdatePoolMembersRequest(lbID, poolID string) *UpdatePoolMembersRequest {
	return &UpdatePoolMembersRequest{
		LoadBalancerID: lbID,
		PoolID:         poolID,
		Members:        make([]*Member, 0),
	}
}

func NewListPoolMembersRequest(lbID, poolID string) *ListPoolMembersRequest {
	return &ListPoolMembersRequest{
		LoadBalancerID: lbID,
		PoolID:         poolID,
	}
}

func NewDeletePoolByIDRequest(lbID, poolID string) *DeletePoolByIDRequest {
	return &DeletePoolByIDRequest{
		LoadBalancerID: lbID,
		PoolID:         poolID,
	}
}

func NewHealthMonitor(checkProtocol HealthCheckProtocol) *HealthMonitor {
	return &HealthMonitor{
		HealthCheckProtocol: checkProtocol,
		HealthyThreshold:    3,
		UnhealthyThreshold:  3,
		Interval:            30,
		Timeout:             5,
	}
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
	return &GetPoolByIDRequest{
		LoadBalancerID: lbID,
		PoolID:         poolID,
	}
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

	LoadBalancerID string
}

type UpdatePoolRequest struct {
	Algorithm     PoolAlgorithm  `json:"algorithm"`
	Stickiness    *bool          `json:"stickiness,omitempty"`    // only for l7, l4 doesn't have this field => nil
	TLSEncryption *bool          `json:"tlsEncryption,omitempty"` // only for l7, l4 doesn't have this field => nil
	HealthMonitor *HealthMonitor `json:"healthMonitor"`

	LoadBalancerID string
	PoolID         string
}

type GetPoolHealthMonitorByIDRequest struct {
	LoadBalancerID string
	PoolID         string
}

type ListPoolMembersRequest struct {
	LoadBalancerID string
	PoolID         string
}

type DeletePoolByIDRequest struct {
	LoadBalancerID string
	PoolID         string
}

type GetPoolByIDRequest struct {
	LoadBalancerID string
	PoolID         string
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
	LoadBalancerID string
}

type UpdatePoolMembersRequest struct {
	Members []*Member `json:"members"`

	LoadBalancerID string
	PoolID         string
}

// normalizeForAPI clears health monitor fields that are irrelevant for the
// configured protocol before the API call. This mutates the receiver.
func (r *CreatePoolRequest) normalizeForAPI() {
	r.HealthMonitor.normalizeForAPI()
}

// normalizeForAPI clears fields that don't apply to the current health check
// protocol (e.g. TCP doesn't use HTTP path/version) and sets a default domain
// name for HTTP/1.1 when none is provided.
func (h *HealthMonitor) normalizeForAPI() {
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

func (r *UpdatePoolRequest) normalizeForAPI() {
	r.HealthMonitor.normalizeForAPI()
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

