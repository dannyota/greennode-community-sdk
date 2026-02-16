package v1

import "github.com/dannyota/greennode-community-sdk/v2/greennode/entity"

// CreateGlobalLoadBalancerResponse is a multi-object envelope returned
// when creating a global load balancer.
type CreateGlobalLoadBalancerResponse struct {
	GlobalLoadBalancer entity.GlobalLoadBalancer `json:"globalLoadBalancer"`
	GlobalListener     entity.GlobalListener     `json:"globalListener"`
	GlobalPool         entity.GlobalPool         `json:"globalPool"`
}

func (r *CreateGlobalLoadBalancerResponse) ToEntityGlobalLoadBalancer() *entity.GlobalLoadBalancer {
	return &r.GlobalLoadBalancer
}

// ListGlobalPackagesResponse wraps a bare JSON array of packages.
type ListGlobalPackagesResponse []entity.GlobalPackage

func (r *ListGlobalPackagesResponse) ToEntityListGlobalPackages() *entity.ListGlobalPackages {
	if r == nil {
		return &entity.ListGlobalPackages{}
	}
	return &entity.ListGlobalPackages{Items: []entity.GlobalPackage(*r)}
}

// ListGlobalRegionsResponse wraps a bare JSON array of regions.
type ListGlobalRegionsResponse []entity.GlobalRegion

func (r *ListGlobalRegionsResponse) ToEntityListGlobalRegions() *entity.ListGlobalRegions {
	if r == nil {
		return &entity.ListGlobalRegions{}
	}
	return &entity.ListGlobalRegions{Items: []entity.GlobalRegion(*r)}
}
