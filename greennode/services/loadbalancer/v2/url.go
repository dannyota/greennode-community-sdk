package v2

import "github.com/dannyota/greennode-community-sdk/v2/greennode/client"

func createLoadBalancerURL(sc client.ServiceClient) string {
	return sc.ServiceURL(
		sc.GetProjectID(),
		"loadBalancers")
}

func resizeLoadBalancerURL(sc client.ServiceClient, opts IResizeLoadBalancerRequest) string {
	return sc.ServiceURL(
		sc.GetProjectID(),
		"loadBalancers",
		opts.GetLoadBalancerID(),
		"resize")
}

func listLoadBalancerPackagesURL(sc client.ServiceClient, opts IListLoadBalancerPackagesRequest) string {
	return sc.ServiceURL(
		sc.GetProjectID(),
		"loadBalancers", "packages") + "?zoneId=" + opts.GetZoneID()
}

func getLoadBalancerByIDURL(sc client.ServiceClient, opts IGetLoadBalancerByIDRequest) string {
	return sc.ServiceURL(
		sc.GetProjectID(),
		"loadBalancers",
		opts.GetLoadBalancerID())
}

func listLoadBalancersURL(sc client.ServiceClient, opts IListLoadBalancersRequest) string {
	query, err := opts.ToListQuery()
	if err != nil {
		query = opts.GetDefaultQuery()
	}

	return sc.ServiceURL(sc.GetProjectID(), "loadBalancers") + query
}

func getPoolHealthMonitorByIDURL(sc client.ServiceClient, opts IGetPoolHealthMonitorByIDRequest) string {
	return sc.ServiceURL(
		sc.GetProjectID(),
		"loadBalancers",
		opts.GetLoadBalancerID(),
		"pools",
		opts.GetPoolID(),
		"healthMonitor")
}

func createPoolURL(sc client.ServiceClient, opts ICreatePoolRequest) string {
	return sc.ServiceURL(
		sc.GetProjectID(),
		"loadBalancers",
		opts.GetLoadBalancerID(),
		"pools")
}

func updatePoolURL(sc client.ServiceClient, opts IUpdatePoolRequest) string {
	return sc.ServiceURL(
		sc.GetProjectID(),
		"loadBalancers",
		opts.GetLoadBalancerID(),
		"pools",
		opts.GetPoolID())
}

func createListenerURL(sc client.ServiceClient, opts ICreateListenerRequest) string {
	return sc.ServiceURL(
		sc.GetProjectID(),
		"loadBalancers",
		opts.GetLoadBalancerID(),
		"listeners")
}

func updateListenerURL(sc client.ServiceClient, opts IUpdateListenerRequest) string {
	return sc.ServiceURL(
		sc.GetProjectID(),
		"loadBalancers",
		opts.GetLoadBalancerID(),
		"listeners",
		opts.GetListenerID())
}

func listListenersByLoadBalancerIDURL(sc client.ServiceClient, opts IListListenersByLoadBalancerIDRequest) string {
	return sc.ServiceURL(
		sc.GetProjectID(),
		"loadBalancers",
		opts.GetLoadBalancerID(),
		"listeners")
}

func listPoolsByLoadBalancerIDURL(sc client.ServiceClient, opts IListPoolsByLoadBalancerIDRequest) string {
	return sc.ServiceURL(
		sc.GetProjectID(),
		"loadBalancers",
		opts.GetLoadBalancerID(),
		"pools")
}

func updatePoolMembersURL(sc client.ServiceClient, opts IUpdatePoolMembersRequest) string {
	return sc.ServiceURL(
		sc.GetProjectID(),
		"loadBalancers",
		opts.GetLoadBalancerID(),
		"pools",
		opts.GetPoolID(),
		"members")
}

func listPoolMembersURL(sc client.ServiceClient, opts IListPoolMembersRequest) string {
	return sc.ServiceURL(
		sc.GetProjectID(),
		"loadBalancers",
		opts.GetLoadBalancerID(),
		"pools",
		opts.GetPoolID(),
		"members")
}

func deletePoolByIDURL(sc client.ServiceClient, opts IDeletePoolByIDRequest) string {
	return sc.ServiceURL(
		sc.GetProjectID(),
		"loadBalancers",
		opts.GetLoadBalancerID(),
		"pools",
		opts.GetPoolID())
}

func deleteListenerByIDURL(sc client.ServiceClient, opts IDeleteListenerByIDRequest) string {
	return sc.ServiceURL(
		sc.GetProjectID(),
		"loadBalancers",
		opts.GetLoadBalancerID(),
		"listeners",
		opts.GetListenerID())
}

func deleteLoadBalancerByIDURL(sc client.ServiceClient, opts IDeleteLoadBalancerByIDRequest) string {
	return sc.ServiceURL(
		sc.GetProjectID(),
		"loadBalancers",
		opts.GetLoadBalancerID())
}

func listTagsURL(sc client.ServiceClient, opts IListTagsRequest) string {
	return sc.ServiceURL(
		sc.GetProjectID(),
		"tag",
		"resource",
		opts.GetLoadBalancerID())
}

func createTagsURL(sc client.ServiceClient, opts ICreateTagsRequest) string {
	return sc.ServiceURL(
		sc.GetProjectID(),
		"tag",
		"resource",
		opts.GetLoadBalancerID())
}

func updateTagsURL(sc client.ServiceClient, opts IUpdateTagsRequest) string {
	return sc.ServiceURL(
		sc.GetProjectID(),
		"tag",
		"resource",
		opts.GetLoadBalancerID())
}

// Policy

func listPoliciesURL(sc client.ServiceClient, opts IListPoliciesRequest) string {
	return sc.ServiceURL(
		sc.GetProjectID(),
		"loadBalancers",
		opts.GetLoadBalancerID(),
		"listeners",
		opts.GetListenerID(),
		"l7policies",
	)
}

func createPolicyURL(sc client.ServiceClient, opts ICreatePolicyRequest) string {
	return sc.ServiceURL(
		sc.GetProjectID(),
		"loadBalancers",
		opts.GetLoadBalancerID(),
		"listeners",
		opts.GetListenerID(),
		"l7policies",
	)
}

func getPolicyByIDURL(sc client.ServiceClient, opts IGetPolicyByIDRequest) string {
	return sc.ServiceURL(
		sc.GetProjectID(),
		"loadBalancers",
		opts.GetLoadBalancerID(),
		"listeners",
		opts.GetListenerID(),
		"l7policies",
		opts.GetPolicyID(),
	)
}

func updatePolicyURL(sc client.ServiceClient, opts IUpdatePolicyRequest) string {
	return sc.ServiceURL(
		sc.GetProjectID(),
		"loadBalancers",
		opts.GetLoadBalancerID(),
		"listeners",
		opts.GetListenerID(),
		"l7policies",
		opts.GetPolicyID(),
	)
}

func deletePolicyByIDURL(sc client.ServiceClient, opts IDeletePolicyByIDRequest) string {
	return sc.ServiceURL(
		sc.GetProjectID(),
		"loadBalancers",
		opts.GetLoadBalancerID(),
		"listeners",
		opts.GetListenerID(),
		"l7policies",
		opts.GetPolicyID(),
	)
}

func reorderPoliciesURL(sc client.ServiceClient, opts IReorderPoliciesRequest) string {
	return sc.ServiceURL(
		sc.GetProjectID(),
		"loadBalancers",
		opts.GetLoadBalancerID(),
		"listeners",
		opts.GetListenerID(),
		"reorderL7Policies")
}

func getPoolByIDURL(sc client.ServiceClient, opts IGetPoolByIDRequest) string {
	return sc.ServiceURL(
		sc.GetProjectID(),
		"loadBalancers",
		opts.GetLoadBalancerID(),
		"pools",
		opts.GetPoolID())
}

func getListenerByIDURL(sc client.ServiceClient, opts IGetListenerByIDRequest) string {
	return sc.ServiceURL(
		sc.GetProjectID(),
		"loadBalancers",
		opts.GetLoadBalancerID(),
		"listeners",
		opts.GetListenerID())
}

func resizeLoadBalancerByIDURL(sc client.ServiceClient, opts IResizeLoadBalancerByIDRequest) string {
	return sc.ServiceURL(
		sc.GetProjectID(),
		"loadBalancers",
		opts.GetLoadBalancerID(),
		"resize")
}

func scaleLoadBalancerURL(sc client.ServiceClient, opts IScaleLoadBalancerRequest) string {
	return sc.ServiceURL(
		sc.GetProjectID(),
		"loadBalancers",
		opts.GetLoadBalancerID(),
		"rebalancing")
}


func listCertificatesURL(sc client.ServiceClient) string {
	return sc.ServiceURL(
		sc.GetProjectID(),
		"cas")
}

func createCertificateURL(sc client.ServiceClient) string {
	return sc.ServiceURL(
		sc.GetProjectID(),
		"cas")
}

func getCertificateByIDURL(sc client.ServiceClient, opts IGetCertificateByIDRequest) string {
	return sc.ServiceURL(
		sc.GetProjectID(),
		"cas",
		opts.GetCertificateID())
}

func deleteCertificateByIDURL(sc client.ServiceClient, opts IDeleteCertificateByIDRequest) string {
	return sc.ServiceURL(
		sc.GetProjectID(),
		"cas",
		opts.GetCertificateID())
}
