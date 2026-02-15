package portal

import (
	"github.com/dannyota/greennode-community-sdk/v2/greennode/client"
	"github.com/dannyota/greennode-community-sdk/v2/greennode/entity"
	portalv1 "github.com/dannyota/greennode-community-sdk/v2/greennode/services/portal/v1"
	portalv2 "github.com/dannyota/greennode-community-sdk/v2/greennode/services/portal/v2"
)

type PortalServiceV1 interface {
	ListZones() (*entity.ListZones, error)
	GetPortalInfo(opts *portalv1.GetPortalInfoRequest) (*entity.Portal, error)
	ListProjects(opts *portalv1.ListProjectsRequest) (*entity.ListPortals, error)
}

type PortalServiceV2 interface {
	ListAllQuotaUsed() (*entity.ListQuotas, error)
	GetQuotaByName(opts *portalv2.GetQuotaByNameRequest) (*entity.Quota, error)
}

func NewPortalServiceV1(svcClient client.ServiceClient) *portalv1.PortalServiceV1 {
	return &portalv1.PortalServiceV1{
		PortalClient: svcClient,
	}
}

func NewPortalServiceV2(svcClient client.ServiceClient) *portalv2.PortalServiceV2 {
	return &portalv2.PortalServiceV2{
		PortalClient: svcClient,
	}
}
