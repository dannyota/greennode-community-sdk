package v2

import (
	"github.com/dannyota/greennode-community-sdk/v2/greennode/client"
	"github.com/dannyota/greennode-community-sdk/v2/greennode/entity"
	sdkerror "github.com/dannyota/greennode-community-sdk/v2/greennode/sdkerror"
)

func (s *PortalServiceV2) ListAllQuotaUsed() (*entity.ListQuotas, sdkerror.Error) {
	url := listAllQuotaUsedUrl(s.PortalClient)
	resp := new(ListAllQuotaUsedResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithOkCodes(200).
		WithJsonResponse(resp).
		WithJsonError(errResp)

	if _, sdkErr := s.PortalClient.Get(url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp).
			WithKVparameters("projectId", s.getProjectId())
	}

	return resp.ToEntityListQuotas(), nil
}

func (s *PortalServiceV2) GetQuotaByName(popts IGetQuotaByNameRequest) (*entity.Quota, sdkerror.Error) {
	listQuotas, sdkErr := s.ListAllQuotaUsed()
	if sdkErr != nil {
		return nil, sdkErr
	}

	quota := listQuotas.FindQuotaByName(string(popts.GetName()))
	if quota == nil {
		return nil, sdkerror.ErrorHandler(nil, sdkerror.WithErrorQuotaNotFound(nil)).WithKVparameters("quotaName", popts.GetName())
	}

	return quota, nil
}

func (s *PortalServiceV2) getProjectId() string {
	return s.PortalClient.GetProjectId()
}
