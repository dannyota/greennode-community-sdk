package common

type VirtualAddressCommon struct {
	VirtualAddressID string
}

func (s *VirtualAddressCommon) GetVirtualAddressID() string {
	return s.VirtualAddressID
}
