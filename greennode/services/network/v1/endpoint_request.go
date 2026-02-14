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

func (s *GetEndpointByIDRequest) AddUserAgent(agent ...string) IGetEndpointByIDRequest {
	s.Agent = append(s.Agent, agent...)
	return s
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

func (s *CreateEndpointRequest) ToMap() map[string]interface{} {
	res := map[string]interface{}{
		"resourceType":      s.ResourceType,
		"action":            s.Action,
		"isBuyMorePoc":      s.ResourceInfo.IsBuyMorePoc,
		"isPoc":             s.ResourceInfo.IsPoc,
		"isEnableAutoRenew": s.ResourceInfo.IsEnableAutoRenew,
		"endpointName":      s.ResourceInfo.EndpointName,
		"categoryID":        s.ResourceInfo.CategoryUuid,
		"serviceId":         s.ResourceInfo.ServiceUuid,
		"packageId":         s.ResourceInfo.PackageUuid,
		"vpcId":             s.ResourceInfo.VpcUuid,
		"subnetId":          s.ResourceInfo.SubnetUuid,
		"regionId":          s.ResourceInfo.RegionUuid,
		"projectId":         s.ResourceInfo.ProjectUuid,
		"description":       s.ResourceInfo.Description,
		"enableAZ":          s.ResourceInfo.EnableAZ,
		"enableDnsName":     s.ResourceInfo.EnableDnsName,
		"networking":        s.ResourceInfo.Networking,
		"scaling":           s.ResourceInfo.Scaling,
	}

	if len(s.Agent) > 0 {
		res["userAgent"] = s.Agent
	}

	return res
}

func (s *CreateEndpointRequest) AddUserAgent(agent ...string) ICreateEndpointRequest {
	s.Agent = append(s.Agent, agent...)
	return s
}

func (s *CreateEndpointRequest) ToRequestBody(svc client.ServiceClient) interface{} {
	s.ResourceType = "endpoint"
	s.Action = "create"
	s.ResourceInfo.EnableAZ = true
	s.ResourceInfo.RegionUuid = svc.GetZoneID()
	s.ResourceInfo.ProjectUuid = svc.GetProjectID()

	return s
}

func (s *CreateEndpointRequest) WithEndpointName(endpointName string) ICreateEndpointRequest {
	s.ResourceInfo.EndpointName = endpointName
	return s
}

func (s *CreateEndpointRequest) WithCategoryUuid(categoryUuid string) ICreateEndpointRequest {
	s.ResourceInfo.CategoryUuid = categoryUuid
	return s
}

func (s *CreateEndpointRequest) WithServiceUuid(serviceUuid string) ICreateEndpointRequest {
	s.ResourceInfo.ServiceUuid = serviceUuid
	return s
}

func (s *CreateEndpointRequest) WithPackageUuid(packageUuid string) ICreateEndpointRequest {
	s.ResourceInfo.PackageUuid = packageUuid
	return s
}

func (s *CreateEndpointRequest) WithVpcUuid(vpcUuid string) ICreateEndpointRequest {
	s.ResourceInfo.VpcUuid = vpcUuid
	return s
}

func (s *CreateEndpointRequest) WithPortalUserID(portalUserID string) ICreateEndpointRequest {
	s.ResourceInfo.PortalUserID = portalUserID
	return s
}

func (s *CreateEndpointRequest) GetPortalUserID() string {
	return s.ResourceInfo.PortalUserID
}

func (s *CreateEndpointRequest) WithSubnetUuid(subnetUuid string) ICreateEndpointRequest {
	s.ResourceInfo.SubnetUuid = subnetUuid
	return s
}

func (s *CreateEndpointRequest) WithDescription(desp string) ICreateEndpointRequest {
	s.ResourceInfo.Description = desp
	return s
}

func (s *CreateEndpointRequest) WithPoc(yes bool) ICreateEndpointRequest {
	s.ResourceInfo.IsPoc = yes
	return s
}

func (s *CreateEndpointRequest) WithEnableDnsName(yes bool) ICreateEndpointRequest {
	s.ResourceInfo.EnableDnsName = yes
	return s
}

func (s *CreateEndpointRequest) WithBuyMorePoc(yes bool) ICreateEndpointRequest {
	s.ResourceInfo.IsBuyMorePoc = yes
	return s
}

func (s *CreateEndpointRequest) WithEnableAutoRenew(yes bool) ICreateEndpointRequest {
	s.ResourceInfo.IsEnableAutoRenew = yes
	return s
}

func (s *CreateEndpointRequest) AddNetworking(zone, subnetUuid string) ICreateEndpointRequest {
	s.ResourceInfo.Networking = append(s.ResourceInfo.Networking, struct {
		Zone       string `json:"zone"`
		SubnetUuid string `json:"subnetUuid"`
	}{
		Zone:       zone,
		SubnetUuid: subnetUuid,
	})
	return s
}

func (s *CreateEndpointRequest) WithScaling(minSize int, maxSize int) ICreateEndpointRequest {
	s.ResourceInfo.Scaling = struct {
		MinSize int `json:"minSize"`
		MaxSize int `json:"maxSize"`
	}{
		MinSize: minSize,
		MaxSize: maxSize,
	}
	return s
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

func (s *DeleteEndpointByIDRequest) AddUserAgent(agent ...string) IDeleteEndpointByIDRequest {
	s.Agent = append(s.Agent, agent...)
	return s
}

func (s *DeleteEndpointByIDRequest) ToRequestBody(svc client.ServiceClient) interface{} {
	s.ProjectUuid = svc.GetProjectID()
	s.RegionUuid = svc.GetZoneID()

	return s
}

func (s *DeleteEndpointByIDRequest) ToMap() map[string]interface{} {
	res := map[string]interface{}{
		"serviceId":  s.EndpointServiceUuid,
		"endpointId": s.EndpointID,
		"projectId":  s.ProjectUuid,
		"regionId":   s.RegionUuid,
		"vpcId":      s.VpcUuid,
	}

	if len(s.Agent) > 0 {
		res["userAgent"] = s.Agent
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

func (s *ListEndpointsRequest) ToMap() map[string]interface{} {
	res := map[string]interface{}{
		"page":  s.Page,
		"size":  s.Size,
		"vpcId": s.VpcID,
		"uuid":  s.Uuid,
	}

	if len(s.Agent) > 0 {
		res["userAgent"] = s.Agent
	}

	return res
}

func (s *ListEndpointsRequest) WithPage(page int) IListEndpointsRequest {
	s.Page = page
	return s
}

func (s *ListEndpointsRequest) WithSize(size int) IListEndpointsRequest {
	s.Size = size
	return s
}

func (s *ListEndpointsRequest) WithVpcID(vpcID string) IListEndpointsRequest {
	s.VpcID = vpcID
	return s
}

func (s *ListEndpointsRequest) WithUuid(uuid string) IListEndpointsRequest {
	s.Uuid = uuid
	return s
}

func (s *ListEndpointsRequest) ToListQuery() (string, error) {
	var params []string
	if s.VpcID != "" {
		params = append(params, fmt.Sprintf(`{"field":"vpcId","value":"%s"}`, s.VpcID))
	}

	if s.Uuid != "" {
		params = append(params, fmt.Sprintf(`{"field":"uuid","value":"%s"}`, s.Uuid))
	}

	paramsFilter := strings.Join(params, ",")
	query := fmt.Sprintf(`{"page":%d,"size":%d,"search":[%s]}`, s.Page, s.Size, paramsFilter)
	query = "params=" + url.QueryEscape(query)

	return query, nil
}

func (s *ListEndpointsRequest) GetDefaultQuery() string {
	query := fmt.Sprintf(`{"page":%d,"size":%d}`, defaultListEndpointsRequestPage, defaultListEndpointsRequestSize)
	query = "params=" + url.QueryEscape(query)
	return query
}

func (s *ListEndpointsRequest) AddUserAgent(agent ...string) IListEndpointsRequest {
	s.Agent = append(s.Agent, agent...)
	return s
}

// _____________________________________________________________________ ListTagsByEndpointIdRequest

type ListTagsByEndpointIDRequest struct {
	common.UserAgent
	common.EndpointCommon
	common.PortalUser

	ProjectID string
	ID        string
}

func (s *ListTagsByEndpointIDRequest) ToListQuery() (string, error) {
	v := url.Values{}
	if s.ID != "" {
		v.Set("resourceUuid", s.ID)
	}
	return v.Encode(), nil
}

func (s *ListTagsByEndpointIDRequest) GetProjectID() string {
	return s.ProjectID
}

func (s *ListTagsByEndpointIDRequest) GetDefaultQuery() string {
	query := fmt.Sprintf(`{"page":%d,"size":%d}`, defaultListEndpointsRequestPage, defaultListEndpointsRequestSize)
	query = "params=" + url.QueryEscape(query)
	return query
}

func (s *ListTagsByEndpointIDRequest) ToMap() map[string]interface{} {
	res := map[string]interface{}{
		"resourceUuid": s.ID,
	}

	if len(s.Agent) > 0 {
		res["userAgent"] = s.Agent
	}

	return res
}

func (s *ListTagsByEndpointIDRequest) GetMapHeaders() map[string]string {
	return s.PortalUser.GetMapHeaders()
}

func (s *ListTagsByEndpointIDRequest) AddUserAgent(agent ...string) IListTagsByEndpointIDRequest {
	s.Agent = append(s.Agent, agent...)
	return s
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

func (s *CreateTagsWithEndpointIDRequest) ToMap() map[string]interface{} {
	res := map[string]interface{}{
		"resourceUuid": s.ID,
	}

	if len(s.Agent) > 0 {
		res["userAgent"] = s.Agent
	}

	res["tags"] = s.Tags

	return res
}

func (s *CreateTagsWithEndpointIDRequest) AddUserAgent(agent ...string) ICreateTagsWithEndpointIDRequest {
	s.Agent = append(s.Agent, agent...)
	return s
}

func (s *CreateTagsWithEndpointIDRequest) GetMapHeaders() map[string]string {
	return s.PortalUser.GetMapHeaders()
}

func (s *CreateTagsWithEndpointIDRequest) AddTag(key, value string) ICreateTagsWithEndpointIDRequest {
	s.Tags = append(s.Tags, struct {
		TagKey   string `json:"tagKey"`
		TagValue string `json:"tagValue"`
	}{
		TagKey:   key,
		TagValue: value,
	})

	return s
}

func (s *CreateTagsWithEndpointIDRequest) ToRequestBody() interface{} {
	return s
}

func (s *CreateTagsWithEndpointIDRequest) GetProjectID() string {
	return s.ProjectID
}

// ____________________________________________________________________ DeleteTagByEndpointIdRequest

type DeleteTagOfEndpointRequest struct {
	common.UserAgent
	common.PortalUser

	ProjectID string
	TagID     string
}

func (s *DeleteTagOfEndpointRequest) ToMap() map[string]interface{} {
	res := map[string]interface{}{
		"tagId": s.TagID,
	}

	if len(s.Agent) > 0 {
		res["userAgent"] = s.Agent
	}

	return res
}

func (s *DeleteTagOfEndpointRequest) AddUserAgent(agent ...string) IDeleteTagOfEndpointRequest {
	s.Agent = append(s.Agent, agent...)
	return s
}

func (s *DeleteTagOfEndpointRequest) GetMapHeaders() map[string]string {
	return s.PortalUser.GetMapHeaders()
}

func (s *DeleteTagOfEndpointRequest) GetTagID() string {
	return s.TagID
}

func (s *DeleteTagOfEndpointRequest) GetProjectID() string {
	return s.ProjectID
}

// _________________________________________________________________ UpdateTagValueOfEndpointRequest

type UpdateTagValueOfEndpointRequest struct {
	common.UserAgent
	common.PortalUser

	TagID     string
	ProjectID string
	TagValue  string `json:"tagValue"`
}

func (s *UpdateTagValueOfEndpointRequest) ToMap() map[string]interface{} {
	res := map[string]interface{}{
		"tagId":    s.TagID,
		"tagValue": s.TagValue,
	}

	if len(s.Agent) > 0 {
		res["userAgent"] = s.Agent
	}

	return res
}

func (s *UpdateTagValueOfEndpointRequest) AddUserAgent(agent ...string) IUpdateTagValueOfEndpointRequest {
	s.Agent = append(s.Agent, agent...)
	return s
}

func (s *UpdateTagValueOfEndpointRequest) GetMapHeaders() map[string]string {
	return s.PortalUser.GetMapHeaders()
}

func (s *UpdateTagValueOfEndpointRequest) GetTagID() string {
	return s.TagID
}

func (s *UpdateTagValueOfEndpointRequest) ToRequestBody() interface{} {
	return s
}

func (s *UpdateTagValueOfEndpointRequest) GetProjectID() string {
	return s.ProjectID
}
