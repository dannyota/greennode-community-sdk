package v1

import "danny.vn/greennode/client"

func createSystemTagURL(sc *client.ServiceClient) string {
	return sc.ServiceURL(
		sc.ProjectID,
		"tags")
}
