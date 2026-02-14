package v2

import (
	lsclient "github.com/dannyota/greennode-community-sdk/v2/greennode/client"
)

func getAccessTokenUrl(psc lsclient.IServiceClient) string {
	return psc.ServiceURL("auth", "token")
}
