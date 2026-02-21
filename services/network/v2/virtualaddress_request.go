package v2

// API Create virtual address cross project
type CreateVirtualAddressCrossProjectRequest struct {
	// Name is the name of the virtual address.
	Name string `json:"name"` // required

	// Description is the description of the virtual address.
	Description string `json:"description"`

	// Contains the required information to create a virtual address cross project.
	CrossProjectRequest struct {
		// The project ID whose virtual address will be created.
		ProjectID string `json:"projectId"` // required

		// The subnet ID where the virtual address will be created.
		SubnetID string `json:"subnetId"` // required

		// The IP address of the virtual address.
		IPAddress string `json:"ipAddress"`
	} `json:"crossProjectRequest"` // required
}

// API Delete virtual address by ID
type DeleteVirtualAddressByIDRequest struct {
	VirtualAddressID string
}

// Api Get virtual address by ID

type GetVirtualAddressByIDRequest struct {
	VirtualAddressID string
}

// Api list address pairs by virtual address ID

type ListAddressPairsByVirtualAddressIDRequest struct {
	VirtualAddressID string
}


func NewCreateVirtualAddressCrossProjectRequest(name, projectID, subnetID string) *CreateVirtualAddressCrossProjectRequest {
	opts := &CreateVirtualAddressCrossProjectRequest{
		Name: name,
	}
	opts.CrossProjectRequest.ProjectID = projectID
	opts.CrossProjectRequest.SubnetID = subnetID
	return opts
}

func NewDeleteVirtualAddressByIDRequest(virtualAddressID string) *DeleteVirtualAddressByIDRequest {
	return &DeleteVirtualAddressByIDRequest{
		VirtualAddressID: virtualAddressID,
	}
}

func NewGetVirtualAddressByIDRequest(virtualAddressID string) *GetVirtualAddressByIDRequest {
	return &GetVirtualAddressByIDRequest{
		VirtualAddressID: virtualAddressID,
	}
}
