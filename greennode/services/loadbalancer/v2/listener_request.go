package v2

const (
	ListenerProtocolTCP   ListenerProtocol = "TCP"
	ListenerProtocolUDP   ListenerProtocol = "UDP"
	ListenerProtocolHTTP  ListenerProtocol = "HTTP"
	ListenerProtocolHTTPS ListenerProtocol = "HTTPS"
)

func NewCreateListenerRequest(name string, protocol ListenerProtocol, port int) *CreateListenerRequest {
	return &CreateListenerRequest{
		ListenerName:        name,
		ListenerProtocol:    protocol,
		ListenerProtocolPort: port,
		AllowedCidrs:        "0.0.0.0/0",
		TimeoutClient:       50,
		TimeoutMember:       50,
		TimeoutConnection:   5,
	}
}

func NewUpdateListenerRequest(lbID, listenerID string) *UpdateListenerRequest {
	return &UpdateListenerRequest{
		LoadBalancerID: lbID,
		ListenerID:     listenerID,
	}
}

func NewListListenersByLoadBalancerIDRequest(lbID string) *ListListenersByLoadBalancerIDRequest {
	return &ListListenersByLoadBalancerIDRequest{
		LoadBalancerID: lbID,
	}
}

func NewDeleteListenerByIDRequest(lbID, listenerID string) *DeleteListenerByIDRequest {
	return &DeleteListenerByIDRequest{
		LoadBalancerID: lbID,
		ListenerID:     listenerID,
	}
}

func NewGetListenerByIDRequest(lbID, listenerID string) *GetListenerByIDRequest {
	return &GetListenerByIDRequest{
		LoadBalancerID: lbID,
		ListenerID:     listenerID,
	}
}

type ListenerProtocol string

type CreateListenerRequest struct {
	AllowedCidrs                string                  `json:"allowedCidrs"`
	ListenerName                string                  `json:"listenerName"`
	ListenerProtocol            ListenerProtocol        `json:"listenerProtocol"`
	ListenerProtocolPort        int                     `json:"listenerProtocolPort"`
	TimeoutClient               int                     `json:"timeoutClient"`
	TimeoutConnection           int                     `json:"timeoutConnection"`
	TimeoutMember               int                     `json:"timeoutMember"`
	DefaultPoolID               *string                 `json:"defaultPoolId"`
	CertificateAuthorities      *[]string               `json:"certificateAuthorities"`
	ClientCertificate           *string                 `json:"clientCertificate"`
	DefaultCertificateAuthority *string                 `json:"defaultCertificateAuthority"`
	InsertHeaders               *[]ListenerInsertHeader `json:"insertHeaders"`

	LoadBalancerID string
}

type UpdateListenerRequest struct {
	AllowedCidrs                string                  `json:"allowedCidrs"`
	DefaultPoolID               string                  `json:"defaultPoolId"`
	TimeoutClient               int                     `json:"timeoutClient"`
	TimeoutConnection           int                     `json:"timeoutConnection"`
	TimeoutMember               int                     `json:"timeoutMember"`
	CertificateAuthorities      *[]string               `json:"certificateAuthorities"`
	ClientCertificate           *string                 `json:"clientCertificate"`
	DefaultCertificateAuthority *string                 `json:"defaultCertificateAuthority"`
	InsertHeaders               *[]ListenerInsertHeader `json:"insertHeaders"`

	LoadBalancerID string
	ListenerID     string
}

type ListListenersByLoadBalancerIDRequest struct {
	LoadBalancerID string
}

type DeleteListenerByIDRequest struct {
	LoadBalancerID string
	ListenerID     string
}

type GetListenerByIDRequest struct {
	LoadBalancerID string
	ListenerID     string
}

// normalizeForAPI clears certificate fields when the listener protocol is not
// HTTPS, since they are only relevant for TLS termination. This mutates the receiver.
func (r *CreateListenerRequest) normalizeForAPI() {
	if r == nil {
		return
	}

	if r.ListenerProtocol == ListenerProtocolHTTPS {
		return
	}

	r.CertificateAuthorities = nil
	r.ClientCertificate = nil
	r.DefaultCertificateAuthority = nil
}
