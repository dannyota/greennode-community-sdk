package v2

import (
	"github.com/dannyota/greennode-community-sdk/v2/greennode/entity"
	"github.com/dannyota/greennode-community-sdk/v2/greennode/services/common"
)

func NewListTagsRequest(lbID string) *ListTagsRequest {
	opt := new(ListTagsRequest)
	opt.LoadBalancerID = lbID
	return opt
}

func (r *ListTagsRequest) AddUserAgent(agent ...string) *ListTagsRequest {
	r.UserAgent.AddUserAgent(agent...)
	return r
}

func NewCreateTagsRequest(lbID string) *CreateTagsRequest {
	opts := new(CreateTagsRequest)
	opts.LoadBalancerID = lbID
	opts.ResourceID = lbID
	opts.ResourceType = "LOAD-BALANCER"
	opts.TagRequestList = make([]common.Tag, 0)

	return opts
}

func NewUpdateTagsRequest(lbID string) *UpdateTagsRequest {
	opts := new(UpdateTagsRequest)
	opts.LoadBalancerID = lbID
	opts.ResourceID = lbID
	opts.ResourceType = "LOAD-BALANCER"
	opts.TagRequestList = make([]common.Tag, 0)

	return opts
}

type ListTagsRequest struct {
	common.UserAgent
	common.LoadBalancerCommon
}

type CreateTagsRequest struct {
	ResourceID     string       `json:"resourceId"`
	ResourceType   string       `json:"resourceType"`
	TagRequestList []common.Tag `json:"tagRequestList"`

	common.UserAgent
	common.LoadBalancerCommon
}

func (r *CreateTagsRequest) AddUserAgent(agent ...string) *CreateTagsRequest {
	r.UserAgent.AddUserAgent(agent...)
	return r
}

type UpdateTagsRequest struct {
	ResourceID     string       `json:"resourceId"`
	ResourceType   string       `json:"resourceType"`
	TagRequestList []common.Tag `json:"tagRequestList"`

	common.UserAgent
	common.LoadBalancerCommon
}

func (r *UpdateTagsRequest) AddUserAgent(agent ...string) *UpdateTagsRequest {
	r.UserAgent.AddUserAgent(agent...)
	return r
}

func (r *CreateTagsRequest) WithTags(tags ...string) *CreateTagsRequest {
	if len(tags)%2 != 0 {
		tags = append(tags, "none")
	}

	for i := 0; i < len(tags); i += 2 {
		r.TagRequestList = append(r.TagRequestList, common.Tag{Key: tags[i], Value: tags[i+1]})
	}
	return r
}

func (r *UpdateTagsRequest) prepare(lstTags *entity.ListTags) {
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

func (r *UpdateTagsRequest) WithTags(tags ...string) *UpdateTagsRequest {
	if len(tags)%2 != 0 {
		tags = append(tags, "none")
	}

	for i := 0; i < len(tags); i += 2 {
		r.TagRequestList = append(r.TagRequestList, common.Tag{Key: tags[i], Value: tags[i+1]})
	}

	return r
}

func (r *UpdateTagsRequest) ToMap() map[string]any {
	res := make(map[string]any)
	for _, tag := range r.TagRequestList {
		res[tag.Key] = tag.Value
	}
	return res
}
