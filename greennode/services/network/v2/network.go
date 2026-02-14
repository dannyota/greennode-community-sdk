package v2

import (
	"github.com/dannyota/greennode-community-sdk/v2/greennode/client"
	"github.com/dannyota/greennode-community-sdk/v2/greennode/entity"
	sdkerror "github.com/dannyota/greennode-community-sdk/v2/greennode/sdkerror"
)

func (s *NetworkServiceV2) GetNetworkById(opts IGetNetworkByIdRequest) (*entity.Network, sdkerror.Error) {
	url := getNetworkByIdUrl(s.VserverClient, opts)
	resp := new(GetNetworkByIdResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithHeader("User-Agent", opts.ParseUserAgent()).
		WithOkCodes(200).
		WithJsonResponse(resp).
		WithJsonError(errResp)

	if _, sdkErr := s.VserverClient.Get(url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.WithErrorNetworkNotFound(errResp)).
			WithKVparameters(
				"networkId", opts.GetNetworkId(),
				"projectId", s.getProjectId())
	}

	return resp.ToEntityNetwork(), nil
}
