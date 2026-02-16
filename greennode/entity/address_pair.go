package entity

type AddressPair struct {
	ID                 string
	VirtualIPAddressID string
	VirtualSubnetID    string
	NetworkInterfaceIP string
	NetworkInterfaceID string
	CIDR               string
}

type ListAddressPairs struct {
	Items []*AddressPair
}

func (l ListAddressPairs) Len() int {
	return len(l.Items)
}

func (l ListAddressPairs) At(idx int) *AddressPair {
	if idx < 0 || idx >= len(l.Items) {
		return nil
	}

	return l.Items[idx]
}
