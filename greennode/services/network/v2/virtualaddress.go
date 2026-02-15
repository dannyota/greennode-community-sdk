package v2

import (
	"github.com/dannyota/greennode-community-sdk/v2/greennode/client"
	"github.com/dannyota/greennode-community-sdk/v2/greennode/entity"
	sdkerror "github.com/dannyota/greennode-community-sdk/v2/greennode/sdkerror"
)

func (s *NetworkServiceV2) CreateVirtualAddressCrossProject(opts *CreateVirtualAddressCrossProjectRequest) (*entity.VirtualAddress, error) {
	url := createVirtualAddressCrossProjectURL(s.VServerClient)
	resp := new(CreateVirtualAddressCrossProjectResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithHeader("User-Agent", opts.ParseUserAgent()).
		WithOkCodes(201).
		WithJSONBody(opts.ToRequestBody()).
		WithJSONResponse(resp).
		WithJSONError(errResp)

	if _, sdkErr := s.VServerClient.Post(url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.EcVServerSubnetNotFound,
			sdkerror.EcVServerVirtualAddressExceedQuota).
			WithKVparameters(opts.ToMap()).
			WithErrorCategories(sdkerror.ErrCatVServer, sdkerror.ErrCatVirtualAddress)
	}

	return resp.ToEntityVirtualAddress(), nil
}

func (s *NetworkServiceV2) DeleteVirtualAddressByID(opts *DeleteVirtualAddressByIDRequest) error {
	url := deleteVirtualAddressByIDURL(s.VServerClient, opts)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithHeader("User-Agent", opts.ParseUserAgent()).
		WithOkCodes(204).
		WithJSONError(errResp)

	if _, sdkErr := s.VServerClient.Delete(url, req); sdkErr != nil {
		return sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.EcVServerVirtualAddressNotFound,
			sdkerror.EcVServerVirtualAddressInUse).
			WithKVparameters(opts.ToMap()).
			WithErrorCategories(sdkerror.ErrCatVServer, sdkerror.ErrCatVirtualAddress)
	}

	return nil
}

func (s *NetworkServiceV2) GetVirtualAddressByID(opts *GetVirtualAddressByIDRequest) (*entity.VirtualAddress, error) {
	url := getVirtualAddressByIDURL(s.VServerClient, opts)
	resp := new(GetVirtualAddressByIDResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithHeader("User-Agent", opts.ParseUserAgent()).
		WithJSONResponse(resp).
		WithOkCodes(200).
		WithJSONError(errResp)

	if _, sdkErr := s.VServerClient.Get(url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.EcVServerVirtualAddressNotFound).
			WithKVparameters(opts.ToMap()).
			WithErrorCategories(sdkerror.ErrCatVServer, sdkerror.ErrCatVirtualAddress)
	}

	return resp.ToEntityVirtualAddress(), nil
}

func (s *NetworkServiceV2) ListAddressPairsByVirtualAddressID(opts *ListAddressPairsByVirtualAddressIDRequest) (*entity.ListAddressPairs, error) {
	url := listAddressPairsByVirtualAddressIDURL(s.VServerClient, opts)
	resp := new(ListAddressPairsByVirtualAddressIDResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithHeader("User-Agent", opts.ParseUserAgent()).
		WithJSONResponse(resp).
		WithOkCodes(200).
		WithJSONError(errResp)

	if _, sdkErr := s.VServerClient.Get(url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.EcVServerVirtualAddressNotFound).
			WithKVparameters(opts.ToMap()).
			WithErrorCategories(sdkerror.ErrCatVServer, sdkerror.ErrCatVirtualAddress)
	}

	return resp.ToEntityListAddressPairs(), nil
}
