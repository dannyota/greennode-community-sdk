package v1

import (
	"github.com/dannyota/greennode-community-sdk/v2/greennode/client"
)

type GetHostedZoneByIDRequest struct {
	HostedZoneID string

}

func (r *GetHostedZoneByIDRequest) GetHostedZoneID() string {
	return r.HostedZoneID
}

func (r *GetHostedZoneByIDRequest) WithHostedZoneID(hostedZoneID string) *GetHostedZoneByIDRequest {
	r.HostedZoneID = hostedZoneID
	return r
}

func NewGetHostedZoneByIDRequest(hostedZoneID string) *GetHostedZoneByIDRequest {
	return &GetHostedZoneByIDRequest{
		HostedZoneID: hostedZoneID,
	}
}


type ListHostedZonesRequest struct {
	Name string

}

func (r *ListHostedZonesRequest) GetName() string {
	return r.Name
}

func (r *ListHostedZonesRequest) WithName(name string) *ListHostedZonesRequest {
	r.Name = name
	return r
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

func (r *CreateHostedZoneRequest) WithDomainName(domainName string) *CreateHostedZoneRequest {
	r.DomainName = domainName
	return r
}

func (r *CreateHostedZoneRequest) WithAssocVpcIDs(assocVpcIDs []string) *CreateHostedZoneRequest {
	r.AssocVpcIDs = assocVpcIDs
	return r
}

func (r *CreateHostedZoneRequest) WithType(zoneType HostedZoneType) *CreateHostedZoneRequest {
	r.Type = zoneType
	return r
}

func (r *CreateHostedZoneRequest) WithDescription(description string) *CreateHostedZoneRequest {
	r.Description = description
	return r
}

func (r *CreateHostedZoneRequest) ToRequestBody(sc client.ServiceClient) map[string]any {
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

func (r *DeleteHostedZoneRequest) GetHostedZoneID() string {
	return r.HostedZoneID
}

func (r *DeleteHostedZoneRequest) WithHostedZoneID(hostedZoneID string) *DeleteHostedZoneRequest {
	r.HostedZoneID = hostedZoneID
	return r
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

func (r *UpdateHostedZoneRequest) GetHostedZoneID() string {
	return r.HostedZoneID
}

func (r *UpdateHostedZoneRequest) WithHostedZoneID(hostedZoneID string) *UpdateHostedZoneRequest {
	r.HostedZoneID = hostedZoneID
	return r
}

func (r *UpdateHostedZoneRequest) WithAssocVpcIDs(assocVpcIDs []string) *UpdateHostedZoneRequest {
	r.AssocVpcIDs = assocVpcIDs
	return r
}

func (r *UpdateHostedZoneRequest) WithDescription(description string) *UpdateHostedZoneRequest {
	r.Description = description
	return r
}

func (r *UpdateHostedZoneRequest) ToRequestBody(sc client.ServiceClient) map[string]any {
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
