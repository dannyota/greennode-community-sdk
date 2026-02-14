package v1

import "github.com/dannyota/greennode-community-sdk/v2/greennode/client"

func createSystemTagUrl(sc client.ServiceClient) string {
	return sc.ServiceURL(
		sc.GetProjectId(),
		"tags")
}
