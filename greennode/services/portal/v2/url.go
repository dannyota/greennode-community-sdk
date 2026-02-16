package v2

import "github.com/dannyota/greennode-community-sdk/v2/greennode/client"

func listAllQuotaUsedURL(sc *client.ServiceClient) string {
	return sc.ServiceURL(
		sc.ProjectID(),
		"quotas",
		"quotaUsed")
}
