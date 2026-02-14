package v2

func NewCreateServerRequest(name, imageID, flavorID, rootDiskType string, rootDiskSize int) ICreateServerRequest {
	opt := new(CreateServerRequest)
	opt.Name = name
	opt.ImageID = imageID
	opt.FlavorID = flavorID
	opt.RootDiskTypeID = rootDiskType
	opt.RootDiskSize = rootDiskSize
	return opt
}

func NewGetServerByIDRequest(serverID string) IGetServerByIDRequest {
	opt := new(GetServerByIDRequest)
	opt.ServerID = serverID
	return opt
}

func NewDeleteServerByIDRequest(serverID string) IDeleteServerByIDRequest {
	opt := new(DeleteServerByIDRequest)
	opt.ServerID = serverID
	opt.DeleteAllVolume = false
	return opt
}

func NewUpdateServerSecgroupsRequest(serverID string, secgroups ...string) IUpdateServerSecgroupsByServerIDRequest {
	opt := new(UpdateServerSecgroupsByServerIDRequest)
	opt.ServerID = serverID
	opt.Secgroups = secgroups
	return opt
}

func NewAttachBlockVolumeRequest(serverID, volumeID string) IAttachBlockVolumeRequest {
	opt := new(AttachBlockVolumeRequest)
	opt.ServerID = serverID
	opt.BlockVolumeID = volumeID
	return opt
}

func NewDetachBlockVolumeRequest(serverID, volumeID string) IDetachBlockVolumeRequest {
	opt := new(DetachBlockVolumeRequest)
	opt.ServerID = serverID
	opt.BlockVolumeID = volumeID
	return opt
}

func NewAttachFloatingIpRequest(serverID, niid string) IAttachFloatingIpRequest {
	opt := new(AttachFloatingIpRequest)
	opt.ServerID = serverID
	opt.InternalNetworkInterfaceID = niid
	opt.NetworkInterfaceID = niid
	return opt
}

func NewDetachFloatingIpRequest(serverID, wanID, niid string) IDetachFloatingIpRequest {
	opt := new(DetachFloatingIpRequest)
	opt.ServerID = serverID
	opt.InternalNetworkInterfaceID = niid
	opt.NetworkInterfaceID = niid
	opt.WanID = wanID
	return opt
}

func NewListServerGroupPoliciesRequest() IListServerGroupPoliciesRequest {
	return new(ListServerGroupPoliciesRequest)
}

func NewDeleteServerGroupByIDRequest(serverGroupID string) IDeleteServerGroupByIDRequest {
	opt := new(DeleteServerGroupByIDRequest)
	opt.ServerGroupID = serverGroupID
	return opt
}

func NewListServerGroupsRequest(page, size int) IListServerGroupsRequest {
	opt := new(ListServerGroupsRequest)
	opt.Page = page
	opt.Size = size
	opt.Name = ""

	return opt
}

func NewCreateServerGroupRequest(name, description, policyID string) ICreateServerGroupRequest {
	opt := new(CreateServerGroupRequest)
	opt.Name = name
	opt.Description = description
	opt.PolicyID = policyID

	return opt
}
