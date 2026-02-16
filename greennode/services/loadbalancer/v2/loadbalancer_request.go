package v2

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"

	"github.com/dannyota/greennode-community-sdk/v2/greennode/services/common"
)

const (
	InternalLoadBalancerScheme LoadBalancerScheme = "Internal"
	InternetLoadBalancerScheme LoadBalancerScheme = "Internet"
	InterVpcLoadBalancerScheme LoadBalancerScheme = "InterVPC"
)

const (
	LoadBalancerTypeLayer4 LoadBalancerType = "Layer 4"
	LoadBalancerTypeLayer7 LoadBalancerType = "Layer 7"
)

type (
	LoadBalancerScheme string
	LoadBalancerType   string
)

// ScalingConfig defines scaling policies (min/max nodes)
type ScalingConfig struct {
	MinNodes int `json:"minSize"`
	MaxNodes int `json:"maxSize"`
}

// NetworkingConfig defines networking topology (subnets)
type NetworkingConfig struct {
	Subnets []string `json:"subnets,omitempty"`
}

func NewCreateLoadBalancerRequest(name, packageID, subnetID string) *CreateLoadBalancerRequest {
	return &CreateLoadBalancerRequest{
		Name:      name,
		PackageID: packageID,
		Scheme:    InternetLoadBalancerScheme,
		SubnetID:  subnetID,
		Type:      LoadBalancerTypeLayer4,
	}
}

func NewResizeLoadBalancerRequest(lbID, packageID string) *ResizeLoadBalancerRequest {
	return &ResizeLoadBalancerRequest{
		LoadBalancerID: lbID,
		PackageID:      packageID,
	}
}

func NewListLoadBalancerPackagesRequest() *ListLoadBalancerPackagesRequest {
	return &ListLoadBalancerPackagesRequest{}
}

func NewGetLoadBalancerByIDRequest(lbID string) *GetLoadBalancerByIDRequest {
	return &GetLoadBalancerByIDRequest{
		LoadBalancerID: lbID,
	}
}

func NewListLoadBalancersRequest(page, size int) *ListLoadBalancersRequest {
	return &ListLoadBalancersRequest{
		Page: page,
		Size: size,
	}
}

func NewDeleteLoadBalancerByIDRequest(lbID string) *DeleteLoadBalancerByIDRequest {
	return &DeleteLoadBalancerByIDRequest{
		LoadBalancerID: lbID,
	}
}

func NewScaleLoadBalancerRequest(lbID string) *ScaleLoadBalancerRequest {
	return &ScaleLoadBalancerRequest{
		LoadBalancerID: lbID,
	}
}

type CreateLoadBalancerRequest struct {
	Name         string                 `json:"name"`
	PackageID    string                 `json:"packageId"`
	Scheme       LoadBalancerScheme     `json:"scheme"`
	AutoScalable bool                   `json:"autoScalable"`
	SubnetID     string                 `json:"subnetId"`
	Type         LoadBalancerType       `json:"type"`
	Listener     *CreateListenerRequest `json:"listener"`
	Pool         *CreatePoolRequest     `json:"pool"`
	Tags         []common.Tag           `json:"tags,omitempty"`
	IsPoc        bool                   `json:"isPoc"`
	ZoneID       *common.Zone           `json:"zoneId"`
}

type ResizeLoadBalancerRequest struct {
	PackageID      string `json:"packageId"`
	LoadBalancerID string
}

type ListLoadBalancerPackagesRequest struct {
	ZoneID common.Zone `q:"zoneId,beempty"`
}

type GetLoadBalancerByIDRequest struct {
	LoadBalancerID string
}

type ListLoadBalancersRequest struct {
	Name string
	Page int
	Size int

	Tags []common.Tag
}

type DeleteLoadBalancerByIDRequest struct {
	LoadBalancerID string
}

type ResizeLoadBalancerByIDRequest struct {
	LoadBalancerID string

	PackageID string `json:"packageId"`
}

type ScaleLoadBalancerRequest struct {
	LoadBalancerID string

	Scaling    *ScalingConfig    `json:"scaling"`
	Networking *NetworkingConfig `json:"networking"`
}

// normalizeForAPI delegates to Pool and Listener normalization, clearing
// protocol-irrelevant fields before the API call. This mutates the receiver.
func (r *CreateLoadBalancerRequest) normalizeForAPI() {
	if r.Pool != nil {
		r.Pool.normalizeForAPI()
	}

	if r.Listener != nil {
		r.Listener.normalizeForAPI()
	}
}

func (r *ListLoadBalancersRequest) ToListQuery() (string, error) {
	v := url.Values{}
	v.Set("name", r.Name)
	v.Set("page", strconv.Itoa(r.Page))
	v.Set("size", strconv.Itoa(r.Size))

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

func (r *ListLoadBalancersRequest) GetDefaultQuery() string {
	return fmt.Sprintf("name=&page=%d&size=%d", defaultPageListLoadBalancer, defaultSizeListLoadBalancer)
}


func NewResizeLoadBalancerByIDRequest(lbID, packageID string) *ResizeLoadBalancerByIDRequest {
	return &ResizeLoadBalancerByIDRequest{
		LoadBalancerID: lbID,
		PackageID:      packageID,
	}
}
