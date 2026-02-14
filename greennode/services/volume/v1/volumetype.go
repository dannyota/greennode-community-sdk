package v1

import (
	"github.com/dannyota/greennode-community-sdk/v2/greennode/client"
	"github.com/dannyota/greennode-community-sdk/v2/greennode/entity"
	sdkerror "github.com/dannyota/greennode-community-sdk/v2/greennode/sdkerror"
)

func (s *VolumeServiceV1) GetVolumeTypeByID(opts IGetVolumeTypeByIDRequest) (*entity.VolumeType, sdkerror.Error) {
	url := getVolumeTypeByIDURL(s.VServerClient, opts)
	resp := new(GetVolumeTypeByIDResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithOkCodes(200).
		WithJSONResponse(resp).
		WithJSONError(errResp)

	if _, sdkErr := s.VServerClient.Get(url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.WithErrorVolumeTypeNotFound(errResp)).
			WithKVparameters(
				"projectId", s.getProjectID(),
				"volumeTypeId", opts.GetVolumeTypeID())
	}

	return resp.ToEntityVolumeType(), nil
}

func (s *VolumeServiceV1) GetDefaultVolumeType() (*entity.VolumeType, sdkerror.Error) {
	url := getDefaultVolumeTypeURL(s.VServerClient)
	resp := new(GetDefaultVolumeTypeResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithOkCodes(200).
		WithJSONResponse(resp).
		WithJSONError(errResp)

	if _, sdkErr := s.VServerClient.Get(url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp).
			WithKVparameters(
				"projectId", s.getProjectID())
	}

	return resp.ToEntityVolumeType(), nil
}

func (s *VolumeServiceV1) GetVolumeTypeZones(opts IGetVolumeTypeZonesRequest) (*entity.ListVolumeTypeZones, sdkerror.Error) {
	url := getVolumeTypeZonesURL(s.VServerClient, opts)
	resp := new(ListVolumeTypeZonesResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithOkCodes(200).
		WithJSONResponse(resp).
		WithJSONError(errResp)

	if _, sdkErr := s.VServerClient.Get(url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.WithErrorVolumeTypeNotFound(errResp)).
			WithKVparameters("projectId", s.getProjectID())
	}

	return resp.ToEntityListVolumeTypeZones(), nil

}

func (s *VolumeServiceV1) GetListVolumeTypes(opts IGetListVolumeTypeRequest) (*entity.ListVolumeType, sdkerror.Error) {
	url := getVolumeTypesURL(s.VServerClient, opts)
	resp := new(ListVolumeTypeResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithOkCodes(200).
		WithJSONResponse(resp).
		WithJSONError(errResp)

	if _, sdkErr := s.VServerClient.Get(url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.WithErrorVolumeTypeNotFound(errResp)).
			WithKVparameters("projectId", s.getProjectID(),
				"volumeTypeZoneId", opts.GetVolumeTypeZoneID())
	}

	return resp.ToEntityListVolumeType(), nil

}
