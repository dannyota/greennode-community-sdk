package sdkerror

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/imroc/req/v3"
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

func ErrorHandler(err error, opts ...func(sdkErr Error)) Error {
	sdkErr := &SdkError{
		error:     err,
		errorCode: EcUnknownError,
		message:   "Unknown error",
	}

	if err != nil && strings.Contains(strings.ToLower(strings.TrimSpace(err.Error())), patternServiceMaintenance) {
		sdkErr.errorCode = EcServiceMaintenance
		sdkErr.message = "Service Maintenance"
		sdkErr.error = fmt.Errorf("service is under maintenance")

		return sdkErr
	}

	for _, opt := range opts {
		opt(sdkErr)
		if sdkErr.errorCode != EcUnknownError {
			return sdkErr
		}
	}

	sdkErr.error = err
	sdkErr.message = ""

	return sdkErr
}

func SdkErrorHandler(sdkErr Error, errResp ErrorResponse, opts ...func(sdkErr Error)) Error {
	if sdkErr == nil && errResp == nil {
		return nil
	}

	if sdkErr != nil && sdkErr.ErrorCode() != EcUnknownError {
		return sdkErr
	}

	// Fill the default error
	if errResp != nil {
		sdkErr.WithErrorCode(EcUnknownError).WithMessage(errResp.GetMessage()).WithErrors(errResp.Err())
	}

	for _, opt := range opts {
		opt(sdkErr)
		if sdkErr.ErrorCode() != EcUnknownError {
			return sdkErr
		}
	}

	return sdkErr
}

func WithErrorInternalServerError() func(Error) {
	return func(sdkErr Error) {
		sdkErr.WithErrorCode(EcInternalServerError).
			WithMessage("Internal Server Error").
			WithErrors(fmt.Errorf("internal server error from making request to external service"))
	}
}

func WithErrorServiceMaintenance() func(Error) {
	return func(sdkErr Error) {
		sdkErr.WithErrorCode(EcServiceMaintenance).
			WithMessage("Service Maintenance").
			WithErrors(fmt.Errorf("service is under maintenance"))
	}
}

func WithErrorPermissionDenied() func(Error) {
	return func(sdkErr Error) {
		sdkErr.WithErrorCode(EcPermissionDenied).
			WithMessage("Permission Denied").
			WithErrors(fmt.Errorf("permission denied when making request to external service"))
	}
}

func WithErrorPurchaseIssue(errResp ErrorResponse) func(sdkError Error) {
	return func(sdkError Error) {
		if errResp == nil {
			return
		}

		errMsg := errResp.GetMessage()
		if strings.Contains(strings.ToLower(strings.TrimSpace(errMsg)), patternPurchaseIssue) {
			sdkError.WithErrorCode(EcPurchaseIssue).
				WithMessage(errMsg).
				WithErrors(errResp.Err()).
				WithErrorCategories(ErrCatPurchase)
		}
	}
}

func WithErrorTagKeyInvalid(errResp ErrorResponse) func(sdkError Error) {
	return func(sdkError Error) {
		if errResp == nil {
			return
		}

		errMsg := errResp.GetMessage()
		if strings.Contains(strings.ToLower(strings.TrimSpace(errMsg)), patternTagKeyInvalid) {
			sdkError.WithErrorCode(EcTagKeyInvalid).
				WithMessage(errMsg).
				WithErrors(errResp.Err())
		}
	}
}

func WithErrorPagingInvalid(errResp ErrorResponse) func(sdkError Error) {
	return func(sdkError Error) {
		if errResp == nil {
			return
		}

		errMsg := errResp.GetMessage()
		if strings.Contains(strings.ToLower(strings.TrimSpace(errMsg)), patternPagingInvalid) {
			sdkError.WithErrorCode(EcPagingInvalid).
				WithMessage(errMsg).
				WithErrors(errResp.Err())
		}
	}
}

func WithErrorUnexpected(response *req.Response) func(Error) {
	statusCode := 0
	url := ""
	err := fmt.Errorf("unexpected error from making request to external service")
	if response != nil {
		if response.Response != nil && response.StatusCode != 0 {
			statusCode = response.StatusCode
		}

		if response.Request != nil && response.Request.URL != nil {
			url = response.Request.URL.String()
		}

		if response.Err != nil {
			err = response.Err
		}
	}

	return func(sdkErr Error) {
		sdkErr.WithErrorCode(EcUnexpectedError).
			WithMessage("Unexpected Error").
			WithErrors(err).
			WithParameters(map[string]any{
				"statusCode": statusCode,
				"url":        url,
			})
	}
}

func WithErrorPaymentMethodNotAllow(errResp ErrorResponse) func(sdkError Error) {
	return func(sdkError Error) {
		if errResp == nil {
			return
		}

		errMsg := strings.ToLower(strings.TrimSpace(errResp.GetMessage()))
		if strings.Contains(errMsg, "ext_pm_payment_method_not_allow") {
			sdkError.WithErrorCode(EcPaymentMethodNotAllow).
				WithMessage(errResp.GetMessage()).
				WithErrors(errResp.Err())
		}
	}
}

func WithErrorCreditNotEnough(errResp ErrorResponse) func(sdkError Error) {
	return func(sdkError Error) {
		if errResp == nil {
			return
		}

		errMsg := strings.ToLower(strings.TrimSpace(errResp.GetMessage()))
		if strings.Contains(errMsg, "ext_pm_credit_not_enough") {
			sdkError.WithErrorCode(EcCreditNotEnough).
				WithMessage(errResp.GetMessage()).
				WithErrors(errResp.Err())
		}
	}
}

func WithErrorProjectConflict(errResp ErrorResponse) func(sdkError Error) {
	return func(sdkError Error) {
		if errResp == nil {
			return
		}

		errMsg := strings.ToLower(strings.TrimSpace(errResp.GetMessage()))
		if regexErrorProjectConflict.FindString(errMsg) != "" {
			sdkError.WithErrorCode(EcProjectConflict).
				WithMessage(errMsg).
				WithErrors(errResp.Err())
		}
	}
}
