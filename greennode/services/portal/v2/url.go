package v2

import "danny.vn/greennode/greennode/client"

func listAllQuotaUsedURL(sc *client.ServiceClient) string {
	return sc.ServiceURL(
		sc.ProjectID,
		"quotas",
		"quotaUsed")
}
