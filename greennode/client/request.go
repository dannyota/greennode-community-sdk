package client

type request struct {
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

func NewRequest() Request {
	return &request{
		okCodes: make(map[int]struct{}),
	}
}

func (r *request) WithOkCodes(okCodes ...int) Request {
	for _, c := range okCodes {
		r.okCodes[c] = struct{}{}
	}
	return r
}

func (r *request) WithUserID(userID string) Request {
	return r.WithHeader("portal-user-id", userID)
}

func (r *request) WithJSONBody(jsonBody any) Request {
	r.jsonBody = jsonBody
	return r
}

func (r *request) WithJSONResponse(jsonResponse any) Request {
	r.jsonResponse = jsonResponse
	return r
}

func (r *request) WithJSONError(jsonError any) Request {
	r.jsonError = jsonError
	return r
}

func (r *request) WithRequestMethod(method requestMethod) Request {
	r.method = method
	return r
}

func (r *request) WithSkipAuth(skipAuth bool) Request {
	r.skipAuth = skipAuth
	return r
}

func (r *request) RequestBody() any {
	return r.jsonBody
}

func (r *request) JSONError() any {
	return r.jsonError
}

func (r *request) RequestMethod() string {
	return string(r.method)
}

func (r *request) MoreHeaders() map[string]string {
	return r.moreHeaders
}

func (r *request) JSONResponse() any {
	return r.jsonResponse
}

func (r *request) SetJSONResponse(jsonResponse any) {
	r.jsonResponse = jsonResponse
}

func (r *request) SetJSONError(jsonError any) {
	r.jsonError = jsonError
}

func (r *request) ContainsOkCode(code ...int) bool {
	for _, c := range code {
		if _, ok := r.okCodes[c]; ok {
			return true
		}
	}
	return false
}

func (r *request) WithHeader(key, value string) Request {
	if key == "" || value == "" {
		return r
	}

	if r.moreHeaders == nil {
		r.moreHeaders = make(map[string]string)
	}

	r.moreHeaders[key] = value
	return r
}

func (r *request) WithMapHeaders(headers map[string]string) Request {
	if r.moreHeaders == nil {
		r.moreHeaders = make(map[string]string)
	}

	for k, v := range headers {
		r.moreHeaders[k] = v
	}

	return r
}

func (r *request) SkipAuthentication() bool {
	return r.skipAuth
}
