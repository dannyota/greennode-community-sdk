package v2

import (
	"context"

	"danny.vn/greennode/client"
	sdkerror "danny.vn/greennode/sdkerror"
)

type PortalServiceV2 struct {
	Client *client.ServiceClient
}

func (s *PortalServiceV2) ListAllQuotaUsed(ctx context.Context) (*ListQuotas, error) {
	url := listAllQuotaUsedURL(s.Client)
	resp := new(ListAllQuotaUsedResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithOKCodes(200).
		WithJSONResponse(resp).
		WithJSONError(errResp)

	if _, sdkErr := s.Client.Get(ctx, url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp).
			WithKVparameters("projectId", s.Client.ProjectID)
	}

	return resp.ToEntityListQuotas(), nil
}

func (s *PortalServiceV2) GetQuotaByName(ctx context.Context, opts *GetQuotaByNameRequest) (*Quota, error) {
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

