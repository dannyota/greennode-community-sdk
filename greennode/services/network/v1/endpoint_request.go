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
	EndpointID string
}

type EndpointNetworking struct {
	Zone       string `json:"zone"`
	SubnetUuid string `json:"subnetUuid"`
}

type EndpointScaling struct {
	MinSize int `json:"minSize"`
	MaxSize int `json:"maxSize"`
}

type CreateEndpointRequest struct {
	ResourceType string `json:"resourceType"`
	Action       string `json:"action"`
	ResourceInfo struct {
		IsBuyMorePoc      bool                 `json:"isBuyMorePoc"`
		IsPoc             bool                 `json:"isPoc"`
		IsEnableAutoRenew bool                 `json:"isEnableAutoRenew"`
		EndpointName      string               `json:"endpointName"`
		CategoryUuid      string               `json:"categoryUuid"`
		ServiceUuid       string               `json:"serviceUuid"`
		PackageUuid       string               `json:"packageUuid"`
		VpcUuid           string               `json:"vpcUuid"`
		PortalUserID      string               `json:"portalUserId"`
		SubnetUuid        string               `json:"subnetUuid"`
		RegionUuid        string               `json:"regionUuid"`
		ProjectUuid       string               `json:"projectUuid"`
		Description       string               `json:"description"`
		EnableAZ          bool                 `json:"enableAZ"`
		EnableDnsName     bool                 `json:"enableDnsName"`
		Networking        []EndpointNetworking `json:"networking"`
		Scaling           EndpointScaling      `json:"scaling"`
	} `json:"resourceInfo"`
}

func (r *CreateEndpointRequest) ToRequestBody(svc *client.ServiceClient) any {
	r.ResourceType = "endpoint"
	r.Action = "create"
	r.ResourceInfo.EnableAZ = true
	r.ResourceInfo.RegionUuid = svc.ZoneID
	r.ResourceInfo.ProjectUuid = svc.ProjectID

	return r
}

type DeleteEndpointByIDRequest struct {
	EndpointServiceUuid string `json:"endpointServiceUuid"`
	EndpointUuid        string `json:"endpointUuid"`
	ProjectUuid         string `json:"projectUuid"`
	RegionUuid          string `json:"regionUuid"`
	VpcUuid             string `json:"vpcUuid"`

	EndpointID string
}

func (r *DeleteEndpointByIDRequest) ToRequestBody(svc *client.ServiceClient) any {
	r.ProjectUuid = svc.ProjectID
	r.RegionUuid = svc.ZoneID

	return r
}

type ListEndpointsRequest struct {
	Page  int
	Size  int
	VpcID string
	Uuid  string
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

func (r *ListEndpointsRequest) getDefaultQuery() string {
	query := fmt.Sprintf(`{"page":%d,"size":%d}`, defaultListEndpointsRequestPage, defaultListEndpointsRequestSize)
	query = "params=" + url.QueryEscape(query)
	return query
}

// _____________________________________________________________________ ListTagsByEndpointIdRequest

type ListTagsByEndpointIDRequest struct {
	EndpointID string
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

func (r *ListTagsByEndpointIDRequest) getDefaultQuery() string {
	query := fmt.Sprintf(`{"page":%d,"size":%d}`, defaultListEndpointsRequestPage, defaultListEndpointsRequestSize)
	query = "params=" + url.QueryEscape(query)
	return query
}

// _________________________________________________________________ CreateTagsWithEndpointIdRequest

type CreateTagsWithEndpointIDRequest struct {
	EndpointID string
	common.PortalUser

	ProjectID    string        `json:"-"`
	ResourceUuid string        `json:"resourceUuid"`
	Tags         []endpointTagResp `json:"tags"`

	SystemTag bool `json:"systemTag"`
}

// ____________________________________________________________________ DeleteTagByEndpointIdRequest

type DeleteTagOfEndpointRequest struct {
	common.PortalUser

	ProjectID string
	TagID     string
}

// _________________________________________________________________ UpdateTagValueOfEndpointRequest

type UpdateTagValueOfEndpointRequest struct {
	common.PortalUser

	TagID     string
	ProjectID string
	TagValue  string `json:"tagValue"`
}

func NewGetEndpointByIDRequest(endpointID string) *GetEndpointByIDRequest {
	return &GetEndpointByIDRequest{
		EndpointID: endpointID,
	}
}

func NewCreateEndpointRequest(name, serviceID, vpcID, subnetID string) *CreateEndpointRequest {
	opts := &CreateEndpointRequest{}
	opts.ResourceInfo.EndpointName = name
	opts.ResourceInfo.ServiceUuid = serviceID
	opts.ResourceInfo.VpcUuid = vpcID
	opts.ResourceInfo.SubnetUuid = subnetID
	opts.ResourceInfo.PackageUuid = defaultPackageID
	opts.ResourceInfo.EnableDnsName = false
	return opts
}

func NewDeleteEndpointByIDRequest(endpointID, vpcID, endpointServiceID string) *DeleteEndpointByIDRequest {
	return &DeleteEndpointByIDRequest{
		EndpointUuid:        endpointID,
		VpcUuid:             vpcID,
		EndpointServiceUuid: endpointServiceID,
		EndpointID:          endpointID,
	}
}

func NewListEndpointsRequest(page, size int) *ListEndpointsRequest {
	return &ListEndpointsRequest{
		Page: page,
		Size: size,
	}
}

func NewListTagsByEndpointIDRequest(userID, projectID, endpointID string) *ListTagsByEndpointIDRequest {
	return &ListTagsByEndpointIDRequest{
		ID:         endpointID,
		ProjectID:  projectID,
		EndpointID: endpointID,
		PortalUser: common.PortalUser{ID: userID},
	}
}

func NewCreateTagsWithEndpointIDRequest(userID, projectID, endpointID string) *CreateTagsWithEndpointIDRequest {
	return &CreateTagsWithEndpointIDRequest{
		ResourceUuid: endpointID,
		SystemTag:    true,
		ProjectID:    projectID,
		EndpointID:   endpointID,
		PortalUser:   common.PortalUser{ID: userID},
	}
}

func NewDeleteTagOfEndpointRequest(userID, projectID, tagID string) *DeleteTagOfEndpointRequest {
	return &DeleteTagOfEndpointRequest{
		TagID:      tagID,
		ProjectID:  projectID,
		PortalUser: common.PortalUser{ID: userID},
	}
}

func NewUpdateTagValueOfEndpointRequest(userID, projectID, tagID, value string) *UpdateTagValueOfEndpointRequest {
	return &UpdateTagValueOfEndpointRequest{
		TagID:      tagID,
		TagValue:   value,
		ProjectID:  projectID,
		PortalUser: common.PortalUser{ID: userID},
	}
}
