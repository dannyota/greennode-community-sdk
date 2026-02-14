package portal

import (
	"github.com/dannyota/greennode-community-sdk/v2/greennode/entity"
	sdkerror "github.com/dannyota/greennode-community-sdk/v2/greennode/sdkerror"
	portalv1 "github.com/dannyota/greennode-community-sdk/v2/greennode/services/portal/v1"
	portalv2 "github.com/dannyota/greennode-community-sdk/v2/greennode/services/portal/v2"
)

type PortalServiceV1 interface {
	ListZones() (*entity.ListZones, sdkerror.Error)
	GetPortalInfo(opts portalv1.IGetPortalInfoRequest) (*entity.Portal, sdkerror.Error)
	ListProjects(opts portalv1.IListProjectsRequest) (*entity.ListPortals, sdkerror.Error)
}

type PortalServiceV2 interface {
	ListAllQuotaUsed() (*entity.ListQuotas, sdkerror.Error)
	GetQuotaByName(opts portalv2.IGetQuotaByNameRequest) (*entity.Quota, sdkerror.Error)
}
