package v2

func NewGetAllAddressPairByVirtualSubnetIDRequest(subnetID string) *GetAllAddressPairByVirtualSubnetIDRequest {
	return &GetAllAddressPairByVirtualSubnetIDRequest{
		VirtualSubnetID: subnetID,
	}
}

type GetAllAddressPairByVirtualSubnetIDRequest struct {
	VirtualSubnetID string
}

func (r *GetAllAddressPairByVirtualSubnetIDRequest) GetVirtualSubnetID() string {
	return r.VirtualSubnetID
}


func NewSetAddressPairInVirtualSubnetRequest(subnetID, networkInterfaceID, CIDR string) *SetAddressPairInVirtualSubnetRequest {
	return &SetAddressPairInVirtualSubnetRequest{
		VirtualSubnetID: subnetID,
		AddressPairRequest: AddressPairRequest{
			CIDR:                       CIDR,
			InternalNetworkInterfaceID: networkInterfaceID,
		},
	}
}

type SetAddressPairInVirtualSubnetRequest struct {
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
	return &DeleteAddressPairRequest{
		AddressPairID: addressPairID,
	}
}

type DeleteAddressPairRequest struct {
	AddressPairID string
}

func (r *DeleteAddressPairRequest) GetAddressPairID() string {
	return r.AddressPairID
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

	VirtualAddressID string
}

func (r *CreateAddressPairRequest) WithMode(mode AddressPairMode) *CreateAddressPairRequest {
	r.Mode = &mode
	return r
}

func NewListAddressPairsByVirtualAddressIDRequest(virtualAddressID string) *ListAddressPairsByVirtualAddressIDRequest {
	return &ListAddressPairsByVirtualAddressIDRequest{
		VirtualAddressID: virtualAddressID,
	}
}

func NewCreateAddressPairRequest(virtualAddressID, internalNicID string) *CreateAddressPairRequest {
	return &CreateAddressPairRequest{
		InternalNetworkInterfaceID: internalNicID,
		VirtualAddressID:           virtualAddressID,
	}
}
