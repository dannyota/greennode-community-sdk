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

func (s *CreateListenerRequest) ToRequestBody() interface{} {
	if s == nil {
		return nil
	}

	if s.ListenerProtocol == ListenerProtocolHTTPS {
		return s
	}

	s.CertificateAuthorities = nil
	s.ClientCertificate = nil
	s.DefaultCertificateAuthority = nil

	return s
}

func (s *CreateListenerRequest) WithAllowedCidrs(cidrs ...string) ICreateListenerRequest {
	if len(cidrs) < 1 {
		return s
	}

	s.AllowedCidrs = strings.Join(cidrs, ",")
	return s
}

func (s *CreateListenerRequest) WithTimeoutClient(toc int) ICreateListenerRequest {
	s.TimeoutClient = toc
	return s
}

func (s *CreateListenerRequest) WithTimeoutConnection(toc int) ICreateListenerRequest {
	s.TimeoutConnection = toc
	return s
}

func (s *CreateListenerRequest) WithTimeoutMember(tom int) ICreateListenerRequest {
	s.TimeoutMember = tom
	return s
}

func (s *CreateListenerRequest) WithLoadBalancerID(lbid string) ICreateListenerRequest {
	s.LoadBalancerID = lbid
	return s
}

func (s *CreateListenerRequest) WithDefaultPoolID(poolID string) ICreateListenerRequest {
	s.DefaultPoolID = &poolID
	return s
}

func (s *CreateListenerRequest) AddCidrs(cidrs ...string) ICreateListenerRequest {
	if len(cidrs) < 1 {
		return s
	}

	if s.AllowedCidrs == "" {
		return s.WithAllowedCidrs(cidrs...)
	} else {
		s.AllowedCidrs = s.AllowedCidrs + "," + strings.Join(cidrs, ",")
	}

	return s
}

func (s *CreateListenerRequest) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"listenerName":         s.ListenerName,
		"listenerProtocol":     s.ListenerProtocol,
		"listenerProtocolPort": s.ListenerProtocolPort,
		"timeoutClient":        s.TimeoutClient,
		"timeoutConnection":    s.TimeoutConnection,
		"timeoutMember":        s.TimeoutMember,
		"allowedCidrs":         s.AllowedCidrs,
		"defaultPoolId":        s.DefaultPoolID,
	}
}
