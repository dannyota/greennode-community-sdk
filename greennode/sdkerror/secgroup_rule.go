package sdkerror

const (
	patternSecgroupRuleNotFound    = "cannot get security group rule with id"
	patternSecgroupRuleExists      = "securitygroupruleexists"
	patternSecgroupRuleExceedQuota = "exceeded secgroup_rule quota"
)

func init() {
	register(EcVServerSecgroupRuleNotFound, &classifier{match: containsAny(patternSecgroupRuleNotFound)})
	register(EcVServerSecgroupRuleAlreadyExists, &classifier{match: containsAny(patternSecgroupRuleExists)})
	register(EcVServerSecgroupRuleExceedQuota, &classifier{
		match: containsAny(patternSecgroupRuleExceedQuota), category: ErrCatQuota,
	})
}
