package v2

import (
	"github.com/dannyota/greennode-community-sdk/v2/greennode/client"
	"github.com/dannyota/greennode-community-sdk/v2/greennode/entity"
	sdkerror "github.com/dannyota/greennode-community-sdk/v2/greennode/sdkerror"
)

func (s *VolumeServiceV2) ListSnapshotsByBlockVolumeID(opts *ListSnapshotsByBlockVolumeIDRequest) (*entity.ListSnapshots, error) {
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

func (s *VolumeServiceV2) CreateSnapshotByBlockVolumeID(opts *CreateSnapshotByBlockVolumeIDRequest) (*entity.Snapshot, error) {
	url := createSnapshotByBlockVolumeIDURL(s.VServerClient, opts)
	resp := new(CreateSnapshotByBlockVolumeIDResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithOkCodes(200).
		WithJSONResponse(resp).
		WithJSONError(errResp).
		WithJSONBody(opts)

	if _, sdkErr := s.VServerClient.Post(url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.EcVServerVolumeNotFound,
			sdkerror.EcVServerSnapshotNameNotValid).
			WithKVparameters(
				"projectId", s.getProjectID(),
				"volumeId", opts.GetBlockVolumeID())
	}

	return resp.ToEntitySnapshot(), nil
}

func (s *VolumeServiceV2) DeleteSnapshotByID(opts *DeleteSnapshotByIDRequest) error {
	url := deleteSnapshotByIDURL(s.VServerClient, opts)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithOkCodes(200).
		WithJSONError(errResp)

	if _, sdkErr := s.VServerClient.Delete(url, req); sdkErr != nil {
		return sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.EcVServerSnapshotNotFound).
			WithKVparameters(
				"projectId", s.getProjectID(),
				"snapshotId", opts.GetSnapshotID())
	}

	return nil
}
