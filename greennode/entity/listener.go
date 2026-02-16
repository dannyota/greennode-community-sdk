package entity

type ListenerInsertHeader struct {
	HeaderName  string `json:"headerName"`
	HeaderValue string `json:"headerValue"`
}

type Listener struct {
	UUID                            string                 `json:"uuid"`
	Name                            string                 `json:"name"`
	Description                     string                 `json:"description"`
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

type ListListeners struct {
	Items []*Listener
}

func (ll *ListListeners) Add(listeners ...*Listener) {
	ll.Items = append(ll.Items, listeners...)
}

func (ll *ListListeners) Len() int {
	return len(ll.Items)
}

func (ll *ListListeners) Empty() bool {
	return ll.Len() < 1
}

func (ll *ListListeners) At(index int) *Listener {
	if index < 0 || index >= len(ll.Items) {
		return nil
	}

	return ll.Items[index]
}
