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

func (r *ListGlobalLoadBalancersRequest) WithName(name string) IListGlobalLoadBalancersRequest {
	r.Name = name
	return r
}

func (r *ListGlobalLoadBalancersRequest) WithTags(tags ...string) IListGlobalLoadBalancersRequest {
	if r.Tags == nil {
		r.Tags = make([]common.Tag, 0)
	}

	if len(tags)%2 != 0 {
		tags = append(tags, "none")
	}

	for i := 0; i < len(tags); i += 2 {
		r.Tags = append(r.Tags, common.Tag{Key: tags[i], Value: tags[i+1]})
	}

	return r
}

func (r *ListGlobalLoadBalancersRequest) ToListQuery() (string, error) {
	v := url.Values{}
	v.Set("name", r.Name)
	v.Set("offset", strconv.Itoa(r.Offset))
	v.Set("limit", strconv.Itoa(r.Limit))

	tuples := make([]string, 0, len(r.Tags))
	for _, tag := range r.Tags {
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

func (r *ListGlobalLoadBalancersRequest) GetDefaultQuery() string {
	return fmt.Sprintf("offset=%d&limit=%d", defaultOffsetListGlobalLoadBalancer, defaultLimitListGlobalLoadBalancer)
}

func (r *ListGlobalLoadBalancersRequest) AddUserAgent(agent ...string) IListGlobalLoadBalancersRequest {
	r.UserAgent.AddUserAgent(agent...)
	return r
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

func (r *CreateGlobalLoadBalancerRequest) WithDescription(desc string) ICreateGlobalLoadBalancerRequest {
	r.Description = desc
	return r
}

func (r *CreateGlobalLoadBalancerRequest) WithName(name string) ICreateGlobalLoadBalancerRequest {
	r.Name = name
	return r
}

func (r *CreateGlobalLoadBalancerRequest) WithType(typeVal GlobalLoadBalancerType) ICreateGlobalLoadBalancerRequest {
	r.Type = typeVal
	return r
}

func (r *CreateGlobalLoadBalancerRequest) WithPackage(packageID string) ICreateGlobalLoadBalancerRequest {
	r.Package = packageID
	return r
}

func (r *CreateGlobalLoadBalancerRequest) WithPaymentFlow(paymentFlow GlobalLoadBalancerPaymentFlow) ICreateGlobalLoadBalancerRequest {
	r.PaymentFlow = paymentFlow
	return r
}

func (r *CreateGlobalLoadBalancerRequest) WithGlobalListener(listener ICreateGlobalListenerRequest) ICreateGlobalLoadBalancerRequest {
	r.GlobalListener = listener
	return r
}

func (r *CreateGlobalLoadBalancerRequest) WithGlobalPool(pool ICreateGlobalPoolRequest) ICreateGlobalLoadBalancerRequest {
	r.GlobalPool = pool
	return r
}

func (r *CreateGlobalLoadBalancerRequest) ToRequestBody() interface{} {
	return r
}

func (r *CreateGlobalLoadBalancerRequest) AddUserAgent(agent ...string) ICreateGlobalLoadBalancerRequest {
	r.UserAgent.AddUserAgent(agent...)
	return r
}

func (r *CreateGlobalLoadBalancerRequest) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"description":    r.Description,
		"name":           r.Name,
		"type":           r.Type,
		"globalListener": r.GlobalListener,
		"globalPool":     r.GlobalPool,
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

func (r *DeleteGlobalLoadBalancerRequest) WithLoadBalancerID(lbID string) IDeleteGlobalLoadBalancerRequest {
	r.LoadBalancerID = lbID
	return r
}

func (r *DeleteGlobalLoadBalancerRequest) AddUserAgent(agent ...string) IDeleteGlobalLoadBalancerRequest {
	r.UserAgent.AddUserAgent(agent...)
	return r
}

func NewDeleteGlobalLoadBalancerRequest(lbID string) IDeleteGlobalLoadBalancerRequest {
	opts := &DeleteGlobalLoadBalancerRequest{
		LoadBalancerCommon: common.LoadBalancerCommon{
			LoadBalancerID: lbID,
		},
	}
	return opts
}

// --------------------------------------------------------------------------

var _ IListGlobalPackagesRequest = &ListGlobalPackagesRequest{}

type ListGlobalPackagesRequest struct {
	common.UserAgent
}

func (r *ListGlobalPackagesRequest) AddUserAgent(agent ...string) IListGlobalPackagesRequest {
	r.UserAgent.AddUserAgent(agent...)
	return r
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

func (r *ListGlobalRegionsRequest) AddUserAgent(agent ...string) IListGlobalRegionsRequest {
	r.UserAgent.AddUserAgent(agent...)
	return r
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

func (r *GetGlobalLoadBalancerUsageHistoriesRequest) WithLoadBalancerID(lbID string) IGetGlobalLoadBalancerUsageHistoriesRequest {
	r.LoadBalancerID = lbID
	return r
}

func (r *GetGlobalLoadBalancerUsageHistoriesRequest) WithFrom(from string) IGetGlobalLoadBalancerUsageHistoriesRequest {
	r.From = from
	return r
}

func (r *GetGlobalLoadBalancerUsageHistoriesRequest) WithTo(to string) IGetGlobalLoadBalancerUsageHistoriesRequest {
	r.To = to
	return r
}

func (r *GetGlobalLoadBalancerUsageHistoriesRequest) WithType(typeVal string) IGetGlobalLoadBalancerUsageHistoriesRequest {
	r.Type = typeVal
	return r
}

func (r *GetGlobalLoadBalancerUsageHistoriesRequest) ToListQuery() (string, error) {
	v := url.Values{}
	if r.From != "" {
		v.Set("from", r.From)
	}
	if r.To != "" {
		v.Set("to", r.To)
	}
	if r.Type != "" {
		v.Set("type", r.Type)
	}
	return v.Encode(), nil
}

func (r *GetGlobalLoadBalancerUsageHistoriesRequest) GetDefaultQuery() string {
	return ""
}

func (r *GetGlobalLoadBalancerUsageHistoriesRequest) AddUserAgent(agent ...string) IGetGlobalLoadBalancerUsageHistoriesRequest {
	r.UserAgent.AddUserAgent(agent...)
	return r
}

func NewGetGlobalLoadBalancerUsageHistoriesRequest(lbID, from, to, usageType string) IGetGlobalLoadBalancerUsageHistoriesRequest {
	opts := &GetGlobalLoadBalancerUsageHistoriesRequest{
		From: from,
		To:   to,
		Type: usageType,
		LoadBalancerCommon: common.LoadBalancerCommon{
			LoadBalancerID: lbID,
		},
	}
	return opts
}

// --------------------------------------------------------------------------

var _ IGetGlobalLoadBalancerByIDRequest = &GetGlobalLoadBalancerByIDRequest{}

type GetGlobalLoadBalancerByIDRequest struct {
	common.UserAgent
	common.LoadBalancerCommon
}

func (r *GetGlobalLoadBalancerByIDRequest) WithLoadBalancerID(lbID string) IGetGlobalLoadBalancerByIDRequest {
	r.LoadBalancerID = lbID
	return r
}

func (r *GetGlobalLoadBalancerByIDRequest) AddUserAgent(agent ...string) IGetGlobalLoadBalancerByIDRequest {
	r.UserAgent.AddUserAgent(agent...)
	return r
}

func NewGetGlobalLoadBalancerByIDRequest(lbID string) IGetGlobalLoadBalancerByIDRequest {
	opts := &GetGlobalLoadBalancerByIDRequest{
		LoadBalancerCommon: common.LoadBalancerCommon{
			LoadBalancerID: lbID,
		},
	}
	return opts
}
