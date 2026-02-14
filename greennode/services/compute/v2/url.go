package v2

import "github.com/dannyota/greennode-community-sdk/v2/greennode/client"

func createServerURL(sc client.ServiceClient) string {
	return sc.ServiceURL(
		sc.GetProjectID(),
		"servers")
}

func getServerByIDURL(sc client.ServiceClient, opts IGetServerByIDRequest) string {
	return sc.ServiceURL(
		sc.GetProjectID(),
		"servers",
		opts.GetServerID())
}

func deleteServerByIDURL(sc client.ServiceClient, opts IDeleteServerByIDRequest) string {
	return sc.ServiceURL(
		sc.GetProjectID(),
		"servers",
		opts.GetServerID())
}

func updateServerSecgroupsByServerIDURL(sc client.ServiceClient, opts IUpdateServerSecgroupsByServerIDRequest) string {
	return sc.ServiceURL(
		sc.GetProjectID(),
		"servers",
		opts.GetServerID(),
		"update-sec-group")
}

func attachBlockVolumeURL(sc client.ServiceClient, opts IAttachBlockVolumeRequest) string {
	return sc.ServiceURL(
		sc.GetProjectID(),
		"volumes",
		opts.GetBlockVolumeID(),
		"servers",
		opts.GetServerID(),
		"attach")
}

func detachBlockVolumeURL(sc client.ServiceClient, opts IDetachBlockVolumeRequest) string {
	return sc.ServiceURL(
		sc.GetProjectID(),
		"volumes",
		opts.GetBlockVolumeID(),
		"servers",
		opts.GetServerID(),
		"detach",
	)
}

func attachFloatingIpURL(sc client.ServiceClient, opts IAttachFloatingIpRequest) string {
	return sc.ServiceURL(
		sc.GetProjectID(),
		"servers",
		opts.GetServerID(),
		"wan-ips",
		"auto",
		"attach")

}

func detachFloatingIpURL(sc client.ServiceClient, opts IDetachFloatingIpRequest) string {
	return sc.ServiceURL(
		sc.GetProjectID(),
		"servers",
		opts.GetServerID(),
		"wan-ips",
		opts.GetWanID(),
		"detach")
}

func listServerGroupPoliciesURL(sc client.ServiceClient) string {
	return sc.ServiceURL(
		sc.GetProjectID(),
		"serverGroups",
		"policies",
	)
}

func deleteServerGroupByIDURL(sc client.ServiceClient, opts IDeleteServerGroupByIDRequest) string {
	return sc.ServiceURL(
		sc.GetProjectID(),
		"serverGroups",
		opts.GetServerGroupID(),
	)
}

func listServerGroupsURL(sc client.ServiceClient, opts IListServerGroupsRequest) string {
	query, err := opts.ToListQuery()
	if err != nil {
		query = opts.GetDefaultQuery()
	}

	return sc.ServiceURL(sc.GetProjectID(), "serverGroups") + query
}

func createServerGroupURL(sc client.ServiceClient, _ ICreateServerGroupRequest) string {
	return sc.ServiceURL(
		sc.GetProjectID(),
		"serverGroups",
	)
}
