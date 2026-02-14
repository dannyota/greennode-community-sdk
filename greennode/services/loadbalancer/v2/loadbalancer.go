package v2

import (
	"github.com/dannyota/greennode-community-sdk/v2/greennode/client"
	"github.com/dannyota/greennode-community-sdk/v2/greennode/entity"
	sdkerror "github.com/dannyota/greennode-community-sdk/v2/greennode/sdkerror"
)

func (s *LoadBalancerServiceV2) CreateLoadBalancer(popts ICreateLoadBalancerRequest) (*entity.LoadBalancer, sdkerror.IError) {
	url := createLoadBalancerUrl(s.VLBClient)
	resp := new(CreateLoadBalancerResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithHeader("User-Agent", popts.ParseUserAgent()).
		WithOkCodes(202).
		WithJsonBody(popts.ToRequestBody()).
		WithJsonResponse(resp).
		WithJsonError(errResp)

	if _, sdkErr := s.VLBClient.Post(url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.WithErrorLoadBalancerExceedQuota(errResp)).
			WithParameters(popts.ToMap()).
			AppendCategories(sdkerror.ErrCatProductVlb)
	}

	return resp.ToEntityLoadBalancer(), nil
}

func (s *LoadBalancerServiceV2) ResizeLoadBalancer(popts IResizeLoadBalancerRequest) (*entity.LoadBalancer, sdkerror.IError) {
	url := resizeLoadBalancerUrl(s.VLBClient, popts)
	resp := new(ResizeLoadBalancerResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithHeader("User-Agent", popts.ParseUserAgent()).
		WithOkCodes(202).
		WithJsonBody(popts.ToRequestBody()).
		WithJsonResponse(resp).
		WithJsonError(errResp)

	if _, sdkErr := s.VLBClient.Put(url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.WithErrorLoadBalancerNotFound(errResp),
			sdkerror.WithErrorLoadBalancerNotReady(errResp))
	}

	return resp.ToEntityLoadBalancer(), nil
}

func (s *LoadBalancerServiceV2) ListLoadBalancerPackages(popts IListLoadBalancerPackagesRequest) (*entity.ListLoadBalancerPackages, sdkerror.IError) {
	url := listLoadBalancerPackagesUrl(s.VLBClient, popts)
	resp := new(ListLoadBalancerPackagesResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithHeader("User-Agent", popts.ParseUserAgent()).
		WithOkCodes(200).
		WithJsonResponse(resp).
		WithJsonError(errResp)

	if _, sdkErr := s.VLBClient.Get(url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp)
	}

	return resp.ToEntityListLoadBalancerPackages(), nil
}

func (s *LoadBalancerServiceV2) GetLoadBalancerById(popts IGetLoadBalancerByIdRequest) (*entity.LoadBalancer, sdkerror.IError) {
	url := getLoadBalancerByIdUrl(s.VLBClient, popts)
	resp := new(GetLoadBalancerByIdResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithHeader("User-Agent", popts.ParseUserAgent()).
		WithOkCodes(200).
		WithJsonResponse(resp).
		WithJsonError(errResp)

	if _, sdkErr := s.VLBClient.Get(url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.WithErrorLoadBalancerNotFound(errResp)).
			WithKVparameters("loadBalancerId", popts.GetLoadBalancerId()).
			AppendCategories(sdkerror.ErrCatProductVlb)
	}

	return resp.ToEntityLoadBalancer(), nil
}

func (s *LoadBalancerServiceV2) ListLoadBalancers(popts IListLoadBalancersRequest) (*entity.ListLoadBalancers, sdkerror.IError) {
	url := listLoadBalancersUrl(s.VLBClient, popts)
	resp := new(ListLoadBalancersResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithHeader("User-Agent", popts.ParseUserAgent()).
		WithOkCodes(200).
		WithJsonResponse(resp).
		WithJsonError(errResp)

	if _, sdkErr := s.VLBClient.Get(url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp).
			AppendCategories(sdkerror.ErrCatProductVlb)
	}

	return resp.ToEntityListLoadBalancers(), nil
}

func (s *LoadBalancerServiceV2) GetPoolHealthMonitorById(popts IGetPoolHealthMonitorByIdRequest) (*entity.HealthMonitor, sdkerror.IError) {
	url := getPoolHealthMonitorByIdUrl(s.VLBClient, popts)
	resp := new(GetPoolHealthMonitorByIdResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithHeader("User-Agent", popts.ParseUserAgent()).
		WithOkCodes(200).
		WithJsonResponse(resp).
		WithJsonError(errResp)

	if _, sdkErr := s.VLBClient.Get(url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.WithErrorLoadBalancerNotFound(errResp),
			sdkerror.WithErrorPoolNotFound(errResp))
	}

	return resp.ToEntityHealthMonitor(), nil
}

func (s *LoadBalancerServiceV2) CreatePool(popts ICreatePoolRequest) (*entity.Pool, sdkerror.IError) {
	url := createPoolUrl(s.VLBClient, popts)
	resp := new(CreatePoolResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithHeader("User-Agent", popts.ParseUserAgent()).
		WithOkCodes(202).
		WithJsonBody(popts.ToRequestBody()).
		WithJsonResponse(resp).
		WithJsonError(errResp)

	if _, sdkErr := s.VLBClient.Post(url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.WithErrorLoadBalancerNotFound(errResp),
			sdkerror.WithErrorLoadBalancerNotFound2(errResp),
			sdkerror.WithErrorLoadBalancerNotReady(errResp),
			sdkerror.WithErrorLoadBalancerDuplicatePoolName(errResp)).
			WithParameters(popts.ToMap()).
			AppendCategories(sdkerror.ErrCatProductVlb)
	}

	return resp.ToEntityPool(), nil
}

func (s *LoadBalancerServiceV2) UpdatePool(popts IUpdatePoolRequest) sdkerror.IError {
	url := updatePoolUrl(s.VLBClient, popts)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithHeader("User-Agent", popts.ParseUserAgent()).
		WithOkCodes(202).
		WithJsonBody(popts.ToRequestBody()).
		WithJsonError(errResp)

	if _, sdkErr := s.VLBClient.Put(url, req); sdkErr != nil {
		return sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.WithErrorLoadBalancerNotFound2(errResp),
			sdkerror.WithErrorLoadBalancerNotReady(errResp),
			sdkerror.WithErrorListenerNotFound(errResp)).
			AppendCategories(sdkerror.ErrCatProductVlb)
	}
	return nil
}

func (s *LoadBalancerServiceV2) CreateListener(popts ICreateListenerRequest) (*entity.Listener, sdkerror.IError) {
	url := createListenerUrl(s.VLBClient, popts)
	resp := new(CreateListenerResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithHeader("User-Agent", popts.ParseUserAgent()).
		WithOkCodes(202).
		WithJsonBody(popts.ToRequestBody()).
		WithJsonResponse(resp).
		WithJsonError(errResp)

	if _, sdkErr := s.VLBClient.Post(url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.WithErrorLoadBalancerNotFound2(errResp),
			sdkerror.WithErrorLoadBalancerNotReady(errResp),
			sdkerror.WithErrorListenerDuplicateName(errResp),
			sdkerror.WithErrorPoolNotFound(errResp),
			sdkerror.WithErrorListenerDuplicateProtocolOrPort(errResp)).
			WithParameters(popts.ToMap()).
			AppendCategories(sdkerror.ErrCatProductVlb)
	}

	return resp.ToEntityListener(), nil
}

func (s *LoadBalancerServiceV2) UpdateListener(popts IUpdateListenerRequest) sdkerror.IError {
	url := updateListenerUrl(s.VLBClient, popts)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithHeader("User-Agent", popts.ParseUserAgent()).
		WithOkCodes(202).
		WithJsonBody(popts.ToRequestBody()).
		WithJsonError(errResp)

	if _, sdkErr := s.VLBClient.Put(url, req); sdkErr != nil {
		return sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.WithErrorLoadBalancerNotFound2(errResp),
			sdkerror.WithErrorLoadBalancerNotReady(errResp),
			sdkerror.WithErrorListenerNotFound(errResp)).
			AppendCategories(sdkerror.ErrCatProductVlb)
	}

	return nil
}

func (s *LoadBalancerServiceV2) ListListenersByLoadBalancerId(popts IListListenersByLoadBalancerIdRequest) (*entity.ListListeners, sdkerror.IError) {
	url := listListenersByLoadBalancerIdUrl(s.VLBClient, popts)
	resp := new(ListListenersByLoadBalancerIdResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithHeader("User-Agent", popts.ParseUserAgent()).
		WithOkCodes(200).
		WithJsonResponse(resp).
		WithJsonError(errResp)

	if _, sdkErr := s.VLBClient.Get(url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.WithErrorLoadBalancerNotFound(errResp)).
			WithKVparameters("loadBalancerId", popts.GetLoadBalancerId()).
			AppendCategories(sdkerror.ErrCatProductVlb)
	}

	return resp.ToEntityListListeners(), nil
}

func (s *LoadBalancerServiceV2) ListPoolsByLoadBalancerId(popts IListPoolsByLoadBalancerIdRequest) (*entity.ListPools, sdkerror.IError) {
	url := listPoolsByLoadBalancerIdUrl(s.VLBClient, popts)
	resp := new(ListPoolsByLoadBalancerIdResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithHeader("User-Agent", popts.ParseUserAgent()).
		WithOkCodes(200).
		WithJsonResponse(resp).
		WithJsonError(errResp)

	if _, sdkErr := s.VLBClient.Get(url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.WithErrorLoadBalancerNotFound(errResp)).
			WithKVparameters("loadBalancerId", popts.GetLoadBalancerId()).
			AppendCategories(sdkerror.ErrCatProductVlb)
	}

	return resp.ToEntityListPools(), nil
}

func (s *LoadBalancerServiceV2) UpdatePoolMembers(popts IUpdatePoolMembersRequest) sdkerror.IError {
	url := updatePoolMembersUrl(s.VLBClient, popts)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithHeader("User-Agent", popts.ParseUserAgent()).
		WithOkCodes(202).
		WithJsonBody(popts.ToRequestBody()).
		WithJsonError(errResp)

	if _, sdkErr := s.VLBClient.Put(url, req); sdkErr != nil {
		return sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.WithErrorLoadBalancerNotFound2(errResp),
			sdkerror.WithErrorLoadBalancerNotReady(errResp),
			sdkerror.WithErrorPoolNotFound(errResp),
			sdkerror.WithErrorMemberMustIdentical(errResp)).
			AppendCategories(sdkerror.ErrCatProductVlb)
	}

	return nil
}

func (s *LoadBalancerServiceV2) ListPoolMembers(popts IListPoolMembersRequest) (*entity.ListMembers, sdkerror.IError) {
	url := listPoolMembersUrl(s.VLBClient, popts)
	resp := new(ListPoolMembersResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithHeader("User-Agent", popts.ParseUserAgent()).
		WithOkCodes(200).
		WithJsonResponse(resp).
		WithJsonError(errResp)

	if _, sdkErr := s.VLBClient.Get(url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.WithErrorLoadBalancerNotFound(errResp),
			sdkerror.WithErrorPoolNotFound(errResp)).
			WithKVparameters("loadBalancerId", popts.GetLoadBalancerId(), "poolId", popts.GetPoolId()).
			AppendCategories(sdkerror.ErrCatProductVlb)
	}

	return resp.ToEntityListMembers(), nil
}

func (s *LoadBalancerServiceV2) DeletePoolById(popts IDeletePoolByIdRequest) sdkerror.IError {
	url := deletePoolByIdUrl(s.VLBClient, popts)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithHeader("User-Agent", popts.ParseUserAgent()).
		WithOkCodes(202).
		WithJsonError(errResp)

	if _, sdkErr := s.VLBClient.Delete(url, req); sdkErr != nil {
		return sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.WithErrorPoolInUse(errResp),
			sdkerror.WithErrorLoadBalancerNotReady(errResp),
			sdkerror.WithErrorPoolNotFound(errResp)).
			AppendCategories(sdkerror.ErrCatProductVlb)
	}

	return nil
}

func (s *LoadBalancerServiceV2) DeleteListenerById(popts IDeleteListenerByIdRequest) sdkerror.IError {
	url := deleteListenerByIdUrl(s.VLBClient, popts)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithHeader("User-Agent", popts.ParseUserAgent()).
		WithOkCodes(202).
		WithJsonError(errResp)

	if _, sdkErr := s.VLBClient.Delete(url, req); sdkErr != nil {
		return sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.WithErrorLoadBalancerNotFound2(errResp),
			sdkerror.WithErrorListenerNotFound(errResp),
			sdkerror.WithErrorLoadBalancerNotReady(errResp)).
			AppendCategories(sdkerror.ErrCatProductVlb)
	}

	return nil
}

func (s *LoadBalancerServiceV2) DeleteLoadBalancerById(popts IDeleteLoadBalancerByIdRequest) sdkerror.IError {
	url := deleteLoadBalancerByIdUrl(s.VLBClient, popts)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithHeader("User-Agent", popts.ParseUserAgent()).
		WithOkCodes(202).
		WithJsonError(errResp)

	if _, sdkErr := s.VLBClient.Delete(url, req); sdkErr != nil {
		return sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.WithErrorLoadBalancerNotFound(errResp),
			sdkerror.WithErrorLoadBalancerNotReady(errResp),
			sdkerror.WithErrorLoadBalancerIsCreating(errResp),
			sdkerror.WithErrorLoadBalancerIsDeleting(errResp)).
			WithKVparameters(
				"loadBalancerId", popts.GetLoadBalancerId(),
				"projectId", s.getProjectId()).
			AppendCategories(sdkerror.ErrCatProductVlb)
	}

	return nil
}

func (s *LoadBalancerServiceV2) GetPoolById(popts IGetPoolByIdRequest) (*entity.Pool, sdkerror.IError) {
	url := getPoolByIdUrl(s.VLBClient, popts)
	resp := new(GetPoolByIdResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithHeader("User-Agent", popts.ParseUserAgent()).
		WithOkCodes(200).
		WithJsonResponse(resp).
		WithJsonError(errResp)

	if _, sdkErr := s.VLBClient.Get(url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.WithErrorLoadBalancerNotFound(errResp),
			sdkerror.WithErrorPoolNotFound(errResp)).
			WithKVparameters(
				"loadBalancerId", popts.GetLoadBalancerId(),
				"poolId", popts.GetPoolId()).
			AppendCategories(sdkerror.ErrCatProductVlb)
	}

	return resp.ToEntityPool(), nil
}

func (s *LoadBalancerServiceV2) GetListenerById(popts IGetListenerByIdRequest) (*entity.Listener, sdkerror.IError) {
	url := getListenerByIdUrl(s.VLBClient, popts)
	resp := new(GetListenerByIdResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithHeader("User-Agent", popts.ParseUserAgent()).
		WithOkCodes(200).
		WithJsonResponse(resp).
		WithJsonError(errResp)

	if _, sdkErr := s.VLBClient.Get(url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.WithErrorLoadBalancerNotFound(errResp),
			sdkerror.WithErrorListenerNotFound(errResp)).
			WithKVparameters(
				"loadBalancerId", popts.GetLoadBalancerId(),
				"listenerId", popts.GetListenerId()).
			AppendCategories(sdkerror.ErrCatProductVlb)
	}

	return resp.ToEntityListener(), nil
}

func (s *LoadBalancerServiceV2) ResizeLoadBalancerById(popts IResizeLoadBalancerByIdRequest) sdkerror.IError {
	url := resizeLoadBalancerByIdUrl(s.VLBClient, popts)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithHeader("User-Agent", popts.ParseUserAgent()).
		WithOkCodes(202).
		WithJsonBody(popts.ToRequestBody()).
		WithJsonError(errResp)

	if _, sdkErr := s.VLBClient.Put(url, req); sdkErr != nil {
		return sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.WithErrorLoadBalancerPackageNotFound(errResp),
			sdkerror.WithErrorLoadBalancerNotFound(errResp),
			sdkerror.WithErrorLoadBalancerNotFound2(errResp),
			sdkerror.WithErrorLoadBalancerNotReady(errResp),
			sdkerror.WithErrorLoadBalancerResizeSamePackage(errResp)).
			WithParameters(popts.ToMap()).
			AppendCategories(sdkerror.ErrCatProductVlb)
	}

	return nil
}

func (s *LoadBalancerServiceV2) ScaleLoadBalancer(popts IScaleLoadBalancerRequest) (*entity.LoadBalancer, sdkerror.IError) {
	url := scaleLoadBalancerUrl(s.VLBClient, popts)
	resp := new(ScaleLoadBalancerResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithHeader("User-Agent", popts.ParseUserAgent()).
		WithOkCodes(202).
		WithJsonBody(popts.ToRequestBody()).
		WithJsonResponse(resp).
		WithJsonError(errResp)

	if _, sdkErr := s.VLBClient.Put(url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.WithErrorLoadBalancerNotFound(errResp),
			sdkerror.WithErrorLoadBalancerNotFound2(errResp),
			sdkerror.WithErrorLoadBalancerNotReady(errResp)).
			WithParameters(popts.ToMap()).
			AppendCategories(sdkerror.ErrCatProductVlb)
	}

	return resp.ToEntityLoadBalancer(), nil
}

// policy

func (s *LoadBalancerServiceV2) ListPolicies(popts IListPoliciesRequest) (*entity.ListPolicies, sdkerror.IError) {
	url := listPoliciesUrl(s.VLBClient, popts)
	resp := new(ListPoliciesResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithHeader("User-Agent", popts.ParseUserAgent()).
		WithOkCodes(200).
		WithJsonResponse(resp).
		WithJsonError(errResp)

	if _, sdkErr := s.VLBClient.Get(url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp)
	}

	return resp.ToEntityListPolicies(), nil
}

func (s *LoadBalancerServiceV2) CreatePolicy(popts ICreatePolicyRequest) (*entity.Policy, sdkerror.IError) {
	url := createPolicyUrl(s.VLBClient, popts)
	resp := new(CreatePolicyResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithHeader("User-Agent", popts.ParseUserAgent()).
		WithOkCodes(202).
		WithJsonBody(popts.ToRequestBody()).
		WithJsonResponse(resp).
		WithJsonError(errResp)

	if _, sdkErr := s.VLBClient.Post(url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.WithErrorLoadBalancerNotFound(errResp),
			sdkerror.WithErrorListenerNotFound(errResp),
		).WithParameters(popts.ToMap())
	}

	return resp.ToEntityPolicy(), nil
}

func (s *LoadBalancerServiceV2) GetPolicyById(popts IGetPolicyByIdRequest) (*entity.Policy, sdkerror.IError) {
	url := getPolicyByIdUrl(s.VLBClient, popts)
	resp := new(GetPolicyResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithHeader("User-Agent", popts.ParseUserAgent()).
		WithOkCodes(200).
		WithJsonResponse(resp).
		WithJsonError(errResp)

	if _, sdkErr := s.VLBClient.Get(url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.WithErrorLoadBalancerNotFound(errResp),
			sdkerror.WithErrorListenerNotFound(errResp),
		).WithKVparameters("policyId", popts.GetPolicyId())
	}

	return resp.ToEntityPolicy(), nil
}

func (s *LoadBalancerServiceV2) UpdatePolicy(popts IUpdatePolicyRequest) sdkerror.IError {
	url := updatePolicyUrl(s.VLBClient, popts)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithHeader("User-Agent", popts.ParseUserAgent()).
		WithOkCodes(202).
		WithJsonBody(popts.ToRequestBody()).
		WithJsonError(errResp)

	if _, sdkErr := s.VLBClient.Put(url, req); sdkErr != nil {
		return sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.WithErrorLoadBalancerNotFound(errResp),
			sdkerror.WithErrorListenerNotFound(errResp),
		)
	}

	return nil
}

func (s *LoadBalancerServiceV2) DeletePolicyById(popts IDeletePolicyByIdRequest) sdkerror.IError {
	url := deletePolicyByIdUrl(s.VLBClient, popts)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithHeader("User-Agent", popts.ParseUserAgent()).
		WithOkCodes(202).
		WithJsonError(errResp)

	if _, sdkErr := s.VLBClient.Delete(url, req); sdkErr != nil {
		return sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.WithErrorLoadBalancerNotFound(errResp),
			sdkerror.WithErrorListenerNotFound(errResp),
		)
	}

	return nil
}

func (s *LoadBalancerServiceV2) ReorderPolicies(popts IReorderPoliciesRequest) sdkerror.IError {
	url := reorderPoliciesUrl(s.VLBClient, popts)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithHeader("User-Agent", popts.ParseUserAgent()).
		WithOkCodes(202).
		WithJsonBody(popts.ToRequestBody()).
		WithJsonError(errResp)
	if _, sdkErr := s.VLBClient.Put(url, req); sdkErr != nil {
		return sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.WithErrorLoadBalancerNotFound(errResp),
			sdkerror.WithErrorListenerNotFound(errResp),
		)
	}
	return nil
}

// --------------------------------------------------------

func (s *LoadBalancerServiceV2) ListCertificates(popts IListCertificatesRequest) (*entity.ListCertificates, sdkerror.IError) {
	url := listCertificatesUrl(s.VLBClient)
	resp := new(ListCertificatesResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithHeader("User-Agent", popts.ParseUserAgent()).
		WithOkCodes(200).
		WithJsonResponse(resp).
		WithJsonError(errResp)

	if _, sdkErr := s.VLBClient.Get(url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp)
	}

	return resp.ToEntityListCertificates(), nil
}

func (s *LoadBalancerServiceV2) GetCertificateById(popts IGetCertificateByIdRequest) (*entity.Certificate, sdkerror.IError) {
	url := getCertificateByIdUrl(s.VLBClient, popts)
	resp := new(GetCertificateByIdResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithHeader("User-Agent", popts.ParseUserAgent()).
		WithOkCodes(200).
		WithJsonResponse(resp).
		WithJsonError(errResp)

	if _, sdkErr := s.VLBClient.Get(url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp)
	}

	return resp.ToEntityCertificate(), nil
}

func (s *LoadBalancerServiceV2) CreateCertificate(popts ICreateCertificateRequest) (*entity.Certificate, sdkerror.IError) {
	url := createCertificateUrl(s.VLBClient)
	resp := new(CreateCertificateResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithHeader("User-Agent", popts.ParseUserAgent()).
		WithOkCodes(201).
		WithJsonBody(popts.ToRequestBody()).
		WithJsonResponse(resp).
		WithJsonError(errResp)

	if _, sdkErr := s.VLBClient.Post(url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp)
	}

	return resp.ToEntityCertificate(), nil
}

func (s *LoadBalancerServiceV2) DeleteCertificateById(popts IDeleteCertificateByIdRequest) sdkerror.IError {
	url := deleteCertificateByIdUrl(s.VLBClient, popts)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithHeader("User-Agent", popts.ParseUserAgent()).
		WithOkCodes(204).
		WithJsonError(errResp)

	if _, sdkErr := s.VLBClient.Delete(url, req); sdkErr != nil {
		return sdkerror.SdkErrorHandler(sdkErr, errResp)
	}

	return nil
}
