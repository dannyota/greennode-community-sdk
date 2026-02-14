package v1

import (
	"github.com/dannyota/greennode-community-sdk/v2/greennode/client"
	"github.com/dannyota/greennode-community-sdk/v2/greennode/entity"
	sdkerror "github.com/dannyota/greennode-community-sdk/v2/greennode/sdkerror"
)

func (s *VolumeServiceV1) GetVolumeTypeById(popts IGetVolumeTypeByIdRequest) (*entity.VolumeType, sdkerror.Error) {
	url := getVolumeTypeByIdUrl(s.VServerClient, popts)
	resp := new(GetVolumeTypeByIdResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithOkCodes(200).
		WithJsonResponse(resp).
		WithJsonError(errResp)

	if _, sdkErr := s.VServerClient.Get(url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.WithErrorVolumeTypeNotFound(errResp)).
			WithKVparameters(
				"projectId", s.getProjectId(),
				"volumeTypeId", popts.GetVolumeTypeId())
	}

	return resp.ToEntityVolumeType(), nil
}

func (s *VolumeServiceV1) GetDefaultVolumeType() (*entity.VolumeType, sdkerror.Error) {
	url := getDefaultVolumeTypeUrl(s.VServerClient)
	resp := new(GetDefaultVolumeTypeResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithOkCodes(200).
		WithJsonResponse(resp).
		WithJsonError(errResp)

	if _, sdkErr := s.VServerClient.Get(url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp).
			WithKVparameters(
				"projectId", s.getProjectId())
	}

	return resp.ToEntityVolumeType(), nil
}

func (s *VolumeServiceV1) GetVolumeTypeZones(popts IGetVolumeTypeZonesRequest) (*entity.ListVolumeTypeZones, sdkerror.Error) {
	url := getVolumeTypeZonesUrl(s.VServerClient, popts)
	resp := new(ListVolumeTypeZonesResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithOkCodes(200).
		WithJsonResponse(resp).
		WithJsonError(errResp)

	if _, sdkErr := s.VServerClient.Get(url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.WithErrorVolumeTypeNotFound(errResp)).
			WithKVparameters("projectId", s.getProjectId())
	}

	return resp.ToEntityListVolumeTypeZones(), nil

}

func (s *VolumeServiceV1) GetListVolumeTypes(popts IGetListVolumeTypeRequest) (*entity.ListVolumeType, sdkerror.Error) {
	url := getVolumeTypesUrl(s.VServerClient, popts)
	resp := new(ListVolumeTypeResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithOkCodes(200).
		WithJsonResponse(resp).
		WithJsonError(errResp)

	if _, sdkErr := s.VServerClient.Get(url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.WithErrorVolumeTypeNotFound(errResp)).
			WithKVparameters("projectId", s.getProjectId(),
				"volumeTypeZoneId", popts.GetVolumeTypeZoneId())
	}

	return resp.ToEntityListVolumeType(), nil

}
