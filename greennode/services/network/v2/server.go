package v2

import (
	lsclient "github.com/dannyota/greennode-community-sdk/v2/greennode/client"
	lsentity "github.com/dannyota/greennode-community-sdk/v2/greennode/entity"
	lserr "github.com/dannyota/greennode-community-sdk/v2/greennode/sdk_error"
)

func (s *NetworkServiceV2) ListAllServersBySecgroupId(popts IListAllServersBySecgroupIdRequest) (*lsentity.ListServers, lserr.IError) {
	url := listAllServersBySecgroupIdUrl(s.VserverClient, popts)
	resp := new(ListAllServersBySecgroupIdResponse)
	errResp := lserr.NewErrorResponse(lserr.NormalErrorType)
	req := lsclient.NewRequest().
		WithHeader("User-Agent", popts.ParseUserAgent()).
		WithOkCodes(200).
		WithJsonResponse(resp).
		WithJsonError(errResp)

	if _, sdkErr := s.VserverClient.Get(url, req); sdkErr != nil {
		return nil, lserr.SdkErrorHandler(sdkErr, errResp,
			lserr.WithErrorSecgroupNotFound(errResp)).
			WithKVparameters(
				"secgroupId", popts.GetSecgroupId(),
				"projectId", s.getProjectId())
	}

	return resp.ToEntityListServers(), nil
}
