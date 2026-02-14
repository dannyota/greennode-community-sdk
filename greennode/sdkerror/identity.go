package sdkerror

import (
	"fmt"
)

const (
	loginFailedPrefixMsg = "There are some problems with your service account key pair, please re-generate a new one. Error message: %s"
)

func WithErrorAuthenticationFailed(perrResp ErrorResponse) func(Error) {
	return func(sdkErr Error) {
		if perrResp == nil {
			return
		}

		if perrResp.GetError() == nil {
			return
		}

		if perrResp.GetError().Error() == "AUTHENTICATION_FAILED" {
			sdkErr.WithErrorCode(EcAuthenticationFailed).
				WithErrors(perrResp.GetError()).
				WithMessage(fmt.Sprintf(loginFailedPrefixMsg, perrResp.GetMessage())).
				WithErrorCategories(ErrCatIam)
		}
	}
}

func WithErrorReauthFuncNotSet() func(Error) {
	return func(sdkErr Error) {
		sdkErr.WithErrorCode(EcReauthFuncNotSet).
			WithMessage("Reauthentication function is not configured").
			WithErrors(fmt.Errorf("reauthentication function is not configured"))
	}
}

func WithErrorTooManyFailedLogin(perrResp ErrorResponse) func(Error) {
	return func(sdkErr Error) {
		if perrResp == nil {
			return
		}

		if perrResp.GetError() == nil {
			return
		}

		if perrResp.GetError().Error() == "TOO_MANY_FAILED_LOGINS" {
			sdkErr.WithErrorCode(EcTooManyFailedLogins).
				WithErrors(perrResp.GetError()).
				WithMessage(fmt.Sprintf(loginFailedPrefixMsg, perrResp.GetMessage())).
				WithErrorCategories(ErrCatIam)
		}
	}
}

func WithErrorUnknownAuthFailure(perrResp ErrorResponse) func(Error) {
	return func(sdkErr Error) {
		if perrResp == nil {
			return
		}

		if perrResp.GetError() == nil {
			return
		}

		if sdkErr.GetErrorCode() == EcUnknownError {
			sdkErr.WithErrorCode(EcUnknownAuthFailure).
				WithMessage(perrResp.GetMessage()).
				WithErrors(perrResp.GetError())
		}
	}
}
