package v1

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"

	"danny.vn/greennode/greennode/services/common"
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

func (r *ListGlobalLoadBalancersRequest) getDefaultQuery() string {
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
	LoadBalancerID string
}

func NewDeleteGlobalLoadBalancerRequest(lbID string) *DeleteGlobalLoadBalancerRequest {
	opts := &DeleteGlobalLoadBalancerRequest{
		LoadBalancerID: lbID,
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

	LoadBalancerID string
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

func (r *GetGlobalLoadBalancerUsageHistoriesRequest) getDefaultQuery() string {
	return ""
}

func NewGetGlobalLoadBalancerUsageHistoriesRequest(lbID, from, to, usageType string) *GetGlobalLoadBalancerUsageHistoriesRequest {
	opts := &GetGlobalLoadBalancerUsageHistoriesRequest{
		From:           from,
		To:             to,
		Type:           usageType,
		LoadBalancerID: lbID,
	}
	return opts
}

type GetGlobalLoadBalancerByIDRequest struct {
	LoadBalancerID string
}

func NewGetGlobalLoadBalancerByIDRequest(lbID string) *GetGlobalLoadBalancerByIDRequest {
	opts := &GetGlobalLoadBalancerByIDRequest{
		LoadBalancerID: lbID,
	}
	return opts
}
