package v2

import (
	"strings"

	"github.com/dannyota/greennode-community-sdk/v2/greennode/entity"
)

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
	AllowedCidrs                string                         `json:"allowedCidrs"`
	ListenerName                string                         `json:"listenerName"`
	ListenerProtocol            ListenerProtocol               `json:"listenerProtocol"`
	ListenerProtocolPort        int                            `json:"listenerProtocolPort"`
	TimeoutClient               int                            `json:"timeoutClient"`
	TimeoutConnection           int                            `json:"timeoutConnection"`
	TimeoutMember               int                            `json:"timeoutMember"`
	DefaultPoolID               *string                        `json:"defaultPoolId"`
	CertificateAuthorities      *[]string                      `json:"certificateAuthorities"`
	ClientCertificate           *string                        `json:"clientCertificate"`
	DefaultCertificateAuthority *string                        `json:"defaultCertificateAuthority"`
	InsertHeaders               *[]entity.ListenerInsertHeader `json:"insertHeaders"`

	LoadBalancerID string
}

type UpdateListenerRequest struct {
	AllowedCidrs                string                         `json:"allowedCidrs"`
	DefaultPoolID               string                         `json:"defaultPoolId"`
	TimeoutClient               int                            `json:"timeoutClient"`
	TimeoutConnection           int                            `json:"timeoutConnection"`
	TimeoutMember               int                            `json:"timeoutMember"`
	CertificateAuthorities      *[]string                      `json:"certificateAuthorities"`
	ClientCertificate           *string                        `json:"clientCertificate"`
	DefaultCertificateAuthority *string                        `json:"defaultCertificateAuthority"`
	InsertHeaders               *[]entity.ListenerInsertHeader `json:"insertHeaders"`

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

func (r *CreateListenerRequest) WithAllowedCidrs(cidrs ...string) *CreateListenerRequest {
	if len(cidrs) < 1 {
		return r
	}

	r.AllowedCidrs = strings.Join(cidrs, ",")
	return r
}

func (r *CreateListenerRequest) WithTimeoutClient(toc int) *CreateListenerRequest {
	r.TimeoutClient = toc
	return r
}

func (r *CreateListenerRequest) WithTimeoutConnection(toc int) *CreateListenerRequest {
	r.TimeoutConnection = toc
	return r
}

func (r *CreateListenerRequest) WithTimeoutMember(tom int) *CreateListenerRequest {
	r.TimeoutMember = tom
	return r
}

func (r *CreateListenerRequest) AddCidrs(cidrs ...string) *CreateListenerRequest {
	if len(cidrs) < 1 {
		return r
	}

	if r.AllowedCidrs == "" {
		return r.WithAllowedCidrs(cidrs...)
	} else {
		r.AllowedCidrs = r.AllowedCidrs + "," + strings.Join(cidrs, ",")
	}

	return r
}

func (r *CreateListenerRequest) WithLoadBalancerID(lbid string) *CreateListenerRequest {
	r.LoadBalancerID = lbid
	return r
}

func (r *CreateListenerRequest) WithDefaultPoolID(poolID string) *CreateListenerRequest {
	r.DefaultPoolID = &poolID
	return r
}

func (r *CreateListenerRequest) WithCertificateAuthorities(ca *[]string) *CreateListenerRequest {
	r.CertificateAuthorities = ca
	return r
}

func (r *CreateListenerRequest) WithClientCertificate(clientCert *string) *CreateListenerRequest {
	r.ClientCertificate = clientCert
	return r
}

func (r *CreateListenerRequest) WithDefaultCertificateAuthority(defaultCA *string) *CreateListenerRequest {
	r.DefaultCertificateAuthority = defaultCA
	return r
}

func (r *CreateListenerRequest) WithInsertHeaders(pheaders ...string) *CreateListenerRequest {
	if len(pheaders) < 1 {
		r.InsertHeaders = nil
		return r
	}

	headers := make([]entity.ListenerInsertHeader, 0)
	for i := 0; i < len(pheaders); i += 2 {
		if i+1 >= len(pheaders) {
			break
		}
		headers = append(headers, entity.ListenerInsertHeader{
			HeaderName:  pheaders[i],
			HeaderValue: pheaders[i+1],
		})
	}
	r.InsertHeaders = &headers
	return r
}

func (r *UpdateListenerRequest) WithCidrs(cidrs ...string) *UpdateListenerRequest {
	if len(cidrs) < 1 {
		return r
	}

	r.AllowedCidrs = strings.Join(cidrs, ",")
	return r
}

func (r *UpdateListenerRequest) WithTimeoutClient(toc int) *UpdateListenerRequest {
	r.TimeoutClient = toc
	return r
}

func (r *UpdateListenerRequest) WithTimeoutConnection(toc int) *UpdateListenerRequest {
	r.TimeoutConnection = toc
	return r
}

func (r *UpdateListenerRequest) WithTimeoutMember(tom int) *UpdateListenerRequest {
	r.TimeoutMember = tom
	return r
}

func (r *UpdateListenerRequest) WithDefaultPoolID(poolID string) *UpdateListenerRequest {
	r.DefaultPoolID = poolID
	return r
}

func (r *UpdateListenerRequest) WithInsertHeaders(pheaders ...string) *UpdateListenerRequest {
	if len(pheaders) < 1 {
		r.InsertHeaders = nil
		return r
	}

	headers := make([]entity.ListenerInsertHeader, 0)
	for i := 0; i < len(pheaders); i += 2 {
		if i+1 >= len(pheaders) {
			break
		}
		headers = append(headers, entity.ListenerInsertHeader{
			HeaderName:  pheaders[i],
			HeaderValue: pheaders[i+1],
		})
	}
	r.InsertHeaders = &headers
	return r
}

func (r *UpdateListenerRequest) WithCertificateAuthorities(ca ...string) *UpdateListenerRequest {
	r.CertificateAuthorities = &ca
	return r
}

func (r *UpdateListenerRequest) WithClientCertificate(clientCert string) *UpdateListenerRequest {
	r.ClientCertificate = &clientCert
	return r
}

func (r *UpdateListenerRequest) WithDefaultCertificateAuthority(defaultCA string) *UpdateListenerRequest {
	r.DefaultCertificateAuthority = &defaultCA
	return r
}
