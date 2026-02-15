package v2

import (
	"context"

	"github.com/dannyota/greennode-community-sdk/v2/greennode/client"
	"github.com/dannyota/greennode-community-sdk/v2/greennode/entity"
	sdkerror "github.com/dannyota/greennode-community-sdk/v2/greennode/sdkerror"
)

func (s *NetworkServiceV2) GetAllAddressPairByVirtualSubnetID(ctx context.Context, opts *GetAllAddressPairByVirtualSubnetIDRequest) ([]*entity.AddressPair, error) {
	url := getAllAddressPairByVirtualSubnetIDURL(s.VServerClient, opts)
	resp := new(GetAllAddressPairByVirtualSubnetIDResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithOkCodes(200).
		WithJSONResponse(resp).
		WithJSONError(errResp)

	if _, sdkErr := s.VServerClient.Get(ctx, url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp)
	}
	return resp.ToListAddressPair(), nil
}

func (s *NetworkServiceV2) SetAddressPairInVirtualSubnet(ctx context.Context, opts *SetAddressPairInVirtualSubnetRequest) (*entity.AddressPair, error) {
	url := setAddressPairInVirtualSubnetURL(s.VServerClient, opts)
	resp := new(SetAddressPairInVirtualSubnetResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithOkCodes(200, 201, 202, 203, 204).
		WithJSONBody(opts.AddressPairRequest).
		WithJSONResponse(resp).
		WithJSONError(errResp)

	if _, sdkErr := s.VServerClient.Post(ctx, url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp)
	}
	return resp.ToAddressPair(), nil
}

func (s *NetworkServiceV2) DeleteAddressPair(ctx context.Context, opts *DeleteAddressPairRequest) error {
	url := deleteAddressPairURL(s.VServerClient, opts)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithOkCodes(200, 201, 202, 203, 204).
		WithJSONError(errResp)

	if _, sdkErr := s.VServerClient.Delete(ctx, url, req); sdkErr != nil {
		return sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.EcVServerVirtualAddressNotFound).
			WithKVparameters("addressPairId", opts.GetAddressPairID())
	}
	return nil
}

func (s *NetworkServiceV2) CreateAddressPair(ctx context.Context, opts *CreateAddressPairRequest) (*entity.AddressPair, error) {
	url := createAddressPairURL(s.VServerClient, opts)
	resp := new(CreateAddressPairResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithOkCodes(201).
		WithJSONBody(opts).
		WithJSONResponse(resp).
		WithJSONError(errResp)

	if _, sdkErr := s.VServerClient.Post(ctx, url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.EcVServerInternalNetworkInterfaceNotFound,
			sdkerror.EcVServerAddressPairExisted).
			WithErrorCategories(sdkerror.ErrCatVServer, sdkerror.ErrCatVirtualAddress)
	}

	return resp.ToAddressPair(), nil
}
