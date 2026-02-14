package v2

import "github.com/dannyota/greennode-community-sdk/v2/greennode/entity"

type GetPoolHealthMonitorByIDResponse struct {
	Data struct {
		UUID                string  `json:"uuid"`
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
	} `json:"data"`
}

func (r *GetPoolHealthMonitorByIDResponse) ToEntityHealthMonitor() *entity.HealthMonitor {
	if r == nil {
		return nil
	}

	return &entity.HealthMonitor{
		Timeout:             r.Data.Timeout,
		CreatedAt:           r.Data.CreatedAt,
		UpdatedAt:           r.Data.UpdatedAt,
		DomainName:          r.Data.DomainName,
		HTTPVersion:         r.Data.HTTPVersion,
		HealthCheckProtocol: r.Data.HealthCheckProtocol,
		Interval:            r.Data.Interval,
		HealthyThreshold:    r.Data.HealthyThreshold,
		UnhealthyThreshold:  r.Data.UnhealthyThreshold,
		HealthCheckMethod:   r.Data.HealthCheckMethod,
		HealthCheckPath:     r.Data.HealthCheckPath,
		SuccessCode:         r.Data.SuccessCode,
		ProgressStatus:      r.Data.ProgressStatus,
		DisplayStatus:       r.Data.DisplayStatus,
	}
}

type CreatePoolResponse struct {
	UUID string `json:"uuid"`
}

type ListPoolsByLoadBalancerIDResponse struct {
	Data []Pool `json:"data"`
}

type ListPoolMembersResponse struct {
	Data []PoolMember `json:"data"`
}

type GetPoolByIDResponse struct {
	Data Pool `json:"data"`
}

type Pool struct {
	UUID              string       `json:"uuid"`
	Name              string       `json:"name"`
	Protocol          string       `json:"protocol"`
	Description       string       `json:"description,omitempty"`
	LoadBalanceMethod string       `json:"loadBalanceMethod"`
	DisplayStatus     string       `json:"displayStatus"`
	CreatedAt         string       `json:"createdAt"`
	UpdatedAt         string       `json:"updatedAt"`
	Stickiness        bool         `json:"stickiness"`
	TLSEncryption     bool         `json:"tlsEncryption"`
	ProgressStatus    string       `json:"progressStatus"`
	Members           []PoolMember `json:"members"`
}

type PoolMember struct {
	Address        string `json:"address"`
	Backup         bool   `json:"backup"`
	CreatedAt      string `json:"createdAt"`
	DisplayStatus  string `json:"displayStatus"`
	MonitorPort    int    `json:"monitorPort"`
	Name           string `json:"name"`
	PoolID         string `json:"poolId"`
	ProgressStatus string `json:"progressStatus"`
	ProtocolPort   int    `json:"protocolPort"`
	SubnetID       string `json:"subnetId"`
	TypeCreate     string `json:"typeCreate"`
	UpdateAt       string `json:"updateAt,omitempty"`
	UUID           string `json:"uuid"`
	Weight         int    `json:"weight"`
}

func (r *CreatePoolResponse) ToEntityPool() *entity.Pool {
	return &entity.Pool{
		UUID: r.UUID,
	}
}

func (r *ListPoolsByLoadBalancerIDResponse) ToEntityListPools() *entity.ListPools {
	listPools := new(entity.ListPools)
	for _, pool := range r.Data {
		listPools.Add(pool.toEntityPool())
	}

	return listPools
}

func (p *PoolMember) toEntityMember() *entity.Member {
	return &entity.Member{
		Address:        p.Address,
		Backup:         p.Backup,
		Name:           p.Name,
		UUID:           p.UUID,
		DisplayStatus:  p.DisplayStatus,
		ProtocolPort:   p.ProtocolPort,
		MonitorPort:    p.MonitorPort,
		SubnetID:       p.SubnetID,
		TypeCreate:     p.TypeCreate,
		CreatedAt:      p.CreatedAt,
		Weight:         p.Weight,
		PoolID:         p.PoolID,
		ProgressStatus: p.ProgressStatus,
	}
}

func (p *Pool) toEntityListMembers() *entity.ListMembers {
	listMembers := &entity.ListMembers{}
	for _, member := range p.Members {
		listMembers.Add(member.toEntityMember())
	}
	return listMembers
}

func (p *Pool) toEntityPool() *entity.Pool {
	return &entity.Pool{
		UUID:              p.UUID,
		Name:              p.Name,
		Protocol:          p.Protocol,
		Description:       p.Description,
		LoadBalanceMethod: p.LoadBalanceMethod,
		Status:            p.DisplayStatus,
		Stickiness:        p.Stickiness,
		TLSEncryption:     p.TLSEncryption,
		Members:           p.toEntityListMembers(),
	}
}

func (r *ListPoolMembersResponse) ToEntityListMembers() *entity.ListMembers {
	listMembers := &entity.ListMembers{}
	for _, member := range r.Data {
		listMembers.Add(member.toEntityMember())
	}
	return listMembers
}

func (r *GetPoolByIDResponse) ToEntityPool() *entity.Pool {
	return r.Data.toEntityPool()
}
