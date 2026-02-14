package sdkerror

import "strings"

const (
	patternSecgroupRuleNotFound    = "cannot get security group rule with id"
	patternSecgroupRuleExists      = "securitygroupruleexists"
	patternSecgroupRuleExceedQuota = "exceeded secgroup_rule quota"
)

func WithErrorSecgroupRuleNotFound(errResp ErrorResponse) func(sdkError Error) {
	return func(sdkError Error) {
		if errResp == nil {
			return
		}

		errMsg := errResp.GetMessage()
		if strings.Contains(strings.ToLower(strings.TrimSpace(errMsg)), patternSecgroupRuleNotFound) {
			sdkError.WithErrorCode(EcVServerSecgroupRuleNotFound).
				WithMessage(errMsg).
				WithErrors(errResp.Err())
		}
	}
}

func WithErrorSecgroupRuleAlreadyExists(errResp ErrorResponse) func(sdkError Error) {
	return func(sdkError Error) {
		if errResp == nil {
			return
		}

		errMsg := errResp.GetMessage()
		if strings.Contains(strings.ToLower(strings.TrimSpace(errMsg)), patternSecgroupRuleExists) {
			sdkError.WithErrorCode(EcVServerSecgroupRuleAlreadyExists).
				WithMessage(errMsg).
				WithErrors(errResp.Err())
		}
	}
}

func WithErrorSecgroupRuleExceedQuota(errResp ErrorResponse) func(sdkError Error) {
	return func(sdkError Error) {
		if errResp == nil {
			return
		}

		errMsg := errResp.GetMessage()
		if strings.Contains(strings.ToLower(strings.TrimSpace(errMsg)), patternSecgroupRuleExceedQuota) {
			sdkError.WithErrorCode(EcVServerSecgroupRuleExceedQuota).
				WithMessage(errMsg).
				WithErrors(errResp.Err()).
				WithErrorCategories(ErrCatQuota)
		}
	}
}
