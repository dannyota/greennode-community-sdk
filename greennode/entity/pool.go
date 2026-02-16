package entity

import "encoding/json"

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
	// parse to string and return
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
