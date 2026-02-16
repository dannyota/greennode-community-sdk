package v2

import (
	"context"

	"github.com/dannyota/greennode-community-sdk/v2/greennode/client"
	sdkerror "github.com/dannyota/greennode-community-sdk/v2/greennode/sdkerror"
	"github.com/dannyota/greennode-community-sdk/v2/greennode/services/common"
)

type VolumeServiceV2 struct {
	Client *client.ServiceClient
}


const (
	defaultPageListBlockVolumesRequest = 1
	defaultSizeListBlockVolumesRequest = 10000

	defaultPageListSnapshotsByBlockVolumeIDRequest = 1
	defaultSizeListSnapshotsByBlockVolumeIDRequest = 10000
)

func (s *VolumeServiceV2) CreateBlockVolume(ctx context.Context, opts *CreateBlockVolumeRequest) (*Volume, error) {
	url := createBlockVolumeURL(s.Client)
	resp := new(CreateBlockVolumeResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithOkCodes(202).
		WithJSONBody(opts).
		WithJSONResponse(resp).
		WithJSONError(errResp)

	if _, sdkErr := s.Client.Post(ctx, url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.EcVServerVolumeTypeNotFound,
			sdkerror.EcVServerVolumeSizeOutOfRange,
			sdkerror.EcVServerVolumeSizeExceedGlobalQuota,
			sdkerror.EcVServerVolumeNameNotValid).
			WithKVparameters("projectId", s.Client.ProjectID).
			WithParameters(common.StructToMap(opts))
	}

	return resp.ToEntityVolume(), nil
}

func (s *VolumeServiceV2) DeleteBlockVolumeByID(ctx context.Context, opts *DeleteBlockVolumeByIDRequest) error {
	url := deleteBlockVolumeByIDURL(s.Client, opts)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithOkCodes(202).
		WithJSONError(errResp)

	if _, sdkErr := s.Client.Delete(ctx, url, req); sdkErr != nil {
		return sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.EcVServerVolumeNotFound).
			WithKVparameters(
				"projectId", s.Client.ProjectID,
				"volumeId", opts.BlockVolumeID)
	}

	return nil
}

func (s *VolumeServiceV2) ListBlockVolumes(ctx context.Context, opts *ListBlockVolumesRequest) (*ListVolumes, error) {
	url := listBlockVolumesURL(s.Client, opts)
	resp := new(ListBlockVolumesResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithOkCodes(200).
		WithJSONResponse(resp).
		WithJSONError(errResp)

	if _, sdkErr := s.Client.Get(ctx, url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.EcPagingInvalid).
			WithKVparameters("projectId", s.Client.ProjectID).
			WithParameters(common.StructToMap(opts))
	}

	return resp.ToEntityListVolumes(), nil
}

func (s *VolumeServiceV2) GetBlockVolumeByID(ctx context.Context, opts *GetBlockVolumeByIDRequest) (*Volume, error) {
	url := getBlockVolumeByIDURL(s.Client, opts)
	resp := new(GetBlockVolumeByIDResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithOkCodes(200).
		WithJSONResponse(resp).
		WithJSONError(errResp)

	if _, sdkErr := s.Client.Get(ctx, url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.EcVServerVolumeNotFound).
			WithKVparameters(
				"projectId", s.Client.ProjectID,
				"volumeId", opts.BlockVolumeID)
	}

	return resp.ToEntityVolume(), nil
}

func (s *VolumeServiceV2) ResizeBlockVolumeByID(ctx context.Context, opts *ResizeBlockVolumeByIDRequest) (*Volume, error) {
	url := resizeBlockVolumeByIDURL(s.Client, opts)
	resp := new(ResizeBlockVolumeByIDResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithOkCodes(202).
		WithJSONBody(opts).
		WithJSONResponse(resp).
		WithJSONError(errResp)

	if _, sdkErr := s.Client.Put(ctx, url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.EcVServerVolumeNotFound,
			sdkerror.EcVServerVolumeSizeOutOfRange,
			sdkerror.EcVServerVolumeMustSameZone,
			sdkerror.EcVServerVolumeUnchanged).
			WithKVparameters(
				"projectId", s.Client.ProjectID,
				"volumeId", opts.BlockVolumeID,
				"size", opts.NewSize)
	}

	return resp.ToEntityVolume(), nil
}

func (s *VolumeServiceV2) GetUnderBlockVolumeID(ctx context.Context, opts *GetUnderBlockVolumeIDRequest) (*Volume, error) {
	url := getUnderBlockVolumeIDURL(s.Client, opts)
	resp := new(GetUnderBlockVolumeIDResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithOkCodes(200).
		WithJSONResponse(resp).
		WithJSONError(errResp)

	if _, sdkErr := s.Client.Get(ctx, url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.EcVServerVolumeNotFound).
			WithKVparameters(
				"projectId", s.Client.ProjectID,
				"volumeId", opts.BlockVolumeID)
	}

	return resp.ToEntityVolume(), nil
}

func (s *VolumeServiceV2) MigrateBlockVolumeByID(ctx context.Context, opts *MigrateBlockVolumeByIDRequest) error {
	url := migrateBlockVolumeByIDURL(s.Client, opts)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	resp := map[string]any{}
	req := client.NewRequest().
		WithOkCodes(204).
		WithJSONBody(opts).
		WithJSONError(errResp).
		WithJSONResponse(&resp)

	if _, err := s.Client.Put(ctx, url, req); err != nil {
		enriched := sdkerror.SdkErrorHandler(err, errResp,
			sdkerror.EcVServerVolumeMigrateInSameZone,
			sdkerror.EcVServerVolumeMigrateMissingInit,
			sdkerror.EcVServerVolumeMigrateNeedProcess,
			sdkerror.EcVServerVolumeMigrateNeedConfirm,
			sdkerror.EcVServerVolumeMigrateBeingProcess,
			sdkerror.EcVServerVolumeMigrateProcessingConfirm,
			sdkerror.EcVServerVolumeMigrateBeingMigrating, // should under WithErrorVolumeMigrateBeingProcess
			sdkerror.EcVServerVolumeMigrateBeingFinish,
			sdkerror.EcVServerVolumeNotFound).
			WithKVparameters(
				"projectId", s.Client.ProjectID,
				"volumeId", opts.BlockVolumeID)

		if opts.ConfirmMigrate {
			switch enriched.ErrorCode() {
			case sdkerror.EcVServerVolumeMigrateMissingInit:
				opts.Action = InitMigrateAction
				return s.MigrateBlockVolumeByID(ctx, opts)
			case sdkerror.EcVServerVolumeMigrateNeedProcess:
				opts.Action = ProcessMigrateAction
				return s.MigrateBlockVolumeByID(ctx, opts)
			case sdkerror.EcVServerVolumeMigrateNeedConfirm:
				opts.Action = ConfirmMigrateAction
				return s.MigrateBlockVolumeByID(ctx, opts)
			}
		}

		return enriched
	}

	return nil
}
