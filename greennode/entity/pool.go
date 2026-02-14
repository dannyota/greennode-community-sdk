package entity

import "encoding/json"

type Pool struct {
	UUID              string
	Name              string
	Protocol          string
	Description       string
	LoadBalanceMethod string
	Status            string
	Stickiness        bool
	TLSEncryption     bool
	Members           *ListMembers
	HealthMonitor     *HealthMonitor
}

type Member struct {
	UUID           string
	Address        string
	ProtocolPort   int
	Weight         int
	MonitorPort    int
	SubnetID       string
	Name           string
	PoolID         string
	TypeCreate     string
	Backup         bool
	DisplayStatus  string
	CreatedAt      string
	UpdatedAt      string
	CreatedBy      string
	ProgressStatus string
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

func (p Pool) GetID() string {
	return p.UUID
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
