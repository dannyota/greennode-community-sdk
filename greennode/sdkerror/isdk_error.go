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
	WithParameters(params map[string]any) Error
	WithKVparameters(params ...any) Error

	Err() error
	GetMessage() string
	ErrorCode() ErrorCode
	StringErrorCode() string
	Parameters() map[string]any
	ErrorCategories() []ErrorCategory
	ErrorMessages() string
	ListParameters() []any

	RemoveCategories(categories ...ErrorCategory) Error

	AppendCategories(categories ...ErrorCategory) Error
}

type ErrorResponse interface {
	GetMessage() string
	Err() error
}
