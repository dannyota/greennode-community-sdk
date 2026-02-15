package sdkerror

const (
	patternSnapshotNameNotValid = "only letters (a-z, a-z, 0-9, '.', '@', '_', '-', space) are allowed. your input data length must be between 5 and 50"
	patternSnapshotNotFound     = "not found snapshot-volume-point"
)

func init() {
	register(EcVServerSnapshotNameNotValid, &classifier{match: containsAny(patternSnapshotNameNotValid)})
	register(EcVServerSnapshotNotFound, &classifier{match: containsAny(patternSnapshotNotFound)})
}
