package server

import (
	"github.com/dannyota/greennode-community-sdk/v2/greennode/entity"
	sdkerror "github.com/dannyota/greennode-community-sdk/v2/greennode/sdkerror"
	serverv1 "github.com/dannyota/greennode-community-sdk/v2/greennode/services/server/v1"
)

type ServerServiceInternalV1 interface {
	CreateSystemTags(popts serverv1.ICreateSystemTagRequest) (*[]entity.SystemTag, sdkerror.Error)
}
