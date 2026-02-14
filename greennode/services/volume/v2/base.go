package v2

import "github.com/dannyota/greennode-community-sdk/v2/greennode/client"

type VolumeServiceV2 struct {
	VServerClient client.IServiceClient
}

func (s *VolumeServiceV2) getProjectId() string {
	return s.VServerClient.GetProjectId()
}

const (
	defaultPageListBlockVolumesRequest = 1
	defaultSizeListBlockVolumesRequest = 10000

	defaultPageListSnapshotsByBlockVolumeIdRequest = 1
	defaultSizeListSnapshotsByBlockVolumeIdRequest = 10000
)
