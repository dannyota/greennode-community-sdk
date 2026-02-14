package v2

import "github.com/dannyota/greennode-community-sdk/v2/greennode/client"

func createServerUrl(sc client.ServiceClient) string {
	return sc.ServiceURL(
		sc.GetProjectId(),
		"servers")
}

func getServerByIdUrl(sc client.ServiceClient, opts IGetServerByIdRequest) string {
	return sc.ServiceURL(
		sc.GetProjectId(),
		"servers",
		opts.GetServerId())
}

func deleteServerByIdUrl(sc client.ServiceClient, opts IDeleteServerByIdRequest) string {
	return sc.ServiceURL(
		sc.GetProjectId(),
		"servers",
		opts.GetServerId())
}

func updateServerSecgroupsByServerIdUrl(sc client.ServiceClient, opts IUpdateServerSecgroupsByServerIdRequest) string {
	return sc.ServiceURL(
		sc.GetProjectId(),
		"servers",
		opts.GetServerId(),
		"update-sec-group")
}

func attachBlockVolumeUrl(sc client.ServiceClient, opts IAttachBlockVolumeRequest) string {
	return sc.ServiceURL(
		sc.GetProjectId(),
		"volumes",
		opts.GetBlockVolumeId(),
		"servers",
		opts.GetServerId(),
		"attach")
}

func detachBlockVolumeUrl(sc client.ServiceClient, opts IDetachBlockVolumeRequest) string {
	return sc.ServiceURL(
		sc.GetProjectId(),
		"volumes",
		opts.GetBlockVolumeId(),
		"servers",
		opts.GetServerId(),
		"detach",
	)
}

func attachFloatingIpUrl(sc client.ServiceClient, opts IAttachFloatingIpRequest) string {
	return sc.ServiceURL(
		sc.GetProjectId(),
		"servers",
		opts.GetServerId(),
		"wan-ips",
		"auto",
		"attach")

}

func detachFloatingIpUrl(sc client.ServiceClient, opts IDetachFloatingIpRequest) string {
	return sc.ServiceURL(
		sc.GetProjectId(),
		"servers",
		opts.GetServerId(),
		"wan-ips",
		opts.GetWanId(),
		"detach")
}

func listServerGroupPoliciesUrl(sc client.ServiceClient) string {
	return sc.ServiceURL(
		sc.GetProjectId(),
		"serverGroups",
		"policies",
	)
}

func deleteServerGroupByIdUrl(sc client.ServiceClient, opts IDeleteServerGroupByIdRequest) string {
	return sc.ServiceURL(
		sc.GetProjectId(),
		"serverGroups",
		opts.GetServerGroupId(),
	)
}

func listServerGroupsUrl(sc client.ServiceClient, opts IListServerGroupsRequest) string {
	query, err := opts.ToListQuery()
	if err != nil {
		query = opts.GetDefaultQuery()
	}

	return sc.ServiceURL(sc.GetProjectId(), "serverGroups") + query
}

func createServerGroupUrl(sc client.ServiceClient, _ ICreateServerGroupRequest) string {
	return sc.ServiceURL(
		sc.GetProjectId(),
		"serverGroups",
	)
}
