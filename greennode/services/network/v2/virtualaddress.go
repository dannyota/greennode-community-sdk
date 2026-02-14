package v2

import (
	"github.com/dannyota/greennode-community-sdk/v2/greennode/client"
	"github.com/dannyota/greennode-community-sdk/v2/greennode/entity"
	sdkerror "github.com/dannyota/greennode-community-sdk/v2/greennode/sdkerror"
)

func (s *NetworkServiceV2) CreateVirtualAddressCrossProject(popts ICreateVirtualAddressCrossProjectRequest) (*entity.VirtualAddress, sdkerror.IError) {
	url := createVirtualAddressCrossProjectUrl(s.VserverClient)
	resp := new(CreateVirtualAddressCrossProjectResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithHeader("User-Agent", popts.ParseUserAgent()).
		WithOkCodes(201).
		WithJsonBody(popts.ToRequestBody()).
		WithJsonResponse(resp).
		WithJsonError(errResp)

	if _, sdkErr := s.VserverClient.Post(url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.WithErrorSubnetNotFound(errResp),
			sdkerror.WithErrorVirtualAddressExceedQuota(errResp)).
			WithKVparameters(popts.ToMap()).
			WithErrorCategories(sdkerror.ErrCatVServer, sdkerror.ErrCatVirtualAddress)
	}

	return resp.ToEntityVirtualAddress(), nil
}

func (s *NetworkServiceV2) DeleteVirtualAddressById(popts IDeleteVirtualAddressByIdRequest) sdkerror.IError {
	url := deleteVirtualAddressByIdUrl(s.VserverClient, popts)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithHeader("User-Agent", popts.ParseUserAgent()).
		WithOkCodes(204).
		WithJsonError(errResp)

	if _, sdkErr := s.VserverClient.Delete(url, req); sdkErr != nil {
		return sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.WithErrorVirtualAddressNotFound(errResp),
			sdkerror.WithErrorVirtualAddressInUse(errResp)).
			WithKVparameters(popts.ToMap()).
			WithErrorCategories(sdkerror.ErrCatVServer, sdkerror.ErrCatVirtualAddress)
	}

	return nil
}

func (s *NetworkServiceV2) GetVirtualAddressById(popts IGetVirtualAddressByIdRequest) (*entity.VirtualAddress, sdkerror.IError) {
	url := getVirtualAddressByIdUrl(s.VserverClient, popts)
	resp := new(GetVirtualAddressByIdResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithHeader("User-Agent", popts.ParseUserAgent()).
		WithJsonResponse(resp).
		WithOkCodes(200).
		WithJsonError(errResp)

	if _, sdkErr := s.VserverClient.Get(url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.WithErrorVirtualAddressNotFound(errResp)).
			WithKVparameters(popts.ToMap()).
			WithErrorCategories(sdkerror.ErrCatVServer, sdkerror.ErrCatVirtualAddress)
	}

	return resp.ToEntityVirtualAddress(), nil
}

func (s *NetworkServiceV2) ListAddressPairsByVirtualAddressId(popts IListAddressPairsByVirtualAddressIdRequest) (*entity.ListAddressPairs, sdkerror.IError) {
	url := listAddressPairsByVirtualAddressIdUrl(s.VserverClient, popts)
	resp := new(ListAddressPairsByVirtualAddressIdResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithHeader("User-Agent", popts.ParseUserAgent()).
		WithJsonResponse(resp).
		WithOkCodes(200).
		WithJsonError(errResp)

	if _, sdkErr := s.VserverClient.Get(url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.WithErrorVirtualAddressNotFound(errResp)).
			WithKVparameters(popts.ToMap()).
			WithErrorCategories(sdkerror.ErrCatVServer, sdkerror.ErrCatVirtualAddress)
	}

	return resp.ToEntityListAddressPairs(), nil
}
