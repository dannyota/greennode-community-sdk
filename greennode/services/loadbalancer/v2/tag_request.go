package v2

import (
	"github.com/dannyota/greennode-community-sdk/v2/greennode/entity"
	"github.com/dannyota/greennode-community-sdk/v2/greennode/services/common"
)

func NewListTagsRequest(lbID string) IListTagsRequest {
	opt := new(ListTagsRequest)
	opt.LoadBalancerID = lbID
	return opt
}

func (r *ListTagsRequest) AddUserAgent(agent ...string) IListTagsRequest {
	r.UserAgent.AddUserAgent(agent...)
	return r
}

func NewCreateTagsRequest(lbID string) ICreateTagsRequest {
	opts := new(CreateTagsRequest)
	opts.LoadBalancerID = lbID
	opts.ResourceID = lbID
	opts.ResourceType = "LOAD-BALANCER"
	opts.TagRequestList = make([]common.Tag, 0)

	return opts
}

func NewUpdateTagsRequest(lbID string) IUpdateTagsRequest {
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

func (r *CreateTagsRequest) AddUserAgent(agent ...string) ICreateTagsRequest {
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

func (r *CreateTagsRequest) ToRequestBody() any {
	return r
}

func (r *UpdateTagsRequest) AddUserAgent(agent ...string) IUpdateTagsRequest {
	r.UserAgent.AddUserAgent(agent...)
	return r
}

func (r *CreateTagsRequest) WithTags(tags ...string) ICreateTagsRequest {
	if len(tags)%2 != 0 {
		tags = append(tags, "none")
	}

	for i := 0; i < len(tags); i += 2 {
		r.TagRequestList = append(r.TagRequestList, common.Tag{Key: tags[i], Value: tags[i+1]})
	}
	return r
}

func (r *UpdateTagsRequest) ToRequestBody(lstTags *entity.ListTags) any {
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

	return r
}

func (r *UpdateTagsRequest) WithTags(tags ...string) IUpdateTagsRequest {
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
