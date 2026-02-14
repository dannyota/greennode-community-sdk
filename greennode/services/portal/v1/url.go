package v1

import "github.com/dannyota/greennode-community-sdk/v2/greennode/client"

func getPortalInfoUrl(sc client.ServiceClient, opts IGetPortalInfoRequest) string {
	return sc.ServiceURL("projects", opts.GetBackEndProjectId(), "detail")
}

func listProjectsUrl(sc client.ServiceClient) string {
	return sc.ServiceURL("projects")
}

func listZonesUrl(sc client.ServiceClient) string {
	return sc.ServiceURL(
		sc.GetProjectId(),
		"zones")
}
