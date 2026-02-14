package v2

func NewCreateServerRequest(name, imageID, flavorID, rootDiskType string, rootDiskSize int) *CreateServerRequest {
	opt := new(CreateServerRequest)
	opt.Name = name
	opt.ImageID = imageID
	opt.FlavorID = flavorID
	opt.RootDiskTypeID = rootDiskType
	opt.RootDiskSize = rootDiskSize
	return opt
}

func NewGetServerByIDRequest(serverID string) *GetServerByIDRequest {
	opt := new(GetServerByIDRequest)
	opt.ServerID = serverID
	return opt
}

func NewDeleteServerByIDRequest(serverID string) *DeleteServerByIDRequest {
	opt := new(DeleteServerByIDRequest)
	opt.ServerID = serverID
	opt.DeleteAllVolume = false
	return opt
}

func NewUpdateServerSecgroupsRequest(serverID string, secgroups ...string) *UpdateServerSecgroupsByServerIDRequest {
	opt := new(UpdateServerSecgroupsByServerIDRequest)
	opt.ServerID = serverID
	opt.Secgroups = secgroups
	return opt
}

func NewAttachBlockVolumeRequest(serverID, volumeID string) *AttachBlockVolumeRequest {
	opt := new(AttachBlockVolumeRequest)
	opt.ServerID = serverID
	opt.BlockVolumeID = volumeID
	return opt
}

func NewDetachBlockVolumeRequest(serverID, volumeID string) *DetachBlockVolumeRequest {
	opt := new(DetachBlockVolumeRequest)
	opt.ServerID = serverID
	opt.BlockVolumeID = volumeID
	return opt
}

func NewAttachFloatingIpRequest(serverID, niid string) *AttachFloatingIpRequest {
	opt := new(AttachFloatingIpRequest)
	opt.ServerID = serverID
	opt.InternalNetworkInterfaceID = niid
	opt.NetworkInterfaceID = niid
	return opt
}

func NewDetachFloatingIpRequest(serverID, wanID, niid string) *DetachFloatingIpRequest {
	opt := new(DetachFloatingIpRequest)
	opt.ServerID = serverID
	opt.InternalNetworkInterfaceID = niid
	opt.NetworkInterfaceID = niid
	opt.WanID = wanID
	return opt
}

func NewListServerGroupPoliciesRequest() *ListServerGroupPoliciesRequest {
	return new(ListServerGroupPoliciesRequest)
}

func NewDeleteServerGroupByIDRequest(serverGroupID string) *DeleteServerGroupByIDRequest {
	opt := new(DeleteServerGroupByIDRequest)
	opt.ServerGroupID = serverGroupID
	return opt
}

func NewListServerGroupsRequest(page, size int) *ListServerGroupsRequest {
	opt := new(ListServerGroupsRequest)
	opt.Page = page
	opt.Size = size
	opt.Name = ""

	return opt
}

func NewCreateServerGroupRequest(name, description, policyID string) *CreateServerGroupRequest {
	opt := new(CreateServerGroupRequest)
	opt.Name = name
	opt.Description = description
	opt.PolicyID = policyID

	return opt
}
