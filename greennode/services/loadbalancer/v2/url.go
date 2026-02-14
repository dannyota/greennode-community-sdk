package v2

import "github.com/dannyota/greennode-community-sdk/v2/greennode/client"

func createLoadBalancerURL(sc client.ServiceClient) string {
	return sc.ServiceURL(
		sc.GetProjectID(),
		"loadBalancers")
}

func resizeLoadBalancerURL(sc client.ServiceClient, opts *ResizeLoadBalancerRequest) string {
	return sc.ServiceURL(
		sc.GetProjectID(),
		"loadBalancers",
		opts.GetLoadBalancerID(),
		"resize")
}

func listLoadBalancerPackagesURL(sc client.ServiceClient, opts *ListLoadBalancerPackagesRequest) string {
	return sc.ServiceURL(
		sc.GetProjectID(),
		"loadBalancers", "packages") + "?zoneId=" + opts.GetZoneID()
}

func getLoadBalancerByIDURL(sc client.ServiceClient, opts *GetLoadBalancerByIDRequest) string {
	return sc.ServiceURL(
		sc.GetProjectID(),
		"loadBalancers",
		opts.GetLoadBalancerID())
}

func listLoadBalancersURL(sc client.ServiceClient, opts *ListLoadBalancersRequest) string {
	query, err := opts.ToListQuery()
	if err != nil {
		query = opts.GetDefaultQuery()
	}

	return sc.ServiceURL(sc.GetProjectID(), "loadBalancers") + query
}

func getPoolHealthMonitorByIDURL(sc client.ServiceClient, opts *GetPoolHealthMonitorByIDRequest) string {
	return sc.ServiceURL(
		sc.GetProjectID(),
		"loadBalancers",
		opts.GetLoadBalancerID(),
		"pools",
		opts.GetPoolID(),
		"healthMonitor")
}

func createPoolURL(sc client.ServiceClient, opts *CreatePoolRequest) string {
	return sc.ServiceURL(
		sc.GetProjectID(),
		"loadBalancers",
		opts.GetLoadBalancerID(),
		"pools")
}

func updatePoolURL(sc client.ServiceClient, opts *UpdatePoolRequest) string {
	return sc.ServiceURL(
		sc.GetProjectID(),
		"loadBalancers",
		opts.GetLoadBalancerID(),
		"pools",
		opts.GetPoolID())
}

func createListenerURL(sc client.ServiceClient, opts *CreateListenerRequest) string {
	return sc.ServiceURL(
		sc.GetProjectID(),
		"loadBalancers",
		opts.GetLoadBalancerID(),
		"listeners")
}

func updateListenerURL(sc client.ServiceClient, opts *UpdateListenerRequest) string {
	return sc.ServiceURL(
		sc.GetProjectID(),
		"loadBalancers",
		opts.GetLoadBalancerID(),
		"listeners",
		opts.GetListenerID())
}

func listListenersByLoadBalancerIDURL(sc client.ServiceClient, opts *ListListenersByLoadBalancerIDRequest) string {
	return sc.ServiceURL(
		sc.GetProjectID(),
		"loadBalancers",
		opts.GetLoadBalancerID(),
		"listeners")
}

func listPoolsByLoadBalancerIDURL(sc client.ServiceClient, opts *ListPoolsByLoadBalancerIDRequest) string {
	return sc.ServiceURL(
		sc.GetProjectID(),
		"loadBalancers",
		opts.GetLoadBalancerID(),
		"pools")
}

func updatePoolMembersURL(sc client.ServiceClient, opts *UpdatePoolMembersRequest) string {
	return sc.ServiceURL(
		sc.GetProjectID(),
		"loadBalancers",
		opts.GetLoadBalancerID(),
		"pools",
		opts.GetPoolID(),
		"members")
}

func listPoolMembersURL(sc client.ServiceClient, opts *ListPoolMembersRequest) string {
	return sc.ServiceURL(
		sc.GetProjectID(),
		"loadBalancers",
		opts.GetLoadBalancerID(),
		"pools",
		opts.GetPoolID(),
		"members")
}

func deletePoolByIDURL(sc client.ServiceClient, opts *DeletePoolByIDRequest) string {
	return sc.ServiceURL(
		sc.GetProjectID(),
		"loadBalancers",
		opts.GetLoadBalancerID(),
		"pools",
		opts.GetPoolID())
}

func deleteListenerByIDURL(sc client.ServiceClient, opts *DeleteListenerByIDRequest) string {
	return sc.ServiceURL(
		sc.GetProjectID(),
		"loadBalancers",
		opts.GetLoadBalancerID(),
		"listeners",
		opts.GetListenerID())
}

func deleteLoadBalancerByIDURL(sc client.ServiceClient, opts *DeleteLoadBalancerByIDRequest) string {
	return sc.ServiceURL(
		sc.GetProjectID(),
		"loadBalancers",
		opts.GetLoadBalancerID())
}

func listTagsURL(sc client.ServiceClient, opts *ListTagsRequest) string {
	return sc.ServiceURL(
		sc.GetProjectID(),
		"tag",
		"resource",
		opts.GetLoadBalancerID())
}

func createTagsURL(sc client.ServiceClient, opts *CreateTagsRequest) string {
	return sc.ServiceURL(
		sc.GetProjectID(),
		"tag",
		"resource",
		opts.GetLoadBalancerID())
}

func updateTagsURL(sc client.ServiceClient, opts *UpdateTagsRequest) string {
	return sc.ServiceURL(
		sc.GetProjectID(),
		"tag",
		"resource",
		opts.GetLoadBalancerID())
}

// Policy

func listPoliciesURL(sc client.ServiceClient, opts *ListPoliciesRequest) string {
	return sc.ServiceURL(
		sc.GetProjectID(),
		"loadBalancers",
		opts.GetLoadBalancerID(),
		"listeners",
		opts.GetListenerID(),
		"l7policies",
	)
}

func createPolicyURL(sc client.ServiceClient, opts *CreatePolicyRequest) string {
	return sc.ServiceURL(
		sc.GetProjectID(),
		"loadBalancers",
		opts.GetLoadBalancerID(),
		"listeners",
		opts.GetListenerID(),
		"l7policies",
	)
}

func getPolicyByIDURL(sc client.ServiceClient, opts *GetPolicyByIDRequest) string {
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

func updatePolicyURL(sc client.ServiceClient, opts *UpdatePolicyRequest) string {
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

func deletePolicyByIDURL(sc client.ServiceClient, opts *DeletePolicyByIDRequest) string {
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

func reorderPoliciesURL(sc client.ServiceClient, opts *ReorderPoliciesRequest) string {
	return sc.ServiceURL(
		sc.GetProjectID(),
		"loadBalancers",
		opts.GetLoadBalancerID(),
		"listeners",
		opts.GetListenerID(),
		"reorderL7Policies")
}

func getPoolByIDURL(sc client.ServiceClient, opts *GetPoolByIDRequest) string {
	return sc.ServiceURL(
		sc.GetProjectID(),
		"loadBalancers",
		opts.GetLoadBalancerID(),
		"pools",
		opts.GetPoolID())
}

func getListenerByIDURL(sc client.ServiceClient, opts *GetListenerByIDRequest) string {
	return sc.ServiceURL(
		sc.GetProjectID(),
		"loadBalancers",
		opts.GetLoadBalancerID(),
		"listeners",
		opts.GetListenerID())
}

func resizeLoadBalancerByIDURL(sc client.ServiceClient, opts *ResizeLoadBalancerByIDRequest) string {
	return sc.ServiceURL(
		sc.GetProjectID(),
		"loadBalancers",
		opts.GetLoadBalancerID(),
		"resize")
}

func scaleLoadBalancerURL(sc client.ServiceClient, opts *ScaleLoadBalancerRequest) string {
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

func getCertificateByIDURL(sc client.ServiceClient, opts *GetCertificateByIDRequest) string {
	return sc.ServiceURL(
		sc.GetProjectID(),
		"cas",
		opts.GetCertificateID())
}

func deleteCertificateByIDURL(sc client.ServiceClient, opts *DeleteCertificateByIDRequest) string {
	return sc.ServiceURL(
		sc.GetProjectID(),
		"cas",
		opts.GetCertificateID())
}
