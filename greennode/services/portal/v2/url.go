package v2

import "github.com/dannyota/greennode-community-sdk/greennode/client"

func listAllQuotaUsedURL(sc *client.ServiceClient) string {
	return sc.ServiceURL(
		sc.ProjectID,
		"quotas",
		"quotaUsed")
}
