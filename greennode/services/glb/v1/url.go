package v1

import "github.com/dannyota/greennode-community-sdk/v2/greennode/client"

// ------------------------------------------------------------

func createGlobalPoolUrl(sc client.ServiceClient, opts ICreateGlobalPoolRequest) string {
	return sc.ServiceURL(
		"global-load-balancers",
		opts.GetLoadBalancerId(),
		"global-pools",
	)
}

func listGlobalPoolsUrl(sc client.ServiceClient, opts IListGlobalPoolsRequest) string {
	return sc.ServiceURL(
		"global-load-balancers",
		opts.GetLoadBalancerId(),
		"global-pools",
	)
}

func updateGlobalPoolUrl(sc client.ServiceClient, opts IUpdateGlobalPoolRequest) string {
	return sc.ServiceURL(
		"global-load-balancers",
		opts.GetLoadBalancerId(),
		"global-pools",
		opts.GetPoolId(),
	)
}

func deleteGlobalPoolUrl(sc client.ServiceClient, opts IDeleteGlobalPoolRequest) string {
	return sc.ServiceURL(
		"global-load-balancers",
		opts.GetLoadBalancerId(),
		"global-pools",
		opts.GetPoolId(),
	)
}

func listGlobalPoolMembersUrl(sc client.ServiceClient, opts IListGlobalPoolMembersRequest) string {
	return sc.ServiceURL(
		"global-load-balancers",
		opts.GetLoadBalancerId(),
		"global-pools",
		opts.GetPoolId(),
		"pool-members",
	)
}

func patchGlobalPoolMembersUrl(sc client.ServiceClient, opts IPatchGlobalPoolMembersRequest) string {
	return sc.ServiceURL(
		"global-load-balancers",
		opts.GetLoadBalancerId(),
		"global-pools",
		opts.GetPoolId(),
		"pool-members",
	)
}

func getGlobalPoolMemberUrl(sc client.ServiceClient, opts IGetGlobalPoolMemberRequest) string {
	return sc.ServiceURL(
		"global-load-balancers",
		opts.GetLoadBalancerId(),
		"global-pools",
		opts.GetPoolId(),
		"pool-members",
		opts.GetPoolMemberId(),
	)
}

func deleteGlobalPoolMemberUrl(sc client.ServiceClient, opts IDeleteGlobalPoolMemberRequest) string {
	return sc.ServiceURL(
		"global-load-balancers",
		opts.GetLoadBalancerId(),
		"global-pools",
		opts.GetPoolId(),
		"pool-members",
		opts.GetPoolMemberId(),
	)
}

func updateGlobalPoolMemberUrl(sc client.ServiceClient, opts IUpdateGlobalPoolMemberRequest) string {
	return sc.ServiceURL(
		"global-load-balancers",
		opts.GetLoadBalancerId(),
		"global-pools",
		opts.GetPoolId(),
		"pool-members",
		opts.GetPoolMemberId(),
	)
}

func listGlobalListenersUrl(sc client.ServiceClient, opts IListGlobalListenersRequest) string {
	return sc.ServiceURL(
		"global-load-balancers",
		opts.GetLoadBalancerId(),
		"global-listeners",
	)
}

func createGlobalListenerUrl(sc client.ServiceClient, opts ICreateGlobalListenerRequest) string {
	return sc.ServiceURL(
		"global-load-balancers",
		opts.GetLoadBalancerId(),
		"global-listeners",
	)
}

func updateGlobalListenerUrl(sc client.ServiceClient, opts IUpdateGlobalListenerRequest) string {
	return sc.ServiceURL(
		"global-load-balancers",
		opts.GetLoadBalancerId(),
		"global-listeners",
		opts.GetListenerId(),
	)
}

func deleteGlobalListenerUrl(sc client.ServiceClient, opts IDeleteGlobalListenerRequest) string {
	return sc.ServiceURL(
		"global-load-balancers",
		opts.GetLoadBalancerId(),
		"global-listeners",
		opts.GetListenerId(),
	)
}

func getGlobalListenerUrl(sc client.ServiceClient, opts IGetGlobalListenerRequest) string {
	return sc.ServiceURL(
		"global-load-balancers",
		opts.GetLoadBalancerId(),
		"global-listeners",
		opts.GetListenerId(),
	)
}

// --------------------------------------------------------

func listGlobalLoadBalancersUrl(sc client.ServiceClient, opts IListGlobalLoadBalancersRequest) string {
	query, err := opts.ToListQuery()
	if err != nil {
		query = opts.GetDefaultQuery()
	}

	return sc.ServiceURL("global-load-balancers") + query
}

func createGlobalLoadBalancerUrl(sc client.ServiceClient, _ ICreateGlobalLoadBalancerRequest) string {
	return sc.ServiceURL("global-load-balancers")
}

func deleteGlobalLoadBalancerUrl(sc client.ServiceClient, opts IDeleteGlobalLoadBalancerRequest) string {
	return sc.ServiceURL(
		"global-load-balancers",
		opts.GetLoadBalancerId(),
	)
}

func getGlobalLoadBalancerByIdUrl(sc client.ServiceClient, opts IGetGlobalLoadBalancerByIdRequest) string {
	return sc.ServiceURL(
		"global-load-balancers",
		opts.GetLoadBalancerId(),
	)
}

func listGlobalPackagesUrl(sc client.ServiceClient, _ IListGlobalPackagesRequest) string {
	return sc.ServiceURL("packages")
}

func listGlobalRegionsUrl(sc client.ServiceClient, _ IListGlobalRegionsRequest) string {
	return sc.ServiceURL("regions")
}

func getGlobalLoadBalancerUsageHistoriesUrl(sc client.ServiceClient, opts IGetGlobalLoadBalancerUsageHistoriesRequest) string {
	query, err := opts.ToListQuery()
	if err != nil {
		query = opts.GetDefaultQuery()
	}

	baseURL := sc.ServiceURL(
		"global-load-balancers",
		opts.GetLoadBalancerId(),
		"usage-histories",
	)

	if query != "" {
		return baseURL + "?" + query
	}
	return baseURL
}
