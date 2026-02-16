package v1

import "github.com/dannyota/greennode-community-sdk/v2/greennode/entity"

// ListGlobalPoolsResponse wraps a bare JSON array of pools into the entity
// list type.
type ListGlobalPoolsResponse []*entity.GlobalPool

func (r *ListGlobalPoolsResponse) ToEntityListGlobalPools() *entity.ListGlobalPools {
	result := &entity.ListGlobalPools{}
	if r != nil {
		result.Items = []*entity.GlobalPool(*r)
	}
	return result
}

// UpdateGlobalPoolResponse returns only the pool ID.
type UpdateGlobalPoolResponse struct {
	ID string `json:"id"`
}

func (r *UpdateGlobalPoolResponse) ToEntityPool() *entity.GlobalPool {
	return &entity.GlobalPool{
		ID: r.ID,
	}
}

// GlobalPoolMemberResponse handles pool member deserialization where the
// nested Members field requires conversion from a JSON array to the
// entity.ListGlobalMembers wrapper type.
type GlobalPoolMemberResponse struct {
	CreatedAt            string                           `json:"createdAt"`
	UpdatedAt            string                           `json:"updatedAt"`
	DeletedAt            *string                          `json:"deletedAt"`
	ID                   string                           `json:"id"`
	Name                 string                           `json:"name"`
	Description          string                           `json:"description"`
	Region               string                           `json:"region"`
	GlobalPoolID         string                           `json:"globalPoolId"`
	GlobalLoadBalancerID string                           `json:"globalLoadBalancerId"`
	TrafficDial          int                              `json:"trafficDial"`
	VpcID                string                           `json:"vpcId"`
	Type                 string                           `json:"type"`
	Status               string                           `json:"status"`
	Members              []*entity.GlobalPoolMemberDetail `json:"members"`
}

func (r *GlobalPoolMemberResponse) ToEntityGlobalPoolMember() *entity.GlobalPoolMember {
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
		Members:              &entity.ListGlobalMembers{Items: r.Members},
	}
}

// ListGlobalPoolMembersResponse wraps a bare JSON array of pool members.
type ListGlobalPoolMembersResponse []*GlobalPoolMemberResponse

func (r *ListGlobalPoolMembersResponse) ToEntityListGlobalPoolMembers() *entity.ListGlobalPoolMembers {
	result := &entity.ListGlobalPoolMembers{}
	if r == nil || len(*r) < 1 {
		return result
	}
	for _, member := range *r {
		result.Items = append(result.Items, member.ToEntityGlobalPoolMember())
	}
	return result
}

// GetGlobalPoolMemberResponse is an alias for GlobalPoolMemberResponse.
type GetGlobalPoolMemberResponse GlobalPoolMemberResponse

func (r *GetGlobalPoolMemberResponse) ToEntityGlobalPoolMember() *entity.GlobalPoolMember {
	return (*GlobalPoolMemberResponse)(r).ToEntityGlobalPoolMember()
}

// UpdateGlobalPoolMemberResponse returns partial pool member fields.
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
