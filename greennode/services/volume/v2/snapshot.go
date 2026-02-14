package v2

import (
	"github.com/dannyota/greennode-community-sdk/v2/greennode/client"
	"github.com/dannyota/greennode-community-sdk/v2/greennode/entity"
	sdkerror "github.com/dannyota/greennode-community-sdk/v2/greennode/sdkerror"
)

func (s *VolumeServiceV2) ListSnapshotsByBlockVolumeId(opts IListSnapshotsByBlockVolumeIdRequest) (*entity.ListSnapshots, sdkerror.Error) {
	url := listSnapshotsByBlockVolumeIdUrl(s.VServerClient, opts)
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
				"volumeId", opts.GetBlockVolumeId())
	}

	return resp.ToEntityListSnapshots(), nil
}

func (s *VolumeServiceV2) CreateSnapshotByBlockVolumeId(opts ICreateSnapshotByBlockVolumeIdRequest) (*entity.Snapshot, sdkerror.Error) {
	url := createSnapshotByBlockVolumeIdUrl(s.VServerClient, opts)
	resp := new(CreateSnapshotByBlockVolumeIdResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithOkCodes(200).
		WithJsonResponse(resp).
		WithJsonError(errResp).
		WithJsonBody(opts.ToRequestBody())

	if _, sdkErr := s.VServerClient.Post(url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.WithErrorVolumeNotFound(errResp),
			sdkerror.WithErrorSnapshotNameNotValid(errResp)).
			WithKVparameters(
				"projectId", s.getProjectId(),
				"volumeId", opts.GetBlockVolumeId())
	}

	return resp.ToEntitySnapshot(), nil
}

func (s *VolumeServiceV2) DeleteSnapshotById(opts IDeleteSnapshotByIdRequest) sdkerror.Error {
	url := deleteSnapshotByIdUrl(s.VServerClient, opts)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithOkCodes(200).
		WithJsonError(errResp)

	if _, sdkErr := s.VServerClient.Delete(url, req); sdkErr != nil {
		return sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.WithErrorSnapshotNameNotFound(errResp)).
			WithKVparameters(
				"projectId", s.getProjectId(),
				"snapshotId", opts.GetSnapshotId())
	}

	return nil
}
