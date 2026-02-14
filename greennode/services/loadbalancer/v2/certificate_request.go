package v2

import (
	"github.com/dannyota/greennode-community-sdk/v2/greennode/services/common"
)

var _ IListCertificatesRequest = &ListCertificatesRequest{}

type ListCertificatesRequest struct {
	common.UserAgent
}

func (s *ListCertificatesRequest) AddUserAgent(agent ...string) IListCertificatesRequest {
	s.UserAgent.AddUserAgent(agent...)
	return s
}

func NewListCertificatesRequest() *ListCertificatesRequest {
	return &ListCertificatesRequest{}
}

// --------------------------------------------------------

var _ IGetCertificateByIdRequest = &GetCertificateByIdRequest{}

type GetCertificateByIdRequest struct {
	common.UserAgent
	CertificateId string
}

func (r *GetCertificateByIdRequest) GetCertificateId() string {
	return r.CertificateId
}

func (s *GetCertificateByIdRequest) AddUserAgent(agent ...string) IGetCertificateByIdRequest {
	s.UserAgent.AddUserAgent(agent...)
	return s
}

func NewGetCertificateByIdRequest(certificateId string) *GetCertificateByIdRequest {
	return &GetCertificateByIdRequest{
		CertificateId: certificateId,
	}
}

// --------------------------------------------------------

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

func (r *CreateCertificateRequest) ToRequestBody() interface{} {
	return r
}

func (r *CreateCertificateRequest) ToMap() map[string]interface{} {
	re := map[string]interface{}{
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

func NewCreateCertificateRequest(name, cert string, typeVal ImportOptsTypeOpt) ICreateCertificateRequest {
	return &CreateCertificateRequest{
		Name:             name,
		Type:             typeVal,
		Certificate:      cert,
		CertificateChain: nil,
		Passphrase:       nil,
		PrivateKey:       nil,
	}
}

func (s *CreateCertificateRequest) AddUserAgent(agent ...string) ICreateCertificateRequest {
	s.UserAgent.AddUserAgent(agent...)
	return s
}

// --------------------------------------------------------

var _ IDeleteCertificateByIdRequest = &DeleteCertificateByIdRequest{}

type DeleteCertificateByIdRequest struct {
	common.UserAgent
	CertificateId string
}

func (r *DeleteCertificateByIdRequest) GetCertificateId() string {
	return r.CertificateId
}

func (s *DeleteCertificateByIdRequest) AddUserAgent(agent ...string) IDeleteCertificateByIdRequest {
	s.UserAgent.AddUserAgent(agent...)
	return s
}

func NewDeleteCertificateByIdRequest(certificateId string) *DeleteCertificateByIdRequest {
	return &DeleteCertificateByIdRequest{
		CertificateId: certificateId,
	}
}
