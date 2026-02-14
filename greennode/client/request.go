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

func NewRequest() IRequest {
	return &request{
		okCodes: make(map[int]struct{}),
	}
}

func (s *request) WithOkCodes(pokCodes ...int) IRequest {
	for _, c := range pokCodes {
		s.okCodes[c] = struct{}{}
	}
	return s
}

func (s *request) WithUserId(puserId string) IRequest {
	return s.WithHeader("portal-user-id", puserId)
}

func (s *request) WithJsonBody(pjsonBody interface{}) IRequest {
	s.JsonBody = pjsonBody
	return s
}

func (s *request) WithJsonResponse(pjsonResponse interface{}) IRequest {
	s.JsonResponse = pjsonResponse
	return s
}

func (s *request) WithJsonError(pjsonError interface{}) IRequest {
	s.JsonError = pjsonError
	return s
}

func (s *request) WithRequestMethod(pmethod requestMethod) IRequest {
	s.Method = pmethod
	return s
}

func (s *request) WithSkipAuth(pskipAuth bool) IRequest {
	s.SkipAuth = pskipAuth
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

func (s *request) SetJsonResponse(pjsonResponse interface{}) {
	s.JsonResponse = pjsonResponse
}

func (s *request) SetJsonError(pjsonError interface{}) {
	s.JsonError = pjsonError
}

func (s *request) ContainsOkCode(pcode ...int) bool {
	for _, c := range pcode {
		if _, ok := s.okCodes[c]; ok {
			return true
		}
	}
	return false
}

func (s *request) WithHeader(pkey, pvalue string) IRequest {
	if pkey == "" || pvalue == "" {
		return s
	}

	if s.MoreHeaders == nil {
		s.MoreHeaders = make(map[string]string)
	}

	s.MoreHeaders[pkey] = pvalue
	return s
}

func (s *request) WithMapHeaders(pheaders map[string]string) IRequest {
	if s.MoreHeaders == nil {
		s.MoreHeaders = make(map[string]string)
	}

	for k, v := range pheaders {
		s.MoreHeaders[k] = v
	}

	return s
}

func (s *request) SkipAuthentication() bool {
	return s.SkipAuth
}
