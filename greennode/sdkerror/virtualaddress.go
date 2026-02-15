package sdkerror

import "regexp"

const (
	patternVirtualAddressNotFound    = `virtual ip address with id [^.]+ is not found`
	patternAddressPairNotFound       = `address pair with uuid: [^.]+ was not existing`
	patternVirtualAddressExceedQuota = "exceeded virtual_ip_address quota"
	patternVirtualAddressInUse       = "ip address is already in use"
	patternVirtualAddressInUse2      = `virtual ip address id [^.]+ is used. please remove address pairs first`
)

var (
	regexErrorVirtualAddressNotFound = regexp.MustCompile(patternVirtualAddressNotFound)
	regexErrorAddressPairNotFound    = regexp.MustCompile(patternAddressPairNotFound)
	regexErrorVirtualAddressInUse    = regexp.MustCompile(patternVirtualAddressInUse2)
)

func init() {
	register(EcVServerVirtualAddressNotFound, &classifier{
		match: matchRegexps(regexErrorVirtualAddressNotFound, regexErrorAddressPairNotFound),
	})
	register(EcVServerVirtualAddressExceedQuota, &classifier{
		match: containsAny(patternVirtualAddressExceedQuota), category: ErrCatQuota,
	})
	register(EcVServerVirtualAddressInUse, &classifier{
		match: matchAnyOf(containsAny(patternVirtualAddressInUse), matchRegexps(regexErrorVirtualAddressInUse)),
		category: ErrCatQuota,
	})
}
