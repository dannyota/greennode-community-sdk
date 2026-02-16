package v2

import (
	"context"

	"github.com/dannyota/greennode-community-sdk/v2/greennode/client"
	"github.com/dannyota/greennode-community-sdk/v2/greennode/entity"
	sdkerror "github.com/dannyota/greennode-community-sdk/v2/greennode/sdkerror"
)

type NetworkServiceV2 struct {
	VServerClient *client.ServiceClient
}

func (s *NetworkServiceV2) getProjectID() string {
	return s.VServerClient.GetProjectID()
}

func (s *NetworkServiceV2) GetNetworkByID(ctx context.Context, opts *GetNetworkByIDRequest) (*entity.Network, error) {
	url := getNetworkByIDURL(s.VServerClient, opts)
	resp := new(GetNetworkByIDResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithOkCodes(200).
		WithJSONResponse(resp).
		WithJSONError(errResp)

	if _, sdkErr := s.VServerClient.Get(ctx, url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.EcVServerNetworkNotFound).
			WithKVparameters(
				"networkId", opts.NetworkID,
				"projectId", s.getProjectID())
	}

	return resp.ToEntityNetwork(), nil
}
