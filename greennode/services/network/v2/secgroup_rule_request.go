package v2

import "github.com/dannyota/greennode-community-sdk/v2/greennode/services/common"

func NewCreateSecgroupRuleRequest(
	direction SecgroupRuleDirection,
	etherType SecgroupRuleEtherType,
	protocol SecgroupRuleProtocol,
	portMinRange, portMaxRange int,
	remoteIpPrefix, securityGroupID, description string) *CreateSecgroupRuleRequest {

	opt := &CreateSecgroupRuleRequest{
		Description:     description,
		Direction:       direction,
		EtherType:       etherType,
		PortRangeMax:    portMaxRange,
		PortRangeMin:    portMinRange,
		Protocol:        protocol,
		RemoteIPPrefix:  remoteIpPrefix,
		SecurityGroupID: securityGroupID,
	}
	opt.SecgroupID = securityGroupID
	return opt
}

func NewDeleteSecgroupRuleByIDRequest(secgroupRuleID string) *DeleteSecgroupRuleByIDRequest {
	opt := new(DeleteSecgroupRuleByIDRequest)
	opt.SecgroupID = "undefined"
	opt.SecgroupRuleID = secgroupRuleID
	return opt
}

func (r *DeleteSecgroupRuleByIDRequest) AddUserAgent(agent ...string) *DeleteSecgroupRuleByIDRequest {
	r.UserAgent.AddUserAgent(agent...)
	return r
}

func NewListSecgroupRulesBySecgroupIDRequest(securityGroupID string) *ListSecgroupRulesBySecgroupIDRequest {
	opt := new(ListSecgroupRulesBySecgroupIDRequest)
	opt.SecgroupID = securityGroupID
	return opt
}

const (
	SecgroupRuleProtocolTCP    SecgroupRuleProtocol = "tcp"
	SecgroupRuleProtocolUDP    SecgroupRuleProtocol = "udp"
	SecgroupRuleProtocolICMP   SecgroupRuleProtocol = "icmp"
	SecgroupRuleProtocolAll    SecgroupRuleProtocol = "any"
	SecgroupRuleProtocolIpInIp SecgroupRuleProtocol = "4"
)

const (
	SecgroupRuleEtherTypeIPv4 SecgroupRuleEtherType = "IPv4"
	SecgroupRuleEtherTypeIPv6 SecgroupRuleEtherType = "IPv6"
)

const (
	SecgroupRuleDirectionIngress SecgroupRuleDirection = "ingress"
	SecgroupRuleDirectionEgress  SecgroupRuleDirection = "egress"
)

type ( //_______________________________________________________________________________________________________________
	CreateSecgroupRuleRequest struct {
		Description     string                `json:"description"`
		Direction       SecgroupRuleDirection `json:"direction"`
		EtherType       SecgroupRuleEtherType `json:"etherType"`
		PortRangeMax    int                   `json:"portRangeMax"`
		PortRangeMin    int                   `json:"portRangeMin"`
		Protocol        SecgroupRuleProtocol  `json:"protocol"`
		RemoteIPPrefix  string                `json:"remoteIpPrefix"`
		SecurityGroupID string                `json:"securityGroupId"`

		common.SecgroupCommon
		common.UserAgent
	}

	SecgroupRuleDirection string
	SecgroupRuleEtherType string
	SecgroupRuleProtocol  string
)

func (r *CreateSecgroupRuleRequest) ToRequestBody() any {
	return r
}

func (r *CreateSecgroupRuleRequest) ToMap() map[string]any {
	return map[string]any{
		"description":     r.Description,
		"direction":       r.Direction,
		"etherType":       r.EtherType,
		"portRangeMax":    r.PortRangeMax,
		"portRangeMin":    r.PortRangeMin,
		"protocol":        r.Protocol,
		"remoteIpPrefix":  r.RemoteIPPrefix,
		"securityGroupId": r.SecurityGroupID,
	}
}

func (r *CreateSecgroupRuleRequest) AddUserAgent(agent ...string) *CreateSecgroupRuleRequest {
	r.UserAgent.AddUserAgent(agent...)
	return r
}

type DeleteSecgroupRuleByIDRequest struct { //__________________________________________________________________________
	SecgroupRuleID string

	common.UserAgent
	common.SecgroupCommon
}

func (r *DeleteSecgroupRuleByIDRequest) GetSecgroupRuleID() string {
	return r.SecgroupRuleID
}

type ListSecgroupRulesBySecgroupIDRequest struct { //___________________________________________________________________
	common.SecgroupCommon
	common.UserAgent
}

func (r *ListSecgroupRulesBySecgroupIDRequest) AddUserAgent(agent ...string) *ListSecgroupRulesBySecgroupIDRequest {
	r.UserAgent.AddUserAgent(agent...)
	return r
}
