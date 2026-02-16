package v1

import (
	"context"

	"github.com/dannyota/greennode-community-sdk/v2/greennode/client"
	"github.com/dannyota/greennode-community-sdk/v2/greennode/entity"
	sdkerror "github.com/dannyota/greennode-community-sdk/v2/greennode/sdkerror"
)

type ServerServiceInternalV1 struct {
	VServerClient *client.ServiceClient
}

func (s *ServerServiceInternalV1) CreateSystemTags(ctx context.Context, opts *CreateSystemTagRequest) (*[]entity.SystemTag, error) {

	url := createSystemTagURL(s.VServerClient)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)

	resp := new([]entity.SystemTag)

	req := client.NewRequest().
		WithOkCodes(200).
		WithJSONBody(opts).
		WithJSONResponse(resp).
		WithJSONError(errResp)

	if _, sdkErr := s.VServerClient.Post(ctx, url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp)
	}

	return resp, nil
}
