package sdkerror

import (
	"regexp"
	"strings"
)

const (
	patternNetworkNotFound                  = "is not found"
	patternSubnetNotFound                   = `subnet with id [^.]+ is not found`
	patternSubnetNotFound2                  = `subnet id: [^.]+ not found`
	patternSubnetNotBelongNetwork           = `subnet id: [^.]+ belong to network id: [^.]+ not found`
	patternInternalNetworkInterfaceNotFound = `internal network interface with id [^.]+ is not found`
	patternWanIpAvailable                   = "wan ip is available"
	pattermWapIdNotFound                    = "cannot get wan ip with id"
	patternAddressPairExisted               = `address pair with internal network interface  id [^.]+ already exists`
)

var (
	regexErrorSubnetNotFound                   = regexp.MustCompile(patternSubnetNotFound)
	regexErrorSubnetNotFound2                  = regexp.MustCompile(patternSubnetNotFound2)
	regexErrorSubnetNotBelongNetwork           = regexp.MustCompile(patternSubnetNotBelongNetwork)
	regexErrorInternalNetworkInterfaceNotFound = regexp.MustCompile(patternInternalNetworkInterfaceNotFound)
	regexErrorAddressPairExisted               = regexp.MustCompile(patternAddressPairExisted)
)

func WithErrorNetworkNotFound(perrResp ErrorResponse) func(sdkError Error) {
	return func(sdkError Error) {
		if perrResp == nil {
			return
		}

		errMsg := perrResp.GetMessage()
		if strings.Contains(strings.ToLower(strings.TrimSpace(errMsg)), patternNetworkNotFound) ||
			strings.ToUpper(strings.TrimSpace(perrResp.GetError().Error())) == "VPC_IS_NOT_FOUND" {
			sdkError.WithErrorCode(EcVServerNetworkNotFound).
				WithMessage(errMsg).
				WithErrors(perrResp.GetError())
		}
	}
}

func WithErrorSubnetNotBelongNetwork(perrResp ErrorResponse) func(sdkError Error) {
	return func(sdkError Error) {
		if perrResp == nil {
			return
		}

		errMsg := strings.ToLower(strings.TrimSpace(perrResp.GetMessage()))
		if regexErrorSubnetNotBelongNetwork.FindString(errMsg) != "" {
			sdkError.WithErrorCode(EcVServerSubnetNotBelongNetwork).
				WithMessage(errMsg).
				WithErrors(perrResp.GetError())
		}
	}
}

func WithErrorSubnetNotFound(perrResp ErrorResponse) func(sdkError Error) {
	return func(sdkError Error) {
		if perrResp == nil {
			return
		}

		errMsg := strings.ToLower(strings.TrimSpace(perrResp.GetMessage()))
		if regexErrorSubnetNotFound.FindString(errMsg) != "" ||
			regexErrorSubnetNotFound2.FindString(errMsg) != "" ||
			strings.ToUpper(strings.TrimSpace(perrResp.GetError().Error())) == "SUBNET_IS_NOT_FOUND" {
			sdkError.WithErrorCode(EcVServerSubnetNotFound).
				WithMessage(errMsg).
				WithErrors(perrResp.GetError())
		}
	}
}

func WithErrorInternalNetworkInterfaceNotFound(perrResp ErrorResponse) func(sdkError Error) {
	return func(sdkError Error) {
		if perrResp == nil {
			return
		}

		errMsg := strings.ToLower(strings.TrimSpace(perrResp.GetMessage()))
		if regexErrorInternalNetworkInterfaceNotFound.FindString(errMsg) != "" {
			sdkError.WithErrorCode(EcVServerInternalNetworkInterfaceNotFound).
				WithMessage(errMsg).
				WithErrors(perrResp.GetError())
		}
	}
}

func WithErrorAddressPairExisted(perrResp ErrorResponse) func(sdkError Error) {
	return func(sdkError Error) {
		if perrResp == nil {
			return
		}

		errMsg := strings.ToLower(strings.TrimSpace(perrResp.GetMessage()))
		if regexErrorAddressPairExisted.FindString(errMsg) != "" {
			sdkError.WithErrorCode(EcVServerAddressPairExisted).
				WithMessage(errMsg).
				WithErrors(perrResp.GetError())
		}
	}
}

func WithErrorWanIpAvailable(perrResp ErrorResponse) func(sdkError Error) {
	return func(sdkError Error) {
		if perrResp == nil {
			return
		}

		errMsg := perrResp.GetMessage()
		if strings.Contains(strings.ToLower(strings.TrimSpace(errMsg)), patternWanIpAvailable) {
			sdkError.WithErrorCode(EcVServerWanIpAvailable).
				WithMessage(errMsg).
				WithErrors(perrResp.GetError())
		}
	}
}

func WithErrorWanIdNotFound(perrResp ErrorResponse) func(sdkError Error) {
	return func(sdkError Error) {
		if perrResp == nil {
			return
		}

		errMsg := perrResp.GetMessage()
		if strings.Contains(strings.ToLower(strings.TrimSpace(errMsg)), pattermWapIdNotFound) {
			sdkError.WithErrorCode(EcVServerWanIdNotFound).
				WithMessage(errMsg).
				WithErrors(perrResp.GetError())
		}
	}
}
