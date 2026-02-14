package portal

import (
	lsentity "github.com/dannyota/greennode-community-sdk/v2/greennode/entity"
	lserr "github.com/dannyota/greennode-community-sdk/v2/greennode/sdk_error"
	lsportalV1 "github.com/dannyota/greennode-community-sdk/v2/greennode/services/portal/v1"
	lsportalV2 "github.com/dannyota/greennode-community-sdk/v2/greennode/services/portal/v2"
)

type IPortalServiceV1 interface {
	ListZones() (*lsentity.ListZones, lserr.IError)
	GetPortalInfo(popts lsportalV1.IGetPortalInfoRequest) (*lsentity.Portal, lserr.IError)
	ListProjects(popts lsportalV1.IListProjectsRequest) (*lsentity.ListPortals, lserr.IError)
}

type IPortalServiceV2 interface {
	ListAllQuotaUsed() (*lsentity.ListQuotas, lserr.IError)
	GetQuotaByName(popts lsportalV2.IGetQuotaByNameRequest) (*lsentity.Quota, lserr.IError)
}
