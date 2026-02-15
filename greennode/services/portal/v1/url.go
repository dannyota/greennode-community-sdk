package v1

import "github.com/dannyota/greennode-community-sdk/v2/greennode/client"

func getPortalInfoURL(sc client.ServiceClient, opts *GetPortalInfoRequest) string {
	return sc.ServiceURL("projects", opts.GetBackEndProjectID(), "detail")
}

func listProjectsURL(sc client.ServiceClient) string {
	return sc.ServiceURL("projects")
}

func listZonesURL(sc client.ServiceClient) string {
	return sc.ServiceURL(
		sc.GetProjectID(),
		"zones")
}
