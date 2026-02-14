package v2

import "github.com/dannyota/greennode-community-sdk/v2/greennode/client"

func createServerUrl(psc client.IServiceClient) string {
	return psc.ServiceURL(
		psc.GetProjectId(),
		"servers")
}

func getServerByIdUrl(psc client.IServiceClient, popts IGetServerByIdRequest) string {
	return psc.ServiceURL(
		psc.GetProjectId(),
		"servers",
		popts.GetServerId())
}

func deleteServerByIdUrl(psc client.IServiceClient, popts IDeleteServerByIdRequest) string {
	return psc.ServiceURL(
		psc.GetProjectId(),
		"servers",
		popts.GetServerId())
}

func updateServerSecgroupsByServerIdUrl(psc client.IServiceClient, popts IUpdateServerSecgroupsByServerIdRequest) string {
	return psc.ServiceURL(
		psc.GetProjectId(),
		"servers",
		popts.GetServerId(),
		"update-sec-group")
}

func attachBlockVolumeUrl(psc client.IServiceClient, popts IAttachBlockVolumeRequest) string {
	return psc.ServiceURL(
		psc.GetProjectId(),
		"volumes",
		popts.GetBlockVolumeId(),
		"servers",
		popts.GetServerId(),
		"attach")
}

func detachBlockVolumeUrl(psc client.IServiceClient, popts IDetachBlockVolumeRequest) string {
	return psc.ServiceURL(
		psc.GetProjectId(),
		"volumes",
		popts.GetBlockVolumeId(),
		"servers",
		popts.GetServerId(),
		"detach",
	)
}

func attachFloatingIpUrl(psc client.IServiceClient, popts IAttachFloatingIpRequest) string {
	return psc.ServiceURL(
		psc.GetProjectId(),
		"servers",
		popts.GetServerId(),
		"wan-ips",
		"auto",
		"attach")

}

func detachFloatingIpUrl(psc client.IServiceClient, popts IDetachFloatingIpRequest) string {
	return psc.ServiceURL(
		psc.GetProjectId(),
		"servers",
		popts.GetServerId(),
		"wan-ips",
		popts.GetWanId(),
		"detach")
}

func listServerGroupPoliciesUrl(psc client.IServiceClient) string {
	return psc.ServiceURL(
		psc.GetProjectId(),
		"serverGroups",
		"policies",
	)
}

func deleteServerGroupByIdUrl(psc client.IServiceClient, popts IDeleteServerGroupByIdRequest) string {
	return psc.ServiceURL(
		psc.GetProjectId(),
		"serverGroups",
		popts.GetServerGroupId(),
	)
}

func listServerGroupsUrl(psc client.IServiceClient, popts IListServerGroupsRequest) string {
	query, err := popts.ToListQuery()
	if err != nil {
		query = popts.GetDefaultQuery()
	}

	return psc.ServiceURL(psc.GetProjectId(), "serverGroups") + query
}

func createServerGroupUrl(psc client.IServiceClient, _ ICreateServerGroupRequest) string {
	return psc.ServiceURL(
		psc.GetProjectId(),
		"serverGroups",
	)
}
