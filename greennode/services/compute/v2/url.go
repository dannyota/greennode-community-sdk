package v2

import "github.com/dannyota/greennode-community-sdk/v2/greennode/client"

func createServerUrl(psc client.ServiceClient) string {
	return psc.ServiceURL(
		psc.GetProjectId(),
		"servers")
}

func getServerByIdUrl(psc client.ServiceClient, popts IGetServerByIdRequest) string {
	return psc.ServiceURL(
		psc.GetProjectId(),
		"servers",
		popts.GetServerId())
}

func deleteServerByIdUrl(psc client.ServiceClient, popts IDeleteServerByIdRequest) string {
	return psc.ServiceURL(
		psc.GetProjectId(),
		"servers",
		popts.GetServerId())
}

func updateServerSecgroupsByServerIdUrl(psc client.ServiceClient, popts IUpdateServerSecgroupsByServerIdRequest) string {
	return psc.ServiceURL(
		psc.GetProjectId(),
		"servers",
		popts.GetServerId(),
		"update-sec-group")
}

func attachBlockVolumeUrl(psc client.ServiceClient, popts IAttachBlockVolumeRequest) string {
	return psc.ServiceURL(
		psc.GetProjectId(),
		"volumes",
		popts.GetBlockVolumeId(),
		"servers",
		popts.GetServerId(),
		"attach")
}

func detachBlockVolumeUrl(psc client.ServiceClient, popts IDetachBlockVolumeRequest) string {
	return psc.ServiceURL(
		psc.GetProjectId(),
		"volumes",
		popts.GetBlockVolumeId(),
		"servers",
		popts.GetServerId(),
		"detach",
	)
}

func attachFloatingIpUrl(psc client.ServiceClient, popts IAttachFloatingIpRequest) string {
	return psc.ServiceURL(
		psc.GetProjectId(),
		"servers",
		popts.GetServerId(),
		"wan-ips",
		"auto",
		"attach")

}

func detachFloatingIpUrl(psc client.ServiceClient, popts IDetachFloatingIpRequest) string {
	return psc.ServiceURL(
		psc.GetProjectId(),
		"servers",
		popts.GetServerId(),
		"wan-ips",
		popts.GetWanId(),
		"detach")
}

func listServerGroupPoliciesUrl(psc client.ServiceClient) string {
	return psc.ServiceURL(
		psc.GetProjectId(),
		"serverGroups",
		"policies",
	)
}

func deleteServerGroupByIdUrl(psc client.ServiceClient, popts IDeleteServerGroupByIdRequest) string {
	return psc.ServiceURL(
		psc.GetProjectId(),
		"serverGroups",
		popts.GetServerGroupId(),
	)
}

func listServerGroupsUrl(psc client.ServiceClient, popts IListServerGroupsRequest) string {
	query, err := popts.ToListQuery()
	if err != nil {
		query = popts.GetDefaultQuery()
	}

	return psc.ServiceURL(psc.GetProjectId(), "serverGroups") + query
}

func createServerGroupUrl(psc client.ServiceClient, _ ICreateServerGroupRequest) string {
	return psc.ServiceURL(
		psc.GetProjectId(),
		"serverGroups",
	)
}
