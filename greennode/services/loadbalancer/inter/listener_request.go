package inter

import (
	"strings"

	"github.com/dannyota/greennode-community-sdk/v2/greennode/services/common"
)

type ICreateListenerRequest interface {
	ToRequestBody() any
	WithAllowedCidrs(cidrs ...string) ICreateListenerRequest
	WithLoadBalancerID(lbid string) ICreateListenerRequest
	WithDefaultPoolID(poolID string) ICreateListenerRequest
	WithTimeoutClient(toc int) ICreateListenerRequest
	WithTimeoutConnection(toc int) ICreateListenerRequest
	WithTimeoutMember(tom int) ICreateListenerRequest
	AddCidrs(cidrs ...string) ICreateListenerRequest
	ParseUserAgent() string
	GetLoadBalancerID() string
	ToMap() map[string]any
}

const (
	ListenerProtocolTCP   ListenerProtocol = "TCP"
	ListenerProtocolUDP   ListenerProtocol = "UDP"
	ListenerProtocolHTTP  ListenerProtocol = "HTTP"
	ListenerProtocolHTTPS ListenerProtocol = "HTTPS"
)

func NewCreateListenerRequest(name string, protocol ListenerProtocol, port int) ICreateListenerRequest {
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

func (r *CreateListenerRequest) ToRequestBody() any {
	if r == nil {
		return nil
	}

	if r.ListenerProtocol == ListenerProtocolHTTPS {
		return r
	}

	r.CertificateAuthorities = nil
	r.ClientCertificate = nil
	r.DefaultCertificateAuthority = nil

	return r
}

func (r *CreateListenerRequest) WithAllowedCidrs(cidrs ...string) ICreateListenerRequest {
	if len(cidrs) < 1 {
		return r
	}

	r.AllowedCidrs = strings.Join(cidrs, ",")
	return r
}

func (r *CreateListenerRequest) WithTimeoutClient(toc int) ICreateListenerRequest {
	r.TimeoutClient = toc
	return r
}

func (r *CreateListenerRequest) WithTimeoutConnection(toc int) ICreateListenerRequest {
	r.TimeoutConnection = toc
	return r
}

func (r *CreateListenerRequest) WithTimeoutMember(tom int) ICreateListenerRequest {
	r.TimeoutMember = tom
	return r
}

func (r *CreateListenerRequest) WithLoadBalancerID(lbid string) ICreateListenerRequest {
	r.LoadBalancerID = lbid
	return r
}

func (r *CreateListenerRequest) WithDefaultPoolID(poolID string) ICreateListenerRequest {
	r.DefaultPoolID = &poolID
	return r
}

func (r *CreateListenerRequest) AddCidrs(cidrs ...string) ICreateListenerRequest {
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
