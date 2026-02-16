package v2

import "github.com/dannyota/greennode-community-sdk/v2/greennode/services/common"

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

func (r *CreateVirtualAddressCrossProjectRequest) WithDescription(description string) *CreateVirtualAddressCrossProjectRequest {
	r.Description = description
	return r
}

// API Delete virtual address by ID
type DeleteVirtualAddressByIDRequest struct {
	common.VirtualAddressCommon
}

// Api Get virtual address by ID

type GetVirtualAddressByIDRequest struct {
	common.VirtualAddressCommon
}

// Api list address pairs by virtual address ID

type ListAddressPairsByVirtualAddressIDRequest struct {
	common.VirtualAddressCommon
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
		VirtualAddressCommon: common.VirtualAddressCommon{
			VirtualAddressID: virtualAddressID,
		},
	}
}

func NewGetVirtualAddressByIDRequest(virtualAddressID string) *GetVirtualAddressByIDRequest {
	return &GetVirtualAddressByIDRequest{
		VirtualAddressCommon: common.VirtualAddressCommon{
			VirtualAddressID: virtualAddressID,
		},
	}
}
