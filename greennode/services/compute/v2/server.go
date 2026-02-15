package v2

import (
	"context"

	"github.com/dannyota/greennode-community-sdk/v2/greennode/client"
	"github.com/dannyota/greennode-community-sdk/v2/greennode/entity"
	sdkerror "github.com/dannyota/greennode-community-sdk/v2/greennode/sdkerror"
	"github.com/dannyota/greennode-community-sdk/v2/greennode/services/common"
)

func (s *ComputeServiceV2) CreateServer(ctx context.Context, opts *CreateServerRequest) (*entity.Server, error) {
	url := createServerURL(s.VServerClient)
	resp := new(CreateServerResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithOkCodes(202).
		WithJSONBody(opts).
		WithJSONResponse(resp).
		WithJSONError(errResp)

	if _, sdkErr := s.VServerClient.Post(ctx, url, req); sdkErr != nil {
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

func (s *ComputeServiceV2) GetServerByID(ctx context.Context, opts *GetServerByIDRequest) (*entity.Server, error) {
	url := getServerByIDURL(s.VServerClient, opts)
	resp := new(GetServerByIDResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithOkCodes(200).
		WithJSONResponse(resp).
		WithJSONError(errResp)

	if _, sdkErr := s.VServerClient.Get(ctx, url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.EcVServerServerNotFound).
			WithParameters(common.StructToMap(opts)).
			WithKVparameters("projectId", s.getProjectID())
	}

	return resp.ToEntityServer(), nil
}

func (s *ComputeServiceV2) DeleteServerByID(ctx context.Context, opts *DeleteServerByIDRequest) error {
	url := deleteServerByIDURL(s.VServerClient, opts)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithOkCodes(202).
		WithJSONBody(opts).
		WithJSONError(errResp)

	if _, sdkErr := s.VServerClient.Delete(ctx, url, req); sdkErr != nil {
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

func (s *ComputeServiceV2) UpdateServerSecgroupsByServerID(ctx context.Context, opts *UpdateServerSecgroupsByServerIDRequest) (*entity.Server, error) {
	url := updateServerSecgroupsByServerIDURL(s.VServerClient, opts)
	resp := new(UpdateServerSecgroupsByServerIDResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithOkCodes(202).
		WithJSONBody(opts).
		WithJSONResponse(resp).
		WithJSONError(errResp)

	if _, sdkErr := s.VServerClient.Put(ctx, url, req); sdkErr != nil {
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

func (s *ComputeServiceV2) AttachBlockVolume(ctx context.Context, opts *AttachBlockVolumeRequest) error {
	url := attachBlockVolumeURL(s.VServerClient, opts)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithOkCodes(202).
		WithJSONBody(map[string]any{}).
		WithJSONError(errResp)

	if _, sdkErr := s.VServerClient.Put(ctx, url, req); sdkErr != nil {
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

func (s *ComputeServiceV2) DetachBlockVolume(ctx context.Context, opts *DetachBlockVolumeRequest) error {
	url := detachBlockVolumeURL(s.VServerClient, opts)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithOkCodes(202).
		WithJSONBody(map[string]any{}).
		WithJSONError(errResp)

	if _, sdkErr := s.VServerClient.Put(ctx, url, req); sdkErr != nil {
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

func (s *ComputeServiceV2) AttachFloatingIp(ctx context.Context, opts *AttachFloatingIpRequest) error {
	url := attachFloatingIpURL(s.VServerClient, opts)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithOkCodes(204).
		WithJSONBody(opts).
		WithJSONError(errResp)

	if _, sdkErr := s.VServerClient.Put(ctx, url, req); sdkErr != nil {
		return sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.EcVServerServerNotFound,
			sdkerror.EcVServerServerCanNotAttachFloatingIp,
			sdkerror.EcVServerInternalNetworkInterfaceNotFound).
			WithParameters(common.StructToMap(opts)).
			WithKVparameters("projectId", s.getProjectID())
	}

	return nil
}

func (s *ComputeServiceV2) DetachFloatingIp(ctx context.Context, opts *DetachFloatingIpRequest) error {
	url := detachFloatingIpURL(s.VServerClient, opts)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithOkCodes(204).
		WithJSONBody(opts).
		WithJSONError(errResp)

	if _, sdkErr := s.VServerClient.Put(ctx, url, req); sdkErr != nil {
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

func (s *ComputeServiceV2) ListServerGroupPolicies(ctx context.Context, opts *ListServerGroupPoliciesRequest) (*entity.ListServerGroupPolicies, error) {
	url := listServerGroupPoliciesURL(s.VServerClient)
	resp := new(ListServerGroupPoliciesResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithOkCodes(200).
		WithJSONResponse(resp).
		WithJSONError(errResp)

	if _, sdkErr := s.VServerClient.Get(ctx, url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp).
			WithKVparameters("projectId", s.getProjectID())
	}

	return resp.ToEntityListServerGroupPolicies(), nil
}

func (s *ComputeServiceV2) DeleteServerGroupByID(ctx context.Context, opts *DeleteServerGroupByIDRequest) error {
	url := deleteServerGroupByIDURL(s.VServerClient, opts)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithOkCodes(204).
		WithJSONError(errResp)

	if _, sdkErr := s.VServerClient.Delete(ctx, url, req); sdkErr != nil {
		return sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.EcVServerServerGroupNotFound,
			sdkerror.EcVServerServerGroupInUse).
			WithParameters(common.StructToMap(opts)).
			WithKVparameters("projectId", s.getProjectID(),
				"serverGroupId", opts.GetServerGroupID())
	}

	return nil
}

func (s *ComputeServiceV2) ListServerGroups(ctx context.Context, opts *ListServerGroupsRequest) (*entity.ListServerGroups, error) {
	url := listServerGroupsURL(s.VServerClient, opts)
	resp := new(ListServerGroupsResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithOkCodes(200).
		WithJSONResponse(resp).
		WithJSONError(errResp)

	if _, sdkErr := s.VServerClient.Get(ctx, url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp).
			WithParameters(common.StructToMap(opts)).
			WithKVparameters("projectId", s.getProjectID())
	}

	return resp.ToEntityListServerGroups(), nil
}

func (s *ComputeServiceV2) CreateServerGroup(ctx context.Context, opts *CreateServerGroupRequest) (*entity.ServerGroup, error) {
	url := createServerGroupURL(s.VServerClient, opts)
	resp := new(CreateServerGroupResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithOkCodes(201).
		WithJSONBody(opts).
		WithJSONResponse(resp).
		WithJSONError(errResp)

	if _, sdkErr := s.VServerClient.Post(ctx, url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.EcVServerServerGroupNameMustBeUnique).
			WithParameters(common.StructToMap(opts)).
			WithKVparameters("projectId", s.getProjectID())
	}

	return resp.ToEntityServerGroup(), nil
}
