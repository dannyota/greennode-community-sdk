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

func NewUpdateListenerRequest(lbID, listenerID string) IUpdateListenerRequest {
	opts := new(UpdateListenerRequest)
	opts.LoadBalancerID = lbID
	opts.ListenerID = listenerID

	return opts
}

func NewListListenersByLoadBalancerIDRequest(lbID string) IListListenersByLoadBalancerIDRequest {
	opts := new(ListListenersByLoadBalancerIDRequest)
	opts.LoadBalancerID = lbID

	return opts
}

func NewDeleteListenerByIDRequest(lbID, listenerID string) IDeleteListenerByIDRequest {
	opts := new(DeleteListenerByIDRequest)
	opts.LoadBalancerID = lbID
	opts.ListenerID = listenerID

	return opts
}

func NewGetListenerByIDRequest(lbID, listenerID string) IGetListenerByIDRequest {
	opts := new(GetListenerByIDRequest)
	opts.LoadBalancerID = lbID
	opts.ListenerID = listenerID

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
	DefaultPoolID               *string                        `json:"defaultPoolId"`
	CertificateAuthorities      *[]string                      `json:"certificateAuthorities"`
	ClientCertificate           *string                        `json:"clientCertificate"`
	DefaultCertificateAuthority *string                        `json:"defaultCertificateAuthority"`
	InsertHeaders               *[]entity.ListenerInsertHeader `json:"insertHeaders"`

	common.LoadBalancerCommon
	common.UserAgent
}

func (r *CreateListenerRequest) AddUserAgent(agent ...string) ICreateListenerRequest {
	r.UserAgent.AddUserAgent(agent...)
	return r
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

	common.LoadBalancerCommon
	common.ListenerCommon
	common.UserAgent
}

func (r *UpdateListenerRequest) AddUserAgent(agent ...string) IUpdateListenerRequest {
	r.UserAgent.AddUserAgent(agent...)
	return r
}

type ListListenersByLoadBalancerIDRequest struct {
	common.LoadBalancerCommon
	common.UserAgent
}

func (r *ListListenersByLoadBalancerIDRequest) AddUserAgent(agent ...string) IListListenersByLoadBalancerIDRequest {
	r.UserAgent.AddUserAgent(agent...)
	return r
}

type DeleteListenerByIDRequest struct {
	common.LoadBalancerCommon
	common.ListenerCommon
	common.UserAgent
}

func (r *DeleteListenerByIDRequest) AddUserAgent(agent ...string) IDeleteListenerByIDRequest {
	r.UserAgent.AddUserAgent(agent...)
	return r
}

type GetListenerByIDRequest struct {
	common.LoadBalancerCommon
	common.ListenerCommon
	common.UserAgent
}

func (r *GetListenerByIDRequest) AddUserAgent(agent ...string) IGetListenerByIDRequest {
	r.UserAgent.AddUserAgent(agent...)
	return r
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

func (r *CreateListenerRequest) WithLoadBalancerID(lbid string) ICreateListenerRequest {
	r.LoadBalancerID = lbid
	return r
}

func (r *CreateListenerRequest) WithDefaultPoolID(poolID string) ICreateListenerRequest {
	r.DefaultPoolID = &poolID
	return r
}

func (r *CreateListenerRequest) WithCertificateAuthorities(ca *[]string) ICreateListenerRequest {
	r.CertificateAuthorities = ca
	return r
}

func (r *CreateListenerRequest) WithClientCertificate(clientCert *string) ICreateListenerRequest {
	r.ClientCertificate = clientCert
	return r
}

func (r *CreateListenerRequest) WithDefaultCertificateAuthority(defaultCA *string) ICreateListenerRequest {
	r.DefaultCertificateAuthority = defaultCA
	return r
}

func (r *CreateListenerRequest) WithInsertHeaders(pheaders ...string) ICreateListenerRequest {
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

func (r *CreateListenerRequest) ToMap() map[string]any {
	return map[string]any{
		"listenerName":                r.ListenerName,
		"listenerProtocol":            r.ListenerProtocol,
		"listenerProtocolPort":        r.ListenerProtocolPort,
		"timeoutClient":               r.TimeoutClient,
		"timeoutConnection":           r.TimeoutConnection,
		"timeoutMember":               r.TimeoutMember,
		"allowedCidrs":                r.AllowedCidrs,
		"defaultPoolId":               r.DefaultPoolID,
		"certificateAuthorities":      r.CertificateAuthorities,
		"clientCertificate":           r.ClientCertificate,
		"defaultCertificateAuthority": r.DefaultCertificateAuthority,
		"insertHeaders":               r.InsertHeaders,
	}
}

func (r *UpdateListenerRequest) ToRequestBody() any {
	return r
}

func (r *UpdateListenerRequest) WithCidrs(cidrs ...string) IUpdateListenerRequest {
	if len(cidrs) < 1 {
		return r
	}

	r.AllowedCidrs = strings.Join(cidrs, ",")
	return r
}

func (r *UpdateListenerRequest) WithTimeoutClient(toc int) IUpdateListenerRequest {
	r.TimeoutClient = toc
	return r
}

func (r *UpdateListenerRequest) WithTimeoutConnection(toc int) IUpdateListenerRequest {
	r.TimeoutConnection = toc
	return r
}

func (r *UpdateListenerRequest) WithTimeoutMember(tom int) IUpdateListenerRequest {
	r.TimeoutMember = tom
	return r
}

func (r *UpdateListenerRequest) WithDefaultPoolID(poolID string) IUpdateListenerRequest {
	r.DefaultPoolID = poolID
	return r
}

func (r *UpdateListenerRequest) WithInsertHeaders(pheaders ...string) IUpdateListenerRequest {
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

func (r *UpdateListenerRequest) WithCertificateAuthorities(ca ...string) IUpdateListenerRequest {
	r.CertificateAuthorities = &ca
	return r
}

func (r *UpdateListenerRequest) WithClientCertificate(clientCert string) IUpdateListenerRequest {
	r.ClientCertificate = &clientCert
	return r
}

func (r *UpdateListenerRequest) WithDefaultCertificateAuthority(defaultCA string) IUpdateListenerRequest {
	r.DefaultCertificateAuthority = &defaultCA
	return r
}
