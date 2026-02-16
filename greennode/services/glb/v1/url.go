package v1

import "github.com/dannyota/greennode-community-sdk/v2/greennode/client"

func createGlobalPoolURL(sc *client.ServiceClient, opts *CreateGlobalPoolRequest) string {
	return sc.ServiceURL(
		"global-load-balancers",
		opts.LoadBalancerID,
		"global-pools",
	)
}

func listGlobalPoolsURL(sc *client.ServiceClient, opts *ListGlobalPoolsRequest) string {
	return sc.ServiceURL(
		"global-load-balancers",
		opts.LoadBalancerID,
		"global-pools",
	)
}

func updateGlobalPoolURL(sc *client.ServiceClient, opts *UpdateGlobalPoolRequest) string {
	return sc.ServiceURL(
		"global-load-balancers",
		opts.LoadBalancerID,
		"global-pools",
		opts.PoolID,
	)
}

func deleteGlobalPoolURL(sc *client.ServiceClient, opts *DeleteGlobalPoolRequest) string {
	return sc.ServiceURL(
		"global-load-balancers",
		opts.LoadBalancerID,
		"global-pools",
		opts.PoolID,
	)
}

func listGlobalPoolMembersURL(sc *client.ServiceClient, opts *ListGlobalPoolMembersRequest) string {
	return sc.ServiceURL(
		"global-load-balancers",
		opts.LoadBalancerID,
		"global-pools",
		opts.PoolID,
		"pool-members",
	)
}

func patchGlobalPoolMembersURL(sc *client.ServiceClient, opts *PatchGlobalPoolMembersRequest) string {
	return sc.ServiceURL(
		"global-load-balancers",
		opts.LoadBalancerID,
		"global-pools",
		opts.PoolID,
		"pool-members",
	)
}

func getGlobalPoolMemberURL(sc *client.ServiceClient, opts *GetGlobalPoolMemberRequest) string {
	return sc.ServiceURL(
		"global-load-balancers",
		opts.LoadBalancerID,
		"global-pools",
		opts.PoolID,
		"pool-members",
		opts.PoolMemberID,
	)
}

func deleteGlobalPoolMemberURL(sc *client.ServiceClient, opts *DeleteGlobalPoolMemberRequest) string {
	return sc.ServiceURL(
		"global-load-balancers",
		opts.LoadBalancerID,
		"global-pools",
		opts.PoolID,
		"pool-members",
		opts.PoolMemberID,
	)
}

func updateGlobalPoolMemberURL(sc *client.ServiceClient, opts *UpdateGlobalPoolMemberRequest) string {
	return sc.ServiceURL(
		"global-load-balancers",
		opts.LoadBalancerID,
		"global-pools",
		opts.PoolID,
		"pool-members",
		opts.PoolMemberID,
	)
}

func listGlobalListenersURL(sc *client.ServiceClient, opts *ListGlobalListenersRequest) string {
	return sc.ServiceURL(
		"global-load-balancers",
		opts.LoadBalancerID,
		"global-listeners",
	)
}

func createGlobalListenerURL(sc *client.ServiceClient, opts *CreateGlobalListenerRequest) string {
	return sc.ServiceURL(
		"global-load-balancers",
		opts.LoadBalancerID,
		"global-listeners",
	)
}

func updateGlobalListenerURL(sc *client.ServiceClient, opts *UpdateGlobalListenerRequest) string {
	return sc.ServiceURL(
		"global-load-balancers",
		opts.LoadBalancerID,
		"global-listeners",
		opts.ListenerID,
	)
}

func deleteGlobalListenerURL(sc *client.ServiceClient, opts *DeleteGlobalListenerRequest) string {
	return sc.ServiceURL(
		"global-load-balancers",
		opts.LoadBalancerID,
		"global-listeners",
		opts.ListenerID,
	)
}

func getGlobalListenerURL(sc *client.ServiceClient, opts *GetGlobalListenerRequest) string {
	return sc.ServiceURL(
		"global-load-balancers",
		opts.LoadBalancerID,
		"global-listeners",
		opts.ListenerID,
	)
}

func listGlobalLoadBalancersURL(sc *client.ServiceClient, opts *ListGlobalLoadBalancersRequest) string {
	query, err := opts.ToListQuery()
	if err != nil {
		query = opts.GetDefaultQuery()
	}

	return sc.ServiceURL("global-load-balancers") + query
}

func createGlobalLoadBalancerURL(sc *client.ServiceClient, _ *CreateGlobalLoadBalancerRequest) string {
	return sc.ServiceURL("global-load-balancers")
}

func deleteGlobalLoadBalancerURL(sc *client.ServiceClient, opts *DeleteGlobalLoadBalancerRequest) string {
	return sc.ServiceURL(
		"global-load-balancers",
		opts.LoadBalancerID,
	)
}

func getGlobalLoadBalancerByIDURL(sc *client.ServiceClient, opts *GetGlobalLoadBalancerByIDRequest) string {
	return sc.ServiceURL(
		"global-load-balancers",
		opts.LoadBalancerID,
	)
}

func listGlobalPackagesURL(sc *client.ServiceClient, _ *ListGlobalPackagesRequest) string {
	return sc.ServiceURL("packages")
}

func listGlobalRegionsURL(sc *client.ServiceClient, _ *ListGlobalRegionsRequest) string {
	return sc.ServiceURL("regions")
}

func getGlobalLoadBalancerUsageHistoriesURL(sc *client.ServiceClient, opts *GetGlobalLoadBalancerUsageHistoriesRequest) string {
	query, err := opts.ToListQuery()
	if err != nil {
		query = opts.GetDefaultQuery()
	}

	baseURL := sc.ServiceURL(
		"global-load-balancers",
		opts.LoadBalancerID,
		"usage-histories",
	)

	if query != "" {
		return baseURL + "?" + query
	}
	return baseURL
}
