package inter

import (
	"strings"

	"github.com/dannyota/greennode-community-sdk/v2/greennode/services/common"
)

const (
	ListenerProtocolTCP   ListenerProtocol = "TCP"
	ListenerProtocolUDP   ListenerProtocol = "UDP"
	ListenerProtocolHTTP  ListenerProtocol = "HTTP"
	ListenerProtocolHTTPS ListenerProtocol = "HTTPS"
)

func NewCreateListenerRequest(name string, protocol ListenerProtocol, port int) *CreateListenerRequest {
	opts := new(CreateListenerRequest)
	opts.ListenerName = name
	opts.ListenerProtocol = protocol
	opts.ListenerProtocolPort = port
	opts.AllowedCidrs = "0.0.0.0/0"
	opts.TimeoutClient = 50
	opts.TimeoutMember = 50
	opts.TimeoutConnection = 5

	return opts
}

type ListenerProtocol string

type CreateListenerRequest struct {
	AllowedCidrs                string           `json:"allowedCidrs"`
	ListenerName                string           `json:"listenerName"`
	ListenerProtocol            ListenerProtocol `json:"listenerProtocol"`
	ListenerProtocolPort        int              `json:"listenerProtocolPort"`
	TimeoutClient               int              `json:"timeoutClient"`
	TimeoutConnection           int              `json:"timeoutConnection"`
	TimeoutMember               int              `json:"timeoutMember"`
	DefaultPoolID               *string          `json:"defaultPoolId"`
	CertificateAuthorities      *[]string        `json:"certificateAuthorities"`
	ClientCertificate           *string          `json:"clientCertificate"`
	DefaultCertificateAuthority *string          `json:"defaultCertificateAuthority"`

	common.LoadBalancerCommon
	common.UserAgent
}

func (r *CreateListenerRequest) prepare() {
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

func (r *CreateListenerRequest) WithLoadBalancerID(lbid string) *CreateListenerRequest {
	r.LoadBalancerID = lbid
	return r
}

func (r *CreateListenerRequest) WithDefaultPoolID(poolID string) *CreateListenerRequest {
	r.DefaultPoolID = &poolID
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

func (r *CreateListenerRequest) ToMap() map[string]any {
	return map[string]any{
		"listenerName":         r.ListenerName,
		"listenerProtocol":     r.ListenerProtocol,
		"listenerProtocolPort": r.ListenerProtocolPort,
		"timeoutClient":        r.TimeoutClient,
		"timeoutConnection":    r.TimeoutConnection,
		"timeoutMember":        r.TimeoutMember,
		"allowedCidrs":         r.AllowedCidrs,
		"defaultPoolId":        r.DefaultPoolID,
	}
}
