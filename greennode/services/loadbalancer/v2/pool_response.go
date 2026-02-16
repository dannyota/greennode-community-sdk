package v2

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

func (r *GetPoolHealthMonitorByIDResponse) ToEntityHealthMonitor() *HealthMonitor {
	if r == nil {
		return nil
	}

	return &HealthMonitor{
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
	Data []poolResp `json:"data"`
}

type ListPoolMembersResponse struct {
	Data []poolMemberResp `json:"data"`
}

type GetPoolByIDResponse struct {
	Data poolResp `json:"data"`
}

type poolResp struct {
	UUID              string           `json:"uuid"`
	Name              string           `json:"name"`
	Protocol          string           `json:"protocol"`
	Description       string           `json:"description,omitempty"`
	LoadBalanceMethod string           `json:"loadBalanceMethod"`
	DisplayStatus     string           `json:"displayStatus"`
	CreatedAt         string           `json:"createdAt"`
	UpdatedAt         string           `json:"updatedAt"`
	Stickiness        bool             `json:"stickiness"`
	TLSEncryption     bool             `json:"tlsEncryption"`
	ProgressStatus    string           `json:"progressStatus"`
	Members           []poolMemberResp `json:"members"`
}

type poolMemberResp struct {
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

func (r *CreatePoolResponse) ToEntityPool() *Pool {
	return &Pool{
		UUID: r.UUID,
	}
}

func (r *ListPoolsByLoadBalancerIDResponse) ToEntityListPools() *ListPools {
	listPools := new(ListPools)
	for _, pool := range r.Data {
		listPools.Add(pool.toEntityPool())
	}

	return listPools
}

func (p *poolMemberResp) toEntityMember() *Member {
	return &Member{
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

func (p *poolResp) toEntityListMembers() *ListMembers {
	listMembers := &ListMembers{}
	for _, member := range p.Members {
		listMembers.Add(member.toEntityMember())
	}
	return listMembers
}

func (p *poolResp) toEntityPool() *Pool {
	return &Pool{
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

func (r *ListPoolMembersResponse) ToEntityListMembers() *ListMembers {
	listMembers := &ListMembers{}
	for _, member := range r.Data {
		listMembers.Add(member.toEntityMember())
	}
	return listMembers
}

func (r *GetPoolByIDResponse) ToEntityPool() *Pool {
	return r.Data.toEntityPool()
}
