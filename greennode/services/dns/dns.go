package dns

import (
	"github.com/dannyota/greennode-community-sdk/v2/greennode/client"
	"github.com/dannyota/greennode-community-sdk/v2/greennode/entity"
	dnsinternalv1 "github.com/dannyota/greennode-community-sdk/v2/greennode/services/dns/internal_system/v1"
	dnsv1 "github.com/dannyota/greennode-community-sdk/v2/greennode/services/dns/v1"
)

type HostedZoneOps interface {
	GetHostedZoneByID(opts *dnsv1.GetHostedZoneByIDRequest) (*entity.HostedZone, error)
	ListHostedZones(opts *dnsv1.ListHostedZonesRequest) (*entity.ListHostedZone, error)
	CreateHostedZone(opts *dnsv1.CreateHostedZoneRequest) (*entity.HostedZone, error)
	UpdateHostedZone(opts *dnsv1.UpdateHostedZoneRequest) error
	DeleteHostedZone(opts *dnsv1.DeleteHostedZoneRequest) error
}

type RecordOps interface {
	ListRecords(opts *dnsv1.ListRecordsRequest) (*entity.ListDnsRecords, error)
	GetRecord(opts *dnsv1.GetRecordRequest) (*entity.DnsRecord, error)
	CreateDnsRecord(opts *dnsv1.CreateDnsRecordRequest) (*entity.DnsRecord, error)
	UpdateRecord(opts *dnsv1.UpdateRecordRequest) error
	DeleteRecord(opts *dnsv1.DeleteRecordRequest) error
}

type VDnsServiceV1 interface {
	HostedZoneOps
	RecordOps
}

type InternalHostedZoneOps interface {
	GetHostedZoneByID(opts *dnsinternalv1.GetHostedZoneByIDRequest, portalUserID string) (*entity.HostedZone, error)
	ListHostedZones(opts *dnsinternalv1.ListHostedZonesRequest, portalUserID string) (*entity.ListHostedZone, error)
	CreateHostedZone(opts *dnsinternalv1.CreateHostedZoneRequest, portalUserID string) (*entity.HostedZone, error)
	UpdateHostedZone(opts *dnsinternalv1.UpdateHostedZoneRequest, portalUserID string) error
	DeleteHostedZone(opts *dnsinternalv1.DeleteHostedZoneRequest, portalUserID string) error
}

type InternalRecordOps interface {
	ListRecords(opts *dnsinternalv1.ListRecordsRequest, portalUserID string) (*entity.ListDnsRecords, error)
	GetRecord(opts *dnsinternalv1.GetRecordRequest, portalUserID string) (*entity.DnsRecord, error)
	CreateDnsRecord(opts *dnsinternalv1.CreateDnsRecordRequest, portalUserID string) (*entity.DnsRecord, error)
	UpdateRecord(opts *dnsinternalv1.UpdateRecordRequest, portalUserID string) error
	DeleteRecord(opts *dnsinternalv1.DeleteRecordRequest, portalUserID string) error
}

type VDnsServiceInternal interface {
	InternalHostedZoneOps
	InternalRecordOps
}

func NewVDnsServiceV1(svcClient client.ServiceClient) *dnsv1.VDnsServiceV1 {
	return &dnsv1.VDnsServiceV1{
		DnsClient: svcClient,
	}
}

func NewVDnsServiceInternal(svcClient client.ServiceClient) *dnsinternalv1.VDnsServiceInternal {
	return &dnsinternalv1.VDnsServiceInternal{
		DnsClient: svcClient,
	}
}
