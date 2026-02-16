package v2

import (
	"context"

	"github.com/dannyota/greennode-community-sdk/v2/greennode/client"
	"github.com/dannyota/greennode-community-sdk/v2/greennode/entity"
	sdkerror "github.com/dannyota/greennode-community-sdk/v2/greennode/sdkerror"
)

type PortalServiceV2 struct {
	PortalClient *client.ServiceClient
}

func (s *PortalServiceV2) ListAllQuotaUsed(ctx context.Context) (*entity.ListQuotas, error) {
	url := listAllQuotaUsedURL(s.PortalClient)
	resp := new(ListAllQuotaUsedResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithOkCodes(200).
		WithJSONResponse(resp).
		WithJSONError(errResp)

	if _, sdkErr := s.PortalClient.Get(ctx, url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp).
			WithKVparameters("projectId", s.getProjectID())
	}

	return resp.ToEntityListQuotas(), nil
}

func (s *PortalServiceV2) GetQuotaByName(ctx context.Context, opts *GetQuotaByNameRequest) (*entity.Quota, error) {
	listQuotas, sdkErr := s.ListAllQuotaUsed(ctx)
	if sdkErr != nil {
		return nil, sdkErr
	}

	quota := listQuotas.FindQuotaByName(string(opts.Name))
	if quota == nil {
		return nil, sdkerror.NewQuotaNotFound().WithKVparameters("quotaName", opts.Name)
	}

	return quota, nil
}

func (s *PortalServiceV2) getProjectID() string {
	return s.PortalClient.GetProjectID()
}
