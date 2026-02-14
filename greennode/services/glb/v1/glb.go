package v1

import (
	"fmt"

	"github.com/dannyota/greennode-community-sdk/v2/greennode/client"
	"github.com/dannyota/greennode-community-sdk/v2/greennode/entity"
	sdkerror "github.com/dannyota/greennode-community-sdk/v2/greennode/sdkerror"
)

func (s *GLBServiceV1) ListGlobalPools(popts IListGlobalPoolsRequest) (*entity.ListGlobalPools, sdkerror.IError) {
	url := listGlobalPoolsUrl(s.VLBClient, popts)
	resp := new(ListGlobalPoolsResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.GlobalLoadBalancerErrorType)
	req := client.NewRequest().
		WithHeader("User-Agent", popts.ParseUserAgent()).
		WithOkCodes(200).
		WithJsonResponse(resp).
		WithJsonError(errResp)

	if _, sdkErr := s.VLBClient.Get(url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.WithErrorGlobalLoadBalancerNotFound(errResp)).
			WithKVparameters("loadBalancerId", popts.GetLoadBalancerId()).
			AppendCategories(sdkerror.ErrCatProductVlb)
	}

	return resp.ToEntityListGlobalPools(), nil
}

// --------------------------------------------------

func (s *GLBServiceV1) CreateGlobalPool(popts ICreateGlobalPoolRequest) (*entity.GlobalPool, sdkerror.IError) {
	url := createGlobalPoolUrl(s.VLBClient, popts)
	resp := new(CreateGlobalPoolResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.GlobalLoadBalancerErrorType)
	req := client.NewRequest().
		WithHeader("User-Agent", popts.ParseUserAgent()).
		WithOkCodes(202).
		WithJsonBody(popts.ToRequestBody()).
		WithJsonResponse(resp).
		WithJsonError(errResp)

	if _, sdkErr := s.VLBClient.Post(url, req); sdkErr != nil {
		fmt.Println("sdkErr: ", sdkErr)
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.WithErrorGlobalLoadBalancerNotFound(errResp)).
			WithParameters(popts.ToMap()).
			AppendCategories(sdkerror.ErrCatProductVlb)
	}

	return resp.ToEntityPool(), nil
}

// --------------------------------------------------

func (s *GLBServiceV1) UpdateGlobalPool(popts IUpdateGlobalPoolRequest) (*entity.GlobalPool, sdkerror.IError) {
	url := updateGlobalPoolUrl(s.VLBClient, popts)
	resp := new(UpdateGlobalPoolResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.GlobalLoadBalancerErrorType)
	req := client.NewRequest().
		WithHeader("User-Agent", popts.ParseUserAgent()).
		WithOkCodes(202).
		WithJsonBody(popts.ToRequestBody()).
		WithJsonResponse(resp).
		WithJsonError(errResp)

	if _, sdkErr := s.VLBClient.Put(url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.WithErrorGlobalLoadBalancerNotFound(errResp)).
			WithParameters(popts.ToMap()).
			AppendCategories(sdkerror.ErrCatProductVlb)
	}

	return resp.ToEntityPool(), nil
}

// --------------------------------------------------

func (s *GLBServiceV1) DeleteGlobalPool(popts IDeleteGlobalPoolRequest) sdkerror.IError {
	url := deleteGlobalPoolUrl(s.VLBClient, popts)
	errResp := sdkerror.NewErrorResponse(sdkerror.GlobalLoadBalancerErrorType)
	req := client.NewRequest().
		WithHeader("User-Agent", popts.ParseUserAgent()).
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

func (s *GLBServiceV1) ListGlobalPoolMembers(popts IListGlobalPoolMembersRequest) (*entity.ListGlobalPoolMembers, sdkerror.IError) {
	url := listGlobalPoolMembersUrl(s.VLBClient, popts)
	resp := new(ListGlobalPoolMembersResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.GlobalLoadBalancerErrorType)
	req := client.NewRequest().
		WithHeader("User-Agent", popts.ParseUserAgent()).
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

func (s *GLBServiceV1) GetGlobalPoolMember(popts IGetGlobalPoolMemberRequest) (*entity.GlobalPoolMember, sdkerror.IError) {
	url := getGlobalPoolMemberUrl(s.VLBClient, popts)
	resp := new(GetGlobalPoolMemberResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.GlobalLoadBalancerErrorType)
	req := client.NewRequest().
		WithHeader("User-Agent", popts.ParseUserAgent()).
		WithOkCodes(200).
		WithJsonResponse(resp).
		WithJsonError(errResp)

	if _, sdkErr := s.VLBClient.Get(url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.WithErrorGlobalLoadBalancerNotFound(errResp)).
			WithKVparameters("loadBalancerId", popts.GetLoadBalancerId()).
			WithKVparameters("poolId", popts.GetPoolId()).
			WithKVparameters("poolMemberId", popts.GetPoolMemberId()).
			AppendCategories(sdkerror.ErrCatProductVlb)
	}

	return resp.ToEntityGlobalPoolMember(), nil
}

// --------------------------------------------------

func (s *GLBServiceV1) DeleteGlobalPoolMember(popts IDeleteGlobalPoolMemberRequest) sdkerror.IError {
	url := deleteGlobalPoolMemberUrl(s.VLBClient, popts)
	errResp := sdkerror.NewErrorResponse(sdkerror.GlobalLoadBalancerErrorType)
	req := client.NewRequest().
		WithHeader("User-Agent", popts.ParseUserAgent()).
		WithOkCodes(202).
		WithJsonError(errResp)

	if _, sdkErr := s.VLBClient.Delete(url, req); sdkErr != nil {
		return sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.WithErrorGlobalLoadBalancerNotFound(errResp)).
			WithKVparameters("loadBalancerId", popts.GetLoadBalancerId()).
			WithKVparameters("poolId", popts.GetPoolId()).
			WithKVparameters("poolMemberId", popts.GetPoolMemberId()).
			AppendCategories(sdkerror.ErrCatProductVlb)
	}

	return nil
}

// --------------------------------------------------

func (s *GLBServiceV1) UpdateGlobalPoolMember(popts IUpdateGlobalPoolMemberRequest) (*entity.GlobalPoolMember, sdkerror.IError) {
	url := updateGlobalPoolMemberUrl(s.VLBClient, popts)
	resp := new(UpdateGlobalPoolMemberResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.GlobalLoadBalancerErrorType)
	req := client.NewRequest().
		WithHeader("User-Agent", popts.ParseUserAgent()).
		WithOkCodes(202).
		WithJsonBody(popts.ToRequestBody()).
		WithJsonResponse(resp).
		WithJsonError(errResp)

	if _, sdkErr := s.VLBClient.Put(url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.WithErrorGlobalLoadBalancerNotFound(errResp)).
			WithParameters(popts.ToMap()).
			AppendCategories(sdkerror.ErrCatProductVlb)
	}

	return resp.ToEntityGlobalPoolMember(), nil
}

// --------------------------------------------------

func (s *GLBServiceV1) PatchGlobalPoolMembers(popts IPatchGlobalPoolMembersRequest) sdkerror.IError {
	url := patchGlobalPoolMembersUrl(s.VLBClient, popts)
	errResp := sdkerror.NewErrorResponse(sdkerror.GlobalLoadBalancerErrorType)
	req := client.NewRequest().
		WithHeader("User-Agent", popts.ParseUserAgent()).
		WithOkCodes(202).
		WithJsonBody(popts.ToRequestBody()).
		WithJsonError(errResp)

	if _, sdkErr := s.VLBClient.Patch(url, req); sdkErr != nil {
		return sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.WithErrorGlobalLoadBalancerNotFound(errResp)).
			WithParameters(popts.ToMap()).
			AppendCategories(sdkerror.ErrCatProductVlb)
	}

	return nil
}

// --------------------------------------------------

func (s *GLBServiceV1) ListGlobalListeners(popts IListGlobalListenersRequest) (*entity.ListGlobalListeners, sdkerror.IError) {
	url := listGlobalListenersUrl(s.VLBClient, popts)
	resp := new(ListGlobalListenersResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.GlobalLoadBalancerErrorType)
	req := client.NewRequest().
		WithHeader("User-Agent", popts.ParseUserAgent()).
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

func (s *GLBServiceV1) CreateGlobalListener(popts ICreateGlobalListenerRequest) (*entity.GlobalListener, sdkerror.IError) {
	url := createGlobalListenerUrl(s.VLBClient, popts)
	resp := new(CreateGlobalListenerResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.GlobalLoadBalancerErrorType)
	req := client.NewRequest().
		WithHeader("User-Agent", popts.ParseUserAgent()).
		WithOkCodes(202).
		WithJsonBody(popts.ToRequestBody()).
		WithJsonResponse(resp).
		WithJsonError(errResp)

	if _, sdkErr := s.VLBClient.Post(url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.WithErrorGlobalLoadBalancerNotFound(errResp)).
			WithParameters(popts.ToMap()).
			AppendCategories(sdkerror.ErrCatProductVlb)
	}

	return resp.ToEntityGlobalListener(), nil
}

// --------------------------------------------------

func (s *GLBServiceV1) UpdateGlobalListener(popts IUpdateGlobalListenerRequest) (*entity.GlobalListener, sdkerror.IError) {
	url := updateGlobalListenerUrl(s.VLBClient, popts)
	resp := new(UpdateGlobalListenerResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.GlobalLoadBalancerErrorType)
	req := client.NewRequest().
		WithHeader("User-Agent", popts.ParseUserAgent()).
		WithOkCodes(202).
		WithJsonBody(popts.ToRequestBody()).
		WithJsonResponse(resp).
		WithJsonError(errResp)

	if _, sdkErr := s.VLBClient.Put(url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.WithErrorGlobalLoadBalancerNotFound(errResp)).
			WithParameters(popts.ToMap()).
			AppendCategories(sdkerror.ErrCatProductVlb)

	}

	return resp.ToEntityGlobalListener(), nil
}

// --------------------------------------------------

func (s *GLBServiceV1) GetGlobalListener(popts IGetGlobalListenerRequest) (*entity.GlobalListener, sdkerror.IError) {
	url := getGlobalListenerUrl(s.VLBClient, popts)
	resp := new(GetGlobalListenerResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.GlobalLoadBalancerErrorType)
	req := client.NewRequest().
		WithHeader("User-Agent", popts.ParseUserAgent()).
		WithOkCodes(200).
		WithJsonResponse(resp).
		WithJsonError(errResp)

	if _, sdkErr := s.VLBClient.Get(url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.WithErrorGlobalLoadBalancerNotFound(errResp)).
			WithKVparameters("loadBalancerId", popts.GetLoadBalancerId()).
			WithKVparameters("listenerId", popts.GetListenerId()).
			AppendCategories(sdkerror.ErrCatProductVlb)
	}

	return resp.ToEntityGlobalListener(), nil
}

// --------------------------------------------------

func (s *GLBServiceV1) DeleteGlobalListener(popts IDeleteGlobalListenerRequest) sdkerror.IError {
	url := deleteGlobalListenerUrl(s.VLBClient, popts)
	errResp := sdkerror.NewErrorResponse(sdkerror.GlobalLoadBalancerErrorType)
	req := client.NewRequest().
		WithHeader("User-Agent", popts.ParseUserAgent()).
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

func (s *GLBServiceV1) ListGlobalLoadBalancers(popts IListGlobalLoadBalancersRequest) (*entity.ListGlobalLoadBalancers, sdkerror.IError) {
	url := listGlobalLoadBalancersUrl(s.VLBClient, popts)
	resp := new(ListGlobalLoadBalancersResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.GlobalLoadBalancerErrorType)
	req := client.NewRequest().
		WithHeader("User-Agent", popts.ParseUserAgent()).
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
	popts ICreateGlobalLoadBalancerRequest,
) (*entity.GlobalLoadBalancer, sdkerror.IError) {
	url := createGlobalLoadBalancerUrl(s.VLBClient, popts)
	resp := new(CreateGlobalLoadBalancerResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.GlobalLoadBalancerErrorType)
	req := client.NewRequest().
		WithHeader("User-Agent", popts.ParseUserAgent()).
		WithOkCodes(202).
		WithJsonBody(popts.ToRequestBody()).
		WithJsonResponse(resp).
		WithJsonError(errResp)

	if _, sdkErr := s.VLBClient.Post(url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp).
			WithParameters(popts.ToMap()).
			AppendCategories(sdkerror.ErrCatProductVlb)
	}

	return resp.ToEntityGlobalLoadBalancer(), nil
}

// --------------------------------------------------

func (s *GLBServiceV1) DeleteGlobalLoadBalancer(popts IDeleteGlobalLoadBalancerRequest) sdkerror.IError {
	url := deleteGlobalLoadBalancerUrl(s.VLBClient, popts)
	errResp := sdkerror.NewErrorResponse(sdkerror.GlobalLoadBalancerErrorType)
	req := client.NewRequest().
		WithHeader("User-Agent", popts.ParseUserAgent()).
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
	popts IGetGlobalLoadBalancerByIdRequest,
) (*entity.GlobalLoadBalancer, sdkerror.IError) {
	url := getGlobalLoadBalancerByIdUrl(s.VLBClient, popts)
	resp := new(GetGlobalLoadBalancerByIdResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.GlobalLoadBalancerErrorType)
	req := client.NewRequest().
		WithHeader("User-Agent", popts.ParseUserAgent()).
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

func (s *GLBServiceV1) ListGlobalPackages(popts IListGlobalPackagesRequest) (*entity.ListGlobalPackages, sdkerror.IError) {
	url := listGlobalPackagesUrl(s.VLBClient, popts)
	resp := new(ListGlobalPackagesResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.GlobalLoadBalancerErrorType)
	req := client.NewRequest().
		WithHeader("User-Agent", popts.ParseUserAgent()).
		WithOkCodes(200).
		WithJsonResponse(resp).
		WithJsonError(errResp)

	if _, sdkErr := s.VLBClient.Get(url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp).
			AppendCategories(sdkerror.ErrCatProductVlb)
	}

	return resp.ToEntityListGlobalPackages(), nil
}

func (s *GLBServiceV1) ListGlobalRegions(popts IListGlobalRegionsRequest) (*entity.ListGlobalRegions, sdkerror.IError) {
	url := listGlobalRegionsUrl(s.VLBClient, popts)
	resp := new(ListGlobalRegionsResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.GlobalLoadBalancerErrorType)
	req := client.NewRequest().
		WithHeader("User-Agent", popts.ParseUserAgent()).
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
	popts IGetGlobalLoadBalancerUsageHistoriesRequest,
) (*entity.ListGlobalLoadBalancerUsageHistories, sdkerror.IError) {
	url := getGlobalLoadBalancerUsageHistoriesUrl(s.VLBClient, popts)
	resp := new(GetGlobalLoadBalancerUsageHistoriesResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.GlobalLoadBalancerErrorType)
	req := client.NewRequest().
		WithHeader("User-Agent", popts.ParseUserAgent()).
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
