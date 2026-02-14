package v2

import (
	"github.com/dannyota/greennode-community-sdk/v2/greennode/client"
	"github.com/dannyota/greennode-community-sdk/v2/greennode/entity"
	sdkerror "github.com/dannyota/greennode-community-sdk/v2/greennode/sdkerror"
)

func (s *VolumeServiceV2) ListSnapshotsByBlockVolumeID(opts IListSnapshotsByBlockVolumeIDRequest) (*entity.ListSnapshots, sdkerror.Error) {
	url := listSnapshotsByBlockVolumeIDURL(s.VServerClient, opts)
	resp := new(ListSnapshotsByBlockVolumeIDResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithOkCodes(200).
		WithJSONResponse(resp).
		WithJSONError(errResp)

	if _, sdkErr := s.VServerClient.Get(url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp).
			WithKVparameters(
				"projectId", s.getProjectID(),
				"volumeId", opts.GetBlockVolumeID())
	}

	return resp.ToEntityListSnapshots(), nil
}

func (s *VolumeServiceV2) CreateSnapshotByBlockVolumeID(opts ICreateSnapshotByBlockVolumeIDRequest) (*entity.Snapshot, sdkerror.Error) {
	url := createSnapshotByBlockVolumeIDURL(s.VServerClient, opts)
	resp := new(CreateSnapshotByBlockVolumeIDResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithOkCodes(200).
		WithJSONResponse(resp).
		WithJSONError(errResp).
		WithJSONBody(opts.ToRequestBody())

	if _, sdkErr := s.VServerClient.Post(url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.WithErrorVolumeNotFound(errResp),
			sdkerror.WithErrorSnapshotNameNotValid(errResp)).
			WithKVparameters(
				"projectId", s.getProjectID(),
				"volumeId", opts.GetBlockVolumeID())
	}

	return resp.ToEntitySnapshot(), nil
}

func (s *VolumeServiceV2) DeleteSnapshotByID(opts IDeleteSnapshotByIDRequest) sdkerror.Error {
	url := deleteSnapshotByIDURL(s.VServerClient, opts)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithOkCodes(200).
		WithJSONError(errResp)

	if _, sdkErr := s.VServerClient.Delete(url, req); sdkErr != nil {
		return sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.WithErrorSnapshotNameNotFound(errResp)).
			WithKVparameters(
				"projectId", s.getProjectID(),
				"snapshotId", opts.GetSnapshotID())
	}

	return nil
}
