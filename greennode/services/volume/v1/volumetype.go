package v1

import (
	"context"

	"github.com/dannyota/greennode-community-sdk/v2/greennode/client"
	"github.com/dannyota/greennode-community-sdk/v2/greennode/entity"
	sdkerror "github.com/dannyota/greennode-community-sdk/v2/greennode/sdkerror"
)

func (s *VolumeServiceV1) GetVolumeTypeByID(ctx context.Context, opts *GetVolumeTypeByIDRequest) (*entity.VolumeType, error) {
	url := getVolumeTypeByIDURL(s.VServerClient, opts)
	resp := new(GetVolumeTypeByIDResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithOkCodes(200).
		WithJSONResponse(resp).
		WithJSONError(errResp)

	if _, sdkErr := s.VServerClient.Get(ctx, url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.EcVServerVolumeTypeNotFound).
			WithKVparameters(
				"projectId", s.getProjectID(),
				"volumeTypeId", opts.GetVolumeTypeID())
	}

	return resp.ToEntityVolumeType(), nil
}

func (s *VolumeServiceV1) GetDefaultVolumeType(ctx context.Context) (*entity.VolumeType, error) {
	url := getDefaultVolumeTypeURL(s.VServerClient)
	resp := new(GetDefaultVolumeTypeResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithOkCodes(200).
		WithJSONResponse(resp).
		WithJSONError(errResp)

	if _, sdkErr := s.VServerClient.Get(ctx, url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp).
			WithKVparameters(
				"projectId", s.getProjectID())
	}

	return resp.ToEntityVolumeType(), nil
}

func (s *VolumeServiceV1) GetVolumeTypeZones(ctx context.Context, opts *GetVolumeTypeZonesRequest) (*entity.ListVolumeTypeZones, error) {
	url := getVolumeTypeZonesURL(s.VServerClient, opts)
	resp := new(ListVolumeTypeZonesResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithOkCodes(200).
		WithJSONResponse(resp).
		WithJSONError(errResp)

	if _, sdkErr := s.VServerClient.Get(ctx, url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.EcVServerVolumeTypeNotFound).
			WithKVparameters("projectId", s.getProjectID())
	}

	return resp.ToEntityListVolumeTypeZones(), nil
}

func (s *VolumeServiceV1) GetListVolumeTypes(ctx context.Context, opts *GetListVolumeTypeRequest) (*entity.ListVolumeType, error) {
	url := getVolumeTypesURL(s.VServerClient, opts)
	resp := new(ListVolumeTypeResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithOkCodes(200).
		WithJSONResponse(resp).
		WithJSONError(errResp)

	if _, sdkErr := s.VServerClient.Get(ctx, url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.EcVServerVolumeTypeNotFound).
			WithKVparameters("projectId", s.getProjectID(),
				"volumeTypeZoneId", opts.GetVolumeTypeZoneID())
	}

	return resp.ToEntityListVolumeType(), nil
}
