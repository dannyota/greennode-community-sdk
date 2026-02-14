package sdkerror

import "fmt"

func WithErrorQuotaNotFound(_ ErrorResponse) func(sdkError Error) {
	return func(sdkError Error) {
		sdkError.WithErrorCode(EcVServerQuotaNotFound).
			WithMessage("Quota not found").
			WithErrors(fmt.Errorf("quota not found"))
	}
}
