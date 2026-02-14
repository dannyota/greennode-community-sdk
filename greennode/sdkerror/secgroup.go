package sdkerror

import "strings"

const (
	patternSecgroupNotFound          = "cannot get security group with id"
	patternSecgroupNameAlreadyExists = "name of security group already exist"
	patternSecgroupExceedQuota       = "exceeded secgroup quota"
	patternSecgroupInUse             = "securitygroupinuse"
)

func WithErrorSecgroupNotFound(perrResp ErrorResponse) func(sdkError Error) {
	return func(sdkError Error) {
		if perrResp == nil {
			return
		}

		errMsg := perrResp.GetMessage()
		if strings.Contains(strings.ToLower(strings.TrimSpace(errMsg)), patternSecgroupNotFound) {
			sdkError.WithErrorCode(EcVServerSecgroupNotFound).
				WithMessage(errMsg).
				WithErrors(perrResp.GetError())
		}
	}
}

func WithErrorSecgroupNameAlreadyExists(perrResp ErrorResponse) func(sdkError Error) {
	return func(sdkError Error) {
		if perrResp == nil {
			return
		}

		errMsg := perrResp.GetMessage()
		if strings.Contains(strings.ToLower(strings.TrimSpace(errMsg)), patternSecgroupNameAlreadyExists) {
			sdkError.WithErrorCode(EcVServerSecgroupNameAlreadyExists).
				WithMessage(errMsg).
				WithErrors(perrResp.GetError())
		}
	}
}

func WithErrorSecgroupExceedQuota(perrResp ErrorResponse) func(sdkError Error) {
	return func(sdkError Error) {
		if perrResp == nil {
			return
		}

		errMsg := perrResp.GetMessage()
		if strings.Contains(strings.ToLower(strings.TrimSpace(errMsg)), patternSecgroupExceedQuota) {
			sdkError.WithErrorCode(EcVServerSecgroupExceedQuota).
				WithMessage(errMsg).
				WithErrors(perrResp.GetError()).
				WithErrorCategories(ErrCatQuota)
		}
	}
}

func WithErrorSecgroupInUse(perrResp ErrorResponse) func(sdkError Error) {
	return func(sdkError Error) {
		if perrResp == nil {
			return
		}

		errMsg := perrResp.GetMessage()
		if strings.Contains(strings.ToLower(strings.TrimSpace(errMsg)), patternSecgroupInUse) {
			sdkError.WithErrorCode(EcVServerSecgroupInUse).
				WithMessage(errMsg).
				WithErrors(perrResp.GetError())
		}
	}
}
