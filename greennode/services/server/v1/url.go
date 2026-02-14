package v1

import lsclient "github.com/dannyota/greennode-community-sdk/v2/greennode/client"

func createSystemTagUrl(psc lsclient.IServiceClient) string {
	return psc.ServiceURL(
		psc.GetProjectId(),
		"tags")
}
