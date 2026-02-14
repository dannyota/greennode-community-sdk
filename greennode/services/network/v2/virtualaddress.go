package v2

import (
	"github.com/dannyota/greennode-community-sdk/v2/greennode/client"
	"github.com/dannyota/greennode-community-sdk/v2/greennode/entity"
	sdkerror "github.com/dannyota/greennode-community-sdk/v2/greennode/sdkerror"
)

func (s *NetworkServiceV2) CreateVirtualAddressCrossProject(opts *CreateVirtualAddressCrossProjectRequest) (*entity.VirtualAddress, sdkerror.Error) {
	url := createVirtualAddressCrossProjectURL(s.VserverClient)
	resp := new(CreateVirtualAddressCrossProjectResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithHeader("User-Agent", opts.ParseUserAgent()).
		WithOkCodes(201).
		WithJSONBody(opts.ToRequestBody()).
		WithJSONResponse(resp).
		WithJSONError(errResp)

	if _, sdkErr := s.VserverClient.Post(url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.WithErrorSubnetNotFound(errResp),
			sdkerror.WithErrorVirtualAddressExceedQuota(errResp)).
			WithKVparameters(opts.ToMap()).
			WithErrorCategories(sdkerror.ErrCatVServer, sdkerror.ErrCatVirtualAddress)
	}

	return resp.ToEntityVirtualAddress(), nil
}

func (s *NetworkServiceV2) DeleteVirtualAddressByID(opts *DeleteVirtualAddressByIDRequest) sdkerror.Error {
	url := deleteVirtualAddressByIDURL(s.VserverClient, opts)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithHeader("User-Agent", opts.ParseUserAgent()).
		WithOkCodes(204).
		WithJSONError(errResp)

	if _, sdkErr := s.VserverClient.Delete(url, req); sdkErr != nil {
		return sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.WithErrorVirtualAddressNotFound(errResp),
			sdkerror.WithErrorVirtualAddressInUse(errResp)).
			WithKVparameters(opts.ToMap()).
			WithErrorCategories(sdkerror.ErrCatVServer, sdkerror.ErrCatVirtualAddress)
	}

	return nil
}

func (s *NetworkServiceV2) GetVirtualAddressByID(opts *GetVirtualAddressByIDRequest) (*entity.VirtualAddress, sdkerror.Error) {
	url := getVirtualAddressByIDURL(s.VserverClient, opts)
	resp := new(GetVirtualAddressByIDResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithHeader("User-Agent", opts.ParseUserAgent()).
		WithJSONResponse(resp).
		WithOkCodes(200).
		WithJSONError(errResp)

	if _, sdkErr := s.VserverClient.Get(url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.WithErrorVirtualAddressNotFound(errResp)).
			WithKVparameters(opts.ToMap()).
			WithErrorCategories(sdkerror.ErrCatVServer, sdkerror.ErrCatVirtualAddress)
	}

	return resp.ToEntityVirtualAddress(), nil
}

func (s *NetworkServiceV2) ListAddressPairsByVirtualAddressID(opts *ListAddressPairsByVirtualAddressIDRequest) (*entity.ListAddressPairs, sdkerror.Error) {
	url := listAddressPairsByVirtualAddressIDURL(s.VserverClient, opts)
	resp := new(ListAddressPairsByVirtualAddressIDResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithHeader("User-Agent", opts.ParseUserAgent()).
		WithJSONResponse(resp).
		WithOkCodes(200).
		WithJSONError(errResp)

	if _, sdkErr := s.VserverClient.Get(url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.WithErrorVirtualAddressNotFound(errResp)).
			WithKVparameters(opts.ToMap()).
			WithErrorCategories(sdkerror.ErrCatVServer, sdkerror.ErrCatVirtualAddress)
	}

	return resp.ToEntityListAddressPairs(), nil
}
