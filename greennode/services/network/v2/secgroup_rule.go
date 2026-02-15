package v2

import (
	"context"

	"github.com/dannyota/greennode-community-sdk/v2/greennode/client"
	"github.com/dannyota/greennode-community-sdk/v2/greennode/entity"
	sdkerror "github.com/dannyota/greennode-community-sdk/v2/greennode/sdkerror"
	"github.com/dannyota/greennode-community-sdk/v2/greennode/services/common"
)

func (s *NetworkServiceV2) CreateSecgroupRule(ctx context.Context, opts *CreateSecgroupRuleRequest) (*entity.SecgroupRule, error) {
	url := createSecgroupRuleURL(s.VServerClient, opts)
	resp := new(CreateSecgroupRuleResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithOkCodes(201).
		WithJSONBody(opts).
		WithJSONResponse(resp).
		WithJSONError(errResp)

	if _, sdkErr := s.VServerClient.Post(ctx, url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.EcVServerSecgroupNotFound,
			sdkerror.EcVServerSecgroupRuleExceedQuota,
			sdkerror.EcVServerSecgroupRuleAlreadyExists).
			WithParameters(common.StructToMap(opts)).
			WithKVparameters("projectId", s.getProjectID())
	}

	return resp.ToEntitySecgroupRule(), nil
}

func (s *NetworkServiceV2) DeleteSecgroupRuleByID(ctx context.Context, opts *DeleteSecgroupRuleByIDRequest) error {
	url := deleteSecgroupRuleByIDURL(s.VServerClient, opts)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithOkCodes(204).
		WithJSONError(errResp)

	if _, sdkErr := s.VServerClient.Delete(ctx, url, req); sdkErr != nil {
		return sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.EcVServerSecgroupRuleNotFound,
			sdkerror.EcVServerSecgroupNotFound).
			WithKVparameters(
				"secgroupId", opts.GetSecgroupID(),
				"secgroupRuleId", opts.GetSecgroupRuleID(),
				"projectId", s.getProjectID())
	}

	return nil
}

func (s *NetworkServiceV2) ListSecgroupRulesBySecgroupID(ctx context.Context, opts *ListSecgroupRulesBySecgroupIDRequest) (*entity.ListSecgroupRules, error) {
	url := listSecgroupRulesBySecgroupIDURL(s.VServerClient, opts)
	resp := new(ListSecgroupRulesBySecgroupIDResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithOkCodes(200).
		WithJSONResponse(resp).
		WithJSONError(errResp)

	if _, sdkErr := s.VServerClient.Get(ctx, url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.EcVServerSecgroupNotFound).
			WithKVparameters("projectId", s.getProjectID(), "secgroupId", opts.GetSecgroupID())
	}

	return resp.ToEntityListSecgroupRules(), nil
}
