package v2

import (
	"context"

	"github.com/dannyota/greennode-community-sdk/v2/greennode/client"
	sdkerror "github.com/dannyota/greennode-community-sdk/v2/greennode/sdkerror"
	"github.com/dannyota/greennode-community-sdk/v2/greennode/services/common"
)

func (s *NetworkServiceV2) CreateSecgroupRule(ctx context.Context, opts *CreateSecgroupRuleRequest) (*SecgroupRule, error) {
	url := createSecgroupRuleURL(s.Client, opts)
	resp := new(CreateSecgroupRuleResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithOkCodes(201).
		WithJSONBody(opts).
		WithJSONResponse(resp).
		WithJSONError(errResp)

	if _, sdkErr := s.Client.Post(ctx, url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.EcVServerSecgroupNotFound,
			sdkerror.EcVServerSecgroupRuleExceedQuota,
			sdkerror.EcVServerSecgroupRuleAlreadyExists).
			WithParameters(common.StructToMap(opts)).
			WithKVparameters("projectId", s.Client.ProjectID)
	}

	return resp.ToEntitySecgroupRule(), nil
}

func (s *NetworkServiceV2) DeleteSecgroupRuleByID(ctx context.Context, opts *DeleteSecgroupRuleByIDRequest) error {
	url := deleteSecgroupRuleByIDURL(s.Client, opts)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithOkCodes(204).
		WithJSONError(errResp)

	if _, sdkErr := s.Client.Delete(ctx, url, req); sdkErr != nil {
		return sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.EcVServerSecgroupRuleNotFound,
			sdkerror.EcVServerSecgroupNotFound).
			WithKVparameters(
				"secgroupId", opts.SecgroupID,
				"secgroupRuleId", opts.SecgroupRuleID,
				"projectId", s.Client.ProjectID)
	}

	return nil
}

func (s *NetworkServiceV2) ListSecgroupRulesBySecgroupID(ctx context.Context, opts *ListSecgroupRulesBySecgroupIDRequest) (*ListSecgroupRules, error) {
	url := listSecgroupRulesBySecgroupIDURL(s.Client, opts)
	resp := new(ListSecgroupRulesBySecgroupIDResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithOkCodes(200).
		WithJSONResponse(resp).
		WithJSONError(errResp)

	if _, sdkErr := s.Client.Get(ctx, url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.EcVServerSecgroupNotFound).
			WithKVparameters("projectId", s.Client.ProjectID, "secgroupId", opts.SecgroupID)
	}

	return resp.ToEntityListSecgroupRules(), nil
}
