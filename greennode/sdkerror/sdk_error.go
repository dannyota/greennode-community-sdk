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

func (s *SdkError) IsError(errCode ErrorCode) bool {
	return s.errorCode == errCode
}

func (s *SdkError) IsErrorAny(errCodes ...ErrorCode) bool {
	for _, perrCode := range errCodes {
		if s.errorCode == perrCode {
			return true
		}
	}

	return false
}

func (s *SdkError) IsCategory(category ErrorCategory) bool {
	if s.categories == nil {
		return false
	}

	_, ok := s.categories[category]
	return ok
}

func (s *SdkError) IsCategories(categories ...ErrorCategory) bool {
	if s.categories == nil {
		return false
	}

	for _, c := range categories {
		if _, ok := s.categories[c]; ok {
			return true
		}
	}
	return false
}

func (s *SdkError) WithErrorCode(errCode ErrorCode) Error {
	s.errorCode = errCode
	return s
}

func (s *SdkError) WithMessage(msg string) Error {
	s.message = msg
	return s
}

func (s *SdkError) WithErrors(errs ...error) Error {
	if len(errs) == 0 {
		return s
	}

	if len(errs) == 1 {
		s.error = errs[0]
		return s
	}

	for _, err := range errs {
		s.error = errors.Join(s.error, err)
	}

	return s
}

func (s *SdkError) WithErrorCategories(categories ...ErrorCategory) Error {
	if s.categories == nil {
		s.categories = make(map[ErrorCategory]struct{})
	}
	for _, c := range categories {
		s.categories[c] = struct{}{}
	}

	return s
}

func (s *SdkError) WithParameters(params map[string]interface{}) Error {
	if s.parameters == nil {
		s.parameters = new(sync.Map)
		return s
	}

	for key, val := range params {
		s.parameters.Store(key, val)
	}

	return s
}

func (s *SdkError) WithKVparameters(params ...interface{}) Error {
	if s.parameters == nil {
		s.parameters = new(sync.Map)
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

		s.parameters.Store(key, params[i+1])
	}

	return s
}

func (s *SdkError) GetError() error {
	return s.error
}

func (s *SdkError) GetMessage() string {
	return s.message
}

func (s *SdkError) GetErrorCode() ErrorCode {
	return s.errorCode
}

func (s *SdkError) GetStringErrorCode() string {
	return string(s.errorCode)
}

func (s *SdkError) GetParameters() map[string]interface{} {
	res := make(map[string]interface{})
	if s.parameters != nil {
		s.parameters.Range(func(key, val interface{}) bool {
			res[key.(string)] = val
			return true
		})
	}

	return res
}

func (s *SdkError) GetErrorCategories() []ErrorCategory {
	result := make([]ErrorCategory, 0, len(s.categories))
	for c := range s.categories {
		result = append(result, c)
	}
	return result
}

func (s *SdkError) GetErrorMessages() string {
	if s.error == nil {
		return s.message
	}

	return fmt.Sprintf("%s: %s", s.message, s.error.Error())
}

func (s *SdkError) GetListParameters() []interface{} {
	var result []interface{}
	if s.parameters == nil {
		return result
	}

	s.parameters.Range(func(key, val interface{}) bool {
		result = append(result, key, val)
		return true
	})

	return result
}

func (s *SdkError) RemoveCategories(categories ...ErrorCategory) Error {
	if s.categories == nil {
		return s
	}

	for _, c := range categories {
		delete(s.categories, c)
	}
	return s
}

func (s *SdkError) AppendCategories(categories ...ErrorCategory) Error {
	if s.categories == nil {
		s.categories = make(map[ErrorCategory]struct{})
	}

	for _, c := range categories {
		s.categories[c] = struct{}{}
	}
	return s
}
