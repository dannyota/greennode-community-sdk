package v2

type CreateListenerResponse struct {
	UUID string `json:"uuid"`
}

type ListListenersByLoadBalancerIDResponse struct {
	Data []listenerResp `json:"data"`
}

type GetListenerByIDResponse struct {
	Data listenerResp `json:"data"`
}

type listenerResp struct {
	UUID                            string                 `json:"uuid"`
	Name                            string                 `json:"name"`
	Description                     string                 `json:"description,omitempty"`
	Protocol                        string                 `json:"protocol"`
	ProtocolPort                    int                    `json:"protocolPort"`
	ConnectionLimit                 int                    `json:"connectionLimit"`
	DefaultPoolID                   string                 `json:"defaultPoolId"`
	DefaultPoolName                 string                 `json:"defaultPoolName"`
	TimeoutClient                   int                    `json:"timeoutClient"`
	TimeoutMember                   int                    `json:"timeoutMember"`
	TimeoutConnection               int                    `json:"timeoutConnection"`
	AllowedCidrs                    string                 `json:"allowedCidrs"`
	CertificateAuthorities          []string               `json:"certificateAuthorities"`
	DisplayStatus                   string                 `json:"displayStatus"`
	CreatedAt                       string                 `json:"createdAt"`
	UpdatedAt                       string                 `json:"updatedAt"`
	DefaultCertificateAuthority     *string                `json:"defaultCertificateAuthority"`
	ClientCertificateAuthentication *string                `json:"clientCertificateAuthentication"`
	ProgressStatus                  string                 `json:"progressStatus"`
	InsertHeaders                   []ListenerInsertHeader `json:"insertHeaders"`
}

func (r *CreateListenerResponse) ToEntityListener() *Listener {
	return &Listener{
		UUID: r.UUID,
	}
}

func (r *ListListenersByLoadBalancerIDResponse) ToEntityListListeners() *ListListeners {
	listeners := &ListListeners{}
	for _, itemListener := range r.Data {
		listeners.Add(itemListener.toEntityListener())
	}
	return listeners
}

func (l *listenerResp) toEntityListener() *Listener {
	if l == nil {
		return nil
	}
	// Convert the slice of insertHeaderResponse to the slice of insertHeader
	insertHeaders := make([]ListenerInsertHeader, len(l.InsertHeaders))
	for i, header := range l.InsertHeaders {
		insertHeaders[i] = ListenerInsertHeader{
			HeaderName:  header.HeaderName,
			HeaderValue: header.HeaderValue,
		}
	}
	return &Listener{
		UUID:                            l.UUID,
		Name:                            l.Name,
		Description:                     l.Description,
		Protocol:                        l.Protocol,
		ProtocolPort:                    l.ProtocolPort,
		ConnectionLimit:                 l.ConnectionLimit,
		DefaultPoolID:                   l.DefaultPoolID,
		DefaultPoolName:                 l.DefaultPoolName,
		TimeoutClient:                   l.TimeoutClient,
		TimeoutMember:                   l.TimeoutMember,
		TimeoutConnection:               l.TimeoutConnection,
		AllowedCidrs:                    l.AllowedCidrs,
		CertificateAuthorities:          l.CertificateAuthorities,
		DisplayStatus:                   l.DisplayStatus,
		CreatedAt:                       l.CreatedAt,
		UpdatedAt:                       l.UpdatedAt,
		DefaultCertificateAuthority:     l.DefaultCertificateAuthority,
		ClientCertificateAuthentication: l.ClientCertificateAuthentication,
		ProgressStatus:                  l.ProgressStatus,
		InsertHeaders:                   insertHeaders,
	}
}

func (r *GetListenerByIDResponse) ToEntityListener() *Listener {
	return r.Data.toEntityListener()
}
