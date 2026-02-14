package v2

import "github.com/dannyota/greennode-community-sdk/v2/greennode/entity"

type ListTagResponse struct {
	Key       string `json:"key"`
	Value     string `json:"value"`
	CreatedAt string `json:"createdAt"`
	SystemTag bool   `json:"systemTag,omitempty"`
}

type ListTagsResponse []ListTagResponse

func (s ListTagResponse) ToEntityTag() *entity.Tag {
	return &entity.Tag{
		Key:       s.Key,
		Value:     s.Value,
		SystemTag: s.SystemTag,
	}
}

func (s *ListTagsResponse) ToEntityListTags() *entity.ListTags {
	result := new(entity.ListTags)
	if s == nil {
		return result
	}

	for _, item := range *s {
		result.Add(item.ToEntityTag())
	}

	return result
}
