package v2

import (
	types "github.com/dannyota/greennode-community-sdk/greennode/types"
	"github.com/dannyota/greennode-community-sdk/greennode/services/common"
)

func NewListTagsRequest(lbID string) *ListTagsRequest {
	return &ListTagsRequest{
		LoadBalancerID: lbID,
	}
}

func NewCreateTagsRequest(lbID string) *CreateTagsRequest {
	return &CreateTagsRequest{
		LoadBalancerID: lbID,
		ResourceID:     lbID,
		ResourceType:   "LOAD-BALANCER",
		TagRequestList: make([]common.Tag, 0),
	}
}

func NewUpdateTagsRequest(lbID string) *UpdateTagsRequest {
	return &UpdateTagsRequest{
		LoadBalancerID: lbID,
		ResourceID:     lbID,
		ResourceType:   "LOAD-BALANCER",
		TagRequestList: make([]common.Tag, 0),
	}
}

type ListTagsRequest struct {
	LoadBalancerID string
}

type CreateTagsRequest struct {
	ResourceID     string       `json:"resourceId"`
	ResourceType   string       `json:"resourceType"`
	TagRequestList []common.Tag `json:"tagRequestList"`

	LoadBalancerID string
}

type UpdateTagsRequest struct {
	ResourceID     string       `json:"resourceId"`
	ResourceType   string       `json:"resourceType"`
	TagRequestList []common.Tag `json:"tagRequestList"`

	LoadBalancerID string
}

func (r *UpdateTagsRequest) prepare(lstTags *types.ListTags) {
	st := map[string]common.Tag{}
	for _, tag := range lstTags.Items {
		st[tag.Key] = common.Tag{
			IsEdited: false,
			Key:      tag.Key,
			Value:    tag.Value,
		}
	}

	for _, tag := range r.TagRequestList {
		st[tag.Key] = tag
	}

	r.TagRequestList = make([]common.Tag, 0)
	for _, tag := range st {
		r.TagRequestList = append(r.TagRequestList, tag)
	}
}
