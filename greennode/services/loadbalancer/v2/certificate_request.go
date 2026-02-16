package v2

type ListCertificatesRequest struct{}


func NewListCertificatesRequest() *ListCertificatesRequest {
	return &ListCertificatesRequest{}
}


type GetCertificateByIDRequest struct {
	CertificateID string
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
	Name        string            `json:"name"`
	Type        ImportOptsTypeOpt `json:"type"`
	Certificate string            `json:"certificate"`

	CertificateChain *string `json:"certificateChain"`
	Passphrase       *string `json:"passphrase"`
	PrivateKey       *string `json:"privateKey"`
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

type DeleteCertificateByIDRequest struct {
	CertificateID string
}

func NewDeleteCertificateByIDRequest(certificateID string) *DeleteCertificateByIDRequest {
	return &DeleteCertificateByIDRequest{
		CertificateID: certificateID,
	}
}
