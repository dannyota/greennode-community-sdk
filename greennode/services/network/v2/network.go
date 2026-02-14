package v2

import (
	"github.com/dannyota/greennode-community-sdk/v2/greennode/client"
	"github.com/dannyota/greennode-community-sdk/v2/greennode/entity"
	sdkerror "github.com/dannyota/greennode-community-sdk/v2/greennode/sdkerror"
)

func (s *NetworkServiceV2) GetNetworkByID(opts IGetNetworkByIDRequest) (*entity.Network, sdkerror.Error) {
	url := getNetworkByIDURL(s.VserverClient, opts)
	resp := new(GetNetworkByIDResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithHeader("User-Agent", opts.ParseUserAgent()).
		WithOkCodes(200).
		WithJSONResponse(resp).
		WithJSONError(errResp)

	if _, sdkErr := s.VserverClient.Get(url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.WithErrorNetworkNotFound(errResp)).
			WithKVparameters(
				"networkId", opts.GetNetworkID(),
				"projectId", s.getProjectID())
	}

	return resp.ToEntityNetwork(), nil
}
