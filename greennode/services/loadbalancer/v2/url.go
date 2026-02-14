package v2

import "github.com/dannyota/greennode-community-sdk/v2/greennode/client"

func createLoadBalancerUrl(psc client.ServiceClient) string {
	return psc.ServiceURL(
		psc.GetProjectId(),
		"loadBalancers")
}

func resizeLoadBalancerUrl(psc client.ServiceClient, popts IResizeLoadBalancerRequest) string {
	return psc.ServiceURL(
		psc.GetProjectId(),
		"loadBalancers",
		popts.GetLoadBalancerId(),
		"resize")
}

func listLoadBalancerPackagesUrl(psc client.ServiceClient, popts IListLoadBalancerPackagesRequest) string {
	return psc.ServiceURL(
		psc.GetProjectId(),
		"loadBalancers", "packages") + "?zoneId=" + popts.GetZoneId()
}

func getLoadBalancerByIdUrl(psc client.ServiceClient, popts IGetLoadBalancerByIdRequest) string {
	return psc.ServiceURL(
		psc.GetProjectId(),
		"loadBalancers",
		popts.GetLoadBalancerId())
}

func listLoadBalancersUrl(psc client.ServiceClient, popts IListLoadBalancersRequest) string {
	query, err := popts.ToListQuery()
	if err != nil {
		query = popts.GetDefaultQuery()
	}

	return psc.ServiceURL(psc.GetProjectId(), "loadBalancers") + query
}

func getPoolHealthMonitorByIdUrl(psc client.ServiceClient, popts IGetPoolHealthMonitorByIdRequest) string {
	return psc.ServiceURL(
		psc.GetProjectId(),
		"loadBalancers",
		popts.GetLoadBalancerId(),
		"pools",
		popts.GetPoolId(),
		"healthMonitor")
}

func createPoolUrl(psc client.ServiceClient, popts ICreatePoolRequest) string {
	return psc.ServiceURL(
		psc.GetProjectId(),
		"loadBalancers",
		popts.GetLoadBalancerId(),
		"pools")
}

func updatePoolUrl(psc client.ServiceClient, popts IUpdatePoolRequest) string {
	return psc.ServiceURL(
		psc.GetProjectId(),
		"loadBalancers",
		popts.GetLoadBalancerId(),
		"pools",
		popts.GetPoolId())
}

func createListenerUrl(psc client.ServiceClient, popts ICreateListenerRequest) string {
	return psc.ServiceURL(
		psc.GetProjectId(),
		"loadBalancers",
		popts.GetLoadBalancerId(),
		"listeners")
}

func updateListenerUrl(psc client.ServiceClient, popts IUpdateListenerRequest) string {
	return psc.ServiceURL(
		psc.GetProjectId(),
		"loadBalancers",
		popts.GetLoadBalancerId(),
		"listeners",
		popts.GetListenerId())
}

func listListenersByLoadBalancerIdUrl(psc client.ServiceClient, popts IListListenersByLoadBalancerIdRequest) string {
	return psc.ServiceURL(
		psc.GetProjectId(),
		"loadBalancers",
		popts.GetLoadBalancerId(),
		"listeners")
}

func listPoolsByLoadBalancerIdUrl(psc client.ServiceClient, popts IListPoolsByLoadBalancerIdRequest) string {
	return psc.ServiceURL(
		psc.GetProjectId(),
		"loadBalancers",
		popts.GetLoadBalancerId(),
		"pools")
}

func updatePoolMembersUrl(psc client.ServiceClient, popts IUpdatePoolMembersRequest) string {
	return psc.ServiceURL(
		psc.GetProjectId(),
		"loadBalancers",
		popts.GetLoadBalancerId(),
		"pools",
		popts.GetPoolId(),
		"members")
}

func listPoolMembersUrl(psc client.ServiceClient, popts IListPoolMembersRequest) string {
	return psc.ServiceURL(
		psc.GetProjectId(),
		"loadBalancers",
		popts.GetLoadBalancerId(),
		"pools",
		popts.GetPoolId(),
		"members")
}

func deletePoolByIdUrl(psc client.ServiceClient, popts IDeletePoolByIdRequest) string {
	return psc.ServiceURL(
		psc.GetProjectId(),
		"loadBalancers",
		popts.GetLoadBalancerId(),
		"pools",
		popts.GetPoolId())
}

func deleteListenerByIdUrl(psc client.ServiceClient, popts IDeleteListenerByIdRequest) string {
	return psc.ServiceURL(
		psc.GetProjectId(),
		"loadBalancers",
		popts.GetLoadBalancerId(),
		"listeners",
		popts.GetListenerId())
}

func deleteLoadBalancerByIdUrl(psc client.ServiceClient, popts IDeleteLoadBalancerByIdRequest) string {
	return psc.ServiceURL(
		psc.GetProjectId(),
		"loadBalancers",
		popts.GetLoadBalancerId())
}

func listTagsUrl(psc client.ServiceClient, popts IListTagsRequest) string {
	return psc.ServiceURL(
		psc.GetProjectId(),
		"tag",
		"resource",
		popts.GetLoadBalancerId())
}

func createTagsUrl(psc client.ServiceClient, popts ICreateTagsRequest) string {
	return psc.ServiceURL(
		psc.GetProjectId(),
		"tag",
		"resource",
		popts.GetLoadBalancerId())
}

func updateTagsUrl(psc client.ServiceClient, popts IUpdateTagsRequest) string {
	return psc.ServiceURL(
		psc.GetProjectId(),
		"tag",
		"resource",
		popts.GetLoadBalancerId())
}

// Policy

func listPoliciesUrl(psc client.ServiceClient, popts IListPoliciesRequest) string {
	return psc.ServiceURL(
		psc.GetProjectId(),
		"loadBalancers",
		popts.GetLoadBalancerId(),
		"listeners",
		popts.GetListenerId(),
		"l7policies",
	)
}

func createPolicyUrl(psc client.ServiceClient, popts ICreatePolicyRequest) string {
	return psc.ServiceURL(
		psc.GetProjectId(),
		"loadBalancers",
		popts.GetLoadBalancerId(),
		"listeners",
		popts.GetListenerId(),
		"l7policies",
	)
}

func getPolicyByIdUrl(psc client.ServiceClient, popts IGetPolicyByIdRequest) string {
	return psc.ServiceURL(
		psc.GetProjectId(),
		"loadBalancers",
		popts.GetLoadBalancerId(),
		"listeners",
		popts.GetListenerId(),
		"l7policies",
		popts.GetPolicyId(),
	)
}

func updatePolicyUrl(psc client.ServiceClient, popts IUpdatePolicyRequest) string {
	return psc.ServiceURL(
		psc.GetProjectId(),
		"loadBalancers",
		popts.GetLoadBalancerId(),
		"listeners",
		popts.GetListenerId(),
		"l7policies",
		popts.GetPolicyId(),
	)
}

func deletePolicyByIdUrl(psc client.ServiceClient, popts IDeletePolicyByIdRequest) string {
	return psc.ServiceURL(
		psc.GetProjectId(),
		"loadBalancers",
		popts.GetLoadBalancerId(),
		"listeners",
		popts.GetListenerId(),
		"l7policies",
		popts.GetPolicyId(),
	)
}

func reorderPoliciesUrl(psc client.ServiceClient, popts IReorderPoliciesRequest) string {
	return psc.ServiceURL(
		psc.GetProjectId(),
		"loadBalancers",
		popts.GetLoadBalancerId(),
		"listeners",
		popts.GetListenerId(),
		"reorderL7Policies")
}

func getPoolByIdUrl(psc client.ServiceClient, popts IGetPoolByIdRequest) string {
	return psc.ServiceURL(
		psc.GetProjectId(),
		"loadBalancers",
		popts.GetLoadBalancerId(),
		"pools",
		popts.GetPoolId())
}

func getListenerByIdUrl(psc client.ServiceClient, popts IGetListenerByIdRequest) string {
	return psc.ServiceURL(
		psc.GetProjectId(),
		"loadBalancers",
		popts.GetLoadBalancerId(),
		"listeners",
		popts.GetListenerId())
}

func resizeLoadBalancerByIdUrl(psc client.ServiceClient, popts IResizeLoadBalancerByIdRequest) string {
	return psc.ServiceURL(
		psc.GetProjectId(),
		"loadBalancers",
		popts.GetLoadBalancerId(),
		"resize")
}

func scaleLoadBalancerUrl(psc client.ServiceClient, popts IScaleLoadBalancerRequest) string {
	return psc.ServiceURL(
		psc.GetProjectId(),
		"loadBalancers",
		popts.GetLoadBalancerId(),
		"rebalancing")
}

// --------------------------------------------------------

func listCertificatesUrl(psc client.ServiceClient) string {
	return psc.ServiceURL(
		psc.GetProjectId(),
		"cas")
}

func createCertificateUrl(psc client.ServiceClient) string {
	return psc.ServiceURL(
		psc.GetProjectId(),
		"cas")
}

func getCertificateByIdUrl(psc client.ServiceClient, popts IGetCertificateByIdRequest) string {
	return psc.ServiceURL(
		psc.GetProjectId(),
		"cas",
		popts.GetCertificateId())
}

func deleteCertificateByIdUrl(psc client.ServiceClient, popts IDeleteCertificateByIdRequest) string {
	return psc.ServiceURL(
		psc.GetProjectId(),
		"cas",
		popts.GetCertificateId())
}
