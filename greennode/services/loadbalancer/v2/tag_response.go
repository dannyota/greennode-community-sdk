package v2

import "github.com/dannyota/greennode-community-sdk/v2/greennode/entity"

type ListTagResponse struct {
	Key       string `json:"key"`
	Value     string `json:"value"`
	CreatedAt string `json:"createdAt"`
	SystemTag bool   `json:"systemTag,omitempty"`
}

type ListTagsResponse []ListTagResponse

func (r ListTagResponse) ToEntityTag() *entity.Tag {
	return &entity.Tag{
		Key:       r.Key,
		Value:     r.Value,
		SystemTag: r.SystemTag,
	}
}

func (r *ListTagsResponse) ToEntityListTags() *entity.ListTags {
	result := new(entity.ListTags)
	if r == nil {
		return result
	}

	for _, item := range *r {
		result.Add(item.ToEntityTag())
	}

	return result
}
