package sdkerror

import "regexp"

const (
	patternLoadBalancerNotBelongToProject  = "does not belong to project"
	patternLoadBalancerNotFound            = "cannot get load balancer with id"
	patternLoadBalancerDuplicatePoolName   = "duplicated pool name"
	patternLoadBalancerNotFound2           = "could not find resource"
	patternListenerDuplicateName           = "duplicated listener name"
	patternListenerNotFound                = "cannot get listener with id"
	patternListenerNotBelongToLoadBalancer = `listener id [^.]+ is not belong to load balancer id [^.]+`
	patternListenerDuplicateProtocolOrPort = "duplicated listener protocol port"
	patternPoolNotFound                    = "cannot get pool with id"
	patternPoolInUse                       = "is used in listener"
	patternLoadBalancerNotReady            = `(?:the )?load balancer id [^.]+ is not ready`
	patternListenerNotReady                = `listener id [^.]+ is not ready`
	patternMemberMustIdentical             = "the members provided are identical to the existing members in the pool"
	patternPoolIsUpdating                  = `pool id [^.]+ is updating`
	patternLoadBalancerExceedQuota         = "exceeded load_balancer quota. current used"
	patternLoadBalancerIsDeleting          = `load balancer id [^.]+ is deleting`
	patternLoadBalancerIsCreating          = `load balancer id [^.]+ is creating`
	patternLoadBalancerResizeSamePackage   = "is the same as the current package"
	patternLoadbalancerPackageNotFound     = "invalid package id"
	patternLoadBalancerIsUpdating          = `load balancer id [^.]+ is updating`
	patternGlobalLoadBalancerNotFound      = "global load balancer is not found"
)

var (
	regexErrorLoadBalancerNotReady            = regexp.MustCompile(patternLoadBalancerNotReady)
	regexErrorListenerNotReady                = regexp.MustCompile(patternListenerNotReady)
	regexErrorPoolIsUpdating                  = regexp.MustCompile(patternPoolIsUpdating)
	regexErrorLoadBalancerIsDeleting          = regexp.MustCompile(patternLoadBalancerIsDeleting)
	regexErrorLoadBalancerIsCreating          = regexp.MustCompile(patternLoadBalancerIsCreating)
	regexErrorLoadBalancerIsUpdating          = regexp.MustCompile(patternLoadBalancerIsUpdating)
	regexErrorListenerNotBelongToLoadBalancer = regexp.MustCompile(patternListenerNotBelongToLoadBalancer)
)

func init() {
	register(EcVLBLoadBalancerNotFound, &classifier{
		match: containsAny(patternLoadBalancerNotBelongToProject, patternLoadBalancerNotFound, patternLoadBalancerNotFound2),
	})
	register(EcVLBLoadBalancerExceedQuota, &classifier{
		match: containsAny(patternLoadBalancerExceedQuota), category: ErrCatQuota,
	})
	register(EcVLBLoadBalancerDuplicatePoolName, &classifier{match: containsAny(patternLoadBalancerDuplicatePoolName)})
	register(EcVLBListenerDuplicateProtocolOrPort, &classifier{match: containsAny(patternListenerDuplicateProtocolOrPort)})
	register(EcVLBListenerDuplicateName, &classifier{match: containsAny(patternListenerDuplicateName)})
	register(EcVLBPoolNotFound, &classifier{match: containsAny(patternPoolNotFound)})
	register(EcVLBPoolInUse, &classifier{match: containsAny(patternPoolInUse)})
	register(EcVLBLoadBalancerNotReady, &classifier{
		match: matchRegexps(regexErrorLoadBalancerNotReady, regexErrorListenerNotReady, regexErrorPoolIsUpdating, regexErrorLoadBalancerIsUpdating),
	})
	register(EcVLBLoadBalancerIsDeleting, &classifier{match: matchRegexps(regexErrorLoadBalancerIsDeleting)})
	register(EcVLBLoadBalancerIsCreating, &classifier{match: matchRegexps(regexErrorLoadBalancerIsCreating)})
	register(EcVLBLoadBalancerIsUpdating, &classifier{match: matchRegexps(regexErrorLoadBalancerIsUpdating)})
	register(EcVLBListenerNotFound, &classifier{
		match: matchAnyOf(containsAny(patternListenerNotFound), matchRegexps(regexErrorListenerNotBelongToLoadBalancer)),
	})
	register(EcVLBMemberMustIdentical, &classifier{match: containsAny(patternMemberMustIdentical)})
	register(EcVLBLoadBalancerResizeSamePackage, &classifier{match: containsAny(patternLoadBalancerResizeSamePackage)})
	register(EcVLBLoadBalancerPackageNotFound, &classifier{match: containsAny(patternLoadbalancerPackageNotFound)})
	register(EcGlobalLoadBalancerNotFound, &classifier{match: containsAny(patternGlobalLoadBalancerNotFound)})
}
