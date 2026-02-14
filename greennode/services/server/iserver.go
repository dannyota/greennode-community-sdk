package server

import (
	lsentity "github.com/dannyota/greennode-community-sdk/v2/greennode/entity"
	lserr "github.com/dannyota/greennode-community-sdk/v2/greennode/sdk_error"
	lsnserverSvcV1 "github.com/dannyota/greennode-community-sdk/v2/greennode/services/server/v1"
)

type IServerServiceInternalV1 interface {
	CreateSystemTags(popts lsnserverSvcV1.ICreateSystemTagRequest) (*[]lsentity.SystemTag, lserr.IError)
}
