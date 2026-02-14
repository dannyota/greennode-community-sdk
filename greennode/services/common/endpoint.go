package common

type EndpointCommon struct {
	EndpointID string
}

func (s *EndpointCommon) GetEndpointID() string {
	return s.EndpointID
}
