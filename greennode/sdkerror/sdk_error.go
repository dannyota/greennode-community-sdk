package sdkerror

import (
	"errors"
	"fmt"
	"sync"
)

var (
	_ Error = new(SdkError)
)

type (
	SdkError struct {
		error      error
		errorCode  ErrorCode
		message    string
		categories map[ErrorCategory]struct{}
		parameters *sync.Map
	}

	ErrorCode string

	ErrorCategory string
)

func (e *SdkError) IsError(errCode ErrorCode) bool {
	return e.errorCode == errCode
}

func (e *SdkError) IsErrorAny(errCodes ...ErrorCode) bool {
	for _, perrCode := range errCodes {
		if e.errorCode == perrCode {
			return true
		}
	}

	return false
}

func (e *SdkError) IsCategory(category ErrorCategory) bool {
	if e.categories == nil {
		return false
	}

	_, ok := e.categories[category]
	return ok
}

func (e *SdkError) IsCategories(categories ...ErrorCategory) bool {
	if e.categories == nil {
		return false
	}

	for _, c := range categories {
		if _, ok := e.categories[c]; ok {
			return true
		}
	}
	return false
}

func (e *SdkError) WithErrorCode(errCode ErrorCode) Error {
	e.errorCode = errCode
	return e
}

func (e *SdkError) WithMessage(msg string) Error {
	e.message = msg
	return e
}

func (e *SdkError) WithErrors(errs ...error) Error {
	if len(errs) == 0 {
		return e
	}

	if len(errs) == 1 {
		e.error = errs[0]
		return e
	}

	for _, err := range errs {
		e.error = errors.Join(e.error, err)
	}

	return e
}

func (e *SdkError) WithErrorCategories(categories ...ErrorCategory) Error {
	if e.categories == nil {
		e.categories = make(map[ErrorCategory]struct{})
	}
	for _, c := range categories {
		e.categories[c] = struct{}{}
	}

	return e
}

func (e *SdkError) WithParameters(params map[string]interface{}) Error {
	if e.parameters == nil {
		e.parameters = new(sync.Map)
		return e
	}

	for key, val := range params {
		e.parameters.Store(key, val)
	}

	return e
}

func (e *SdkError) WithKVparameters(params ...interface{}) Error {
	if e.parameters == nil {
		e.parameters = new(sync.Map)
	}

	// Always make sure that the length of pparams is even
	if len(params)%2 != 0 {
		params = append(params, nil)
	}

	for i := 0; i < len(params); i += 2 {
		key, ok := params[i].(string)
		if !ok {
			continue
		}

		e.parameters.Store(key, params[i+1])
	}

	return e
}

func (e *SdkError) Err() error {
	return e.error
}

func (e *SdkError) GetMessage() string {
	return e.message
}

func (e *SdkError) ErrorCode() ErrorCode {
	return e.errorCode
}

func (e *SdkError) StringErrorCode() string {
	return string(e.errorCode)
}

func (e *SdkError) Parameters() map[string]interface{} {
	res := make(map[string]interface{})
	if e.parameters != nil {
		e.parameters.Range(func(key, val interface{}) bool {
			res[key.(string)] = val
			return true
		})
	}

	return res
}

func (e *SdkError) ErrorCategories() []ErrorCategory {
	result := make([]ErrorCategory, 0, len(e.categories))
	for c := range e.categories {
		result = append(result, c)
	}
	return result
}

func (e *SdkError) ErrorMessages() string {
	if e.error == nil {
		return e.message
	}

	return fmt.Sprintf("%s: %s", e.message, e.error.Error())
}

func (e *SdkError) ListParameters() []interface{} {
	var result []interface{}
	if e.parameters == nil {
		return result
	}

	e.parameters.Range(func(key, val interface{}) bool {
		result = append(result, key, val)
		return true
	})

	return result
}

func (e *SdkError) RemoveCategories(categories ...ErrorCategory) Error {
	if e.categories == nil {
		return e
	}

	for _, c := range categories {
		delete(e.categories, c)
	}
	return e
}

func (e *SdkError) AppendCategories(categories ...ErrorCategory) Error {
	if e.categories == nil {
		e.categories = make(map[ErrorCategory]struct{})
	}

	for _, c := range categories {
		e.categories[c] = struct{}{}
	}
	return e
}
