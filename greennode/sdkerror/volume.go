package sdkerror

import (
	"regexp"
	"strings"
)

const ( // "Cannot get volume type with id vtype-6790f903-38d2-454d-919e-5b49184b5927"
	patternVolumeNameNotValid              = "only letters (a-z, a-z, 0-9, '.', '@', '_', '-', space) are allowed. your input data length must be between 5 and 50" // "Volume name is not valid"
	patternVolumeSizeOutOfRange            = "field volume_size must from"
	patternVolumeNewSizeOutOfRange         = "field new_volume_size must from"
	patternVolumeNotFound                  = `volume with id [^.]+ is not found`
	patternVolumeNotFound2                 = "cannot get volume with id"
	patternVolumeAvailable                 = "this volume is available"
	patternVolumeAlreadyAttached           = "already attached to instance"
	patternVolumeAlreadyAttachedThisServer = "this volume has been attached"
	patternVolumeInProcess                 = "is in-process"
	patternVolumeUnchaged                  = "volume size or volume type must be changed"
	patternVolumeMustSameZone              = "new volume type must be same zone"
	patternVolumeMigrateMissingInit        = "the action must be init-migrate or migrate or confirm-migrate"
	patternVolumeMigrateNeedProcess        = "this volume cannot initialize migration because state is ready to migrate difference"
	patternVolumeMigrateNeedConfirm        = "this volume cannot initialize migration because state is confirm final migration"
	patternVolumeMigrateBeingProcess       = "this volume cannot initialize migration because state is migrating difference"
	patternVolumeMigrateBeingMigrating     = "this volume cannot initialize migration because state is migrating"
	patternVolumeMigrateBeingFinish        = "this volume cannot migrate difference data because state is confirm final migration"
	patternVolumeMigrateProcessingConfirm  = "this volume cannot initialize migration because state is processing to confirm"
	patternVolumeMigrateInSameZone         = "new volume type must be different zone"
	patternVolumeIsMigrating               = "is migrating"
	patternVolumeSizeExceedGlobalQuota     = "exceeded volume_size quota"
	patternVolumeExceedQuota               = "exceeded volume quota. current used"
)

var (
	regexErrorVolumeNotFound = regexp.MustCompile(patternVolumeNotFound)
)

func WithErrorVolumeNameNotValid(errResp ErrorResponse) func(sdkError Error) {
	return func(sdkError Error) {
		if errResp == nil {
			return
		}

		errMsg := errResp.GetMessage()
		if strings.Contains(strings.ToLower(strings.TrimSpace(errMsg)), patternVolumeNameNotValid) {
			sdkError.WithErrorCode(EcVServerVolumeNameNotValid).
				WithMessage(errMsg).
				WithErrors(errResp.GetError())
		}
	}
}

func WithErrorVolumeSizeOutOfRange(errResp ErrorResponse) func(sdkError Error) {
	return func(sdkError Error) {
		if errResp == nil {
			return
		}

		errMsg := errResp.GetMessage()
		if strings.Contains(strings.ToLower(strings.TrimSpace(errMsg)), patternVolumeSizeOutOfRange) ||
			strings.Contains(strings.ToLower(strings.TrimSpace(errMsg)), patternVolumeNewSizeOutOfRange) {
			sdkError.WithErrorCode(EcVServerVolumeSizeOutOfRange).
				WithMessage(errMsg).
				WithErrors(errResp.GetError())
		}
	}
}

func WithErrorVolumeSizeExceedGlobalQuota(errResp ErrorResponse) func(sdkError Error) {
	return func(sdkError Error) {
		if errResp == nil {
			return
		}

		errMsg := errResp.GetMessage()
		if strings.Contains(strings.ToLower(strings.TrimSpace(errMsg)), patternVolumeSizeExceedGlobalQuota) {
			sdkError.WithErrorCode(EcVServerVolumeSizeExceedGlobalQuota).
				WithMessage(errMsg).
				WithErrors(errResp.GetError()).
				WithErrorCategories(ErrCatQuota)
		}
	}
}

func WithErrorVolumeExceedQuota(errResp ErrorResponse) func(sdkError Error) {
	return func(sdkError Error) {
		if errResp == nil {
			return
		}

		errMsg := errResp.GetMessage()
		if strings.Contains(strings.ToLower(strings.TrimSpace(errMsg)), patternVolumeExceedQuota) {
			sdkError.WithErrorCode(EcVServerVolumeExceedQuota).
				WithMessage(errMsg).
				WithErrors(errResp.GetError()).
				WithErrorCategories(ErrCatQuota)
		}
	}
}

func WithErrorVolumeNotFound(errResp ErrorResponse) func(sdkError Error) {
	return func(sdkError Error) {
		if errResp == nil {
			return
		}

		errMsg := strings.ToLower(strings.TrimSpace(errResp.GetMessage()))
		if regexErrorVolumeNotFound.FindString(errMsg) != "" ||
			strings.Contains(strings.ToLower(strings.TrimSpace(errMsg)), patternVolumeNotFound2) {
			sdkError.WithErrorCode(EcVServerVolumeNotFound).
				WithMessage(errMsg).
				WithErrors(errResp.GetError())
		}
	}
}

// WithErrorVolumeAvailable indicates that the volume is AVAILABLE state but try to make detach this volume out of server
func WithErrorVolumeAvailable(errResp ErrorResponse) func(sdkError Error) {
	return func(sdkError Error) {
		if errResp == nil {
			return
		}

		errMsg := errResp.GetMessage()
		if strings.Contains(strings.ToLower(strings.TrimSpace(errMsg)), patternVolumeAvailable) {
			sdkError.WithErrorCode(EcVServerVolumeAvailable).
				WithMessage(errMsg).
				WithErrors(errResp.GetError())
		}
	}
}

func WithErrorVolumeAlreadyAttached(errResp ErrorResponse) func(sdkError Error) {
	return func(sdkError Error) {
		if errResp == nil {
			return
		}

		errMsg := errResp.GetMessage()
		if strings.Contains(strings.ToLower(strings.TrimSpace(errMsg)), patternVolumeAlreadyAttached) {
			sdkError.WithErrorCode(EcVServerVolumeAlreadyAttached).
				WithMessage(errMsg).
				WithErrors(errResp.GetError())
		}
	}
}

func WithErrorVolumeAlreadyAttachedThisServer(errResp ErrorResponse) func(sdkError Error) {
	return func(sdkError Error) {
		if errResp == nil {
			return
		}

		errMsg := errResp.GetMessage()
		if strings.Contains(strings.ToLower(strings.TrimSpace(errMsg)), patternVolumeAlreadyAttachedThisServer) {
			sdkError.WithErrorCode(EcVServerVolumeAlreadyAttachedThisServer).
				WithMessage(errMsg).
				WithErrors(errResp.GetError())
		}
	}
}

func WithErrorVolumeInProcess(errResp ErrorResponse) func(sdkError Error) {
	return func(sdkError Error) {
		if errResp == nil {
			return
		}

		errMsg := errResp.GetMessage()
		if strings.Contains(strings.ToLower(strings.TrimSpace(errMsg)), patternVolumeInProcess) {
			sdkError.WithErrorCode(EcVServerVolumeInProcess).
				WithMessage(errMsg).
				WithErrors(errResp.GetError())
		}
	}
}

func WithErrorVolumeUnchanged(errResp ErrorResponse) func(sdkError Error) {
	return func(sdkError Error) {
		if errResp == nil {
			return
		}

		errMsg := errResp.GetMessage()
		if strings.Contains(strings.ToLower(strings.TrimSpace(errMsg)), patternVolumeUnchaged) {
			sdkError.WithErrorCode(EcVServerVolumeUnchanged).
				WithMessage(errMsg).
				WithErrors(errResp.GetError())
		}
	}
}

func WithErrorVolumeMustSameZone(errResp ErrorResponse) func(sdkError Error) {
	return func(sdkError Error) {
		if errResp == nil {
			return
		}

		errMsg := errResp.GetMessage()
		if strings.Contains(strings.ToLower(strings.TrimSpace(errMsg)), patternVolumeMustSameZone) {
			sdkError.WithErrorCode(EcVServerVolumeMustSameZone).
				WithMessage(errMsg).
				WithErrors(errResp.GetError())
		}
	}
}

func WithErrorVolumeMigrateMissingInit(errResp ErrorResponse) func(sdkError Error) {
	return func(sdkError Error) {
		if errResp == nil {
			return
		}

		errMsg := errResp.GetMessage()
		if strings.Contains(strings.ToLower(strings.TrimSpace(errMsg)), patternVolumeMigrateMissingInit) {
			sdkError.WithErrorCode(EcVServerVolumeMigrateMissingInit).
				WithMessage(errMsg).
				WithErrors(errResp.GetError())
		}
	}
}

func WithErrorVolumeMigrateNeedProcess(errResp ErrorResponse) func(sdkError Error) {
	return func(sdkError Error) {
		if errResp == nil {
			return
		}

		errMsg := errResp.GetMessage()
		if strings.Contains(strings.ToLower(strings.TrimSpace(errMsg)), patternVolumeMigrateNeedProcess) {
			sdkError.WithErrorCode(EcVServerVolumeMigrateNeedProcess).
				WithMessage(errMsg).
				WithErrors(errResp.GetError())
		}
	}
}

func WithErrorVolumeMigrateNeedConfirm(errResp ErrorResponse) func(sdkError Error) {
	return func(sdkError Error) {
		if errResp == nil {
			return
		}

		errMsg := errResp.GetMessage()
		if strings.Contains(strings.ToLower(strings.TrimSpace(errMsg)), patternVolumeMigrateNeedConfirm) {
			sdkError.WithErrorCode(EcVServerVolumeMigrateNeedConfirm).
				WithMessage(errMsg).
				WithErrors(errResp.GetError())
		}
	}
}

func WithErrorVolumeMigrateBeingProcess(errResp ErrorResponse) func(sdkError Error) {
	return func(sdkError Error) {
		if errResp == nil {
			return
		}

		errMsg := errResp.GetMessage()
		if strings.Contains(strings.ToLower(strings.TrimSpace(errMsg)), patternVolumeMigrateBeingProcess) {
			sdkError.WithErrorCode(EcVServerVolumeMigrateBeingProcess).
				WithMessage(errMsg).
				WithErrors(errResp.GetError())
		}
	}
}

func WithErrorVolumeMigrateBeingFinish(errResp ErrorResponse) func(sdkError Error) {
	return func(sdkError Error) {
		if errResp == nil {
			return
		}

		errMsg := errResp.GetMessage()
		if strings.Contains(strings.ToLower(strings.TrimSpace(errMsg)), patternVolumeMigrateBeingFinish) {
			sdkError.WithErrorCode(EcVServerVolumeMigrateBeingFinish).
				WithMessage(errMsg).
				WithErrors(errResp.GetError())
		}
	}
}

func WithErrorVolumeMigrateProcessingConfirm(errResp ErrorResponse) func(sdkError Error) {
	return func(sdkError Error) {
		if errResp == nil {
			return
		}

		errMsg := errResp.GetMessage()
		if strings.Contains(strings.ToLower(strings.TrimSpace(errMsg)), patternVolumeMigrateProcessingConfirm) {
			sdkError.WithErrorCode(EcVServerVolumeMigrateProcessingConfirm).
				WithMessage(errMsg).
				WithErrors(errResp.GetError())
		}
	}
}

//

func WithErrorVolumeMigrateBeingMigrating(errResp ErrorResponse) func(sdkError Error) {
	return func(sdkError Error) {
		if errResp == nil {
			return
		}

		errMsg := errResp.GetMessage()
		if strings.Contains(strings.ToLower(strings.TrimSpace(errMsg)), patternVolumeMigrateBeingMigrating) {
			sdkError.WithErrorCode(EcVServerVolumeMigrateBeingMigrating).
				WithMessage(errMsg).
				WithErrors(errResp.GetError())
		}
	}
}

func WithErrorVolumeMigrateInSameZone(errResp ErrorResponse) func(sdkError Error) {
	return func(sdkError Error) {
		if errResp == nil {
			return
		}

		errMsg := errResp.GetMessage()
		if strings.Contains(strings.ToLower(strings.TrimSpace(errMsg)), patternVolumeMigrateInSameZone) {
			sdkError.WithErrorCode(EcVServerVolumeMigrateInSameZone).
				WithMessage(errMsg).
				WithErrors(errResp.GetError())
		}
	}
}

func WithErrorVolumeIsMigrating(errResp ErrorResponse) func(sdkError Error) {
	return func(sdkError Error) {
		if errResp == nil {
			return
		}

		errMsg := errResp.GetMessage()
		if strings.Contains(strings.ToLower(strings.TrimSpace(errMsg)), patternVolumeIsMigrating) {
			sdkError.WithErrorCode(EcVServerVolumeIsMigrating).
				WithMessage(errMsg).
				WithErrors(errResp.GetError())
		}
	}
}

// patternVolumeIsMigrating
