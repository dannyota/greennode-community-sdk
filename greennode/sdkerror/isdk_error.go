package sdkerror

type Error interface {
	IsError(errCode ErrorCode) bool
	IsErrorAny(errCodes ...ErrorCode) bool
	IsCategory(category ErrorCategory) bool
	IsCategories(categories ...ErrorCategory) bool

	WithErrorCode(errCode ErrorCode) Error
	WithMessage(msg string) Error
	WithErrors(errs ...error) Error
	WithErrorCategories(categories ...ErrorCategory) Error
	WithParameters(params map[string]interface{}) Error
	WithKVparameters(params ...interface{}) Error

	GetError() error
	GetMessage() string
	GetErrorCode() ErrorCode
	GetStringErrorCode() string
	GetParameters() map[string]interface{}
	GetErrorCategories() []ErrorCategory
	GetErrorMessages() string
	GetListParameters() []interface{}

	RemoveCategories(categories ...ErrorCategory) Error

	AppendCategories(categories ...ErrorCategory) Error
}

type ErrorResponse interface {
	GetMessage() string
	GetError() error
}
