package v1

import (
	"strings"

	"github.com/dannyota/greennode-community-sdk/v2/greennode/services/common"
	// "strings"
)

type GlobalListenerProtocol string

const (
	GlobalListenerProtocolTCP GlobalListenerProtocol = "TCP"
)

var _ IListGlobalListenersRequest = &ListGlobalListenersRequest{}

type ListGlobalListenersRequest struct {
	common.UserAgent
	common.LoadBalancerCommon
}

func (s *ListGlobalListenersRequest) WithLoadBalancerID(lbID string) IListGlobalListenersRequest {
	s.LoadBalancerID = lbID
	return s
}

func (s *ListGlobalListenersRequest) AddUserAgent(agent ...string) IListGlobalListenersRequest {
	s.UserAgent.AddUserAgent(agent...)
	return s
}

func NewListGlobalListenersRequest(lbID string) IListGlobalListenersRequest {
	opts := &ListGlobalListenersRequest{}
	opts.LoadBalancerID = lbID
	return opts
}

// --------------------------------------------------

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

func (s *CreateGlobalListenerRequest) WithAllowedCidrs(cidrs ...string) ICreateGlobalListenerRequest {
	s.AllowedCidrs = strings.Join(cidrs, ",")
	return s
}

func (s *CreateGlobalListenerRequest) WithDescription(desc string) ICreateGlobalListenerRequest {
	s.Description = desc
	return s
}

func (s *CreateGlobalListenerRequest) WithHeaders(headers ...string) ICreateGlobalListenerRequest {
	s.Headers = headers
	return s
}

func (s *CreateGlobalListenerRequest) WithName(name string) ICreateGlobalListenerRequest {
	s.Name = name
	return s
}

func (s *CreateGlobalListenerRequest) WithPort(port int) ICreateGlobalListenerRequest {
	s.Port = port
	return s
}

func (s *CreateGlobalListenerRequest) WithProtocol(protocol GlobalListenerProtocol) ICreateGlobalListenerRequest {
	s.Protocol = protocol
	return s
}

func (s *CreateGlobalListenerRequest) WithTimeoutClient(toc int) ICreateGlobalListenerRequest {
	s.TimeoutClient = toc
	return s
}

func (s *CreateGlobalListenerRequest) WithTimeoutConnection(toc int) ICreateGlobalListenerRequest {
	s.TimeoutConnection = toc
	return s
}

func (s *CreateGlobalListenerRequest) WithTimeoutMember(tom int) ICreateGlobalListenerRequest {
	s.TimeoutMember = tom
	return s
}

func (s *CreateGlobalListenerRequest) WithGlobalPoolID(poolID string) ICreateGlobalListenerRequest {
	s.GlobalPoolID = poolID
	return s
}

func (s *CreateGlobalListenerRequest) WithLoadBalancerID(lbid string) ICreateGlobalListenerRequest {
	s.LoadBalancerID = lbid
	return s
}

func (s *CreateGlobalListenerRequest) ToRequestBody() interface{} {
	return s
}

func (s *CreateGlobalListenerRequest) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"allowedCidrs":      s.AllowedCidrs,
		"description":       s.Description,
		"headers":           s.Headers,
		"name":              s.Name,
		"port":              s.Port,
		"protocol":          s.Protocol,
		"timeoutClient":     s.TimeoutClient,
		"timeoutConnection": s.TimeoutConnection,
		"timeoutMember":     s.TimeoutMember,
		"globalPoolId":      s.GlobalPoolID,
	}
}

func (s *CreateGlobalListenerRequest) AddUserAgent(agent ...string) ICreateGlobalListenerRequest {
	s.UserAgent.AddUserAgent(agent...)
	return s
}

func NewCreateGlobalListenerRequest(lbID, name string) ICreateGlobalListenerRequest {
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

// --------------------------------------------------

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

func (s *UpdateGlobalListenerRequest) WithAllowedCidrs(cidrs ...string) IUpdateGlobalListenerRequest {
	s.AllowedCidrs = strings.Join(cidrs, ",")
	return s
}

func (s *UpdateGlobalListenerRequest) WithTimeoutClient(toc int) IUpdateGlobalListenerRequest {
	s.TimeoutClient = toc
	return s
}

func (s *UpdateGlobalListenerRequest) WithTimeoutMember(tom int) IUpdateGlobalListenerRequest {
	s.TimeoutMember = tom
	return s
}

func (s *UpdateGlobalListenerRequest) WithTimeoutConnection(toc int) IUpdateGlobalListenerRequest {
	s.TimeoutConnection = toc
	return s
}

func (s *UpdateGlobalListenerRequest) WithHeaders(headers ...string) IUpdateGlobalListenerRequest {
	h := strings.Join(headers, ",")
	s.Headers = &h
	return s
}

func (s *UpdateGlobalListenerRequest) WithGlobalPoolID(poolID string) IUpdateGlobalListenerRequest {
	s.GlobalPoolID = poolID
	return s
}

func (s *UpdateGlobalListenerRequest) WithLoadBalancerID(lbid string) IUpdateGlobalListenerRequest {
	s.LoadBalancerID = lbid
	return s
}

func (s *UpdateGlobalListenerRequest) WithListenerID(lid string) IUpdateGlobalListenerRequest {
	s.ListenerID = lid
	return s
}

func (s *UpdateGlobalListenerRequest) ToRequestBody() interface{} {
	return s
}

func (s *UpdateGlobalListenerRequest) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"allowedCidrs":      s.AllowedCidrs,
		"timeoutClient":     s.TimeoutClient,
		"timeoutMember":     s.TimeoutMember,
		"timeoutConnection": s.TimeoutConnection,
		"headers":           s.Headers,
		"globalPoolId":      s.GlobalPoolID,
	}
}

func (s *UpdateGlobalListenerRequest) AddUserAgent(agent ...string) IUpdateGlobalListenerRequest {
	s.UserAgent.AddUserAgent(agent...)
	return s
}

func NewUpdateGlobalListenerRequest(lbID, lID string) IUpdateGlobalListenerRequest {
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

// --------------------------------------------------

var _ IDeleteGlobalListenerRequest = &DeleteGlobalListenerRequest{}

type DeleteGlobalListenerRequest struct {
	common.UserAgent
	common.LoadBalancerCommon
	common.ListenerCommon
}

func (s *DeleteGlobalListenerRequest) WithLoadBalancerID(lbid string) IDeleteGlobalListenerRequest {
	s.LoadBalancerID = lbid
	return s
}

func (s *DeleteGlobalListenerRequest) WithListenerID(lid string) IDeleteGlobalListenerRequest {
	s.ListenerID = lid
	return s
}

func (s *DeleteGlobalListenerRequest) AddUserAgent(agent ...string) IDeleteGlobalListenerRequest {
	s.UserAgent.AddUserAgent(agent...)
	return s
}

func NewDeleteGlobalListenerRequest(lbID, lID string) IDeleteGlobalListenerRequest {
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

// --------------------------------------------------

var _ IGetGlobalListenerRequest = &GetGlobalListenerRequest{}

type GetGlobalListenerRequest struct {
	common.UserAgent
	common.LoadBalancerCommon
	common.ListenerCommon
}

func (s *GetGlobalListenerRequest) WithLoadBalancerID(lbid string) IGetGlobalListenerRequest {
	s.LoadBalancerID = lbid
	return s
}

func (s *GetGlobalListenerRequest) WithListenerID(lid string) IGetGlobalListenerRequest {
	s.ListenerID = lid
	return s
}

func (s *GetGlobalListenerRequest) AddUserAgent(agent ...string) IGetGlobalListenerRequest {
	s.UserAgent.AddUserAgent(agent...)
	return s
}

func NewGetGlobalListenerRequest(lbID, lID string) IGetGlobalListenerRequest {
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
