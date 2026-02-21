package v2

import types "github.com/dannyota/greennode-community-sdk/greennode/types"

type ListTagResponse struct {
	Key       string `json:"key"`
	Value     string `json:"value"`
	CreatedAt string `json:"createdAt"`
	SystemTag bool   `json:"systemTag,omitempty"`
}

type ListTagsResponse []ListTagResponse

func (r ListTagResponse) ToEntityTag() *types.Tag {
	return &types.Tag{
		Key:       r.Key,
		Value:     r.Value,
		SystemTag: r.SystemTag,
	}
}

func (r *ListTagsResponse) ToEntityListTags() *types.ListTags {
	result := new(types.ListTags)
	if r == nil {
		return result
	}

	for _, item := range *r {
		result.Add(item.ToEntityTag())
	}

	return result
}
