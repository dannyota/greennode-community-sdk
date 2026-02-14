package dns

import (
	"github.com/dannyota/greennode-community-sdk/v2/greennode/entity"
	sdkerror "github.com/dannyota/greennode-community-sdk/v2/greennode/sdkerror"
	dnsinternalv1 "github.com/dannyota/greennode-community-sdk/v2/greennode/services/dns/internal_system/v1"
	dnsv1 "github.com/dannyota/greennode-community-sdk/v2/greennode/services/dns/v1"
)

type VDnsServiceV1 interface {
	GetHostedZoneByID(opts dnsv1.IGetHostedZoneByIDRequest) (*entity.HostedZone, sdkerror.Error)
	ListHostedZones(opts dnsv1.IListHostedZonesRequest) (*entity.ListHostedZone, sdkerror.Error)
	CreateHostedZone(opts dnsv1.ICreateHostedZoneRequest) (*entity.HostedZone, sdkerror.Error)
	UpdateHostedZone(opts dnsv1.IUpdateHostedZoneRequest) sdkerror.Error
	DeleteHostedZone(opts dnsv1.IDeleteHostedZoneRequest) sdkerror.Error

	ListRecords(opts dnsv1.IListRecordsRequest) (*entity.ListDnsRecords, sdkerror.Error)
	GetRecord(opts dnsv1.IGetRecordRequest) (*entity.DnsRecord, sdkerror.Error)
	CreateDnsRecord(opts dnsv1.ICreateDnsRecordRequest) (*entity.DnsRecord, sdkerror.Error)
	UpdateRecord(opts dnsv1.IUpdateRecordRequest) sdkerror.Error
	DeleteRecord(opts dnsv1.IDeleteRecordRequest) sdkerror.Error
}

type VDnsServiceInternal interface {
	GetHostedZoneByID(opts dnsinternalv1.IGetHostedZoneByIDRequest, portalUserID string) (*entity.HostedZone, sdkerror.Error)
	ListHostedZones(opts dnsinternalv1.IListHostedZonesRequest, portalUserID string) (*entity.ListHostedZone, sdkerror.Error)
	CreateHostedZone(opts dnsinternalv1.ICreateHostedZoneRequest, portalUserID string) (*entity.HostedZone, sdkerror.Error)
	UpdateHostedZone(opts dnsinternalv1.IUpdateHostedZoneRequest, portalUserID string) sdkerror.Error
	DeleteHostedZone(opts dnsinternalv1.IDeleteHostedZoneRequest, portalUserID string) sdkerror.Error

	ListRecords(opts dnsinternalv1.IListRecordsRequest, portalUserID string) (*entity.ListDnsRecords, sdkerror.Error)
	GetRecord(opts dnsinternalv1.IGetRecordRequest, portalUserID string) (*entity.DnsRecord, sdkerror.Error)
	CreateDnsRecord(opts dnsinternalv1.ICreateDnsRecordRequest, portalUserID string) (*entity.DnsRecord, sdkerror.Error)
	UpdateRecord(opts dnsinternalv1.IUpdateRecordRequest, portalUserID string) sdkerror.Error
	DeleteRecord(opts dnsinternalv1.IDeleteRecordRequest, portalUserID string) sdkerror.Error
}
