package v1

import (
	"context"

	"github.com/dannyota/greennode-community-sdk/greennode/client"
	sdkerror "github.com/dannyota/greennode-community-sdk/greennode/sdkerror"
)

type ServerServiceInternalV1 struct {
	Client *client.ServiceClient
}

func (s *ServerServiceInternalV1) CreateSystemTags(ctx context.Context, opts *CreateSystemTagRequest) (*[]SystemTag, error) {

	url := createSystemTagURL(s.Client)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)

	resp := new([]SystemTag)

	req := client.NewRequest().
		WithOKCodes(200).
		WithJSONBody(opts).
		WithJSONResponse(resp).
		WithJSONError(errResp)

	if _, sdkErr := s.Client.Post(ctx, url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp)
	}

	return resp, nil
}
