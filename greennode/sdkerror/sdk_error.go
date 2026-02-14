package sdkerror

import (
	"errors"
	"fmt"
	"sync"
)

var (
	_ IError = new(SdkError)
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

func (s *SdkError) IsError(perrCode ErrorCode) bool {
	return s.errorCode == perrCode
}

func (s *SdkError) IsErrorAny(perrCodes ...ErrorCode) bool {
	for _, perrCode := range perrCodes {
		if s.errorCode == perrCode {
			return true
		}
	}

	return false
}

func (s *SdkError) IsCategory(pcategory ErrorCategory) bool {
	if s.categories == nil {
		return false
	}

	_, ok := s.categories[pcategory]
	return ok
}

func (s *SdkError) IsCategories(pcategories ...ErrorCategory) bool {
	if s.categories == nil {
		return false
	}

	for _, c := range pcategories {
		if _, ok := s.categories[c]; ok {
			return true
		}
	}
	return false
}

func (s *SdkError) WithErrorCode(perrCode ErrorCode) IError {
	s.errorCode = perrCode
	return s
}

func (s *SdkError) WithMessage(pmsg string) IError {
	s.message = pmsg
	return s
}

func (s *SdkError) WithErrors(perrs ...error) IError {
	if len(perrs) == 0 {
		return s
	}

	if len(perrs) == 1 {
		s.error = perrs[0]
		return s
	}

	for _, err := range perrs {
		s.error = errors.Join(s.error, err)
	}

	return s
}

func (s *SdkError) WithErrorCategories(pcategories ...ErrorCategory) IError {
	if s.categories == nil {
		s.categories = make(map[ErrorCategory]struct{})
	}
	for _, c := range pcategories {
		s.categories[c] = struct{}{}
	}

	return s
}

func (s *SdkError) WithParameters(pparams map[string]interface{}) IError {
	if s.parameters == nil {
		s.parameters = new(sync.Map)
		return s
	}

	for key, val := range pparams {
		s.parameters.Store(key, val)
	}

	return s
}

func (s *SdkError) WithKVparameters(pparams ...interface{}) IError {
	if s.parameters == nil {
		s.parameters = new(sync.Map)
	}

	// Always make sure that the length of pparams is even
	if len(pparams)%2 != 0 {
		pparams = append(pparams, nil)
	}

	for i := 0; i < len(pparams); i += 2 {
		key, ok := pparams[i].(string)
		if !ok {
			continue
		}

		s.parameters.Store(key, pparams[i+1])
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

func (s *SdkError) RemoveCategories(pcategories ...ErrorCategory) IError {
	if s.categories == nil {
		return s
	}

	for _, c := range pcategories {
		delete(s.categories, c)
	}
	return s
}

func (s *SdkError) AppendCategories(pcategories ...ErrorCategory) IError {
	if s.categories == nil {
		s.categories = make(map[ErrorCategory]struct{})
	}

	for _, c := range pcategories {
		s.categories[c] = struct{}{}
	}
	return s
}
