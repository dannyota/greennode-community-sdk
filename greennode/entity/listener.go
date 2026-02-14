package entity

type ListenerInsertHeader struct {
	HeaderName  string `json:"headerName"`
	HeaderValue string `json:"headerValue"`
}

type Listener struct {
	UUID                            string
	Name                            string
	Description                     string
	Protocol                        string
	ProtocolPort                    int
	ConnectionLimit                 int
	DefaultPoolID                   string
	DefaultPoolName                 string
	TimeoutClient                   int
	TimeoutMember                   int
	TimeoutConnection               int
	AllowedCidrs                    string
	CertificateAuthorities          []string
	DisplayStatus                   string
	CreatedAt                       string
	UpdatedAt                       string
	DefaultCertificateAuthority     *string
	ClientCertificateAuthentication *string
	ProgressStatus                  string
	InsertHeaders                   []ListenerInsertHeader
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

func (l *Listener) GetID() string {
	return l.UUID
}

func (l *Listener) GetDefaultPoolID() string {
	return l.DefaultPoolID
}

func (ll *ListListeners) At(index int) *Listener {
	if index < 0 || index >= len(ll.Items) {
		return nil
	}

	return ll.Items[index]
}
