package option

import "net/http"

// ClientOption configures the SDK client.
// Use the With* constructors; do not implement this interface directly.
type ClientOption interface {
	applyOption(*Settings)
}

// Settings is the resolved set of options (internal use by greennode package).
type Settings struct {
	HTTPClient *http.Client
	Transport  http.RoundTripper
	UserAgent  string
}

// Apply merges all options into a Settings value.
func Apply(opts []ClientOption) *Settings {
	s := &Settings{}
	for _, o := range opts {
		o.applyOption(s)
	}
	return s
}

// WithHTTPClient returns a ClientOption that specifies the HTTP client to use
// as the basis of communications. When used, it takes precedence over
// WithTransport. The SDK's retry and authentication logic still applies.
func WithHTTPClient(c *http.Client) ClientOption {
	return withHTTPClient{c}
}

type withHTTPClient struct{ client *http.Client }

func (w withHTTPClient) applyOption(s *Settings) { s.HTTPClient = w.client }

// WithTransport returns a ClientOption that sets a custom http.RoundTripper
// on the SDK's default http.Client. Useful for adding rate limiters, tracing,
// or proxy transports without replacing the whole http.Client.
// Ignored if WithHTTPClient is also provided.
func WithTransport(t http.RoundTripper) ClientOption {
	return withTransport{t}
}

type withTransport struct{ transport http.RoundTripper }

func (w withTransport) applyOption(s *Settings) { s.Transport = w.transport }

// WithUserAgent returns a ClientOption that sets the User-Agent header on all
// requests. Takes precedence over Config.UserAgent when both are set.
func WithUserAgent(ua string) ClientOption {
	return withUserAgent{ua}
}

type withUserAgent struct{ ua string }

func (w withUserAgent) applyOption(s *Settings) { s.UserAgent = w.ua }
