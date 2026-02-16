package v1

import (
	"github.com/dannyota/greennode-community-sdk/v2/greennode/client"
)

type GetHostedZoneByIDRequest struct {
	HostedZoneID string
}

func NewGetHostedZoneByIDRequest(hostedZoneID string) *GetHostedZoneByIDRequest {
	return &GetHostedZoneByIDRequest{
		HostedZoneID: hostedZoneID,
	}
}

type ListHostedZonesRequest struct {
	Name string
}

func NewListHostedZonesRequest() *ListHostedZonesRequest {
	return &ListHostedZonesRequest{}
}

type CreateHostedZoneRequest struct {
	DomainName  string         `json:"domainName"`
	AssocVpcIDs []string       `json:"assocVpcIds"`
	Type        HostedZoneType `json:"type"`
	Description string         `json:"description"`
}

func (r *CreateHostedZoneRequest) ToRequestBody(sc *client.ServiceClient) map[string]any {
	return map[string]any{
		"domainName":  r.DomainName,
		"assocVpcIds": r.AssocVpcIDs,
		"type":        r.Type,
		"description": r.Description,
	}
}

func NewCreateHostedZoneRequest(domainName string, assocVpcIDs []string, zoneType HostedZoneType) *CreateHostedZoneRequest {
	return &CreateHostedZoneRequest{
		DomainName:  domainName,
		AssocVpcIDs: assocVpcIDs,
		Type:        zoneType,
	}
}

type DeleteHostedZoneRequest struct {
	HostedZoneID string
}

func NewDeleteHostedZoneRequest(hostedZoneID string) *DeleteHostedZoneRequest {
	return &DeleteHostedZoneRequest{
		HostedZoneID: hostedZoneID,
	}
}

type UpdateHostedZoneRequest struct {
	HostedZoneID string   `json:"-"`
	AssocVpcIDs  []string `json:"assocVpcIds"`
	Description  string   `json:"description"`
}

func (r *UpdateHostedZoneRequest) ToRequestBody(sc *client.ServiceClient) map[string]any {
	return map[string]any{
		"assocVpcIds": r.AssocVpcIDs,
		"description": r.Description,
	}
}

func NewUpdateHostedZoneRequest(hostedZoneID string) *UpdateHostedZoneRequest {
	return &UpdateHostedZoneRequest{
		HostedZoneID: hostedZoneID,
	}
}
