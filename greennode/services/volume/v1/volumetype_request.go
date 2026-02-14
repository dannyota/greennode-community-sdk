package v1

import (
	"fmt"
	"net/url"

	"github.com/dannyota/greennode-community-sdk/v2/greennode/services/common"
)

func NewGetVolumeTypeByIdRequest(pvolumeTypeId string) IGetVolumeTypeByIdRequest {
	opt := new(GetVolumeTypeByIdRequest)
	opt.VolumeTypeId = pvolumeTypeId
	return opt
}

func NewListVolumeTypeRequest(volumeTypeZoneId string) IGetListVolumeTypeRequest {
	opt := new(GetListVolumeTypeRequest)
	opt.VolumeTypeZoneId = volumeTypeZoneId
	return opt
}

func NewGetVolumeTypeZonesRequest(zoneId string) IGetVolumeTypeZonesRequest {
	opt := new(GetVolumeTypeZonesRequest)
	opt.ZoneId = zoneId
	return opt
}

type GetVolumeTypeByIdRequest struct {
	common.VolumeTypeCommon
}

type GetVolumeTypeZonesRequest struct {
	ZoneId string
}

type GetListVolumeTypeRequest struct {
	VolumeTypeZoneId string
}

func (s *GetVolumeTypeZonesRequest) GetDefaultQuery() string {
	return fmt.Sprintf("zoneId=%s", defaultZoneGetVolumeTypeZonesRequest)
}

func (s *GetVolumeTypeZonesRequest) ToQuery() (string, error) {
	v := url.Values{}
	if s.ZoneId != "" {
		v.Set("zoneId", s.ZoneId)
	}
	return v.Encode(), nil
}

func (s *GetListVolumeTypeRequest) GetVolumeTypeZoneId() string {
	return s.VolumeTypeZoneId
}
