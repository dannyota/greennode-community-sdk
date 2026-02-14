package portal

import (
	"github.com/dannyota/greennode-community-sdk/v2/greennode/entity"
	sdkerror "github.com/dannyota/greennode-community-sdk/v2/greennode/sdkerror"
	portalv1 "github.com/dannyota/greennode-community-sdk/v2/greennode/services/portal/v1"
	portalv2 "github.com/dannyota/greennode-community-sdk/v2/greennode/services/portal/v2"
)

type IPortalServiceV1 interface {
	ListZones() (*entity.ListZones, sdkerror.IError)
	GetPortalInfo(popts portalv1.IGetPortalInfoRequest) (*entity.Portal, sdkerror.IError)
	ListProjects(popts portalv1.IListProjectsRequest) (*entity.ListPortals, sdkerror.IError)
}

type IPortalServiceV2 interface {
	ListAllQuotaUsed() (*entity.ListQuotas, sdkerror.IError)
	GetQuotaByName(popts portalv2.IGetQuotaByNameRequest) (*entity.Quota, sdkerror.IError)
}
