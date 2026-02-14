package sdk_error

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

func WithErrorVirtualAddressNotFound(perrResp IErrorRespone) func(sdkError IError) {
	return func(sdkError IError) {
		if perrResp == nil {
			return
		}

		errMsg := strings.ToLower(strings.TrimSpace(perrResp.GetMessage()))
		if regexErrorVirtualAddressNotFound.FindString(errMsg) != "" {
			sdkError.WithErrorCode(EcVServerVirtualAddressNotFound).
				WithMessage(errMsg).
				WithErrors(perrResp.GetError())
		}
	}
}

func WithErrorAddressPairNotFound(perrResp IErrorRespone) func(sdkError IError) {
	return func(sdkError IError) {
		if perrResp == nil {
			return
		}

		errMsg := strings.ToLower(strings.TrimSpace(perrResp.GetMessage()))
		if regexErrorAddressPairNotFound.FindString(errMsg) != "" {
			sdkError.WithErrorCode(EcVServerVirtualAddressNotFound).
				WithMessage(errMsg).
				WithErrors(perrResp.GetError())
		}
	}
}

func WithErrorVirtualAddressExceedQuota(perrResp IErrorRespone) func(sdkError IError) {
	return func(sdkError IError) {
		if perrResp == nil {
			return
		}

		errMsg := perrResp.GetMessage()
		if strings.Contains(strings.ToLower(strings.TrimSpace(errMsg)), patternVirtualAddressExceedQuota) {
			sdkError.WithErrorCode(EcVServerVirtualAddressExceedQuota).
				WithMessage(errMsg).
				WithErrors(perrResp.GetError()).
				WithErrorCategories(ErrCatQuota)
		}
	}
}

func WithErrorVirtualAddressInUse(perrResp IErrorRespone) func(sdkError IError) {
	return func(sdkError IError) {
		if perrResp == nil {
			return
		}

		errMsg := strings.ToLower(strings.TrimSpace(perrResp.GetMessage()))
		if strings.Contains(errMsg, patternVirtualAddressInUse) ||
			regexErrorVirtualAddressInUse.FindString(errMsg) != "" {
			sdkError.WithErrorCode(EcVServerVirtualAddressInUse).
				WithMessage(perrResp.GetMessage()).
				WithErrors(perrResp.GetError()).
				WithErrorCategories(ErrCatQuota)
		}
	}
}
