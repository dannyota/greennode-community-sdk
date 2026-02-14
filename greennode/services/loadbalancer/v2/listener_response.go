package v2

import "github.com/dannyota/greennode-community-sdk/v2/greennode/entity"

type CreateListenerResponse struct {
	UUID string `json:"uuid"`
}

type ListListenersByLoadBalancerIDResponse struct {
	Data []Listener `json:"data"`
}

type GetListenerByIDResponse struct {
	Data Listener `json:"data"`
}

type Listener struct {
	UUID                            string                        `json:"uuid"`
	Name                            string                        `json:"name"`
	Description                     string                        `json:"description,omitempty"`
	Protocol                        string                        `json:"protocol"`
	ProtocolPort                    int                           `json:"protocolPort"`
	ConnectionLimit                 int                           `json:"connectionLimit"`
	DefaultPoolID                   string                        `json:"defaultPoolId"`
	DefaultPoolName                 string                        `json:"defaultPoolName"`
	TimeoutClient                   int                           `json:"timeoutClient"`
	TimeoutMember                   int                           `json:"timeoutMember"`
	TimeoutConnection               int                           `json:"timeoutConnection"`
	AllowedCidrs                    string                        `json:"allowedCidrs"`
	CertificateAuthorities          []string                      `json:"certificateAuthorities"`
	DisplayStatus                   string                        `json:"displayStatus"`
	CreatedAt                       string                        `json:"createdAt"`
	UpdatedAt                       string                        `json:"updatedAt"`
	DefaultCertificateAuthority     *string                       `json:"defaultCertificateAuthority"`
	ClientCertificateAuthentication *string                       `json:"clientCertificateAuthentication"`
	ProgressStatus                  string                        `json:"progressStatus"`
	InsertHeaders                   []entity.ListenerInsertHeader `json:"insertHeaders"`
}

func (s *CreateListenerResponse) ToEntityListener() *entity.Listener {
	return &entity.Listener{
		UUID: s.UUID,
	}
}

func (s *ListListenersByLoadBalancerIDResponse) ToEntityListListeners() *entity.ListListeners {
	listeners := &entity.ListListeners{}
	for _, itemListener := range s.Data {
		listeners.Add(itemListener.toEntityListener())
	}
	return listeners
}

func (s *Listener) toEntityListener() *entity.Listener {
	if s == nil {
		return nil
	}
	// Convert the slice of insertHeaderResponse to the slice of insertHeader
	insertHeaders := make([]entity.ListenerInsertHeader, len(s.InsertHeaders))
	for i, header := range s.InsertHeaders {
		insertHeaders[i] = entity.ListenerInsertHeader{
			HeaderName:  header.HeaderName,
			HeaderValue: header.HeaderValue,
		}
	}
	return &entity.Listener{
		UUID:                            s.UUID,
		Name:                            s.Name,
		Description:                     s.Description,
		Protocol:                        s.Protocol,
		ProtocolPort:                    s.ProtocolPort,
		ConnectionLimit:                 s.ConnectionLimit,
		DefaultPoolID:                   s.DefaultPoolID,
		DefaultPoolName:                 s.DefaultPoolName,
		TimeoutClient:                   s.TimeoutClient,
		TimeoutMember:                   s.TimeoutMember,
		TimeoutConnection:               s.TimeoutConnection,
		AllowedCidrs:                    s.AllowedCidrs,
		CertificateAuthorities:          s.CertificateAuthorities,
		DisplayStatus:                   s.DisplayStatus,
		CreatedAt:                       s.CreatedAt,
		UpdatedAt:                       s.UpdatedAt,
		DefaultCertificateAuthority:     s.DefaultCertificateAuthority,
		ClientCertificateAuthentication: s.ClientCertificateAuthentication,
		ProgressStatus:                  s.ProgressStatus,
		InsertHeaders:                   insertHeaders,
	}
}

func (s *GetListenerByIDResponse) ToEntityListener() *entity.Listener {
	return s.Data.toEntityListener()
}
