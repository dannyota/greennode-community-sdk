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

func (s *CreateVirtualAddressCrossProjectRequest) ToRequestBody() interface{} {
	return s
}

func (s *CreateVirtualAddressCrossProjectRequest) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"name":        s.Name,
		"description": s.Description,
		"crossProjectRequest": map[string]interface{}{
			"projectId": s.CrossProjectRequest.ProjectID,
			"subnetId":  s.CrossProjectRequest.SubnetID,
			"ipAddress": s.CrossProjectRequest.IpAddress,
		},
	}
}

func (s *CreateVirtualAddressCrossProjectRequest) AddUserAgent(agent ...string) ICreateVirtualAddressCrossProjectRequest {
	s.UserAgent.AddUserAgent(agent...)
	return s
}

func (s *CreateVirtualAddressCrossProjectRequest) WithDescription(description string) ICreateVirtualAddressCrossProjectRequest {
	s.Description = description
	return s
}

// API Delete virtual address by ID
type DeleteVirtualAddressByIDRequest struct {
	common.VirtualAddressCommon
	common.UserAgent
}

func (s *DeleteVirtualAddressByIDRequest) AddUserAgent(agent ...string) IDeleteVirtualAddressByIDRequest {
	s.UserAgent.AddUserAgent(agent...)
	return s
}

func (s *DeleteVirtualAddressByIDRequest) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"virtualAddressId": s.VirtualAddressID,
	}
}

// Api Get virtual address by ID

type GetVirtualAddressByIDRequest struct {
	common.VirtualAddressCommon
	common.UserAgent
}

func (s *GetVirtualAddressByIDRequest) AddUserAgent(agent ...string) IGetVirtualAddressByIDRequest {
	s.UserAgent.AddUserAgent(agent...)
	return s
}

func (s *GetVirtualAddressByIDRequest) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"virtualAddressId": s.VirtualAddressID,
	}
}

// Api list address pairs by virtual address ID

type ListAddressPairsByVirtualAddressIDRequest struct {
	common.VirtualAddressCommon
	common.UserAgent
}

func (s *ListAddressPairsByVirtualAddressIDRequest) AddUserAgent(agent ...string) IListAddressPairsByVirtualAddressIDRequest {
	s.UserAgent.AddUserAgent(agent...)
	return s
}

func (s *ListAddressPairsByVirtualAddressIDRequest) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"virtualAddressId": s.VirtualAddressID,
	}
}
