package sdkerror

const (
	patternVolumeTypeNotFound = "cannot get volume type with id"
)

func init() {
	register(EcVServerVolumeTypeNotFound, &classifier{match: containsAny(patternVolumeTypeNotFound)})
}
