package client

type request struct {
	JsonBody     interface{}
	JsonResponse interface{}
	JsonError    interface{}
	MoreHeaders  map[string]string
	okCodes      map[int]struct{}
	Method       requestMethod
	SkipAuth     bool
}

type requestMethod string

const (
	MethodGet    = requestMethod("GET")
	MethodPost   = requestMethod("POST")
	MethodPut    = requestMethod("PUT")
	MethodPatch  = requestMethod("PATCH")
	MethodDelete = requestMethod("DELETE")
)

func NewRequest() Request {
	return &request{
		okCodes: make(map[int]struct{}),
	}
}

func (s *request) WithOkCodes(okCodes ...int) Request {
	for _, c := range okCodes {
		s.okCodes[c] = struct{}{}
	}
	return s
}

func (s *request) WithUserId(userId string) Request {
	return s.WithHeader("portal-user-id", userId)
}

func (s *request) WithJsonBody(jsonBody interface{}) Request {
	s.JsonBody = jsonBody
	return s
}

func (s *request) WithJsonResponse(jsonResponse interface{}) Request {
	s.JsonResponse = jsonResponse
	return s
}

func (s *request) WithJsonError(jsonError interface{}) Request {
	s.JsonError = jsonError
	return s
}

func (s *request) WithRequestMethod(method requestMethod) Request {
	s.Method = method
	return s
}

func (s *request) WithSkipAuth(skipAuth bool) Request {
	s.SkipAuth = skipAuth
	return s
}

func (s *request) GetRequestBody() interface{} {
	return s.JsonBody
}

func (s *request) GetJsonError() interface{} {
	return s.JsonError
}

func (s *request) GetRequestMethod() string {
	return string(s.Method)
}

func (s *request) GetMoreHeaders() map[string]string {
	return s.MoreHeaders
}

func (s *request) GetJsonResponse() interface{} {
	return s.JsonResponse
}

func (s *request) SetJsonResponse(jsonResponse interface{}) {
	s.JsonResponse = jsonResponse
}

func (s *request) SetJsonError(jsonError interface{}) {
	s.JsonError = jsonError
}

func (s *request) ContainsOkCode(code ...int) bool {
	for _, c := range code {
		if _, ok := s.okCodes[c]; ok {
			return true
		}
	}
	return false
}

func (s *request) WithHeader(key, value string) Request {
	if key == "" || value == "" {
		return s
	}

	if s.MoreHeaders == nil {
		s.MoreHeaders = make(map[string]string)
	}

	s.MoreHeaders[key] = value
	return s
}

func (s *request) WithMapHeaders(headers map[string]string) Request {
	if s.MoreHeaders == nil {
		s.MoreHeaders = make(map[string]string)
	}

	for k, v := range headers {
		s.MoreHeaders[k] = v
	}

	return s
}

func (s *request) SkipAuthentication() bool {
	return s.SkipAuth
}
