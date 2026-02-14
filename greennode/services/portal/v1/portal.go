package v1

import (
	"github.com/dannyota/greennode-community-sdk/v2/greennode/client"
	"github.com/dannyota/greennode-community-sdk/v2/greennode/entity"
	sdkerror "github.com/dannyota/greennode-community-sdk/v2/greennode/sdkerror"
)

func (s *PortalServiceV1) GetPortalInfo(popts IGetPortalInfoRequest) (*entity.Portal, sdkerror.Error) {
	url := getPortalInfoUrl(s.PortalClient, popts)
	resp := new(GetPortalInfoResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithOkCodes(200).
		WithJsonResponse(resp).
		WithJsonError(errResp)

	if _, sdkErr := s.PortalClient.Get(url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp).
			WithKVparameters("backendProjectId", popts.GetBackEndProjectId())
	}

	return resp.ToEntityPortal(), nil
}

func (s *PortalServiceV1) ListProjects(popts IListProjectsRequest) (*entity.ListPortals, sdkerror.Error) {
	url := listProjectsUrl(s.PortalClient)
	resp := new(ListProjectsResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithHeader("User-Agent", popts.ParseUserAgent()).
		WithOkCodes(200).
		WithJsonResponse(resp).
		WithJsonError(errResp)

	if _, sdkErr := s.PortalClient.Get(url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp)
	}

	return resp.ToEntityListPortals(), nil
}
