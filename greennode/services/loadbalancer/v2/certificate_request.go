package v2

import (
	"github.com/dannyota/greennode-community-sdk/v2/greennode/services/common"
)

type IListCertificatesRequest interface {
	ParseUserAgent() string
	AddUserAgent(agent ...string) IListCertificatesRequest
}

type IGetCertificateByIDRequest interface {
	GetCertificateID() string
	ParseUserAgent() string
	AddUserAgent(agent ...string) IGetCertificateByIDRequest
}

type ICreateCertificateRequest interface {
	ToRequestBody() any
	ParseUserAgent() string
	ToMap() map[string]any
	AddUserAgent(agent ...string) ICreateCertificateRequest

	WithCertificateChain(chain string) ICreateCertificateRequest
	WithPassphrase(passphrase string) ICreateCertificateRequest
	WithPrivateKey(privateKey string) ICreateCertificateRequest
}

type IDeleteCertificateByIDRequest interface {
	GetCertificateID() string
	ParseUserAgent() string
	AddUserAgent(agent ...string) IDeleteCertificateByIDRequest
}

var _ IListCertificatesRequest = &ListCertificatesRequest{}

type ListCertificatesRequest struct {
	common.UserAgent
}

func (r *ListCertificatesRequest) AddUserAgent(agent ...string) IListCertificatesRequest {
	r.UserAgent.AddUserAgent(agent...)
	return r
}

func NewListCertificatesRequest() *ListCertificatesRequest {
	return &ListCertificatesRequest{}
}


var _ IGetCertificateByIDRequest = &GetCertificateByIDRequest{}

type GetCertificateByIDRequest struct {
	common.UserAgent
	CertificateID string
}

func (r *GetCertificateByIDRequest) GetCertificateID() string {
	return r.CertificateID
}

func (r *GetCertificateByIDRequest) AddUserAgent(agent ...string) IGetCertificateByIDRequest {
	r.UserAgent.AddUserAgent(agent...)
	return r
}

func NewGetCertificateByIDRequest(certificateID string) *GetCertificateByIDRequest {
	return &GetCertificateByIDRequest{
		CertificateID: certificateID,
	}
}


type (
	ImportOptsTypeOpt string
)

const (
	ImportOptsTypeOptTLS ImportOptsTypeOpt = "TLS/SSL"
	ImportOptsTypeOptCA  ImportOptsTypeOpt = "CA"
)

var _ ICreateCertificateRequest = &CreateCertificateRequest{}

type CreateCertificateRequest struct {
	common.UserAgent
	Name        string            `json:"name"`
	Type        ImportOptsTypeOpt `json:"type"`
	Certificate string            `json:"certificate"`

	CertificateChain *string `json:"certificateChain"`
	Passphrase       *string `json:"passphrase"`
	PrivateKey       *string `json:"privateKey"`
}

func (r *CreateCertificateRequest) WithCertificateChain(chain string) ICreateCertificateRequest {
	r.CertificateChain = &chain
	return r
}

func (r *CreateCertificateRequest) WithPassphrase(passphrase string) ICreateCertificateRequest {
	r.Passphrase = &passphrase
	return r
}

func (r *CreateCertificateRequest) WithPrivateKey(privateKey string) ICreateCertificateRequest {
	r.PrivateKey = &privateKey
	return r
}

func (r *CreateCertificateRequest) ToRequestBody() any {
	return r
}

func (r *CreateCertificateRequest) ToMap() map[string]any {
	re := map[string]any{
		"name":        r.Name,
		"type":        r.Type,
		"certificate": r.Certificate,
	}
	if r.Type == ImportOptsTypeOptTLS {
		re["certificateChain"] = r.CertificateChain
		re["passphrase"] = r.Passphrase
		re["privateKey"] = r.PrivateKey
	}
	return re
}

func NewCreateCertificateRequest(name, cert string, typeVal ImportOptsTypeOpt) *CreateCertificateRequest {
	return &CreateCertificateRequest{
		Name:             name,
		Type:             typeVal,
		Certificate:      cert,
		CertificateChain: nil,
		Passphrase:       nil,
		PrivateKey:       nil,
	}
}

func (r *CreateCertificateRequest) AddUserAgent(agent ...string) ICreateCertificateRequest {
	r.UserAgent.AddUserAgent(agent...)
	return r
}


var _ IDeleteCertificateByIDRequest = &DeleteCertificateByIDRequest{}

type DeleteCertificateByIDRequest struct {
	common.UserAgent
	CertificateID string
}

func (r *DeleteCertificateByIDRequest) GetCertificateID() string {
	return r.CertificateID
}

func (r *DeleteCertificateByIDRequest) AddUserAgent(agent ...string) IDeleteCertificateByIDRequest {
	r.UserAgent.AddUserAgent(agent...)
	return r
}

func NewDeleteCertificateByIDRequest(certificateID string) *DeleteCertificateByIDRequest {
	return &DeleteCertificateByIDRequest{
		CertificateID: certificateID,
	}
}
