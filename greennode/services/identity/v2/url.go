package v2

import (
	"github.com/dannyota/greennode-community-sdk/v2/greennode/client"
)

func getAccessTokenUrl(psc client.ServiceClient) string {
	return psc.ServiceURL("auth", "token")
}
