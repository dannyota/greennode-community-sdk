package sdkerror

import "strings"

func WithErrorEndpointStatusInvalid(perrResp ErrorResponse) func(sdkError Error) {
	return func(sdkError Error) {
		if perrResp == nil {
			return
		}

		if strings.ToUpper(strings.TrimSpace(perrResp.GetError().Error())) == "ENDPOINT_STATUS_INVALID" {
			sdkError.WithErrorCode(EcVNetworkEndpointStatusInvalid).
				WithMessage(perrResp.GetMessage()).
				WithErrors(perrResp.GetError())
		}
	}
}

func WithErrorEndpointOfVpcExists(perrResp ErrorResponse) func(sdkError Error) {
	return func(sdkError Error) {
		if perrResp == nil {
			return
		}

		if strings.ToUpper(strings.TrimSpace(perrResp.GetError().Error())) == "ENDPOINT_OF_VPC_IS_EXISTS" {
			sdkError.WithErrorCode(EcVNetworkEndpointOfVpcExists).
				WithMessage(perrResp.GetMessage()).
				WithErrors(perrResp.GetError())
		}
	}
}

func WithErrorEndpointPackageNotBelongToEndpointService(perrResp ErrorResponse) func(sdkError Error) {
	return func(sdkError Error) {
		if perrResp == nil {
			return
		}

		if strings.ToUpper(strings.TrimSpace(perrResp.GetError().Error())) == "ENDPOINT_PACKAGE_NOT_BELONG_TO_ENDPOINT_SERVICE" {
			sdkError.WithErrorCode(EcVNetworkEndpointPackageNotBelongToEndpointService).
				WithMessage(perrResp.GetMessage()).
				WithErrors(perrResp.GetError())
		}
	}
}

func WithErrorContainInvalidCharacter(perrResp ErrorResponse) func(sdkError Error) {
	return func(sdkError Error) {
		if perrResp == nil {
			return
		}

		if strings.ToUpper(strings.TrimSpace(perrResp.GetError().Error())) == "CONTAIN_INVALID_CHARACTER" {
			sdkError.WithErrorCode(EcVNetworkContainInvalidCharacter).
				WithMessage(perrResp.GetMessage()).
				WithErrors(perrResp.GetError())
		}
	}
}

func WithErrorLockOnProcess(perrResp ErrorResponse) func(sdkError Error) {
	return func(sdkError Error) {
		if perrResp == nil {
			return
		}

		if strings.ToUpper(strings.TrimSpace(perrResp.GetError().Error())) == "LOCK_ON_PROCESS" {
			sdkError.WithErrorCode(EcVNetworkLockOnProcess).
				WithMessage(perrResp.GetMessage()).
				WithErrors(perrResp.GetError())
		}
	}
}

func WithErrorEndpointTagNotFound(perrResp ErrorResponse) func(sdkError Error) {
	return func(sdkError Error) {
		if perrResp == nil {
			return
		}

		if strings.ToUpper(strings.TrimSpace(perrResp.GetError().Error())) == "TAG_RESOURCE_WAS_DELETED" {
			sdkError.WithErrorCode(EcVNetworkEndpointTagNotFound).
				WithMessage(perrResp.GetMessage()).
				WithErrors(perrResp.GetError())
		}
	}
}

func WithErrorEndpointTagExisted(perrResp ErrorResponse) func(sdkError Error) {
	return func(sdkError Error) {
		if perrResp == nil {
			return
		}

		if strings.ToUpper(strings.TrimSpace(perrResp.GetError().Error())) == "TAG_EXISTED" {
			sdkError.WithErrorCode(EcVNetworkEndpointTagExisted).
				WithMessage(perrResp.GetMessage()).
				WithErrors(perrResp.GetError())
		}
	}
}
