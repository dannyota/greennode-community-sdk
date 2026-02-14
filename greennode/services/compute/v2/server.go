package v2

import (
	"github.com/dannyota/greennode-community-sdk/v2/greennode/client"
	"github.com/dannyota/greennode-community-sdk/v2/greennode/entity"
	sdkerror "github.com/dannyota/greennode-community-sdk/v2/greennode/sdkerror"
)

func (s *ComputeServiceV2) CreateServer(opts ICreateServerRequest) (*entity.Server, sdkerror.Error) {
	url := createServerURL(s.VServerClient)
	resp := new(CreateServerResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithHeader("User-Agent", opts.ParseUserAgent()).
		WithOkCodes(202).
		WithJSONBody(opts.ToRequestBody()).
		WithJSONResponse(resp).
		WithJSONError(errResp)

	if _, sdkErr := s.VServerClient.Post(url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.WithErrorPurchaseIssue(errResp),
			sdkerror.WithErrorSubnetNotFound(errResp),
			sdkerror.WithErrorImageNotFound(errResp),
			sdkerror.WithErrorServerExceedQuota(errResp),
			sdkerror.WithErrorServerExceedCpuQuota(errResp),
			sdkerror.WithErrorServerFlavorSystemExceedQuota(errResp),
			sdkerror.WithErrorVolumeTypeNotFound(errResp),
			sdkerror.WithErrorNetworkNotFound(errResp),
			sdkerror.WithErrorVolumeExceedQuota(errResp),
			sdkerror.WithErrorVolumeSizeExceedGlobalQuota(errResp),
			sdkerror.WithErrorSecgroupNotFound(errResp),
			sdkerror.WithErrorServerExceedFloatingIpQuota(errResp),
			sdkerror.WithErrorServerImageNotSupported(errResp),
			sdkerror.WithErrorServerFlavorNotSupported(errResp),
			sdkerror.WithErrorProjectConflict(errResp),
			sdkerror.WithErrorServerCreateBillingPaymentMethodNotAllowed(errResp)).
			WithParameters(opts.ToMap()).
			WithKVparameters("projectId", s.getProjectID()).
			WithErrorCategories(sdkerror.ErrCatVServer)
	}

	return resp.ToEntityServer(), nil
}

func (s *ComputeServiceV2) GetServerByID(opts IGetServerByIDRequest) (*entity.Server, sdkerror.Error) {
	url := getServerByIDURL(s.VServerClient, opts)
	resp := new(GetServerByIDResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithHeader("User-Agent", opts.ParseUserAgent()).
		WithOkCodes(200).
		WithJSONResponse(resp).
		WithJSONError(errResp)

	if _, sdkErr := s.VServerClient.Get(url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.WithErrorServerNotFound(errResp)).
			WithParameters(opts.ToMap()).
			WithKVparameters("projectId", s.getProjectID())
	}

	return resp.ToEntityServer(), nil
}

func (s *ComputeServiceV2) DeleteServerByID(opts IDeleteServerByIDRequest) sdkerror.Error {
	url := deleteServerByIDURL(s.VServerClient, opts)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithHeader("User-Agent", opts.ParseUserAgent()).
		WithOkCodes(202).
		WithJSONBody(opts.ToRequestBody()).
		WithJSONError(errResp)

	if _, sdkErr := s.VServerClient.Delete(url, req); sdkErr != nil {
		return sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.WithErrorServerNotFound(errResp),
			sdkerror.WithErrorServerDeleteDeletingServer(errResp),
			sdkerror.WithErrorServerUpdatingSecgroups(errResp),
			sdkerror.WithErrorServerDeleteBillingServer(errResp),
			sdkerror.WithErrorServerDeleteCreatingServer(errResp),
			sdkerror.WithErrorVolumeInProcess(errResp)).
			WithKVparameters("projectId", s.getProjectID(),
				"serverId", opts.GetServerID())
	}

	return nil
}

func (s *ComputeServiceV2) UpdateServerSecgroupsByServerID(opts IUpdateServerSecgroupsByServerIDRequest) (*entity.Server, sdkerror.Error) {
	url := updateServerSecgroupsByServerIDURL(s.VServerClient, opts)
	resp := new(UpdateServerSecgroupsByServerIDResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithHeader("User-Agent", opts.ParseUserAgent()).
		WithOkCodes(202).
		WithJSONBody(opts.ToRequestBody()).
		WithJSONResponse(resp).
		WithJSONError(errResp)

	if _, sdkErr := s.VServerClient.Put(url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.WithErrorServerNotFound(errResp),
			sdkerror.WithErrorServerExpired(errResp),
			sdkerror.WithErrorServerUpdatingSecgroups(errResp),
			sdkerror.WithErrorSecgroupNotFound(errResp)).
			WithKVparameters("projectId", s.getProjectID(),
				"serverId", opts.GetServerID(),
				"secgroupIds", opts.GetListSecgroupsIDs())
	}

	return resp.ToEntityServer(), nil
}

func (s *ComputeServiceV2) AttachBlockVolume(opts IAttachBlockVolumeRequest) sdkerror.Error {
	url := attachBlockVolumeURL(s.VServerClient, opts)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithOkCodes(202).
		WithJSONBody(map[string]interface{}{}).
		WithJSONError(errResp)

	if _, sdkErr := s.VServerClient.Put(url, req); sdkErr != nil {
		return sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.WithErrorVolumeNotFound(errResp),
			sdkerror.WithErrorServerNotFound(errResp),
			sdkerror.WithErrorVolumeAvailable(errResp),
			sdkerror.WithErrorVolumeInProcess(errResp),
			sdkerror.WithErrorVolumeAlreadyAttached(errResp),
			sdkerror.WithErrorServerAttachEncryptedVolume(errResp),
			sdkerror.WithErrorVolumeAlreadyAttachedThisServer(errResp),
			sdkerror.WithErrorServerAttachVolumeQuotaExceeded(errResp)).
			WithKVparameters("projectId", s.getProjectID(),
				"volumeId", opts.GetBlockVolumeID(),
				"serverId", opts.GetServerID())
	}

	return nil
}

func (s *ComputeServiceV2) DetachBlockVolume(opts IDetachBlockVolumeRequest) sdkerror.Error {
	url := detachBlockVolumeURL(s.VServerClient, opts)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithOkCodes(202).
		WithJSONBody(map[string]interface{}{}).
		WithJSONError(errResp)

	if _, sdkErr := s.VServerClient.Put(url, req); sdkErr != nil {
		return sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.WithErrorVolumeNotFound(errResp),
			sdkerror.WithErrorVolumeInProcess(errResp),
			sdkerror.WithErrorServerNotFound(errResp),
			sdkerror.WithErrorVolumeIsMigrating(errResp),
			sdkerror.WithErrorVolumeAvailable(errResp)).
			WithKVparameters("projectId", s.getProjectID(),
				"volumeId", opts.GetBlockVolumeID(),
				"serverId", opts.GetServerID())
	}

	return nil
}

func (s *ComputeServiceV2) AttachFloatingIp(opts IAttachFloatingIpRequest) sdkerror.Error {
	url := attachFloatingIpURL(s.VServerClient, opts)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithHeader("User-Agent", opts.ParseUserAgent()).
		WithOkCodes(204).
		WithJSONBody(opts.ToRequestBody()).
		WithJSONError(errResp)

	if _, sdkErr := s.VServerClient.Put(url, req); sdkErr != nil {
		return sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.WithErrorServerNotFound(errResp),
			sdkerror.WithErrorServerCanNotAttachFloatingIp(errResp),
			sdkerror.WithErrorInternalNetworkInterfaceNotFound(errResp)).
			WithParameters(opts.ToMap()).
			WithKVparameters("projectId", s.getProjectID())
	}

	return nil
}

func (s *ComputeServiceV2) DetachFloatingIp(opts IDetachFloatingIpRequest) sdkerror.Error {
	url := detachFloatingIpURL(s.VServerClient, opts)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithHeader("User-Agent", opts.ParseUserAgent()).
		WithOkCodes(204).
		WithJSONBody(opts.ToRequestBody()).
		WithJSONError(errResp)

	if _, sdkErr := s.VServerClient.Put(url, req); sdkErr != nil {
		return sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.WithErrorWanIpAvailable(errResp),
			sdkerror.WithErrorServerNotFound(errResp),
			sdkerror.WithErrorWanIDNotFound(errResp),
			sdkerror.WithErrorInternalNetworkInterfaceNotFound(errResp)).
			WithParameters(opts.ToMap()).
			WithKVparameters("projectId", s.getProjectID())
	}

	return nil
}

func (s *ComputeServiceV2) ListServerGroupPolicies(opts IListServerGroupPoliciesRequest) (*entity.ListServerGroupPolicies, sdkerror.Error) {
	url := listServerGroupPoliciesURL(s.VServerClient)
	resp := new(ListServerGroupPoliciesResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithOkCodes(200).
		WithJSONResponse(resp).
		WithJSONError(errResp)

	if opts != nil {
		req = req.WithHeader("User-Agent", opts.ParseUserAgent())
	}

	if _, sdkErr := s.VServerClient.Get(url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp).
			WithKVparameters("projectId", s.getProjectID())
	}

	return resp.ToEntityListServerGroupPolicies(), nil
}

func (s *ComputeServiceV2) DeleteServerGroupByID(opts IDeleteServerGroupByIDRequest) sdkerror.Error {
	url := deleteServerGroupByIDURL(s.VServerClient, opts)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithHeader("User-Agent", opts.ParseUserAgent()).
		WithOkCodes(204).
		WithJSONError(errResp)

	if _, sdkErr := s.VServerClient.Delete(url, req); sdkErr != nil {
		return sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.WithErrorServerGroupNotFound(errResp),
			sdkerror.WithErrorServerGroupInUse(errResp)).
			WithParameters(opts.ToMap()).
			WithKVparameters("projectId", s.getProjectID(),
				"serverGroupId", opts.GetServerGroupID())
	}

	return nil
}

func (s *ComputeServiceV2) ListServerGroups(opts IListServerGroupsRequest) (*entity.ListServerGroups, sdkerror.Error) {
	url := listServerGroupsURL(s.VServerClient, opts)
	resp := new(ListServerGroupsResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithHeader("User-Agent", opts.ParseUserAgent()).
		WithOkCodes(200).
		WithJSONResponse(resp).
		WithJSONError(errResp)

	if _, sdkErr := s.VServerClient.Get(url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp).
			WithParameters(opts.ToMap()).
			WithKVparameters("projectId", s.getProjectID())
	}

	return resp.ToEntityListServerGroups(), nil
}

func (s *ComputeServiceV2) CreateServerGroup(opts ICreateServerGroupRequest) (*entity.ServerGroup, sdkerror.Error) {
	url := createServerGroupURL(s.VServerClient, opts)
	resp := new(CreateServerGroupResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithHeader("User-Agent", opts.ParseUserAgent()).
		WithOkCodes(201).
		WithJSONBody(opts.ToRequestBody()).
		WithJSONResponse(resp).
		WithJSONError(errResp)

	if _, sdkErr := s.VServerClient.Post(url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.WithErrorServerGroupNameMustBeUnique(errResp)).
			WithParameters(opts.ToMap()).
			WithKVparameters("projectId", s.getProjectID())
	}

	return resp.ToEntityServerGroup(), nil
}
