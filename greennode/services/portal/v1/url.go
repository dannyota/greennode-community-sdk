package v1

import "github.com/dannyota/greennode-community-sdk/v2/greennode/client"

func getPortalInfoUrl(psc client.IServiceClient, popts IGetPortalInfoRequest) string {
	return psc.ServiceURL("projects", popts.GetBackEndProjectId(), "detail")
}

func listProjectsUrl(psc client.IServiceClient) string {
	return psc.ServiceURL("projects")
}

func listZonesUrl(psc client.IServiceClient) string {
	return psc.ServiceURL(
		psc.GetProjectId(),
		"zones")
}
