package dns

import (
	"github.com/dannyota/greennode-community-sdk/v2/greennode/entity"
	sdkerror "github.com/dannyota/greennode-community-sdk/v2/greennode/sdkerror"
	dnsinternalv1 "github.com/dannyota/greennode-community-sdk/v2/greennode/services/dns/internal_system/v1"
	dnsv1 "github.com/dannyota/greennode-community-sdk/v2/greennode/services/dns/v1"
)

type VDnsServiceV1 interface {
	GetHostedZoneById(popts dnsv1.IGetHostedZoneByIdRequest) (*entity.HostedZone, sdkerror.Error)
	ListHostedZones(popts dnsv1.IListHostedZonesRequest) (*entity.ListHostedZone, sdkerror.Error)
	CreateHostedZone(popts dnsv1.ICreateHostedZoneRequest) (*entity.HostedZone, sdkerror.Error)
	UpdateHostedZone(popts dnsv1.IUpdateHostedZoneRequest) sdkerror.Error
	DeleteHostedZone(popts dnsv1.IDeleteHostedZoneRequest) sdkerror.Error

	ListRecords(popts dnsv1.IListRecordsRequest) (*entity.ListDnsRecords, sdkerror.Error)
	GetRecord(popts dnsv1.IGetRecordRequest) (*entity.DnsRecord, sdkerror.Error)
	CreateDnsRecord(popts dnsv1.ICreateDnsRecordRequest) (*entity.DnsRecord, sdkerror.Error)
	UpdateRecord(popts dnsv1.IUpdateRecordRequest) sdkerror.Error
	DeleteRecord(popts dnsv1.IDeleteRecordRequest) sdkerror.Error
}

type VDnsServiceInternal interface {
	GetHostedZoneById(popts dnsinternalv1.IGetHostedZoneByIdRequest, portalUserId string) (*entity.HostedZone, sdkerror.Error)
	ListHostedZones(popts dnsinternalv1.IListHostedZonesRequest, portalUserId string) (*entity.ListHostedZone, sdkerror.Error)
	CreateHostedZone(popts dnsinternalv1.ICreateHostedZoneRequest, portalUserId string) (*entity.HostedZone, sdkerror.Error)
	UpdateHostedZone(popts dnsinternalv1.IUpdateHostedZoneRequest, portalUserId string) sdkerror.Error
	DeleteHostedZone(popts dnsinternalv1.IDeleteHostedZoneRequest, portalUserId string) sdkerror.Error

	ListRecords(popts dnsinternalv1.IListRecordsRequest, portalUserId string) (*entity.ListDnsRecords, sdkerror.Error)
	GetRecord(popts dnsinternalv1.IGetRecordRequest, portalUserId string) (*entity.DnsRecord, sdkerror.Error)
	CreateDnsRecord(popts dnsinternalv1.ICreateDnsRecordRequest, portalUserId string) (*entity.DnsRecord, sdkerror.Error)
	UpdateRecord(popts dnsinternalv1.IUpdateRecordRequest, portalUserId string) sdkerror.Error
	DeleteRecord(popts dnsinternalv1.IDeleteRecordRequest, portalUserId string) sdkerror.Error
}
