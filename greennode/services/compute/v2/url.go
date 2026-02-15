package v2

import "github.com/dannyota/greennode-community-sdk/v2/greennode/client"

func createServerURL(sc client.ServiceClient) string {
	return sc.ServiceURL(
		sc.GetProjectID(),
		"servers")
}

func getServerByIDURL(sc client.ServiceClient, opts *GetServerByIDRequest) string {
	return sc.ServiceURL(
		sc.GetProjectID(),
		"servers",
		opts.GetServerID())
}

func deleteServerByIDURL(sc client.ServiceClient, opts *DeleteServerByIDRequest) string {
	return sc.ServiceURL(
		sc.GetProjectID(),
		"servers",
		opts.GetServerID())
}

func updateServerSecgroupsByServerIDURL(sc client.ServiceClient, opts *UpdateServerSecgroupsByServerIDRequest) string {
	return sc.ServiceURL(
		sc.GetProjectID(),
		"servers",
		opts.GetServerID(),
		"update-sec-group")
}

func attachBlockVolumeURL(sc client.ServiceClient, opts *AttachBlockVolumeRequest) string {
	return sc.ServiceURL(
		sc.GetProjectID(),
		"volumes",
		opts.GetBlockVolumeID(),
		"servers",
		opts.GetServerID(),
		"attach")
}

func detachBlockVolumeURL(sc client.ServiceClient, opts *DetachBlockVolumeRequest) string {
	return sc.ServiceURL(
		sc.GetProjectID(),
		"volumes",
		opts.GetBlockVolumeID(),
		"servers",
		opts.GetServerID(),
		"detach",
	)
}

func attachFloatingIpURL(sc client.ServiceClient, opts *AttachFloatingIpRequest) string {
	return sc.ServiceURL(
		sc.GetProjectID(),
		"servers",
		opts.GetServerID(),
		"wan-ips",
		"auto",
		"attach")

}

func detachFloatingIpURL(sc client.ServiceClient, opts *DetachFloatingIpRequest) string {
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

func deleteServerGroupByIDURL(sc client.ServiceClient, opts *DeleteServerGroupByIDRequest) string {
	return sc.ServiceURL(
		sc.GetProjectID(),
		"serverGroups",
		opts.GetServerGroupID(),
	)
}

func listServerGroupsURL(sc client.ServiceClient, opts *ListServerGroupsRequest) string {
	query, err := opts.ToListQuery()
	if err != nil {
		query = opts.GetDefaultQuery()
	}

	return sc.ServiceURL(sc.GetProjectID(), "serverGroups") + query
}

func createServerGroupURL(sc client.ServiceClient, _ *CreateServerGroupRequest) string {
	return sc.ServiceURL(
		sc.GetProjectID(),
		"serverGroups",
	)
}
