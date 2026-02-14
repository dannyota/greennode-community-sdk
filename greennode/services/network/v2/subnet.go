package v2

import (
	"github.com/dannyota/greennode-community-sdk/v2/greennode/client"
	"github.com/dannyota/greennode-community-sdk/v2/greennode/entity"
	sdkerror "github.com/dannyota/greennode-community-sdk/v2/greennode/sdkerror"
)

func (s *NetworkServiceV2) GetSubnetById(popts IGetSubnetByIdRequest) (*entity.Subnet, sdkerror.IError) {
	url := getSubnetByIdUrl(s.VserverClient, popts)
	resp := new(GetSubnetByIdResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithHeader("User-Agent", popts.ParseUserAgent()).
		WithOkCodes(200).
		WithJsonResponse(resp).
		WithJsonError(errResp)

	if _, sdkErr := s.VserverClient.Get(url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.WithErrorSubnetNotBelongNetwork(sdkErr),
			sdkerror.WithErrorSubnetNotFound(errResp)).
			WithKVparameters(
				"subnetId", popts.GetSubnetId(),
				"projectId", s.getProjectId())
	}

	return resp.ToEntitySubnet(), nil
}

func (s *NetworkServiceV2) UpdateSubnetById(popts IUpdateSubnetByIdRequest) (*entity.Subnet, sdkerror.IError) {
	url := updateSubnetByIdUrl(s.VserverClient, popts)
	resp := new(UpdateSubnetByIdResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithHeader("User-Agent", popts.ParseUserAgent()).
		WithOkCodes(200, 201, 202, 203, 204).
		WithJsonBody(popts.ToRequestBody()).
		WithJsonResponse(resp).
		WithJsonError(errResp)

	_, sdkErr := s.VserverClient.Patch(url, req)
	if sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp)
	}

	return resp.ToEntitySubnet(), nil
}
