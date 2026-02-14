package v1

import "github.com/dannyota/greennode-community-sdk/v2/greennode/entity"

type SystemTagResponse struct {
	Key       string `json:"key"`
	Value     string `json:"value"`
	CreatedAt string `json:"createdAt"`
	SystemTag bool   `json:"systemTag"`
}

func (s *SystemTagResponse) toSystemTag() entity.SystemTag {
	return entity.SystemTag{
		Key:       s.Key,
		Value:     s.Value,
		CreatedAt: s.CreatedAt,
		SystemTag: s.SystemTag,
	}
}
