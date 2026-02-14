package sdkerror

import (
	"regexp"
	"strings"
)

const (
	patternLoadBalancerNotBelongToProject  = "does not belong to project"
	patternLoadBalancerNotFound            = "cannot get load balancer with id" // "Cannot get load balancer with id"
	patternLoadBalancerDuplicatePoolName   = "duplicated pool name"
	patternLoadBalancerNotFound2           = "could not find resource"
	patternListenerDuplicateName           = "duplicated listener name"
	patternListenerNotFound                = "cannot get listener with id"
	patternListenerNotBelongToLoadBalancer = `listener id [^.]+ is not belong to load balancer id [^.]+`
	patternListenerDuplicateProtocolOrPort = "duplicated listener protocol port"
	patternPoolNotFound                    = "cannot get pool with id"
	patternPoolInUse                       = "is used in listener"
	patternLoadBalancerNotReady            = `(?:the )?load balancer id [^.]+ is not ready`
	patternListenerNotReady                = `listener id [^.]+ is not ready`
	patternMemberMustIdentical             = "the members provided are identical to the existing members in the pool"
	patternPoolIsUpdating                  = `pool id [^.]+ is updating`
	patternLoadBalancerExceedQuota         = "exceeded load_balancer quota. current used"
	patternLoadBalancerIsDeleting          = `load balancer id [^.]+ is deleting`
	patternLoadBalancerIsCreating          = `load balancer id [^.]+ is creating`
	patternLoadBalancerResizeSamePackage   = "is the same as the current package"
	patternLoadbalancerPackageNotFound     = "invalid package id"
	patternLoadBalancerIsUpdating          = `load balancer id [^.]+ is updating`
)

var (
	regexErrorLoadBalancerNotReady            = regexp.MustCompile(patternLoadBalancerNotReady)
	regexErrorListenerNotReady                = regexp.MustCompile(patternListenerNotReady)
	regexErrorPoolIsUpdating                  = regexp.MustCompile(patternPoolIsUpdating)
	regexErrorLoadBalancerIsDeleting          = regexp.MustCompile(patternLoadBalancerIsDeleting)
	regexErrorLoadBalancerIsCreating          = regexp.MustCompile(patternLoadBalancerIsCreating)
	regexErrorLoadBalancerIsUpdating          = regexp.MustCompile(patternLoadBalancerIsUpdating)
	regexErrorListenerNotBelongToLoadBalancer = regexp.MustCompile(patternListenerNotBelongToLoadBalancer)
)

func WithErrorLoadBalancerNotFound(errResp ErrorResponse) func(sdkError Error) {
	return func(sdkError Error) {
		if errResp == nil {
			return
		}

		errMsg := errResp.GetMessage()
		if strings.Contains(strings.ToLower(strings.TrimSpace(errMsg)), patternLoadBalancerNotBelongToProject) ||
			strings.Contains(strings.ToLower(strings.TrimSpace(errMsg)), patternLoadBalancerNotFound) {
			sdkError.WithErrorCode(EcVLBLoadBalancerNotFound).
				WithMessage(errMsg).
				WithErrors(errResp.GetError())
		}
	}
}

// WithErrorLoadBalancerNotFound2 indicate the issue creating Pool with non-existed LoadBalancer
func WithErrorLoadBalancerNotFound2(errResp ErrorResponse) func(sdkError Error) {
	return func(sdkError Error) {
		if errResp == nil {
			return
		}

		errMsg := errResp.GetMessage()
		if strings.Contains(strings.ToLower(strings.TrimSpace(errMsg)), patternLoadBalancerNotFound2) {
			sdkError.WithErrorCode(EcVLBLoadBalancerNotFound).
				WithMessage(errMsg).
				WithErrors(errResp.GetError())
		}
	}
}

func WithErrorLoadBalancerExceedQuota(errResp ErrorResponse) func(sdkError Error) {
	return func(sdkError Error) {
		if errResp == nil {
			return
		}

		errMsg := errResp.GetMessage()
		if strings.Contains(strings.ToLower(strings.TrimSpace(errMsg)), patternLoadBalancerExceedQuota) {
			sdkError.WithErrorCode(EcVLBLoadBalancerExceedQuota).
				WithMessage(errMsg).
				WithErrors(errResp.GetError()).
				WithErrorCategories(ErrCatQuota)
		}
	}
}

func WithErrorLoadBalancerDuplicatePoolName(errResp ErrorResponse) func(sdkError Error) {
	return func(sdkError Error) {
		if errResp == nil {
			return
		}

		errMsg := errResp.GetMessage()
		if strings.Contains(strings.ToLower(strings.TrimSpace(errMsg)), patternLoadBalancerDuplicatePoolName) {
			sdkError.WithErrorCode(EcVLBLoadBalancerDuplicatePoolName).
				WithMessage(errMsg).
				WithErrors(errResp.GetError())
		}
	}
}

func WithErrorListenerDuplicateProtocolOrPort(errResp ErrorResponse) func(sdkError Error) {
	return func(sdkError Error) {
		if errResp == nil {
			return
		}

		errMsg := errResp.GetMessage()
		if strings.Contains(strings.ToLower(strings.TrimSpace(errMsg)), patternListenerDuplicateProtocolOrPort) {
			sdkError.WithErrorCode(EcVLBListenerDuplicateProtocolOrPort).
				WithMessage(errMsg).
				WithErrors(errResp.GetError())
		}
	}
}

func WithErrorListenerDuplicateName(errResp ErrorResponse) func(sdkError Error) {
	return func(sdkError Error) {
		if errResp == nil {
			return
		}

		errMsg := errResp.GetMessage()
		if strings.Contains(strings.ToLower(strings.TrimSpace(errMsg)), patternListenerDuplicateName) {
			sdkError.WithErrorCode(EcVLBListenerDuplicateName).
				WithMessage(errMsg).
				WithErrors(errResp.GetError())
		}
	}
}

func WithErrorPoolNotFound(errResp ErrorResponse) func(sdkError Error) {
	return func(sdkError Error) {
		if errResp == nil {
			return
		}

		errMsg := errResp.GetMessage()
		if strings.Contains(strings.ToLower(strings.TrimSpace(errMsg)), patternPoolNotFound) {
			sdkError.WithErrorCode(EcVLBPoolNotFound).
				WithMessage(errMsg).
				WithErrors(errResp.GetError())
		}
	}
}

func WithErrorPoolInUse(errResp ErrorResponse) func(sdkError Error) {
	return func(sdkError Error) {
		if errResp == nil {
			return
		}

		errMsg := errResp.GetMessage()
		if strings.Contains(strings.ToLower(strings.TrimSpace(errMsg)), patternPoolInUse) {
			sdkError.WithErrorCode(EcVLBPoolInUse).
				WithMessage(errMsg).
				WithErrors(errResp.GetError())
		}
	}
}

func WithErrorLoadBalancerNotReady(errResp ErrorResponse) func(sdkError Error) {
	return func(sdkError Error) {
		if errResp == nil {
			return
		}

		errMsg := strings.ToLower(strings.TrimSpace(errResp.GetMessage()))
		if regexErrorLoadBalancerNotReady.FindString(errMsg) != "" ||
			regexErrorListenerNotReady.FindString(errMsg) != "" ||
			regexErrorPoolIsUpdating.FindString(errMsg) != "" ||
			regexErrorLoadBalancerIsUpdating.FindString(errMsg) != "" {
			sdkError.WithErrorCode(EcVLBLoadBalancerNotReady).
				WithMessage(errMsg).
				WithErrors(errResp.GetError())
		}
	}
}

func WithErrorLoadBalancerIsDeleting(errResp ErrorResponse) func(sdkError Error) {
	return func(sdkError Error) {
		if errResp == nil {
			return
		}

		errMsg := strings.ToLower(strings.TrimSpace(errResp.GetMessage()))
		if regexErrorLoadBalancerIsDeleting.FindString(errMsg) != "" {
			sdkError.WithErrorCode(EcVLBLoadBalancerIsDeleting).
				WithMessage(errMsg).
				WithErrors(errResp.GetError())
		}
	}
}

func WithErrorLoadBalancerIsCreating(errResp ErrorResponse) func(sdkError Error) {
	return func(sdkError Error) {
		if errResp == nil {
			return
		}

		errMsg := strings.ToLower(strings.TrimSpace(errResp.GetMessage()))
		if regexErrorLoadBalancerIsCreating.FindString(errMsg) != "" {
			sdkError.WithErrorCode(EcVLBLoadBalancerIsCreating).
				WithMessage(errMsg).
				WithErrors(errResp.GetError())
		}
	}
}

func WithErrorLoadBalancerIsUpdating(errResp ErrorResponse) func(sdkError Error) {
	return func(sdkError Error) {
		if errResp == nil {
			return
		}

		errMsg := strings.ToLower(strings.TrimSpace(errResp.GetMessage()))
		if regexErrorLoadBalancerIsUpdating.FindString(errMsg) != "" {
			sdkError.WithErrorCode(EcVLBLoadBalancerIsUpdating).
				WithMessage(errMsg).
				WithErrors(errResp.GetError())
		}
	}
}

func WithErrorListenerNotFound(errResp ErrorResponse) func(sdkError Error) {
	return func(sdkError Error) {
		if errResp == nil {
			return
		}

		errMsg := strings.ToLower(strings.TrimSpace(errResp.GetMessage()))
		if strings.Contains(errMsg, patternListenerNotFound) ||
			regexErrorListenerNotBelongToLoadBalancer.FindString(errMsg) != "" {
			sdkError.WithErrorCode(EcVLBListenerNotFound).
				WithMessage(errMsg).
				WithErrors(errResp.GetError())
		}
	}
}

func WithErrorMemberMustIdentical(errResp ErrorResponse) func(sdkError Error) {
	return func(sdkError Error) {
		if errResp == nil {
			return
		}

		errMsg := errResp.GetMessage()
		if strings.Contains(strings.ToLower(strings.TrimSpace(errMsg)), patternMemberMustIdentical) {
			sdkError.WithErrorCode(EcVLBMemberMustIdentical).
				WithMessage(errMsg).
				WithErrors(errResp.GetError())
		}
	}
}

func WithErrorLoadBalancerResizeSamePackage(errResp ErrorResponse) func(sdkError Error) {
	return func(sdkError Error) {
		if errResp == nil {
			return
		}

		errMsg := errResp.GetMessage()
		if strings.Contains(strings.ToLower(strings.TrimSpace(errMsg)), patternLoadBalancerResizeSamePackage) {
			sdkError.WithErrorCode(EcVLBLoadBalancerResizeSamePackage).
				WithMessage(errMsg).
				WithErrors(errResp.GetError())
		}
	}
}

func WithErrorLoadBalancerPackageNotFound(errResp ErrorResponse) func(sdkError Error) {
	return func(sdkError Error) {
		if errResp == nil {
			return
		}

		errMsg := errResp.GetMessage()
		if strings.Contains(strings.ToLower(strings.TrimSpace(errMsg)), patternLoadbalancerPackageNotFound) {
			sdkError.WithErrorCode(EcVLBLoadBalancerPackageNotFound).
				WithMessage(errMsg).
				WithErrors(errResp.GetError())
		}
	}
}

func WithErrorGlobalLoadBalancerNotFound(errResp ErrorResponse) func(sdkError Error) {
	return func(sdkError Error) {
		if errResp == nil {
			return
		}

		errMsg := errResp.GetMessage()
		if strings.Contains(strings.ToLower(strings.TrimSpace(errMsg)), strings.ToLower("Global load balancer is not found")) {
			sdkError.WithErrorCode(EcGlobalLoadBalancerNotFound).
				WithMessage(errMsg).
				WithErrors(errResp.GetError())
		}
	}
}
