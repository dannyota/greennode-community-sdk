package intervpc

import (
	"context"

	"danny.vn/greennode/client"
	sdkerror "danny.vn/greennode/sdkerror"
	"danny.vn/greennode/services/common"
	lbv2 "danny.vn/greennode/services/loadbalancer/v2"
)

type LoadBalancerServiceInternal struct {
	Client *client.ServiceClient
}

func (s *LoadBalancerServiceInternal) CreateLoadBalancer(ctx context.Context, opts *CreateLoadBalancerRequest) (*lbv2.LoadBalancer, error) {
	url := createLoadBalancerURL(s.Client)
	resp := new(CreateLoadBalancerResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	opts.ProjectID = s.Client.ProjectID
	opts.normalizeForAPI()
	req := client.NewRequest().
		WithUserID(opts.PortalUser.ID).
		WithOKCodes(202).
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
