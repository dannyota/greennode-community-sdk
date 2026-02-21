package v2

import (
	"context"

	"github.com/dannyota/greennode-community-sdk/greennode/client"
	sdkerror "github.com/dannyota/greennode-community-sdk/greennode/sdkerror"
)

func (s *NetworkServiceV2) GetSecgroupByID(ctx context.Context, opts *GetSecgroupByIDRequest) (*Secgroup, error) {
	url := getSecgroupByIDURL(s.Client, opts)
	resp := new(GetSecgroupByIDResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithOKCodes(200).
		WithJSONResponse(resp).
		WithJSONError(errResp)

	if _, sdkErr := s.Client.Get(ctx, url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.EcVServerSecgroupNotFound).
			WithKVparameters(
				"secgroupId", opts.SecgroupID,
				"projectId", s.Client.ProjectID)
	}

	return resp.ToEntitySecgroup(), nil
}

func (s *NetworkServiceV2) CreateSecgroup(ctx context.Context, opts *CreateSecgroupRequest) (*Secgroup, error) {
	url := createSecgroupURL(s.Client)
	resp := new(CreateSecgroupResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithOKCodes(201).
		WithJSONBody(opts).
		WithJSONResponse(resp).
		WithJSONError(errResp)

	if _, sdkErr := s.Client.Post(ctx, url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.EcVServerSecgroupNameAlreadyExists,
			sdkerror.EcVServerSecgroupRuleExceedQuota,
			sdkerror.EcVServerSecgroupExceedQuota).
			WithKVparameters(
				"secgroupName", opts.Name,
				"projectId", s.Client.ProjectID)
	}

	return resp.ToEntitySecgroup(), nil
}

func (s *NetworkServiceV2) ListSecgroup(ctx context.Context, opts *ListSecgroupRequest) (*ListSecgroups, error) {
	url := listSecgroupURL(s.Client, opts)
	resp := new(ListSecgroupResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithOKCodes(200).
		WithJSONResponse(resp).
		WithJSONError(errResp)

	if _, sdkErr := s.Client.Get(ctx, url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp)
	}

	return resp.ToListEntitySecgroups(), nil
}

func (s *NetworkServiceV2) DeleteSecgroupByID(ctx context.Context, opts *DeleteSecgroupByIDRequest) error {
	url := deleteSecgroupByIDURL(s.Client, opts)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithOKCodes(204).
		WithJSONError(errResp)

	if _, sdkErr := s.Client.Delete(ctx, url, req); sdkErr != nil {
		return sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.EcVServerSecgroupInUse,
			sdkerror.EcVServerSecgroupNotFound).
			WithKVparameters(
				"secgroupId", opts.SecgroupID,
				"projectId", s.Client.ProjectID)
	}

	return nil
}
