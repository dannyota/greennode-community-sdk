package v1

import (
	"github.com/dannyota/greennode-community-sdk/v2/greennode/client"
)

type ListRecordsRequest struct {
	HostedZoneID string
	Name         string

}

func (r *ListRecordsRequest) GetHostedZoneID() string {
	return r.HostedZoneID
}

func (r *ListRecordsRequest) GetName() string {
	return r.Name
}

func (r *ListRecordsRequest) WithHostedZoneID(hostedZoneID string) *ListRecordsRequest {
	r.HostedZoneID = hostedZoneID
	return r
}

func (r *ListRecordsRequest) WithName(name string) *ListRecordsRequest {
	r.Name = name
	return r
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

func (r *GetRecordRequest) GetHostedZoneID() string {
	return r.HostedZoneID
}

func (r *GetRecordRequest) GetRecordID() string {
	return r.RecordID
}

func (r *GetRecordRequest) WithHostedZoneID(hostedZoneID string) *GetRecordRequest {
	r.HostedZoneID = hostedZoneID
	return r
}

func (r *GetRecordRequest) WithRecordID(recordID string) *GetRecordRequest {
	r.RecordID = recordID
	return r
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

func (r *UpdateRecordRequest) GetHostedZoneID() string {
	return r.HostedZoneID
}

func (r *UpdateRecordRequest) GetRecordID() string {
	return r.RecordID
}

func (r *UpdateRecordRequest) WithHostedZoneID(hostedZoneID string) *UpdateRecordRequest {
	r.HostedZoneID = hostedZoneID
	return r
}

func (r *UpdateRecordRequest) WithRecordID(recordID string) *UpdateRecordRequest {
	r.RecordID = recordID
	return r
}

func (r *UpdateRecordRequest) WithSubDomain(subDomain string) *UpdateRecordRequest {
	r.SubDomain = subDomain
	return r
}

func (r *UpdateRecordRequest) WithTTL(ttl int) *UpdateRecordRequest {
	r.TTL = ttl
	return r
}

func (r *UpdateRecordRequest) WithType(recordType DnsRecordType) *UpdateRecordRequest {
	r.Type = recordType
	return r
}

func (r *UpdateRecordRequest) WithRoutingPolicy(routingPolicy RoutingPolicy) *UpdateRecordRequest {
	r.RoutingPolicy = routingPolicy
	return r
}

func (r *UpdateRecordRequest) WithEnableStickySession(enable bool) *UpdateRecordRequest {
	r.EnableStickySession = &enable
	return r
}

func (r *UpdateRecordRequest) WithValue(value []RecordValueRequest) *UpdateRecordRequest {
	r.Value = value
	return r
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

func (r *DeleteRecordRequest) GetHostedZoneID() string {
	return r.HostedZoneID
}

func (r *DeleteRecordRequest) GetRecordID() string {
	return r.RecordID
}

func (r *DeleteRecordRequest) WithHostedZoneID(hostedZoneID string) *DeleteRecordRequest {
	r.HostedZoneID = hostedZoneID
	return r
}

func (r *DeleteRecordRequest) WithRecordID(recordID string) *DeleteRecordRequest {
	r.RecordID = recordID
	return r
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

func (r *CreateDnsRecordRequest) GetHostedZoneID() string {
	return r.HostedZoneID
}

func (r *CreateDnsRecordRequest) WithHostedZoneID(hostedZoneID string) *CreateDnsRecordRequest {
	r.HostedZoneID = hostedZoneID
	return r
}

func (r *CreateDnsRecordRequest) WithSubDomain(subDomain string) *CreateDnsRecordRequest {
	r.SubDomain = subDomain
	return r
}

func (r *CreateDnsRecordRequest) WithTTL(ttl int) *CreateDnsRecordRequest {
	r.TTL = ttl
	return r
}

func (r *CreateDnsRecordRequest) WithType(recordType DnsRecordType) *CreateDnsRecordRequest {
	r.Type = recordType
	return r
}

func (r *CreateDnsRecordRequest) WithRoutingPolicy(routingPolicy RoutingPolicy) *CreateDnsRecordRequest {
	r.RoutingPolicy = routingPolicy
	return r
}

func (r *CreateDnsRecordRequest) WithEnableStickySession(enable bool) *CreateDnsRecordRequest {
	r.EnableStickySession = &enable
	return r
}

func (r *CreateDnsRecordRequest) WithValue(value []RecordValueRequest) *CreateDnsRecordRequest {
	r.Value = value
	return r
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
