package v1

import (
	lsclient "github.com/dannyota/greennode-community-sdk/v2/greennode/client"
)

const (
	defaultZoneGetVolumeTypeZonesRequest = "HCM03-1A"
)

type VolumeServiceV1 struct {
	VServerClient lsclient.IServiceClient
}

func (s *VolumeServiceV1) getProjectId() string {
	return s.VServerClient.GetProjectId()
}
