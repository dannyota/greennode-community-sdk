package v1

import (
	"fmt"
	"net/url"
	"strings"
	"strconv"

	"github.com/dannyota/greennode-community-sdk/v2/greennode/services/common"
)

type (
	GlobalLoadBalancerType        string
	GlobalLoadBalancerPaymentFlow string
)

const (
	GlobalLoadBalancerTypeLayer4           GlobalLoadBalancerType        = "Layer 4"
	GlobalLoadBalancerPaymentFlowAutomated GlobalLoadBalancerPaymentFlow = "automated"
)

// --------------------------------------------------------------------------

func NewListGlobalLoadBalancersRequest(offset, limit int) IListGlobalLoadBalancersRequest {
	opts := &ListGlobalLoadBalancersRequest{
		Name:   "",
		Offset: offset,
		Limit:  limit,
		Tags:   make([]common.Tag, 0),
	}
	return opts
}

type ListGlobalLoadBalancersRequest struct {
	Name   string
	Offset int
	Limit  int

	Tags []common.Tag
	common.UserAgent
}

func (s *ListGlobalLoadBalancersRequest) WithName(name string) IListGlobalLoadBalancersRequest {
	s.Name = name
	return s
}

func (s *ListGlobalLoadBalancersRequest) WithTags(tags ...string) IListGlobalLoadBalancersRequest {
	if s.Tags == nil {
		s.Tags = make([]common.Tag, 0)
	}

	if len(tags)%2 != 0 {
		tags = append(tags, "none")
	}

	for i := 0; i < len(tags); i += 2 {
		s.Tags = append(s.Tags, common.Tag{Key: tags[i], Value: tags[i+1]})
	}

	return s
}

func (s *ListGlobalLoadBalancersRequest) ToListQuery() (string, error) {
	v := url.Values{}
	v.Set("name", s.Name)
	v.Set("offset", strconv.Itoa(s.Offset))
	v.Set("limit", strconv.Itoa(s.Limit))

	tuples := make([]string, 0, len(s.Tags))
	for _, tag := range s.Tags {
		if tag.Key == "" {
			continue
		}

		tuple := "tags=key:" + tag.Key
		if tag.Value != "" {
			tuple += ",value:" + tag.Value
		}
		tuples = append(tuples, tuple)
	}

	if len(tuples) > 0 {
		return v.Encode() + "&" + strings.Join(tuples, "&"), nil
	}

	return v.Encode(), nil
}

func (s *ListGlobalLoadBalancersRequest) GetDefaultQuery() string {
	return fmt.Sprintf("offset=%d&limit=%d", defaultOffsetListGlobalLoadBalancer, defaultLimitListGlobalLoadBalancer)
}

func (s *ListGlobalLoadBalancersRequest) AddUserAgent(agent ...string) IListGlobalLoadBalancersRequest {
	s.UserAgent.AddUserAgent(agent...)
	return s
}

// --------------------------------------------------------------------------

var _ ICreateGlobalLoadBalancerRequest = &CreateGlobalLoadBalancerRequest{}

type CreateGlobalLoadBalancerRequest struct {
	Description    string                        `json:"description"`
	Name           string                        `json:"name"`
	Type           GlobalLoadBalancerType        `json:"type"`
	Package        string                        `json:"package"`
	PaymentFlow    GlobalLoadBalancerPaymentFlow `json:"paymentFlow"`
	GlobalListener ICreateGlobalListenerRequest  `json:"globalListener"`
	GlobalPool     ICreateGlobalPoolRequest      `json:"globalPool"`

	common.UserAgent
}

func (s *CreateGlobalLoadBalancerRequest) WithDescription(desc string) ICreateGlobalLoadBalancerRequest {
	s.Description = desc
	return s
}

func (s *CreateGlobalLoadBalancerRequest) WithName(name string) ICreateGlobalLoadBalancerRequest {
	s.Name = name
	return s
}

func (s *CreateGlobalLoadBalancerRequest) WithType(typeVal GlobalLoadBalancerType) ICreateGlobalLoadBalancerRequest {
	s.Type = typeVal
	return s
}

func (s *CreateGlobalLoadBalancerRequest) WithPackage(packageId string) ICreateGlobalLoadBalancerRequest {
	s.Package = packageId
	return s
}

func (s *CreateGlobalLoadBalancerRequest) WithPaymentFlow(paymentFlow GlobalLoadBalancerPaymentFlow) ICreateGlobalLoadBalancerRequest {
	s.PaymentFlow = paymentFlow
	return s
}

func (s *CreateGlobalLoadBalancerRequest) WithGlobalListener(listener ICreateGlobalListenerRequest) ICreateGlobalLoadBalancerRequest {
	s.GlobalListener = listener
	return s
}

func (s *CreateGlobalLoadBalancerRequest) WithGlobalPool(pool ICreateGlobalPoolRequest) ICreateGlobalLoadBalancerRequest {
	s.GlobalPool = pool
	return s
}

func (s *CreateGlobalLoadBalancerRequest) ToRequestBody() interface{} {
	return s
}

func (s *CreateGlobalLoadBalancerRequest) AddUserAgent(agent ...string) ICreateGlobalLoadBalancerRequest {
	s.UserAgent.AddUserAgent(agent...)
	return s
}

func (s *CreateGlobalLoadBalancerRequest) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"description":    s.Description,
		"name":           s.Name,
		"type":           s.Type,
		"globalListener": s.GlobalListener,
		"globalPool":     s.GlobalPool,
	}
}

func NewCreateGlobalLoadBalancerRequest(name string) ICreateGlobalLoadBalancerRequest {
	opts := &CreateGlobalLoadBalancerRequest{
		Description:    "",
		Name:           name,
		Type:           GlobalLoadBalancerTypeLayer4,
		Package:        "",
		PaymentFlow:    GlobalLoadBalancerPaymentFlowAutomated,
		GlobalListener: nil,
		GlobalPool:     nil,
	}
	return opts
}

// --------------------------------------------------------------------------

var _ IDeleteGlobalLoadBalancerRequest = &DeleteGlobalLoadBalancerRequest{}

type DeleteGlobalLoadBalancerRequest struct {
	common.UserAgent
	common.LoadBalancerCommon
}

func (s *DeleteGlobalLoadBalancerRequest) WithLoadBalancerId(lbId string) IDeleteGlobalLoadBalancerRequest {
	s.LoadBalancerId = lbId
	return s
}

func (s *DeleteGlobalLoadBalancerRequest) AddUserAgent(agent ...string) IDeleteGlobalLoadBalancerRequest {
	s.UserAgent.AddUserAgent(agent...)
	return s
}

func NewDeleteGlobalLoadBalancerRequest(lbId string) IDeleteGlobalLoadBalancerRequest {
	opts := &DeleteGlobalLoadBalancerRequest{
		LoadBalancerCommon: common.LoadBalancerCommon{
			LoadBalancerId: lbId,
		},
	}
	return opts
}

// --------------------------------------------------------------------------

var _ IListGlobalPackagesRequest = &ListGlobalPackagesRequest{}

type ListGlobalPackagesRequest struct {
	common.UserAgent
}

func (s *ListGlobalPackagesRequest) AddUserAgent(agent ...string) IListGlobalPackagesRequest {
	s.UserAgent.AddUserAgent(agent...)
	return s
}

func NewListGlobalPackagesRequest() IListGlobalPackagesRequest {
	opts := &ListGlobalPackagesRequest{}
	return opts
}

// --------------------------------------------------------------------------

var _ IListGlobalRegionsRequest = &ListGlobalRegionsRequest{}

type ListGlobalRegionsRequest struct {
	common.UserAgent
}

func (s *ListGlobalRegionsRequest) AddUserAgent(agent ...string) IListGlobalRegionsRequest {
	s.UserAgent.AddUserAgent(agent...)
	return s
}

func NewListGlobalRegionsRequest() IListGlobalRegionsRequest {
	opts := &ListGlobalRegionsRequest{}
	return opts
}

// --------------------------------------------------------------------------

var _ IGetGlobalLoadBalancerUsageHistoriesRequest = &GetGlobalLoadBalancerUsageHistoriesRequest{}

type GetGlobalLoadBalancerUsageHistoriesRequest struct {
	From string
	To   string
	Type string

	common.UserAgent
	common.LoadBalancerCommon
}

func (s *GetGlobalLoadBalancerUsageHistoriesRequest) WithLoadBalancerId(lbId string) IGetGlobalLoadBalancerUsageHistoriesRequest {
	s.LoadBalancerId = lbId
	return s
}

func (s *GetGlobalLoadBalancerUsageHistoriesRequest) WithFrom(from string) IGetGlobalLoadBalancerUsageHistoriesRequest {
	s.From = from
	return s
}

func (s *GetGlobalLoadBalancerUsageHistoriesRequest) WithTo(to string) IGetGlobalLoadBalancerUsageHistoriesRequest {
	s.To = to
	return s
}

func (s *GetGlobalLoadBalancerUsageHistoriesRequest) WithType(typeVal string) IGetGlobalLoadBalancerUsageHistoriesRequest {
	s.Type = typeVal
	return s
}

func (s *GetGlobalLoadBalancerUsageHistoriesRequest) ToListQuery() (string, error) {
	v := url.Values{}
	if s.From != "" {
		v.Set("from", s.From)
	}
	if s.To != "" {
		v.Set("to", s.To)
	}
	if s.Type != "" {
		v.Set("type", s.Type)
	}
	return v.Encode(), nil
}

func (s *GetGlobalLoadBalancerUsageHistoriesRequest) GetDefaultQuery() string {
	return ""
}

func (s *GetGlobalLoadBalancerUsageHistoriesRequest) AddUserAgent(agent ...string) IGetGlobalLoadBalancerUsageHistoriesRequest {
	s.UserAgent.AddUserAgent(agent...)
	return s
}

func NewGetGlobalLoadBalancerUsageHistoriesRequest(lbId, from, to, usageType string) IGetGlobalLoadBalancerUsageHistoriesRequest {
	opts := &GetGlobalLoadBalancerUsageHistoriesRequest{
		From: from,
		To:   to,
		Type: usageType,
		LoadBalancerCommon: common.LoadBalancerCommon{
			LoadBalancerId: lbId,
		},
	}
	return opts
}

// --------------------------------------------------------------------------

var _ IGetGlobalLoadBalancerByIdRequest = &GetGlobalLoadBalancerByIdRequest{}

type GetGlobalLoadBalancerByIdRequest struct {
	common.UserAgent
	common.LoadBalancerCommon
}

func (s *GetGlobalLoadBalancerByIdRequest) WithLoadBalancerId(lbId string) IGetGlobalLoadBalancerByIdRequest {
	s.LoadBalancerId = lbId
	return s
}

func (s *GetGlobalLoadBalancerByIdRequest) AddUserAgent(agent ...string) IGetGlobalLoadBalancerByIdRequest {
	s.UserAgent.AddUserAgent(agent...)
	return s
}

func NewGetGlobalLoadBalancerByIdRequest(lbId string) IGetGlobalLoadBalancerByIdRequest {
	opts := &GetGlobalLoadBalancerByIdRequest{
		LoadBalancerCommon: common.LoadBalancerCommon{
			LoadBalancerId: lbId,
		},
	}
	return opts
}
