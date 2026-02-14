package sdkerror

type Error interface {
	IsError(perrCode ErrorCode) bool
	IsErrorAny(perrCodes ...ErrorCode) bool
	IsCategory(pcategory ErrorCategory) bool
	IsCategories(pcategories ...ErrorCategory) bool

	WithErrorCode(perrCode ErrorCode) Error
	WithMessage(pmsg string) Error
	WithErrors(perrs ...error) Error
	WithErrorCategories(pcategories ...ErrorCategory) Error
	WithParameters(pparams map[string]interface{}) Error
	WithKVparameters(pparams ...interface{}) Error

	GetError() error
	GetMessage() string
	GetErrorCode() ErrorCode
	GetStringErrorCode() string
	GetParameters() map[string]interface{}
	GetErrorCategories() []ErrorCategory
	GetErrorMessages() string
	GetListParameters() []interface{}

	RemoveCategories(pcategories ...ErrorCategory) Error

	AppendCategories(pcategories ...ErrorCategory) Error
}

type ErrorResponse interface {
	GetMessage() string
	GetError() error
}
