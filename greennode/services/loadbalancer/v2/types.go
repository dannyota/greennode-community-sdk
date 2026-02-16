package v2

import "encoding/json"

type LoadBalancer struct {
	UUID               string  `json:"uuid"`
	Name               string  `json:"name"`
	DisplayStatus      string  `json:"displayStatus"`
	Address            string  `json:"address"`
	PrivateSubnetID    string  `json:"privateSubnetId"`
	PrivateSubnetCidr  string  `json:"privateSubnetCidr"`
	Type               string  `json:"type"`
	DisplayType        string  `json:"displayType"`
	LoadBalancerSchema string  `json:"loadBalancerSchema"`
	PackageID          string  `json:"packageId"`
	Description        string  `json:"description"`
	Location           string  `json:"location"`
	CreatedAt          string  `json:"createdAt"`
	UpdatedAt          string  `json:"updatedAt"`
	ProgressStatus     string  `json:"progressStatus"`
	Status             string  `json:"status"`
	BackendSubnetID    string  `json:"backendSubnetId"`
	Internal           bool    `json:"internal"`
	AutoScalable       bool    `json:"autoScalable"`
	ZoneID             string  `json:"zoneId"`
	MinSize            int     `json:"minSize"`
	MaxSize            int     `json:"maxSize"`
	TotalNodes         int     `json:"totalNodes"`
	Nodes              []*Node `json:"nodes"`
}

type Node struct {
	Status   string `json:"status"`
	ZoneID   string `json:"zoneId"`
	ZoneName string `json:"zoneName"`
	SubnetID string `json:"subnetId"`
}

type ListLoadBalancers struct {
	Items     []*LoadBalancer
	Page      int
	PageSize  int
	TotalPage int
	TotalItem int
}

func (l *ListLoadBalancers) Len() int {
	return len(l.Items)
}

func (l *ListLoadBalancers) Empty() bool {
	return l.Len() < 1
}

func (l *ListLoadBalancers) Add(item *LoadBalancer) {
	l.Items = append(l.Items, item)
}

func (l *ListLoadBalancers) At(idx int) *LoadBalancer {
	if idx < 0 || idx >= l.Len() {
		return nil
	}

	return l.Items[idx]
}

type ListLoadBalancerPackages struct {
	Items []*LoadBalancerPackage
}

type LoadBalancerPackage struct {
	UUID             string `json:"uuid"`
	Name             string `json:"name"`
	Type             string `json:"type"`
	ConnectionNumber int    `json:"connectionNumber"`
	DataTransfer     int    `json:"dataTransfer"`
	Mode             string `json:"mode"`
	LbType           string `json:"lbType"`
	DisplayLbType    string `json:"displayLbType"`
}

type ListenerInsertHeader struct {
	HeaderName  string `json:"headerName"`
	HeaderValue string `json:"headerValue"`
}

type Listener struct {
	UUID                            string                 `json:"uuid"`
	Name                            string                 `json:"name"`
	Description                     string                 `json:"description"`
	Protocol                        string                 `json:"protocol"`
	ProtocolPort                    int                    `json:"protocolPort"`
	ConnectionLimit                 int                    `json:"connectionLimit"`
	DefaultPoolID                   string                 `json:"defaultPoolId"`
	DefaultPoolName                 string                 `json:"defaultPoolName"`
	TimeoutClient                   int                    `json:"timeoutClient"`
	TimeoutMember                   int                    `json:"timeoutMember"`
	TimeoutConnection               int                    `json:"timeoutConnection"`
	AllowedCidrs                    string                 `json:"allowedCidrs"`
	CertificateAuthorities          []string               `json:"certificateAuthorities"`
	DisplayStatus                   string                 `json:"displayStatus"`
	CreatedAt                       string                 `json:"createdAt"`
	UpdatedAt                       string                 `json:"updatedAt"`
	DefaultCertificateAuthority     *string                `json:"defaultCertificateAuthority"`
	ClientCertificateAuthentication *string                `json:"clientCertificateAuthentication"`
	ProgressStatus                  string                 `json:"progressStatus"`
	InsertHeaders                   []ListenerInsertHeader `json:"insertHeaders"`
}

type ListListeners struct {
	Items []*Listener
}

func (ll *ListListeners) Add(listeners ...*Listener) {
	ll.Items = append(ll.Items, listeners...)
}

func (ll *ListListeners) Len() int {
	return len(ll.Items)
}

func (ll *ListListeners) Empty() bool {
	return ll.Len() < 1
}

func (ll *ListListeners) At(index int) *Listener {
	if index < 0 || index >= len(ll.Items) {
		return nil
	}

	return ll.Items[index]
}

type Pool struct {
	UUID              string         `json:"uuid"`
	Name              string         `json:"name"`
	Protocol          string         `json:"protocol"`
	Description       string         `json:"description"`
	LoadBalanceMethod string         `json:"loadBalanceMethod"`
	Status            string         `json:"displayStatus"`
	Stickiness        bool           `json:"stickiness"`
	TLSEncryption     bool           `json:"tlsEncryption"`
	Members           *ListMembers   `json:"members"`
	HealthMonitor     *HealthMonitor `json:"healthMonitor"`
}

type Member struct {
	UUID           string `json:"uuid"`
	Address        string `json:"address"`
	ProtocolPort   int    `json:"protocolPort"`
	Weight         int    `json:"weight"`
	MonitorPort    int    `json:"monitorPort"`
	SubnetID       string `json:"subnetId"`
	Name           string `json:"name"`
	PoolID         string `json:"poolId"`
	TypeCreate     string `json:"typeCreate"`
	Backup         bool   `json:"backup"`
	DisplayStatus  string `json:"displayStatus"`
	CreatedAt      string `json:"createdAt"`
	UpdatedAt      string `json:"updateAt"`
	CreatedBy      string `json:"createdBy"`
	ProgressStatus string `json:"progressStatus"`
}

type HealthMonitor struct {
	Timeout             int     `json:"timeout"`
	CreatedAt           string  `json:"createdAt"`
	UpdatedAt           string  `json:"updatedAt"`
	DomainName          *string `json:"domainName"`
	HTTPVersion         *string `json:"httpVersion"`
	HealthCheckProtocol string  `json:"healthCheckProtocol"`
	Interval            int     `json:"interval"`
	HealthyThreshold    int     `json:"healthyThreshold"`
	UnhealthyThreshold  int     `json:"unhealthyThreshold"`
	HealthCheckMethod   *string `json:"healthCheckMethod"`
	HealthCheckPath     *string `json:"healthCheckPath"`
	SuccessCode         *string `json:"successCode"`
	ProgressStatus      string  `json:"progressStatus"`
	DisplayStatus       string  `json:"displayStatus"`
}

func (h HealthMonitor) String() string {
	out, err := json.Marshal(h)
	if err != nil {
		return "Error parsing to string"
	}
	return string(out)
}

type ListPools struct {
	Items []*Pool
}

type ListMembers struct {
	Items []*Member
}

func (l *ListPools) Add(pools ...*Pool) {
	l.Items = append(l.Items, pools...)
}

func (l *ListMembers) Add(members ...*Member) {
	l.Items = append(l.Items, members...)
}

func (l *ListPools) Len() int {
	return len(l.Items)
}

func (l *ListPools) Empty() bool {
	return l.Len() < 1
}

func (l *ListPools) At(index int) *Pool {
	if index < 0 || index >= l.Len() {
		return nil
	}

	return l.Items[index]
}

type Policy struct {
	UUID             string    `json:"uuid"`
	Name             string    `json:"name"`
	Description      string    `json:"description"`
	RedirectPoolID   string    `json:"redirectPoolId"`
	RedirectPoolName string    `json:"redirectPoolName"`
	Action           string    `json:"action"`
	RedirectURL      string    `json:"redirectUrl"`
	RedirectHTTPCode int       `json:"redirectHttpCode"`
	KeepQueryString  bool      `json:"keepQueryString"`
	Position         int       `json:"position"`
	L7Rules          []*L7Rule `json:"l7Rules"`
	DisplayStatus    string    `json:"displayStatus"`
	CreatedAt        string    `json:"createdAt"`
	UpdatedAt        string    `json:"updatedAt"`
	ProgressStatus   string    `json:"progressStatus"`
}

type L7Rule struct {
	UUID               string `json:"uuid"`
	CompareType        string `json:"compareType"`
	RuleValue          string `json:"ruleValue"`
	RuleType           string `json:"ruleType"`
	ProvisioningStatus string `json:"provisioningStatus"`
	OperatingStatus    string `json:"operatingStatus"`
}

type ListPolicies struct {
	Items []*Policy
}

type Certificate struct {
	UUID               string `json:"uuid"`
	Name               string `json:"name"`
	CertificateType    string `json:"certificateType"`
	ExpiredAt          string `json:"expiredAt"`
	ImportedAt         string `json:"importedAt"`
	NotAfter           int64  `json:"notAfter"`
	KeyAlgorithm       string `json:"keyAlgorithm"`
	Serial             string `json:"serial"`
	Subject            string `json:"subject"`
	DomainName         string `json:"domainName"`
	InUse              bool   `json:"inUse"`
	Issuer             string `json:"issuer"`
	SignatureAlgorithm string `json:"signatureAlgorithm"`
	NotBefore          int64  `json:"notBefore"`
}

type ListCertificates struct {
	Certificates []Certificate `json:"certificates"`
}
