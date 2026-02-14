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

func WithErrorNetworkNotFound(errResp ErrorResponse) func(sdkError Error) {
	return func(sdkError Error) {
		if errResp == nil {
			return
		}

		errMsg := errResp.GetMessage()
		if strings.Contains(strings.ToLower(strings.TrimSpace(errMsg)), patternNetworkNotFound) ||
			strings.ToUpper(strings.TrimSpace(errResp.Err().Error())) == "VPC_IS_NOT_FOUND" {
			sdkError.WithErrorCode(EcVServerNetworkNotFound).
				WithMessage(errMsg).
				WithErrors(errResp.Err())
		}
	}
}

func WithErrorSubnetNotBelongNetwork(errResp ErrorResponse) func(sdkError Error) {
	return func(sdkError Error) {
		if errResp == nil {
			return
		}

		errMsg := strings.ToLower(strings.TrimSpace(errResp.GetMessage()))
		if regexErrorSubnetNotBelongNetwork.FindString(errMsg) != "" {
			sdkError.WithErrorCode(EcVServerSubnetNotBelongNetwork).
				WithMessage(errMsg).
				WithErrors(errResp.Err())
		}
	}
}

func WithErrorSubnetNotFound(errResp ErrorResponse) func(sdkError Error) {
	return func(sdkError Error) {
		if errResp == nil {
			return
		}

		errMsg := strings.ToLower(strings.TrimSpace(errResp.GetMessage()))
		if regexErrorSubnetNotFound.FindString(errMsg) != "" ||
			regexErrorSubnetNotFound2.FindString(errMsg) != "" ||
			strings.ToUpper(strings.TrimSpace(errResp.Err().Error())) == "SUBNET_IS_NOT_FOUND" {
			sdkError.WithErrorCode(EcVServerSubnetNotFound).
				WithMessage(errMsg).
				WithErrors(errResp.Err())
		}
	}
}

func WithErrorInternalNetworkInterfaceNotFound(errResp ErrorResponse) func(sdkError Error) {
	return func(sdkError Error) {
		if errResp == nil {
			return
		}

		errMsg := strings.ToLower(strings.TrimSpace(errResp.GetMessage()))
		if regexErrorInternalNetworkInterfaceNotFound.FindString(errMsg) != "" {
			sdkError.WithErrorCode(EcVServerInternalNetworkInterfaceNotFound).
				WithMessage(errMsg).
				WithErrors(errResp.Err())
		}
	}
}

func WithErrorAddressPairExisted(errResp ErrorResponse) func(sdkError Error) {
	return func(sdkError Error) {
		if errResp == nil {
			return
		}

		errMsg := strings.ToLower(strings.TrimSpace(errResp.GetMessage()))
		if regexErrorAddressPairExisted.FindString(errMsg) != "" {
			sdkError.WithErrorCode(EcVServerAddressPairExisted).
				WithMessage(errMsg).
				WithErrors(errResp.Err())
		}
	}
}

func WithErrorWanIpAvailable(errResp ErrorResponse) func(sdkError Error) {
	return func(sdkError Error) {
		if errResp == nil {
			return
		}

		errMsg := errResp.GetMessage()
		if strings.Contains(strings.ToLower(strings.TrimSpace(errMsg)), patternWanIpAvailable) {
			sdkError.WithErrorCode(EcVServerWanIpAvailable).
				WithMessage(errMsg).
				WithErrors(errResp.Err())
		}
	}
}

func WithErrorWanIDNotFound(errResp ErrorResponse) func(sdkError Error) {
	return func(sdkError Error) {
		if errResp == nil {
			return
		}

		errMsg := errResp.GetMessage()
		if strings.Contains(strings.ToLower(strings.TrimSpace(errMsg)), pattermWapIDNotFound) {
			sdkError.WithErrorCode(EcVServerWanIDNotFound).
				WithMessage(errMsg).
				WithErrors(errResp.Err())
		}
	}
}
