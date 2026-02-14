package v1

import "github.com/dannyota/greennode-community-sdk/v2/greennode/client"

// ------------------------------------------------------------

func createGlobalPoolUrl(psc client.ServiceClient, popts ICreateGlobalPoolRequest) string {
	return psc.ServiceURL(
		"global-load-balancers",
		popts.GetLoadBalancerId(),
		"global-pools",
	)
}

func listGlobalPoolsUrl(psc client.ServiceClient, popts IListGlobalPoolsRequest) string {
	return psc.ServiceURL(
		"global-load-balancers",
		popts.GetLoadBalancerId(),
		"global-pools",
	)
}

func updateGlobalPoolUrl(psc client.ServiceClient, popts IUpdateGlobalPoolRequest) string {
	return psc.ServiceURL(
		"global-load-balancers",
		popts.GetLoadBalancerId(),
		"global-pools",
		popts.GetPoolId(),
	)
}

func deleteGlobalPoolUrl(psc client.ServiceClient, popts IDeleteGlobalPoolRequest) string {
	return psc.ServiceURL(
		"global-load-balancers",
		popts.GetLoadBalancerId(),
		"global-pools",
		popts.GetPoolId(),
	)
}

func listGlobalPoolMembersUrl(psc client.ServiceClient, popts IListGlobalPoolMembersRequest) string {
	return psc.ServiceURL(
		"global-load-balancers",
		popts.GetLoadBalancerId(),
		"global-pools",
		popts.GetPoolId(),
		"pool-members",
	)
}

func patchGlobalPoolMembersUrl(psc client.ServiceClient, popts IPatchGlobalPoolMembersRequest) string {
	return psc.ServiceURL(
		"global-load-balancers",
		popts.GetLoadBalancerId(),
		"global-pools",
		popts.GetPoolId(),
		"pool-members",
	)
}

func getGlobalPoolMemberUrl(psc client.ServiceClient, popts IGetGlobalPoolMemberRequest) string {
	return psc.ServiceURL(
		"global-load-balancers",
		popts.GetLoadBalancerId(),
		"global-pools",
		popts.GetPoolId(),
		"pool-members",
		popts.GetPoolMemberId(),
	)
}

func deleteGlobalPoolMemberUrl(psc client.ServiceClient, popts IDeleteGlobalPoolMemberRequest) string {
	return psc.ServiceURL(
		"global-load-balancers",
		popts.GetLoadBalancerId(),
		"global-pools",
		popts.GetPoolId(),
		"pool-members",
		popts.GetPoolMemberId(),
	)
}

func updateGlobalPoolMemberUrl(psc client.ServiceClient, popts IUpdateGlobalPoolMemberRequest) string {
	return psc.ServiceURL(
		"global-load-balancers",
		popts.GetLoadBalancerId(),
		"global-pools",
		popts.GetPoolId(),
		"pool-members",
		popts.GetPoolMemberId(),
	)
}

func listGlobalListenersUrl(psc client.ServiceClient, popts IListGlobalListenersRequest) string {
	return psc.ServiceURL(
		"global-load-balancers",
		popts.GetLoadBalancerId(),
		"global-listeners",
	)
}

func createGlobalListenerUrl(psc client.ServiceClient, popts ICreateGlobalListenerRequest) string {
	return psc.ServiceURL(
		"global-load-balancers",
		popts.GetLoadBalancerId(),
		"global-listeners",
	)
}

func updateGlobalListenerUrl(psc client.ServiceClient, popts IUpdateGlobalListenerRequest) string {
	return psc.ServiceURL(
		"global-load-balancers",
		popts.GetLoadBalancerId(),
		"global-listeners",
		popts.GetListenerId(),
	)
}

func deleteGlobalListenerUrl(psc client.ServiceClient, popts IDeleteGlobalListenerRequest) string {
	return psc.ServiceURL(
		"global-load-balancers",
		popts.GetLoadBalancerId(),
		"global-listeners",
		popts.GetListenerId(),
	)
}

func getGlobalListenerUrl(psc client.ServiceClient, popts IGetGlobalListenerRequest) string {
	return psc.ServiceURL(
		"global-load-balancers",
		popts.GetLoadBalancerId(),
		"global-listeners",
		popts.GetListenerId(),
	)
}

// --------------------------------------------------------

func listGlobalLoadBalancersUrl(psc client.ServiceClient, popts IListGlobalLoadBalancersRequest) string {
	query, err := popts.ToListQuery()
	if err != nil {
		query = popts.GetDefaultQuery()
	}

	return psc.ServiceURL("global-load-balancers") + query
}

func createGlobalLoadBalancerUrl(psc client.ServiceClient, _ ICreateGlobalLoadBalancerRequest) string {
	return psc.ServiceURL("global-load-balancers")
}

func deleteGlobalLoadBalancerUrl(psc client.ServiceClient, popts IDeleteGlobalLoadBalancerRequest) string {
	return psc.ServiceURL(
		"global-load-balancers",
		popts.GetLoadBalancerId(),
	)
}

func getGlobalLoadBalancerByIdUrl(psc client.ServiceClient, popts IGetGlobalLoadBalancerByIdRequest) string {
	return psc.ServiceURL(
		"global-load-balancers",
		popts.GetLoadBalancerId(),
	)
}

func listGlobalPackagesUrl(psc client.ServiceClient, _ IListGlobalPackagesRequest) string {
	return psc.ServiceURL("packages")
}

func listGlobalRegionsUrl(psc client.ServiceClient, _ IListGlobalRegionsRequest) string {
	return psc.ServiceURL("regions")
}

func getGlobalLoadBalancerUsageHistoriesUrl(psc client.ServiceClient, popts IGetGlobalLoadBalancerUsageHistoriesRequest) string {
	query, err := popts.ToListQuery()
	if err != nil {
		query = popts.GetDefaultQuery()
	}

	baseURL := psc.ServiceURL(
		"global-load-balancers",
		popts.GetLoadBalancerId(),
		"usage-histories",
	)

	if query != "" {
		return baseURL + "?" + query
	}
	return baseURL
}
