package v2

func NewCreateVirtualAddressCrossProjectRequest(name, projectID, subnetID string) *CreateVirtualAddressCrossProjectRequest {
	opts := new(CreateVirtualAddressCrossProjectRequest)
	opts.Name = name
	opts.CrossProjectRequest.ProjectID = projectID
	opts.CrossProjectRequest.SubnetID = subnetID
	return opts
}

func NewDeleteVirtualAddressByIDRequest(virtualAddressID string) *DeleteVirtualAddressByIDRequest {
	opts := new(DeleteVirtualAddressByIDRequest)
	opts.VirtualAddressID = virtualAddressID
	return opts
}

func NewGetVirtualAddressByIDRequest(virtualAddressID string) *GetVirtualAddressByIDRequest {
	opts := new(GetVirtualAddressByIDRequest)
	opts.VirtualAddressID = virtualAddressID
	return opts
}

func NewListAddressPairsByVirtualAddressIDRequest(virtualAddressID string) *ListAddressPairsByVirtualAddressIDRequest {
	opts := new(ListAddressPairsByVirtualAddressIDRequest)
	opts.VirtualAddressID = virtualAddressID
	return opts
}

func NewCreateAddressPairRequest(virtualAddressID, internalNicID string) *CreateAddressPairRequest {
	opts := new(CreateAddressPairRequest)
	opts.VirtualAddressID = virtualAddressID
	opts.InternalNetworkInterfaceID = internalNicID
	return opts
}

func NewCreateSecgroupRequest(name, description string) *CreateSecgroupRequest {
	return &CreateSecgroupRequest{
		Name:        name,
		Description: description,
	}
}

func NewDeleteSecgroupByIDRequest(secgroupID string) *DeleteSecgroupByIDRequest {
	opts := new(DeleteSecgroupByIDRequest)
	opts.SecgroupID = secgroupID
	return opts
}

func NewGetSecgroupByIDRequest(secgroupID string) *GetSecgroupByIDRequest {
	opt := new(GetSecgroupByIDRequest)
	opt.SecgroupID = secgroupID
	return opt
}

func NewListSecgroupRequest() *ListSecgroupRequest {
	return &ListSecgroupRequest{}
}
