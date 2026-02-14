package sdk_error

import "strings"

const (
	patternSecgroupRuleNotFound    = "cannot get security group rule with id"
	patternSecgroupRuleExists      = "securitygroupruleexists"
	patternSecgroupRuleExceedQuota = "exceeded secgroup_rule quota"
)

func WithErrorSecgroupRuleNotFound(perrResp IErrorResponse) func(sdkError IError) {
	return func(sdkError IError) {
		if perrResp == nil {
			return
		}

		errMsg := perrResp.GetMessage()
		if strings.Contains(strings.ToLower(strings.TrimSpace(errMsg)), patternSecgroupRuleNotFound) {
			sdkError.WithErrorCode(EcVServerSecgroupRuleNotFound).
				WithMessage(errMsg).
				WithErrors(perrResp.GetError())
		}
	}
}

func WithErrorSecgroupRuleAlreadyExists(perrResp IErrorResponse) func(sdkError IError) {
	return func(sdkError IError) {
		if perrResp == nil {
			return
		}

		errMsg := perrResp.GetMessage()
		if strings.Contains(strings.ToLower(strings.TrimSpace(errMsg)), patternSecgroupRuleExists) {
			sdkError.WithErrorCode(EcVServerSecgroupRuleAlreadyExists).
				WithMessage(errMsg).
				WithErrors(perrResp.GetError())
		}
	}
}

func WithErrorSecgroupRuleExceedQuota(perrResp IErrorResponse) func(sdkError IError) {
	return func(sdkError IError) {
		if perrResp == nil {
			return
		}

		errMsg := perrResp.GetMessage()
		if strings.Contains(strings.ToLower(strings.TrimSpace(errMsg)), patternSecgroupRuleExceedQuota) {
			sdkError.WithErrorCode(EcVServerSecgroupRuleExceedQuota).
				WithMessage(errMsg).
				WithErrors(perrResp.GetError()).
				WithErrorCategories(ErrCatQuota)
		}
	}
}
