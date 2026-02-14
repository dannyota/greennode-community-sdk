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

func ErrorHandler(perr error, popts ...func(psdkErr Error)) Error {
	sdkErr := &SdkError{
		error:     perr,
		errorCode: EcUnknownError,
		message:   "Unknown error",
	}

	if perr != nil && strings.Contains(strings.ToLower(strings.TrimSpace(perr.Error())), patternServiceMaintenance) {
		sdkErr.errorCode = EcServiceMaintenance
		sdkErr.message = "Service Maintenance"
		sdkErr.error = fmt.Errorf("service is under maintenance")

		return sdkErr
	}

	for _, opt := range popts {
		opt(sdkErr)
		if sdkErr.errorCode != EcUnknownError {
			return sdkErr
		}
	}

	sdkErr.error = perr
	sdkErr.message = ""

	return sdkErr
}

func SdkErrorHandler(psdkErr Error, perrResp ErrorResponse, popts ...func(psdkErr Error)) Error {
	if psdkErr == nil && perrResp == nil {
		return nil
	}

	if psdkErr != nil && psdkErr.GetErrorCode() != EcUnknownError {
		return psdkErr
	}

	// Fill the default error
	if perrResp != nil {
		psdkErr.WithErrorCode(EcUnknownError).WithMessage(perrResp.GetMessage()).WithErrors(perrResp.GetError())
	}

	for _, opt := range popts {
		opt(psdkErr)
		if psdkErr.GetErrorCode() != EcUnknownError {
			return psdkErr
		}
	}

	return psdkErr
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

func WithErrorPurchaseIssue(perrResp ErrorResponse) func(sdkError Error) {
	return func(sdkError Error) {
		if perrResp == nil {
			return
		}

		errMsg := perrResp.GetMessage()
		if strings.Contains(strings.ToLower(strings.TrimSpace(errMsg)), patternPurchaseIssue) {
			sdkError.WithErrorCode(EcPurchaseIssue).
				WithMessage(errMsg).
				WithErrors(perrResp.GetError()).
				WithErrorCategories(ErrCatPurchase)
		}
	}
}

func WithErrorTagKeyInvalid(perrResp ErrorResponse) func(sdkError Error) {
	return func(sdkError Error) {
		if perrResp == nil {
			return
		}

		errMsg := perrResp.GetMessage()
		if strings.Contains(strings.ToLower(strings.TrimSpace(errMsg)), patternTagKeyInvalid) {
			sdkError.WithErrorCode(EcTagKeyInvalid).
				WithMessage(errMsg).
				WithErrors(perrResp.GetError())
		}
	}
}

func WithErrorPagingInvalid(perrResp ErrorResponse) func(sdkError Error) {
	return func(sdkError Error) {
		if perrResp == nil {
			return
		}

		errMsg := perrResp.GetMessage()
		if strings.Contains(strings.ToLower(strings.TrimSpace(errMsg)), patternPagingInvalid) {
			sdkError.WithErrorCode(EcPagingInvalid).
				WithMessage(errMsg).
				WithErrors(perrResp.GetError())
		}
	}
}

func WithErrorUnexpected(presponse *req.Response) func(Error) {
	statusCode := 0
	url := ""
	err := fmt.Errorf("unexpected error from making request to external service")
	if presponse != nil {
		if presponse.Response != nil && presponse.StatusCode != 0 {
			statusCode = presponse.StatusCode
		}

		if presponse.Request != nil && presponse.Request.URL != nil {
			url = presponse.Request.URL.String()
		}

		if presponse.Err != nil {
			err = presponse.Err
		}
	}

	return func(sdkErr Error) {
		sdkErr.WithErrorCode(EcUnexpectedError).
			WithMessage("Unexpected Error").
			WithErrors(err).
			WithParameters(map[string]interface{}{
				"statusCode": statusCode,
				"url":        url,
			})
	}
}

func WithErrorPaymentMethodNotAllow(perrResp ErrorResponse) func(sdkError Error) {
	return func(sdkError Error) {
		if perrResp == nil {
			return
		}

		errMsg := strings.ToLower(strings.TrimSpace(perrResp.GetMessage()))
		if strings.Contains(errMsg, "ext_pm_payment_method_not_allow") {
			sdkError.WithErrorCode(EcPaymentMethodNotAllow).
				WithMessage(perrResp.GetMessage()).
				WithErrors(perrResp.GetError())
		}
	}
}

func WithErrorCreditNotEnough(perrResp ErrorResponse) func(sdkError Error) {
	return func(sdkError Error) {
		if perrResp == nil {
			return
		}

		errMsg := strings.ToLower(strings.TrimSpace(perrResp.GetMessage()))
		if strings.Contains(errMsg, "ext_pm_credit_not_enough") {
			sdkError.WithErrorCode(EcCreditNotEnough).
				WithMessage(perrResp.GetMessage()).
				WithErrors(perrResp.GetError())
		}
	}
}

func WithErrorProjectConflict(perrResp ErrorResponse) func(sdkError Error) {
	return func(sdkError Error) {
		if perrResp == nil {
			return
		}

		errMsg := strings.ToLower(strings.TrimSpace(perrResp.GetMessage()))
		if regexErrorProjectConflict.FindString(errMsg) != "" {
			sdkError.WithErrorCode(EcProjectConflict).
				WithMessage(errMsg).
				WithErrors(perrResp.GetError())
		}
	}
}
