package v1

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/dannyota/greennode-community-sdk/v2/greennode/client"
	"github.com/dannyota/greennode-community-sdk/v2/greennode/services/common"
)

const (
	defaultPackageID = "d391c404-51b0-4f4d-a946-ec35c8a6e2be"

	VStorageServiceID = "b9ba2b16-389e-48b7-9e75-4c991239da27"
	VCRServiceID      = "4f823540-4d64-4bf5-a3e2-322b098b601b"
	VServerServiceID  = "e3f1b087-2d8b-4258-8b49-f5c6a4d27411"
	IAMServiceID      = "f3d11a4c-f071-4009-88a6-4a21346c8708"
	VMonitorServiceID = "9cc0d21a-5c27-4295-9787-eea213e255a1"

	defaultListEndpointsRequestPage = 1
	defaultListEndpointsRequestSize = 10
)

type GetEndpointByIDRequest struct {
	common.EndpointCommon
}

type CreateEndpointRequest struct {
	ResourceType string `json:"resourceType"`
	Action       string `json:"action"`
	ResourceInfo struct {
		IsBuyMorePoc      bool   `json:"isBuyMorePoc"`
		IsPoc             bool   `json:"isPoc"`
		IsEnableAutoRenew bool   `json:"isEnableAutoRenew"`
		EndpointName      string `json:"endpointName"`
		CategoryUuid      string `json:"categoryUuid"`
		ServiceUuid       string `json:"serviceUuid"`
		PackageUuid       string `json:"packageUuid"`
		VpcUuid           string `json:"vpcUuid"`
		PortalUserID      string `json:"portalUserId"`
		SubnetUuid        string `json:"subnetUuid"`
		RegionUuid        string `json:"regionUuid"`
		ProjectUuid       string `json:"projectUuid"`
		Description       string `json:"description"`
		EnableAZ          bool   `json:"enableAZ"`
		EnableDnsName     bool   `json:"enableDnsName"`
		Networking        []struct {
			Zone       string `json:"zone"`
			SubnetUuid string `json:"subnetUuid"`
		} `json:"networking"`
		Scaling struct {
			MinSize int `json:"minSize"`
			MaxSize int `json:"maxSize"`
		} `json:"scaling"`
	} `json:"resourceInfo"`
}

func (r *CreateEndpointRequest) ToRequestBody(svc client.ServiceClient) any {
	r.ResourceType = "endpoint"
	r.Action = "create"
	r.ResourceInfo.EnableAZ = true
	r.ResourceInfo.RegionUuid = svc.GetZoneID()
	r.ResourceInfo.ProjectUuid = svc.GetProjectID()

	return r
}

func (r *CreateEndpointRequest) WithEndpointName(endpointName string) *CreateEndpointRequest {
	r.ResourceInfo.EndpointName = endpointName
	return r
}

func (r *CreateEndpointRequest) WithCategoryUuid(categoryUuid string) *CreateEndpointRequest {
	r.ResourceInfo.CategoryUuid = categoryUuid
	return r
}

func (r *CreateEndpointRequest) WithServiceUuid(serviceUuid string) *CreateEndpointRequest {
	r.ResourceInfo.ServiceUuid = serviceUuid
	return r
}

func (r *CreateEndpointRequest) WithPackageUuid(packageUuid string) *CreateEndpointRequest {
	r.ResourceInfo.PackageUuid = packageUuid
	return r
}

func (r *CreateEndpointRequest) WithVpcUuid(vpcUuid string) *CreateEndpointRequest {
	r.ResourceInfo.VpcUuid = vpcUuid
	return r
}

func (r *CreateEndpointRequest) WithPortalUserID(portalUserID string) *CreateEndpointRequest {
	r.ResourceInfo.PortalUserID = portalUserID
	return r
}

func (r *CreateEndpointRequest) GetPortalUserID() string {
	return r.ResourceInfo.PortalUserID
}

func (r *CreateEndpointRequest) WithSubnetUuid(subnetUuid string) *CreateEndpointRequest {
	r.ResourceInfo.SubnetUuid = subnetUuid
	return r
}

func (r *CreateEndpointRequest) WithDescription(desp string) *CreateEndpointRequest {
	r.ResourceInfo.Description = desp
	return r
}

func (r *CreateEndpointRequest) WithPoc(yes bool) *CreateEndpointRequest {
	r.ResourceInfo.IsPoc = yes
	return r
}

func (r *CreateEndpointRequest) WithEnableDnsName(yes bool) *CreateEndpointRequest {
	r.ResourceInfo.EnableDnsName = yes
	return r
}

func (r *CreateEndpointRequest) WithBuyMorePoc(yes bool) *CreateEndpointRequest {
	r.ResourceInfo.IsBuyMorePoc = yes
	return r
}

func (r *CreateEndpointRequest) WithEnableAutoRenew(yes bool) *CreateEndpointRequest {
	r.ResourceInfo.IsEnableAutoRenew = yes
	return r
}

func (r *CreateEndpointRequest) AddNetworking(zone, subnetUuid string) *CreateEndpointRequest {
	r.ResourceInfo.Networking = append(r.ResourceInfo.Networking, struct {
		Zone       string `json:"zone"`
		SubnetUuid string `json:"subnetUuid"`
	}{
		Zone:       zone,
		SubnetUuid: subnetUuid,
	})
	return r
}

func (r *CreateEndpointRequest) WithScaling(minSize int, maxSize int) *CreateEndpointRequest {
	r.ResourceInfo.Scaling = struct {
		MinSize int `json:"minSize"`
		MaxSize int `json:"maxSize"`
	}{
		MinSize: minSize,
		MaxSize: maxSize,
	}
	return r
}

type DeleteEndpointByIDRequest struct {
	EndpointServiceUuid string `json:"endpointServiceUuid"`
	EndpointUuid        string `json:"endpointUuid"`
	ProjectUuid         string `json:"projectUuid"`
	RegionUuid          string `json:"regionUuid"`
	VpcUuid             string `json:"vpcUuid"`

	common.EndpointCommon
}

func (r *DeleteEndpointByIDRequest) ToRequestBody(svc client.ServiceClient) any {
	r.ProjectUuid = svc.GetProjectID()
	r.RegionUuid = svc.GetZoneID()

	return r
}

type ListEndpointsRequest struct {
	Page  int
	Size  int
	VpcID string
	Uuid  string
}

func (r *ListEndpointsRequest) WithPage(page int) *ListEndpointsRequest {
	r.Page = page
	return r
}

func (r *ListEndpointsRequest) WithSize(size int) *ListEndpointsRequest {
	r.Size = size
	return r
}

func (r *ListEndpointsRequest) WithVpcID(vpcID string) *ListEndpointsRequest {
	r.VpcID = vpcID
	return r
}

func (r *ListEndpointsRequest) WithUuid(uuid string) *ListEndpointsRequest {
	r.Uuid = uuid
	return r
}

func (r *ListEndpointsRequest) ToListQuery() (string, error) {
	var params []string
	if r.VpcID != "" {
		params = append(params, fmt.Sprintf(`{"field":"vpcId","value":"%s"}`, r.VpcID))
	}

	if r.Uuid != "" {
		params = append(params, fmt.Sprintf(`{"field":"uuid","value":"%s"}`, r.Uuid))
	}

	paramsFilter := strings.Join(params, ",")
	query := fmt.Sprintf(`{"page":%d,"size":%d,"search":[%s]}`, r.Page, r.Size, paramsFilter)
	query = "params=" + url.QueryEscape(query)

	return query, nil
}

func (r *ListEndpointsRequest) GetDefaultQuery() string {
	query := fmt.Sprintf(`{"page":%d,"size":%d}`, defaultListEndpointsRequestPage, defaultListEndpointsRequestSize)
	query = "params=" + url.QueryEscape(query)
	return query
}

// _____________________________________________________________________ ListTagsByEndpointIdRequest

type ListTagsByEndpointIDRequest struct {
	common.EndpointCommon
	common.PortalUser

	ProjectID string
	ID        string
}

func (r *ListTagsByEndpointIDRequest) ToListQuery() (string, error) {
	v := url.Values{}
	if r.ID != "" {
		v.Set("resourceUuid", r.ID)
	}
	return v.Encode(), nil
}

func (r *ListTagsByEndpointIDRequest) GetProjectID() string {
	return r.ProjectID
}

func (r *ListTagsByEndpointIDRequest) GetDefaultQuery() string {
	query := fmt.Sprintf(`{"page":%d,"size":%d}`, defaultListEndpointsRequestPage, defaultListEndpointsRequestSize)
	query = "params=" + url.QueryEscape(query)
	return query
}

func (r *ListTagsByEndpointIDRequest) GetMapHeaders() map[string]string {
	return r.PortalUser.GetMapHeaders()
}

// _________________________________________________________________ CreateTagsWithEndpointIdRequest

type CreateTagsWithEndpointIDRequest struct {
	common.EndpointCommon
	common.PortalUser

	ProjectID    string
	ResourceUuid string `json:"resourceUuid"`
	Tags         []struct {
		TagKey   string `json:"tagKey"`
		TagValue string `json:"tagValue"`
	} `json:"tags"`

	SystemTag bool `json:"systemTag"`
}

func (r *CreateTagsWithEndpointIDRequest) GetMapHeaders() map[string]string {
	return r.PortalUser.GetMapHeaders()
}

func (r *CreateTagsWithEndpointIDRequest) AddTag(key, value string) *CreateTagsWithEndpointIDRequest {
	r.Tags = append(r.Tags, struct {
		TagKey   string `json:"tagKey"`
		TagValue string `json:"tagValue"`
	}{
		TagKey:   key,
		TagValue: value,
	})

	return r
}

func (r *CreateTagsWithEndpointIDRequest) GetProjectID() string {
	return r.ProjectID
}

// ____________________________________________________________________ DeleteTagByEndpointIdRequest

type DeleteTagOfEndpointRequest struct {
	common.PortalUser

	ProjectID string
	TagID     string
}

func (r *DeleteTagOfEndpointRequest) GetMapHeaders() map[string]string {
	return r.PortalUser.GetMapHeaders()
}

func (r *DeleteTagOfEndpointRequest) GetTagID() string {
	return r.TagID
}

func (r *DeleteTagOfEndpointRequest) GetProjectID() string {
	return r.ProjectID
}

// _________________________________________________________________ UpdateTagValueOfEndpointRequest

type UpdateTagValueOfEndpointRequest struct {
	common.PortalUser

	TagID     string
	ProjectID string
	TagValue  string `json:"tagValue"`
}

func (r *UpdateTagValueOfEndpointRequest) GetMapHeaders() map[string]string {
	return r.PortalUser.GetMapHeaders()
}

func (r *UpdateTagValueOfEndpointRequest) GetTagID() string {
	return r.TagID
}

func (r *UpdateTagValueOfEndpointRequest) GetProjectID() string {
	return r.ProjectID
}
