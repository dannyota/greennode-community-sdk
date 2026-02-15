package v2

import (
	"github.com/dannyota/greennode-community-sdk/v2/greennode/services/common"
)

type ListCertificatesRequest struct {
	common.UserAgent
}

func (r *ListCertificatesRequest) AddUserAgent(agent ...string) *ListCertificatesRequest {
	r.UserAgent.AddUserAgent(agent...)
	return r
}

func NewListCertificatesRequest() *ListCertificatesRequest {
	return &ListCertificatesRequest{}
}


type GetCertificateByIDRequest struct {
	common.UserAgent
	CertificateID string
}

func (r *GetCertificateByIDRequest) GetCertificateID() string {
	return r.CertificateID
}

func (r *GetCertificateByIDRequest) AddUserAgent(agent ...string) *GetCertificateByIDRequest {
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

type CreateCertificateRequest struct {
	common.UserAgent
	Name        string            `json:"name"`
	Type        ImportOptsTypeOpt `json:"type"`
	Certificate string            `json:"certificate"`

	CertificateChain *string `json:"certificateChain"`
	Passphrase       *string `json:"passphrase"`
	PrivateKey       *string `json:"privateKey"`
}

func (r *CreateCertificateRequest) WithCertificateChain(chain string) *CreateCertificateRequest {
	r.CertificateChain = &chain
	return r
}

func (r *CreateCertificateRequest) WithPassphrase(passphrase string) *CreateCertificateRequest {
	r.Passphrase = &passphrase
	return r
}

func (r *CreateCertificateRequest) WithPrivateKey(privateKey string) *CreateCertificateRequest {
	r.PrivateKey = &privateKey
	return r
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

func (r *CreateCertificateRequest) AddUserAgent(agent ...string) *CreateCertificateRequest {
	r.UserAgent.AddUserAgent(agent...)
	return r
}


type DeleteCertificateByIDRequest struct {
	common.UserAgent
	CertificateID string
}

func (r *DeleteCertificateByIDRequest) GetCertificateID() string {
	return r.CertificateID
}

func (r *DeleteCertificateByIDRequest) AddUserAgent(agent ...string) *DeleteCertificateByIDRequest {
	r.UserAgent.AddUserAgent(agent...)
	return r
}

func NewDeleteCertificateByIDRequest(certificateID string) *DeleteCertificateByIDRequest {
	return &DeleteCertificateByIDRequest{
		CertificateID: certificateID,
	}
}
