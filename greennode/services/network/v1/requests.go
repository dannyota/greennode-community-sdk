package v1

func NewGetEndpointByIdRequest(endpointId string) IGetEndpointByIdRequest {
	opt := new(GetEndpointByIdRequest)
	opt.EndpointId = endpointId
	return opt
}

func NewCreateEndpointRequest(name, serviceId, vpcId, subnetId string) ICreateEndpointRequest {
	opts := new(CreateEndpointRequest)
	opts.ResourceInfo.EndpointName = name
	opts.ResourceInfo.ServiceUuid = serviceId
	opts.ResourceInfo.VpcUuid = vpcId
	opts.ResourceInfo.SubnetUuid = subnetId
	opts.ResourceInfo.PackageUuid = defaultPackageId
	opts.ResourceInfo.EnableDnsName = false
	return opts
}

func NewDeleteEndpointByIdRequest(endpointId, vpcId, endpointServiceId string) IDeleteEndpointByIdRequest {
	opt := new(DeleteEndpointByIdRequest)
	opt.EndpointId = endpointId
	opt.EndpointUuid = endpointId
	opt.VpcUuid = vpcId
	opt.EndpointServiceUuid = endpointServiceId

	return opt
}

func NewListEndpointsRequest(page, size int) IListEndpointsRequest {
	return &ListEndpointsRequest{
		Page: page,
		Size: size,
	}
}

func NewListTagsByEndpointIdRequest(userId, projectId, endpointId string) IListTagsByEndpointIdRequest {
	opt := new(ListTagsByEndpointIdRequest)
	opt.Id = endpointId
	opt.EndpointId = endpointId
	opt.ProjectId = projectId
	opt.SetPortalUserId(userId)
	return opt
}

func NewCreateTagsWithEndpointIdRequest(userId, projectId, endpointId string) ICreateTagsWithEndpointIdRequest {
	opt := new(CreateTagsWithEndpointIdRequest)
	opt.ResourceUuid = endpointId
	opt.EndpointId = endpointId
	opt.SystemTag = true
	opt.ProjectId = projectId
	opt.SetPortalUserId(userId)

	return opt
}

func NewDeleteTagOfEndpointRequest(userId, projectId, tagId string) IDeleteTagOfEndpointRequest {
	opt := new(DeleteTagOfEndpointRequest)
	opt.TagId = tagId
	opt.ProjectId = projectId
	opt.SetPortalUserId(userId)

	return opt
}

func NewUpdateTagValueOfEndpointRequest(userId, projectId, tagId, value string) IUpdateTagValueOfEndpointRequest {
	opt := new(UpdateTagValueOfEndpointRequest)
	opt.TagId = tagId
	opt.TagValue = value
	opt.ProjectId = projectId
	opt.SetPortalUserId(userId)

	return opt
}
