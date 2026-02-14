package v2

import "github.com/dannyota/greennode-community-sdk/v2/greennode/services/common"

func NewCreateSecgroupRuleRequest(
	direction SecgroupRuleDirection,
	etherType SecgroupRuleEtherType,
	protocol SecgroupRuleProtocol,
	portMinRange, portMaxRange int,
	remoteIpPrefix, securityGroupId, description string) ICreateSecgroupRuleRequest {

	opt := &CreateSecgroupRuleRequest{
		Description:     description,
		Direction:       direction,
		EtherType:       etherType,
		PortRangeMax:    portMaxRange,
		PortRangeMin:    portMinRange,
		Protocol:        protocol,
		RemoteIPPrefix:  remoteIpPrefix,
		SecurityGroupID: securityGroupId,
	}
	opt.SecgroupId = securityGroupId
	return opt
}

func NewDeleteSecgroupRuleByIdRequest(secgroupRuleId string) IDeleteSecgroupRuleByIdRequest {
	opt := new(DeleteSecgroupRuleByIdRequest)
	opt.SecgroupId = "undefined"
	opt.SecgroupRuleId = secgroupRuleId
	return opt
}

func (s *DeleteSecgroupRuleByIdRequest) AddUserAgent(agent ...string) IDeleteSecgroupRuleByIdRequest {
	s.UserAgent.AddUserAgent(agent...)
	return s
}

func NewListSecgroupRulesBySecgroupIdRequest(securityGroupId string) IListSecgroupRulesBySecgroupIdRequest {
	opt := new(ListSecgroupRulesBySecgroupIdRequest)
	opt.SecgroupId = securityGroupId
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

func (s *CreateSecgroupRuleRequest) ToRequestBody() interface{} {
	return s
}

func (s *CreateSecgroupRuleRequest) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"description":     s.Description,
		"direction":       s.Direction,
		"etherType":       s.EtherType,
		"portRangeMax":    s.PortRangeMax,
		"portRangeMin":    s.PortRangeMin,
		"protocol":        s.Protocol,
		"remoteIpPrefix":  s.RemoteIPPrefix,
		"securityGroupId": s.SecurityGroupID,
	}
}

func (s *CreateSecgroupRuleRequest) AddUserAgent(agent ...string) ICreateSecgroupRuleRequest {
	s.UserAgent.AddUserAgent(agent...)
	return s
}

type DeleteSecgroupRuleByIdRequest struct { //__________________________________________________________________________
	SecgroupRuleId string

	common.UserAgent
	common.SecgroupCommon
}

func (s *DeleteSecgroupRuleByIdRequest) GetSecgroupRuleId() string {
	return s.SecgroupRuleId
}

type ListSecgroupRulesBySecgroupIdRequest struct { //___________________________________________________________________
	common.SecgroupCommon
	common.UserAgent
}

func (s *ListSecgroupRulesBySecgroupIdRequest) AddUserAgent(agent ...string) IListSecgroupRulesBySecgroupIdRequest {
	s.UserAgent.AddUserAgent(agent...)
	return s
}
