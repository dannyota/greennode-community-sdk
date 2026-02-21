package v1

import "github.com/dannyota/greennode-community-sdk/greennode/client"

func createSystemTagURL(sc *client.ServiceClient) string {
	return sc.ServiceURL(
		sc.ProjectID,
		"tags")
}
