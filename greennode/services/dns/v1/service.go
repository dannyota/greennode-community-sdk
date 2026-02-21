package v1

import "danny.vn/greennode/greennode/client"

type VDnsServiceV1 struct {
	Client *client.ServiceClient
}

type VDnsServiceInternal struct {
	Client *client.ServiceClient
}
