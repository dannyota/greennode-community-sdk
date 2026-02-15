package v2

import "github.com/dannyota/greennode-community-sdk/v2/greennode/services/common"

func NewGetAllAddressPairByVirtualSubnetIDRequest(subnetID string) *GetAllAddressPairByVirtualSubnetIDRequest {
	opt := new(GetAllAddressPairByVirtualSubnetIDRequest)
	opt.VirtualSubnetID = subnetID
	return opt
}

type GetAllAddressPairByVirtualSubnetIDRequest struct {
	common.UserAgent
	VirtualSubnetID string
}

func (r *GetAllAddressPairByVirtualSubnetIDRequest) GetVirtualSubnetID() string {
	return r.VirtualSubnetID
}


func NewSetAddressPairInVirtualSubnetRequest(subnetID, networkInterfaceID, CIDR string) *SetAddressPairInVirtualSubnetRequest {
	opt := new(SetAddressPairInVirtualSubnetRequest)
	opt.VirtualSubnetID = subnetID
	opt.AddressPairRequest = AddressPairRequest{
		CIDR:                       CIDR,
		InternalNetworkInterfaceID: networkInterfaceID,
	}
	return opt
}

type SetAddressPairInVirtualSubnetRequest struct {
	common.UserAgent
	VirtualSubnetID    string
	AddressPairRequest AddressPairRequest
}

func (r *SetAddressPairInVirtualSubnetRequest) GetVirtualSubnetID() string {
	return r.VirtualSubnetID
}

type AddressPairRequest struct {
	CIDR                       string `json:"cidr"`
	InternalNetworkInterfaceID string `json:"internalNetworkInterfaceId"`
}


func NewDeleteAddressPairRequest(addressPairID string) *DeleteAddressPairRequest {
	opt := new(DeleteAddressPairRequest)
	opt.AddressPairID = addressPairID
	return opt
}

type DeleteAddressPairRequest struct {
	common.UserAgent
	AddressPairID string
}

func (r *DeleteAddressPairRequest) GetAddressPairID() string {
	return r.AddressPairID
}

func (r *DeleteAddressPairRequest) AddUserAgent(agent ...string) *DeleteAddressPairRequest {
	r.UserAgent.AddUserAgent(agent...)
	return r
}


// Api create address pair

type AddressPairMode string

const (
	AddressPairModeActiveActive AddressPairMode = "active-active"
)

type CreateAddressPairRequest struct {
	// Is the ID of the network interface that the address pair will be attached to.
	InternalNetworkInterfaceID string `json:"internalNetworkInterfaceId"` // required

	// Is the pair mode of the address pair.
	Mode *AddressPairMode `json:"mode,omitempty"`

	common.InternalNetworkInterfaceCommon
	common.UserAgent
	common.VirtualAddressCommon
}

func (r *CreateAddressPairRequest) ToMap() map[string]any {
	mode := "active-standby"
	if r.Mode != nil {
		mode = string(*r.Mode)
	}

	return map[string]any{
		"internalNetworkInterfaceId": r.InternalNetworkInterfaceID,
		"mode":                       mode,
	}
}

func (r *CreateAddressPairRequest) AddUserAgent(agent ...string) *CreateAddressPairRequest {
	r.UserAgent.AddUserAgent(agent...)
	return r
}

func (r *CreateAddressPairRequest) WithMode(mode AddressPairMode) *CreateAddressPairRequest {
	r.Mode = &mode
	return r
}
