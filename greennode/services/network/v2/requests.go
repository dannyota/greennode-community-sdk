package v2

func NewCreateVirtualAddressCrossProjectRequest(name, projectID, subnetID string) ICreateVirtualAddressCrossProjectRequest {
	opts := new(CreateVirtualAddressCrossProjectRequest)
	opts.Name = name
	opts.CrossProjectRequest.ProjectID = projectID
	opts.CrossProjectRequest.SubnetID = subnetID
	return opts
}

func NewDeleteVirtualAddressByIDRequest(virtualAddressID string) IDeleteVirtualAddressByIDRequest {
	opts := new(DeleteVirtualAddressByIDRequest)
	opts.VirtualAddressID = virtualAddressID
	return opts
}

func NewGetVirtualAddressByIDRequest(virtualAddressID string) IGetVirtualAddressByIDRequest {
	opts := new(GetVirtualAddressByIDRequest)
	opts.VirtualAddressID = virtualAddressID
	return opts
}

func NewListAddressPairsByVirtualAddressIDRequest(virtualAddressID string) IListAddressPairsByVirtualAddressIDRequest {
	opts := new(ListAddressPairsByVirtualAddressIDRequest)
	opts.VirtualAddressID = virtualAddressID
	return opts
}

func NewCreateAddressPairRequest(virtualAddressID, internalNicID string) ICreateAddressPairRequest {
	opts := new(CreateAddressPairRequest)
	opts.VirtualAddressID = virtualAddressID
	opts.InternalNetworkInterfaceID = internalNicID
	return opts
}

func NewCreateSecgroupRequest(name, description string) ICreateSecgroupRequest {
	return &CreateSecgroupRequest{
		Name:        name,
		Description: description,
	}
}

func NewDeleteSecgroupByIDRequest(secgroupID string) IDeleteSecgroupByIDRequest {
	opts := new(DeleteSecgroupByIDRequest)
	opts.SecgroupID = secgroupID
	return opts
}

func NewGetSecgroupByIDRequest(secgroupID string) IGetSecgroupByIDRequest {
	opt := new(GetSecgroupByIDRequest)
	opt.SecgroupID = secgroupID
	return opt
}

func NewListSecgroupRequest() IListSecgroupRequest {
	return &ListSecgroupRequest{}
}
