package v1

import (
	"context"

	"danny.vn/greennode/greennode/client"
	sdkerror "danny.vn/greennode/greennode/sdkerror"
	"danny.vn/greennode/greennode/services/common"
)

type GLBServiceV1 struct {
	Client *client.ServiceClient
}

const (
	defaultOffsetListGlobalLoadBalancer = 0
	defaultLimitListGlobalLoadBalancer  = 10
)

func (s *GLBServiceV1) ListGlobalPools(ctx context.Context, opts *ListGlobalPoolsRequest) (*ListGlobalPools, error) {
	url := listGlobalPoolsURL(s.Client, opts)
	resp := new(ListGlobalPoolsResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.GlobalLoadBalancerErrorType)
	req := client.NewRequest().
		WithOKCodes(200).
		WithJSONResponse(resp).
		WithJSONError(errResp)

	if _, sdkErr := s.Client.Get(ctx, url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.EcGlobalLoadBalancerNotFound).
			WithKVparameters("loadBalancerId", opts.LoadBalancerID).
			AppendCategories(sdkerror.ErrCatProductVlb)
	}

	return resp.ToEntityListGlobalPools(), nil
}

func (s *GLBServiceV1) CreateGlobalPool(ctx context.Context, opts *CreateGlobalPoolRequest) (*GlobalPool, error) {
	url := createGlobalPoolURL(s.Client, opts)
	resp := new(GlobalPool)
	errResp := sdkerror.NewErrorResponse(sdkerror.GlobalLoadBalancerErrorType)
	req := client.NewRequest().
		WithOKCodes(202).
		WithJSONBody(opts).
		WithJSONResponse(resp).
		WithJSONError(errResp)

	if _, sdkErr := s.Client.Post(ctx, url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.EcGlobalLoadBalancerNotFound).
			WithParameters(common.StructToMap(opts)).
			AppendCategories(sdkerror.ErrCatProductVlb)
	}

	return resp, nil
}

func (s *GLBServiceV1) UpdateGlobalPool(ctx context.Context, opts *UpdateGlobalPoolRequest) (*GlobalPool, error) {
	url := updateGlobalPoolURL(s.Client, opts)
	resp := new(UpdateGlobalPoolResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.GlobalLoadBalancerErrorType)
	req := client.NewRequest().
		WithOKCodes(202).
		WithJSONBody(opts).
		WithJSONResponse(resp).
		WithJSONError(errResp)

	if _, sdkErr := s.Client.Put(ctx, url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.EcGlobalLoadBalancerNotFound).
			WithParameters(common.StructToMap(opts)).
			AppendCategories(sdkerror.ErrCatProductVlb)
	}

	return resp.ToEntityPool(), nil
}

func (s *GLBServiceV1) DeleteGlobalPool(ctx context.Context, opts *DeleteGlobalPoolRequest) error {
	url := deleteGlobalPoolURL(s.Client, opts)
	errResp := sdkerror.NewErrorResponse(sdkerror.GlobalLoadBalancerErrorType)
	req := client.NewRequest().
		WithOKCodes(202).
		WithJSONError(errResp)

	if _, sdkErr := s.Client.Delete(ctx, url, req); sdkErr != nil {
		return sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.EcGlobalLoadBalancerNotFound).
			AppendCategories(sdkerror.ErrCatProductVlb)
	}

	return nil
}

func (s *GLBServiceV1) ListGlobalPoolMembers(ctx context.Context, opts *ListGlobalPoolMembersRequest) (*ListGlobalPoolMembers, error) {
	url := listGlobalPoolMembersURL(s.Client, opts)
	resp := new(ListGlobalPoolMembersResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.GlobalLoadBalancerErrorType)
	req := client.NewRequest().
		WithOKCodes(200).
		WithJSONResponse(resp).
		WithJSONError(errResp)

	if _, sdkErr := s.Client.Get(ctx, url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.EcGlobalLoadBalancerNotFound).
			AppendCategories(sdkerror.ErrCatProductVlb)
	}

	return resp.ToEntityListGlobalPoolMembers(), nil
}

func (s *GLBServiceV1) GetGlobalPoolMember(ctx context.Context, opts *GetGlobalPoolMemberRequest) (*GlobalPoolMember, error) {
	url := getGlobalPoolMemberURL(s.Client, opts)
	resp := new(GetGlobalPoolMemberResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.GlobalLoadBalancerErrorType)
	req := client.NewRequest().
		WithOKCodes(200).
		WithJSONResponse(resp).
		WithJSONError(errResp)

	if _, sdkErr := s.Client.Get(ctx, url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.EcGlobalLoadBalancerNotFound).
			WithKVparameters("loadBalancerId", opts.LoadBalancerID).
			WithKVparameters("poolId", opts.PoolID).
			WithKVparameters("poolMemberId", opts.PoolMemberID).
			AppendCategories(sdkerror.ErrCatProductVlb)
	}

	return resp.ToEntityGlobalPoolMember(), nil
}

func (s *GLBServiceV1) DeleteGlobalPoolMember(ctx context.Context, opts *DeleteGlobalPoolMemberRequest) error {
	url := deleteGlobalPoolMemberURL(s.Client, opts)
	errResp := sdkerror.NewErrorResponse(sdkerror.GlobalLoadBalancerErrorType)
	req := client.NewRequest().
		WithOKCodes(202).
		WithJSONError(errResp)

	if _, sdkErr := s.Client.Delete(ctx, url, req); sdkErr != nil {
		return sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.EcGlobalLoadBalancerNotFound).
			WithKVparameters("loadBalancerId", opts.LoadBalancerID).
			WithKVparameters("poolId", opts.PoolID).
			WithKVparameters("poolMemberId", opts.PoolMemberID).
			AppendCategories(sdkerror.ErrCatProductVlb)
	}

	return nil
}

func (s *GLBServiceV1) UpdateGlobalPoolMember(ctx context.Context, opts *UpdateGlobalPoolMemberRequest) (*GlobalPoolMember, error) {
	url := updateGlobalPoolMemberURL(s.Client, opts)
	resp := new(UpdateGlobalPoolMemberResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.GlobalLoadBalancerErrorType)
	req := client.NewRequest().
		WithOKCodes(202).
		WithJSONBody(opts).
		WithJSONResponse(resp).
		WithJSONError(errResp)

	if _, sdkErr := s.Client.Put(ctx, url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.EcGlobalLoadBalancerNotFound).
			WithParameters(common.StructToMap(opts)).
			AppendCategories(sdkerror.ErrCatProductVlb)
	}

	return resp.ToEntityGlobalPoolMember(), nil
}

func (s *GLBServiceV1) PatchGlobalPoolMembers(ctx context.Context, opts *PatchGlobalPoolMembersRequest) error {
	url := patchGlobalPoolMembersURL(s.Client, opts)
	errResp := sdkerror.NewErrorResponse(sdkerror.GlobalLoadBalancerErrorType)
	req := client.NewRequest().
		WithOKCodes(202).
		WithJSONBody(opts).
		WithJSONError(errResp)

	if _, sdkErr := s.Client.Patch(ctx, url, req); sdkErr != nil {
		return sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.EcGlobalLoadBalancerNotFound).
			WithParameters(common.StructToMap(opts)).
			AppendCategories(sdkerror.ErrCatProductVlb)
	}

	return nil
}

func (s *GLBServiceV1) ListGlobalListeners(ctx context.Context, opts *ListGlobalListenersRequest) (*ListGlobalListeners, error) {
	url := listGlobalListenersURL(s.Client, opts)
	resp := new(ListGlobalListenersResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.GlobalLoadBalancerErrorType)
	req := client.NewRequest().
		WithOKCodes(200).
		WithJSONResponse(resp).
		WithJSONError(errResp)

	if _, sdkErr := s.Client.Get(ctx, url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.EcGlobalLoadBalancerNotFound).
			AppendCategories(sdkerror.ErrCatProductVlb)
	}

	return resp.ToEntityListGlobalListeners(), nil
}

func (s *GLBServiceV1) CreateGlobalListener(ctx context.Context, opts *CreateGlobalListenerRequest) (*GlobalListener, error) {
	url := createGlobalListenerURL(s.Client, opts)
	resp := new(GlobalListener)
	errResp := sdkerror.NewErrorResponse(sdkerror.GlobalLoadBalancerErrorType)
	req := client.NewRequest().
		WithOKCodes(202).
		WithJSONBody(opts).
		WithJSONResponse(resp).
		WithJSONError(errResp)

	if _, sdkErr := s.Client.Post(ctx, url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.EcGlobalLoadBalancerNotFound).
			WithParameters(common.StructToMap(opts)).
			AppendCategories(sdkerror.ErrCatProductVlb)
	}

	return resp, nil
}

func (s *GLBServiceV1) UpdateGlobalListener(ctx context.Context, opts *UpdateGlobalListenerRequest) (*GlobalListener, error) {
	url := updateGlobalListenerURL(s.Client, opts)
	resp := new(GlobalListener)
	errResp := sdkerror.NewErrorResponse(sdkerror.GlobalLoadBalancerErrorType)
	req := client.NewRequest().
		WithOKCodes(202).
		WithJSONBody(opts).
		WithJSONResponse(resp).
		WithJSONError(errResp)

	if _, sdkErr := s.Client.Put(ctx, url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.EcGlobalLoadBalancerNotFound).
			WithParameters(common.StructToMap(opts)).
			AppendCategories(sdkerror.ErrCatProductVlb)
	}

	return resp, nil
}

func (s *GLBServiceV1) GetGlobalListener(ctx context.Context, opts *GetGlobalListenerRequest) (*GlobalListener, error) {
	url := getGlobalListenerURL(s.Client, opts)
	resp := new(GlobalListener)
	errResp := sdkerror.NewErrorResponse(sdkerror.GlobalLoadBalancerErrorType)
	req := client.NewRequest().
		WithOKCodes(200).
		WithJSONResponse(resp).
		WithJSONError(errResp)

	if _, sdkErr := s.Client.Get(ctx, url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.EcGlobalLoadBalancerNotFound).
			WithKVparameters("loadBalancerId", opts.LoadBalancerID).
			WithKVparameters("listenerId", opts.ListenerID).
			AppendCategories(sdkerror.ErrCatProductVlb)
	}

	return resp, nil
}

func (s *GLBServiceV1) DeleteGlobalListener(ctx context.Context, opts *DeleteGlobalListenerRequest) error {
	url := deleteGlobalListenerURL(s.Client, opts)
	errResp := sdkerror.NewErrorResponse(sdkerror.GlobalLoadBalancerErrorType)
	req := client.NewRequest().
		WithOKCodes(202).
		WithJSONError(errResp)

	if _, sdkErr := s.Client.Delete(ctx, url, req); sdkErr != nil {
		return sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.EcGlobalLoadBalancerNotFound).
			AppendCategories(sdkerror.ErrCatProductVlb)
	}

	return nil
}

func (s *GLBServiceV1) ListGlobalLoadBalancers(ctx context.Context, opts *ListGlobalLoadBalancersRequest) (*ListGlobalLoadBalancers, error) {
	url := listGlobalLoadBalancersURL(s.Client, opts)
	resp := new(ListGlobalLoadBalancers)
	errResp := sdkerror.NewErrorResponse(sdkerror.GlobalLoadBalancerErrorType)
	req := client.NewRequest().
		WithOKCodes(200).
		WithJSONResponse(resp).
		WithJSONError(errResp)

	if _, sdkErr := s.Client.Get(ctx, url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp).
			AppendCategories(sdkerror.ErrCatProductVlb)
	}

	return resp, nil
}

func (s *GLBServiceV1) CreateGlobalLoadBalancer(
	ctx context.Context,
	opts *CreateGlobalLoadBalancerRequest,
) (*GlobalLoadBalancer, error) {
	url := createGlobalLoadBalancerURL(s.Client, opts)
	resp := new(CreateGlobalLoadBalancerResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.GlobalLoadBalancerErrorType)
	req := client.NewRequest().
		WithOKCodes(202).
		WithJSONBody(opts).
		WithJSONResponse(resp).
		WithJSONError(errResp)

	if _, sdkErr := s.Client.Post(ctx, url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp).
			WithParameters(common.StructToMap(opts)).
			AppendCategories(sdkerror.ErrCatProductVlb)
	}

	return resp.ToEntityGlobalLoadBalancer(), nil
}

func (s *GLBServiceV1) DeleteGlobalLoadBalancer(ctx context.Context, opts *DeleteGlobalLoadBalancerRequest) error {
	url := deleteGlobalLoadBalancerURL(s.Client, opts)
	errResp := sdkerror.NewErrorResponse(sdkerror.GlobalLoadBalancerErrorType)
	req := client.NewRequest().
		WithOKCodes(202).
		WithJSONError(errResp)

	if _, sdkErr := s.Client.Delete(ctx, url, req); sdkErr != nil {
		return sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.EcGlobalLoadBalancerNotFound).
			AppendCategories(sdkerror.ErrCatProductVlb)
	}

	return nil
}

func (s *GLBServiceV1) GetGlobalLoadBalancerByID(
	ctx context.Context,
	opts *GetGlobalLoadBalancerByIDRequest,
) (*GlobalLoadBalancer, error) {
	url := getGlobalLoadBalancerByIDURL(s.Client, opts)
	resp := new(GlobalLoadBalancer)
	errResp := sdkerror.NewErrorResponse(sdkerror.GlobalLoadBalancerErrorType)
	req := client.NewRequest().
		WithOKCodes(200).
		WithJSONResponse(resp).
		WithJSONError(errResp)

	if _, sdkErr := s.Client.Get(ctx, url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.EcGlobalLoadBalancerNotFound).
			AppendCategories(sdkerror.ErrCatProductVlb)
	}

	return resp, nil
}

func (s *GLBServiceV1) ListGlobalPackages(ctx context.Context, opts *ListGlobalPackagesRequest) (*ListGlobalPackages, error) {
	url := listGlobalPackagesURL(s.Client, opts)
	resp := new(ListGlobalPackagesResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.GlobalLoadBalancerErrorType)
	req := client.NewRequest().
		WithOKCodes(200).
		WithJSONResponse(resp).
		WithJSONError(errResp)

	if _, sdkErr := s.Client.Get(ctx, url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp).
			AppendCategories(sdkerror.ErrCatProductVlb)
	}

	return resp.ToEntityListGlobalPackages(), nil
}

func (s *GLBServiceV1) ListGlobalRegions(ctx context.Context, opts *ListGlobalRegionsRequest) (*ListGlobalRegions, error) {
	url := listGlobalRegionsURL(s.Client, opts)
	resp := new(ListGlobalRegionsResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.GlobalLoadBalancerErrorType)
	req := client.NewRequest().
		WithOKCodes(200).
		WithJSONResponse(resp).
		WithJSONError(errResp)

	if _, sdkErr := s.Client.Get(ctx, url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp).
			AppendCategories(sdkerror.ErrCatProductVlb)
	}

	return resp.ToEntityListGlobalRegions(), nil
}

func (s *GLBServiceV1) GetGlobalLoadBalancerUsageHistories(
	ctx context.Context,
	opts *GetGlobalLoadBalancerUsageHistoriesRequest,
) (*ListGlobalLoadBalancerUsageHistories, error) {
	url := getGlobalLoadBalancerUsageHistoriesURL(s.Client, opts)
	resp := new(ListGlobalLoadBalancerUsageHistories)
	errResp := sdkerror.NewErrorResponse(sdkerror.GlobalLoadBalancerErrorType)
	req := client.NewRequest().
		WithOKCodes(200).
		WithJSONResponse(resp).
		WithJSONError(errResp)

	if _, sdkErr := s.Client.Get(ctx, url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.EcGlobalLoadBalancerNotFound).
			AppendCategories(sdkerror.ErrCatProductVlb)
	}

	return resp, nil
}
