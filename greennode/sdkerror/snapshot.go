package sdkerror

import "strings"

const (
	patternSnapshotNameNotValid = "only letters (a-z, a-z, 0-9, '.', '@', '_', '-', space) are allowed. your input data length must be between 5 and 50" // "Volume name is not valid"
	patternSnapshotNotFound     = "not found snapshot-volume-point"
)

func WithErrorSnapshotNameNotValid(perrResp ErrorResponse) func(sdkError Error) {
	return func(sdkError Error) {
		if perrResp == nil {
			return
		}

		errMsg := perrResp.GetMessage()
		if strings.Contains(strings.ToLower(strings.TrimSpace(errMsg)), patternSnapshotNameNotValid) {
			sdkError.WithErrorCode(EcVServerSnapshotNameNotValid).
				WithMessage(errMsg).
				WithErrors(perrResp.GetError())
		}
	}
}

func WithErrorSnapshotNameNotFound(perrResp ErrorResponse) func(sdkError Error) {
	return func(sdkError Error) {
		if perrResp == nil {
			return
		}

		errMsg := perrResp.GetMessage()
		if strings.Contains(strings.ToLower(strings.TrimSpace(errMsg)), patternSnapshotNotFound) {
			sdkError.WithErrorCode(EcVServerSnapshotNotFound).
				WithMessage(errMsg).
				WithErrors(perrResp.GetError())
		}
	}
}
