package v1

type GetVolumeTypeByIDResponse struct {
	VolumeTypes []VolumeType `json:"volumeTypes"`
}

type GetDefaultVolumeTypeResponse struct {
	ID     string `json:"volumeTypeId"`
	ZoneID string `json:"volumeTypeZoneId"`
}

type ListVolumeTypeZonesResponse struct {
	VolumeTypeZones []VolumeTypeZone `json:"volumeTypeZones"`
}

type ListVolumeTypeResponse struct {
	VolumeTypes []VolumeType `json:"volumeTypes"`
}

func (r *GetVolumeTypeByIDResponse) ToEntityVolumeType() *VolumeType {
	if len(r.VolumeTypes) == 0 {
		return nil
	}

	return &r.VolumeTypes[0]
}

func (r *GetDefaultVolumeTypeResponse) ToEntityVolumeType() *VolumeType {
	return &VolumeType{
		ID:     r.ID,
		ZoneID: r.ZoneID,
	}
}

func (r *ListVolumeTypeZonesResponse) ToEntityListVolumeTypeZones() *ListVolumeTypeZones {
	sl := new(ListVolumeTypeZones)

	for i := range r.VolumeTypeZones {
		sl.VolumeTypeZones = append(sl.VolumeTypeZones, &r.VolumeTypeZones[i])
	}

	return sl
}

func (r *ListVolumeTypeResponse) ToEntityListVolumeType() *ListVolumeTypes {
	sl := new(ListVolumeTypes)

	for i := range r.VolumeTypes {
		sl.VolumeTypes = append(sl.VolumeTypes, &r.VolumeTypes[i])
	}

	return sl
}
