package common

type SubnetCommon struct {
	SubnetID string
}

func (s *SubnetCommon) GetSubnetID() string {
	return s.SubnetID
}
