package v2

import (
	"github.com/dannyota/greennode-community-sdk/v2/greennode/client"
)

func getAccessTokenUrl(psc client.IServiceClient) string {
	return psc.ServiceURL("auth", "token")
}
