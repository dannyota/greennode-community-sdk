package v2

import (
	"github.com/dannyota/greennode-community-sdk/v2/greennode/client"
	"github.com/dannyota/greennode-community-sdk/v2/greennode/entity"
	sdkerror "github.com/dannyota/greennode-community-sdk/v2/greennode/sdkerror"
)

func (s *VolumeServiceV2) CreateBlockVolume(opts *CreateBlockVolumeRequest) (*entity.Volume, error) {
	url := createBlockVolumeURL(s.VServerClient)
	resp := new(CreateBlockVolumeResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithOkCodes(202).
		WithJSONBody(opts.ToRequestBody()).
		WithJSONResponse(resp).
		WithJSONError(errResp)

	if _, sdkErr := s.VServerClient.Post(url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.EcVServerVolumeTypeNotFound,
			sdkerror.EcVServerVolumeSizeOutOfRange,
			sdkerror.EcVServerVolumeSizeExceedGlobalQuota,
			sdkerror.EcVServerVolumeNameNotValid).
			WithKVparameters("projectId", s.getProjectID()).
			WithParameters(opts.ToMap())
	}

	return resp.ToEntityVolume(), nil
}

func (s *VolumeServiceV2) DeleteBlockVolumeByID(opts *DeleteBlockVolumeByIDRequest) error {
	url := deleteBlockVolumeByIDURL(s.VServerClient, opts)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithOkCodes(202).
		WithJSONError(errResp)

	if _, sdkErr := s.VServerClient.Delete(url, req); sdkErr != nil {
		return sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.EcVServerVolumeNotFound).
			WithKVparameters(
				"projectId", s.getProjectID(),
				"volumeId", opts.GetBlockVolumeID())
	}

	return nil
}

func (s *VolumeServiceV2) ListBlockVolumes(opts *ListBlockVolumesRequest) (*entity.ListVolumes, error) {
	url := listBlockVolumesURL(s.VServerClient, opts)
	resp := new(ListBlockVolumesResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithOkCodes(200).
		WithJSONResponse(resp).
		WithJSONError(errResp)

	if _, sdkErr := s.VServerClient.Get(url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.EcPagingInvalid).
			WithKVparameters("projectId", s.getProjectID()).
			WithParameters(opts.ToMap())
	}

	return resp.ToEntityListVolumes(), nil
}

func (s *VolumeServiceV2) GetBlockVolumeByID(opts *GetBlockVolumeByIDRequest) (*entity.Volume, error) {
	url := getBlockVolumeByIDURL(s.VServerClient, opts)
	resp := new(GetBlockVolumeByIDResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithOkCodes(200).
		WithJSONResponse(resp).
		WithJSONError(errResp)

	if _, sdkErr := s.VServerClient.Get(url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.EcVServerVolumeNotFound).
			WithKVparameters(
				"projectId", s.getProjectID(),
				"volumeId", opts.GetBlockVolumeID())
	}

	return resp.ToEntityVolume(), nil
}

func (s *VolumeServiceV2) ResizeBlockVolumeByID(opts *ResizeBlockVolumeByIDRequest) (*entity.Volume, error) {
	url := resizeBlockVolumeByIDURL(s.VServerClient, opts)
	resp := new(ResizeBlockVolumeByIDResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithOkCodes(202).
		WithJSONBody(opts.ToRequestBody()).
		WithJSONResponse(resp).
		WithJSONError(errResp)

	if _, sdkErr := s.VServerClient.Put(url, req); sdkErr != nil {
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

func (s *VolumeServiceV2) GetUnderBlockVolumeID(opts *GetUnderBlockVolumeIDRequest) (*entity.Volume, error) {
	url := getUnderBlockVolumeIDURL(s.VServerClient, opts)
	resp := new(GetUnderBlockVolumeIDResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithOkCodes(200).
		WithJSONResponse(resp).
		WithJSONError(errResp)

	if _, sdkErr := s.VServerClient.Get(url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.EcVServerVolumeNotFound).
			WithKVparameters(
				"projectId", s.getProjectID(),
				"volumeId", opts.GetBlockVolumeID())
	}

	return resp.ToEntityVolume(), nil
}

func (s *VolumeServiceV2) MigrateBlockVolumeByID(opts *MigrateBlockVolumeByIDRequest) error {
	url := migrateBlockVolumeByIDURL(s.VServerClient, opts)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	resp := map[string]any{}
	req := client.NewRequest().
		WithOkCodes(204).
		WithJSONBody(opts.ToRequestBody()).
		WithJSONError(errResp).
		WithJSONResponse(&resp)

	if _, err := s.VServerClient.Put(url, req); err != nil {
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
				return s.MigrateBlockVolumeByID(opts)
			case sdkerror.EcVServerVolumeMigrateNeedProcess:
				opts = opts.WithAction(ProcessMigrateAction)
				return s.MigrateBlockVolumeByID(opts)
			case sdkerror.EcVServerVolumeMigrateNeedConfirm:
				opts = opts.WithAction(ConfirmMigrateAction)
				return s.MigrateBlockVolumeByID(opts)
			}
		}

		return enriched
	}

	return nil
}
