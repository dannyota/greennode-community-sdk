package v2

import "github.com/dannyota/greennode-community-sdk/v2/greennode/services/common"

/**
 * The group of Virtual Address APIs
 */

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
		IpAddress string `json:"ipAddress"`
	} `json:"crossProjectRequest"` // required

	// others...
	common.UserAgent
}

func (r *CreateVirtualAddressCrossProjectRequest) ToRequestBody() interface{} {
	return r
}

func (r *CreateVirtualAddressCrossProjectRequest) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"name":        r.Name,
		"description": r.Description,
		"crossProjectRequest": map[string]interface{}{
			"projectId": r.CrossProjectRequest.ProjectID,
			"subnetId":  r.CrossProjectRequest.SubnetID,
			"ipAddress": r.CrossProjectRequest.IpAddress,
		},
	}
}

func (r *CreateVirtualAddressCrossProjectRequest) AddUserAgent(agent ...string) ICreateVirtualAddressCrossProjectRequest {
	r.UserAgent.AddUserAgent(agent...)
	return r
}

func (r *CreateVirtualAddressCrossProjectRequest) WithDescription(description string) ICreateVirtualAddressCrossProjectRequest {
	r.Description = description
	return r
}

// API Delete virtual address by ID
type DeleteVirtualAddressByIDRequest struct {
	common.VirtualAddressCommon
	common.UserAgent
}

func (r *DeleteVirtualAddressByIDRequest) AddUserAgent(agent ...string) IDeleteVirtualAddressByIDRequest {
	r.UserAgent.AddUserAgent(agent...)
	return r
}

func (r *DeleteVirtualAddressByIDRequest) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"virtualAddressId": r.VirtualAddressID,
	}
}

// Api Get virtual address by ID

type GetVirtualAddressByIDRequest struct {
	common.VirtualAddressCommon
	common.UserAgent
}

func (r *GetVirtualAddressByIDRequest) AddUserAgent(agent ...string) IGetVirtualAddressByIDRequest {
	r.UserAgent.AddUserAgent(agent...)
	return r
}

func (r *GetVirtualAddressByIDRequest) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"virtualAddressId": r.VirtualAddressID,
	}
}

// Api list address pairs by virtual address ID

type ListAddressPairsByVirtualAddressIDRequest struct {
	common.VirtualAddressCommon
	common.UserAgent
}

func (r *ListAddressPairsByVirtualAddressIDRequest) AddUserAgent(agent ...string) IListAddressPairsByVirtualAddressIDRequest {
	r.UserAgent.AddUserAgent(agent...)
	return r
}

func (r *ListAddressPairsByVirtualAddressIDRequest) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"virtualAddressId": r.VirtualAddressID,
	}
}
