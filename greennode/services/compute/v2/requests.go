package v2

func NewCreateServerRequest(name, imageId, flavorId, rootDiskType string, rootDiskSize int) ICreateServerRequest {
	opt := new(CreateServerRequest)
	opt.Name = name
	opt.ImageId = imageId
	opt.FlavorId = flavorId
	opt.RootDiskTypeId = rootDiskType
	opt.RootDiskSize = rootDiskSize
	return opt
}

func NewGetServerByIdRequest(serverId string) IGetServerByIdRequest {
	opt := new(GetServerByIdRequest)
	opt.ServerId = serverId
	return opt
}

func NewDeleteServerByIdRequest(serverId string) IDeleteServerByIdRequest {
	opt := new(DeleteServerByIdRequest)
	opt.ServerId = serverId
	opt.DeleteAllVolume = false
	return opt
}

func NewUpdateServerSecgroupsRequest(serverId string, secgroups ...string) IUpdateServerSecgroupsByServerIdRequest {
	opt := new(UpdateServerSecgroupsByServerIdRequest)
	opt.ServerId = serverId
	opt.Secgroups = secgroups
	return opt
}

func NewAttachBlockVolumeRequest(serverId, volumeId string) IAttachBlockVolumeRequest {
	opt := new(AttachBlockVolumeRequest)
	opt.ServerId = serverId
	opt.BlockVolumeId = volumeId
	return opt
}

func NewDetachBlockVolumeRequest(serverId, volumeId string) IDetachBlockVolumeRequest {
	opt := new(DetachBlockVolumeRequest)
	opt.ServerId = serverId
	opt.BlockVolumeId = volumeId
	return opt
}

func NewAttachFloatingIpRequest(serverId, niid string) IAttachFloatingIpRequest {
	opt := new(AttachFloatingIpRequest)
	opt.ServerId = serverId
	opt.InternalNetworkInterfaceId = niid
	opt.NetworkInterfaceId = niid
	return opt
}

func NewDetachFloatingIpRequest(serverId, wanId, niid string) IDetachFloatingIpRequest {
	opt := new(DetachFloatingIpRequest)
	opt.ServerId = serverId
	opt.InternalNetworkInterfaceId = niid
	opt.NetworkInterfaceId = niid
	opt.WanId = wanId
	return opt
}

func NewListServerGroupPoliciesRequest() IListServerGroupPoliciesRequest {
	return new(ListServerGroupPoliciesRequest)
}

func NewDeleteServerGroupByIdRequest(serverGroupId string) IDeleteServerGroupByIdRequest {
	opt := new(DeleteServerGroupByIdRequest)
	opt.ServerGroupId = serverGroupId
	return opt
}

func NewListServerGroupsRequest(page, size int) IListServerGroupsRequest {
	opt := new(ListServerGroupsRequest)
	opt.Page = page
	opt.Size = size
	opt.Name = ""

	return opt
}

func NewCreateServerGroupRequest(name, description, policyId string) ICreateServerGroupRequest {
	opt := new(CreateServerGroupRequest)
	opt.Name = name
	opt.Description = description
	opt.PolicyId = policyId

	return opt
}
