package v2

import (
	"github.com/dannyota/greennode-community-sdk/v2/greennode/entity"
	"github.com/dannyota/greennode-community-sdk/v2/greennode/services/common"
)

func NewListTagsRequest(lbId string) IListTagsRequest {
	opt := new(ListTagsRequest)
	opt.LoadBalancerId = lbId
	return opt
}

func (s *ListTagsRequest) AddUserAgent(agent ...string) IListTagsRequest {
	s.UserAgent.AddUserAgent(agent...)
	return s
}

func NewCreateTagsRequest(lbId string) ICreateTagsRequest {
	opts := new(CreateTagsRequest)
	opts.LoadBalancerId = lbId
	opts.ResourceID = lbId
	opts.ResourceType = "LOAD-BALANCER"
	opts.TagRequestList = make([]common.Tag, 0)

	return opts
}

func NewUpdateTagsRequest(lbId string) IUpdateTagsRequest {
	opts := new(UpdateTagsRequest)
	opts.LoadBalancerId = lbId
	opts.ResourceID = lbId
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

func (s *CreateTagsRequest) AddUserAgent(agent ...string) ICreateTagsRequest {
	s.UserAgent.AddUserAgent(agent...)
	return s
}

type UpdateTagsRequest struct {
	ResourceID     string       `json:"resourceId"`
	ResourceType   string       `json:"resourceType"`
	TagRequestList []common.Tag `json:"tagRequestList"`

	common.UserAgent
	common.LoadBalancerCommon
}

func (s *CreateTagsRequest) ToRequestBody() interface{} {
	return s
}

func (s *UpdateTagsRequest) AddUserAgent(agent ...string) IUpdateTagsRequest {
	s.UserAgent.AddUserAgent(agent...)
	return s
}

func (s *CreateTagsRequest) WithTags(tags ...string) ICreateTagsRequest {
	if len(tags)%2 != 0 {
		tags = append(tags, "none")
	}

	for i := 0; i < len(tags); i += 2 {
		s.TagRequestList = append(s.TagRequestList, common.Tag{Key: tags[i], Value: tags[i+1]})
	}
	return s
}

func (s *UpdateTagsRequest) ToRequestBody(lstTags *entity.ListTags) interface{} {
	st := map[string]common.Tag{}
	for _, tag := range lstTags.Items {
		st[tag.Key] = common.Tag{
			IsEdited: false,
			Key:      tag.Key,
			Value:    tag.Value,
		}
	}

	for _, tag := range s.TagRequestList {
		st[tag.Key] = tag
	}

	s.TagRequestList = make([]common.Tag, 0)
	for _, tag := range st {
		s.TagRequestList = append(s.TagRequestList, tag)
	}

	return s
}

func (s *UpdateTagsRequest) WithTags(tags ...string) IUpdateTagsRequest {
	if len(tags)%2 != 0 {
		tags = append(tags, "none")
	}

	for i := 0; i < len(tags); i += 2 {
		s.TagRequestList = append(s.TagRequestList, common.Tag{Key: tags[i], Value: tags[i+1]})
	}

	return s
}

func (s *UpdateTagsRequest) ToMap() map[string]interface{} {
	res := make(map[string]interface{})
	for _, tag := range s.TagRequestList {
		res[tag.Key] = tag.Value
	}
	return res
}
