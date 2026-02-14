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
	common.UserAgent
	common.EndpointCommon
}

func (r *GetEndpointByIDRequest) AddUserAgent(agent ...string) IGetEndpointByIDRequest {
	r.Agent = append(r.Agent, agent...)
	return r
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

	common.UserAgent
}

func (r *CreateEndpointRequest) ToMap() map[string]any {
	res := map[string]any{
		"resourceType":      r.ResourceType,
		"action":            r.Action,
		"isBuyMorePoc":      r.ResourceInfo.IsBuyMorePoc,
		"isPoc":             r.ResourceInfo.IsPoc,
		"isEnableAutoRenew": r.ResourceInfo.IsEnableAutoRenew,
		"endpointName":      r.ResourceInfo.EndpointName,
		"categoryID":        r.ResourceInfo.CategoryUuid,
		"serviceId":         r.ResourceInfo.ServiceUuid,
		"packageId":         r.ResourceInfo.PackageUuid,
		"vpcId":             r.ResourceInfo.VpcUuid,
		"subnetId":          r.ResourceInfo.SubnetUuid,
		"regionId":          r.ResourceInfo.RegionUuid,
		"projectId":         r.ResourceInfo.ProjectUuid,
		"description":       r.ResourceInfo.Description,
		"enableAZ":          r.ResourceInfo.EnableAZ,
		"enableDnsName":     r.ResourceInfo.EnableDnsName,
		"networking":        r.ResourceInfo.Networking,
		"scaling":           r.ResourceInfo.Scaling,
	}

	if len(r.Agent) > 0 {
		res["userAgent"] = r.Agent
	}

	return res
}

func (r *CreateEndpointRequest) AddUserAgent(agent ...string) ICreateEndpointRequest {
	r.Agent = append(r.Agent, agent...)
	return r
}

func (r *CreateEndpointRequest) ToRequestBody(svc client.ServiceClient) any {
	r.ResourceType = "endpoint"
	r.Action = "create"
	r.ResourceInfo.EnableAZ = true
	r.ResourceInfo.RegionUuid = svc.GetZoneID()
	r.ResourceInfo.ProjectUuid = svc.GetProjectID()

	return r
}

func (r *CreateEndpointRequest) WithEndpointName(endpointName string) ICreateEndpointRequest {
	r.ResourceInfo.EndpointName = endpointName
	return r
}

func (r *CreateEndpointRequest) WithCategoryUuid(categoryUuid string) ICreateEndpointRequest {
	r.ResourceInfo.CategoryUuid = categoryUuid
	return r
}

func (r *CreateEndpointRequest) WithServiceUuid(serviceUuid string) ICreateEndpointRequest {
	r.ResourceInfo.ServiceUuid = serviceUuid
	return r
}

func (r *CreateEndpointRequest) WithPackageUuid(packageUuid string) ICreateEndpointRequest {
	r.ResourceInfo.PackageUuid = packageUuid
	return r
}

func (r *CreateEndpointRequest) WithVpcUuid(vpcUuid string) ICreateEndpointRequest {
	r.ResourceInfo.VpcUuid = vpcUuid
	return r
}

func (r *CreateEndpointRequest) WithPortalUserID(portalUserID string) ICreateEndpointRequest {
	r.ResourceInfo.PortalUserID = portalUserID
	return r
}

func (r *CreateEndpointRequest) GetPortalUserID() string {
	return r.ResourceInfo.PortalUserID
}

func (r *CreateEndpointRequest) WithSubnetUuid(subnetUuid string) ICreateEndpointRequest {
	r.ResourceInfo.SubnetUuid = subnetUuid
	return r
}

func (r *CreateEndpointRequest) WithDescription(desp string) ICreateEndpointRequest {
	r.ResourceInfo.Description = desp
	return r
}

func (r *CreateEndpointRequest) WithPoc(yes bool) ICreateEndpointRequest {
	r.ResourceInfo.IsPoc = yes
	return r
}

func (r *CreateEndpointRequest) WithEnableDnsName(yes bool) ICreateEndpointRequest {
	r.ResourceInfo.EnableDnsName = yes
	return r
}

func (r *CreateEndpointRequest) WithBuyMorePoc(yes bool) ICreateEndpointRequest {
	r.ResourceInfo.IsBuyMorePoc = yes
	return r
}

func (r *CreateEndpointRequest) WithEnableAutoRenew(yes bool) ICreateEndpointRequest {
	r.ResourceInfo.IsEnableAutoRenew = yes
	return r
}

func (r *CreateEndpointRequest) AddNetworking(zone, subnetUuid string) ICreateEndpointRequest {
	r.ResourceInfo.Networking = append(r.ResourceInfo.Networking, struct {
		Zone       string `json:"zone"`
		SubnetUuid string `json:"subnetUuid"`
	}{
		Zone:       zone,
		SubnetUuid: subnetUuid,
	})
	return r
}

func (r *CreateEndpointRequest) WithScaling(minSize int, maxSize int) ICreateEndpointRequest {
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

	common.UserAgent
	common.EndpointCommon
}

func (r *DeleteEndpointByIDRequest) AddUserAgent(agent ...string) IDeleteEndpointByIDRequest {
	r.Agent = append(r.Agent, agent...)
	return r
}

func (r *DeleteEndpointByIDRequest) ToRequestBody(svc client.ServiceClient) any {
	r.ProjectUuid = svc.GetProjectID()
	r.RegionUuid = svc.GetZoneID()

	return r
}

func (r *DeleteEndpointByIDRequest) ToMap() map[string]any {
	res := map[string]any{
		"serviceId":  r.EndpointServiceUuid,
		"endpointId": r.EndpointID,
		"projectId":  r.ProjectUuid,
		"regionId":   r.RegionUuid,
		"vpcId":      r.VpcUuid,
	}

	if len(r.Agent) > 0 {
		res["userAgent"] = r.Agent
	}

	return res
}

type ListEndpointsRequest struct {
	Page  int
	Size  int
	VpcID string
	Uuid  string
	common.UserAgent
}

func (r *ListEndpointsRequest) ToMap() map[string]any {
	res := map[string]any{
		"page":  r.Page,
		"size":  r.Size,
		"vpcId": r.VpcID,
		"uuid":  r.Uuid,
	}

	if len(r.Agent) > 0 {
		res["userAgent"] = r.Agent
	}

	return res
}

func (r *ListEndpointsRequest) WithPage(page int) IListEndpointsRequest {
	r.Page = page
	return r
}

func (r *ListEndpointsRequest) WithSize(size int) IListEndpointsRequest {
	r.Size = size
	return r
}

func (r *ListEndpointsRequest) WithVpcID(vpcID string) IListEndpointsRequest {
	r.VpcID = vpcID
	return r
}

func (r *ListEndpointsRequest) WithUuid(uuid string) IListEndpointsRequest {
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

func (r *ListEndpointsRequest) AddUserAgent(agent ...string) IListEndpointsRequest {
	r.Agent = append(r.Agent, agent...)
	return r
}

// _____________________________________________________________________ ListTagsByEndpointIdRequest

type ListTagsByEndpointIDRequest struct {
	common.UserAgent
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

func (r *ListTagsByEndpointIDRequest) ToMap() map[string]any {
	res := map[string]any{
		"resourceUuid": r.ID,
	}

	if len(r.Agent) > 0 {
		res["userAgent"] = r.Agent
	}

	return res
}

func (r *ListTagsByEndpointIDRequest) GetMapHeaders() map[string]string {
	return r.PortalUser.GetMapHeaders()
}

func (r *ListTagsByEndpointIDRequest) AddUserAgent(agent ...string) IListTagsByEndpointIDRequest {
	r.Agent = append(r.Agent, agent...)
	return r
}

// _________________________________________________________________ CreateTagsWithEndpointIdRequest

type CreateTagsWithEndpointIDRequest struct {
	common.UserAgent
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

func (r *CreateTagsWithEndpointIDRequest) ToMap() map[string]any {
	res := map[string]any{
		"resourceUuid": r.ID,
	}

	if len(r.Agent) > 0 {
		res["userAgent"] = r.Agent
	}

	res["tags"] = r.Tags

	return res
}

func (r *CreateTagsWithEndpointIDRequest) AddUserAgent(agent ...string) ICreateTagsWithEndpointIDRequest {
	r.Agent = append(r.Agent, agent...)
	return r
}

func (r *CreateTagsWithEndpointIDRequest) GetMapHeaders() map[string]string {
	return r.PortalUser.GetMapHeaders()
}

func (r *CreateTagsWithEndpointIDRequest) AddTag(key, value string) ICreateTagsWithEndpointIDRequest {
	r.Tags = append(r.Tags, struct {
		TagKey   string `json:"tagKey"`
		TagValue string `json:"tagValue"`
	}{
		TagKey:   key,
		TagValue: value,
	})

	return r
}

func (r *CreateTagsWithEndpointIDRequest) ToRequestBody() any {
	return r
}

func (r *CreateTagsWithEndpointIDRequest) GetProjectID() string {
	return r.ProjectID
}

// ____________________________________________________________________ DeleteTagByEndpointIdRequest

type DeleteTagOfEndpointRequest struct {
	common.UserAgent
	common.PortalUser

	ProjectID string
	TagID     string
}

func (r *DeleteTagOfEndpointRequest) ToMap() map[string]any {
	res := map[string]any{
		"tagId": r.TagID,
	}

	if len(r.Agent) > 0 {
		res["userAgent"] = r.Agent
	}

	return res
}

func (r *DeleteTagOfEndpointRequest) AddUserAgent(agent ...string) IDeleteTagOfEndpointRequest {
	r.Agent = append(r.Agent, agent...)
	return r
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
	common.UserAgent
	common.PortalUser

	TagID     string
	ProjectID string
	TagValue  string `json:"tagValue"`
}

func (r *UpdateTagValueOfEndpointRequest) ToMap() map[string]any {
	res := map[string]any{
		"tagId":    r.TagID,
		"tagValue": r.TagValue,
	}

	if len(r.Agent) > 0 {
		res["userAgent"] = r.Agent
	}

	return res
}

func (r *UpdateTagValueOfEndpointRequest) AddUserAgent(agent ...string) IUpdateTagValueOfEndpointRequest {
	r.Agent = append(r.Agent, agent...)
	return r
}

func (r *UpdateTagValueOfEndpointRequest) GetMapHeaders() map[string]string {
	return r.PortalUser.GetMapHeaders()
}

func (r *UpdateTagValueOfEndpointRequest) GetTagID() string {
	return r.TagID
}

func (r *UpdateTagValueOfEndpointRequest) ToRequestBody() any {
	return r
}

func (r *UpdateTagValueOfEndpointRequest) GetProjectID() string {
	return r.ProjectID
}
