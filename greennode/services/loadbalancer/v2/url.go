package v2

import "github.com/dannyota/greennode-community-sdk/greennode/client"

func createLoadBalancerURL(sc *client.ServiceClient) string {
	return sc.ServiceURL(
		sc.ProjectID,
		"loadBalancers")
}

func resizeLoadBalancerURL(sc *client.ServiceClient, opts *ResizeLoadBalancerRequest) string {
	return sc.ServiceURL(
		sc.ProjectID,
		"loadBalancers",
		opts.LoadBalancerID,
		"resize")
}

func listLoadBalancerPackagesURL(sc *client.ServiceClient, opts *ListLoadBalancerPackagesRequest) string {
	return sc.ServiceURL(
		sc.ProjectID,
		"loadBalancers", "packages") + "?zoneId=" + string(opts.ZoneID)
}

func getLoadBalancerByIDURL(sc *client.ServiceClient, opts *GetLoadBalancerByIDRequest) string {
	return sc.ServiceURL(
		sc.ProjectID,
		"loadBalancers",
		opts.LoadBalancerID)
}

func listLoadBalancersURL(sc *client.ServiceClient, opts *ListLoadBalancersRequest) string {
	query, err := opts.ToListQuery()
	if err != nil {
		query = opts.getDefaultQuery()
	}

	return sc.ServiceURL(sc.ProjectID, "loadBalancers") + query
}

func getPoolHealthMonitorByIDURL(sc *client.ServiceClient, opts *GetPoolHealthMonitorByIDRequest) string {
	return sc.ServiceURL(
		sc.ProjectID,
		"loadBalancers",
		opts.LoadBalancerID,
		"pools",
		opts.PoolID,
		"healthMonitor")
}

func createPoolURL(sc *client.ServiceClient, opts *CreatePoolRequest) string {
	return sc.ServiceURL(
		sc.ProjectID,
		"loadBalancers",
		opts.LoadBalancerID,
		"pools")
}

func updatePoolURL(sc *client.ServiceClient, opts *UpdatePoolRequest) string {
	return sc.ServiceURL(
		sc.ProjectID,
		"loadBalancers",
		opts.LoadBalancerID,
		"pools",
		opts.PoolID)
}

func createListenerURL(sc *client.ServiceClient, opts *CreateListenerRequest) string {
	return sc.ServiceURL(
		sc.ProjectID,
		"loadBalancers",
		opts.LoadBalancerID,
		"listeners")
}

func updateListenerURL(sc *client.ServiceClient, opts *UpdateListenerRequest) string {
	return sc.ServiceURL(
		sc.ProjectID,
		"loadBalancers",
		opts.LoadBalancerID,
		"listeners",
		opts.ListenerID)
}

func listListenersByLoadBalancerIDURL(sc *client.ServiceClient, opts *ListListenersByLoadBalancerIDRequest) string {
	return sc.ServiceURL(
		sc.ProjectID,
		"loadBalancers",
		opts.LoadBalancerID,
		"listeners")
}

func listPoolsByLoadBalancerIDURL(sc *client.ServiceClient, opts *ListPoolsByLoadBalancerIDRequest) string {
	return sc.ServiceURL(
		sc.ProjectID,
		"loadBalancers",
		opts.LoadBalancerID,
		"pools")
}

func updatePoolMembersURL(sc *client.ServiceClient, opts *UpdatePoolMembersRequest) string {
	return sc.ServiceURL(
		sc.ProjectID,
		"loadBalancers",
		opts.LoadBalancerID,
		"pools",
		opts.PoolID,
		"members")
}

func listPoolMembersURL(sc *client.ServiceClient, opts *ListPoolMembersRequest) string {
	return sc.ServiceURL(
		sc.ProjectID,
		"loadBalancers",
		opts.LoadBalancerID,
		"pools",
		opts.PoolID,
		"members")
}

func deletePoolByIDURL(sc *client.ServiceClient, opts *DeletePoolByIDRequest) string {
	return sc.ServiceURL(
		sc.ProjectID,
		"loadBalancers",
		opts.LoadBalancerID,
		"pools",
		opts.PoolID)
}

func deleteListenerByIDURL(sc *client.ServiceClient, opts *DeleteListenerByIDRequest) string {
	return sc.ServiceURL(
		sc.ProjectID,
		"loadBalancers",
		opts.LoadBalancerID,
		"listeners",
		opts.ListenerID)
}

func deleteLoadBalancerByIDURL(sc *client.ServiceClient, opts *DeleteLoadBalancerByIDRequest) string {
	return sc.ServiceURL(
		sc.ProjectID,
		"loadBalancers",
		opts.LoadBalancerID)
}

func listTagsURL(sc *client.ServiceClient, opts *ListTagsRequest) string {
	return sc.ServiceURL(
		sc.ProjectID,
		"tag",
		"resource",
		opts.LoadBalancerID)
}

func createTagsURL(sc *client.ServiceClient, opts *CreateTagsRequest) string {
	return sc.ServiceURL(
		sc.ProjectID,
		"tag",
		"resource",
		opts.LoadBalancerID)
}

func updateTagsURL(sc *client.ServiceClient, opts *UpdateTagsRequest) string {
	return sc.ServiceURL(
		sc.ProjectID,
		"tag",
		"resource",
		opts.LoadBalancerID)
}

// Policy

func listPoliciesURL(sc *client.ServiceClient, opts *ListPoliciesRequest) string {
	return sc.ServiceURL(
		sc.ProjectID,
		"loadBalancers",
		opts.LoadBalancerID,
		"listeners",
		opts.ListenerID,
		"l7policies",
	)
}

func createPolicyURL(sc *client.ServiceClient, opts *CreatePolicyRequest) string {
	return sc.ServiceURL(
		sc.ProjectID,
		"loadBalancers",
		opts.LoadBalancerID,
		"listeners",
		opts.ListenerID,
		"l7policies",
	)
}

func getPolicyByIDURL(sc *client.ServiceClient, opts *GetPolicyByIDRequest) string {
	return sc.ServiceURL(
		sc.ProjectID,
		"loadBalancers",
		opts.LoadBalancerID,
		"listeners",
		opts.ListenerID,
		"l7policies",
		opts.PolicyID,
	)
}

func updatePolicyURL(sc *client.ServiceClient, opts *UpdatePolicyRequest) string {
	return sc.ServiceURL(
		sc.ProjectID,
		"loadBalancers",
		opts.LoadBalancerID,
		"listeners",
		opts.ListenerID,
		"l7policies",
		opts.PolicyID,
	)
}

func deletePolicyByIDURL(sc *client.ServiceClient, opts *DeletePolicyByIDRequest) string {
	return sc.ServiceURL(
		sc.ProjectID,
		"loadBalancers",
		opts.LoadBalancerID,
		"listeners",
		opts.ListenerID,
		"l7policies",
		opts.PolicyID,
	)
}

func reorderPoliciesURL(sc *client.ServiceClient, opts *ReorderPoliciesRequest) string {
	return sc.ServiceURL(
		sc.ProjectID,
		"loadBalancers",
		opts.LoadBalancerID,
		"listeners",
		opts.ListenerID,
		"reorderL7Policies")
}

func getPoolByIDURL(sc *client.ServiceClient, opts *GetPoolByIDRequest) string {
	return sc.ServiceURL(
		sc.ProjectID,
		"loadBalancers",
		opts.LoadBalancerID,
		"pools",
		opts.PoolID)
}

func getListenerByIDURL(sc *client.ServiceClient, opts *GetListenerByIDRequest) string {
	return sc.ServiceURL(
		sc.ProjectID,
		"loadBalancers",
		opts.LoadBalancerID,
		"listeners",
		opts.ListenerID)
}

func resizeLoadBalancerByIDURL(sc *client.ServiceClient, opts *ResizeLoadBalancerByIDRequest) string {
	return sc.ServiceURL(
		sc.ProjectID,
		"loadBalancers",
		opts.LoadBalancerID,
		"resize")
}

func scaleLoadBalancerURL(sc *client.ServiceClient, opts *ScaleLoadBalancerRequest) string {
	return sc.ServiceURL(
		sc.ProjectID,
		"loadBalancers",
		opts.LoadBalancerID,
		"rebalancing")
}


func listCertificatesURL(sc *client.ServiceClient) string {
	return sc.ServiceURL(
		sc.ProjectID,
		"cas")
}

func createCertificateURL(sc *client.ServiceClient) string {
	return sc.ServiceURL(
		sc.ProjectID,
		"cas")
}

func getCertificateByIDURL(sc *client.ServiceClient, opts *GetCertificateByIDRequest) string {
	return sc.ServiceURL(
		sc.ProjectID,
		"cas",
		opts.CertificateID)
}

func deleteCertificateByIDURL(sc *client.ServiceClient, opts *DeleteCertificateByIDRequest) string {
	return sc.ServiceURL(
		sc.ProjectID,
		"cas",
		opts.CertificateID)
}
