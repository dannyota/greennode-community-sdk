package v1

import "github.com/dannyota/greennode-community-sdk/v2/greennode/client"

// ------------------------------------------------------------

func createGlobalPoolURL(sc client.ServiceClient, opts ICreateGlobalPoolRequest) string {
	return sc.ServiceURL(
		"global-load-balancers",
		opts.GetLoadBalancerID(),
		"global-pools",
	)
}

func listGlobalPoolsURL(sc client.ServiceClient, opts IListGlobalPoolsRequest) string {
	return sc.ServiceURL(
		"global-load-balancers",
		opts.GetLoadBalancerID(),
		"global-pools",
	)
}

func updateGlobalPoolURL(sc client.ServiceClient, opts IUpdateGlobalPoolRequest) string {
	return sc.ServiceURL(
		"global-load-balancers",
		opts.GetLoadBalancerID(),
		"global-pools",
		opts.GetPoolID(),
	)
}

func deleteGlobalPoolURL(sc client.ServiceClient, opts IDeleteGlobalPoolRequest) string {
	return sc.ServiceURL(
		"global-load-balancers",
		opts.GetLoadBalancerID(),
		"global-pools",
		opts.GetPoolID(),
	)
}

func listGlobalPoolMembersURL(sc client.ServiceClient, opts IListGlobalPoolMembersRequest) string {
	return sc.ServiceURL(
		"global-load-balancers",
		opts.GetLoadBalancerID(),
		"global-pools",
		opts.GetPoolID(),
		"pool-members",
	)
}

func patchGlobalPoolMembersURL(sc client.ServiceClient, opts IPatchGlobalPoolMembersRequest) string {
	return sc.ServiceURL(
		"global-load-balancers",
		opts.GetLoadBalancerID(),
		"global-pools",
		opts.GetPoolID(),
		"pool-members",
	)
}

func getGlobalPoolMemberURL(sc client.ServiceClient, opts IGetGlobalPoolMemberRequest) string {
	return sc.ServiceURL(
		"global-load-balancers",
		opts.GetLoadBalancerID(),
		"global-pools",
		opts.GetPoolID(),
		"pool-members",
		opts.GetPoolMemberID(),
	)
}

func deleteGlobalPoolMemberURL(sc client.ServiceClient, opts IDeleteGlobalPoolMemberRequest) string {
	return sc.ServiceURL(
		"global-load-balancers",
		opts.GetLoadBalancerID(),
		"global-pools",
		opts.GetPoolID(),
		"pool-members",
		opts.GetPoolMemberID(),
	)
}

func updateGlobalPoolMemberURL(sc client.ServiceClient, opts IUpdateGlobalPoolMemberRequest) string {
	return sc.ServiceURL(
		"global-load-balancers",
		opts.GetLoadBalancerID(),
		"global-pools",
		opts.GetPoolID(),
		"pool-members",
		opts.GetPoolMemberID(),
	)
}

func listGlobalListenersURL(sc client.ServiceClient, opts IListGlobalListenersRequest) string {
	return sc.ServiceURL(
		"global-load-balancers",
		opts.GetLoadBalancerID(),
		"global-listeners",
	)
}

func createGlobalListenerURL(sc client.ServiceClient, opts ICreateGlobalListenerRequest) string {
	return sc.ServiceURL(
		"global-load-balancers",
		opts.GetLoadBalancerID(),
		"global-listeners",
	)
}

func updateGlobalListenerURL(sc client.ServiceClient, opts IUpdateGlobalListenerRequest) string {
	return sc.ServiceURL(
		"global-load-balancers",
		opts.GetLoadBalancerID(),
		"global-listeners",
		opts.GetListenerID(),
	)
}

func deleteGlobalListenerURL(sc client.ServiceClient, opts IDeleteGlobalListenerRequest) string {
	return sc.ServiceURL(
		"global-load-balancers",
		opts.GetLoadBalancerID(),
		"global-listeners",
		opts.GetListenerID(),
	)
}

func getGlobalListenerURL(sc client.ServiceClient, opts IGetGlobalListenerRequest) string {
	return sc.ServiceURL(
		"global-load-balancers",
		opts.GetLoadBalancerID(),
		"global-listeners",
		opts.GetListenerID(),
	)
}

// --------------------------------------------------------

func listGlobalLoadBalancersURL(sc client.ServiceClient, opts IListGlobalLoadBalancersRequest) string {
	query, err := opts.ToListQuery()
	if err != nil {
		query = opts.GetDefaultQuery()
	}

	return sc.ServiceURL("global-load-balancers") + query
}

func createGlobalLoadBalancerURL(sc client.ServiceClient, _ ICreateGlobalLoadBalancerRequest) string {
	return sc.ServiceURL("global-load-balancers")
}

func deleteGlobalLoadBalancerURL(sc client.ServiceClient, opts IDeleteGlobalLoadBalancerRequest) string {
	return sc.ServiceURL(
		"global-load-balancers",
		opts.GetLoadBalancerID(),
	)
}

func getGlobalLoadBalancerByIDURL(sc client.ServiceClient, opts IGetGlobalLoadBalancerByIDRequest) string {
	return sc.ServiceURL(
		"global-load-balancers",
		opts.GetLoadBalancerID(),
	)
}

func listGlobalPackagesURL(sc client.ServiceClient, _ IListGlobalPackagesRequest) string {
	return sc.ServiceURL("packages")
}

func listGlobalRegionsURL(sc client.ServiceClient, _ IListGlobalRegionsRequest) string {
	return sc.ServiceURL("regions")
}

func getGlobalLoadBalancerUsageHistoriesURL(sc client.ServiceClient, opts IGetGlobalLoadBalancerUsageHistoriesRequest) string {
	query, err := opts.ToListQuery()
	if err != nil {
		query = opts.GetDefaultQuery()
	}

	baseURL := sc.ServiceURL(
		"global-load-balancers",
		opts.GetLoadBalancerID(),
		"usage-histories",
	)

	if query != "" {
		return baseURL + "?" + query
	}
	return baseURL
}
