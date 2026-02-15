package v1

import (
	"github.com/dannyota/greennode-community-sdk/v2/greennode/client"
)

const (
	defaultZoneGetVolumeTypeZonesRequest = "HCM03-1A"
)

type VolumeServiceV1 struct {
	VServerClient *client.ServiceClient
}

func (s *VolumeServiceV1) getProjectID() string {
	return s.VServerClient.GetProjectID()
}
