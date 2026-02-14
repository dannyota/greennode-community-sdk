package v2

import (
	"github.com/dannyota/greennode-community-sdk/v2/greennode/client"
	"github.com/dannyota/greennode-community-sdk/v2/greennode/entity"
	sdkerror "github.com/dannyota/greennode-community-sdk/v2/greennode/sdk_error"
)

func (s *VolumeServiceV2) CreateBlockVolume(popts ICreateBlockVolumeRequest) (*entity.Volume, sdkerror.IError) {
	url := createBlockVolumeUrl(s.VServerClient)
	resp := new(CreateBlockVolumeResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithOkCodes(202).
		WithJsonBody(popts.ToRequestBody()).
		WithJsonResponse(resp).
		WithJsonError(errResp)

	if _, sdkErr := s.VServerClient.Post(url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.WithErrorVolumeTypeNotFound(errResp),
			sdkerror.WithErrorVolumeSizeOutOfRange(errResp),
			sdkerror.WithErrorVolumeSizeExceedGlobalQuota(errResp),
			sdkerror.WithErrorVolumeNameNotValid(errResp)).
			WithKVparameters("projectId", s.getProjectId()).
			WithParameters(popts.ToMap())
	}

	return resp.ToEntityVolume(), nil
}

func (s *VolumeServiceV2) DeleteBlockVolumeById(popts IDeleteBlockVolumeByIdRequest) sdkerror.IError {
	url := deleteBlockVolumeByIdUrl(s.VServerClient, popts)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithOkCodes(202).
		WithJsonError(errResp)

	if _, sdkErr := s.VServerClient.Delete(url, req); sdkErr != nil {
		return sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.WithErrorVolumeNotFound(errResp)).
			WithKVparameters(
				"projectId", s.getProjectId(),
				"volumeId", popts.GetBlockVolumeId())
	}

	return nil
}

func (s *VolumeServiceV2) ListBlockVolumes(popts IListBlockVolumesRequest) (*entity.ListVolumes, sdkerror.IError) {
	url := listBlockVolumesUrl(s.VServerClient, popts)
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
			WithParameters(popts.ToMap())
	}

	return resp.ToEntityListVolumes(), nil
}

func (s *VolumeServiceV2) GetBlockVolumeById(popts IGetBlockVolumeByIdRequest) (*entity.Volume, sdkerror.IError) {
	url := getBlockVolumeByIdUrl(s.VServerClient, popts)
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
				"volumeId", popts.GetBlockVolumeId())
	}

	return resp.ToEntityVolume(), nil
}

func (s *VolumeServiceV2) ResizeBlockVolumeById(popts IResizeBlockVolumeByIdRequest) (*entity.Volume, sdkerror.IError) {
	url := resizeBlockVolumeByIdUrl(s.VServerClient, popts)
	resp := new(ResizeBlockVolumeByIdResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithOkCodes(202).
		WithJsonBody(popts.ToRequestBody()).
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
				"volumeId", popts.GetBlockVolumeId(),
				"size", popts.GetSize())
	}

	return resp.ToEntityVolume(), nil
}

func (s *VolumeServiceV2) GetUnderBlockVolumeId(popts IGetUnderBlockVolumeIdRequest) (*entity.Volume, sdkerror.IError) {
	url := getUnderBlockVolumeIdUrl(s.VServerClient, popts)
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
				"volumeId", popts.GetBlockVolumeId())
	}

	return resp.ToEntityVolume(), nil
}

func (s *VolumeServiceV2) MigrateBlockVolumeById(popts IMigrateBlockVolumeByIdRequest) sdkerror.IError {
	url := migrateBlockVolumeByIdUrl(s.VServerClient, popts)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	resp := map[string]interface{}{}
	req := client.NewRequest().
		WithOkCodes(204).
		WithJsonBody(popts.ToRequestBody()).
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
				"volumeId", popts.GetBlockVolumeId())

		if popts.IsConfirm() {
			switch sdkErr.GetErrorCode() {
			case sdkerror.EcVServerVolumeMigrateMissingInit:
				popts = popts.WithAction(InitMigrateAction)
				return s.MigrateBlockVolumeById(popts)
			case sdkerror.EcVServerVolumeMigrateNeedProcess:
				popts = popts.WithAction(ProcessMigrateAction)
				return s.MigrateBlockVolumeById(popts)
			case sdkerror.EcVServerVolumeMigrateNeedConfirm:
				popts = popts.WithAction(ConfirmMigrateAction)
				return s.MigrateBlockVolumeById(popts)
			}
		}

		return sdkErr
	}

	return nil
}
