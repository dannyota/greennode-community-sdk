package v2

import "github.com/dannyota/greennode-community-sdk/v2/greennode/client"

func listAllQuotaUsedUrl(sc client.ServiceClient) string {
	return sc.ServiceURL(
		sc.GetProjectId(),
		"quotas",
		"quotaUsed")
}
