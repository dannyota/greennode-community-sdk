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

func (r *ListAllQuotaUsedResponse) ToEntityListQuotas() *entity.ListQuotas {
	listQuotas := &entity.ListQuotas{
		Items: make([]*entity.Quota, 0),
	}
	for _, q := range r.Data {
		listQuotas.Items = append(listQuotas.Items, q.ToEntityQuota())
	}

	return listQuotas
}

func (q *Quota) ToEntityQuota() *entity.Quota {
	var (
		used int
		err  error
	)

	if used, err = strconv.Atoi(q.Used); err != nil {
		used = 0
	}

	return &entity.Quota{
		Description: q.Description,
		Name:        string(q.Name),
		Type:        string(q.Type),
		Limit:       q.Limit,
		Used:        used,
	}
}
