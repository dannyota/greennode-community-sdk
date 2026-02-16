package inter

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
	HealthMonitor *HealthMonitor `json:"healthMonitor"`
	Members       []*Member     `json:"members"`

	LoadBalancerID string `json:"-"`
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

// normalizeForAPI clears health monitor fields that are irrelevant for the
// configured protocol before the API call. This mutates the receiver.
func (r *CreatePoolRequest) normalizeForAPI() {
	r.HealthMonitor.normalizeForAPI()
}

// normalizeForAPI clears fields that don't apply to the current health check
// protocol, enforces sane defaults, and sets a default domain name for
// HTTP/1.1 when none is provided.
func (h *HealthMonitor) normalizeForAPI() {
	if h.HealthyThreshold < 1 {
		h.HealthyThreshold = 3
	}
	if h.UnhealthyThreshold < 1 {
		h.UnhealthyThreshold = 3
	}
	if h.Interval < 1 {
		h.Interval = 30
	}
	if h.Timeout < 1 {
		h.Timeout = 5
	}

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


