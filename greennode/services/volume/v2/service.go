package v2

import (
	"context"

	"github.com/dannyota/greennode-community-sdk/v2/greennode/client"
	"github.com/dannyota/greennode-community-sdk/v2/greennode/entity"
	sdkerror "github.com/dannyota/greennode-community-sdk/v2/greennode/sdkerror"
	"github.com/dannyota/greennode-community-sdk/v2/greennode/services/common"
)

type VolumeServiceV2 struct {
	VServerClient *client.ServiceClient
}

func (s *VolumeServiceV2) getProjectID() string {
	return s.VServerClient.GetProjectID()
}

const (
	defaultPageListBlockVolumesRequest = 1
	defaultSizeListBlockVolumesRequest = 10000

	defaultPageListSnapshotsByBlockVolumeIDRequest = 1
	defaultSizeListSnapshotsByBlockVolumeIDRequest = 10000
)

func (s *VolumeServiceV2) CreateBlockVolume(ctx context.Context, opts *CreateBlockVolumeRequest) (*entity.Volume, error) {
	url := createBlockVolumeURL(s.VServerClient)
	resp := new(CreateBlockVolumeResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithOkCodes(202).
		WithJSONBody(opts).
		WithJSONResponse(resp).
		WithJSONError(errResp)

	if _, sdkErr := s.VServerClient.Post(ctx, url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.EcVServerVolumeTypeNotFound,
			sdkerror.EcVServerVolumeSizeOutOfRange,
			sdkerror.EcVServerVolumeSizeExceedGlobalQuota,
			sdkerror.EcVServerVolumeNameNotValid).
			WithKVparameters("projectId", s.getProjectID()).
			WithParameters(common.StructToMap(opts))
	}

	return resp.ToEntityVolume(), nil
}

func (s *VolumeServiceV2) DeleteBlockVolumeByID(ctx context.Context, opts *DeleteBlockVolumeByIDRequest) error {
	url := deleteBlockVolumeByIDURL(s.VServerClient, opts)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithOkCodes(202).
		WithJSONError(errResp)

	if _, sdkErr := s.VServerClient.Delete(ctx, url, req); sdkErr != nil {
		return sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.EcVServerVolumeNotFound).
			WithKVparameters(
				"projectId", s.getProjectID(),
				"volumeId", opts.GetBlockVolumeID())
	}

	return nil
}

func (s *VolumeServiceV2) ListBlockVolumes(ctx context.Context, opts *ListBlockVolumesRequest) (*entity.ListVolumes, error) {
	url := listBlockVolumesURL(s.VServerClient, opts)
	resp := new(ListBlockVolumesResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithOkCodes(200).
		WithJSONResponse(resp).
		WithJSONError(errResp)

	if _, sdkErr := s.VServerClient.Get(ctx, url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.EcPagingInvalid).
			WithKVparameters("projectId", s.getProjectID()).
			WithParameters(common.StructToMap(opts))
	}

	return resp.ToEntityListVolumes(), nil
}

func (s *VolumeServiceV2) GetBlockVolumeByID(ctx context.Context, opts *GetBlockVolumeByIDRequest) (*entity.Volume, error) {
	url := getBlockVolumeByIDURL(s.VServerClient, opts)
	resp := new(GetBlockVolumeByIDResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithOkCodes(200).
		WithJSONResponse(resp).
		WithJSONError(errResp)

	if _, sdkErr := s.VServerClient.Get(ctx, url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.EcVServerVolumeNotFound).
			WithKVparameters(
				"projectId", s.getProjectID(),
				"volumeId", opts.GetBlockVolumeID())
	}

	return resp.ToEntityVolume(), nil
}

func (s *VolumeServiceV2) ResizeBlockVolumeByID(ctx context.Context, opts *ResizeBlockVolumeByIDRequest) (*entity.Volume, error) {
	url := resizeBlockVolumeByIDURL(s.VServerClient, opts)
	resp := new(ResizeBlockVolumeByIDResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithOkCodes(202).
		WithJSONBody(opts).
		WithJSONResponse(resp).
		WithJSONError(errResp)

	if _, sdkErr := s.VServerClient.Put(ctx, url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.EcVServerVolumeNotFound,
			sdkerror.EcVServerVolumeSizeOutOfRange,
			sdkerror.EcVServerVolumeMustSameZone,
			sdkerror.EcVServerVolumeUnchanged).
			WithKVparameters(
				"projectId", s.getProjectID(),
				"volumeId", opts.GetBlockVolumeID(),
				"size", opts.GetSize())
	}

	return resp.ToEntityVolume(), nil
}

func (s *VolumeServiceV2) GetUnderBlockVolumeID(ctx context.Context, opts *GetUnderBlockVolumeIDRequest) (*entity.Volume, error) {
	url := getUnderBlockVolumeIDURL(s.VServerClient, opts)
	resp := new(GetUnderBlockVolumeIDResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithOkCodes(200).
		WithJSONResponse(resp).
		WithJSONError(errResp)

	if _, sdkErr := s.VServerClient.Get(ctx, url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.EcVServerVolumeNotFound).
			WithKVparameters(
				"projectId", s.getProjectID(),
				"volumeId", opts.GetBlockVolumeID())
	}

	return resp.ToEntityVolume(), nil
}

func (s *VolumeServiceV2) MigrateBlockVolumeByID(ctx context.Context, opts *MigrateBlockVolumeByIDRequest) error {
	url := migrateBlockVolumeByIDURL(s.VServerClient, opts)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	resp := map[string]any{}
	req := client.NewRequest().
		WithOkCodes(204).
		WithJSONBody(opts).
		WithJSONError(errResp).
		WithJSONResponse(&resp)

	if _, err := s.VServerClient.Put(ctx, url, req); err != nil {
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
				"projectId", s.getProjectID(),
				"volumeId", opts.GetBlockVolumeID())

		if opts.IsConfirm() {
			switch enriched.ErrorCode() {
			case sdkerror.EcVServerVolumeMigrateMissingInit:
				opts = opts.WithAction(InitMigrateAction)
				return s.MigrateBlockVolumeByID(ctx, opts)
			case sdkerror.EcVServerVolumeMigrateNeedProcess:
				opts = opts.WithAction(ProcessMigrateAction)
				return s.MigrateBlockVolumeByID(ctx, opts)
			case sdkerror.EcVServerVolumeMigrateNeedConfirm:
				opts = opts.WithAction(ConfirmMigrateAction)
				return s.MigrateBlockVolumeByID(ctx, opts)
			}
		}

		return enriched
	}

	return nil
}
