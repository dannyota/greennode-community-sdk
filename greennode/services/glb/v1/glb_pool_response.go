package v1

import "github.com/dannyota/greennode-community-sdk/v2/greennode/entity"

type GlobalPoolResponse struct {
	CreatedAt                 string                      `json:"createdAt"`
	UpdatedAt                 string                      `json:"updatedAt"`
	DeletedAt                 *string                     `json:"deletedAt"`
	ID                        string                      `json:"id"`
	Name                      string                      `json:"name"`
	Description               string                      `json:"description"`
	GlobalLoadBalancerID      string                      `json:"globalLoadBalancerId"`
	Algorithm                 string                      `json:"algorithm"`
	StickySession             *string                     `json:"stickySession"`
	TLSEnabled                *string                     `json:"tlsEnabled"`
	Protocol                  string                      `json:"protocol"`
	Status                    string                      `json:"status"`
	Health                    *HealthResponse             `json:"health"`
	GlobalPoolMembersResponse *[]GlobalPoolMemberResponse `json:"globalPoolMembers"`
}

func (r *GlobalPoolResponse) ToEntityPool() *entity.GlobalPool {
	return &entity.GlobalPool{
		CreatedAt:            r.CreatedAt,
		UpdatedAt:            r.UpdatedAt,
		DeletedAt:            r.DeletedAt,
		ID:                   r.ID,
		Name:                 r.Name,
		Description:          r.Description,
		GlobalLoadBalancerID: r.GlobalLoadBalancerID,
		Algorithm:            r.Algorithm,
		StickySession:        r.StickySession,
		TLSEnabled:           r.TLSEnabled,
		Protocol:             r.Protocol,
		Status:               r.Status,
		Health:               r.Health.ToEntityGlobalPoolHealthMonitor(),
	}
}

type HealthResponse struct {
	CreatedAt            string  `json:"createdAt"`
	UpdatedAt            string  `json:"updatedAt"`
	DeletedAt            *string `json:"deletedAt"`
	ID                   string  `json:"id"`
	GlobalPoolID         string  `json:"globalPoolId"`
	GlobalLoadBalancerID string  `json:"globalLoadBalancerId"`
	Protocol             string  `json:"protocol"`
	Path                 *string `json:"path"`
	Timeout              int     `json:"timeout"`
	IntervalTime         int     `json:"intervalTime"`
	HealthyThreshold     int     `json:"healthyThreshold"`
	UnhealthyThreshold   int     `json:"unhealthyThreshold"`
	HTTPVersion          *string `json:"httpVersion"`
	HTTPMethod           *string `json:"httpMethod"`
	DomainName           *string `json:"domainName"`
	SuccessCode          *string `json:"successCode"`
	Status               string  `json:"status"`
}

func (r *HealthResponse) ToEntityGlobalPoolHealthMonitor() *entity.GlobalPoolHealthMonitor {
	return &entity.GlobalPoolHealthMonitor{
		CreatedAt:            r.CreatedAt,
		UpdatedAt:            r.UpdatedAt,
		DeletedAt:            r.DeletedAt,
		ID:                   r.ID,
		GlobalPoolID:         r.GlobalPoolID,
		GlobalLoadBalancerID: r.GlobalLoadBalancerID,
		Protocol:             r.Protocol,
		Path:                 r.Path,
		Timeout:              r.Timeout,
		IntervalTime:         r.IntervalTime,
		HealthyThreshold:     r.HealthyThreshold,
		UnhealthyThreshold:   r.UnhealthyThreshold,
		DomainName:           r.DomainName,
		HTTPVersion:          r.HTTPVersion,
		HTTPMethod:           r.HTTPMethod,
		SuccessCode:          r.SuccessCode,
		Status:               r.Status,
	}
}

type GlobalPoolMemberResponse struct {
	CreatedAt            string                  `json:"createdAt"`
	UpdatedAt            string                  `json:"updatedAt"`
	DeletedAt            *string                 `json:"deletedAt"`
	ID                   string                  `json:"id"`
	Name                 string                  `json:"name"`
	Description          string                  `json:"description"`
	Region               string                  `json:"region"`
	GlobalPoolID         string                  `json:"globalPoolId"`
	GlobalLoadBalancerID string                  `json:"globalLoadBalancerId"`
	TrafficDial          int                     `json:"trafficDial"`
	VpcID                string                  `json:"vpcId"`
	Type                 string                  `json:"type"`
	Status               string                  `json:"status"`
	Members              []*GlobalMemberResponse `json:"members"`
}

func (r *GlobalPoolMemberResponse) ToEntityGlobalPoolMember() *entity.GlobalPoolMember {
	members := make([]*entity.GlobalPoolMemberDetail, 0, len(r.Members))
	for _, member := range r.Members {
		members = append(members, member.ToEntityGlobalMember())
	}
	return &entity.GlobalPoolMember{
		CreatedAt:            r.CreatedAt,
		UpdatedAt:            r.UpdatedAt,
		DeletedAt:            r.DeletedAt,
		ID:                   r.ID,
		Name:                 r.Name,
		Description:          r.Description,
		Region:               r.Region,
		GlobalPoolID:         r.GlobalPoolID,
		GlobalLoadBalancerID: r.GlobalLoadBalancerID,
		TrafficDial:          r.TrafficDial,
		VpcID:                r.VpcID,
		Type:                 r.Type,
		Status:               r.Status,
		Members:              &entity.ListGlobalMembers{Items: members},
	}
}

type GlobalMemberResponse struct {
	CreatedAt            string  `json:"createdAt"`
	UpdatedAt            string  `json:"updatedAt"`
	DeletedAt            *string `json:"deletedAt"`
	ID                   string  `json:"id"`
	Name                 string  `json:"name"`
	Description          string  `json:"description"`
	GlobalPoolMemberID   string  `json:"globalPoolMemberId"`
	GlobalLoadBalancerID string  `json:"globalLoadBalancerId"`
	SubnetID             string  `json:"subnetId"`
	Address              string  `json:"address"`
	Weight               int     `json:"weight"`
	Port                 int     `json:"port"`
	MonitorPort          int     `json:"monitorPort"`
	BackupRole           bool    `json:"backupRole"`
	Status               string  `json:"status"`
}

func (r *GlobalMemberResponse) ToEntityGlobalMember() *entity.GlobalPoolMemberDetail {
	return &entity.GlobalPoolMemberDetail{
		CreatedAt:            r.CreatedAt,
		UpdatedAt:            r.UpdatedAt,
		DeletedAt:            r.DeletedAt,
		ID:                   r.ID,
		Name:                 r.Name,
		Description:          r.Description,
		GlobalPoolMemberID:   r.GlobalPoolMemberID,
		GlobalLoadBalancerID: r.GlobalLoadBalancerID,
		SubnetID:             r.SubnetID,
		Address:              r.Address,
		Weight:               r.Weight,
		Port:                 r.Port,
		MonitorPort:          r.MonitorPort,
		BackupRole:           r.BackupRole,
		Status:               r.Status,
	}
}

// ----------------------------------------------------------

type ListGlobalPoolsResponse []*GlobalPoolResponse

func (r *ListGlobalPoolsResponse) ToEntityListGlobalPools() *entity.ListGlobalPools {
	result := &entity.ListGlobalPools{
		Items: make([]*entity.GlobalPool, 0),
	}

	if r == nil || len(*r) < 1 {
		return result
	}

	for _, pool := range *r {
		result.Items = append(result.Items, pool.ToEntityPool())
	}

	return result
}

// ----------------------------------------------------------

type CreateGlobalPoolResponse struct {
	ID                   string                      `json:"id"`
	Name                 string                      `json:"name"`
	Description          string                      `json:"description"`
	GlobalLoadBalancerID string                      `json:"globalLoadBalancerId"`
	Algorithm            string                      `json:"algorithm"`
	StickySession        *string                     `json:"stickySession"`
	TLSEnabled           *string                     `json:"tlsEnabled"`
	Protocol             string                      `json:"protocol"`
	Health               *HealthResponse             `json:"health"`
	GlobalPoolMembers    []*GlobalPoolMemberResponse `json:"globalPoolMembers"`
}

func (r *CreateGlobalPoolResponse) ToEntityPool() *entity.GlobalPool {
	return &entity.GlobalPool{
		ID:                   r.ID,
		Name:                 r.Name,
		Description:          r.Description,
		GlobalLoadBalancerID: r.GlobalLoadBalancerID,
		Algorithm:            r.Algorithm,
		StickySession:        r.StickySession,
		TLSEnabled:           r.TLSEnabled,
		Protocol:             r.Protocol,
	}
}

// ----------------------------------------------------------

type UpdateGlobalPoolResponse struct {
	ID string `json:"id"`
}

func (r *UpdateGlobalPoolResponse) ToEntityPool() *entity.GlobalPool {
	return &entity.GlobalPool{
		ID: r.ID,
	}
}

// ----------------------------------------------------------

type ListGlobalPoolMembersResponse []*GlobalPoolMemberResponse

func (r *ListGlobalPoolMembersResponse) ToEntityListGlobalPoolMembers() *entity.ListGlobalPoolMembers {
	result := &entity.ListGlobalPoolMembers{
		Items: make([]*entity.GlobalPoolMember, 0),
	}

	if r == nil || len(*r) < 1 {
		return result
	}

	for _, member := range *r {
		result.Items = append(result.Items, member.ToEntityGlobalPoolMember())
	}

	return result
}

// ----------------------------------------------------------

type GetGlobalPoolMemberResponse GlobalPoolMemberResponse

func (r *GetGlobalPoolMemberResponse) ToEntityGlobalPoolMember() *entity.GlobalPoolMember {
	members := make([]*entity.GlobalPoolMemberDetail, 0, len(r.Members))
	for _, member := range r.Members {
		members = append(members, member.ToEntityGlobalMember())
	}
	return &entity.GlobalPoolMember{
		CreatedAt:            r.CreatedAt,
		UpdatedAt:            r.UpdatedAt,
		DeletedAt:            r.DeletedAt,
		ID:                   r.ID,
		Name:                 r.Name,
		Description:          r.Description,
		Region:               r.Region,
		GlobalPoolID:         r.GlobalPoolID,
		GlobalLoadBalancerID: r.GlobalLoadBalancerID,
		TrafficDial:          r.TrafficDial,
		VpcID:                r.VpcID,
		Type:                 r.Type,
		Status:               r.Status,
		Members:              &entity.ListGlobalMembers{Items: members},
	}
}

// ----------------------------------------------------------

type UpdateGlobalPoolMemberResponse struct {
	ID                   string `json:"id"`
	GlobalPoolID         string `json:"globalPoolId"`
	GlobalLoadBalancerID string `json:"globalLoadBalancerId"`
	TrafficDial          int    `json:"trafficDial"`
}

func (r *UpdateGlobalPoolMemberResponse) ToEntityGlobalPoolMember() *entity.GlobalPoolMember {
	return &entity.GlobalPoolMember{
		ID:                   r.ID,
		GlobalPoolID:         r.GlobalPoolID,
		GlobalLoadBalancerID: r.GlobalLoadBalancerID,
		TrafficDial:          r.TrafficDial,
	}
}

// ----------------------------------------------------------
