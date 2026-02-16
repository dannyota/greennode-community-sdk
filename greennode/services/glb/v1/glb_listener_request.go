package v1

type GlobalListenerProtocol string

const (
	GlobalListenerProtocolTCP GlobalListenerProtocol = "TCP"
)

type ListGlobalListenersRequest struct {
	LoadBalancerID string
}

func NewListGlobalListenersRequest(lbID string) *ListGlobalListenersRequest {
	opts := &ListGlobalListenersRequest{}
	opts.LoadBalancerID = lbID
	return opts
}

type CreateGlobalListenerRequest struct {
	AllowedCidrs      string                 `json:"allowedCidrs"`
	Description       string                 `json:"description"`
	Headers           []string               `json:"headers"`
	Name              string                 `json:"name"`
	Port              int                    `json:"port"`
	Protocol          GlobalListenerProtocol `json:"protocol"`
	TimeoutClient     int                    `json:"timeoutClient"`
	TimeoutConnection int                    `json:"timeoutConnection"`
	TimeoutMember     int                    `json:"timeoutMember"`
	GlobalPoolID      string                 `json:"globalPoolId"`

	LoadBalancerID string
}

func NewCreateGlobalListenerRequest(lbID, name string) *CreateGlobalListenerRequest {
	opts := &CreateGlobalListenerRequest{
		AllowedCidrs:      "0.0.0.0/0",
		Description:       "",
		Headers:           nil,
		Name:              name,
		Port:              80,
		Protocol:          GlobalListenerProtocolTCP,
		TimeoutClient:     50,
		TimeoutConnection: 5,
		TimeoutMember:     50,
		GlobalPoolID:      "",
		LoadBalancerID:    lbID,
	}
	return opts
}

type UpdateGlobalListenerRequest struct {
	AllowedCidrs      string  `json:"allowedCidrs"`
	TimeoutClient     int     `json:"timeoutClient"`
	TimeoutMember     int     `json:"timeoutMember"`
	TimeoutConnection int     `json:"timeoutConnection"`
	Headers           *string `json:"headers"`
	GlobalPoolID      string  `json:"globalPoolId"`

	LoadBalancerID string
	ListenerID     string
}

func NewUpdateGlobalListenerRequest(lbID, lID string) *UpdateGlobalListenerRequest {
	opts := &UpdateGlobalListenerRequest{
		AllowedCidrs:      "0.0.0.0/0",
		TimeoutClient:     50,
		TimeoutMember:     50,
		TimeoutConnection: 5,
		Headers:           nil,
		GlobalPoolID:      "",
		LoadBalancerID:    lbID,
		ListenerID:        lID,
	}
	return opts
}

type DeleteGlobalListenerRequest struct {
	LoadBalancerID string
	ListenerID     string
}

func NewDeleteGlobalListenerRequest(lbID, lID string) *DeleteGlobalListenerRequest {
	opts := &DeleteGlobalListenerRequest{
		LoadBalancerID: lbID,
		ListenerID:     lID,
	}
	return opts
}

type GetGlobalListenerRequest struct {
	LoadBalancerID string
	ListenerID     string
}

func NewGetGlobalListenerRequest(lbID, lID string) *GetGlobalListenerRequest {
	opts := &GetGlobalListenerRequest{
		LoadBalancerID: lbID,
		ListenerID:     lID,
	}
	return opts
}
