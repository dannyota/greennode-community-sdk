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
		LoadBalancerCommon: common.LoadBalancerCommon{
			LoadBalancerID: lbID,
		},
		PackageID: packageID,
	}
}

func NewListLoadBalancerPackagesRequest() *ListLoadBalancerPackagesRequest {
	return &ListLoadBalancerPackagesRequest{}
}

func NewGetLoadBalancerByIDRequest(lbID string) *GetLoadBalancerByIDRequest {
	opts := new(GetLoadBalancerByIDRequest)
	opts.LoadBalancerID = lbID
	return opts
}

func NewListLoadBalancersRequest(page, size int) *ListLoadBalancersRequest {
	opts := new(ListLoadBalancersRequest)
	opts.Page = page
	opts.Size = size
	return opts
}

func (r *ListLoadBalancersRequest) AddUserAgent(agent ...string) *ListLoadBalancersRequest {
	r.UserAgent.AddUserAgent(agent...)
	return r
}

func NewDeleteLoadBalancerByIDRequest(lbID string) *DeleteLoadBalancerByIDRequest {
	opts := new(DeleteLoadBalancerByIDRequest)
	opts.LoadBalancerID = lbID
	return opts
}

func NewScaleLoadBalancerRequest(lbID string) *ScaleLoadBalancerRequest {
	return &ScaleLoadBalancerRequest{
		LoadBalancerCommon: common.LoadBalancerCommon{
			LoadBalancerID: lbID,
		},
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

	common.UserAgent
}

type ResizeLoadBalancerRequest struct {
	PackageID string `json:"packageId"`
	common.UserAgent
	common.LoadBalancerCommon
}

type ListLoadBalancerPackagesRequest struct {
	common.UserAgent
	ZoneID common.Zone `q:"zoneId,beempty"`
}

type GetLoadBalancerByIDRequest struct {
	common.UserAgent
	common.LoadBalancerCommon
}

type ListLoadBalancersRequest struct {
	Name string
	Page int
	Size int

	Tags []common.Tag
	common.UserAgent
}

type DeleteLoadBalancerByIDRequest struct {
	common.UserAgent
	common.LoadBalancerCommon
}

func (r *DeleteLoadBalancerByIDRequest) AddUserAgent(agent ...string) *DeleteLoadBalancerByIDRequest {
	r.UserAgent.AddUserAgent(agent...)
	return r
}

type ResizeLoadBalancerByIDRequest struct {
	common.UserAgent
	common.LoadBalancerCommon

	PackageID string `json:"packageId"`
}

func (r *ResizeLoadBalancerByIDRequest) AddUserAgent(agent ...string) *ResizeLoadBalancerByIDRequest {
	r.UserAgent.AddUserAgent(agent...)
	return r
}

type ScaleLoadBalancerRequest struct {
	common.UserAgent
	common.LoadBalancerCommon

	Scaling    *ScalingConfig    `json:"scaling"`
	Networking *NetworkingConfig `json:"networking"`
}

func (r *ScaleLoadBalancerRequest) AddUserAgent(agent ...string) *ScaleLoadBalancerRequest {
	r.UserAgent.AddUserAgent(agent...)
	return r
}

func (r *ScaleLoadBalancerRequest) WithScaling(scaling *ScalingConfig) *ScaleLoadBalancerRequest {
	r.Scaling = scaling
	return r
}

func (r *ScaleLoadBalancerRequest) WithNetworking(networking *NetworkingConfig) *ScaleLoadBalancerRequest {
	r.Networking = networking
	return r
}

func (r *CreateLoadBalancerRequest) prepare() {
	if r.Pool != nil {
		r.Pool.prepare()
	}

	if r.Listener != nil {
		r.Listener.prepare()
	}
}

func (r *CreateLoadBalancerRequest) AddUserAgent(agent ...string) *CreateLoadBalancerRequest {
	r.UserAgent.AddUserAgent(agent...)
	return r
}
func (r *CreateLoadBalancerRequest) WithListener(listener *CreateListenerRequest) *CreateLoadBalancerRequest {
	r.Listener = listener
	return r
}

func (r *CreateLoadBalancerRequest) WithPool(pool *CreatePoolRequest) *CreateLoadBalancerRequest {
	r.Pool = pool
	return r
}

func (r *CreateLoadBalancerRequest) WithTags(tags ...string) *CreateLoadBalancerRequest {
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

func (r *CreateLoadBalancerRequest) WithScheme(scheme LoadBalancerScheme) *CreateLoadBalancerRequest {
	r.Scheme = scheme
	return r
}

func (r *CreateLoadBalancerRequest) WithAutoScalable(autoScalable bool) *CreateLoadBalancerRequest {
	r.AutoScalable = autoScalable
	return r
}

func (r *CreateLoadBalancerRequest) WithPackageID(packageID string) *CreateLoadBalancerRequest {
	r.PackageID = packageID
	return r
}

func (r *CreateLoadBalancerRequest) WithSubnetID(subnetID string) *CreateLoadBalancerRequest {
	r.SubnetID = subnetID
	return r
}

func (r *CreateLoadBalancerRequest) WithType(typeVal LoadBalancerType) *CreateLoadBalancerRequest {
	r.Type = typeVal
	return r
}

func (r *CreateLoadBalancerRequest) WithPoc(isPoc bool) *CreateLoadBalancerRequest {
	r.IsPoc = isPoc
	return r
}

func (r *CreateLoadBalancerRequest) WithZoneID(zoneID common.Zone) *CreateLoadBalancerRequest {
	r.ZoneID = &zoneID
	return r
}

func (r *ResizeLoadBalancerRequest) AddUserAgent(agent ...string) *ResizeLoadBalancerRequest {
	r.UserAgent.AddUserAgent(agent...)
	return r
}

func (r *ResizeLoadBalancerRequest) WithPackageID(packageID string) *ResizeLoadBalancerRequest {
	r.PackageID = packageID
	return r
}

func (r *ListLoadBalancerPackagesRequest) AddUserAgent(agent ...string) *ListLoadBalancerPackagesRequest {
	r.UserAgent.AddUserAgent(agent...)
	return r
}

func (r *ListLoadBalancerPackagesRequest) WithZoneID(zoneID common.Zone) *ListLoadBalancerPackagesRequest {
	r.ZoneID = zoneID
	return r
}

func (r *ListLoadBalancerPackagesRequest) GetZoneID() string {
	return string(r.ZoneID)
}

func (r *GetLoadBalancerByIDRequest) AddUserAgent(agent ...string) *GetLoadBalancerByIDRequest {
	r.UserAgent.AddUserAgent(agent...)
	return r
}

func (r *ListLoadBalancersRequest) WithName(name string) *ListLoadBalancersRequest {
	r.Name = name
	return r
}

func (r *ListLoadBalancersRequest) WithTags(tags ...string) *ListLoadBalancersRequest {
	if r.Tags == nil {
		r.Tags = make([]common.Tag, 0)
	}

	if len(tags)%2 != 0 {
		tags = append(tags, "")
	}

	for i := 0; i < len(tags); i += 2 {
		r.Tags = append(r.Tags, common.Tag{Key: tags[i], Value: tags[i+1]})
	}

	return r
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


