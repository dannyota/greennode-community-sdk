package sdkerror

func init() {
	register(EcVNetworkEndpointStatusInvalid, &classifier{match: matchErrCode("ENDPOINT_STATUS_INVALID")})
	register(EcVNetworkEndpointOfVpcExists, &classifier{match: matchErrCode("ENDPOINT_OF_VPC_IS_EXISTS")})
	register(EcVNetworkEndpointPackageNotBelongToEndpointService, &classifier{match: matchErrCode("ENDPOINT_PACKAGE_NOT_BELONG_TO_ENDPOINT_SERVICE")})
	register(EcVNetworkContainInvalidCharacter, &classifier{match: matchErrCode("CONTAIN_INVALID_CHARACTER")})
	register(EcVNetworkLockOnProcess, &classifier{match: matchErrCode("LOCK_ON_PROCESS")})
	register(EcVNetworkEndpointTagNotFound, &classifier{match: matchErrCode("TAG_RESOURCE_WAS_DELETED")})
	register(EcVNetworkEndpointTagExisted, &classifier{match: matchErrCode("TAG_EXISTED")})
}
