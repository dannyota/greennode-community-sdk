package sdkerror

import (
	"regexp"
	"strings"
)

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
	regexErrorVirtualAddressInUse    = regexp.MustCompile(patternVirtualAddressInUse)
)

func WithErrorVirtualAddressNotFound(errResp ErrorResponse) func(sdkError Error) {
	return func(sdkError Error) {
		if errResp == nil {
			return
		}

		errMsg := strings.ToLower(strings.TrimSpace(errResp.GetMessage()))
		if regexErrorVirtualAddressNotFound.FindString(errMsg) != "" {
			sdkError.WithErrorCode(EcVServerVirtualAddressNotFound).
				WithMessage(errMsg).
				WithErrors(errResp.Err())
		}
	}
}

func WithErrorAddressPairNotFound(errResp ErrorResponse) func(sdkError Error) {
	return func(sdkError Error) {
		if errResp == nil {
			return
		}

		errMsg := strings.ToLower(strings.TrimSpace(errResp.GetMessage()))
		if regexErrorAddressPairNotFound.FindString(errMsg) != "" {
			sdkError.WithErrorCode(EcVServerVirtualAddressNotFound).
				WithMessage(errMsg).
				WithErrors(errResp.Err())
		}
	}
}

func WithErrorVirtualAddressExceedQuota(errResp ErrorResponse) func(sdkError Error) {
	return func(sdkError Error) {
		if errResp == nil {
			return
		}

		errMsg := errResp.GetMessage()
		if strings.Contains(strings.ToLower(strings.TrimSpace(errMsg)), patternVirtualAddressExceedQuota) {
			sdkError.WithErrorCode(EcVServerVirtualAddressExceedQuota).
				WithMessage(errMsg).
				WithErrors(errResp.Err()).
				WithErrorCategories(ErrCatQuota)
		}
	}
}

func WithErrorVirtualAddressInUse(errResp ErrorResponse) func(sdkError Error) {
	return func(sdkError Error) {
		if errResp == nil {
			return
		}

		errMsg := strings.ToLower(strings.TrimSpace(errResp.GetMessage()))
		if strings.Contains(errMsg, patternVirtualAddressInUse) ||
			regexErrorVirtualAddressInUse.FindString(errMsg) != "" {
			sdkError.WithErrorCode(EcVServerVirtualAddressInUse).
				WithMessage(errResp.GetMessage()).
				WithErrors(errResp.Err()).
				WithErrorCategories(ErrCatQuota)
		}
	}
}
