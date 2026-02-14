package v2

import (
	"github.com/dannyota/greennode-community-sdk/v2/greennode/client"
	"github.com/dannyota/greennode-community-sdk/v2/greennode/entity"
	sdkerror "github.com/dannyota/greennode-community-sdk/v2/greennode/sdkerror"
)

func (s *NetworkServiceV2) GetSubnetById(opts IGetSubnetByIdRequest) (*entity.Subnet, sdkerror.Error) {
	url := getSubnetByIdUrl(s.VserverClient, opts)
	resp := new(GetSubnetByIdResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithHeader("User-Agent", opts.ParseUserAgent()).
		WithOkCodes(200).
		WithJsonResponse(resp).
		WithJsonError(errResp)

	if _, sdkErr := s.VserverClient.Get(url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.WithErrorSubnetNotBelongNetwork(sdkErr),
			sdkerror.WithErrorSubnetNotFound(errResp)).
			WithKVparameters(
				"subnetId", opts.GetSubnetId(),
				"projectId", s.getProjectId())
	}

	return resp.ToEntitySubnet(), nil
}

func (s *NetworkServiceV2) UpdateSubnetById(opts IUpdateSubnetByIdRequest) (*entity.Subnet, sdkerror.Error) {
	url := updateSubnetByIdUrl(s.VserverClient, opts)
	resp := new(UpdateSubnetByIdResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithHeader("User-Agent", opts.ParseUserAgent()).
		WithOkCodes(200, 201, 202, 203, 204).
		WithJsonBody(opts.ToRequestBody()).
		WithJsonResponse(resp).
		WithJsonError(errResp)

	_, sdkErr := s.VserverClient.Patch(url, req)
	if sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp)
	}

	return resp.ToEntitySubnet(), nil
}
