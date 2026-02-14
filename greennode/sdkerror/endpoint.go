package sdkerror

import "strings"

func WithErrorEndpointStatusInvalid(errResp ErrorResponse) func(sdkError Error) {
	return func(sdkError Error) {
		if errResp == nil {
			return
		}

		if strings.ToUpper(strings.TrimSpace(errResp.GetError().Error())) == "ENDPOINT_STATUS_INVALID" {
			sdkError.WithErrorCode(EcVNetworkEndpointStatusInvalid).
				WithMessage(errResp.GetMessage()).
				WithErrors(errResp.GetError())
		}
	}
}

func WithErrorEndpointOfVpcExists(errResp ErrorResponse) func(sdkError Error) {
	return func(sdkError Error) {
		if errResp == nil {
			return
		}

		if strings.ToUpper(strings.TrimSpace(errResp.GetError().Error())) == "ENDPOINT_OF_VPC_IS_EXISTS" {
			sdkError.WithErrorCode(EcVNetworkEndpointOfVpcExists).
				WithMessage(errResp.GetMessage()).
				WithErrors(errResp.GetError())
		}
	}
}

func WithErrorEndpointPackageNotBelongToEndpointService(errResp ErrorResponse) func(sdkError Error) {
	return func(sdkError Error) {
		if errResp == nil {
			return
		}

		if strings.ToUpper(strings.TrimSpace(errResp.GetError().Error())) == "ENDPOINT_PACKAGE_NOT_BELONG_TO_ENDPOINT_SERVICE" {
			sdkError.WithErrorCode(EcVNetworkEndpointPackageNotBelongToEndpointService).
				WithMessage(errResp.GetMessage()).
				WithErrors(errResp.GetError())
		}
	}
}

func WithErrorContainInvalidCharacter(errResp ErrorResponse) func(sdkError Error) {
	return func(sdkError Error) {
		if errResp == nil {
			return
		}

		if strings.ToUpper(strings.TrimSpace(errResp.GetError().Error())) == "CONTAIN_INVALID_CHARACTER" {
			sdkError.WithErrorCode(EcVNetworkContainInvalidCharacter).
				WithMessage(errResp.GetMessage()).
				WithErrors(errResp.GetError())
		}
	}
}

func WithErrorLockOnProcess(errResp ErrorResponse) func(sdkError Error) {
	return func(sdkError Error) {
		if errResp == nil {
			return
		}

		if strings.ToUpper(strings.TrimSpace(errResp.GetError().Error())) == "LOCK_ON_PROCESS" {
			sdkError.WithErrorCode(EcVNetworkLockOnProcess).
				WithMessage(errResp.GetMessage()).
				WithErrors(errResp.GetError())
		}
	}
}

func WithErrorEndpointTagNotFound(errResp ErrorResponse) func(sdkError Error) {
	return func(sdkError Error) {
		if errResp == nil {
			return
		}

		if strings.ToUpper(strings.TrimSpace(errResp.GetError().Error())) == "TAG_RESOURCE_WAS_DELETED" {
			sdkError.WithErrorCode(EcVNetworkEndpointTagNotFound).
				WithMessage(errResp.GetMessage()).
				WithErrors(errResp.GetError())
		}
	}
}

func WithErrorEndpointTagExisted(errResp ErrorResponse) func(sdkError Error) {
	return func(sdkError Error) {
		if errResp == nil {
			return
		}

		if strings.ToUpper(strings.TrimSpace(errResp.GetError().Error())) == "TAG_EXISTED" {
			sdkError.WithErrorCode(EcVNetworkEndpointTagExisted).
				WithMessage(errResp.GetMessage()).
				WithErrors(errResp.GetError())
		}
	}
}
