package v1

import (
	lsclient "github.com/dannyota/greennode-community-sdk/v2/greennode/client"
)

type ServerServiceInternalV1 struct {
	VServerClient lsclient.IServiceClient
}
