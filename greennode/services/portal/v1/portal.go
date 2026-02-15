package v1

import (
	"github.com/dannyota/greennode-community-sdk/v2/greennode/client"
	"github.com/dannyota/greennode-community-sdk/v2/greennode/entity"
	sdkerror "github.com/dannyota/greennode-community-sdk/v2/greennode/sdkerror"
)

func (s *PortalServiceV1) GetPortalInfo(opts *GetPortalInfoRequest) (*entity.Portal, sdkerror.Error) {
	url := getPortalInfoURL(s.PortalClient, opts)
	resp := new(GetPortalInfoResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithOkCodes(200).
		WithJSONResponse(resp).
		WithJSONError(errResp)

	if _, sdkErr := s.PortalClient.Get(url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp).
			WithKVparameters("backendProjectId", opts.GetBackEndProjectID())
	}

	return resp.ToEntityPortal(), nil
}

func (s *PortalServiceV1) ListProjects(opts *ListProjectsRequest) (*entity.ListPortals, sdkerror.Error) {
	url := listProjectsURL(s.PortalClient)
	resp := new(ListProjectsResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithHeader("User-Agent", opts.ParseUserAgent()).
		WithOkCodes(200).
		WithJSONResponse(resp).
		WithJSONError(errResp)

	if _, sdkErr := s.PortalClient.Get(url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp)
	}

	return resp.ToEntityListPortals(), nil
}
