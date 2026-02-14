package v1

import "github.com/dannyota/greennode-community-sdk/v2/greennode/client"

func createSystemTagUrl(psc client.IServiceClient) string {
	return psc.ServiceURL(
		psc.GetProjectId(),
		"tags")
}
