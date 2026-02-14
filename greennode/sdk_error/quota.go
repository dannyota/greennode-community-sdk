package sdk_error

import "fmt"

func WithErrorQuotaNotFound(_ IErrorRespone) func(sdkError IError) {
	return func(sdkError IError) {
		sdkError.WithErrorCode(EcVServerQuotaNotFound).
			WithMessage("Quota not found").
			WithErrors(fmt.Errorf("quota not found"))
	}
}
