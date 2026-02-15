package dns

import (
	"github.com/dannyota/greennode-community-sdk/v2/greennode/client"
	"github.com/dannyota/greennode-community-sdk/v2/greennode/entity"
	sdkerror "github.com/dannyota/greennode-community-sdk/v2/greennode/sdkerror"
	dnsinternalv1 "github.com/dannyota/greennode-community-sdk/v2/greennode/services/dns/internal_system/v1"
	dnsv1 "github.com/dannyota/greennode-community-sdk/v2/greennode/services/dns/v1"
)

type VDnsServiceV1 interface {
	GetHostedZoneByID(opts *dnsv1.GetHostedZoneByIDRequest) (*entity.HostedZone, sdkerror.Error)
	ListHostedZones(opts *dnsv1.ListHostedZonesRequest) (*entity.ListHostedZone, sdkerror.Error)
	CreateHostedZone(opts *dnsv1.CreateHostedZoneRequest) (*entity.HostedZone, sdkerror.Error)
	UpdateHostedZone(opts *dnsv1.UpdateHostedZoneRequest) sdkerror.Error
	DeleteHostedZone(opts *dnsv1.DeleteHostedZoneRequest) sdkerror.Error

	ListRecords(opts *dnsv1.ListRecordsRequest) (*entity.ListDnsRecords, sdkerror.Error)
	GetRecord(opts *dnsv1.GetRecordRequest) (*entity.DnsRecord, sdkerror.Error)
	CreateDnsRecord(opts *dnsv1.CreateDnsRecordRequest) (*entity.DnsRecord, sdkerror.Error)
	UpdateRecord(opts *dnsv1.UpdateRecordRequest) sdkerror.Error
	DeleteRecord(opts *dnsv1.DeleteRecordRequest) sdkerror.Error
}

type VDnsServiceInternal interface {
	GetHostedZoneByID(opts *dnsinternalv1.GetHostedZoneByIDRequest, portalUserID string) (*entity.HostedZone, sdkerror.Error)
	ListHostedZones(opts *dnsinternalv1.ListHostedZonesRequest, portalUserID string) (*entity.ListHostedZone, sdkerror.Error)
	CreateHostedZone(opts *dnsinternalv1.CreateHostedZoneRequest, portalUserID string) (*entity.HostedZone, sdkerror.Error)
	UpdateHostedZone(opts *dnsinternalv1.UpdateHostedZoneRequest, portalUserID string) sdkerror.Error
	DeleteHostedZone(opts *dnsinternalv1.DeleteHostedZoneRequest, portalUserID string) sdkerror.Error

	ListRecords(opts *dnsinternalv1.ListRecordsRequest, portalUserID string) (*entity.ListDnsRecords, sdkerror.Error)
	GetRecord(opts *dnsinternalv1.GetRecordRequest, portalUserID string) (*entity.DnsRecord, sdkerror.Error)
	CreateDnsRecord(opts *dnsinternalv1.CreateDnsRecordRequest, portalUserID string) (*entity.DnsRecord, sdkerror.Error)
	UpdateRecord(opts *dnsinternalv1.UpdateRecordRequest, portalUserID string) sdkerror.Error
	DeleteRecord(opts *dnsinternalv1.DeleteRecordRequest, portalUserID string) sdkerror.Error
}

func NewVDnsServiceV1(svcClient client.ServiceClient) VDnsServiceV1 {
	return &dnsv1.VDnsServiceV1{
		DnsClient: svcClient,
	}
}

func NewVDnsServiceInternal(svcClient client.ServiceClient) VDnsServiceInternal {
	return &dnsinternalv1.VDnsServiceInternal{
		DnsClient: svcClient,
	}
}
