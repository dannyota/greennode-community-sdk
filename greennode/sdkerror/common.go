package sdkerror

import (
	"fmt"
	"net/http"
	"regexp"
	"strings"
)

const (
	patternPurchaseIssue      = "you do not have sufficient credits to complete the purchase"
	patternPagingInvalid      = "page or size invalid"
	patternTagKeyInvalid      = "the value for the tag key contains illegal characters"
	patternServiceMaintenance = "this service is in maintenance"
	patternProjectConflict    = `project [^.]+ is not belong to user`
)

var (
	regexErrorProjectConflict = regexp.MustCompile(patternProjectConflict)
)

func init() {
	register(EcPurchaseIssue, &classifier{
		match: containsAny(patternPurchaseIssue), category: ErrCatPurchase,
	})
	register(EcTagKeyInvalid, &classifier{match: containsAny(patternTagKeyInvalid)})
	register(EcPagingInvalid, &classifier{match: containsAny(patternPagingInvalid)})
	register(EcPaymentMethodNotAllow, &classifier{match: containsAny("ext_pm_payment_method_not_allow")})
	register(EcCreditNotEnough, &classifier{match: containsAny("ext_pm_credit_not_enough")})
	register(EcProjectConflict, &classifier{match: matchRegexps(regexErrorProjectConflict)})
}

func ErrorHandler(err error) *SdkError {
	sdkErr := &SdkError{
		error:     err,
		errorCode: EcUnknownError,
		message:   "Unknown error",
	}

	if err != nil && strings.Contains(strings.ToLower(strings.TrimSpace(err.Error())), patternServiceMaintenance) {
		sdkErr.errorCode = EcServiceMaintenance
		sdkErr.message = "Service Maintenance"
		sdkErr.error = fmt.Errorf("service is under maintenance")
	}

	return sdkErr
}

func SdkErrorHandler(err error, errResp ErrorResponse, codes ...ErrorCode) *SdkError {
	if err == nil && errResp == nil {
		return nil
	}

	var sdkErr *SdkError
	if err != nil {
		if e, ok := err.(*SdkError); ok {
			sdkErr = e
		} else {
			sdkErr = ErrorHandler(err)
		}
	} else {
		sdkErr = ErrorHandler(nil)
	}

	if sdkErr.errorCode != EcUnknownError {
		return sdkErr
	}

	if errResp != nil {
		sdkErr.WithErrorCode(EcUnknownError).WithMessage(errResp.GetMessage()).WithErrors(errResp.Err())
	}

	if errResp == nil {
		return sdkErr
	}

	lowerMsg := strings.ToLower(strings.TrimSpace(errResp.GetMessage()))

	for _, code := range codes {
		c, ok := classifierRegistry[code]
		if !ok {
			continue
		}

		if c.catchAll {
			if sdkErr.errorCode != EcUnknownError {
				continue
			}
			sdkErr.WithErrorCode(code).WithMessage(errResp.GetMessage()).WithErrors(errResp.Err())
			if c.category != "" {
				sdkErr.AppendCategories(c.category)
			}
			return sdkErr
		}

		if c.match(lowerMsg, errResp) {
			msg := errResp.GetMessage()
			if c.msgFmt != "" {
				msg = fmt.Sprintf(c.msgFmt, errResp.GetMessage())
			}
			sdkErr.WithErrorCode(code).WithMessage(msg).WithErrors(errResp.Err())
			if c.category != "" {
				sdkErr.AppendCategories(c.category)
			}
			return sdkErr
		}
	}

	return sdkErr
}

func NewInternalServerError() *SdkError {
	return &SdkError{
		errorCode: EcInternalServerError,
		message:   "Internal Server Error",
		error:     fmt.Errorf("internal server error from making request to external service"),
	}
}

func NewServiceMaintenance() *SdkError {
	return &SdkError{
		errorCode: EcServiceMaintenance,
		message:   "Service Maintenance",
		error:     fmt.Errorf("service is under maintenance"),
	}
}

func NewPermissionDenied() *SdkError {
	return &SdkError{
		errorCode: EcPermissionDenied,
		message:   "Permission Denied",
		error:     fmt.Errorf("permission denied when making request to external service"),
	}
}

func NewReauthFuncNotSet() *SdkError {
	return &SdkError{
		errorCode: EcReauthFuncNotSet,
		message:   "Reauthentication function is not configured",
		error:     fmt.Errorf("reauthentication function is not configured"),
	}
}

func NewUnexpectedError(response *http.Response) *SdkError {
	statusCode := 0
	url := ""
	if response != nil {
		statusCode = response.StatusCode

		if response.Request != nil && response.Request.URL != nil {
			url = response.Request.URL.String()
		}
	}

	sdkErr := &SdkError{
		errorCode: EcUnexpectedError,
		message:   "Unexpected Error",
		error:     fmt.Errorf("unexpected error from making request to external service"),
	}
	sdkErr.WithParameters(map[string]any{
		"statusCode": statusCode,
		"url":        url,
	})
	return sdkErr
}

func NewQuotaNotFound() *SdkError {
	return &SdkError{
		errorCode: EcVServerQuotaNotFound,
		message:   "Quota not found",
		error:     fmt.Errorf("quota not found"),
	}
}
