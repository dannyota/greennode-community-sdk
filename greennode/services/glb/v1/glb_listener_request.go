package v1

import (
	"strings"
)

type GlobalListenerProtocol string

const (
	GlobalListenerProtocolTCP GlobalListenerProtocol = "TCP"
)

type ListGlobalListenersRequest struct {
	LoadBalancerID string
}

func (r *ListGlobalListenersRequest) WithLoadBalancerID(lbID string) *ListGlobalListenersRequest {
	r.LoadBalancerID = lbID
	return r
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

func (r *CreateGlobalListenerRequest) WithAllowedCidrs(cidrs ...string) *CreateGlobalListenerRequest {
	r.AllowedCidrs = strings.Join(cidrs, ",")
	return r
}

func (r *CreateGlobalListenerRequest) WithDescription(desc string) *CreateGlobalListenerRequest {
	r.Description = desc
	return r
}

func (r *CreateGlobalListenerRequest) WithHeaders(headers ...string) *CreateGlobalListenerRequest {
	r.Headers = headers
	return r
}

func (r *CreateGlobalListenerRequest) WithName(name string) *CreateGlobalListenerRequest {
	r.Name = name
	return r
}

func (r *CreateGlobalListenerRequest) WithPort(port int) *CreateGlobalListenerRequest {
	r.Port = port
	return r
}

func (r *CreateGlobalListenerRequest) WithProtocol(protocol GlobalListenerProtocol) *CreateGlobalListenerRequest {
	r.Protocol = protocol
	return r
}

func (r *CreateGlobalListenerRequest) WithTimeoutClient(toc int) *CreateGlobalListenerRequest {
	r.TimeoutClient = toc
	return r
}

func (r *CreateGlobalListenerRequest) WithTimeoutConnection(toc int) *CreateGlobalListenerRequest {
	r.TimeoutConnection = toc
	return r
}

func (r *CreateGlobalListenerRequest) WithTimeoutMember(tom int) *CreateGlobalListenerRequest {
	r.TimeoutMember = tom
	return r
}

func (r *CreateGlobalListenerRequest) WithGlobalPoolID(poolID string) *CreateGlobalListenerRequest {
	r.GlobalPoolID = poolID
	return r
}

func (r *CreateGlobalListenerRequest) WithLoadBalancerID(lbid string) *CreateGlobalListenerRequest {
	r.LoadBalancerID = lbid
	return r
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

func (r *UpdateGlobalListenerRequest) WithAllowedCidrs(cidrs ...string) *UpdateGlobalListenerRequest {
	r.AllowedCidrs = strings.Join(cidrs, ",")
	return r
}

func (r *UpdateGlobalListenerRequest) WithTimeoutClient(toc int) *UpdateGlobalListenerRequest {
	r.TimeoutClient = toc
	return r
}

func (r *UpdateGlobalListenerRequest) WithTimeoutMember(tom int) *UpdateGlobalListenerRequest {
	r.TimeoutMember = tom
	return r
}

func (r *UpdateGlobalListenerRequest) WithTimeoutConnection(toc int) *UpdateGlobalListenerRequest {
	r.TimeoutConnection = toc
	return r
}

func (r *UpdateGlobalListenerRequest) WithHeaders(headers ...string) *UpdateGlobalListenerRequest {
	h := strings.Join(headers, ",")
	r.Headers = &h
	return r
}

func (r *UpdateGlobalListenerRequest) WithGlobalPoolID(poolID string) *UpdateGlobalListenerRequest {
	r.GlobalPoolID = poolID
	return r
}

func (r *UpdateGlobalListenerRequest) WithLoadBalancerID(lbid string) *UpdateGlobalListenerRequest {
	r.LoadBalancerID = lbid
	return r
}

func (r *UpdateGlobalListenerRequest) WithListenerID(lid string) *UpdateGlobalListenerRequest {
	r.ListenerID = lid
	return r
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

func (r *DeleteGlobalListenerRequest) WithLoadBalancerID(lbid string) *DeleteGlobalListenerRequest {
	r.LoadBalancerID = lbid
	return r
}

func (r *DeleteGlobalListenerRequest) WithListenerID(lid string) *DeleteGlobalListenerRequest {
	r.ListenerID = lid
	return r
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

func (r *GetGlobalListenerRequest) WithLoadBalancerID(lbid string) *GetGlobalListenerRequest {
	r.LoadBalancerID = lbid
	return r
}

func (r *GetGlobalListenerRequest) WithListenerID(lid string) *GetGlobalListenerRequest {
	r.ListenerID = lid
	return r
}

func NewGetGlobalListenerRequest(lbID, lID string) *GetGlobalListenerRequest {
	opts := &GetGlobalListenerRequest{
		LoadBalancerID: lbID,
		ListenerID:     lID,
	}
	return opts
}
