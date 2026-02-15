package v1

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"

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

func NewListGlobalLoadBalancersRequest(offset, limit int) *ListGlobalLoadBalancersRequest {
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
}

func (r *ListGlobalLoadBalancersRequest) WithName(name string) *ListGlobalLoadBalancersRequest {
	r.Name = name
	return r
}

func (r *ListGlobalLoadBalancersRequest) WithTags(tags ...string) *ListGlobalLoadBalancersRequest {
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

type CreateGlobalLoadBalancerRequest struct {
	Description    string                        `json:"description"`
	Name           string                        `json:"name"`
	Type           GlobalLoadBalancerType        `json:"type"`
	Package        string                        `json:"package"`
	PaymentFlow    GlobalLoadBalancerPaymentFlow `json:"paymentFlow"`
	GlobalListener *CreateGlobalListenerRequest  `json:"globalListener"`
	GlobalPool     *CreateGlobalPoolRequest      `json:"globalPool"`
}

func (r *CreateGlobalLoadBalancerRequest) WithDescription(desc string) *CreateGlobalLoadBalancerRequest {
	r.Description = desc
	return r
}

func (r *CreateGlobalLoadBalancerRequest) WithName(name string) *CreateGlobalLoadBalancerRequest {
	r.Name = name
	return r
}

func (r *CreateGlobalLoadBalancerRequest) WithType(typeVal GlobalLoadBalancerType) *CreateGlobalLoadBalancerRequest {
	r.Type = typeVal
	return r
}

func (r *CreateGlobalLoadBalancerRequest) WithPackage(packageID string) *CreateGlobalLoadBalancerRequest {
	r.Package = packageID
	return r
}

func (r *CreateGlobalLoadBalancerRequest) WithPaymentFlow(paymentFlow GlobalLoadBalancerPaymentFlow) *CreateGlobalLoadBalancerRequest {
	r.PaymentFlow = paymentFlow
	return r
}

func (r *CreateGlobalLoadBalancerRequest) WithGlobalListener(listener *CreateGlobalListenerRequest) *CreateGlobalLoadBalancerRequest {
	r.GlobalListener = listener
	return r
}

func (r *CreateGlobalLoadBalancerRequest) WithGlobalPool(pool *CreateGlobalPoolRequest) *CreateGlobalLoadBalancerRequest {
	r.GlobalPool = pool
	return r
}

func NewCreateGlobalLoadBalancerRequest(name string) *CreateGlobalLoadBalancerRequest {
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

type DeleteGlobalLoadBalancerRequest struct {
	common.LoadBalancerCommon
}

func (r *DeleteGlobalLoadBalancerRequest) WithLoadBalancerID(lbID string) *DeleteGlobalLoadBalancerRequest {
	r.LoadBalancerID = lbID
	return r
}

func NewDeleteGlobalLoadBalancerRequest(lbID string) *DeleteGlobalLoadBalancerRequest {
	opts := &DeleteGlobalLoadBalancerRequest{
		LoadBalancerCommon: common.LoadBalancerCommon{
			LoadBalancerID: lbID,
		},
	}
	return opts
}

type ListGlobalPackagesRequest struct{}

func NewListGlobalPackagesRequest() *ListGlobalPackagesRequest {
	opts := &ListGlobalPackagesRequest{}
	return opts
}

type ListGlobalRegionsRequest struct{}

func NewListGlobalRegionsRequest() *ListGlobalRegionsRequest {
	opts := &ListGlobalRegionsRequest{}
	return opts
}

type GetGlobalLoadBalancerUsageHistoriesRequest struct {
	From string
	To   string
	Type string

	common.LoadBalancerCommon
}

func (r *GetGlobalLoadBalancerUsageHistoriesRequest) WithLoadBalancerID(lbID string) *GetGlobalLoadBalancerUsageHistoriesRequest {
	r.LoadBalancerID = lbID
	return r
}

func (r *GetGlobalLoadBalancerUsageHistoriesRequest) WithFrom(from string) *GetGlobalLoadBalancerUsageHistoriesRequest {
	r.From = from
	return r
}

func (r *GetGlobalLoadBalancerUsageHistoriesRequest) WithTo(to string) *GetGlobalLoadBalancerUsageHistoriesRequest {
	r.To = to
	return r
}

func (r *GetGlobalLoadBalancerUsageHistoriesRequest) WithType(typeVal string) *GetGlobalLoadBalancerUsageHistoriesRequest {
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

func NewGetGlobalLoadBalancerUsageHistoriesRequest(lbID, from, to, usageType string) *GetGlobalLoadBalancerUsageHistoriesRequest {
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

type GetGlobalLoadBalancerByIDRequest struct {
	common.LoadBalancerCommon
}

func (r *GetGlobalLoadBalancerByIDRequest) WithLoadBalancerID(lbID string) *GetGlobalLoadBalancerByIDRequest {
	r.LoadBalancerID = lbID
	return r
}

func NewGetGlobalLoadBalancerByIDRequest(lbID string) *GetGlobalLoadBalancerByIDRequest {
	opts := &GetGlobalLoadBalancerByIDRequest{
		LoadBalancerCommon: common.LoadBalancerCommon{
			LoadBalancerID: lbID,
		},
	}
	return opts
}
