package v1

func NewGetEndpointByIDRequest(endpointID string) *GetEndpointByIDRequest {
	opt := new(GetEndpointByIDRequest)
	opt.EndpointID = endpointID
	return opt
}

func NewCreateEndpointRequest(name, serviceID, vpcID, subnetID string) *CreateEndpointRequest {
	opts := new(CreateEndpointRequest)
	opts.ResourceInfo.EndpointName = name
	opts.ResourceInfo.ServiceUuid = serviceID
	opts.ResourceInfo.VpcUuid = vpcID
	opts.ResourceInfo.SubnetUuid = subnetID
	opts.ResourceInfo.PackageUuid = defaultPackageID
	opts.ResourceInfo.EnableDnsName = false
	return opts
}

func NewDeleteEndpointByIDRequest(endpointID, vpcID, endpointServiceID string) *DeleteEndpointByIDRequest {
	opt := new(DeleteEndpointByIDRequest)
	opt.EndpointID = endpointID
	opt.EndpointUuid = endpointID
	opt.VpcUuid = vpcID
	opt.EndpointServiceUuid = endpointServiceID

	return opt
}

func NewListEndpointsRequest(page, size int) *ListEndpointsRequest {
	return &ListEndpointsRequest{
		Page: page,
		Size: size,
	}
}

func NewListTagsByEndpointIDRequest(userID, projectID, endpointID string) *ListTagsByEndpointIDRequest {
	opt := new(ListTagsByEndpointIDRequest)
	opt.ID = endpointID
	opt.EndpointID = endpointID
	opt.ProjectID = projectID
	opt.SetPortalUserID(userID)
	return opt
}

func NewCreateTagsWithEndpointIDRequest(userID, projectID, endpointID string) *CreateTagsWithEndpointIDRequest {
	opt := new(CreateTagsWithEndpointIDRequest)
	opt.ResourceUuid = endpointID
	opt.EndpointID = endpointID
	opt.SystemTag = true
	opt.ProjectID = projectID
	opt.SetPortalUserID(userID)

	return opt
}

func NewDeleteTagOfEndpointRequest(userID, projectID, tagID string) *DeleteTagOfEndpointRequest {
	opt := new(DeleteTagOfEndpointRequest)
	opt.TagID = tagID
	opt.ProjectID = projectID
	opt.SetPortalUserID(userID)

	return opt
}

func NewUpdateTagValueOfEndpointRequest(userID, projectID, tagID, value string) *UpdateTagValueOfEndpointRequest {
	opt := new(UpdateTagValueOfEndpointRequest)
	opt.TagID = tagID
	opt.TagValue = value
	opt.ProjectID = projectID
	opt.SetPortalUserID(userID)

	return opt
}
