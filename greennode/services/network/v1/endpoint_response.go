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

func (s *Endpoint) toEntityEndpoint() *entity.Endpoint {
	return &entity.Endpoint{
		ID:          s.Uuid,
		Name:        s.EndpointName,
		VpcID:       s.VpcID,
		IPv4Address: s.EndpointIp,
		EndpointURL: s.EndpointURL,
		Status:      s.Status,
	}
}

type GetEndpointByIDResponse struct {
	Data Endpoint `json:"data"`
}

func (s *GetEndpointByIDResponse) ToEntityEndpoint() *entity.Endpoint {
	return s.Data.toEntityEndpoint()
}

type CreateEndpointResponse struct {
	Data struct {
		Uuid string `json:"uuid,omitempty"`
		Name string `json:"name,omitempty"`
	} `json:"data"`
}

func (s *CreateEndpointResponse) ToEntityEndpoint() *entity.Endpoint {
	return &entity.Endpoint{
		ID:   s.Data.Uuid,
		Name: s.Data.Name,
	}
}

type ListEndpointsResponse struct {
	Data      []Endpoint `json:"data"`
	Page      int        `json:"page"`
	Size      int        `json:"size"`
	TotalPage int        `json:"totalPage"`
	Total     int        `json:"total"`
}

func (s *ListEndpointsResponse) ToEntityListEndpoints() *entity.ListEndpoints {
	items := make([]*entity.Endpoint, 0, len(s.Data))
	for _, item := range s.Data {
		items = append(items, item.toEntityEndpoint())
	}
	return &entity.ListEndpoints{
		Items:     items,
		Page:      s.Page,
		PageSize:  s.Size,
		TotalPage: s.TotalPage,
		TotalItem: s.Total,
	}
}

type ListTagsByEndpointIDResponse struct {
	Data []EndpointTag `json:"data"`
}

func (s *ListTagsByEndpointIDResponse) ToEntityListTags() *entity.ListTags {
	items := make([]*entity.Tag, 0, len(s.Data))
	for _, item := range s.Data {
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
