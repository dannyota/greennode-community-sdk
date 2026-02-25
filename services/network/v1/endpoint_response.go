package v1

import "danny.vn/greennode/types"

type endpointCategoryResp struct {
	Uuid      string `json:"uuid,omitempty"`
	Name      string `json:"name,omitempty"`
	IsDefault bool   `json:"isDefault,omitempty"`
}

type endpointServiceDetailResp struct {
	EndpointAuthURL    string `json:"endpoint_auth_url,omitempty"`
	EndpointURL        string `json:"endpoint_url,omitempty"`
	TargetCIDR         string `json:"target_cidr,omitempty"`
	EndpointEncryptURL string `json:"endpoint_encrypt_url,omitempty"`
}

type endpointServiceResp struct {
	Uuid         string                     `json:"uuid,omitempty"`
	Name         string                     `json:"name,omitempty"`
	EndpointURL  string                     `json:"endpointUrl,omitempty"`
	EndpointType string                     `json:"endpointType,omitempty"`
	Detail       *endpointServiceDetailResp `json:"endpointDetailInformation,omitempty"`
}

type endpointVPCResp struct {
	Uuid      string `json:"uuid,omitempty"`
	Name      string `json:"name,omitempty"`
	CIDR      string `json:"cidr,omitempty"`
	Status    string `json:"status,omitempty"`
	DnsStatus string `json:"dnsStatus,omitempty"`
}

type endpointSubnetResp struct {
	Uuid   string `json:"uuid,omitempty"`
	Name   string `json:"name,omitempty"`
	Status string `json:"status,omitempty"`
	CIDR   string `json:"cidr,omitempty"`
	ZoneID string `json:"zoneId,omitempty"`
}

type endpointPackageResp struct {
	Uuid        string `json:"uuid,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

type endpointProjectResp struct {
	Uuid             string `json:"id,omitempty"`
	BackendProjectID string `json:"backendProjectId,omitempty"`
	PortalUserID     int    `json:"portalUserId,omitempty"`
	VServerProjectID string `json:"vserverProjectId,omitempty"`
}

type endpointResp struct {
	Uuid              string `json:"uuid,omitempty"`
	EndpointName      string `json:"endpointName,omitempty"`
	EndpointServiceID string `json:"endpointServiceId,omitempty"`
	VpcID             string `json:"vpcId,omitempty"`
	EndpointURL       string `json:"endpointUrl,omitempty"`
	EndpointAuthURL   string `json:"endpointAuthUrl,omitempty"`
	EndpointIp        string `json:"endpointIp,omitempty"`
	Status            string `json:"status,omitempty"`
	BillingStatus     string `json:"billingStatus,omitempty"`
	EndpointType      string `json:"endpointType,omitempty"`
	Version           string `json:"version,omitempty"`
	Description       string `json:"description,omitempty"`
	CreatedAt         string `json:"createdAt,omitempty"`
	UpdatedAt         string `json:"updatedAt,omitempty"`
	ZoneUuid          string `json:"zoneUuid,omitempty"`
	EnableDnsName     bool   `json:"enableDnsName,omitempty"`

	EndpointDomains []string             `json:"endpointDomains,omitempty"`
	Category        *endpointCategoryResp `json:"category,omitempty"`
	Service         *endpointServiceResp  `json:"service,omitempty"`
	VPC             *endpointVPCResp      `json:"vpc,omitempty"`
	Subnet          *endpointSubnetResp   `json:"subnet,omitempty"`
	Package         *endpointPackageResp  `json:"packageId,omitempty"`
	Project         *endpointProjectResp  `json:"project,omitempty"`
}

type endpointTagResp struct {
	Uuid         string `json:"uuid,omitempty"`
	TagKey       string `json:"tagKey,omitempty"`
	TagValue     string `json:"tagValue,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	ResourceType string `json:"resourceType,omitempty"`
	SystemTag    bool   `json:"systemTag,omitempty"`
	CreatedAt    string `json:"createdAt,omitempty"`
	UpdatedAt    string `json:"updatedAt,omitempty"`
}

func (e *endpointResp) toEntityEndpoint() *Endpoint {
	ep := &Endpoint{
		UUID:              e.Uuid,
		Name:              e.EndpointName,
		VpcID:             e.VpcID,
		IPv4Address:       e.EndpointIp,
		EndpointURL:       e.EndpointURL,
		EndpointAuthURL:   e.EndpointAuthURL,
		Status:            e.Status,
		EndpointServiceID: e.EndpointServiceID,
		BillingStatus:     e.BillingStatus,
		EndpointType:      e.EndpointType,
		Version:           e.Version,
		Description:       e.Description,
		CreatedAt:         e.CreatedAt,
		UpdatedAt:         e.UpdatedAt,
		ZoneUuid:          e.ZoneUuid,
		EnableDnsName:     e.EnableDnsName,
		EndpointDomains:   e.EndpointDomains,
	}

	if e.Category != nil {
		ep.Category = &EndpointCategory{
			UUID:      e.Category.Uuid,
			Name:      e.Category.Name,
			IsDefault: e.Category.IsDefault,
		}
	}

	if e.Service != nil {
		svc := &EndpointService{
			UUID:         e.Service.Uuid,
			Name:         e.Service.Name,
			EndpointURL:  e.Service.EndpointURL,
			EndpointType: e.Service.EndpointType,
		}
		if e.Service.Detail != nil {
			svc.Detail = &EndpointServiceDetail{
				EndpointAuthURL:    e.Service.Detail.EndpointAuthURL,
				EndpointURL:        e.Service.Detail.EndpointURL,
				TargetCIDR:         e.Service.Detail.TargetCIDR,
				EndpointEncryptURL: e.Service.Detail.EndpointEncryptURL,
			}
		}
		ep.Service = svc
	}

	if e.VPC != nil {
		ep.VPC = &EndpointVPC{
			UUID:      e.VPC.Uuid,
			Name:      e.VPC.Name,
			CIDR:      e.VPC.CIDR,
			Status:    e.VPC.Status,
			DnsStatus: e.VPC.DnsStatus,
		}
	}

	if e.Subnet != nil {
		ep.Subnet = &EndpointSubnet{
			UUID:   e.Subnet.Uuid,
			Name:   e.Subnet.Name,
			Status: e.Subnet.Status,
			CIDR:   e.Subnet.CIDR,
			ZoneID: e.Subnet.ZoneID,
		}
	}

	if e.Package != nil {
		ep.Package = &EndpointPackage{
			UUID:        e.Package.Uuid,
			Name:        e.Package.Name,
			Description: e.Package.Description,
		}
	}

	if e.Project != nil {
		ep.Project = &EndpointProject{
			ID:               e.Project.Uuid,
			BackendProjectID: e.Project.BackendProjectID,
			PortalUserID:     e.Project.PortalUserID,
			VServerProjectID: e.Project.VServerProjectID,
		}
	}

	return ep
}

type GetEndpointByIDResponse struct {
	Data endpointResp `json:"data"`
}

func (r *GetEndpointByIDResponse) ToEntityEndpoint() *Endpoint {
	return r.Data.toEntityEndpoint()
}

type CreateEndpointResponse struct {
	Data struct {
		Uuid string `json:"uuid,omitempty"`
		Name string `json:"name,omitempty"`
	} `json:"data"`
}

func (r *CreateEndpointResponse) ToEntityEndpoint() *Endpoint {
	return &Endpoint{
		UUID: r.Data.Uuid,
		Name: r.Data.Name,
	}
}

type ListEndpointsResponse struct {
	Data      []endpointResp `json:"data"`
	Page      int            `json:"page"`
	Size      int            `json:"size"`
	TotalPage int            `json:"totalPage"`
	Total     int            `json:"total"`
}

func (r *ListEndpointsResponse) ToEntityListEndpoints() *ListEndpoints {
	items := make([]*Endpoint, 0, len(r.Data))
	for _, item := range r.Data {
		items = append(items, item.toEntityEndpoint())
	}
	return &ListEndpoints{
		Items:     items,
		Page:      r.Page,
		PageSize:  r.Size,
		TotalPage: r.TotalPage,
		TotalItem: r.Total,
	}
}

type ListTagsByEndpointIDResponse struct {
	Data []endpointTagResp `json:"data"`
}

func (r *ListTagsByEndpointIDResponse) ToEntityListTags() *types.ListTags {
	items := make([]*types.Tag, 0, len(r.Data))
	for _, item := range r.Data {
		items = append(items, &types.Tag{
			Key:        item.TagKey,
			Value:      item.TagValue,
			SystemTag:  item.SystemTag,
			ResourceID: item.ResourceUuid,
			TagID:      item.Uuid,
		})
	}
	return &types.ListTags{
		Items: items,
	}
}
