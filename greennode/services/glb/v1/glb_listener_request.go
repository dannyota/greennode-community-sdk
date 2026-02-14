package v1

import (
	"strings"

	"github.com/dannyota/greennode-community-sdk/v2/greennode/services/common"
	// "strings"
)

type IListGlobalListenersRequest interface {
	WithLoadBalancerID(lbID string) IListGlobalListenersRequest
	GetLoadBalancerID() string

	AddUserAgent(agent ...string) IListGlobalListenersRequest
	ParseUserAgent() string
}

type IGetGlobalListenerRequest interface {
	WithLoadBalancerID(lbID string) IGetGlobalListenerRequest
	WithListenerID(listenerID string) IGetGlobalListenerRequest
	GetLoadBalancerID() string
	GetListenerID() string

	AddUserAgent(agent ...string) IGetGlobalListenerRequest
	ParseUserAgent() string
}

type ICreateGlobalListenerRequest interface {
	WithAllowedCidrs(cidrs ...string) ICreateGlobalListenerRequest
	WithDescription(desc string) ICreateGlobalListenerRequest
	WithHeaders(headers ...string) ICreateGlobalListenerRequest
	WithName(name string) ICreateGlobalListenerRequest
	WithPort(port int) ICreateGlobalListenerRequest
	WithProtocol(protocol GlobalListenerProtocol) ICreateGlobalListenerRequest
	WithTimeoutClient(toc int) ICreateGlobalListenerRequest
	WithTimeoutConnection(toc int) ICreateGlobalListenerRequest
	WithTimeoutMember(tom int) ICreateGlobalListenerRequest
	WithGlobalPoolID(poolID string) ICreateGlobalListenerRequest

	WithLoadBalancerID(lbid string) ICreateGlobalListenerRequest
	GetLoadBalancerID() string

	AddUserAgent(agent ...string) ICreateGlobalListenerRequest
	ParseUserAgent() string
	ToRequestBody() any
	ToMap() map[string]any
}

type IUpdateGlobalListenerRequest interface {
	WithAllowedCidrs(cidrs ...string) IUpdateGlobalListenerRequest
	WithTimeoutClient(toc int) IUpdateGlobalListenerRequest
	WithTimeoutMember(tom int) IUpdateGlobalListenerRequest
	WithTimeoutConnection(toc int) IUpdateGlobalListenerRequest
	WithHeaders(headers ...string) IUpdateGlobalListenerRequest
	WithGlobalPoolID(poolID string) IUpdateGlobalListenerRequest

	WithLoadBalancerID(lbID string) IUpdateGlobalListenerRequest
	WithListenerID(listenerID string) IUpdateGlobalListenerRequest
	GetLoadBalancerID() string
	GetListenerID() string

	AddUserAgent(agent ...string) IUpdateGlobalListenerRequest
	ParseUserAgent() string
	ToRequestBody() any
	ToMap() map[string]any
}

type IDeleteGlobalListenerRequest interface {
	WithLoadBalancerID(lbID string) IDeleteGlobalListenerRequest
	WithListenerID(listenerID string) IDeleteGlobalListenerRequest
	GetLoadBalancerID() string
	GetListenerID() string

	AddUserAgent(agent ...string) IDeleteGlobalListenerRequest
	ParseUserAgent() string
}

type GlobalListenerProtocol string

const (
	GlobalListenerProtocolTCP GlobalListenerProtocol = "TCP"
)

var _ IListGlobalListenersRequest = &ListGlobalListenersRequest{}

type ListGlobalListenersRequest struct {
	common.UserAgent
	common.LoadBalancerCommon
}

func (r *ListGlobalListenersRequest) WithLoadBalancerID(lbID string) IListGlobalListenersRequest {
	r.LoadBalancerID = lbID
	return r
}

func (r *ListGlobalListenersRequest) AddUserAgent(agent ...string) IListGlobalListenersRequest {
	r.UserAgent.AddUserAgent(agent...)
	return r
}

func NewListGlobalListenersRequest(lbID string) *ListGlobalListenersRequest {
	opts := &ListGlobalListenersRequest{}
	opts.LoadBalancerID = lbID
	return opts
}


var _ ICreateGlobalListenerRequest = &CreateGlobalListenerRequest{}

// WithAllowedCidrs(pcidrs ...string) ICreateGlobalListenerRequest
// WithDescription(pdesc string) ICreateGlobalListenerRequest
// WithHeaders(pheaders ...string) ICreateGlobalListenerRequest
// WithName(pname string) ICreateGlobalListenerRequest
// WithPort(pport int) ICreateGlobalListenerRequest
// WithProtocol(pprotocol GlobalListenerProtocol) ICreateGlobalListenerRequest
// WithTimeoutClient(ptoc int) ICreateGlobalListenerRequest
// WithTimeoutConnection(ptoc int) ICreateGlobalListenerRequest
// WithTimeoutMember(ptom int) ICreateGlobalListenerRequest
// WithDefaultPoolId(ppoolId string) ICreateGlobalListenerRequest
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

	common.UserAgent
	common.LoadBalancerCommon
}

func (r *CreateGlobalListenerRequest) WithAllowedCidrs(cidrs ...string) ICreateGlobalListenerRequest {
	r.AllowedCidrs = strings.Join(cidrs, ",")
	return r
}

func (r *CreateGlobalListenerRequest) WithDescription(desc string) ICreateGlobalListenerRequest {
	r.Description = desc
	return r
}

func (r *CreateGlobalListenerRequest) WithHeaders(headers ...string) ICreateGlobalListenerRequest {
	r.Headers = headers
	return r
}

func (r *CreateGlobalListenerRequest) WithName(name string) ICreateGlobalListenerRequest {
	r.Name = name
	return r
}

func (r *CreateGlobalListenerRequest) WithPort(port int) ICreateGlobalListenerRequest {
	r.Port = port
	return r
}

func (r *CreateGlobalListenerRequest) WithProtocol(protocol GlobalListenerProtocol) ICreateGlobalListenerRequest {
	r.Protocol = protocol
	return r
}

func (r *CreateGlobalListenerRequest) WithTimeoutClient(toc int) ICreateGlobalListenerRequest {
	r.TimeoutClient = toc
	return r
}

func (r *CreateGlobalListenerRequest) WithTimeoutConnection(toc int) ICreateGlobalListenerRequest {
	r.TimeoutConnection = toc
	return r
}

func (r *CreateGlobalListenerRequest) WithTimeoutMember(tom int) ICreateGlobalListenerRequest {
	r.TimeoutMember = tom
	return r
}

func (r *CreateGlobalListenerRequest) WithGlobalPoolID(poolID string) ICreateGlobalListenerRequest {
	r.GlobalPoolID = poolID
	return r
}

func (r *CreateGlobalListenerRequest) WithLoadBalancerID(lbid string) ICreateGlobalListenerRequest {
	r.LoadBalancerID = lbid
	return r
}

func (r *CreateGlobalListenerRequest) ToRequestBody() any {
	return r
}

func (r *CreateGlobalListenerRequest) ToMap() map[string]any {
	return map[string]any{
		"allowedCidrs":      r.AllowedCidrs,
		"description":       r.Description,
		"headers":           r.Headers,
		"name":              r.Name,
		"port":              r.Port,
		"protocol":          r.Protocol,
		"timeoutClient":     r.TimeoutClient,
		"timeoutConnection": r.TimeoutConnection,
		"timeoutMember":     r.TimeoutMember,
		"globalPoolId":      r.GlobalPoolID,
	}
}

func (r *CreateGlobalListenerRequest) AddUserAgent(agent ...string) ICreateGlobalListenerRequest {
	r.UserAgent.AddUserAgent(agent...)
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
		LoadBalancerCommon: common.LoadBalancerCommon{
			LoadBalancerID: lbID,
		},
	}
	return opts
}


var _ IUpdateGlobalListenerRequest = &UpdateGlobalListenerRequest{}

type UpdateGlobalListenerRequest struct {
	AllowedCidrs      string  `json:"allowedCidrs"`
	TimeoutClient     int     `json:"timeoutClient"`
	TimeoutMember     int     `json:"timeoutMember"`
	TimeoutConnection int     `json:"timeoutConnection"`
	Headers           *string `json:"headers"`
	GlobalPoolID      string  `json:"globalPoolId"`

	common.UserAgent
	common.LoadBalancerCommon
	common.ListenerCommon
}

func (r *UpdateGlobalListenerRequest) WithAllowedCidrs(cidrs ...string) IUpdateGlobalListenerRequest {
	r.AllowedCidrs = strings.Join(cidrs, ",")
	return r
}

func (r *UpdateGlobalListenerRequest) WithTimeoutClient(toc int) IUpdateGlobalListenerRequest {
	r.TimeoutClient = toc
	return r
}

func (r *UpdateGlobalListenerRequest) WithTimeoutMember(tom int) IUpdateGlobalListenerRequest {
	r.TimeoutMember = tom
	return r
}

func (r *UpdateGlobalListenerRequest) WithTimeoutConnection(toc int) IUpdateGlobalListenerRequest {
	r.TimeoutConnection = toc
	return r
}

func (r *UpdateGlobalListenerRequest) WithHeaders(headers ...string) IUpdateGlobalListenerRequest {
	h := strings.Join(headers, ",")
	r.Headers = &h
	return r
}

func (r *UpdateGlobalListenerRequest) WithGlobalPoolID(poolID string) IUpdateGlobalListenerRequest {
	r.GlobalPoolID = poolID
	return r
}

func (r *UpdateGlobalListenerRequest) WithLoadBalancerID(lbid string) IUpdateGlobalListenerRequest {
	r.LoadBalancerID = lbid
	return r
}

func (r *UpdateGlobalListenerRequest) WithListenerID(lid string) IUpdateGlobalListenerRequest {
	r.ListenerID = lid
	return r
}

func (r *UpdateGlobalListenerRequest) ToRequestBody() any {
	return r
}

func (r *UpdateGlobalListenerRequest) ToMap() map[string]any {
	return map[string]any{
		"allowedCidrs":      r.AllowedCidrs,
		"timeoutClient":     r.TimeoutClient,
		"timeoutMember":     r.TimeoutMember,
		"timeoutConnection": r.TimeoutConnection,
		"headers":           r.Headers,
		"globalPoolId":      r.GlobalPoolID,
	}
}

func (r *UpdateGlobalListenerRequest) AddUserAgent(agent ...string) IUpdateGlobalListenerRequest {
	r.UserAgent.AddUserAgent(agent...)
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
		LoadBalancerCommon: common.LoadBalancerCommon{
			LoadBalancerID: lbID,
		},
		ListenerCommon: common.ListenerCommon{
			ListenerID: lID,
		},
	}
	return opts
}


var _ IDeleteGlobalListenerRequest = &DeleteGlobalListenerRequest{}

type DeleteGlobalListenerRequest struct {
	common.UserAgent
	common.LoadBalancerCommon
	common.ListenerCommon
}

func (r *DeleteGlobalListenerRequest) WithLoadBalancerID(lbid string) IDeleteGlobalListenerRequest {
	r.LoadBalancerID = lbid
	return r
}

func (r *DeleteGlobalListenerRequest) WithListenerID(lid string) IDeleteGlobalListenerRequest {
	r.ListenerID = lid
	return r
}

func (r *DeleteGlobalListenerRequest) AddUserAgent(agent ...string) IDeleteGlobalListenerRequest {
	r.UserAgent.AddUserAgent(agent...)
	return r
}

func NewDeleteGlobalListenerRequest(lbID, lID string) *DeleteGlobalListenerRequest {
	opts := &DeleteGlobalListenerRequest{
		LoadBalancerCommon: common.LoadBalancerCommon{
			LoadBalancerID: lbID,
		},
		ListenerCommon: common.ListenerCommon{
			ListenerID: lID,
		},
	}
	return opts
}


var _ IGetGlobalListenerRequest = &GetGlobalListenerRequest{}

type GetGlobalListenerRequest struct {
	common.UserAgent
	common.LoadBalancerCommon
	common.ListenerCommon
}

func (r *GetGlobalListenerRequest) WithLoadBalancerID(lbid string) IGetGlobalListenerRequest {
	r.LoadBalancerID = lbid
	return r
}

func (r *GetGlobalListenerRequest) WithListenerID(lid string) IGetGlobalListenerRequest {
	r.ListenerID = lid
	return r
}

func (r *GetGlobalListenerRequest) AddUserAgent(agent ...string) IGetGlobalListenerRequest {
	r.UserAgent.AddUserAgent(agent...)
	return r
}

func NewGetGlobalListenerRequest(lbID, lID string) *GetGlobalListenerRequest {
	opts := &GetGlobalListenerRequest{
		LoadBalancerCommon: common.LoadBalancerCommon{
			LoadBalancerID: lbID,
		},
		ListenerCommon: common.ListenerCommon{
			ListenerID: lID,
		},
	}
	return opts
}
