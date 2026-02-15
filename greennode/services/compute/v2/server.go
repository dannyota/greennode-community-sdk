package v2

import (
	"github.com/dannyota/greennode-community-sdk/v2/greennode/client"
	"github.com/dannyota/greennode-community-sdk/v2/greennode/entity"
	sdkerror "github.com/dannyota/greennode-community-sdk/v2/greennode/sdkerror"
	"github.com/dannyota/greennode-community-sdk/v2/greennode/services/common"
)

func (s *ComputeServiceV2) CreateServer(opts *CreateServerRequest) (*entity.Server, error) {
	url := createServerURL(s.VServerClient)
	resp := new(CreateServerResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithHeader("User-Agent", opts.ParseUserAgent()).
		WithOkCodes(202).
		WithJSONBody(opts).
		WithJSONResponse(resp).
		WithJSONError(errResp)

	if _, sdkErr := s.VServerClient.Post(url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.EcPurchaseIssue,
			sdkerror.EcVServerSubnetNotFound,
			sdkerror.EcVServerImageNotFound,
			sdkerror.EcVServerServerExceedQuota,
			sdkerror.EcVServerServerExceedCpuQuota,
			sdkerror.EcVServerServerFlavorSystemExceedQuota,
			sdkerror.EcVServerVolumeTypeNotFound,
			sdkerror.EcVServerNetworkNotFound,
			sdkerror.EcVServerVolumeExceedQuota,
			sdkerror.EcVServerVolumeSizeExceedGlobalQuota,
			sdkerror.EcVServerSecgroupNotFound,
			sdkerror.EcVServerServerExceedFloatingIpQuota,
			sdkerror.EcVServerServerImageNotSupported,
			sdkerror.EcVServerFlavorNotSupported,
			sdkerror.EcProjectConflict,
			sdkerror.EcVServerCreateBillingPaymentMethodNotAllowed).
			WithParameters(common.StructToMap(opts)).
			WithKVparameters("projectId", s.getProjectID()).
			WithErrorCategories(sdkerror.ErrCatVServer)
	}

	return resp.ToEntityServer(), nil
}

func (s *ComputeServiceV2) GetServerByID(opts *GetServerByIDRequest) (*entity.Server, error) {
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
			sdkerror.EcVServerServerNotFound).
			WithParameters(common.StructToMap(opts)).
			WithKVparameters("projectId", s.getProjectID())
	}

	return resp.ToEntityServer(), nil
}

func (s *ComputeServiceV2) DeleteServerByID(opts *DeleteServerByIDRequest) error {
	url := deleteServerByIDURL(s.VServerClient, opts)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithHeader("User-Agent", opts.ParseUserAgent()).
		WithOkCodes(202).
		WithJSONBody(opts).
		WithJSONError(errResp)

	if _, sdkErr := s.VServerClient.Delete(url, req); sdkErr != nil {
		return sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.EcVServerServerNotFound,
			sdkerror.EcVServerServerDeleteDeletingServer,
			sdkerror.EcVServerServerUpdatingSecgroups,
			sdkerror.EcVServerServerDeleteBillingServer,
			sdkerror.EcVServerServerDeleteCreatingServer,
			sdkerror.EcVServerVolumeInProcess).
			WithKVparameters("projectId", s.getProjectID(),
				"serverId", opts.GetServerID())
	}

	return nil
}

func (s *ComputeServiceV2) UpdateServerSecgroupsByServerID(opts *UpdateServerSecgroupsByServerIDRequest) (*entity.Server, error) {
	url := updateServerSecgroupsByServerIDURL(s.VServerClient, opts)
	resp := new(UpdateServerSecgroupsByServerIDResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithHeader("User-Agent", opts.ParseUserAgent()).
		WithOkCodes(202).
		WithJSONBody(opts).
		WithJSONResponse(resp).
		WithJSONError(errResp)

	if _, sdkErr := s.VServerClient.Put(url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.EcVServerServerNotFound,
			sdkerror.EcVServerServerExpired,
			sdkerror.EcVServerServerUpdatingSecgroups,
			sdkerror.EcVServerSecgroupNotFound).
			WithKVparameters("projectId", s.getProjectID(),
				"serverId", opts.GetServerID(),
				"secgroupIds", opts.GetListSecgroupsIDs())
	}

	return resp.ToEntityServer(), nil
}

func (s *ComputeServiceV2) AttachBlockVolume(opts *AttachBlockVolumeRequest) error {
	url := attachBlockVolumeURL(s.VServerClient, opts)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithOkCodes(202).
		WithJSONBody(map[string]any{}).
		WithJSONError(errResp)

	if _, sdkErr := s.VServerClient.Put(url, req); sdkErr != nil {
		return sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.EcVServerVolumeNotFound,
			sdkerror.EcVServerServerNotFound,
			sdkerror.EcVServerVolumeAvailable,
			sdkerror.EcVServerVolumeInProcess,
			sdkerror.EcVServerVolumeAlreadyAttached,
			sdkerror.EcVServerServerAttachEncryptedVolume,
			sdkerror.EcVServerVolumeAlreadyAttachedThisServer,
			sdkerror.EcVServerServerVolumeAttachQuotaExceeded).
			WithKVparameters("projectId", s.getProjectID(),
				"volumeId", opts.GetBlockVolumeID(),
				"serverId", opts.GetServerID())
	}

	return nil
}

func (s *ComputeServiceV2) DetachBlockVolume(opts *DetachBlockVolumeRequest) error {
	url := detachBlockVolumeURL(s.VServerClient, opts)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithOkCodes(202).
		WithJSONBody(map[string]any{}).
		WithJSONError(errResp)

	if _, sdkErr := s.VServerClient.Put(url, req); sdkErr != nil {
		return sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.EcVServerVolumeNotFound,
			sdkerror.EcVServerVolumeInProcess,
			sdkerror.EcVServerServerNotFound,
			sdkerror.EcVServerVolumeIsMigrating,
			sdkerror.EcVServerVolumeAvailable).
			WithKVparameters("projectId", s.getProjectID(),
				"volumeId", opts.GetBlockVolumeID(),
				"serverId", opts.GetServerID())
	}

	return nil
}

func (s *ComputeServiceV2) AttachFloatingIp(opts *AttachFloatingIpRequest) error {
	url := attachFloatingIpURL(s.VServerClient, opts)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithHeader("User-Agent", opts.ParseUserAgent()).
		WithOkCodes(204).
		WithJSONBody(opts).
		WithJSONError(errResp)

	if _, sdkErr := s.VServerClient.Put(url, req); sdkErr != nil {
		return sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.EcVServerServerNotFound,
			sdkerror.EcVServerServerCanNotAttachFloatingIp,
			sdkerror.EcVServerInternalNetworkInterfaceNotFound).
			WithParameters(common.StructToMap(opts)).
			WithKVparameters("projectId", s.getProjectID())
	}

	return nil
}

func (s *ComputeServiceV2) DetachFloatingIp(opts *DetachFloatingIpRequest) error {
	url := detachFloatingIpURL(s.VServerClient, opts)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithHeader("User-Agent", opts.ParseUserAgent()).
		WithOkCodes(204).
		WithJSONBody(opts).
		WithJSONError(errResp)

	if _, sdkErr := s.VServerClient.Put(url, req); sdkErr != nil {
		return sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.EcVServerWanIpAvailable,
			sdkerror.EcVServerServerNotFound,
			sdkerror.EcVServerWanIDNotFound,
			sdkerror.EcVServerInternalNetworkInterfaceNotFound).
			WithParameters(common.StructToMap(opts)).
			WithKVparameters("projectId", s.getProjectID())
	}

	return nil
}

func (s *ComputeServiceV2) ListServerGroupPolicies(opts *ListServerGroupPoliciesRequest) (*entity.ListServerGroupPolicies, error) {
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

func (s *ComputeServiceV2) DeleteServerGroupByID(opts *DeleteServerGroupByIDRequest) error {
	url := deleteServerGroupByIDURL(s.VServerClient, opts)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithHeader("User-Agent", opts.ParseUserAgent()).
		WithOkCodes(204).
		WithJSONError(errResp)

	if _, sdkErr := s.VServerClient.Delete(url, req); sdkErr != nil {
		return sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.EcVServerServerGroupNotFound,
			sdkerror.EcVServerServerGroupInUse).
			WithParameters(common.StructToMap(opts)).
			WithKVparameters("projectId", s.getProjectID(),
				"serverGroupId", opts.GetServerGroupID())
	}

	return nil
}

func (s *ComputeServiceV2) ListServerGroups(opts *ListServerGroupsRequest) (*entity.ListServerGroups, error) {
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
			WithParameters(common.StructToMap(opts)).
			WithKVparameters("projectId", s.getProjectID())
	}

	return resp.ToEntityListServerGroups(), nil
}

func (s *ComputeServiceV2) CreateServerGroup(opts *CreateServerGroupRequest) (*entity.ServerGroup, error) {
	url := createServerGroupURL(s.VServerClient, opts)
	resp := new(CreateServerGroupResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithHeader("User-Agent", opts.ParseUserAgent()).
		WithOkCodes(201).
		WithJSONBody(opts).
		WithJSONResponse(resp).
		WithJSONError(errResp)

	if _, sdkErr := s.VServerClient.Post(url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.EcVServerServerGroupNameMustBeUnique).
			WithParameters(common.StructToMap(opts)).
			WithKVparameters("projectId", s.getProjectID())
	}

	return resp.ToEntityServerGroup(), nil
}
