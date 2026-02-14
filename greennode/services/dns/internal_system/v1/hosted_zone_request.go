package v1

import (
	"github.com/dannyota/greennode-community-sdk/v2/greennode/client"
	"github.com/dannyota/greennode-community-sdk/v2/greennode/services/common"
)

type IGetHostedZoneByIDRequest interface {
	GetHostedZoneID() string
	WithHostedZoneID(hostedZoneID string) IGetHostedZoneByIDRequest
	ToMap() map[string]any

	ParseUserAgent() string
}

type GetHostedZoneByIDRequest struct {
	HostedZoneID string

	common.UserAgent
}

func (r *GetHostedZoneByIDRequest) GetHostedZoneID() string {
	return r.HostedZoneID
}

func (r *GetHostedZoneByIDRequest) WithHostedZoneID(hostedZoneID string) IGetHostedZoneByIDRequest {
	r.HostedZoneID = hostedZoneID
	return r
}

func (r *GetHostedZoneByIDRequest) ToMap() map[string]any {
	return map[string]any{
		"hostedZoneId": r.HostedZoneID,
	}
}

func NewGetHostedZoneByIDRequest(hostedZoneID string) IGetHostedZoneByIDRequest {
	return &GetHostedZoneByIDRequest{
		HostedZoneID: hostedZoneID,
	}
}


type IListHostedZonesRequest interface {
	GetName() string
	WithName(name string) IListHostedZonesRequest
	ToMap() map[string]any

	ParseUserAgent() string
}

type ListHostedZonesRequest struct {
	Name string

	common.UserAgent
}

func (r *ListHostedZonesRequest) GetName() string {
	return r.Name
}

func (r *ListHostedZonesRequest) WithName(name string) IListHostedZonesRequest {
	r.Name = name
	return r
}

func (r *ListHostedZonesRequest) ToMap() map[string]any {
	return map[string]any{
		"name": r.Name,
	}
}

func NewListHostedZonesRequest() IListHostedZonesRequest {
	return &ListHostedZonesRequest{}
}


type ICreateHostedZoneRequest interface {
	WithDomainName(domainName string) ICreateHostedZoneRequest
	WithAssocVpcIDs(assocVpcIDs []string) ICreateHostedZoneRequest
	WithType(zoneType HostedZoneType) ICreateHostedZoneRequest
	WithDescription(description string) ICreateHostedZoneRequest
	ToRequestBody(sc client.ServiceClient) map[string]any
	ToMap() map[string]any

	ParseUserAgent() string
}

type CreateHostedZoneRequest struct {
	DomainName  string         `json:"domainName"`
	AssocVpcIDs []string       `json:"assocVpcIds"`
	Type        HostedZoneType `json:"type"`
	Description string         `json:"description"`

	common.UserAgent
}

func (r *CreateHostedZoneRequest) WithDomainName(domainName string) ICreateHostedZoneRequest {
	r.DomainName = domainName
	return r
}

func (r *CreateHostedZoneRequest) WithAssocVpcIDs(assocVpcIDs []string) ICreateHostedZoneRequest {
	r.AssocVpcIDs = assocVpcIDs
	return r
}

func (r *CreateHostedZoneRequest) WithType(zoneType HostedZoneType) ICreateHostedZoneRequest {
	r.Type = zoneType
	return r
}

func (r *CreateHostedZoneRequest) WithDescription(description string) ICreateHostedZoneRequest {
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

func (r *CreateHostedZoneRequest) ToMap() map[string]any {
	return map[string]any{
		"domainName":  r.DomainName,
		"assocVpcIds": r.AssocVpcIDs,
		"type":        r.Type,
		"description": r.Description,
	}
}

func NewCreateHostedZoneRequest(domainName string, assocVpcIDs []string, zoneType HostedZoneType) ICreateHostedZoneRequest {
	return &CreateHostedZoneRequest{
		DomainName:  domainName,
		AssocVpcIDs: assocVpcIDs,
		Type:        zoneType,
	}
}


type IDeleteHostedZoneRequest interface {
	GetHostedZoneID() string
	WithHostedZoneID(hostedZoneID string) IDeleteHostedZoneRequest
	ToMap() map[string]any

	ParseUserAgent() string
}

type DeleteHostedZoneRequest struct {
	HostedZoneID string

	common.UserAgent
}

func (r *DeleteHostedZoneRequest) GetHostedZoneID() string {
	return r.HostedZoneID
}

func (r *DeleteHostedZoneRequest) WithHostedZoneID(hostedZoneID string) IDeleteHostedZoneRequest {
	r.HostedZoneID = hostedZoneID
	return r
}

func (r *DeleteHostedZoneRequest) ToMap() map[string]any {
	return map[string]any{
		"hostedZoneId": r.HostedZoneID,
	}
}

func NewDeleteHostedZoneRequest(hostedZoneID string) IDeleteHostedZoneRequest {
	return &DeleteHostedZoneRequest{
		HostedZoneID: hostedZoneID,
	}
}


type IUpdateHostedZoneRequest interface {
	GetHostedZoneID() string
	WithHostedZoneID(hostedZoneID string) IUpdateHostedZoneRequest
	WithAssocVpcIDs(assocVpcIDs []string) IUpdateHostedZoneRequest
	WithDescription(description string) IUpdateHostedZoneRequest
	ToRequestBody(sc client.ServiceClient) map[string]any
	ToMap() map[string]any

	ParseUserAgent() string
}

type UpdateHostedZoneRequest struct {
	HostedZoneID string   `json:"-"`
	AssocVpcIDs  []string `json:"assocVpcIds"`
	Description  string   `json:"description"`

	common.UserAgent
}

func (r *UpdateHostedZoneRequest) GetHostedZoneID() string {
	return r.HostedZoneID
}

func (r *UpdateHostedZoneRequest) WithHostedZoneID(hostedZoneID string) IUpdateHostedZoneRequest {
	r.HostedZoneID = hostedZoneID
	return r
}

func (r *UpdateHostedZoneRequest) WithAssocVpcIDs(assocVpcIDs []string) IUpdateHostedZoneRequest {
	r.AssocVpcIDs = assocVpcIDs
	return r
}

func (r *UpdateHostedZoneRequest) WithDescription(description string) IUpdateHostedZoneRequest {
	r.Description = description
	return r
}

func (r *UpdateHostedZoneRequest) ToRequestBody(sc client.ServiceClient) map[string]any {
	return map[string]any{
		"assocVpcIds": r.AssocVpcIDs,
		"description": r.Description,
	}
}

func (r *UpdateHostedZoneRequest) ToMap() map[string]any {
	return map[string]any{
		"hostedZoneId": r.HostedZoneID,
		"assocVpcIds":  r.AssocVpcIDs,
		"description":  r.Description,
	}
}

func NewUpdateHostedZoneRequest(hostedZoneID string) IUpdateHostedZoneRequest {
	return &UpdateHostedZoneRequest{
		HostedZoneID: hostedZoneID,
	}
}
