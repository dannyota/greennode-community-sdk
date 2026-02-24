package v2

import (
	"context"

	"danny.vn/greennode/client"
	sdkerror "danny.vn/greennode/sdkerror"
	"danny.vn/greennode/services/common"
)

const (
	defaultPageListNetworks = 1
	defaultSizeListNetworks = 10
)

type NetworkServiceV2 struct {
	Client *client.ServiceClient
}

func (s *NetworkServiceV2) ListNetworks(ctx context.Context, opts *ListNetworksRequest) (*ListNetworks, error) {
	url := listNetworksURL(s.Client, opts)
	resp := new(ListNetworksResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithOKCodes(200).
		WithJSONResponse(resp).
		WithJSONError(errResp)

	if _, sdkErr := s.Client.Get(ctx, url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp).
			WithParameters(common.StructToMap(opts)).
			WithKVparameters("projectId", s.Client.ProjectID)
	}

	return resp.ToEntityListNetworks(), nil
}

func (s *NetworkServiceV2) GetNetworkByID(ctx context.Context, opts *GetNetworkByIDRequest) (*Network, error) {
	url := getNetworkByIDURL(s.Client, opts)
	resp := new(GetNetworkByIDResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithOKCodes(200).
		WithJSONResponse(resp).
		WithJSONError(errResp)

	if _, sdkErr := s.Client.Get(ctx, url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.EcVServerNetworkNotFound).
			WithKVparameters(
				"networkId", opts.NetworkID,
				"projectId", s.Client.ProjectID)
	}

	return resp.ToEntityNetwork(), nil
}
