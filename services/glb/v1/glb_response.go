package v1

// CreateGlobalLoadBalancerResponse is a multi-object envelope returned
// when creating a global load balancer.
type CreateGlobalLoadBalancerResponse struct {
	GlobalLoadBalancer GlobalLoadBalancer `json:"globalLoadBalancer"`
	GlobalListener     GlobalListener     `json:"globalListener"`
	GlobalPool         GlobalPool         `json:"globalPool"`
}

func (r *CreateGlobalLoadBalancerResponse) ToEntityGlobalLoadBalancer() *GlobalLoadBalancer {
	return &r.GlobalLoadBalancer
}

// ListGlobalPackagesResponse wraps a bare JSON array of packages.
type ListGlobalPackagesResponse []GlobalPackage

func (r *ListGlobalPackagesResponse) ToEntityListGlobalPackages() *ListGlobalPackages {
	if r == nil {
		return &ListGlobalPackages{}
	}
	return &ListGlobalPackages{Items: []GlobalPackage(*r)}
}

// ListGlobalRegionsResponse wraps a bare JSON array of regions.
type ListGlobalRegionsResponse []GlobalRegion

func (r *ListGlobalRegionsResponse) ToEntityListGlobalRegions() *ListGlobalRegions {
	if r == nil {
		return &ListGlobalRegions{}
	}
	return &ListGlobalRegions{Items: []GlobalRegion(*r)}
}
