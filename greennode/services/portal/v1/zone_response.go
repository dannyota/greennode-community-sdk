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

func (s *ListZoneResponse) ToEntityListZones() *entity.ListZones {
	listZones := &entity.ListZones{
		Items: make([]*entity.Zone, 0),
	}
	for _, q := range s.Data {
		listZones.Items = append(listZones.Items, q.ToEntityZone())
	}

	return listZones
}

func (s *Zone) ToEntityZone() *entity.Zone {

	return &entity.Zone{
		Uuid:          s.Uuid,
		Name:          s.Name,
		OpenstackZone: s.OpenstackZone,
	}
}
