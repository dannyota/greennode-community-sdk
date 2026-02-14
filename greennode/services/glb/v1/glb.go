package v1

import (
	"fmt"

	"github.com/dannyota/greennode-community-sdk/v2/greennode/client"
	"github.com/dannyota/greennode-community-sdk/v2/greennode/entity"
	sdkerror "github.com/dannyota/greennode-community-sdk/v2/greennode/sdkerror"
)

func (s *GLBServiceV1) ListGlobalPools(opts IListGlobalPoolsRequest) (*entity.ListGlobalPools, sdkerror.Error) {
	url := listGlobalPoolsUrl(s.VLBClient, opts)
	resp := new(ListGlobalPoolsResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.GlobalLoadBalancerErrorType)
	req := client.NewRequest().
		WithHeader("User-Agent", opts.ParseUserAgent()).
		WithOkCodes(200).
		WithJsonResponse(resp).
		WithJsonError(errResp)

	if _, sdkErr := s.VLBClient.Get(url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.WithErrorGlobalLoadBalancerNotFound(errResp)).
			WithKVparameters("loadBalancerId", opts.GetLoadBalancerId()).
			AppendCategories(sdkerror.ErrCatProductVlb)
	}

	return resp.ToEntityListGlobalPools(), nil
}

// --------------------------------------------------

func (s *GLBServiceV1) CreateGlobalPool(opts ICreateGlobalPoolRequest) (*entity.GlobalPool, sdkerror.Error) {
	url := createGlobalPoolUrl(s.VLBClient, opts)
	resp := new(CreateGlobalPoolResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.GlobalLoadBalancerErrorType)
	req := client.NewRequest().
		WithHeader("User-Agent", opts.ParseUserAgent()).
		WithOkCodes(202).
		WithJsonBody(opts.ToRequestBody()).
		WithJsonResponse(resp).
		WithJsonError(errResp)

	if _, sdkErr := s.VLBClient.Post(url, req); sdkErr != nil {
		fmt.Println("sdkErr: ", sdkErr)
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.WithErrorGlobalLoadBalancerNotFound(errResp)).
			WithParameters(opts.ToMap()).
			AppendCategories(sdkerror.ErrCatProductVlb)
	}

	return resp.ToEntityPool(), nil
}

// --------------------------------------------------

func (s *GLBServiceV1) UpdateGlobalPool(opts IUpdateGlobalPoolRequest) (*entity.GlobalPool, sdkerror.Error) {
	url := updateGlobalPoolUrl(s.VLBClient, opts)
	resp := new(UpdateGlobalPoolResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.GlobalLoadBalancerErrorType)
	req := client.NewRequest().
		WithHeader("User-Agent", opts.ParseUserAgent()).
		WithOkCodes(202).
		WithJsonBody(opts.ToRequestBody()).
		WithJsonResponse(resp).
		WithJsonError(errResp)

	if _, sdkErr := s.VLBClient.Put(url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.WithErrorGlobalLoadBalancerNotFound(errResp)).
			WithParameters(opts.ToMap()).
			AppendCategories(sdkerror.ErrCatProductVlb)
	}

	return resp.ToEntityPool(), nil
}

// --------------------------------------------------

func (s *GLBServiceV1) DeleteGlobalPool(opts IDeleteGlobalPoolRequest) sdkerror.Error {
	url := deleteGlobalPoolUrl(s.VLBClient, opts)
	errResp := sdkerror.NewErrorResponse(sdkerror.GlobalLoadBalancerErrorType)
	req := client.NewRequest().
		WithHeader("User-Agent", opts.ParseUserAgent()).
		WithOkCodes(202).
		WithJsonError(errResp)

	if _, sdkErr := s.VLBClient.Delete(url, req); sdkErr != nil {
		return sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.WithErrorGlobalLoadBalancerNotFound(errResp)).
			AppendCategories(sdkerror.ErrCatProductVlb)
	}

	return nil
}

// --------------------------------------------------

func (s *GLBServiceV1) ListGlobalPoolMembers(opts IListGlobalPoolMembersRequest) (*entity.ListGlobalPoolMembers, sdkerror.Error) {
	url := listGlobalPoolMembersUrl(s.VLBClient, opts)
	resp := new(ListGlobalPoolMembersResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.GlobalLoadBalancerErrorType)
	req := client.NewRequest().
		WithHeader("User-Agent", opts.ParseUserAgent()).
		WithOkCodes(200).
		WithJsonResponse(resp).
		WithJsonError(errResp)

	if _, sdkErr := s.VLBClient.Get(url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.WithErrorGlobalLoadBalancerNotFound(errResp)).
			AppendCategories(sdkerror.ErrCatProductVlb)
	}

	return resp.ToEntityListGlobalPoolMembers(), nil
}

// --------------------------------------------------

func (s *GLBServiceV1) GetGlobalPoolMember(opts IGetGlobalPoolMemberRequest) (*entity.GlobalPoolMember, sdkerror.Error) {
	url := getGlobalPoolMemberUrl(s.VLBClient, opts)
	resp := new(GetGlobalPoolMemberResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.GlobalLoadBalancerErrorType)
	req := client.NewRequest().
		WithHeader("User-Agent", opts.ParseUserAgent()).
		WithOkCodes(200).
		WithJsonResponse(resp).
		WithJsonError(errResp)

	if _, sdkErr := s.VLBClient.Get(url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.WithErrorGlobalLoadBalancerNotFound(errResp)).
			WithKVparameters("loadBalancerId", opts.GetLoadBalancerId()).
			WithKVparameters("poolId", opts.GetPoolId()).
			WithKVparameters("poolMemberId", opts.GetPoolMemberId()).
			AppendCategories(sdkerror.ErrCatProductVlb)
	}

	return resp.ToEntityGlobalPoolMember(), nil
}

// --------------------------------------------------

func (s *GLBServiceV1) DeleteGlobalPoolMember(opts IDeleteGlobalPoolMemberRequest) sdkerror.Error {
	url := deleteGlobalPoolMemberUrl(s.VLBClient, opts)
	errResp := sdkerror.NewErrorResponse(sdkerror.GlobalLoadBalancerErrorType)
	req := client.NewRequest().
		WithHeader("User-Agent", opts.ParseUserAgent()).
		WithOkCodes(202).
		WithJsonError(errResp)

	if _, sdkErr := s.VLBClient.Delete(url, req); sdkErr != nil {
		return sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.WithErrorGlobalLoadBalancerNotFound(errResp)).
			WithKVparameters("loadBalancerId", opts.GetLoadBalancerId()).
			WithKVparameters("poolId", opts.GetPoolId()).
			WithKVparameters("poolMemberId", opts.GetPoolMemberId()).
			AppendCategories(sdkerror.ErrCatProductVlb)
	}

	return nil
}

// --------------------------------------------------

func (s *GLBServiceV1) UpdateGlobalPoolMember(opts IUpdateGlobalPoolMemberRequest) (*entity.GlobalPoolMember, sdkerror.Error) {
	url := updateGlobalPoolMemberUrl(s.VLBClient, opts)
	resp := new(UpdateGlobalPoolMemberResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.GlobalLoadBalancerErrorType)
	req := client.NewRequest().
		WithHeader("User-Agent", opts.ParseUserAgent()).
		WithOkCodes(202).
		WithJsonBody(opts.ToRequestBody()).
		WithJsonResponse(resp).
		WithJsonError(errResp)

	if _, sdkErr := s.VLBClient.Put(url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.WithErrorGlobalLoadBalancerNotFound(errResp)).
			WithParameters(opts.ToMap()).
			AppendCategories(sdkerror.ErrCatProductVlb)
	}

	return resp.ToEntityGlobalPoolMember(), nil
}

// --------------------------------------------------

func (s *GLBServiceV1) PatchGlobalPoolMembers(opts IPatchGlobalPoolMembersRequest) sdkerror.Error {
	url := patchGlobalPoolMembersUrl(s.VLBClient, opts)
	errResp := sdkerror.NewErrorResponse(sdkerror.GlobalLoadBalancerErrorType)
	req := client.NewRequest().
		WithHeader("User-Agent", opts.ParseUserAgent()).
		WithOkCodes(202).
		WithJsonBody(opts.ToRequestBody()).
		WithJsonError(errResp)

	if _, sdkErr := s.VLBClient.Patch(url, req); sdkErr != nil {
		return sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.WithErrorGlobalLoadBalancerNotFound(errResp)).
			WithParameters(opts.ToMap()).
			AppendCategories(sdkerror.ErrCatProductVlb)
	}

	return nil
}

// --------------------------------------------------

func (s *GLBServiceV1) ListGlobalListeners(opts IListGlobalListenersRequest) (*entity.ListGlobalListeners, sdkerror.Error) {
	url := listGlobalListenersUrl(s.VLBClient, opts)
	resp := new(ListGlobalListenersResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.GlobalLoadBalancerErrorType)
	req := client.NewRequest().
		WithHeader("User-Agent", opts.ParseUserAgent()).
		WithOkCodes(200).
		WithJsonResponse(resp).
		WithJsonError(errResp)

	if _, sdkErr := s.VLBClient.Get(url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.WithErrorGlobalLoadBalancerNotFound(errResp)).
			AppendCategories(sdkerror.ErrCatProductVlb)
	}

	return resp.ToEntityListGlobalListeners(), nil
}

// --------------------------------------------------

func (s *GLBServiceV1) CreateGlobalListener(opts ICreateGlobalListenerRequest) (*entity.GlobalListener, sdkerror.Error) {
	url := createGlobalListenerUrl(s.VLBClient, opts)
	resp := new(CreateGlobalListenerResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.GlobalLoadBalancerErrorType)
	req := client.NewRequest().
		WithHeader("User-Agent", opts.ParseUserAgent()).
		WithOkCodes(202).
		WithJsonBody(opts.ToRequestBody()).
		WithJsonResponse(resp).
		WithJsonError(errResp)

	if _, sdkErr := s.VLBClient.Post(url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.WithErrorGlobalLoadBalancerNotFound(errResp)).
			WithParameters(opts.ToMap()).
			AppendCategories(sdkerror.ErrCatProductVlb)
	}

	return resp.ToEntityGlobalListener(), nil
}

// --------------------------------------------------

func (s *GLBServiceV1) UpdateGlobalListener(opts IUpdateGlobalListenerRequest) (*entity.GlobalListener, sdkerror.Error) {
	url := updateGlobalListenerUrl(s.VLBClient, opts)
	resp := new(UpdateGlobalListenerResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.GlobalLoadBalancerErrorType)
	req := client.NewRequest().
		WithHeader("User-Agent", opts.ParseUserAgent()).
		WithOkCodes(202).
		WithJsonBody(opts.ToRequestBody()).
		WithJsonResponse(resp).
		WithJsonError(errResp)

	if _, sdkErr := s.VLBClient.Put(url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.WithErrorGlobalLoadBalancerNotFound(errResp)).
			WithParameters(opts.ToMap()).
			AppendCategories(sdkerror.ErrCatProductVlb)

	}

	return resp.ToEntityGlobalListener(), nil
}

// --------------------------------------------------

func (s *GLBServiceV1) GetGlobalListener(opts IGetGlobalListenerRequest) (*entity.GlobalListener, sdkerror.Error) {
	url := getGlobalListenerUrl(s.VLBClient, opts)
	resp := new(GetGlobalListenerResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.GlobalLoadBalancerErrorType)
	req := client.NewRequest().
		WithHeader("User-Agent", opts.ParseUserAgent()).
		WithOkCodes(200).
		WithJsonResponse(resp).
		WithJsonError(errResp)

	if _, sdkErr := s.VLBClient.Get(url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.WithErrorGlobalLoadBalancerNotFound(errResp)).
			WithKVparameters("loadBalancerId", opts.GetLoadBalancerId()).
			WithKVparameters("listenerId", opts.GetListenerId()).
			AppendCategories(sdkerror.ErrCatProductVlb)
	}

	return resp.ToEntityGlobalListener(), nil
}

// --------------------------------------------------

func (s *GLBServiceV1) DeleteGlobalListener(opts IDeleteGlobalListenerRequest) sdkerror.Error {
	url := deleteGlobalListenerUrl(s.VLBClient, opts)
	errResp := sdkerror.NewErrorResponse(sdkerror.GlobalLoadBalancerErrorType)
	req := client.NewRequest().
		WithHeader("User-Agent", opts.ParseUserAgent()).
		WithOkCodes(202).
		WithJsonError(errResp)

	if _, sdkErr := s.VLBClient.Delete(url, req); sdkErr != nil {
		return sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.WithErrorGlobalLoadBalancerNotFound(errResp)).
			AppendCategories(sdkerror.ErrCatProductVlb)
	}

	return nil
}

// --------------------------------------------------

func (s *GLBServiceV1) ListGlobalLoadBalancers(opts IListGlobalLoadBalancersRequest) (*entity.ListGlobalLoadBalancers, sdkerror.Error) {
	url := listGlobalLoadBalancersUrl(s.VLBClient, opts)
	resp := new(ListGlobalLoadBalancersResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.GlobalLoadBalancerErrorType)
	req := client.NewRequest().
		WithHeader("User-Agent", opts.ParseUserAgent()).
		WithOkCodes(200).
		WithJsonResponse(resp).
		WithJsonError(errResp)

	if _, sdkErr := s.VLBClient.Get(url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp).
			AppendCategories(sdkerror.ErrCatProductVlb)
	}

	return resp.ToEntityListGlobalLoadBalancers(), nil
}

// --------------------------------------------------

func (s *GLBServiceV1) CreateGlobalLoadBalancer(
	opts ICreateGlobalLoadBalancerRequest,
) (*entity.GlobalLoadBalancer, sdkerror.Error) {
	url := createGlobalLoadBalancerUrl(s.VLBClient, opts)
	resp := new(CreateGlobalLoadBalancerResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.GlobalLoadBalancerErrorType)
	req := client.NewRequest().
		WithHeader("User-Agent", opts.ParseUserAgent()).
		WithOkCodes(202).
		WithJsonBody(opts.ToRequestBody()).
		WithJsonResponse(resp).
		WithJsonError(errResp)

	if _, sdkErr := s.VLBClient.Post(url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp).
			WithParameters(opts.ToMap()).
			AppendCategories(sdkerror.ErrCatProductVlb)
	}

	return resp.ToEntityGlobalLoadBalancer(), nil
}

// --------------------------------------------------

func (s *GLBServiceV1) DeleteGlobalLoadBalancer(opts IDeleteGlobalLoadBalancerRequest) sdkerror.Error {
	url := deleteGlobalLoadBalancerUrl(s.VLBClient, opts)
	errResp := sdkerror.NewErrorResponse(sdkerror.GlobalLoadBalancerErrorType)
	req := client.NewRequest().
		WithHeader("User-Agent", opts.ParseUserAgent()).
		WithOkCodes(202).
		WithJsonError(errResp)

	if _, sdkErr := s.VLBClient.Delete(url, req); sdkErr != nil {
		return sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.WithErrorGlobalLoadBalancerNotFound(errResp)).
			AppendCategories(sdkerror.ErrCatProductVlb)
	}

	return nil
}

// --------------------------------------------------

func (s *GLBServiceV1) GetGlobalLoadBalancerById(
	opts IGetGlobalLoadBalancerByIdRequest,
) (*entity.GlobalLoadBalancer, sdkerror.Error) {
	url := getGlobalLoadBalancerByIdUrl(s.VLBClient, opts)
	resp := new(GetGlobalLoadBalancerByIdResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.GlobalLoadBalancerErrorType)
	req := client.NewRequest().
		WithHeader("User-Agent", opts.ParseUserAgent()).
		WithOkCodes(200).
		WithJsonResponse(resp).
		WithJsonError(errResp)

	if _, sdkErr := s.VLBClient.Get(url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.WithErrorGlobalLoadBalancerNotFound(errResp)).
			AppendCategories(sdkerror.ErrCatProductVlb)
	}

	return resp.ToEntityGlobalLoadBalancer(), nil
}

func (s *GLBServiceV1) ListGlobalPackages(opts IListGlobalPackagesRequest) (*entity.ListGlobalPackages, sdkerror.Error) {
	url := listGlobalPackagesUrl(s.VLBClient, opts)
	resp := new(ListGlobalPackagesResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.GlobalLoadBalancerErrorType)
	req := client.NewRequest().
		WithHeader("User-Agent", opts.ParseUserAgent()).
		WithOkCodes(200).
		WithJsonResponse(resp).
		WithJsonError(errResp)

	if _, sdkErr := s.VLBClient.Get(url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp).
			AppendCategories(sdkerror.ErrCatProductVlb)
	}

	return resp.ToEntityListGlobalPackages(), nil
}

func (s *GLBServiceV1) ListGlobalRegions(opts IListGlobalRegionsRequest) (*entity.ListGlobalRegions, sdkerror.Error) {
	url := listGlobalRegionsUrl(s.VLBClient, opts)
	resp := new(ListGlobalRegionsResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.GlobalLoadBalancerErrorType)
	req := client.NewRequest().
		WithHeader("User-Agent", opts.ParseUserAgent()).
		WithOkCodes(200).
		WithJsonResponse(resp).
		WithJsonError(errResp)

	if _, sdkErr := s.VLBClient.Get(url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp).
			AppendCategories(sdkerror.ErrCatProductVlb)
	}

	return resp.ToEntityListGlobalRegions(), nil
}

func (s *GLBServiceV1) GetGlobalLoadBalancerUsageHistories(
	opts IGetGlobalLoadBalancerUsageHistoriesRequest,
) (*entity.ListGlobalLoadBalancerUsageHistories, sdkerror.Error) {
	url := getGlobalLoadBalancerUsageHistoriesUrl(s.VLBClient, opts)
	resp := new(GetGlobalLoadBalancerUsageHistoriesResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.GlobalLoadBalancerErrorType)
	req := client.NewRequest().
		WithHeader("User-Agent", opts.ParseUserAgent()).
		WithOkCodes(200).
		WithJsonResponse(resp).
		WithJsonError(errResp)

	if _, sdkErr := s.VLBClient.Get(url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.WithErrorGlobalLoadBalancerNotFound(errResp)).
			AppendCategories(sdkerror.ErrCatProductVlb)
	}

	return resp.ToEntityGlobalLoadBalancerUsageHistories(), nil
}
