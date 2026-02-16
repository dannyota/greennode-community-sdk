package v2

import (
	"context"

	"github.com/dannyota/greennode-community-sdk/v2/greennode/client"
	sdkerror "github.com/dannyota/greennode-community-sdk/v2/greennode/sdkerror"
	"github.com/dannyota/greennode-community-sdk/v2/greennode/services/common"
)

func (s *NetworkServiceV2) CreateVirtualAddressCrossProject(ctx context.Context, opts *CreateVirtualAddressCrossProjectRequest) (*VirtualAddress, error) {
	url := createVirtualAddressCrossProjectURL(s.Client)
	resp := new(CreateVirtualAddressCrossProjectResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithOkCodes(201).
		WithJSONBody(opts).
		WithJSONResponse(resp).
		WithJSONError(errResp)

	if _, sdkErr := s.Client.Post(ctx, url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.EcVServerSubnetNotFound,
			sdkerror.EcVServerVirtualAddressExceedQuota).
			WithKVparameters(common.StructToMap(opts)).
			WithErrorCategories(sdkerror.ErrCatVServer, sdkerror.ErrCatVirtualAddress)
	}

	return resp.ToEntityVirtualAddress(), nil
}

func (s *NetworkServiceV2) DeleteVirtualAddressByID(ctx context.Context, opts *DeleteVirtualAddressByIDRequest) error {
	url := deleteVirtualAddressByIDURL(s.Client, opts)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithOkCodes(204).
		WithJSONError(errResp)

	if _, sdkErr := s.Client.Delete(ctx, url, req); sdkErr != nil {
		return sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.EcVServerVirtualAddressNotFound,
			sdkerror.EcVServerVirtualAddressInUse).
			WithKVparameters(common.StructToMap(opts)).
			WithErrorCategories(sdkerror.ErrCatVServer, sdkerror.ErrCatVirtualAddress)
	}

	return nil
}

func (s *NetworkServiceV2) GetVirtualAddressByID(ctx context.Context, opts *GetVirtualAddressByIDRequest) (*VirtualAddress, error) {
	url := getVirtualAddressByIDURL(s.Client, opts)
	resp := new(GetVirtualAddressByIDResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithJSONResponse(resp).
		WithOkCodes(200).
		WithJSONError(errResp)

	if _, sdkErr := s.Client.Get(ctx, url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.EcVServerVirtualAddressNotFound).
			WithKVparameters(common.StructToMap(opts)).
			WithErrorCategories(sdkerror.ErrCatVServer, sdkerror.ErrCatVirtualAddress)
	}

	return resp.ToEntityVirtualAddress(), nil
}

func (s *NetworkServiceV2) ListAddressPairsByVirtualAddressID(ctx context.Context, opts *ListAddressPairsByVirtualAddressIDRequest) (*ListAddressPairs, error) {
	url := listAddressPairsByVirtualAddressIDURL(s.Client, opts)
	resp := new(ListAddressPairsByVirtualAddressIDResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithJSONResponse(resp).
		WithOkCodes(200).
		WithJSONError(errResp)

	if _, sdkErr := s.Client.Get(ctx, url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.EcVServerVirtualAddressNotFound).
			WithKVparameters(common.StructToMap(opts)).
			WithErrorCategories(sdkerror.ErrCatVServer, sdkerror.ErrCatVirtualAddress)
	}

	return resp.ToEntityListAddressPairs(), nil
}
