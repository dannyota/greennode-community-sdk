package v2

type certResponseData struct {
	UUID               string `json:"uuid"`
	Name               string `json:"name"`
	CertificateType    string `json:"certificateType"`
	ExpiredAt          string `json:"expiredAt"`
	ImportedAt         string `json:"importedAt"`
	NotAfter           int64  `json:"notAfter"`
	KeyAlgorithm       string `json:"keyAlgorithm"`
	Serial             string `json:"serial"`
	Subject            string `json:"subject"`
	DomainName         string `json:"domainName"`
	InUse              bool   `json:"inUse"`
	Issuer             string `json:"issuer"`
	SignatureAlgorithm string `json:"signatureAlgorithm"`
	NotBefore          int64  `json:"notBefore"`
	// SubjectAlternativeNames
}

func (r *certResponseData) ToEntityCertificate() *Certificate {
	return &Certificate{
		UUID:               r.UUID,
		Name:               r.Name,
		CertificateType:    r.CertificateType,
		ExpiredAt:          r.ExpiredAt,
		ImportedAt:         r.ImportedAt,
		NotAfter:           r.NotAfter,
		KeyAlgorithm:       r.KeyAlgorithm,
		Serial:             r.Serial,
		Subject:            r.Subject,
		DomainName:         r.DomainName,
		InUse:              r.InUse,
		Issuer:             r.Issuer,
		SignatureAlgorithm: r.SignatureAlgorithm,
		NotBefore:          r.NotBefore,
	}
}


type ListCertificatesResponse struct {
	ListData  []certResponseData `json:"listData"`
	Page      int                `json:"page"`
	PageSize  int                `json:"pageSize"`
	TotalPage int                `json:"totalPage"`
	TotalItem int                `json:"totalItem"`
}

func (r *ListCertificatesResponse) ToEntityListCertificates() *ListCertificates {
	certs := make([]Certificate, 0, len(r.ListData))
	for _, cert := range r.ListData {
		certs = append(certs, *cert.ToEntityCertificate())
	}

	return &ListCertificates{
		Certificates: certs,
	}
}


type GetCertificateByIDResponse struct {
	certResponseData
}


type CreateCertificateResponse struct {
	Data certResponseData `json:"data"`
}

func (r *CreateCertificateResponse) ToEntityCertificate() *Certificate {
	return r.Data.ToEntityCertificate()
}
