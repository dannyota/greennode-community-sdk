package common

type EndpointCommon struct {
	EndpointID string
}

func (e *EndpointCommon) GetEndpointID() string {
	return e.EndpointID
}
