package v1

import "github.com/dannyota/greennode-community-sdk/v2/greennode/client"

// ------------------------------------------------------------

func createGlobalPoolUrl(psc client.IServiceClient, popts ICreateGlobalPoolRequest) string {
	return psc.ServiceURL(
		"global-load-balancers",
		popts.GetLoadBalancerId(),
		"global-pools",
	)
}

func listGlobalPoolsUrl(psc client.IServiceClient, popts IListGlobalPoolsRequest) string {
	return psc.ServiceURL(
		"global-load-balancers",
		popts.GetLoadBalancerId(),
		"global-pools",
	)
}

func updateGlobalPoolUrl(psc client.IServiceClient, popts IUpdateGlobalPoolRequest) string {
	return psc.ServiceURL(
		"global-load-balancers",
		popts.GetLoadBalancerId(),
		"global-pools",
		popts.GetPoolId(),
	)
}

func deleteGlobalPoolUrl(psc client.IServiceClient, popts IDeleteGlobalPoolRequest) string {
	return psc.ServiceURL(
		"global-load-balancers",
		popts.GetLoadBalancerId(),
		"global-pools",
		popts.GetPoolId(),
	)
}

func listGlobalPoolMembersUrl(psc client.IServiceClient, popts IListGlobalPoolMembersRequest) string {
	return psc.ServiceURL(
		"global-load-balancers",
		popts.GetLoadBalancerId(),
		"global-pools",
		popts.GetPoolId(),
		"pool-members",
	)
}

func patchGlobalPoolMembersUrl(psc client.IServiceClient, popts IPatchGlobalPoolMembersRequest) string {
	return psc.ServiceURL(
		"global-load-balancers",
		popts.GetLoadBalancerId(),
		"global-pools",
		popts.GetPoolId(),
		"pool-members",
	)
}

func getGlobalPoolMemberUrl(psc client.IServiceClient, popts IGetGlobalPoolMemberRequest) string {
	return psc.ServiceURL(
		"global-load-balancers",
		popts.GetLoadBalancerId(),
		"global-pools",
		popts.GetPoolId(),
		"pool-members",
		popts.GetPoolMemberId(),
	)
}

func deleteGlobalPoolMemberUrl(psc client.IServiceClient, popts IDeleteGlobalPoolMemberRequest) string {
	return psc.ServiceURL(
		"global-load-balancers",
		popts.GetLoadBalancerId(),
		"global-pools",
		popts.GetPoolId(),
		"pool-members",
		popts.GetPoolMemberId(),
	)
}

func updateGlobalPoolMemberUrl(psc client.IServiceClient, popts IUpdateGlobalPoolMemberRequest) string {
	return psc.ServiceURL(
		"global-load-balancers",
		popts.GetLoadBalancerId(),
		"global-pools",
		popts.GetPoolId(),
		"pool-members",
		popts.GetPoolMemberId(),
	)
}

func listGlobalListenersUrl(psc client.IServiceClient, popts IListGlobalListenersRequest) string {
	return psc.ServiceURL(
		"global-load-balancers",
		popts.GetLoadBalancerId(),
		"global-listeners",
	)
}

func createGlobalListenerUrl(psc client.IServiceClient, popts ICreateGlobalListenerRequest) string {
	return psc.ServiceURL(
		"global-load-balancers",
		popts.GetLoadBalancerId(),
		"global-listeners",
	)
}

func updateGlobalListenerUrl(psc client.IServiceClient, popts IUpdateGlobalListenerRequest) string {
	return psc.ServiceURL(
		"global-load-balancers",
		popts.GetLoadBalancerId(),
		"global-listeners",
		popts.GetListenerId(),
	)
}

func deleteGlobalListenerUrl(psc client.IServiceClient, popts IDeleteGlobalListenerRequest) string {
	return psc.ServiceURL(
		"global-load-balancers",
		popts.GetLoadBalancerId(),
		"global-listeners",
		popts.GetListenerId(),
	)
}

func getGlobalListenerUrl(psc client.IServiceClient, popts IGetGlobalListenerRequest) string {
	return psc.ServiceURL(
		"global-load-balancers",
		popts.GetLoadBalancerId(),
		"global-listeners",
		popts.GetListenerId(),
	)
}

// --------------------------------------------------------

func listGlobalLoadBalancersUrl(psc client.IServiceClient, popts IListGlobalLoadBalancersRequest) string {
	query, err := popts.ToListQuery()
	if err != nil {
		query = popts.GetDefaultQuery()
	}

	return psc.ServiceURL("global-load-balancers") + query
}

func createGlobalLoadBalancerUrl(psc client.IServiceClient, _ ICreateGlobalLoadBalancerRequest) string {
	return psc.ServiceURL("global-load-balancers")
}

func deleteGlobalLoadBalancerUrl(psc client.IServiceClient, popts IDeleteGlobalLoadBalancerRequest) string {
	return psc.ServiceURL(
		"global-load-balancers",
		popts.GetLoadBalancerId(),
	)
}

func getGlobalLoadBalancerByIdUrl(psc client.IServiceClient, popts IGetGlobalLoadBalancerByIdRequest) string {
	return psc.ServiceURL(
		"global-load-balancers",
		popts.GetLoadBalancerId(),
	)
}

func listGlobalPackagesUrl(psc client.IServiceClient, _ IListGlobalPackagesRequest) string {
	return psc.ServiceURL("packages")
}

func listGlobalRegionsUrl(psc client.IServiceClient, _ IListGlobalRegionsRequest) string {
	return psc.ServiceURL("regions")
}

func getGlobalLoadBalancerUsageHistoriesUrl(psc client.IServiceClient, popts IGetGlobalLoadBalancerUsageHistoriesRequest) string {
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
