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
	return &DeleteSecgroupRuleByIDRequest{
		SecgroupRuleID: secgroupRuleID,
		SecgroupCommon: common.SecgroupCommon{
			SecgroupID: "undefined",
		},
	}
}

func NewListSecgroupRulesBySecgroupIDRequest(securityGroupID string) *ListSecgroupRulesBySecgroupIDRequest {
	return &ListSecgroupRulesBySecgroupIDRequest{
		SecgroupCommon: common.SecgroupCommon{
			SecgroupID: securityGroupID,
		},
	}
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
	}

	SecgroupRuleDirection string
	SecgroupRuleEtherType string
	SecgroupRuleProtocol  string
)

type DeleteSecgroupRuleByIDRequest struct { //__________________________________________________________________________
	SecgroupRuleID string

	common.SecgroupCommon
}

func (r *DeleteSecgroupRuleByIDRequest) GetSecgroupRuleID() string {
	return r.SecgroupRuleID
}

type ListSecgroupRulesBySecgroupIDRequest struct { //___________________________________________________________________
	common.SecgroupCommon
}
