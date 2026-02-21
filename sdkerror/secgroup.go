package sdkerror

const (
	patternSecgroupNotFound          = "cannot get security group with id"
	patternSecgroupNameAlreadyExists = "name of security group already exist"
	patternSecgroupExceedQuota       = "exceeded secgroup quota"
	patternSecgroupInUse             = "securitygroupinuse"
)

func init() {
	register(EcVServerSecgroupNotFound, &classifier{match: containsAny(patternSecgroupNotFound)})
	register(EcVServerSecgroupNameAlreadyExists, &classifier{match: containsAny(patternSecgroupNameAlreadyExists)})
	register(EcVServerSecgroupExceedQuota, &classifier{
		match: containsAny(patternSecgroupExceedQuota), category: ErrCatQuota,
	})
	register(EcVServerSecgroupInUse, &classifier{match: containsAny(patternSecgroupInUse)})
}
