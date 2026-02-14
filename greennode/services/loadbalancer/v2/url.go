package v2

import "github.com/dannyota/greennode-community-sdk/v2/greennode/client"

func createLoadBalancerUrl(psc client.IServiceClient) string {
	return psc.ServiceURL(
		psc.GetProjectId(),
		"loadBalancers")
}

func resizeLoadBalancerUrl(psc client.IServiceClient, popts IResizeLoadBalancerRequest) string {
	return psc.ServiceURL(
		psc.GetProjectId(),
		"loadBalancers",
		popts.GetLoadBalancerId(),
		"resize")
}

func listLoadBalancerPackagesUrl(psc client.IServiceClient, popts IListLoadBalancerPackagesRequest) string {
	return psc.ServiceURL(
		psc.GetProjectId(),
		"loadBalancers", "packages") + "?zoneId=" + popts.GetZoneId()
}

func getLoadBalancerByIdUrl(psc client.IServiceClient, popts IGetLoadBalancerByIdRequest) string {
	return psc.ServiceURL(
		psc.GetProjectId(),
		"loadBalancers",
		popts.GetLoadBalancerId())
}

func listLoadBalancersUrl(psc client.IServiceClient, popts IListLoadBalancersRequest) string {
	query, err := popts.ToListQuery()
	if err != nil {
		query = popts.GetDefaultQuery()
	}

	return psc.ServiceURL(psc.GetProjectId(), "loadBalancers") + query
}

func getPoolHealthMonitorByIdUrl(psc client.IServiceClient, popts IGetPoolHealthMonitorByIdRequest) string {
	return psc.ServiceURL(
		psc.GetProjectId(),
		"loadBalancers",
		popts.GetLoadBalancerId(),
		"pools",
		popts.GetPoolId(),
		"healthMonitor")
}

func createPoolUrl(psc client.IServiceClient, popts ICreatePoolRequest) string {
	return psc.ServiceURL(
		psc.GetProjectId(),
		"loadBalancers",
		popts.GetLoadBalancerId(),
		"pools")
}

func updatePoolUrl(psc client.IServiceClient, popts IUpdatePoolRequest) string {
	return psc.ServiceURL(
		psc.GetProjectId(),
		"loadBalancers",
		popts.GetLoadBalancerId(),
		"pools",
		popts.GetPoolId())
}

func createListenerUrl(psc client.IServiceClient, popts ICreateListenerRequest) string {
	return psc.ServiceURL(
		psc.GetProjectId(),
		"loadBalancers",
		popts.GetLoadBalancerId(),
		"listeners")
}

func updateListenerUrl(psc client.IServiceClient, popts IUpdateListenerRequest) string {
	return psc.ServiceURL(
		psc.GetProjectId(),
		"loadBalancers",
		popts.GetLoadBalancerId(),
		"listeners",
		popts.GetListenerId())
}

func listListenersByLoadBalancerIdUrl(psc client.IServiceClient, popts IListListenersByLoadBalancerIdRequest) string {
	return psc.ServiceURL(
		psc.GetProjectId(),
		"loadBalancers",
		popts.GetLoadBalancerId(),
		"listeners")
}

func listPoolsByLoadBalancerIdUrl(psc client.IServiceClient, popts IListPoolsByLoadBalancerIdRequest) string {
	return psc.ServiceURL(
		psc.GetProjectId(),
		"loadBalancers",
		popts.GetLoadBalancerId(),
		"pools")
}

func updatePoolMembersUrl(psc client.IServiceClient, popts IUpdatePoolMembersRequest) string {
	return psc.ServiceURL(
		psc.GetProjectId(),
		"loadBalancers",
		popts.GetLoadBalancerId(),
		"pools",
		popts.GetPoolId(),
		"members")
}

func listPoolMembersUrl(psc client.IServiceClient, popts IListPoolMembersRequest) string {
	return psc.ServiceURL(
		psc.GetProjectId(),
		"loadBalancers",
		popts.GetLoadBalancerId(),
		"pools",
		popts.GetPoolId(),
		"members")
}

func deletePoolByIdUrl(psc client.IServiceClient, popts IDeletePoolByIdRequest) string {
	return psc.ServiceURL(
		psc.GetProjectId(),
		"loadBalancers",
		popts.GetLoadBalancerId(),
		"pools",
		popts.GetPoolId())
}

func deleteListenerByIdUrl(psc client.IServiceClient, popts IDeleteListenerByIdRequest) string {
	return psc.ServiceURL(
		psc.GetProjectId(),
		"loadBalancers",
		popts.GetLoadBalancerId(),
		"listeners",
		popts.GetListenerId())
}

func deleteLoadBalancerByIdUrl(psc client.IServiceClient, popts IDeleteLoadBalancerByIdRequest) string {
	return psc.ServiceURL(
		psc.GetProjectId(),
		"loadBalancers",
		popts.GetLoadBalancerId())
}

func listTagsUrl(psc client.IServiceClient, popts IListTagsRequest) string {
	return psc.ServiceURL(
		psc.GetProjectId(),
		"tag",
		"resource",
		popts.GetLoadBalancerId())
}

func createTagsUrl(psc client.IServiceClient, popts ICreateTagsRequest) string {
	return psc.ServiceURL(
		psc.GetProjectId(),
		"tag",
		"resource",
		popts.GetLoadBalancerId())
}

func updateTagsUrl(psc client.IServiceClient, popts IUpdateTagsRequest) string {
	return psc.ServiceURL(
		psc.GetProjectId(),
		"tag",
		"resource",
		popts.GetLoadBalancerId())
}

// Policy

func listPoliciesUrl(psc client.IServiceClient, popts IListPoliciesRequest) string {
	return psc.ServiceURL(
		psc.GetProjectId(),
		"loadBalancers",
		popts.GetLoadBalancerId(),
		"listeners",
		popts.GetListenerId(),
		"l7policies",
	)
}

func createPolicyUrl(psc client.IServiceClient, popts ICreatePolicyRequest) string {
	return psc.ServiceURL(
		psc.GetProjectId(),
		"loadBalancers",
		popts.GetLoadBalancerId(),
		"listeners",
		popts.GetListenerId(),
		"l7policies",
	)
}

func getPolicyByIdUrl(psc client.IServiceClient, popts IGetPolicyByIdRequest) string {
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

func updatePolicyUrl(psc client.IServiceClient, popts IUpdatePolicyRequest) string {
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

func deletePolicyByIdUrl(psc client.IServiceClient, popts IDeletePolicyByIdRequest) string {
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

func reorderPoliciesUrl(psc client.IServiceClient, popts IReorderPoliciesRequest) string {
	return psc.ServiceURL(
		psc.GetProjectId(),
		"loadBalancers",
		popts.GetLoadBalancerId(),
		"listeners",
		popts.GetListenerId(),
		"reorderL7Policies")
}

func getPoolByIdUrl(psc client.IServiceClient, popts IGetPoolByIdRequest) string {
	return psc.ServiceURL(
		psc.GetProjectId(),
		"loadBalancers",
		popts.GetLoadBalancerId(),
		"pools",
		popts.GetPoolId())
}

func getListenerByIdUrl(psc client.IServiceClient, popts IGetListenerByIdRequest) string {
	return psc.ServiceURL(
		psc.GetProjectId(),
		"loadBalancers",
		popts.GetLoadBalancerId(),
		"listeners",
		popts.GetListenerId())
}

func resizeLoadBalancerByIdUrl(psc client.IServiceClient, popts IResizeLoadBalancerByIdRequest) string {
	return psc.ServiceURL(
		psc.GetProjectId(),
		"loadBalancers",
		popts.GetLoadBalancerId(),
		"resize")
}

func scaleLoadBalancerUrl(psc client.IServiceClient, popts IScaleLoadBalancerRequest) string {
	return psc.ServiceURL(
		psc.GetProjectId(),
		"loadBalancers",
		popts.GetLoadBalancerId(),
		"rebalancing")
}

// --------------------------------------------------------

func listCertificatesUrl(psc client.IServiceClient) string {
	return psc.ServiceURL(
		psc.GetProjectId(),
		"cas")
}

func createCertificateUrl(psc client.IServiceClient) string {
	return psc.ServiceURL(
		psc.GetProjectId(),
		"cas")
}

func getCertificateByIdUrl(psc client.IServiceClient, popts IGetCertificateByIdRequest) string {
	return psc.ServiceURL(
		psc.GetProjectId(),
		"cas",
		popts.GetCertificateId())
}

func deleteCertificateByIdUrl(psc client.IServiceClient, popts IDeleteCertificateByIdRequest) string {
	return psc.ServiceURL(
		psc.GetProjectId(),
		"cas",
		popts.GetCertificateId())
}
