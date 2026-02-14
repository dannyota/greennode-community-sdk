package v1

import (
	"github.com/dannyota/greennode-community-sdk/v2/greennode/entity"
)

type Zone struct {
	Uuid          string `json:"uuid"`
	Name          string `json:"name"`
	OpenstackZone string `json:"openstackZone"`
}

type ListZoneResponse struct {
	Data []Zone `json:"data"`
}

func (r *ListZoneResponse) ToEntityListZones() *entity.ListZones {
	listZones := &entity.ListZones{
		Items: make([]*entity.Zone, 0),
	}
	for _, q := range r.Data {
		listZones.Items = append(listZones.Items, q.ToEntityZone())
	}

	return listZones
}

func (z *Zone) ToEntityZone() *entity.Zone {

	return &entity.Zone{
		Uuid:          z.Uuid,
		Name:          z.Name,
		OpenstackZone: z.OpenstackZone,
	}
}
