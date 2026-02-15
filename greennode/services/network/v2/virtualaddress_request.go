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

	// others...
	common.UserAgent
}

func (r *CreateVirtualAddressCrossProjectRequest) AddUserAgent(agent ...string) *CreateVirtualAddressCrossProjectRequest {
	r.UserAgent.AddUserAgent(agent...)
	return r
}

func (r *CreateVirtualAddressCrossProjectRequest) WithDescription(description string) *CreateVirtualAddressCrossProjectRequest {
	r.Description = description
	return r
}

// API Delete virtual address by ID
type DeleteVirtualAddressByIDRequest struct {
	common.VirtualAddressCommon
	common.UserAgent
}

func (r *DeleteVirtualAddressByIDRequest) AddUserAgent(agent ...string) *DeleteVirtualAddressByIDRequest {
	r.UserAgent.AddUserAgent(agent...)
	return r
}

// Api Get virtual address by ID

type GetVirtualAddressByIDRequest struct {
	common.VirtualAddressCommon
	common.UserAgent
}

func (r *GetVirtualAddressByIDRequest) AddUserAgent(agent ...string) *GetVirtualAddressByIDRequest {
	r.UserAgent.AddUserAgent(agent...)
	return r
}

// Api list address pairs by virtual address ID

type ListAddressPairsByVirtualAddressIDRequest struct {
	common.VirtualAddressCommon
	common.UserAgent
}

func (r *ListAddressPairsByVirtualAddressIDRequest) AddUserAgent(agent ...string) *ListAddressPairsByVirtualAddressIDRequest {
	r.UserAgent.AddUserAgent(agent...)
	return r
}

