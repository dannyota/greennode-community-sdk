package v2

import lsclient "github.com/dannyota/greennode-community-sdk/v2/greennode/client"

func listAllQuotaUsedUrl(psc lsclient.IServiceClient) string {
	return psc.ServiceURL(
		psc.GetProjectId(),
		"quotas",
		"quotaUsed")
}
