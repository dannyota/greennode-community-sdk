package v2

import "github.com/dannyota/greennode-community-sdk/v2/greennode/services/common"

func NewGetAllAddressPairByVirtualSubnetIdRequest(subnetId string) IGetAllAddressPairByVirtualSubnetIdRequest {
	opt := new(GetAllAddressPairByVirtualSubnetIdRequest)
	opt.VirtualSubnetId = subnetId
	return opt
}

type GetAllAddressPairByVirtualSubnetIdRequest struct {
	common.UserAgent
	VirtualSubnetId string
}

func (s *GetAllAddressPairByVirtualSubnetIdRequest) GetVirtualSubnetId() string {
	return s.VirtualSubnetId
}

// --------------------------------------------------------

func NewSetAddressPairInVirtualSubnetRequest(subnetId, networkInterfaceID, CIDR string) ISetAddressPairInVirtualSubnetRequest {
	opt := new(SetAddressPairInVirtualSubnetRequest)
	opt.VirtualSubnetId = subnetId
	opt.AddressPairRequest = AddressPairRequest{
		CIDR:                       CIDR,
		InternalNetworkInterfaceId: networkInterfaceID,
	}
	return opt
}

type SetAddressPairInVirtualSubnetRequest struct {
	common.UserAgent
	VirtualSubnetId    string
	AddressPairRequest AddressPairRequest
}

func (s *SetAddressPairInVirtualSubnetRequest) GetVirtualSubnetId() string {
	return s.VirtualSubnetId
}

func (s *SetAddressPairInVirtualSubnetRequest) ToRequestBody() interface{} {
	return s.AddressPairRequest
}

type AddressPairRequest struct {
	CIDR                       string `json:"cidr"`
	InternalNetworkInterfaceId string `json:"internalNetworkInterfaceId"`
}

// --------------------------------------------------------

func NewDeleteAddressPairRequest(addressPairID string) IDeleteAddressPairRequest {
	opt := new(DeleteAddressPairRequest)
	opt.AddressPairID = addressPairID
	return opt
}

type DeleteAddressPairRequest struct {
	common.UserAgent
	AddressPairID string
}

func (s *DeleteAddressPairRequest) GetAddressPairID() string {
	return s.AddressPairID
}

func (s *DeleteAddressPairRequest) AddUserAgent(agent ...string) IDeleteAddressPairRequest {
	s.UserAgent.AddUserAgent(agent...)
	return s
}

// --------------------------------------------------------

// Api create address pair

type AddressPairMode string

const (
	AddressPairModeActiveActive AddressPairMode = "active-active"
)

type CreateAddressPairRequest struct {
	// Is the ID of the network interface that the address pair will be attached to.
	InternalNetworkInterfaceId string `json:"internalNetworkInterfaceId"` // required

	// Is the pair mode of the address pair.
	Mode *AddressPairMode `json:"mode,omitempty"`

	common.InternalNetworkInterfaceCommon
	common.UserAgent
	common.VirtualAddressCommon
}

func (s *CreateAddressPairRequest) ToRequestBody() interface{} {
	return s
}

func (s *CreateAddressPairRequest) ToMap() map[string]interface{} {
	mode := "active-standby"
	if s.Mode != nil {
		mode = string(*s.Mode)
	}

	return map[string]interface{}{
		"internalNetworkInterfaceId": s.InternalNetworkInterfaceId,
		"mode":                       mode,
	}
}

func (s *CreateAddressPairRequest) AddUserAgent(agent ...string) ICreateAddressPairRequest {
	s.UserAgent.AddUserAgent(agent...)
	return s
}

func (s *CreateAddressPairRequest) WithMode(mode AddressPairMode) ICreateAddressPairRequest {
	s.Mode = &mode
	return s
}
