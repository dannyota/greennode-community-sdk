package v1

import (
	"fmt"
	"net/url"
)

func NewGetVolumeTypeByIDRequest(volumeTypeID string) *GetVolumeTypeByIDRequest {
	return &GetVolumeTypeByIDRequest{
		VolumeTypeID: volumeTypeID,
	}
}

func NewListVolumeTypeRequest(volumeTypeZoneID string) *GetListVolumeTypeRequest {
	return &GetListVolumeTypeRequest{
		VolumeTypeZoneID: volumeTypeZoneID,
	}
}

func NewGetVolumeTypeZonesRequest(zoneID string) *GetVolumeTypeZonesRequest {
	return &GetVolumeTypeZonesRequest{
		ZoneID: zoneID,
	}
}

type GetVolumeTypeByIDRequest struct {
	VolumeTypeID string
}

type GetVolumeTypeZonesRequest struct {
	ZoneID string
}

type GetListVolumeTypeRequest struct {
	VolumeTypeZoneID string
}

func (r *GetVolumeTypeZonesRequest) getDefaultQuery() string {
	return fmt.Sprintf("zoneId=%s", defaultZoneGetVolumeTypeZonesRequest)
}

func (r *GetVolumeTypeZonesRequest) ToQuery() (string, error) {
	v := url.Values{}
	if r.ZoneID != "" {
		v.Set("zoneId", r.ZoneID)
	}
	return v.Encode(), nil
}

