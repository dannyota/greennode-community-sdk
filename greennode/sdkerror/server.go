package sdkerror

import (
	"regexp"
	"strings"
)

const (
	patternServerNotFound                      = "cannot get server with id"                 // "Cannot get volume type with id vtype-6790f903-38d2-454d-919e-5b49184b5927"
	patternServerCreating                      = "cannot delete server with status creating" // "Server is creating"
	patternServerExceedQuota                   = "exceeded vm quota"                         // "The number of servers exceeds the quota"
	patternServerDeleting                      = "cannot delete server with status deleting" // "Server is deleting"
	patternServerBilling                       = "cannot delete server with status creating-billing"
	patternBillingPaymentMethodNotAllowed      = "payment method is not allowed for the user"
	patternServerAttachVolumeQuotaExceeded     = "exceeded volume_per_server quota"
	patternServerAttachEncryptedVolume         = "cannot attach encryption volume"
	patternServerExpired                       = "server is expired"
	patternServerFlavorSystemExceedQuota       = "there are no more remaining flavor with id"
	patternServerUpdatingSecgroups             = "cannot change security group of server with status changing-security-group"
	patternServerExceedCpuQuota                = "exceeded vcpu quota. current used"
	patternServerImageNotSupported             = "doesn't support image with id"
	patternImageNotSupport                     = "don't support image"
	patternServerCanNotAttachFloatingIp        = "the server only allows attaching 1 floating ip"
	patternServerFlavorNotSupported            = `flavor [^.]+ don't support image [^.]+`
	patternServerDeleteServerUpdatingSecgroups = "cannot delete server with status changing-security-group"
	patternServerExceedFloatingIpQuota         = "exceeded floating_ip quota"
	patternImageNotFound                       = "cannot get image with id"
	patternServerGroupNotFound                 = "not found server group"
	patternServerGroupInUse                    = "server group is in use"
	patternServerGroupNameMustBeUnique         = "name must be unique"
)

var (
	regexErrorServerFlavorNotSupported = regexp.MustCompile(patternServerFlavorNotSupported)
)

func WithErrorServerFlavorNotSupported(errResp ErrorResponse) func(sdkError Error) {
	return func(sdkError Error) {
		if errResp == nil {
			return
		}

		errMsg := strings.ToLower(strings.TrimSpace(errResp.GetMessage()))
		if regexErrorServerFlavorNotSupported.FindString(errMsg) != "" {
			sdkError.WithErrorCode(EcVServerFlavorNotSupported).
				WithMessage(errMsg).
				WithErrors(errResp.Err())
		}
	}
}

func WithErrorServerNotFound(errResp ErrorResponse) func(sdkError Error) {
	return func(sdkError Error) {
		if errResp == nil {
			return
		}

		errMsg := errResp.GetMessage()
		if strings.Contains(strings.ToLower(strings.TrimSpace(errMsg)), patternServerNotFound) {
			sdkError.WithErrorCode(EcVServerServerNotFound).
				WithMessage(errMsg).
				WithErrors(errResp.Err())
		}
	}
}

func WithErrorImageNotFound(errResp ErrorResponse) func(sdkError Error) {
	return func(sdkError Error) {
		if errResp == nil {
			return
		}

		errMsg := strings.ToLower(strings.TrimSpace(errResp.GetMessage()))
		if strings.Contains(errMsg, patternImageNotFound) {
			sdkError.WithErrorCode(EcVServerImageNotFound).
				WithMessage(errResp.GetMessage()).
				WithErrors(errResp.Err())
		}
	}
}

func WithErrorServerFlavorSystemExceedQuota(errResp ErrorResponse) func(sdkError Error) {
	return func(sdkError Error) {
		if errResp == nil {
			return
		}

		errMsg := errResp.GetMessage()
		if strings.Contains(strings.ToLower(strings.TrimSpace(errMsg)), patternServerFlavorSystemExceedQuota) {
			sdkError.WithErrorCode(EcVServerServerFlavorSystemExceedQuota).
				WithMessage(errMsg).
				WithErrors(errResp.Err()).
				WithErrorCategories(ErrCatInfra)
		}
	}
}

func WithErrorServerDeleteCreatingServer(errResp ErrorResponse) func(sdkError Error) {
	return func(sdkError Error) {
		if errResp == nil {
			return
		}

		errMsg := errResp.GetMessage()
		if strings.Contains(strings.ToLower(strings.TrimSpace(errMsg)), patternServerCreating) {
			sdkError.WithErrorCode(EcVServerServerDeleteCreatingServer).
				WithMessage(errMsg).
				WithErrors(errResp.Err())
		}
	}
}

func WithErrorServerExpired(errResp ErrorResponse) func(sdkError Error) {
	return func(sdkError Error) {
		if errResp == nil {
			return
		}

		errMsg := errResp.GetMessage()
		if strings.Contains(strings.ToLower(strings.TrimSpace(errMsg)), patternServerExpired) {
			sdkError.WithErrorCode(EcVServerServerExpired).
				WithMessage(errMsg).
				WithErrors(errResp.Err())
		}
	}
}

func WithErrorServerUpdatingSecgroups(errResp ErrorResponse) func(sdkError Error) {
	return func(sdkError Error) {
		if errResp == nil {
			return
		}

		errMsg := errResp.GetMessage()
		stdErrMsg := strings.ToLower(strings.TrimSpace(errMsg))
		if strings.Contains(stdErrMsg, patternServerUpdatingSecgroups) ||
			strings.Contains(stdErrMsg, patternServerDeleteServerUpdatingSecgroups) {
			sdkError.WithErrorCode(EcVServerServerUpdatingSecgroups).
				WithMessage(errMsg).
				WithErrors(errResp.Err())
		}
	}
}

func WithErrorServerExceedCpuQuota(errResp ErrorResponse) func(sdkError Error) {
	return func(sdkError Error) {
		if errResp == nil {
			return
		}

		errMsg := errResp.GetMessage()
		if strings.Contains(strings.ToLower(strings.TrimSpace(errMsg)), patternServerExceedCpuQuota) {
			sdkError.WithErrorCode(EcVServerServerExceedCpuQuota).
				WithMessage(errMsg).
				WithErrors(errResp.Err()).
				WithErrorCategories(ErrCatQuota)
		}
	}
}

func WithErrorServerExceedFloatingIpQuota(errResp ErrorResponse) func(sdkError Error) {
	return func(sdkError Error) {
		if errResp == nil {
			return
		}

		errMsg := errResp.GetMessage()
		if strings.Contains(strings.ToLower(strings.TrimSpace(errMsg)), patternServerExceedFloatingIpQuota) {
			sdkError.WithErrorCode(EcVServerServerExceedFloatingIpQuota).
				WithMessage(errMsg).
				WithErrors(errResp.Err()).
				WithErrorCategories(ErrCatQuota)
		}
	}
}

func WithErrorServerImageNotSupported(errResp ErrorResponse) func(sdkError Error) {
	return func(sdkError Error) {
		if errResp == nil {
			return
		}

		errMsg := errResp.GetMessage()
		lowerErrMsg := strings.ToLower(strings.TrimSpace(errMsg))
		if strings.Contains(lowerErrMsg, patternServerImageNotSupported) ||
			strings.Contains(lowerErrMsg, patternImageNotSupport) {
			sdkError.WithErrorCode(EcVServerServerImageNotSupported).
				WithMessage(errMsg).
				WithErrors(errResp.Err())
		}
	}
}

func WithErrorServerExceedQuota(errResp ErrorResponse) func(sdkError Error) {
	return func(sdkError Error) {
		if errResp == nil {
			return
		}

		errMsg := errResp.GetMessage()
		if strings.Contains(strings.ToLower(strings.TrimSpace(errMsg)), patternServerExceedQuota) {
			sdkError.WithErrorCode(EcVServerServerExceedQuota).
				WithMessage(errMsg).
				WithErrors(errResp.Err()).
				WithErrorCategories(ErrCatQuota)
		}
	}
}

func WithErrorServerDeleteDeletingServer(errResp ErrorResponse) func(sdkError Error) {
	return func(sdkError Error) {
		if errResp == nil {
			return
		}

		errMsg := errResp.GetMessage()
		if strings.Contains(strings.ToLower(strings.TrimSpace(errMsg)), patternServerDeleting) {
			sdkError.WithErrorCode(EcVServerServerDeleteDeletingServer).
				WithMessage(errMsg).
				WithErrors(errResp.Err())
		}
	}
}

func WithErrorServerDeleteBillingServer(errResp ErrorResponse) func(sdkError Error) {
	return func(sdkError Error) {
		if errResp == nil {
			return
		}

		errMsg := errResp.GetMessage()
		if strings.Contains(strings.ToLower(strings.TrimSpace(errMsg)), patternServerBilling) {
			sdkError.WithErrorCode(EcVServerServerDeleteBillingServer).
				WithMessage(errMsg).
				WithErrors(errResp.Err())
		}
	}
}

func WithErrorServerCreateBillingPaymentMethodNotAllowed(errResp ErrorResponse) func(sdkError Error) {
	return func(sdkError Error) {
		if errResp == nil {
			return
		}

		errMsg := strings.ToLower(strings.TrimSpace(errResp.GetMessage()))
		if strings.Contains(errMsg, patternBillingPaymentMethodNotAllowed) {
			sdkError.WithErrorCode(EcVServerCreateBillingPaymentMethodNotAllowed).
				WithMessage(errResp.GetMessage()).
				WithErrors(errResp.Err())
		}
	}
}

func WithErrorServerAttachVolumeQuotaExceeded(errResp ErrorResponse) func(sdkError Error) {
	return func(sdkError Error) {
		if errResp == nil {
			return
		}

		errMsg := errResp.GetMessage()
		if strings.Contains(strings.ToLower(strings.TrimSpace(errMsg)), patternServerAttachVolumeQuotaExceeded) {
			sdkError.WithErrorCode(EcVServerServerVolumeAttachQuotaExceeded).
				WithMessage(errMsg).
				WithErrors(errResp.Err()).
				WithErrorCategories(ErrCatQuota)
		}
	}
}

func WithErrorServerAttachEncryptedVolume(errResp ErrorResponse) func(sdkError Error) {
	return func(sdkError Error) {
		if errResp == nil {
			return
		}

		errMsg := errResp.GetMessage()
		if strings.Contains(strings.ToLower(strings.TrimSpace(errMsg)), patternServerAttachEncryptedVolume) {
			sdkError.WithErrorCode(EcVServerServerAttachEncryptedVolume).
				WithMessage(errMsg).
				WithErrors(errResp.Err())
		}
	}
}

func WithErrorServerCanNotAttachFloatingIp(errResp ErrorResponse) func(sdkError Error) {
	return func(sdkError Error) {
		if errResp == nil {
			return
		}

		errMsg := errResp.GetMessage()
		if strings.Contains(strings.ToLower(strings.TrimSpace(errMsg)), patternServerCanNotAttachFloatingIp) {
			sdkError.WithErrorCode(EcVServerServerCanNotAttachFloatingIp).
				WithMessage(errMsg).
				WithErrors(errResp.Err())
		}
	}
}

func WithErrorServerGroupNotFound(errResp ErrorResponse) func(sdkError Error) {
	return func(sdkError Error) {
		if errResp == nil {
			return
		}

		errMsg := errResp.GetMessage()
		if strings.Contains(strings.ToLower(strings.TrimSpace(errMsg)), patternServerGroupNotFound) {
			sdkError.WithErrorCode(EcVServerServerGroupNotFound).
				WithMessage(errMsg).
				WithErrors(errResp.Err())
		}
	}
}

func WithErrorServerGroupInUse(errResp ErrorResponse) func(sdkError Error) {
	return func(sdkError Error) {
		if errResp == nil {
			return
		}

		errMsg := errResp.GetMessage()
		if strings.Contains(strings.ToLower(strings.TrimSpace(errMsg)), patternServerGroupInUse) {
			sdkError.WithErrorCode(EcVServerServerGroupInUse).
				WithMessage(errMsg).
				WithErrors(errResp.Err())
		}
	}
}

func WithErrorServerGroupNameMustBeUnique(errResp ErrorResponse) func(sdkError Error) {
	return func(sdkError Error) {
		if errResp == nil {
			return
		}

		errMsg := errResp.GetMessage()
		if strings.Contains(strings.ToLower(strings.TrimSpace(errMsg)), patternServerGroupNameMustBeUnique) {
			sdkError.WithErrorCode(EcVServerServerGroupNameMustBeUnique).
				WithMessage(errMsg).
				WithErrors(errResp.Err())
		}
	}
}
