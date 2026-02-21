package v2

import (
	"danny.vn/greennode/greennode/client"
)

func getAccessTokenURL(sc *client.ServiceClient) string {
	return sc.ServiceURL("auth", "token")
}
