package v1

import "github.com/dannyota/greennode-community-sdk/v2/greennode/services/common"

const (
	Volume       ResourceType = "VOLUME"
	Server       ResourceType = "SERVER"
	LoadBalancer ResourceType = "LOAD-BALANCER"
)

type ResourceType string
type CreateSystemTagRequest struct { // __________________________________________________________________________________
	ResourceId   string       `json:"resourceId"`
	ResourceType ResourceType `json:"resourceType"`
	Tags         []struct {
		Key   string `json:"key"`
		Value string `json:"value"`
	} `json:"tagRequestList"`
	common.UserAgent
}

func (s *CreateSystemTagRequest) ToRequestBody() interface{} {
	return s
}

func (s *CreateSystemTagRequest) AddUserAgent(agent ...string) ICreateSystemTagRequest {
	s.UserAgent.AddUserAgent(agent...)
	return s
}

func (s *CreateSystemTagRequest) GetResourceId() string {
	return s.ResourceId
}

func (s *CreateSystemTagRequest) GetResourceType() ResourceType {
	return s.ResourceType
}

func (s *CreateSystemTagRequest) AddTag(key, value string) ICreateSystemTagRequest {
	s.Tags = append(s.Tags, struct {
		Key   string `json:"key"`
		Value string `json:"value"`
	}{
		Key:   key,
		Value: value,
	})

	return s
}

func (s *CreateSystemTagRequest) ToMap() map[string]interface{} {
	res := map[string]interface{}{
		"resourceId":   s.ResourceId,
		"resourceType": s.ResourceType,
	}

	if len(s.Agent) > 0 {
		res["userAgent"] = s.Agent
	}

	res["tags"] = s.Tags

	return res
}
