package v2

import "danny.vn/greennode/client"

func createServerURL(sc *client.ServiceClient) string {
	return sc.ServiceURL(
		sc.ProjectID,
		"servers")
}

func getServerByIDURL(sc *client.ServiceClient, opts *GetServerByIDRequest) string {
	return sc.ServiceURL(
		sc.ProjectID,
		"servers",
		opts.ServerID)
}

func deleteServerByIDURL(sc *client.ServiceClient, opts *DeleteServerByIDRequest) string {
	return sc.ServiceURL(
		sc.ProjectID,
		"servers",
		opts.ServerID)
}

func updateServerSecgroupsByServerIDURL(sc *client.ServiceClient, opts *UpdateServerSecgroupsByServerIDRequest) string {
	return sc.ServiceURL(
		sc.ProjectID,
		"servers",
		opts.ServerID,
		"update-sec-group")
}

func attachBlockVolumeURL(sc *client.ServiceClient, opts *AttachBlockVolumeRequest) string {
	return sc.ServiceURL(
		sc.ProjectID,
		"volumes",
		opts.BlockVolumeID,
		"servers",
		opts.ServerID,
		"attach")
}

func detachBlockVolumeURL(sc *client.ServiceClient, opts *DetachBlockVolumeRequest) string {
	return sc.ServiceURL(
		sc.ProjectID,
		"volumes",
		opts.BlockVolumeID,
		"servers",
		opts.ServerID,
		"detach",
	)
}

func attachFloatingIpURL(sc *client.ServiceClient, opts *AttachFloatingIpRequest) string {
	return sc.ServiceURL(
		sc.ProjectID,
		"servers",
		opts.ServerID,
		"wan-ips",
		"auto",
		"attach")

}

func detachFloatingIpURL(sc *client.ServiceClient, opts *DetachFloatingIpRequest) string {
	return sc.ServiceURL(
		sc.ProjectID,
		"servers",
		opts.ServerID,
		"wan-ips",
		opts.WanID,
		"detach")
}

func listServersURL(sc *client.ServiceClient, opts *ListServersRequest) string {
	query, err := opts.ToListQuery()
	if err != nil {
		query = opts.getDefaultQuery()
	}

	base := sc.ServiceURL(sc.ProjectID, "servers")
	if query != "" {
		return base + "?" + query
	}
	return base
}

func listSSHKeysURL(sc *client.ServiceClient, opts *ListSSHKeysRequest) string {
	query, err := opts.ToListQuery()
	if err != nil {
		query = opts.getDefaultQuery()
	}

	base := sc.ServiceURL(sc.ProjectID, "sshKeys")
	if query != "" {
		return base + "?" + query
	}
	return base
}

func listServerGroupPoliciesURL(sc *client.ServiceClient) string {
	return sc.ServiceURL(
		sc.ProjectID,
		"serverGroups",
		"policies",
	)
}

func deleteServerGroupByIDURL(sc *client.ServiceClient, opts *DeleteServerGroupByIDRequest) string {
	return sc.ServiceURL(
		sc.ProjectID,
		"serverGroups",
		opts.ServerGroupID,
	)
}

func listServerGroupsURL(sc *client.ServiceClient, opts *ListServerGroupsRequest) string {
	query, err := opts.ToListQuery()
	if err != nil {
		query = opts.getDefaultQuery()
	}

	base := sc.ServiceURL(sc.ProjectID, "serverGroups")
	if query != "" {
		return base + "?" + query
	}
	return base
}

func createServerGroupURL(sc *client.ServiceClient, _ *CreateServerGroupRequest) string {
	return sc.ServiceURL(
		sc.ProjectID,
		"serverGroups",
	)
}
