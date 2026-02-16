package v1

import "github.com/dannyota/greennode-community-sdk/v2/greennode/entity"

type GetVolumeTypeByIDResponse struct {
	VolumeTypes []entity.VolumeType `json:"volumeTypes"`
}

type GetDefaultVolumeTypeResponse struct {
	ID     string `json:"volumeTypeId"`
	ZoneID string `json:"volumeTypeZoneId"`
}

type ListVolumeTypeZonesResponse struct {
	VolumeTypeZones []entity.VolumeTypeZone `json:"volumeTypeZones"`
}

type ListVolumeTypeResponse struct {
	VolumeTypes []entity.VolumeType `json:"volumeTypes"`
}

func (r *GetVolumeTypeByIDResponse) ToEntityVolumeType() *entity.VolumeType {
	if len(r.VolumeTypes) == 0 {
		return nil
	}

	return &r.VolumeTypes[0]
}

func (r *GetDefaultVolumeTypeResponse) ToEntityVolumeType() *entity.VolumeType {
	return &entity.VolumeType{
		ID:     r.ID,
		ZoneID: r.ZoneID,
	}
}

func (r *ListVolumeTypeZonesResponse) ToEntityListVolumeTypeZones() *entity.ListVolumeTypeZones {
	sl := new(entity.ListVolumeTypeZones)

	for i := range r.VolumeTypeZones {
		sl.VolumeTypeZones = append(sl.VolumeTypeZones, &r.VolumeTypeZones[i])
	}

	return sl
}

func (r *ListVolumeTypeResponse) ToEntityListVolumeType() *entity.ListVolumeType {
	sl := new(entity.ListVolumeType)

	for i := range r.VolumeTypes {
		sl.VolumeTypes = append(sl.VolumeTypes, &r.VolumeTypes[i])
	}

	return sl
}
