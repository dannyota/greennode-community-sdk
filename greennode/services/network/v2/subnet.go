package v2

import (
	"github.com/dannyota/greennode-community-sdk/v2/greennode/client"
	"github.com/dannyota/greennode-community-sdk/v2/greennode/entity"
	sdkerror "github.com/dannyota/greennode-community-sdk/v2/greennode/sdkerror"
)

func (s *NetworkServiceV2) GetSubnetByID(opts IGetSubnetByIDRequest) (*entity.Subnet, sdkerror.Error) {
	url := getSubnetByIDURL(s.VserverClient, opts)
	resp := new(GetSubnetByIDResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithHeader("User-Agent", opts.ParseUserAgent()).
		WithOkCodes(200).
		WithJSONResponse(resp).
		WithJSONError(errResp)

	if _, sdkErr := s.VserverClient.Get(url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.WithErrorSubnetNotBelongNetwork(sdkErr),
			sdkerror.WithErrorSubnetNotFound(errResp)).
			WithKVparameters(
				"subnetId", opts.GetSubnetID(),
				"projectId", s.getProjectID())
	}

	return resp.ToEntitySubnet(), nil
}

func (s *NetworkServiceV2) UpdateSubnetByID(opts IUpdateSubnetByIDRequest) (*entity.Subnet, sdkerror.Error) {
	url := updateSubnetByIDURL(s.VserverClient, opts)
	resp := new(UpdateSubnetByIDResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithHeader("User-Agent", opts.ParseUserAgent()).
		WithOkCodes(200, 201, 202, 203, 204).
		WithJSONBody(opts.ToRequestBody()).
		WithJSONResponse(resp).
		WithJSONError(errResp)

	_, sdkErr := s.VserverClient.Patch(url, req)
	if sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp)
	}

	return resp.ToEntitySubnet(), nil
}
