package sdkerror

import (
	"fmt"
)

const (
	loginFailedPrefixMsg = "There are some problems with your service account key pair, please re-generate a new one. Error message: %s"
)

func WithErrorAuthenticationFailed(errResp ErrorResponse) func(Error) {
	return func(sdkErr Error) {
		if errResp == nil {
			return
		}

		if errResp.Err() == nil {
			return
		}

		if errResp.Err().Error() == "AUTHENTICATION_FAILED" {
			sdkErr.WithErrorCode(EcAuthenticationFailed).
				WithErrors(errResp.Err()).
				WithMessage(fmt.Sprintf(loginFailedPrefixMsg, errResp.GetMessage())).
				WithErrorCategories(ErrCatIAM)
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

func WithErrorTooManyFailedLogin(errResp ErrorResponse) func(Error) {
	return func(sdkErr Error) {
		if errResp == nil {
			return
		}

		if errResp.Err() == nil {
			return
		}

		if errResp.Err().Error() == "TOO_MANY_FAILED_LOGINS" {
			sdkErr.WithErrorCode(EcTooManyFailedLogins).
				WithErrors(errResp.Err()).
				WithMessage(fmt.Sprintf(loginFailedPrefixMsg, errResp.GetMessage())).
				WithErrorCategories(ErrCatIAM)
		}
	}
}

func WithErrorUnknownAuthFailure(errResp ErrorResponse) func(Error) {
	return func(sdkErr Error) {
		if errResp == nil {
			return
		}

		if errResp.Err() == nil {
			return
		}

		if sdkErr.ErrorCode() == EcUnknownError {
			sdkErr.WithErrorCode(EcUnknownAuthFailure).
				WithMessage(errResp.GetMessage()).
				WithErrors(errResp.Err())
		}
	}
}
