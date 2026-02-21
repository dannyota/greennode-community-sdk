package v2

import (
	"context"

	"github.com/dannyota/greennode-community-sdk/greennode/client"
	sdkerror "github.com/dannyota/greennode-community-sdk/greennode/sdkerror"
)

type NetworkServiceV2 struct {
	Client *client.ServiceClient
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
