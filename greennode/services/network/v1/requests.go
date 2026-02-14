package v1

func NewGetEndpointByIDRequest(endpointID string) IGetEndpointByIDRequest {
	opt := new(GetEndpointByIDRequest)
	opt.EndpointID = endpointID
	return opt
}

func NewCreateEndpointRequest(name, serviceID, vpcID, subnetID string) ICreateEndpointRequest {
	opts := new(CreateEndpointRequest)
	opts.ResourceInfo.EndpointName = name
	opts.ResourceInfo.ServiceUuid = serviceID
	opts.ResourceInfo.VpcUuid = vpcID
	opts.ResourceInfo.SubnetUuid = subnetID
	opts.ResourceInfo.PackageUuid = defaultPackageID
	opts.ResourceInfo.EnableDnsName = false
	return opts
}

func NewDeleteEndpointByIDRequest(endpointID, vpcID, endpointServiceID string) IDeleteEndpointByIDRequest {
	opt := new(DeleteEndpointByIDRequest)
	opt.EndpointID = endpointID
	opt.EndpointUuid = endpointID
	opt.VpcUuid = vpcID
	opt.EndpointServiceUuid = endpointServiceID

	return opt
}

func NewListEndpointsRequest(page, size int) IListEndpointsRequest {
	return &ListEndpointsRequest{
		Page: page,
		Size: size,
	}
}

func NewListTagsByEndpointIDRequest(userID, projectID, endpointID string) IListTagsByEndpointIDRequest {
	opt := new(ListTagsByEndpointIDRequest)
	opt.ID = endpointID
	opt.EndpointID = endpointID
	opt.ProjectID = projectID
	opt.SetPortalUserID(userID)
	return opt
}

func NewCreateTagsWithEndpointIDRequest(userID, projectID, endpointID string) ICreateTagsWithEndpointIDRequest {
	opt := new(CreateTagsWithEndpointIDRequest)
	opt.ResourceUuid = endpointID
	opt.EndpointID = endpointID
	opt.SystemTag = true
	opt.ProjectID = projectID
	opt.SetPortalUserID(userID)

	return opt
}

func NewDeleteTagOfEndpointRequest(userID, projectID, tagID string) IDeleteTagOfEndpointRequest {
	opt := new(DeleteTagOfEndpointRequest)
	opt.TagID = tagID
	opt.ProjectID = projectID
	opt.SetPortalUserID(userID)

	return opt
}

func NewUpdateTagValueOfEndpointRequest(userID, projectID, tagID, value string) IUpdateTagValueOfEndpointRequest {
	opt := new(UpdateTagValueOfEndpointRequest)
	opt.TagID = tagID
	opt.TagValue = value
	opt.ProjectID = projectID
	opt.SetPortalUserID(userID)

	return opt
}
