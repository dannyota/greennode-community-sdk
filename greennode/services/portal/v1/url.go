package v1

import "github.com/dannyota/greennode-community-sdk/greennode/client"

func getPortalInfoURL(sc *client.ServiceClient, opts *GetPortalInfoRequest) string {
	return sc.ServiceURL("projects", opts.BackEndProjectID, "detail")
}

func listProjectsURL(sc *client.ServiceClient) string {
	return sc.ServiceURL("projects")
}

func listZonesURL(sc *client.ServiceClient) string {
	return sc.ServiceURL(
		sc.ProjectID,
		"zones")
}
