package common

type NetworkCommon struct {
	NetworkID string
}

func (s *NetworkCommon) GetNetworkID() string {
	return s.NetworkID
}

type InternalNetworkInterfaceCommon struct {
	InternalNetworkInterfaceID string
}

func (s *InternalNetworkInterfaceCommon) GetInternalNetworkInterfaceID() string {
	return s.InternalNetworkInterfaceID
}

type WanCommon struct {
	WanID string
}

func (s *WanCommon) GetWanID() string {
	return s.WanID
}

type SecgroupCommon struct {
	SecgroupID string
}

func (s *SecgroupCommon) GetSecgroupID() string {
	return s.SecgroupID
}
