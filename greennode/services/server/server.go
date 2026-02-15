package server

import (
	"github.com/dannyota/greennode-community-sdk/v2/greennode/client"
	"github.com/dannyota/greennode-community-sdk/v2/greennode/entity"
	serverv1 "github.com/dannyota/greennode-community-sdk/v2/greennode/services/server/v1"
)

type ServerServiceInternalV1 interface {
	CreateSystemTags(opts *serverv1.CreateSystemTagRequest) (*[]entity.SystemTag, error)
}

func NewServerServiceInternalV1(svcClient client.ServiceClient) *serverv1.ServerServiceInternalV1 {
	return &serverv1.ServerServiceInternalV1{
		VServerClient: svcClient,
	}
}
