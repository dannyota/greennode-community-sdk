package v2

import (
	"context"

	"github.com/dannyota/greennode-community-sdk/v2/greennode/client"
	"github.com/dannyota/greennode-community-sdk/v2/greennode/entity"
	sdkerror "github.com/dannyota/greennode-community-sdk/v2/greennode/sdkerror"
	"github.com/dannyota/greennode-community-sdk/v2/greennode/services/common"
)

type LoadBalancerServiceV2 struct {
	Client       *client.ServiceClient
	ServerClient *client.ServiceClient
}

func (s *LoadBalancerServiceV2) getProjectID() string {
	return s.Client.ProjectID()
}

const (
	defaultPageListLoadBalancer = 1
	defaultSizeListLoadBalancer = 10
)

func (s *LoadBalancerServiceV2) CreateLoadBalancer(ctx context.Context, opts *CreateLoadBalancerRequest) (*entity.LoadBalancer, error) {
	url := createLoadBalancerURL(s.Client)
	resp := new(CreateLoadBalancerResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	opts.normalizeForAPI()
	req := client.NewRequest().
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

func (s *LoadBalancerServiceV2) ResizeLoadBalancer(ctx context.Context, opts *ResizeLoadBalancerRequest) (*entity.LoadBalancer, error) {
	url := resizeLoadBalancerURL(s.Client, opts)
	resp := new(ResizeLoadBalancerResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithOkCodes(202).
		WithJSONBody(opts).
		WithJSONResponse(resp).
		WithJSONError(errResp)

	if _, sdkErr := s.Client.Put(ctx, url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.EcVLBLoadBalancerNotFound,
			sdkerror.EcVLBLoadBalancerNotReady)
	}

	return resp.ToEntityLoadBalancer(), nil
}

func (s *LoadBalancerServiceV2) ListLoadBalancerPackages(ctx context.Context, opts *ListLoadBalancerPackagesRequest) (*entity.ListLoadBalancerPackages, error) {
	url := listLoadBalancerPackagesURL(s.Client, opts)
	resp := new(ListLoadBalancerPackagesResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithOkCodes(200).
		WithJSONResponse(resp).
		WithJSONError(errResp)

	if _, sdkErr := s.Client.Get(ctx, url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp)
	}

	return resp.ToEntityListLoadBalancerPackages(), nil
}

func (s *LoadBalancerServiceV2) GetLoadBalancerByID(ctx context.Context, opts *GetLoadBalancerByIDRequest) (*entity.LoadBalancer, error) {
	url := getLoadBalancerByIDURL(s.Client, opts)
	resp := new(GetLoadBalancerByIDResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithOkCodes(200).
		WithJSONResponse(resp).
		WithJSONError(errResp)

	if _, sdkErr := s.Client.Get(ctx, url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.EcVLBLoadBalancerNotFound).
			WithKVparameters("loadBalancerId", opts.LoadBalancerID).
			AppendCategories(sdkerror.ErrCatProductVlb)
	}

	return resp.ToEntityLoadBalancer(), nil
}

func (s *LoadBalancerServiceV2) ListLoadBalancers(ctx context.Context, opts *ListLoadBalancersRequest) (*entity.ListLoadBalancers, error) {
	url := listLoadBalancersURL(s.Client, opts)
	resp := new(ListLoadBalancersResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithOkCodes(200).
		WithJSONResponse(resp).
		WithJSONError(errResp)

	if _, sdkErr := s.Client.Get(ctx, url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp).
			AppendCategories(sdkerror.ErrCatProductVlb)
	}

	return resp.ToEntityListLoadBalancers(), nil
}

func (s *LoadBalancerServiceV2) GetPoolHealthMonitorByID(ctx context.Context, opts *GetPoolHealthMonitorByIDRequest) (*entity.HealthMonitor, error) {
	url := getPoolHealthMonitorByIDURL(s.Client, opts)
	resp := new(GetPoolHealthMonitorByIDResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithOkCodes(200).
		WithJSONResponse(resp).
		WithJSONError(errResp)

	if _, sdkErr := s.Client.Get(ctx, url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.EcVLBLoadBalancerNotFound,
			sdkerror.EcVLBPoolNotFound)
	}

	return resp.ToEntityHealthMonitor(), nil
}

func (s *LoadBalancerServiceV2) CreatePool(ctx context.Context, opts *CreatePoolRequest) (*entity.Pool, error) {
	url := createPoolURL(s.Client, opts)
	resp := new(CreatePoolResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	opts.normalizeForAPI()
	req := client.NewRequest().
		WithOkCodes(202).
		WithJSONBody(opts).
		WithJSONResponse(resp).
		WithJSONError(errResp)

	if _, sdkErr := s.Client.Post(ctx, url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.EcVLBLoadBalancerNotFound,
			sdkerror.EcVLBLoadBalancerNotReady,
			sdkerror.EcVLBLoadBalancerDuplicatePoolName).
			WithParameters(common.StructToMap(opts)).
			AppendCategories(sdkerror.ErrCatProductVlb)
	}

	return resp.ToEntityPool(), nil
}

func (s *LoadBalancerServiceV2) UpdatePool(ctx context.Context, opts *UpdatePoolRequest) error {
	url := updatePoolURL(s.Client, opts)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	opts.normalizeForAPI()
	req := client.NewRequest().
		WithOkCodes(202).
		WithJSONBody(opts).
		WithJSONError(errResp)

	if _, sdkErr := s.Client.Put(ctx, url, req); sdkErr != nil {
		return sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.EcVLBLoadBalancerNotFound,
			sdkerror.EcVLBLoadBalancerNotReady,
			sdkerror.EcVLBListenerNotFound).
			AppendCategories(sdkerror.ErrCatProductVlb)
	}
	return nil
}

func (s *LoadBalancerServiceV2) CreateListener(ctx context.Context, opts *CreateListenerRequest) (*entity.Listener, error) {
	url := createListenerURL(s.Client, opts)
	resp := new(CreateListenerResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	opts.normalizeForAPI()
	req := client.NewRequest().
		WithOkCodes(202).
		WithJSONBody(opts).
		WithJSONResponse(resp).
		WithJSONError(errResp)

	if _, sdkErr := s.Client.Post(ctx, url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.EcVLBLoadBalancerNotFound,
			sdkerror.EcVLBLoadBalancerNotReady,
			sdkerror.EcVLBListenerDuplicateName,
			sdkerror.EcVLBPoolNotFound,
			sdkerror.EcVLBListenerDuplicateProtocolOrPort).
			WithParameters(common.StructToMap(opts)).
			AppendCategories(sdkerror.ErrCatProductVlb)
	}

	return resp.ToEntityListener(), nil
}

func (s *LoadBalancerServiceV2) UpdateListener(ctx context.Context, opts *UpdateListenerRequest) error {
	url := updateListenerURL(s.Client, opts)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithOkCodes(202).
		WithJSONBody(opts).
		WithJSONError(errResp)

	if _, sdkErr := s.Client.Put(ctx, url, req); sdkErr != nil {
		return sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.EcVLBLoadBalancerNotFound,
			sdkerror.EcVLBLoadBalancerNotReady,
			sdkerror.EcVLBListenerNotFound).
			AppendCategories(sdkerror.ErrCatProductVlb)
	}

	return nil
}

func (s *LoadBalancerServiceV2) ListListenersByLoadBalancerID(ctx context.Context, opts *ListListenersByLoadBalancerIDRequest) (*entity.ListListeners, error) {
	url := listListenersByLoadBalancerIDURL(s.Client, opts)
	resp := new(ListListenersByLoadBalancerIDResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithOkCodes(200).
		WithJSONResponse(resp).
		WithJSONError(errResp)

	if _, sdkErr := s.Client.Get(ctx, url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.EcVLBLoadBalancerNotFound).
			WithKVparameters("loadBalancerId", opts.LoadBalancerID).
			AppendCategories(sdkerror.ErrCatProductVlb)
	}

	return resp.ToEntityListListeners(), nil
}

func (s *LoadBalancerServiceV2) ListPoolsByLoadBalancerID(ctx context.Context, opts *ListPoolsByLoadBalancerIDRequest) (*entity.ListPools, error) {
	url := listPoolsByLoadBalancerIDURL(s.Client, opts)
	resp := new(ListPoolsByLoadBalancerIDResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithOkCodes(200).
		WithJSONResponse(resp).
		WithJSONError(errResp)

	if _, sdkErr := s.Client.Get(ctx, url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.EcVLBLoadBalancerNotFound).
			WithKVparameters("loadBalancerId", opts.LoadBalancerID).
			AppendCategories(sdkerror.ErrCatProductVlb)
	}

	return resp.ToEntityListPools(), nil
}

func (s *LoadBalancerServiceV2) UpdatePoolMembers(ctx context.Context, opts *UpdatePoolMembersRequest) error {
	url := updatePoolMembersURL(s.Client, opts)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithOkCodes(202).
		WithJSONBody(opts).
		WithJSONError(errResp)

	if _, sdkErr := s.Client.Put(ctx, url, req); sdkErr != nil {
		return sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.EcVLBLoadBalancerNotFound,
			sdkerror.EcVLBLoadBalancerNotReady,
			sdkerror.EcVLBPoolNotFound,
			sdkerror.EcVLBMemberMustIdentical).
			AppendCategories(sdkerror.ErrCatProductVlb)
	}

	return nil
}

func (s *LoadBalancerServiceV2) ListPoolMembers(ctx context.Context, opts *ListPoolMembersRequest) (*entity.ListMembers, error) {
	url := listPoolMembersURL(s.Client, opts)
	resp := new(ListPoolMembersResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithOkCodes(200).
		WithJSONResponse(resp).
		WithJSONError(errResp)

	if _, sdkErr := s.Client.Get(ctx, url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.EcVLBLoadBalancerNotFound,
			sdkerror.EcVLBPoolNotFound).
			WithKVparameters("loadBalancerId", opts.LoadBalancerID, "poolId", opts.PoolID).
			AppendCategories(sdkerror.ErrCatProductVlb)
	}

	return resp.ToEntityListMembers(), nil
}

func (s *LoadBalancerServiceV2) DeletePoolByID(ctx context.Context, opts *DeletePoolByIDRequest) error {
	url := deletePoolByIDURL(s.Client, opts)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithOkCodes(202).
		WithJSONError(errResp)

	if _, sdkErr := s.Client.Delete(ctx, url, req); sdkErr != nil {
		return sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.EcVLBPoolInUse,
			sdkerror.EcVLBLoadBalancerNotReady,
			sdkerror.EcVLBPoolNotFound).
			AppendCategories(sdkerror.ErrCatProductVlb)
	}

	return nil
}

func (s *LoadBalancerServiceV2) DeleteListenerByID(ctx context.Context, opts *DeleteListenerByIDRequest) error {
	url := deleteListenerByIDURL(s.Client, opts)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithOkCodes(202).
		WithJSONError(errResp)

	if _, sdkErr := s.Client.Delete(ctx, url, req); sdkErr != nil {
		return sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.EcVLBLoadBalancerNotFound,
			sdkerror.EcVLBListenerNotFound,
			sdkerror.EcVLBLoadBalancerNotReady).
			AppendCategories(sdkerror.ErrCatProductVlb)
	}

	return nil
}

func (s *LoadBalancerServiceV2) DeleteLoadBalancerByID(ctx context.Context, opts *DeleteLoadBalancerByIDRequest) error {
	url := deleteLoadBalancerByIDURL(s.Client, opts)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithOkCodes(202).
		WithJSONError(errResp)

	if _, sdkErr := s.Client.Delete(ctx, url, req); sdkErr != nil {
		return sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.EcVLBLoadBalancerNotFound,
			sdkerror.EcVLBLoadBalancerNotReady,
			sdkerror.EcVLBLoadBalancerIsCreating,
			sdkerror.EcVLBLoadBalancerIsDeleting).
			WithKVparameters(
				"loadBalancerId", opts.LoadBalancerID,
				"projectId", s.getProjectID()).
			AppendCategories(sdkerror.ErrCatProductVlb)
	}

	return nil
}

func (s *LoadBalancerServiceV2) GetPoolByID(ctx context.Context, opts *GetPoolByIDRequest) (*entity.Pool, error) {
	url := getPoolByIDURL(s.Client, opts)
	resp := new(GetPoolByIDResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithOkCodes(200).
		WithJSONResponse(resp).
		WithJSONError(errResp)

	if _, sdkErr := s.Client.Get(ctx, url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.EcVLBLoadBalancerNotFound,
			sdkerror.EcVLBPoolNotFound).
			WithKVparameters(
				"loadBalancerId", opts.LoadBalancerID,
				"poolId", opts.PoolID).
			AppendCategories(sdkerror.ErrCatProductVlb)
	}

	return resp.ToEntityPool(), nil
}

func (s *LoadBalancerServiceV2) GetListenerByID(ctx context.Context, opts *GetListenerByIDRequest) (*entity.Listener, error) {
	url := getListenerByIDURL(s.Client, opts)
	resp := new(GetListenerByIDResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithOkCodes(200).
		WithJSONResponse(resp).
		WithJSONError(errResp)

	if _, sdkErr := s.Client.Get(ctx, url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.EcVLBLoadBalancerNotFound,
			sdkerror.EcVLBListenerNotFound).
			WithKVparameters(
				"loadBalancerId", opts.LoadBalancerID,
				"listenerId", opts.ListenerID).
			AppendCategories(sdkerror.ErrCatProductVlb)
	}

	return resp.ToEntityListener(), nil
}

func (s *LoadBalancerServiceV2) ResizeLoadBalancerByID(ctx context.Context, opts *ResizeLoadBalancerByIDRequest) error {
	url := resizeLoadBalancerByIDURL(s.Client, opts)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithOkCodes(202).
		WithJSONBody(opts).
		WithJSONError(errResp)

	if _, sdkErr := s.Client.Put(ctx, url, req); sdkErr != nil {
		return sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.EcVLBLoadBalancerPackageNotFound,
			sdkerror.EcVLBLoadBalancerNotFound,
			sdkerror.EcVLBLoadBalancerNotReady,
			sdkerror.EcVLBLoadBalancerResizeSamePackage).
			WithParameters(common.StructToMap(opts)).
			AppendCategories(sdkerror.ErrCatProductVlb)
	}

	return nil
}

func (s *LoadBalancerServiceV2) ScaleLoadBalancer(ctx context.Context, opts *ScaleLoadBalancerRequest) (*entity.LoadBalancer, error) {
	url := scaleLoadBalancerURL(s.Client, opts)
	resp := new(ScaleLoadBalancerResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithOkCodes(202).
		WithJSONBody(opts).
		WithJSONResponse(resp).
		WithJSONError(errResp)

	if _, sdkErr := s.Client.Put(ctx, url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.EcVLBLoadBalancerNotFound,
			sdkerror.EcVLBLoadBalancerNotReady).
			WithParameters(common.StructToMap(opts)).
			AppendCategories(sdkerror.ErrCatProductVlb)
	}

	return resp.ToEntityLoadBalancer(), nil
}

// policy

func (s *LoadBalancerServiceV2) ListPolicies(ctx context.Context, opts *ListPoliciesRequest) (*entity.ListPolicies, error) {
	url := listPoliciesURL(s.Client, opts)
	resp := new(ListPoliciesResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithOkCodes(200).
		WithJSONResponse(resp).
		WithJSONError(errResp)

	if _, sdkErr := s.Client.Get(ctx, url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp)
	}

	return resp.ToEntityListPolicies(), nil
}

func (s *LoadBalancerServiceV2) CreatePolicy(ctx context.Context, opts *CreatePolicyRequest) (*entity.Policy, error) {
	url := createPolicyURL(s.Client, opts)
	resp := new(CreatePolicyResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithOkCodes(202).
		WithJSONBody(opts.toRequestBody()).
		WithJSONResponse(resp).
		WithJSONError(errResp)

	if _, sdkErr := s.Client.Post(ctx, url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.EcVLBLoadBalancerNotFound,
			sdkerror.EcVLBListenerNotFound,
		).WithParameters(common.StructToMap(opts))
	}

	return resp.ToEntityPolicy(), nil
}

func (s *LoadBalancerServiceV2) GetPolicyByID(ctx context.Context, opts *GetPolicyByIDRequest) (*entity.Policy, error) {
	url := getPolicyByIDURL(s.Client, opts)
	resp := new(GetPolicyResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithOkCodes(200).
		WithJSONResponse(resp).
		WithJSONError(errResp)

	if _, sdkErr := s.Client.Get(ctx, url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.EcVLBLoadBalancerNotFound,
			sdkerror.EcVLBListenerNotFound,
		).WithKVparameters("policyId", opts.PolicyID)
	}

	return resp.ToEntityPolicy(), nil
}

func (s *LoadBalancerServiceV2) UpdatePolicy(ctx context.Context, opts *UpdatePolicyRequest) error {
	url := updatePolicyURL(s.Client, opts)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithOkCodes(202).
		WithJSONBody(opts.toRequestBody()).
		WithJSONError(errResp)

	if _, sdkErr := s.Client.Put(ctx, url, req); sdkErr != nil {
		return sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.EcVLBLoadBalancerNotFound,
			sdkerror.EcVLBListenerNotFound,
		)
	}

	return nil
}

func (s *LoadBalancerServiceV2) DeletePolicyByID(ctx context.Context, opts *DeletePolicyByIDRequest) error {
	url := deletePolicyByIDURL(s.Client, opts)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithOkCodes(202).
		WithJSONError(errResp)

	if _, sdkErr := s.Client.Delete(ctx, url, req); sdkErr != nil {
		return sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.EcVLBLoadBalancerNotFound,
			sdkerror.EcVLBListenerNotFound,
		)
	}

	return nil
}

func (s *LoadBalancerServiceV2) ReorderPolicies(ctx context.Context, opts *ReorderPoliciesRequest) error {
	url := reorderPoliciesURL(s.Client, opts)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithOkCodes(202).
		WithJSONBody(opts.toRequestBody()).
		WithJSONError(errResp)
	if _, sdkErr := s.Client.Put(ctx, url, req); sdkErr != nil {
		return sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.EcVLBLoadBalancerNotFound,
			sdkerror.EcVLBListenerNotFound,
		)
	}
	return nil
}


func (s *LoadBalancerServiceV2) ListCertificates(ctx context.Context, opts *ListCertificatesRequest) (*entity.ListCertificates, error) {
	url := listCertificatesURL(s.Client)
	resp := new(ListCertificatesResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithOkCodes(200).
		WithJSONResponse(resp).
		WithJSONError(errResp)

	if _, sdkErr := s.Client.Get(ctx, url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp)
	}

	return resp.ToEntityListCertificates(), nil
}

func (s *LoadBalancerServiceV2) GetCertificateByID(ctx context.Context, opts *GetCertificateByIDRequest) (*entity.Certificate, error) {
	url := getCertificateByIDURL(s.Client, opts)
	resp := new(GetCertificateByIDResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithOkCodes(200).
		WithJSONResponse(resp).
		WithJSONError(errResp)

	if _, sdkErr := s.Client.Get(ctx, url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp)
	}

	return resp.ToEntityCertificate(), nil
}

func (s *LoadBalancerServiceV2) CreateCertificate(ctx context.Context, opts *CreateCertificateRequest) (*entity.Certificate, error) {
	url := createCertificateURL(s.Client)
	resp := new(CreateCertificateResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithOkCodes(201).
		WithJSONBody(opts).
		WithJSONResponse(resp).
		WithJSONError(errResp)

	if _, sdkErr := s.Client.Post(ctx, url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp)
	}

	return resp.ToEntityCertificate(), nil
}

func (s *LoadBalancerServiceV2) DeleteCertificateByID(ctx context.Context, opts *DeleteCertificateByIDRequest) error {
	url := deleteCertificateByIDURL(s.Client, opts)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithOkCodes(204).
		WithJSONError(errResp)

	if _, sdkErr := s.Client.Delete(ctx, url, req); sdkErr != nil {
		return sdkerror.SdkErrorHandler(sdkErr, errResp)
	}

	return nil
}
