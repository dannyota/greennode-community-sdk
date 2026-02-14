package v2

import "github.com/dannyota/greennode-community-sdk/v2/greennode/client"

func createLoadBalancerUrl(sc client.ServiceClient) string {
	return sc.ServiceURL(
		sc.GetProjectId(),
		"loadBalancers")
}

func resizeLoadBalancerUrl(sc client.ServiceClient, opts IResizeLoadBalancerRequest) string {
	return sc.ServiceURL(
		sc.GetProjectId(),
		"loadBalancers",
		opts.GetLoadBalancerId(),
		"resize")
}

func listLoadBalancerPackagesUrl(sc client.ServiceClient, opts IListLoadBalancerPackagesRequest) string {
	return sc.ServiceURL(
		sc.GetProjectId(),
		"loadBalancers", "packages") + "?zoneId=" + opts.GetZoneId()
}

func getLoadBalancerByIdUrl(sc client.ServiceClient, opts IGetLoadBalancerByIdRequest) string {
	return sc.ServiceURL(
		sc.GetProjectId(),
		"loadBalancers",
		opts.GetLoadBalancerId())
}

func listLoadBalancersUrl(sc client.ServiceClient, opts IListLoadBalancersRequest) string {
	query, err := opts.ToListQuery()
	if err != nil {
		query = opts.GetDefaultQuery()
	}

	return sc.ServiceURL(sc.GetProjectId(), "loadBalancers") + query
}

func getPoolHealthMonitorByIdUrl(sc client.ServiceClient, opts IGetPoolHealthMonitorByIdRequest) string {
	return sc.ServiceURL(
		sc.GetProjectId(),
		"loadBalancers",
		opts.GetLoadBalancerId(),
		"pools",
		opts.GetPoolId(),
		"healthMonitor")
}

func createPoolUrl(sc client.ServiceClient, opts ICreatePoolRequest) string {
	return sc.ServiceURL(
		sc.GetProjectId(),
		"loadBalancers",
		opts.GetLoadBalancerId(),
		"pools")
}

func updatePoolUrl(sc client.ServiceClient, opts IUpdatePoolRequest) string {
	return sc.ServiceURL(
		sc.GetProjectId(),
		"loadBalancers",
		opts.GetLoadBalancerId(),
		"pools",
		opts.GetPoolId())
}

func createListenerUrl(sc client.ServiceClient, opts ICreateListenerRequest) string {
	return sc.ServiceURL(
		sc.GetProjectId(),
		"loadBalancers",
		opts.GetLoadBalancerId(),
		"listeners")
}

func updateListenerUrl(sc client.ServiceClient, opts IUpdateListenerRequest) string {
	return sc.ServiceURL(
		sc.GetProjectId(),
		"loadBalancers",
		opts.GetLoadBalancerId(),
		"listeners",
		opts.GetListenerId())
}

func listListenersByLoadBalancerIdUrl(sc client.ServiceClient, opts IListListenersByLoadBalancerIdRequest) string {
	return sc.ServiceURL(
		sc.GetProjectId(),
		"loadBalancers",
		opts.GetLoadBalancerId(),
		"listeners")
}

func listPoolsByLoadBalancerIdUrl(sc client.ServiceClient, opts IListPoolsByLoadBalancerIdRequest) string {
	return sc.ServiceURL(
		sc.GetProjectId(),
		"loadBalancers",
		opts.GetLoadBalancerId(),
		"pools")
}

func updatePoolMembersUrl(sc client.ServiceClient, opts IUpdatePoolMembersRequest) string {
	return sc.ServiceURL(
		sc.GetProjectId(),
		"loadBalancers",
		opts.GetLoadBalancerId(),
		"pools",
		opts.GetPoolId(),
		"members")
}

func listPoolMembersUrl(sc client.ServiceClient, opts IListPoolMembersRequest) string {
	return sc.ServiceURL(
		sc.GetProjectId(),
		"loadBalancers",
		opts.GetLoadBalancerId(),
		"pools",
		opts.GetPoolId(),
		"members")
}

func deletePoolByIdUrl(sc client.ServiceClient, opts IDeletePoolByIdRequest) string {
	return sc.ServiceURL(
		sc.GetProjectId(),
		"loadBalancers",
		opts.GetLoadBalancerId(),
		"pools",
		opts.GetPoolId())
}

func deleteListenerByIdUrl(sc client.ServiceClient, opts IDeleteListenerByIdRequest) string {
	return sc.ServiceURL(
		sc.GetProjectId(),
		"loadBalancers",
		opts.GetLoadBalancerId(),
		"listeners",
		opts.GetListenerId())
}

func deleteLoadBalancerByIdUrl(sc client.ServiceClient, opts IDeleteLoadBalancerByIdRequest) string {
	return sc.ServiceURL(
		sc.GetProjectId(),
		"loadBalancers",
		opts.GetLoadBalancerId())
}

func listTagsUrl(sc client.ServiceClient, opts IListTagsRequest) string {
	return sc.ServiceURL(
		sc.GetProjectId(),
		"tag",
		"resource",
		opts.GetLoadBalancerId())
}

func createTagsUrl(sc client.ServiceClient, opts ICreateTagsRequest) string {
	return sc.ServiceURL(
		sc.GetProjectId(),
		"tag",
		"resource",
		opts.GetLoadBalancerId())
}

func updateTagsUrl(sc client.ServiceClient, opts IUpdateTagsRequest) string {
	return sc.ServiceURL(
		sc.GetProjectId(),
		"tag",
		"resource",
		opts.GetLoadBalancerId())
}

// Policy

func listPoliciesUrl(sc client.ServiceClient, opts IListPoliciesRequest) string {
	return sc.ServiceURL(
		sc.GetProjectId(),
		"loadBalancers",
		opts.GetLoadBalancerId(),
		"listeners",
		opts.GetListenerId(),
		"l7policies",
	)
}

func createPolicyUrl(sc client.ServiceClient, opts ICreatePolicyRequest) string {
	return sc.ServiceURL(
		sc.GetProjectId(),
		"loadBalancers",
		opts.GetLoadBalancerId(),
		"listeners",
		opts.GetListenerId(),
		"l7policies",
	)
}

func getPolicyByIdUrl(sc client.ServiceClient, opts IGetPolicyByIdRequest) string {
	return sc.ServiceURL(
		sc.GetProjectId(),
		"loadBalancers",
		opts.GetLoadBalancerId(),
		"listeners",
		opts.GetListenerId(),
		"l7policies",
		opts.GetPolicyId(),
	)
}

func updatePolicyUrl(sc client.ServiceClient, opts IUpdatePolicyRequest) string {
	return sc.ServiceURL(
		sc.GetProjectId(),
		"loadBalancers",
		opts.GetLoadBalancerId(),
		"listeners",
		opts.GetListenerId(),
		"l7policies",
		opts.GetPolicyId(),
	)
}

func deletePolicyByIdUrl(sc client.ServiceClient, opts IDeletePolicyByIdRequest) string {
	return sc.ServiceURL(
		sc.GetProjectId(),
		"loadBalancers",
		opts.GetLoadBalancerId(),
		"listeners",
		opts.GetListenerId(),
		"l7policies",
		opts.GetPolicyId(),
	)
}

func reorderPoliciesUrl(sc client.ServiceClient, opts IReorderPoliciesRequest) string {
	return sc.ServiceURL(
		sc.GetProjectId(),
		"loadBalancers",
		opts.GetLoadBalancerId(),
		"listeners",
		opts.GetListenerId(),
		"reorderL7Policies")
}

func getPoolByIdUrl(sc client.ServiceClient, opts IGetPoolByIdRequest) string {
	return sc.ServiceURL(
		sc.GetProjectId(),
		"loadBalancers",
		opts.GetLoadBalancerId(),
		"pools",
		opts.GetPoolId())
}

func getListenerByIdUrl(sc client.ServiceClient, opts IGetListenerByIdRequest) string {
	return sc.ServiceURL(
		sc.GetProjectId(),
		"loadBalancers",
		opts.GetLoadBalancerId(),
		"listeners",
		opts.GetListenerId())
}

func resizeLoadBalancerByIdUrl(sc client.ServiceClient, opts IResizeLoadBalancerByIdRequest) string {
	return sc.ServiceURL(
		sc.GetProjectId(),
		"loadBalancers",
		opts.GetLoadBalancerId(),
		"resize")
}

func scaleLoadBalancerUrl(sc client.ServiceClient, opts IScaleLoadBalancerRequest) string {
	return sc.ServiceURL(
		sc.GetProjectId(),
		"loadBalancers",
		opts.GetLoadBalancerId(),
		"rebalancing")
}

// --------------------------------------------------------

func listCertificatesUrl(sc client.ServiceClient) string {
	return sc.ServiceURL(
		sc.GetProjectId(),
		"cas")
}

func createCertificateUrl(sc client.ServiceClient) string {
	return sc.ServiceURL(
		sc.GetProjectId(),
		"cas")
}

func getCertificateByIdUrl(sc client.ServiceClient, opts IGetCertificateByIdRequest) string {
	return sc.ServiceURL(
		sc.GetProjectId(),
		"cas",
		opts.GetCertificateId())
}

func deleteCertificateByIdUrl(sc client.ServiceClient, opts IDeleteCertificateByIdRequest) string {
	return sc.ServiceURL(
		sc.GetProjectId(),
		"cas",
		opts.GetCertificateId())
}
