package v2

import (
	"github.com/dannyota/greennode-community-sdk/v2/greennode/client"
	"github.com/dannyota/greennode-community-sdk/v2/greennode/entity"
	sdkerror "github.com/dannyota/greennode-community-sdk/v2/greennode/sdkerror"
)

func (s *VolumeServiceV2) ListSnapshotsByBlockVolumeId(popts IListSnapshotsByBlockVolumeIdRequest) (*entity.ListSnapshots, sdkerror.IError) {
	url := listSnapshotsByBlockVolumeIdUrl(s.VServerClient, popts)
	resp := new(ListSnapshotsByBlockVolumeIdResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithOkCodes(200).
		WithJsonResponse(resp).
		WithJsonError(errResp)

	if _, sdkErr := s.VServerClient.Get(url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp).
			WithKVparameters(
				"projectId", s.getProjectId(),
				"volumeId", popts.GetBlockVolumeId())
	}

	return resp.ToEntityListSnapshots(), nil
}

func (s *VolumeServiceV2) CreateSnapshotByBlockVolumeId(popts ICreateSnapshotByBlockVolumeIdRequest) (*entity.Snapshot, sdkerror.IError) {
	url := createSnapshotByBlockVolumeIdUrl(s.VServerClient, popts)
	resp := new(CreateSnapshotByBlockVolumeIdResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithOkCodes(200).
		WithJsonResponse(resp).
		WithJsonError(errResp).
		WithJsonBody(popts.ToRequestBody())

	if _, sdkErr := s.VServerClient.Post(url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.WithErrorVolumeNotFound(errResp),
			sdkerror.WithErrorSnapshotNameNotValid(errResp)).
			WithKVparameters(
				"projectId", s.getProjectId(),
				"volumeId", popts.GetBlockVolumeId())
	}

	return resp.ToEntitySnapshot(), nil
}

func (s *VolumeServiceV2) DeleteSnapshotById(popts IDeleteSnapshotByIdRequest) sdkerror.IError {
	url := deleteSnapshotByIdUrl(s.VServerClient, popts)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithOkCodes(200).
		WithJsonError(errResp)

	if _, sdkErr := s.VServerClient.Delete(url, req); sdkErr != nil {
		return sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.WithErrorSnapshotNameNotFound(errResp)).
			WithKVparameters(
				"projectId", s.getProjectId(),
				"snapshotId", popts.GetSnapshotId())
	}

	return nil
}
