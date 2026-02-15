package v1

import "github.com/dannyota/greennode-community-sdk/v2/greennode/client"

func createSystemTagURL(sc *client.ServiceClient) string {
	return sc.ServiceURL(
		sc.GetProjectID(),
		"tags")
}
