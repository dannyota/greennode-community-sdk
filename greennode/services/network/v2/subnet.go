package v2

import (
	"github.com/dannyota/greennode-community-sdk/v2/greennode/client"
	"github.com/dannyota/greennode-community-sdk/v2/greennode/entity"
	sdkerror "github.com/dannyota/greennode-community-sdk/v2/greennode/sdkerror"
)

func (s *NetworkServiceV2) GetSubnetByID(opts *GetSubnetByIDRequest) (*entity.Subnet, error) {
	url := getSubnetByIDURL(s.VServerClient, opts)
	resp := new(GetSubnetByIDResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithOkCodes(200).
		WithJSONResponse(resp).
		WithJSONError(errResp)

	if _, sdkErr := s.VServerClient.Get(url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.EcVServerSubnetNotBelongNetwork,
			sdkerror.EcVServerSubnetNotFound).
			WithKVparameters(
				"subnetId", opts.GetSubnetID(),
				"projectId", s.getProjectID())
	}

	return resp.ToEntitySubnet(), nil
}

func (s *NetworkServiceV2) UpdateSubnetByID(opts *UpdateSubnetByIDRequest) (*entity.Subnet, error) {
	url := updateSubnetByIDURL(s.VServerClient, opts)
	resp := new(UpdateSubnetByIDResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithOkCodes(200, 201, 202, 203, 204).
		WithJSONBody(opts.UpdateSubnetBody).
		WithJSONResponse(resp).
		WithJSONError(errResp)

	_, sdkErr := s.VServerClient.Patch(url, req)
	if sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp)
	}

	return resp.ToEntitySubnet(), nil
}
