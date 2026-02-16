package client

type Request struct {
	jsonBody     any
	jsonResponse any
	jsonError    any
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

func NewRequest() *Request {
	return &Request{
		okCodes: make(map[int]struct{}),
	}
}

func (r *Request) WithOKCodes(okCodes ...int) *Request {
	for _, c := range okCodes {
		r.okCodes[c] = struct{}{}
	}
	return r
}

func (r *Request) WithUserID(userID string) *Request {
	return r.WithHeader("portal-user-id", userID)
}

func (r *Request) WithJSONBody(jsonBody any) *Request {
	r.jsonBody = jsonBody
	return r
}

func (r *Request) WithJSONResponse(jsonResponse any) *Request {
	r.jsonResponse = jsonResponse
	return r
}

func (r *Request) WithJSONError(jsonError any) *Request {
	r.jsonError = jsonError
	return r
}

func (r *Request) WithSkipAuth(skipAuth bool) *Request {
	r.skipAuth = skipAuth
	return r
}

func (r *Request) containsOKCode(code ...int) bool {
	for _, c := range code {
		if _, ok := r.okCodes[c]; ok {
			return true
		}
	}
	return false
}

func (r *Request) WithHeader(key, value string) *Request {
	if key == "" || value == "" {
		return r
	}

	if r.moreHeaders == nil {
		r.moreHeaders = make(map[string]string)
	}

	r.moreHeaders[key] = value
	return r
}
