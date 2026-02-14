package v2

import (
	"github.com/dannyota/greennode-community-sdk/v2/greennode/client"
	"github.com/dannyota/greennode-community-sdk/v2/greennode/entity"
	sdkerror "github.com/dannyota/greennode-community-sdk/v2/greennode/sdkerror"
)

func (s *NetworkServiceV2) CreateVirtualAddressCrossProject(opts ICreateVirtualAddressCrossProjectRequest) (*entity.VirtualAddress, sdkerror.Error) {
	url := createVirtualAddressCrossProjectUrl(s.VserverClient)
	resp := new(CreateVirtualAddressCrossProjectResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithHeader("User-Agent", opts.ParseUserAgent()).
		WithOkCodes(201).
		WithJsonBody(opts.ToRequestBody()).
		WithJsonResponse(resp).
		WithJsonError(errResp)

	if _, sdkErr := s.VserverClient.Post(url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.WithErrorSubnetNotFound(errResp),
			sdkerror.WithErrorVirtualAddressExceedQuota(errResp)).
			WithKVparameters(opts.ToMap()).
			WithErrorCategories(sdkerror.ErrCatVServer, sdkerror.ErrCatVirtualAddress)
	}

	return resp.ToEntityVirtualAddress(), nil
}

func (s *NetworkServiceV2) DeleteVirtualAddressById(opts IDeleteVirtualAddressByIdRequest) sdkerror.Error {
	url := deleteVirtualAddressByIdUrl(s.VserverClient, opts)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithHeader("User-Agent", opts.ParseUserAgent()).
		WithOkCodes(204).
		WithJsonError(errResp)

	if _, sdkErr := s.VserverClient.Delete(url, req); sdkErr != nil {
		return sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.WithErrorVirtualAddressNotFound(errResp),
			sdkerror.WithErrorVirtualAddressInUse(errResp)).
			WithKVparameters(opts.ToMap()).
			WithErrorCategories(sdkerror.ErrCatVServer, sdkerror.ErrCatVirtualAddress)
	}

	return nil
}

func (s *NetworkServiceV2) GetVirtualAddressById(opts IGetVirtualAddressByIdRequest) (*entity.VirtualAddress, sdkerror.Error) {
	url := getVirtualAddressByIdUrl(s.VserverClient, opts)
	resp := new(GetVirtualAddressByIdResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithHeader("User-Agent", opts.ParseUserAgent()).
		WithJsonResponse(resp).
		WithOkCodes(200).
		WithJsonError(errResp)

	if _, sdkErr := s.VserverClient.Get(url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.WithErrorVirtualAddressNotFound(errResp)).
			WithKVparameters(opts.ToMap()).
			WithErrorCategories(sdkerror.ErrCatVServer, sdkerror.ErrCatVirtualAddress)
	}

	return resp.ToEntityVirtualAddress(), nil
}

func (s *NetworkServiceV2) ListAddressPairsByVirtualAddressId(opts IListAddressPairsByVirtualAddressIdRequest) (*entity.ListAddressPairs, sdkerror.Error) {
	url := listAddressPairsByVirtualAddressIdUrl(s.VserverClient, opts)
	resp := new(ListAddressPairsByVirtualAddressIdResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithHeader("User-Agent", opts.ParseUserAgent()).
		WithJsonResponse(resp).
		WithOkCodes(200).
		WithJsonError(errResp)

	if _, sdkErr := s.VserverClient.Get(url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.WithErrorVirtualAddressNotFound(errResp)).
			WithKVparameters(opts.ToMap()).
			WithErrorCategories(sdkerror.ErrCatVServer, sdkerror.ErrCatVirtualAddress)
	}

	return resp.ToEntityListAddressPairs(), nil
}
