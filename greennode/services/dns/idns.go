package dns

import (
	"github.com/dannyota/greennode-community-sdk/v2/greennode/entity"
	sdkerror "github.com/dannyota/greennode-community-sdk/v2/greennode/sdk_error"
	dnsinternalv1 "github.com/dannyota/greennode-community-sdk/v2/greennode/services/dns/internal_system/v1"
	dnsv1 "github.com/dannyota/greennode-community-sdk/v2/greennode/services/dns/v1"
)

type IVDnsServiceV1 interface {
	GetHostedZoneById(popts dnsv1.IGetHostedZoneByIdRequest) (*entity.HostedZone, sdkerror.IError)
	ListHostedZones(popts dnsv1.IListHostedZonesRequest) (*entity.ListHostedZone, sdkerror.IError)
	CreateHostedZone(popts dnsv1.ICreateHostedZoneRequest) (*entity.HostedZone, sdkerror.IError)
	UpdateHostedZone(popts dnsv1.IUpdateHostedZoneRequest) sdkerror.IError
	DeleteHostedZone(popts dnsv1.IDeleteHostedZoneRequest) sdkerror.IError

	ListRecords(popts dnsv1.IListRecordsRequest) (*entity.ListDnsRecords, sdkerror.IError)
	GetRecord(popts dnsv1.IGetRecordRequest) (*entity.DnsRecord, sdkerror.IError)
	CreateDnsRecord(popts dnsv1.ICreateDnsRecordRequest) (*entity.DnsRecord, sdkerror.IError)
	UpdateRecord(popts dnsv1.IUpdateRecordRequest) sdkerror.IError
	DeleteRecord(popts dnsv1.IDeleteRecordRequest) sdkerror.IError
}

type IVDnsServiceInternal interface {
	GetHostedZoneById(popts dnsinternalv1.IGetHostedZoneByIdRequest, portalUserId string) (*entity.HostedZone, sdkerror.IError)
	ListHostedZones(popts dnsinternalv1.IListHostedZonesRequest, portalUserId string) (*entity.ListHostedZone, sdkerror.IError)
	CreateHostedZone(popts dnsinternalv1.ICreateHostedZoneRequest, portalUserId string) (*entity.HostedZone, sdkerror.IError)
	UpdateHostedZone(popts dnsinternalv1.IUpdateHostedZoneRequest, portalUserId string) sdkerror.IError
	DeleteHostedZone(popts dnsinternalv1.IDeleteHostedZoneRequest, portalUserId string) sdkerror.IError

	ListRecords(popts dnsinternalv1.IListRecordsRequest, portalUserId string) (*entity.ListDnsRecords, sdkerror.IError)
	GetRecord(popts dnsinternalv1.IGetRecordRequest, portalUserId string) (*entity.DnsRecord, sdkerror.IError)
	CreateDnsRecord(popts dnsinternalv1.ICreateDnsRecordRequest, portalUserId string) (*entity.DnsRecord, sdkerror.IError)
	UpdateRecord(popts dnsinternalv1.IUpdateRecordRequest, portalUserId string) sdkerror.IError
	DeleteRecord(popts dnsinternalv1.IDeleteRecordRequest, portalUserId string) sdkerror.IError
}
