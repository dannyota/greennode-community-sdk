package v2

import (
	"github.com/dannyota/greennode-community-sdk/v2/greennode/client"
)

func getAccessTokenURL(sc client.ServiceClient) string {
	return sc.ServiceURL("auth", "token")
}
