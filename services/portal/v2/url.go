package v2

import "danny.vn/greennode/client"

func listRegionsURL(sc *client.ServiceClient) string {
	return sc.ServiceURL(
		sc.ProjectID,
		"region")
}

func listAllQuotaUsedURL(sc *client.ServiceClient) string {
	return sc.ServiceURL(
		sc.ProjectID,
		"quotas",
		"quotaUsed")
}
