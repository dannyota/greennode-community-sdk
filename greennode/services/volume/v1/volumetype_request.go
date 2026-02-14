package v1

import (
	"fmt"
	"net/url"

	"github.com/dannyota/greennode-community-sdk/v2/greennode/services/common"
)

func NewGetVolumeTypeByIDRequest(volumeTypeID string) IGetVolumeTypeByIDRequest {
	opt := new(GetVolumeTypeByIDRequest)
	opt.VolumeTypeID = volumeTypeID
	return opt
}

func NewListVolumeTypeRequest(volumeTypeZoneID string) IGetListVolumeTypeRequest {
	opt := new(GetListVolumeTypeRequest)
	opt.VolumeTypeZoneID = volumeTypeZoneID
	return opt
}

func NewGetVolumeTypeZonesRequest(zoneID string) IGetVolumeTypeZonesRequest {
	opt := new(GetVolumeTypeZonesRequest)
	opt.ZoneID = zoneID
	return opt
}

type GetVolumeTypeByIDRequest struct {
	common.VolumeTypeCommon
}

type GetVolumeTypeZonesRequest struct {
	ZoneID string
}

type GetListVolumeTypeRequest struct {
	VolumeTypeZoneID string
}

func (r *GetVolumeTypeZonesRequest) GetDefaultQuery() string {
	return fmt.Sprintf("zoneId=%s", defaultZoneGetVolumeTypeZonesRequest)
}

func (r *GetVolumeTypeZonesRequest) ToQuery() (string, error) {
	v := url.Values{}
	if r.ZoneID != "" {
		v.Set("zoneId", r.ZoneID)
	}
	return v.Encode(), nil
}

func (r *GetListVolumeTypeRequest) GetVolumeTypeZoneID() string {
	return r.VolumeTypeZoneID
}
