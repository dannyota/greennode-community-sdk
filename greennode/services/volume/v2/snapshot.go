package v2

import (
	"context"

	"github.com/dannyota/greennode-community-sdk/v2/greennode/client"
	sdkerror "github.com/dannyota/greennode-community-sdk/v2/greennode/sdkerror"
)

func (s *VolumeServiceV2) ListSnapshotsByBlockVolumeID(ctx context.Context, opts *ListSnapshotsByBlockVolumeIDRequest) (*ListSnapshots, error) {
	url := listSnapshotsByBlockVolumeIDURL(s.Client, opts)
	resp := new(ListSnapshotsByBlockVolumeIDResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithOkCodes(200).
		WithJSONResponse(resp).
		WithJSONError(errResp)

	if _, sdkErr := s.Client.Get(ctx, url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp).
			WithKVparameters(
				"projectId", s.getProjectID(),
				"volumeId", opts.BlockVolumeID)
	}

	return resp.ToEntityListSnapshots(), nil
}

func (s *VolumeServiceV2) CreateSnapshotByBlockVolumeID(ctx context.Context, opts *CreateSnapshotByBlockVolumeIDRequest) (*Snapshot, error) {
	url := createSnapshotByBlockVolumeIDURL(s.Client, opts)
	resp := new(CreateSnapshotByBlockVolumeIDResponse)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithOkCodes(200).
		WithJSONResponse(resp).
		WithJSONError(errResp).
		WithJSONBody(opts)

	if _, sdkErr := s.Client.Post(ctx, url, req); sdkErr != nil {
		return nil, sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.EcVServerVolumeNotFound,
			sdkerror.EcVServerSnapshotNameNotValid).
			WithKVparameters(
				"projectId", s.getProjectID(),
				"volumeId", opts.BlockVolumeID)
	}

	return resp.ToEntitySnapshot(), nil
}

func (s *VolumeServiceV2) DeleteSnapshotByID(ctx context.Context, opts *DeleteSnapshotByIDRequest) error {
	url := deleteSnapshotByIDURL(s.Client, opts)
	errResp := sdkerror.NewErrorResponse(sdkerror.NormalErrorType)
	req := client.NewRequest().
		WithOkCodes(200).
		WithJSONError(errResp)

	if _, sdkErr := s.Client.Delete(ctx, url, req); sdkErr != nil {
		return sdkerror.SdkErrorHandler(sdkErr, errResp,
			sdkerror.EcVServerSnapshotNotFound).
			WithKVparameters(
				"projectId", s.getProjectID(),
				"snapshotId", opts.SnapshotID)
	}

	return nil
}
