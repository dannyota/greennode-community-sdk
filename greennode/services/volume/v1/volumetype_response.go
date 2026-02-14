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

func (r *GetVolumeTypeByIDResponse) ToEntityVolumeType() *entity.VolumeType {
	if len(r.VolumeTypes) == 0 {
		return nil
	}

	return r.VolumeTypes[0].toEntityVolumeType()
}

func (r *GetDefaultVolumeTypeResponse) ToEntityVolumeType() *entity.VolumeType {
	return &entity.VolumeType{
		ID:     r.ID,
		ZoneID: r.ZoneID,
	}
}

func (v VolumeType) toEntityVolumeType() *entity.VolumeType {
	return &entity.VolumeType{
		ID:         v.ID,
		Iops:       v.Iops,
		MaxSize:    v.MaxSize,
		MinSize:    v.MinSize,
		Name:       v.Name,
		ThroughPut: v.ThroughPut,
		ZoneID:     v.ZoneID,
	}
}

func (v VolumeTypeZone) toEntityVolumeTypeZone() *entity.VolumeTypeZone {
	return &entity.VolumeTypeZone{
		ID:       v.ID,
		Name:     v.Name,
		PoolName: v.PoolName,
	}
}

func (r *ListVolumeTypeZonesResponse) ToEntityListVolumeTypeZones() *entity.ListVolumeTypeZones {
	sl := new(entity.ListVolumeTypeZones)

	for _, item := range r.VolumeTypeZones {
		sl.VolumeTypeZones = append(sl.VolumeTypeZones, item.toEntityVolumeTypeZone())
	}

	return sl
}

func (r *ListVolumeTypeResponse) ToEntityListVolumeType() *entity.ListVolumeType {
	sl := new(entity.ListVolumeType)

	for _, item := range r.VolumeTypes {
		sl.VolumeTypes = append(sl.VolumeTypes, item.toEntityVolumeType())
	}

	return sl
}
