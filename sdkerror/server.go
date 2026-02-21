package sdkerror

import "regexp"

const (
	patternServerNotFound                      = "cannot get server with id"
	patternServerCreating                      = "cannot delete server with status creating"
	patternServerExceedQuota                   = "exceeded vm quota"
	patternServerDeleting                      = "cannot delete server with status deleting"
	patternServerBilling                       = "cannot delete server with status creating-billing"
	patternBillingPaymentMethodNotAllowed      = "payment method is not allowed for the user"
	patternServerAttachVolumeQuotaExceeded     = "exceeded volume_per_server quota"
	patternServerAttachEncryptedVolume         = "cannot attach encryption volume"
	patternServerExpired                       = "server is expired"
	patternServerFlavorSystemExceedQuota       = "there are no more remaining flavor with id"
	patternServerUpdatingSecgroups             = "cannot change security group of server with status changing-security-group"
	patternServerExceedCpuQuota                = "exceeded vcpu quota. current used"
	patternServerImageNotSupported             = "doesn't support image with id"
	patternImageNotSupport                     = "don't support image"
	patternServerCanNotAttachFloatingIp        = "the server only allows attaching 1 floating ip"
	patternServerFlavorNotSupported            = `flavor [^.]+ don't support image [^.]+`
	patternServerDeleteServerUpdatingSecgroups = "cannot delete server with status changing-security-group"
	patternServerExceedFloatingIpQuota         = "exceeded floating_ip quota"
	patternImageNotFound                       = "cannot get image with id"
	patternServerGroupNotFound                 = "not found server group"
	patternServerGroupInUse                    = "server group is in use"
	patternServerGroupNameMustBeUnique         = "name must be unique"
)

var (
	regexErrorServerFlavorNotSupported = regexp.MustCompile(patternServerFlavorNotSupported)
)

func init() {
	register(EcVServerFlavorNotSupported, &classifier{match: matchRegexps(regexErrorServerFlavorNotSupported)})
	register(EcVServerServerNotFound, &classifier{match: containsAny(patternServerNotFound)})
	register(EcVServerImageNotFound, &classifier{match: containsAny(patternImageNotFound)})
	register(EcVServerServerFlavorSystemExceedQuota, &classifier{
		match: containsAny(patternServerFlavorSystemExceedQuota), category: ErrCatInfra,
	})
	register(EcVServerServerDeleteCreatingServer, &classifier{match: containsAny(patternServerCreating)})
	register(EcVServerServerExpired, &classifier{match: containsAny(patternServerExpired)})
	register(EcVServerServerUpdatingSecgroups, &classifier{
		match: containsAny(patternServerUpdatingSecgroups, patternServerDeleteServerUpdatingSecgroups),
	})
	register(EcVServerServerExceedCpuQuota, &classifier{
		match: containsAny(patternServerExceedCpuQuota), category: ErrCatQuota,
	})
	register(EcVServerServerExceedFloatingIpQuota, &classifier{
		match: containsAny(patternServerExceedFloatingIpQuota), category: ErrCatQuota,
	})
	register(EcVServerServerImageNotSupported, &classifier{
		match: containsAny(patternServerImageNotSupported, patternImageNotSupport),
	})
	register(EcVServerServerExceedQuota, &classifier{
		match: containsAny(patternServerExceedQuota), category: ErrCatQuota,
	})
	register(EcVServerServerDeleteDeletingServer, &classifier{match: containsAny(patternServerDeleting)})
	register(EcVServerServerDeleteBillingServer, &classifier{match: containsAny(patternServerBilling)})
	register(EcVServerCreateBillingPaymentMethodNotAllowed, &classifier{match: containsAny(patternBillingPaymentMethodNotAllowed)})
	register(EcVServerServerVolumeAttachQuotaExceeded, &classifier{
		match: containsAny(patternServerAttachVolumeQuotaExceeded), category: ErrCatQuota,
	})
	register(EcVServerServerAttachEncryptedVolume, &classifier{match: containsAny(patternServerAttachEncryptedVolume)})
	register(EcVServerServerCanNotAttachFloatingIp, &classifier{match: containsAny(patternServerCanNotAttachFloatingIp)})
	register(EcVServerServerGroupNotFound, &classifier{match: containsAny(patternServerGroupNotFound)})
	register(EcVServerServerGroupInUse, &classifier{match: containsAny(patternServerGroupInUse)})
	register(EcVServerServerGroupNameMustBeUnique, &classifier{match: containsAny(patternServerGroupNameMustBeUnique)})
}
