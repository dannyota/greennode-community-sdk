package v2

import (
	"github.com/dannyota/greennode-community-sdk/v2/greennode/client"
	"github.com/dannyota/greennode-community-sdk/v2/greennode/entity"
	sdkerror "github.com/dannyota/greennode-community-sdk/v2/greennode/sdkerror"
)

func (s *NetworkServiceV2) GetSecgroupById(popts IGetSecgroupByIdRequest) (*entity.Secgroup, sdkerror.IError) {
	url := getSecgroupByIdUrl(s.VserverClient, popts)
	resp := new(GetSecgroupByIdResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithHeader("User-Agent", popts.ParseUserAgent()).
		WithOkCodes(200).
		WithJsonResponse(resp).
		WithJsonError(errResp)

	if _, sdkErr := s.VserverClient.Get(url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.WithErrorSecgroupNotFound(errResp)).
			WithKVparameters(
				"secgroupId", popts.GetSecgroupId(),
				"projectId", s.getProjectId())
	}

	return resp.ToEntitySecgroup(), nil
}

func (s *NetworkServiceV2) CreateSecgroup(popts ICreateSecgroupRequest) (*entity.Secgroup, sdkerror.IError) {
	url := createSecgroupUrl(s.VserverClient)
	resp := new(CreateSecgroupResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithHeader("User-Agent", popts.ParseUserAgent()).
		WithOkCodes(201).
		WithJsonBody(popts.ToRequestBody()).
		WithJsonResponse(resp).
		WithJsonError(errResp)

	if _, sdkErr := s.VserverClient.Post(url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.WithErrorSecgroupNameAlreadyExists(errResp),
			sdkerror.WithErrorSecgroupRuleExceedQuota(errResp),
			sdkerror.WithErrorSecgroupExceedQuota(errResp)).
			WithKVparameters(
				"secgroupName", popts.GetSecgroupName(),
				"projectId", s.getProjectId())
	}

	return resp.ToEntitySecgroup(), nil
}

func (s *NetworkServiceV2) ListSecgroup(popts IListSecgroupRequest) (*entity.ListSecgroups, sdkerror.IError) {
	url := listSecgroupUrl(s.VserverClient, popts)
	resp := new(ListSecgroupResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithHeader("User-Agent", popts.ParseUserAgent()).
		WithOkCodes(200).
		WithJsonResponse(resp).
		WithJsonError(errResp)

	if _, sdkErr := s.VserverClient.Get(url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp)
	}

	return resp.ToListEntitySecgroups(), nil
}

func (s *NetworkServiceV2) DeleteSecgroupById(popts IDeleteSecgroupByIdRequest) sdkerror.IError {
	url := deleteSecgroupByIdUrl(s.VserverClient, popts)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithHeader("User-Agent", popts.ParseUserAgent()).
		WithOkCodes(204).
		WithJsonError(errResp)

	if _, sdkErr := s.VserverClient.Delete(url, req); sdkErr != nil {
		return sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.WithErrorSecgroupInUse(errResp),
			sdkerror.WithErrorSecgroupNotFound(errResp)).
			WithKVparameters(
				"secgroupId", popts.GetSecgroupId(),
				"projectId", s.getProjectId())
	}

	return nil
}
