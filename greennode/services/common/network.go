package common

type NetworkCommon struct {
	NetworkID string
}

func (n *NetworkCommon) GetNetworkID() string {
	return n.NetworkID
}

type InternalNetworkInterfaceCommon struct {
	InternalNetworkInterfaceID string
}

func (i *InternalNetworkInterfaceCommon) GetInternalNetworkInterfaceID() string {
	return i.InternalNetworkInterfaceID
}

type WanCommon struct {
	WanID string
}

func (w *WanCommon) GetWanID() string {
	return w.WanID
}

type SecgroupCommon struct {
	SecgroupID string
}

func (s *SecgroupCommon) GetSecgroupID() string {
	return s.SecgroupID
}
