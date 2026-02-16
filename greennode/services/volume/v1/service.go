package v1

import (
	"context"

	"github.com/dannyota/greennode-community-sdk/v2/greennode/client"
	sdkerror "github.com/dannyota/greennode-community-sdk/v2/greennode/sdkerror"
)

const (
	defaultZoneGetVolumeTypeZonesRequest = "HCM03-1A"
)

type VolumeServiceV1 struct {
	Client *client.ServiceClient
}


func (s *VolumeServiceV1) GetVolumeTypeByID(ctx context.Context, opts *GetVolumeTypeByIDRequest) (*VolumeType, error) {
	url := getVolumeTypeByIDURL(s.Client, opts)
	resp := new(GetVolumeTypeByIDResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithOkCodes(200).
		WithJSONResponse(resp).
		WithJSONError(errResp)

	if _, sdkErr := s.Client.Get(ctx, url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.EcVServerVolumeTypeNotFound).
			WithKVparameters(
				"projectId", s.Client.ProjectID,
				"volumeTypeId", opts.VolumeTypeID)
	}

	return resp.ToEntityVolumeType(), nil
}

func (s *VolumeServiceV1) GetDefaultVolumeType(ctx context.Context) (*VolumeType, error) {
	url := getDefaultVolumeTypeURL(s.Client)
	resp := new(GetDefaultVolumeTypeResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithOkCodes(200).
		WithJSONResponse(resp).
		WithJSONError(errResp)

	if _, sdkErr := s.Client.Get(ctx, url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp).
			WithKVparameters(
				"projectId", s.Client.ProjectID)
	}

	return resp.ToEntityVolumeType(), nil
}

func (s *VolumeServiceV1) GetVolumeTypeZones(ctx context.Context, opts *GetVolumeTypeZonesRequest) (*ListVolumeTypeZones, error) {
	url := getVolumeTypeZonesURL(s.Client, opts)
	resp := new(ListVolumeTypeZonesResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithOkCodes(200).
		WithJSONResponse(resp).
		WithJSONError(errResp)

	if _, sdkErr := s.Client.Get(ctx, url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.EcVServerVolumeTypeNotFound).
			WithKVparameters("projectId", s.Client.ProjectID)
	}

	return resp.ToEntityListVolumeTypeZones(), nil
}

func (s *VolumeServiceV1) GetListVolumeTypes(ctx context.Context, opts *GetListVolumeTypeRequest) (*ListVolumeTypes, error) {
	url := getVolumeTypesURL(s.Client, opts)
	resp := new(ListVolumeTypeResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithOkCodes(200).
		WithJSONResponse(resp).
		WithJSONError(errResp)

	if _, sdkErr := s.Client.Get(ctx, url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.EcVServerVolumeTypeNotFound).
			WithKVparameters("projectId", s.Client.ProjectID,
				"volumeTypeZoneId", opts.VolumeTypeZoneID)
	}

	return resp.ToEntityListVolumeType(), nil
}
