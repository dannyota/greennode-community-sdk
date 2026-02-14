package common

type VirtualAddressCommon struct {
	VirtualAddressID string
}

func (v *VirtualAddressCommon) GetVirtualAddressID() string {
	return v.VirtualAddressID
}
