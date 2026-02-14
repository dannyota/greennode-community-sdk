package entity

type AddressPair struct {
	ID                 string
	VirtualIpAddressID string
	VirtualSubnetID    string
	NetworkInterfaceIp string
	NetworkInterfaceID string
	CIDR               string
}

func (ap *AddressPair) GetID() string {
	return ap.ID
}

type ListAddressPairs struct {
	Items []*AddressPair
}

func (l *ListAddressPairs) Len() int {
	return len(l.Items)
}

func (l *ListAddressPairs) At(idx int) *AddressPair {
	if idx < 0 || idx >= len(l.Items) {
		return nil
	}

	return l.Items[idx]
}
