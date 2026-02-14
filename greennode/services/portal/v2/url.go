package v2

import "github.com/dannyota/greennode-community-sdk/v2/greennode/client"

func listAllQuotaUsedUrl(psc client.IServiceClient) string {
	return psc.ServiceURL(
		psc.GetProjectId(),
		"quotas",
		"quotaUsed")
}
