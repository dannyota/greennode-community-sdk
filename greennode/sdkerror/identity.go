package sdkerror

const (
	loginFailedPrefixMsg = "There are some problems with your service account key pair, please re-generate a new one. Error message: %s"
)

func init() {
	register(EcAuthenticationFailed, &classifier{
		match:  matchErrCode("AUTHENTICATION_FAILED"),
		msgFmt: loginFailedPrefixMsg, category: ErrCatIAM,
	})
	register(EcTooManyFailedLogins, &classifier{
		match:  matchErrCode("TOO_MANY_FAILED_LOGINS"),
		msgFmt: loginFailedPrefixMsg, category: ErrCatIAM,
	})
	register(EcUnknownAuthFailure, &classifier{catchAll: true})
}
