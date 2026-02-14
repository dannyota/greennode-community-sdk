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

func (s *GetPoolHealthMonitorByIDResponse) ToEntityHealthMonitor() *entity.HealthMonitor {
	if s == nil {
		return nil
	}

	return &entity.HealthMonitor{
		Timeout:             s.Data.Timeout,
		CreatedAt:           s.Data.CreatedAt,
		UpdatedAt:           s.Data.UpdatedAt,
		DomainName:          s.Data.DomainName,
		HTTPVersion:         s.Data.HTTPVersion,
		HealthCheckProtocol: s.Data.HealthCheckProtocol,
		Interval:            s.Data.Interval,
		HealthyThreshold:    s.Data.HealthyThreshold,
		UnhealthyThreshold:  s.Data.UnhealthyThreshold,
		HealthCheckMethod:   s.Data.HealthCheckMethod,
		HealthCheckPath:     s.Data.HealthCheckPath,
		SuccessCode:         s.Data.SuccessCode,
		ProgressStatus:      s.Data.ProgressStatus,
		DisplayStatus:       s.Data.DisplayStatus,
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

func (s *CreatePoolResponse) ToEntityPool() *entity.Pool {
	return &entity.Pool{
		UUID: s.UUID,
	}
}

func (s *ListPoolsByLoadBalancerIDResponse) ToEntityListPools() *entity.ListPools {
	listPools := new(entity.ListPools)
	for _, pool := range s.Data {
		listPools.Add(pool.toEntityPool())
	}

	return listPools
}

func (s *PoolMember) toEntityMember() *entity.Member {
	return &entity.Member{
		Address:        s.Address,
		Backup:         s.Backup,
		Name:           s.Name,
		UUID:           s.UUID,
		DisplayStatus:  s.DisplayStatus,
		ProtocolPort:   s.ProtocolPort,
		MonitorPort:    s.MonitorPort,
		SubnetID:       s.SubnetID,
		TypeCreate:     s.TypeCreate,
		CreatedAt:      s.CreatedAt,
		Weight:         s.Weight,
		PoolID:         s.PoolID,
		ProgressStatus: s.ProgressStatus,
	}
}

func (s *Pool) toEntityListMembers() *entity.ListMembers {
	listMembers := &entity.ListMembers{}
	for _, member := range s.Members {
		listMembers.Add(member.toEntityMember())
	}
	return listMembers
}

func (s *Pool) toEntityPool() *entity.Pool {
	return &entity.Pool{
		UUID:              s.UUID,
		Name:              s.Name,
		Protocol:          s.Protocol,
		Description:       s.Description,
		LoadBalanceMethod: s.LoadBalanceMethod,
		Status:            s.DisplayStatus,
		Stickiness:        s.Stickiness,
		TLSEncryption:     s.TLSEncryption,
		Members:           s.toEntityListMembers(),
	}
}

func (s *ListPoolMembersResponse) ToEntityListMembers() *entity.ListMembers {
	listMembers := &entity.ListMembers{}
	for _, member := range s.Data {
		listMembers.Add(member.toEntityMember())
	}
	return listMembers
}

func (s *GetPoolByIDResponse) ToEntityPool() *entity.Pool {
	return s.Data.toEntityPool()
}
