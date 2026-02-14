package v2

import (
	"github.com/dannyota/greennode-community-sdk/v2/greennode/client"
	"github.com/dannyota/greennode-community-sdk/v2/greennode/entity"
	sdkerror "github.com/dannyota/greennode-community-sdk/v2/greennode/sdkerror"
)

func (s *VolumeServiceV2) CreateBlockVolume(opts ICreateBlockVolumeRequest) (*entity.Volume, sdkerror.Error) {
	url := createBlockVolumeUrl(s.VServerClient)
	resp := new(CreateBlockVolumeResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithOkCodes(202).
		WithJsonBody(opts.ToRequestBody()).
		WithJsonResponse(resp).
		WithJsonError(errResp)

	if _, sdkErr := s.VServerClient.Post(url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.WithErrorVolumeTypeNotFound(errResp),
			sdkerror.WithErrorVolumeSizeOutOfRange(errResp),
			sdkerror.WithErrorVolumeSizeExceedGlobalQuota(errResp),
			sdkerror.WithErrorVolumeNameNotValid(errResp)).
			WithKVparameters("projectId", s.getProjectId()).
			WithParameters(opts.ToMap())
	}

	return resp.ToEntityVolume(), nil
}

func (s *VolumeServiceV2) DeleteBlockVolumeById(opts IDeleteBlockVolumeByIdRequest) sdkerror.Error {
	url := deleteBlockVolumeByIdUrl(s.VServerClient, opts)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithOkCodes(202).
		WithJsonError(errResp)

	if _, sdkErr := s.VServerClient.Delete(url, req); sdkErr != nil {
		return sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.WithErrorVolumeNotFound(errResp)).
			WithKVparameters(
				"projectId", s.getProjectId(),
				"volumeId", opts.GetBlockVolumeId())
	}

	return nil
}

func (s *VolumeServiceV2) ListBlockVolumes(opts IListBlockVolumesRequest) (*entity.ListVolumes, sdkerror.Error) {
	url := listBlockVolumesUrl(s.VServerClient, opts)
	resp := new(ListBlockVolumesResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithOkCodes(200).
		WithJsonResponse(resp).
		WithJsonError(errResp)

	if _, sdkErr := s.VServerClient.Get(url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.WithErrorPagingInvalid(errResp)).
			WithKVparameters("projectId", s.getProjectId()).
			WithParameters(opts.ToMap())
	}

	return resp.ToEntityListVolumes(), nil
}

func (s *VolumeServiceV2) GetBlockVolumeById(opts IGetBlockVolumeByIdRequest) (*entity.Volume, sdkerror.Error) {
	url := getBlockVolumeByIdUrl(s.VServerClient, opts)
	resp := new(GetBlockVolumeByIdResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithOkCodes(200).
		WithJsonResponse(resp).
		WithJsonError(errResp)

	if _, sdkErr := s.VServerClient.Get(url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.WithErrorVolumeNotFound(errResp)).
			WithKVparameters(
				"projectId", s.getProjectId(),
				"volumeId", opts.GetBlockVolumeId())
	}

	return resp.ToEntityVolume(), nil
}

func (s *VolumeServiceV2) ResizeBlockVolumeById(opts IResizeBlockVolumeByIdRequest) (*entity.Volume, sdkerror.Error) {
	url := resizeBlockVolumeByIdUrl(s.VServerClient, opts)
	resp := new(ResizeBlockVolumeByIdResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithOkCodes(202).
		WithJsonBody(opts.ToRequestBody()).
		WithJsonResponse(resp).
		WithJsonError(errResp)

	if _, sdkErr := s.VServerClient.Put(url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.WithErrorVolumeNotFound(errResp),
			sdkerror.WithErrorVolumeSizeOutOfRange(errResp),
			sdkerror.WithErrorVolumeMustSameZone(errResp),
			sdkerror.WithErrorVolumeUnchanged(errResp)).
			WithKVparameters(
				"projectId", s.getProjectId(),
				"volumeId", opts.GetBlockVolumeId(),
				"size", opts.GetSize())
	}

	return resp.ToEntityVolume(), nil
}

func (s *VolumeServiceV2) GetUnderBlockVolumeId(opts IGetUnderBlockVolumeIdRequest) (*entity.Volume, sdkerror.Error) {
	url := getUnderBlockVolumeIdUrl(s.VServerClient, opts)
	resp := new(GetUnderBlockVolumeIdResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithOkCodes(200).
		WithJsonResponse(resp).
		WithJsonError(errResp)

	if _, sdkErr := s.VServerClient.Get(url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.WithErrorVolumeNotFound(errResp)).
			WithKVparameters(
				"projectId", s.getProjectId(),
				"volumeId", opts.GetBlockVolumeId())
	}

	return resp.ToEntityVolume(), nil
}

func (s *VolumeServiceV2) MigrateBlockVolumeById(opts IMigrateBlockVolumeByIdRequest) sdkerror.Error {
	url := migrateBlockVolumeByIdUrl(s.VServerClient, opts)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	resp := map[string]interface{}{}
	req := client.NewRequest().
		WithOkCodes(204).
		WithJsonBody(opts.ToRequestBody()).
		WithJsonError(errResp).
		WithJsonResponse(&resp)

	if _, sdkErr := s.VServerClient.Put(url, req); sdkErr != nil {
		sdkErr = sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.WithErrorVolumeMigrateInSameZone(errResp),
			sdkerror.WithErrorVolumeMigrateMissingInit(errResp),
			sdkerror.WithErrorVolumeMigrateNeedProcess(errResp),
			sdkerror.WithErrorVolumeMigrateNeedConfirm(errResp),
			sdkerror.WithErrorVolumeMigrateBeingProcess(errResp),
			sdkerror.WithErrorVolumeMigrateProcessingConfirm(errResp),
			sdkerror.WithErrorVolumeMigrateBeingMigrating(errResp), // should under WithErrorVolumeMigrateBeingProcess
			sdkerror.WithErrorVolumeMigrateBeingFinish(errResp),
			sdkerror.WithErrorVolumeNotFound(errResp)).
			WithKVparameters(
				"projectId", s.getProjectId(),
				"volumeId", opts.GetBlockVolumeId())

		if opts.IsConfirm() {
			switch sdkErr.GetErrorCode() {
			case sdkerror.EcVServerVolumeMigrateMissingInit:
				opts = opts.WithAction(InitMigrateAction)
				return s.MigrateBlockVolumeById(opts)
			case sdkerror.EcVServerVolumeMigrateNeedProcess:
				opts = opts.WithAction(ProcessMigrateAction)
				return s.MigrateBlockVolumeById(opts)
			case sdkerror.EcVServerVolumeMigrateNeedConfirm:
				opts = opts.WithAction(ConfirmMigrateAction)
				return s.MigrateBlockVolumeById(opts)
			}
		}

		return sdkErr
	}

	return nil
}
