package sdkerror

import "regexp"

const (
	patternNetworkNotFound                  = "is not found"
	patternSubnetNotFound                   = `subnet with id [^.]+ is not found`
	patternSubnetNotFound2                  = `subnet id: [^.]+ not found`
	patternSubnetNotBelongNetwork           = `subnet id: [^.]+ belong to network id: [^.]+ not found`
	patternInternalNetworkInterfaceNotFound = `internal network interface with id [^.]+ is not found`
	patternWanIpAvailable                   = "wan ip is available"
	pattermWapIDNotFound                    = "cannot get wan ip with id"
	patternAddressPairExisted               = `address pair with internal network interface  id [^.]+ already exists`
)

var (
	regexErrorSubnetNotFound                   = regexp.MustCompile(patternSubnetNotFound)
	regexErrorSubnetNotFound2                  = regexp.MustCompile(patternSubnetNotFound2)
	regexErrorSubnetNotBelongNetwork           = regexp.MustCompile(patternSubnetNotBelongNetwork)
	regexErrorInternalNetworkInterfaceNotFound = regexp.MustCompile(patternInternalNetworkInterfaceNotFound)
	regexErrorAddressPairExisted               = regexp.MustCompile(patternAddressPairExisted)
)

func init() {
	register(EcVServerNetworkNotFound, &classifier{
		match: matchAnyOf(containsAny(patternNetworkNotFound), matchErrCode("VPC_IS_NOT_FOUND")),
	})
	register(EcVServerSubnetNotBelongNetwork, &classifier{match: matchRegexps(regexErrorSubnetNotBelongNetwork)})
	register(EcVServerSubnetNotFound, &classifier{
		match: matchAnyOf(matchRegexps(regexErrorSubnetNotFound, regexErrorSubnetNotFound2), matchErrCode("SUBNET_IS_NOT_FOUND")),
	})
	register(EcVServerInternalNetworkInterfaceNotFound, &classifier{match: matchRegexps(regexErrorInternalNetworkInterfaceNotFound)})
	register(EcVServerAddressPairExisted, &classifier{match: matchRegexps(regexErrorAddressPairExisted)})
	register(EcVServerWanIpAvailable, &classifier{match: containsAny(patternWanIpAvailable)})
	register(EcVServerWanIDNotFound, &classifier{match: containsAny(pattermWapIDNotFound)})
}
