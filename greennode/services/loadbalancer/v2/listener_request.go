package v2

import (
	"strings"

	"github.com/dannyota/greennode-community-sdk/v2/greennode/entity"
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

func NewUpdateListenerRequest(lbId, listenerId string) IUpdateListenerRequest {
	opts := new(UpdateListenerRequest)
	opts.LoadBalancerId = lbId
	opts.ListenerId = listenerId

	return opts
}

func NewListListenersByLoadBalancerIdRequest(lbId string) IListListenersByLoadBalancerIdRequest {
	opts := new(ListListenersByLoadBalancerIdRequest)
	opts.LoadBalancerId = lbId

	return opts
}

func NewDeleteListenerByIdRequest(lbId, listenerId string) IDeleteListenerByIdRequest {
	opts := new(DeleteListenerByIdRequest)
	opts.LoadBalancerId = lbId
	opts.ListenerId = listenerId

	return opts
}

func NewGetListenerByIdRequest(lbId, listenerId string) IGetListenerByIdRequest {
	opts := new(GetListenerByIdRequest)
	opts.LoadBalancerId = lbId
	opts.ListenerId = listenerId

	return opts
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
	DefaultPoolId               *string                        `json:"defaultPoolId"`
	CertificateAuthorities      *[]string                      `json:"certificateAuthorities"`
	ClientCertificate           *string                        `json:"clientCertificate"`
	DefaultCertificateAuthority *string                        `json:"defaultCertificateAuthority"`
	InsertHeaders               *[]entity.ListenerInsertHeader `json:"insertHeaders"`

	common.LoadBalancerCommon
	common.UserAgent
}

func (s *CreateListenerRequest) AddUserAgent(agent ...string) ICreateListenerRequest {
	s.UserAgent.AddUserAgent(agent...)
	return s
}

type UpdateListenerRequest struct {
	AllowedCidrs                string                         `json:"allowedCidrs"`
	DefaultPoolId               string                         `json:"defaultPoolId"`
	TimeoutClient               int                            `json:"timeoutClient"`
	TimeoutConnection           int                            `json:"timeoutConnection"`
	TimeoutMember               int                            `json:"timeoutMember"`
	CertificateAuthorities      *[]string                      `json:"certificateAuthorities"`
	ClientCertificate           *string                        `json:"clientCertificate"`
	DefaultCertificateAuthority *string                        `json:"defaultCertificateAuthority"`
	InsertHeaders               *[]entity.ListenerInsertHeader `json:"insertHeaders"`

	common.LoadBalancerCommon
	common.ListenerCommon
	common.UserAgent
}

func (s *UpdateListenerRequest) AddUserAgent(agent ...string) IUpdateListenerRequest {
	s.UserAgent.AddUserAgent(agent...)
	return s
}

type ListListenersByLoadBalancerIdRequest struct {
	common.LoadBalancerCommon
	common.UserAgent
}

func (s *ListListenersByLoadBalancerIdRequest) AddUserAgent(agent ...string) IListListenersByLoadBalancerIdRequest {
	s.UserAgent.AddUserAgent(agent...)
	return s
}

type DeleteListenerByIdRequest struct {
	common.LoadBalancerCommon
	common.ListenerCommon
	common.UserAgent
}

func (s *DeleteListenerByIdRequest) AddUserAgent(agent ...string) IDeleteListenerByIdRequest {
	s.UserAgent.AddUserAgent(agent...)
	return s
}

type GetListenerByIdRequest struct {
	common.LoadBalancerCommon
	common.ListenerCommon
	common.UserAgent
}

func (s *GetListenerByIdRequest) AddUserAgent(agent ...string) IGetListenerByIdRequest {
	s.UserAgent.AddUserAgent(agent...)
	return s
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

func (s *CreateListenerRequest) WithLoadBalancerId(lbid string) ICreateListenerRequest {
	s.LoadBalancerId = lbid
	return s
}

func (s *CreateListenerRequest) WithDefaultPoolId(poolId string) ICreateListenerRequest {
	s.DefaultPoolId = &poolId
	return s
}

func (s *CreateListenerRequest) WithCertificateAuthorities(ca *[]string) ICreateListenerRequest {
	s.CertificateAuthorities = ca
	return s
}

func (s *CreateListenerRequest) WithClientCertificate(clientCert *string) ICreateListenerRequest {
	s.ClientCertificate = clientCert
	return s
}

func (s *CreateListenerRequest) WithDefaultCertificateAuthority(defaultCA *string) ICreateListenerRequest {
	s.DefaultCertificateAuthority = defaultCA
	return s
}

func (s *CreateListenerRequest) WithInsertHeaders(pheaders ...string) ICreateListenerRequest {
	if len(pheaders) < 1 {
		s.InsertHeaders = nil
		return s
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
	s.InsertHeaders = &headers
	return s
}

func (s *CreateListenerRequest) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"listenerName":                s.ListenerName,
		"listenerProtocol":            s.ListenerProtocol,
		"listenerProtocolPort":        s.ListenerProtocolPort,
		"timeoutClient":               s.TimeoutClient,
		"timeoutConnection":           s.TimeoutConnection,
		"timeoutMember":               s.TimeoutMember,
		"allowedCidrs":                s.AllowedCidrs,
		"defaultPoolId":               s.DefaultPoolId,
		"certificateAuthorities":      s.CertificateAuthorities,
		"clientCertificate":           s.ClientCertificate,
		"defaultCertificateAuthority": s.DefaultCertificateAuthority,
		"insertHeaders":               s.InsertHeaders,
	}
}

func (s *UpdateListenerRequest) ToRequestBody() interface{} {
	return s
}

func (s *UpdateListenerRequest) WithCidrs(cidrs ...string) IUpdateListenerRequest {
	if len(cidrs) < 1 {
		return s
	}

	s.AllowedCidrs = strings.Join(cidrs, ",")
	return s
}

func (s *UpdateListenerRequest) WithTimeoutClient(toc int) IUpdateListenerRequest {
	s.TimeoutClient = toc
	return s
}

func (s *UpdateListenerRequest) WithTimeoutConnection(toc int) IUpdateListenerRequest {
	s.TimeoutConnection = toc
	return s
}

func (s *UpdateListenerRequest) WithTimeoutMember(tom int) IUpdateListenerRequest {
	s.TimeoutMember = tom
	return s
}

func (s *UpdateListenerRequest) WithDefaultPoolId(poolId string) IUpdateListenerRequest {
	s.DefaultPoolId = poolId
	return s
}

func (s *UpdateListenerRequest) WithInsertHeaders(pheaders ...string) IUpdateListenerRequest {
	if len(pheaders) < 1 {
		s.InsertHeaders = nil
		return s
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
	s.InsertHeaders = &headers
	return s
}

func (s *UpdateListenerRequest) WithCertificateAuthorities(ca ...string) IUpdateListenerRequest {
	s.CertificateAuthorities = &ca
	return s
}

func (s *UpdateListenerRequest) WithClientCertificate(clientCert string) IUpdateListenerRequest {
	s.ClientCertificate = &clientCert
	return s
}

func (s *UpdateListenerRequest) WithDefaultCertificateAuthority(defaultCA string) IUpdateListenerRequest {
	s.DefaultCertificateAuthority = &defaultCA
	return s
}
