package v2

func NewCreateVirtualAddressCrossProjectRequest(name, projectId, subnetId string) ICreateVirtualAddressCrossProjectRequest {
	opts := new(CreateVirtualAddressCrossProjectRequest)
	opts.Name = name
	opts.CrossProjectRequest.ProjectId = projectId
	opts.CrossProjectRequest.SubnetId = subnetId
	return opts
}

func NewDeleteVirtualAddressByIdRequest(virtualAddressId string) IDeleteVirtualAddressByIdRequest {
	opts := new(DeleteVirtualAddressByIdRequest)
	opts.VirtualAddressId = virtualAddressId
	return opts
}

func NewGetVirtualAddressByIdRequest(virtualAddressId string) IGetVirtualAddressByIdRequest {
	opts := new(GetVirtualAddressByIdRequest)
	opts.VirtualAddressId = virtualAddressId
	return opts
}

func NewListAddressPairsByVirtualAddressIdRequest(virtualAddressId string) IListAddressPairsByVirtualAddressIdRequest {
	opts := new(ListAddressPairsByVirtualAddressIdRequest)
	opts.VirtualAddressId = virtualAddressId
	return opts
}

func NewCreateAddressPairRequest(virtualAddressId, internalNicId string) ICreateAddressPairRequest {
	opts := new(CreateAddressPairRequest)
	opts.VirtualAddressId = virtualAddressId
	opts.InternalNetworkInterfaceId = internalNicId
	return opts
}

func NewCreateSecgroupRequest(name, description string) ICreateSecgroupRequest {
	return &CreateSecgroupRequest{
		Name:        name,
		Description: description,
	}
}

func NewDeleteSecgroupByIdRequest(secgroupId string) IDeleteSecgroupByIdRequest {
	opts := new(DeleteSecgroupByIdRequest)
	opts.SecgroupId = secgroupId
	return opts
}

func NewGetSecgroupByIdRequest(secgroupId string) IGetSecgroupByIdRequest {
	opt := new(GetSecgroupByIdRequest)
	opt.SecgroupId = secgroupId
	return opt
}

func NewListSecgroupRequest() IListSecgroupRequest {
	return &ListSecgroupRequest{}
}
