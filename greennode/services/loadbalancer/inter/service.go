package inter

import (
	"context"

	"github.com/dannyota/greennode-community-sdk/v2/greennode/client"
	"github.com/dannyota/greennode-community-sdk/v2/greennode/entity"
	sdkerror "github.com/dannyota/greennode-community-sdk/v2/greennode/sdkerror"
	"github.com/dannyota/greennode-community-sdk/v2/greennode/services/common"
)

type LoadBalancerServiceInternal struct {
	Client *client.ServiceClient
}

func (s *LoadBalancerServiceInternal) CreateLoadBalancer(ctx context.Context, opts *CreateLoadBalancerRequest) (*entity.LoadBalancer, error) {
	url := createLoadBalancerURL(s.Client)
	resp := new(CreateLoadBalancerResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	opts.ProjectID = s.Client.ProjectID()
	opts.normalizeForAPI()
	req := client.NewRequest().
		WithMapHeaders(opts.GetMapHeaders()).
		WithOkCodes(202).
		WithJSONBody(opts).
		WithJSONResponse(resp).
		WithJSONError(errResp)

	if _, sdkErr := s.Client.Post(ctx, url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.EcVLBLoadBalancerExceedQuota).
			WithParameters(common.StructToMap(opts)).
			AppendCategories(sdkerror.ErrCatProductVlb)
	}

	return resp.ToEntityLoadBalancer(), nil
}
