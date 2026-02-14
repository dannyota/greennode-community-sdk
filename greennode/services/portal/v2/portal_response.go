package v2

import (
	"strconv"

	"github.com/dannyota/greennode-community-sdk/v2/greennode/entity"
)

const (
	QtVolumeAttachLimit = QuotaName("VOLUME_PER_SERVER") // The maximum number of volumes that you can attach to an server
)

type (
	Quota struct {
		Description string    `json:"description,omitempty"`
		Name        QuotaName `json:"quotaName"`
		Type        QuotaType `json:"type"`
		Used        string    `json:"used"`
		Limit       int       `json:"limit"`
	}

	QuotaType string
	QuotaName string
)

type ListAllQuotaUsedResponse struct {
	Data []Quota `json:"data"`
}

func (s *ListAllQuotaUsedResponse) ToEntityListQuotas() *entity.ListQuotas {
	listQuotas := &entity.ListQuotas{
		Items: make([]*entity.Quota, 0),
	}
	for _, q := range s.Data {
		listQuotas.Items = append(listQuotas.Items, q.ToEntityQuota())
	}

	return listQuotas
}

func (s *Quota) ToEntityQuota() *entity.Quota {
	var (
		used int
		err  error
	)

	if used, err = strconv.Atoi(s.Used); err != nil {
		used = 0
	}

	return &entity.Quota{
		Description: s.Description,
		Name:        string(s.Name),
		Type:        string(s.Type),
		Limit:       s.Limit,
		Used:        used,
	}
}
