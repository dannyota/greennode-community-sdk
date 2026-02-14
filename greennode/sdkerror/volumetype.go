package sdkerror

import "strings"

const (
	patternVolumeTypeNotFound = "cannot get volume type with id"
)

func WithErrorVolumeTypeNotFound(errResp ErrorResponse) func(sdkError Error) {
	return func(sdkError Error) {
		if errResp == nil {
			return
		}

		errMsg := errResp.GetMessage()
		if strings.Contains(strings.ToLower(strings.TrimSpace(errMsg)), patternVolumeTypeNotFound) {
			sdkError.WithErrorCode(EcVServerVolumeTypeNotFound).
				WithMessage(errMsg).
				WithErrors(errResp.Err())
		}
	}
}
