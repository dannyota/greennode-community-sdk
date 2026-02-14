package v1

import "github.com/dannyota/greennode-community-sdk/v2/greennode/entity"

type Endpoint struct {
	Uuid              string `json:"uuid,omitempty"`
	EndpointName      string `json:"endpointName,omitempty"`
	EndpointServiceID string `json:"endpointServiceId,omitempty"`
	VpcID             string `json:"vpcId,omitempty"`
	EndpointURL       string `json:"endpointUrl,omitempty"`
	EndpointIp        string `json:"endpointIp,omitempty"`
	Status            string `json:"status,omitempty"`
}

type EndpointTag struct {
	Uuid         string `json:"uuid,omitempty"`
	TagKey       string `json:"tagKey,omitempty"`
	TagValue     string `json:"tagValue,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	ResourceType string `json:"resourceType,omitempty"`
	SystemTag    bool   `json:"systemTag,omitempty"`
	CreatedAt    string `json:"createdAt,omitempty"`
	UpdatedAt    string `json:"updatedAt,omitempty"`
}

func (e *Endpoint) toEntityEndpoint() *entity.Endpoint {
	return &entity.Endpoint{
		ID:          e.Uuid,
		Name:        e.EndpointName,
		VpcID:       e.VpcID,
		IPv4Address: e.EndpointIp,
		EndpointURL: e.EndpointURL,
		Status:      e.Status,
	}
}

type GetEndpointByIDResponse struct {
	Data Endpoint `json:"data"`
}

func (r *GetEndpointByIDResponse) ToEntityEndpoint() *entity.Endpoint {
	return r.Data.toEntityEndpoint()
}

type CreateEndpointResponse struct {
	Data struct {
		Uuid string `json:"uuid,omitempty"`
		Name string `json:"name,omitempty"`
	} `json:"data"`
}

func (r *CreateEndpointResponse) ToEntityEndpoint() *entity.Endpoint {
	return &entity.Endpoint{
		ID:   r.Data.Uuid,
		Name: r.Data.Name,
	}
}

type ListEndpointsResponse struct {
	Data      []Endpoint `json:"data"`
	Page      int        `json:"page"`
	Size      int        `json:"size"`
	TotalPage int        `json:"totalPage"`
	Total     int        `json:"total"`
}

func (r *ListEndpointsResponse) ToEntityListEndpoints() *entity.ListEndpoints {
	items := make([]*entity.Endpoint, 0, len(r.Data))
	for _, item := range r.Data {
		items = append(items, item.toEntityEndpoint())
	}
	return &entity.ListEndpoints{
		Items:     items,
		Page:      r.Page,
		PageSize:  r.Size,
		TotalPage: r.TotalPage,
		TotalItem: r.Total,
	}
}

type ListTagsByEndpointIDResponse struct {
	Data []EndpointTag `json:"data"`
}

func (r *ListTagsByEndpointIDResponse) ToEntityListTags() *entity.ListTags {
	items := make([]*entity.Tag, 0, len(r.Data))
	for _, item := range r.Data {
		items = append(items, &entity.Tag{
			Key:        item.TagKey,
			Value:      item.TagValue,
			SystemTag:  item.SystemTag,
			ResourceID: item.ResourceUuid,
			TagID:      item.Uuid,
		})
	}
	return &entity.ListTags{
		Items: items,
	}
}
