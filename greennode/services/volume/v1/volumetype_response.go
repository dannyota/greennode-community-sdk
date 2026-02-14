package v1

import "github.com/dannyota/greennode-community-sdk/v2/greennode/entity"

type GetVolumeTypeByIDResponse struct {
	VolumeTypes []VolumeType `json:"volumeTypes"`
}

type GetDefaultVolumeTypeResponse struct {
	ID     string `json:"volumeTypeId"`
	ZoneID string `json:"volumeTypeZoneId"`
}

type (
	VolumeTypeZone struct {
		ID       string   `json:"id"`
		Name     string   `json:"name"`
		PoolName []string `json:"poolName,omitempty"`
	}

	VolumeType struct {
		ID         string `json:"id"`
		Iops       int    `json:"iops"`
		MaxSize    int    `json:"maxSize"`
		MinSize    int    `json:"minSize"`
		Name       string `json:"name"`
		ThroughPut int    `json:"throughPut,omitempty"`
		ZoneID     string `json:"zoneId,omitempty"`
	}
)

type ListVolumeTypeZonesResponse struct {
	VolumeTypeZones []VolumeTypeZone `json:"volumeTypeZones"`
}

type ListVolumeTypeResponse struct {
	VolumeTypes []VolumeType `json:"volumeTypes"`
}

func (s *GetVolumeTypeByIDResponse) ToEntityVolumeType() *entity.VolumeType {
	if len(s.VolumeTypes) == 0 {
		return nil
	}

	return s.VolumeTypes[0].toEntityVolumeType()
}

func (s *GetDefaultVolumeTypeResponse) ToEntityVolumeType() *entity.VolumeType {
	return &entity.VolumeType{
		ID:     s.ID,
		ZoneID: s.ZoneID,
	}
}

func (s VolumeType) toEntityVolumeType() *entity.VolumeType {
	return &entity.VolumeType{
		ID:         s.ID,
		Iops:       s.Iops,
		MaxSize:    s.MaxSize,
		MinSize:    s.MinSize,
		Name:       s.Name,
		ThroughPut: s.ThroughPut,
		ZoneID:     s.ZoneID,
	}
}

func (s VolumeTypeZone) toEntityVolumeTypeZone() *entity.VolumeTypeZone {
	return &entity.VolumeTypeZone{
		ID:       s.ID,
		Name:     s.Name,
		PoolName: s.PoolName,
	}
}

func (s *ListVolumeTypeZonesResponse) ToEntityListVolumeTypeZones() *entity.ListVolumeTypeZones {
	sl := new(entity.ListVolumeTypeZones)

	for _, item := range s.VolumeTypeZones {
		sl.VolumeTypeZones = append(sl.VolumeTypeZones, item.toEntityVolumeTypeZone())
	}

	return sl
}

func (s *ListVolumeTypeResponse) ToEntityListVolumeType() *entity.ListVolumeType {
	sl := new(entity.ListVolumeType)

	for _, item := range s.VolumeTypes {
		sl.VolumeTypes = append(sl.VolumeTypes, item.toEntityVolumeType())
	}

	return sl
}
