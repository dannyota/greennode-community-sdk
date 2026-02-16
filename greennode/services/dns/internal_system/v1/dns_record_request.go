package v1

import (
	"github.com/dannyota/greennode-community-sdk/v2/greennode/client"
)

type ListRecordsRequest struct {
	HostedZoneID string
	Name         string
}

func NewListRecordsRequest(hostedZoneID string) *ListRecordsRequest {
	return &ListRecordsRequest{
		HostedZoneID: hostedZoneID,
	}
}

type GetRecordRequest struct {
	HostedZoneID string
	RecordID     string
}

func NewGetRecordRequest(hostedZoneID, recordID string) *GetRecordRequest {
	return &GetRecordRequest{
		HostedZoneID: hostedZoneID,
		RecordID:     recordID,
	}
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
}

func (r *UpdateRecordRequest) ToRequestBody(sc *client.ServiceClient) map[string]any {
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

func NewUpdateRecordRequest(hostedZoneID, recordID string) *UpdateRecordRequest {
	return &UpdateRecordRequest{
		HostedZoneID: hostedZoneID,
		RecordID:     recordID,
	}
}

type DeleteRecordRequest struct {
	HostedZoneID string
	RecordID     string
}

func NewDeleteRecordRequest(hostedZoneID, recordID string) *DeleteRecordRequest {
	return &DeleteRecordRequest{
		HostedZoneID: hostedZoneID,
		RecordID:     recordID,
	}
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
}

func (r *CreateDnsRecordRequest) ToRequestBody(sc *client.ServiceClient) map[string]any {
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

func NewCreateDnsRecordRequest(hostedZoneID, subDomain string, ttl int, recordType DnsRecordType, routingPolicy RoutingPolicy, value []RecordValueRequest) *CreateDnsRecordRequest {
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
