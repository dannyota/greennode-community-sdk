package sdkerror

import "strings"

const (
	patternSecgroupNotFound          = "cannot get security group with id"
	patternSecgroupNameAlreadyExists = "name of security group already exist"
	patternSecgroupExceedQuota       = "exceeded secgroup quota"
	patternSecgroupInUse             = "securitygroupinuse"
)

func WithErrorSecgroupNotFound(errResp ErrorResponse) func(sdkError Error) {
	return func(sdkError Error) {
		if errResp == nil {
			return
		}

		errMsg := errResp.GetMessage()
		if strings.Contains(strings.ToLower(strings.TrimSpace(errMsg)), patternSecgroupNotFound) {
			sdkError.WithErrorCode(EcVServerSecgroupNotFound).
				WithMessage(errMsg).
				WithErrors(errResp.GetError())
		}
	}
}

func WithErrorSecgroupNameAlreadyExists(errResp ErrorResponse) func(sdkError Error) {
	return func(sdkError Error) {
		if errResp == nil {
			return
		}

		errMsg := errResp.GetMessage()
		if strings.Contains(strings.ToLower(strings.TrimSpace(errMsg)), patternSecgroupNameAlreadyExists) {
			sdkError.WithErrorCode(EcVServerSecgroupNameAlreadyExists).
				WithMessage(errMsg).
				WithErrors(errResp.GetError())
		}
	}
}

func WithErrorSecgroupExceedQuota(errResp ErrorResponse) func(sdkError Error) {
	return func(sdkError Error) {
		if errResp == nil {
			return
		}

		errMsg := errResp.GetMessage()
		if strings.Contains(strings.ToLower(strings.TrimSpace(errMsg)), patternSecgroupExceedQuota) {
			sdkError.WithErrorCode(EcVServerSecgroupExceedQuota).
				WithMessage(errMsg).
				WithErrors(errResp.GetError()).
				WithErrorCategories(ErrCatQuota)
		}
	}
}

func WithErrorSecgroupInUse(errResp ErrorResponse) func(sdkError Error) {
	return func(sdkError Error) {
		if errResp == nil {
			return
		}

		errMsg := errResp.GetMessage()
		if strings.Contains(strings.ToLower(strings.TrimSpace(errMsg)), patternSecgroupInUse) {
			sdkError.WithErrorCode(EcVServerSecgroupInUse).
				WithMessage(errMsg).
				WithErrors(errResp.GetError())
		}
	}
}
