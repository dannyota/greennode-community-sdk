package v2

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"

	"github.com/dannyota/greennode-community-sdk/v2/greennode/entity"
	"github.com/dannyota/greennode-community-sdk/v2/greennode/services/common"
)

type ICreateLoadBalancerRequest interface {
	ToRequestBody() any
	AddUserAgent(agent ...string) ICreateLoadBalancerRequest
	WithListener(listener ICreateListenerRequest) ICreateLoadBalancerRequest
	WithPool(pool ICreatePoolRequest) ICreateLoadBalancerRequest
	WithTags(tags ...string) ICreateLoadBalancerRequest
	WithScheme(scheme LoadBalancerScheme) ICreateLoadBalancerRequest
	WithAutoScalable(autoScalable bool) ICreateLoadBalancerRequest
	WithPackageID(packageID string) ICreateLoadBalancerRequest
	WithSubnetID(subnetID string) ICreateLoadBalancerRequest
	WithType(typeVal LoadBalancerType) ICreateLoadBalancerRequest
	WithPoc(poc bool) ICreateLoadBalancerRequest
	WithZoneID(zoneID common.Zone) ICreateLoadBalancerRequest
	ParseUserAgent() string
	ToMap() map[string]any
}

type IResizeLoadBalancerRequest interface {
	ToRequestBody() any
	AddUserAgent(agent ...string) IResizeLoadBalancerRequest
	WithPackageID(packageID string) IResizeLoadBalancerRequest
	ParseUserAgent() string

	GetLoadBalancerID() string
}

type IListLoadBalancerPackagesRequest interface {
	WithZoneID(zoneID common.Zone) IListLoadBalancerPackagesRequest
	GetZoneID() string
	AddUserAgent(agent ...string) IListLoadBalancerPackagesRequest
	ParseUserAgent() string
	ToMap() map[string]any
}

type IGetLoadBalancerByIDRequest interface {
	AddUserAgent(agent ...string) IGetLoadBalancerByIDRequest
	ParseUserAgent() string
	GetLoadBalancerID() string
}

type IListLoadBalancersRequest interface {
	WithName(name string) IListLoadBalancersRequest
	WithTags(tags ...string) IListLoadBalancersRequest
	ToListQuery() (string, error)
	ParseUserAgent() string
	GetDefaultQuery() string
	AddUserAgent(agent ...string) IListLoadBalancersRequest
}

type IDeleteLoadBalancerByIDRequest interface {
	GetLoadBalancerID() string
	ParseUserAgent() string
	AddUserAgent(agent ...string) IDeleteLoadBalancerByIDRequest
}

type IResizeLoadBalancerByIDRequest interface {
	GetLoadBalancerID() string
	ToMap() map[string]any
	ParseUserAgent() string
	ToRequestBody() any
	AddUserAgent(agent ...string) IResizeLoadBalancerByIDRequest
}

type IScaleLoadBalancerRequest interface {
	GetLoadBalancerID() string
	ToMap() map[string]any
	ParseUserAgent() string
	ToRequestBody() any
	AddUserAgent(agent ...string) IScaleLoadBalancerRequest
	WithScaling(scaling *ScalingConfig) IScaleLoadBalancerRequest
	WithNetworking(networking *NetworkingConfig) IScaleLoadBalancerRequest
}

type IUpdateTagsRequest interface {
	GetLoadBalancerID() string
	ToRequestBody(lstTags *entity.ListTags) any
	ParseUserAgent() string
	WithTags(tags ...string) IUpdateTagsRequest
	ToMap() map[string]any
	AddUserAgent(agent ...string) IUpdateTagsRequest
}

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

func (r *ListLoadBalancersRequest) AddUserAgent(agent ...string) IListLoadBalancersRequest {
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
	Listener     ICreateListenerRequest `json:"listener"`
	Pool         ICreatePoolRequest     `json:"pool"`
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

func (r *DeleteLoadBalancerByIDRequest) AddUserAgent(agent ...string) IDeleteLoadBalancerByIDRequest {
	r.UserAgent.AddUserAgent(agent...)
	return r
}

type ResizeLoadBalancerByIDRequest struct {
	common.UserAgent
	common.LoadBalancerCommon

	PackageID string `json:"packageId"`
}

func (r *ResizeLoadBalancerByIDRequest) AddUserAgent(agent ...string) IResizeLoadBalancerByIDRequest {
	r.UserAgent.AddUserAgent(agent...)
	return r
}

type ScaleLoadBalancerRequest struct {
	common.UserAgent
	common.LoadBalancerCommon

	Scaling    *ScalingConfig    `json:"scaling"`
	Networking *NetworkingConfig `json:"networking"`
}

func (r *ScaleLoadBalancerRequest) AddUserAgent(agent ...string) IScaleLoadBalancerRequest {
	r.UserAgent.AddUserAgent(agent...)
	return r
}

func (r *ScaleLoadBalancerRequest) WithScaling(scaling *ScalingConfig) IScaleLoadBalancerRequest {
	r.Scaling = scaling
	return r
}

func (r *ScaleLoadBalancerRequest) WithNetworking(networking *NetworkingConfig) IScaleLoadBalancerRequest {
	r.Networking = networking
	return r
}

func (r *ScaleLoadBalancerRequest) ToRequestBody() any {
	return r
}

func (r *ScaleLoadBalancerRequest) ToMap() map[string]any {
	result := map[string]any{}
	if r.Scaling != nil {
		result["scaling"] = map[string]any{
			"minSize": r.Scaling.MinNodes,
			"maxSize": r.Scaling.MaxNodes,
		}
	}
	if r.Networking != nil {
		result["networking"] = map[string]any{
			"subnets": r.Networking.Subnets,
		}
	}
	return result
}

func (r *CreateLoadBalancerRequest) ToMap() map[string]any {
	err := map[string]any{
		"name":         r.Name,
		"packageId":    r.PackageID,
		"scheme":       r.Scheme,
		"autoScalable": r.AutoScalable,
		"subnetId":     r.SubnetID,
		"type":         r.Type,
		"tags":         r.Tags,
	}

	if r.Listener != nil {
		err["listener"] = r.Listener.ToMap()
	}

	if r.Pool != nil {
		err["pool"] = r.Pool.ToMap()
	}

	return err
}

func (r *CreateLoadBalancerRequest) ToRequestBody() any {
	if r.Pool != nil {
		r.Pool = r.Pool.ToRequestBody().(*CreatePoolRequest)
	}

	if r.Listener != nil {
		r.Listener = r.Listener.ToRequestBody().(*CreateListenerRequest)
	}

	return r
}

func (r *CreateLoadBalancerRequest) AddUserAgent(agent ...string) ICreateLoadBalancerRequest {
	r.UserAgent.AddUserAgent(agent...)
	return r
}
func (r *CreateLoadBalancerRequest) WithListener(listener ICreateListenerRequest) ICreateLoadBalancerRequest {
	r.Listener = listener
	return r
}

func (r *CreateLoadBalancerRequest) WithPool(pool ICreatePoolRequest) ICreateLoadBalancerRequest {
	r.Pool = pool
	return r
}

func (r *CreateLoadBalancerRequest) WithTags(tags ...string) ICreateLoadBalancerRequest {
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

func (r *CreateLoadBalancerRequest) WithScheme(scheme LoadBalancerScheme) ICreateLoadBalancerRequest {
	r.Scheme = scheme
	return r
}

func (r *CreateLoadBalancerRequest) WithAutoScalable(autoScalable bool) ICreateLoadBalancerRequest {
	r.AutoScalable = autoScalable
	return r
}

func (r *CreateLoadBalancerRequest) WithPackageID(packageID string) ICreateLoadBalancerRequest {
	r.PackageID = packageID
	return r
}

func (r *CreateLoadBalancerRequest) WithSubnetID(subnetID string) ICreateLoadBalancerRequest {
	r.SubnetID = subnetID
	return r
}

func (r *CreateLoadBalancerRequest) WithType(typeVal LoadBalancerType) ICreateLoadBalancerRequest {
	r.Type = typeVal
	return r
}

func (r *CreateLoadBalancerRequest) WithPoc(isPoc bool) ICreateLoadBalancerRequest {
	r.IsPoc = isPoc
	return r
}

func (r *CreateLoadBalancerRequest) WithZoneID(zoneID common.Zone) ICreateLoadBalancerRequest {
	r.ZoneID = &zoneID
	return r
}

func (r *ResizeLoadBalancerRequest) ToRequestBody() any {
	return r
}

func (r *ResizeLoadBalancerRequest) AddUserAgent(agent ...string) IResizeLoadBalancerRequest {
	r.UserAgent.AddUserAgent(agent...)
	return r
}

func (r *ResizeLoadBalancerRequest) WithPackageID(packageID string) IResizeLoadBalancerRequest {
	r.PackageID = packageID
	return r
}

func (r *ListLoadBalancerPackagesRequest) AddUserAgent(agent ...string) IListLoadBalancerPackagesRequest {
	r.UserAgent.AddUserAgent(agent...)
	return r
}

func (r *ListLoadBalancerPackagesRequest) WithZoneID(zoneID common.Zone) IListLoadBalancerPackagesRequest {
	r.ZoneID = zoneID
	return r
}

func (r *ListLoadBalancerPackagesRequest) GetZoneID() string {
	return string(r.ZoneID)
}

func (r *ListLoadBalancerPackagesRequest) ToMap() map[string]any {
	return map[string]any{}
}

func (r *GetLoadBalancerByIDRequest) AddUserAgent(agent ...string) IGetLoadBalancerByIDRequest {
	r.UserAgent.AddUserAgent(agent...)
	return r
}

func (r *ListLoadBalancersRequest) WithName(name string) IListLoadBalancersRequest {
	r.Name = name
	return r
}

func (r *ListLoadBalancersRequest) WithTags(tags ...string) IListLoadBalancersRequest {
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

func (r *ResizeLoadBalancerByIDRequest) ToMap() map[string]any {
	return map[string]any{
		"packageId":      r.PackageID,
		"loadBalancerId": r.LoadBalancerID,
	}
}

func (r *ResizeLoadBalancerByIDRequest) ToRequestBody() any {
	return r
}
