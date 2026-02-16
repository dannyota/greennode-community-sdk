package v2

import (
	"context"

	"github.com/dannyota/greennode-community-sdk/v2/greennode/client"
	"github.com/dannyota/greennode-community-sdk/v2/greennode/entity"
	sdkerror "github.com/dannyota/greennode-community-sdk/v2/greennode/sdkerror"
)

func (s *NetworkServiceV2) GetSecgroupByID(ctx context.Context, opts *GetSecgroupByIDRequest) (*entity.Secgroup, error) {
	url := getSecgroupByIDURL(s.VServerClient, opts)
	resp := new(GetSecgroupByIDResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithOkCodes(200).
		WithJSONResponse(resp).
		WithJSONError(errResp)

	if _, sdkErr := s.VServerClient.Get(ctx, url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.EcVServerSecgroupNotFound).
			WithKVparameters(
				"secgroupId", opts.SecgroupID,
				"projectId", s.getProjectID())
	}

	return resp.ToEntitySecgroup(), nil
}

func (s *NetworkServiceV2) CreateSecgroup(ctx context.Context, opts *CreateSecgroupRequest) (*entity.Secgroup, error) {
	url := createSecgroupURL(s.VServerClient)
	resp := new(CreateSecgroupResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithOkCodes(201).
		WithJSONBody(opts).
		WithJSONResponse(resp).
		WithJSONError(errResp)

	if _, sdkErr := s.VServerClient.Post(ctx, url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.EcVServerSecgroupNameAlreadyExists,
			sdkerror.EcVServerSecgroupRuleExceedQuota,
			sdkerror.EcVServerSecgroupExceedQuota).
			WithKVparameters(
				"secgroupName", opts.GetSecgroupName(),
				"projectId", s.getProjectID())
	}

	return resp.ToEntitySecgroup(), nil
}

func (s *NetworkServiceV2) ListSecgroup(ctx context.Context, opts *ListSecgroupRequest) (*entity.ListSecgroups, error) {
	url := listSecgroupURL(s.VServerClient, opts)
	resp := new(ListSecgroupResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithOkCodes(200).
		WithJSONResponse(resp).
		WithJSONError(errResp)

	if _, sdkErr := s.VServerClient.Get(ctx, url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp)
	}

	return resp.ToListEntitySecgroups(), nil
}

func (s *NetworkServiceV2) DeleteSecgroupByID(ctx context.Context, opts *DeleteSecgroupByIDRequest) error {
	url := deleteSecgroupByIDURL(s.VServerClient, opts)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithOkCodes(204).
		WithJSONError(errResp)

	if _, sdkErr := s.VServerClient.Delete(ctx, url, req); sdkErr != nil {
		return sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.EcVServerSecgroupInUse,
			sdkerror.EcVServerSecgroupNotFound).
			WithKVparameters(
				"secgroupId", opts.SecgroupID,
				"projectId", s.getProjectID())
	}

	return nil
}
