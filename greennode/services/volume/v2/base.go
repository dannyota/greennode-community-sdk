package v2

import "github.com/dannyota/greennode-community-sdk/v2/greennode/client"

type VolumeServiceV2 struct {
	VServerClient client.ServiceClient
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
