package client

type request struct {
	jsonBody     interface{}
	jsonResponse interface{}
	jsonError    interface{}
	moreHeaders  map[string]string
	okCodes      map[int]struct{}
	method       requestMethod
	skipAuth     bool
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

func (s *request) WithUserID(userID string) Request {
	return s.WithHeader("portal-user-id", userID)
}

func (s *request) WithJSONBody(jsonBody interface{}) Request {
	s.jsonBody = jsonBody
	return s
}

func (s *request) WithJSONResponse(jsonResponse interface{}) Request {
	s.jsonResponse = jsonResponse
	return s
}

func (s *request) WithJSONError(jsonError interface{}) Request {
	s.jsonError = jsonError
	return s
}

func (s *request) WithRequestMethod(method requestMethod) Request {
	s.method = method
	return s
}

func (s *request) WithSkipAuth(skipAuth bool) Request {
	s.skipAuth = skipAuth
	return s
}

func (s *request) RequestBody() interface{} {
	return s.jsonBody
}

func (s *request) JSONError() interface{} {
	return s.jsonError
}

func (s *request) RequestMethod() string {
	return string(s.method)
}

func (s *request) MoreHeaders() map[string]string {
	return s.moreHeaders
}

func (s *request) JSONResponse() interface{} {
	return s.jsonResponse
}

func (s *request) SetJSONResponse(jsonResponse interface{}) {
	s.jsonResponse = jsonResponse
}

func (s *request) SetJSONError(jsonError interface{}) {
	s.jsonError = jsonError
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

	if s.moreHeaders == nil {
		s.moreHeaders = make(map[string]string)
	}

	s.moreHeaders[key] = value
	return s
}

func (s *request) WithMapHeaders(headers map[string]string) Request {
	if s.moreHeaders == nil {
		s.moreHeaders = make(map[string]string)
	}

	for k, v := range headers {
		s.moreHeaders[k] = v
	}

	return s
}

func (s *request) SkipAuthentication() bool {
	return s.skipAuth
}
