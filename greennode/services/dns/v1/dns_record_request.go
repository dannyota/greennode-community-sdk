package v1

import (
	"github.com/dannyota/greennode-community-sdk/v2/greennode/client"
	"github.com/dannyota/greennode-community-sdk/v2/greennode/services/common"
)

type IListRecordsRequest interface {
	GetHostedZoneID() string
	GetName() string
	WithHostedZoneID(hostedZoneID string) IListRecordsRequest
	WithName(name string) IListRecordsRequest
	ToMap() map[string]any

	ParseUserAgent() string
}

type ListRecordsRequest struct {
	HostedZoneID string
	Name         string

	common.UserAgent
}

func (r *ListRecordsRequest) GetHostedZoneID() string {
	return r.HostedZoneID
}

func (r *ListRecordsRequest) GetName() string {
	return r.Name
}

func (r *ListRecordsRequest) WithHostedZoneID(hostedZoneID string) IListRecordsRequest {
	r.HostedZoneID = hostedZoneID
	return r
}

func (r *ListRecordsRequest) WithName(name string) IListRecordsRequest {
	r.Name = name
	return r
}

func (r *ListRecordsRequest) ToMap() map[string]any {
	return map[string]any{
		"hostedZoneId": r.HostedZoneID,
		"name":         r.Name,
	}
}

func NewListRecordsRequest(hostedZoneID string) IListRecordsRequest {
	return &ListRecordsRequest{
		HostedZoneID: hostedZoneID,
	}
}


type IGetRecordRequest interface {
	GetHostedZoneID() string
	GetRecordID() string
	WithHostedZoneID(hostedZoneID string) IGetRecordRequest
	WithRecordID(recordID string) IGetRecordRequest
	ToMap() map[string]any

	ParseUserAgent() string
}

type GetRecordRequest struct {
	HostedZoneID string
	RecordID     string

	common.UserAgent
}

func (r *GetRecordRequest) GetHostedZoneID() string {
	return r.HostedZoneID
}

func (r *GetRecordRequest) GetRecordID() string {
	return r.RecordID
}

func (r *GetRecordRequest) WithHostedZoneID(hostedZoneID string) IGetRecordRequest {
	r.HostedZoneID = hostedZoneID
	return r
}

func (r *GetRecordRequest) WithRecordID(recordID string) IGetRecordRequest {
	r.RecordID = recordID
	return r
}

func (r *GetRecordRequest) ToMap() map[string]any {
	return map[string]any{
		"hostedZoneId": r.HostedZoneID,
		"recordId":     r.RecordID,
	}
}

func NewGetRecordRequest(hostedZoneID, recordID string) IGetRecordRequest {
	return &GetRecordRequest{
		HostedZoneID: hostedZoneID,
		RecordID:     recordID,
	}
}


type IUpdateRecordRequest interface {
	GetHostedZoneID() string
	GetRecordID() string
	WithHostedZoneID(hostedZoneID string) IUpdateRecordRequest
	WithRecordID(recordID string) IUpdateRecordRequest
	WithSubDomain(subDomain string) IUpdateRecordRequest
	WithTTL(ttl int) IUpdateRecordRequest
	WithType(recordType DnsRecordType) IUpdateRecordRequest
	WithRoutingPolicy(routingPolicy RoutingPolicy) IUpdateRecordRequest
	WithEnableStickySession(enable bool) IUpdateRecordRequest
	WithValue(value []RecordValueRequest) IUpdateRecordRequest
	ToRequestBody(sc client.ServiceClient) map[string]any
	ToMap() map[string]any

	ParseUserAgent() string
}

type UpdateRecordRequest struct {
	HostedZoneID        string               `json:"-"`
	RecordID            string               `json:"-"`
	SubDomain           string               `json:"subDomain"`
	TTL                 int                  `json:"ttl"`
	Type                DnsRecordType        `json:"type"`
	RoutingPolicy       RoutingPolicy        `json:"routingPolicy"`
	EnableStickySession *bool                `json:"enableStickySession,omitempty"`
	Value               []RecordValueRequest `json:"value"`

	common.UserAgent
}

func (r *UpdateRecordRequest) GetHostedZoneID() string {
	return r.HostedZoneID
}

func (r *UpdateRecordRequest) GetRecordID() string {
	return r.RecordID
}

func (r *UpdateRecordRequest) WithHostedZoneID(hostedZoneID string) IUpdateRecordRequest {
	r.HostedZoneID = hostedZoneID
	return r
}

func (r *UpdateRecordRequest) WithRecordID(recordID string) IUpdateRecordRequest {
	r.RecordID = recordID
	return r
}

func (r *UpdateRecordRequest) WithSubDomain(subDomain string) IUpdateRecordRequest {
	r.SubDomain = subDomain
	return r
}

func (r *UpdateRecordRequest) WithTTL(ttl int) IUpdateRecordRequest {
	r.TTL = ttl
	return r
}

func (r *UpdateRecordRequest) WithType(recordType DnsRecordType) IUpdateRecordRequest {
	r.Type = recordType
	return r
}

func (r *UpdateRecordRequest) WithRoutingPolicy(routingPolicy RoutingPolicy) IUpdateRecordRequest {
	r.RoutingPolicy = routingPolicy
	return r
}

func (r *UpdateRecordRequest) WithEnableStickySession(enable bool) IUpdateRecordRequest {
	r.EnableStickySession = &enable
	return r
}

func (r *UpdateRecordRequest) WithValue(value []RecordValueRequest) IUpdateRecordRequest {
	r.Value = value
	return r
}

func (r *UpdateRecordRequest) ToRequestBody(sc client.ServiceClient) map[string]any {
	body := map[string]any{
		"subDomain":     r.SubDomain,
		"ttl":           r.TTL,
		"type":          r.Type,
		"routingPolicy": r.RoutingPolicy,
		"value":         r.Value,
	}
	if r.EnableStickySession != nil {
		body["enableStickySession"] = *r.EnableStickySession
	}
	return body
}

func (r *UpdateRecordRequest) ToMap() map[string]any {
	m := map[string]any{
		"hostedZoneId":  r.HostedZoneID,
		"recordId":      r.RecordID,
		"subDomain":     r.SubDomain,
		"ttl":           r.TTL,
		"type":          r.Type,
		"routingPolicy": r.RoutingPolicy,
		"value":         r.Value,
	}
	if r.EnableStickySession != nil {
		m["enableStickySession"] = *r.EnableStickySession
	}
	return m
}

func NewUpdateRecordRequest(hostedZoneID, recordID string) IUpdateRecordRequest {
	return &UpdateRecordRequest{
		HostedZoneID: hostedZoneID,
		RecordID:     recordID,
	}
}


type IDeleteRecordRequest interface {
	GetHostedZoneID() string
	GetRecordID() string
	WithHostedZoneID(hostedZoneID string) IDeleteRecordRequest
	WithRecordID(recordID string) IDeleteRecordRequest
	ToMap() map[string]any

	ParseUserAgent() string
}

type DeleteRecordRequest struct {
	HostedZoneID string
	RecordID     string

	common.UserAgent
}

func (r *DeleteRecordRequest) GetHostedZoneID() string {
	return r.HostedZoneID
}

func (r *DeleteRecordRequest) GetRecordID() string {
	return r.RecordID
}

func (r *DeleteRecordRequest) WithHostedZoneID(hostedZoneID string) IDeleteRecordRequest {
	r.HostedZoneID = hostedZoneID
	return r
}

func (r *DeleteRecordRequest) WithRecordID(recordID string) IDeleteRecordRequest {
	r.RecordID = recordID
	return r
}

func (r *DeleteRecordRequest) ToMap() map[string]any {
	return map[string]any{
		"hostedZoneId": r.HostedZoneID,
		"recordId":     r.RecordID,
	}
}

func NewDeleteRecordRequest(hostedZoneID, recordID string) IDeleteRecordRequest {
	return &DeleteRecordRequest{
		HostedZoneID: hostedZoneID,
		RecordID:     recordID,
	}
}


type ICreateDnsRecordRequest interface {
	GetHostedZoneID() string
	WithHostedZoneID(hostedZoneID string) ICreateDnsRecordRequest
	WithSubDomain(subDomain string) ICreateDnsRecordRequest
	WithTTL(ttl int) ICreateDnsRecordRequest
	WithType(recordType DnsRecordType) ICreateDnsRecordRequest
	WithRoutingPolicy(routingPolicy RoutingPolicy) ICreateDnsRecordRequest
	WithEnableStickySession(enable bool) ICreateDnsRecordRequest
	WithValue(value []RecordValueRequest) ICreateDnsRecordRequest
	ToRequestBody(sc client.ServiceClient) map[string]any
	ToMap() map[string]any

	ParseUserAgent() string
}

type RecordValueRequest struct {
	Value    string  `json:"value"`
	Location *string `json:"location,omitempty"`
	Weight   *int    `json:"weight,omitempty"`
}

type CreateDnsRecordRequest struct {
	HostedZoneID        string               `json:"-"`
	SubDomain           string               `json:"subDomain"`
	TTL                 int                  `json:"ttl"`
	Type                DnsRecordType        `json:"type"`
	RoutingPolicy       RoutingPolicy        `json:"routingPolicy"`
	EnableStickySession *bool                `json:"enableStickySession,omitempty"`
	Value               []RecordValueRequest `json:"value"`

	common.UserAgent
}

func (r *CreateDnsRecordRequest) GetHostedZoneID() string {
	return r.HostedZoneID
}

func (r *CreateDnsRecordRequest) WithHostedZoneID(hostedZoneID string) ICreateDnsRecordRequest {
	r.HostedZoneID = hostedZoneID
	return r
}

func (r *CreateDnsRecordRequest) WithSubDomain(subDomain string) ICreateDnsRecordRequest {
	r.SubDomain = subDomain
	return r
}

func (r *CreateDnsRecordRequest) WithTTL(ttl int) ICreateDnsRecordRequest {
	r.TTL = ttl
	return r
}

func (r *CreateDnsRecordRequest) WithType(recordType DnsRecordType) ICreateDnsRecordRequest {
	r.Type = recordType
	return r
}

func (r *CreateDnsRecordRequest) WithRoutingPolicy(routingPolicy RoutingPolicy) ICreateDnsRecordRequest {
	r.RoutingPolicy = routingPolicy
	return r
}

func (r *CreateDnsRecordRequest) WithEnableStickySession(enable bool) ICreateDnsRecordRequest {
	r.EnableStickySession = &enable
	return r
}

func (r *CreateDnsRecordRequest) WithValue(value []RecordValueRequest) ICreateDnsRecordRequest {
	r.Value = value
	return r
}

func (r *CreateDnsRecordRequest) ToRequestBody(sc client.ServiceClient) map[string]any {
	body := map[string]any{
		"subDomain":     r.SubDomain,
		"ttl":           r.TTL,
		"type":          r.Type,
		"routingPolicy": r.RoutingPolicy,
		"value":         r.Value,
	}
	if r.EnableStickySession != nil {
		body["enableStickySession"] = *r.EnableStickySession
	}
	return body
}

func (r *CreateDnsRecordRequest) ToMap() map[string]any {
	m := map[string]any{
		"hostedZoneId":  r.HostedZoneID,
		"subDomain":     r.SubDomain,
		"ttl":           r.TTL,
		"type":          r.Type,
		"routingPolicy": r.RoutingPolicy,
		"value":         r.Value,
	}
	if r.EnableStickySession != nil {
		m["enableStickySession"] = *r.EnableStickySession
	}
	return m
}

func NewCreateDnsRecordRequest(hostedZoneID, subDomain string, ttl int, recordType DnsRecordType, routingPolicy RoutingPolicy, value []RecordValueRequest) ICreateDnsRecordRequest {
	return &CreateDnsRecordRequest{
		HostedZoneID:  hostedZoneID,
		SubDomain:     subDomain,
		TTL:           ttl,
		Type:          recordType,
		RoutingPolicy: routingPolicy,
		Value:         value,
	}
}

// NewRecordValueRequest creates a record value with optional location and weight pointers
func NewRecordValueRequest(value string, location *string, weight *int) RecordValueRequest {
	return RecordValueRequest{
		Value:    value,
		Location: location,
		Weight:   weight,
	}
}
