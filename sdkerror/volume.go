package sdkerror

import "regexp"

const (
	patternVolumeNameNotValid              = "only letters (a-z, a-z, 0-9, '.', '@', '_', '-', space) are allowed. your input data length must be between 5 and 50"
	patternVolumeSizeOutOfRange            = "field volume_size must from"
	patternVolumeNewSizeOutOfRange         = "field new_volume_size must from"
	patternVolumeNotFound                  = `volume with id [^.]+ is not found`
	patternVolumeNotFound2                 = "cannot get volume with id"
	patternVolumeAvailable                 = "this volume is available"
	patternVolumeAlreadyAttached           = "already attached to instance"
	patternVolumeAlreadyAttachedThisServer = "this volume has been attached"
	patternVolumeInProcess                 = "is in-process"
	patternVolumeUnchanged                  = "volume size or volume type must be changed"
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

func init() {
	register(EcVServerVolumeNameNotValid, &classifier{match: containsAny(patternVolumeNameNotValid)})
	register(EcVServerVolumeSizeOutOfRange, &classifier{
		match: containsAny(patternVolumeSizeOutOfRange, patternVolumeNewSizeOutOfRange),
	})
	register(EcVServerVolumeSizeExceedGlobalQuota, &classifier{
		match: containsAny(patternVolumeSizeExceedGlobalQuota), category: ErrCatQuota,
	})
	register(EcVServerVolumeExceedQuota, &classifier{
		match: containsAny(patternVolumeExceedQuota), category: ErrCatQuota,
	})
	register(EcVServerVolumeNotFound, &classifier{
		match: matchAnyOf(matchRegexps(regexErrorVolumeNotFound), containsAny(patternVolumeNotFound2)),
	})
	register(EcVServerVolumeAvailable, &classifier{match: containsAny(patternVolumeAvailable)})
	register(EcVServerVolumeAlreadyAttached, &classifier{match: containsAny(patternVolumeAlreadyAttached)})
	register(EcVServerVolumeAlreadyAttachedThisServer, &classifier{match: containsAny(patternVolumeAlreadyAttachedThisServer)})
	register(EcVServerVolumeInProcess, &classifier{match: containsAny(patternVolumeInProcess)})
	register(EcVServerVolumeUnchanged, &classifier{match: containsAny(patternVolumeUnchanged)})
	register(EcVServerVolumeMustSameZone, &classifier{match: containsAny(patternVolumeMustSameZone)})
	register(EcVServerVolumeMigrateMissingInit, &classifier{match: containsAny(patternVolumeMigrateMissingInit)})
	register(EcVServerVolumeMigrateNeedProcess, &classifier{match: containsAny(patternVolumeMigrateNeedProcess)})
	register(EcVServerVolumeMigrateNeedConfirm, &classifier{match: containsAny(patternVolumeMigrateNeedConfirm)})
	register(EcVServerVolumeMigrateBeingProcess, &classifier{match: containsAny(patternVolumeMigrateBeingProcess)})
	register(EcVServerVolumeMigrateBeingFinish, &classifier{match: containsAny(patternVolumeMigrateBeingFinish)})
	register(EcVServerVolumeMigrateProcessingConfirm, &classifier{match: containsAny(patternVolumeMigrateProcessingConfirm)})
	register(EcVServerVolumeMigrateBeingMigrating, &classifier{match: containsAny(patternVolumeMigrateBeingMigrating)})
	register(EcVServerVolumeMigrateInSameZone, &classifier{match: containsAny(patternVolumeMigrateInSameZone)})
	register(EcVServerVolumeIsMigrating, &classifier{match: containsAny(patternVolumeIsMigrating)})
}
