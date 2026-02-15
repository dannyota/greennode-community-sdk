package v1

import "github.com/dannyota/greennode-community-sdk/v2/greennode/services/common"

const (
	Volume       ResourceType = "VOLUME"
	Server       ResourceType = "SERVER"
	LoadBalancer ResourceType = "LOAD-BALANCER"
)

type ResourceType string
type CreateSystemTagRequest struct { // __________________________________________________________________________________
	ResourceID   string       `json:"resourceId"`
	ResourceType ResourceType `json:"resourceType"`
	Tags         []struct {
		Key   string `json:"key"`
		Value string `json:"value"`
	} `json:"tagRequestList"`
	common.UserAgent
}

func (r *CreateSystemTagRequest) AddUserAgent(agent ...string) *CreateSystemTagRequest {
	r.UserAgent.AddUserAgent(agent...)
	return r
}

func (r *CreateSystemTagRequest) GetResourceID() string {
	return r.ResourceID
}

func (r *CreateSystemTagRequest) GetResourceType() ResourceType {
	return r.ResourceType
}

func (r *CreateSystemTagRequest) AddTag(key, value string) *CreateSystemTagRequest {
	r.Tags = append(r.Tags, struct {
		Key   string `json:"key"`
		Value string `json:"value"`
	}{
		Key:   key,
		Value: value,
	})

	return r
}

