package sdkerror

import (
	"regexp"
	"strings"
)

// matchFunc tests whether an error response matches a known error pattern.
// lowerMsg is strings.ToLower(strings.TrimSpace(errResp.GetMessage())).
// resp is the original ErrorResponse (for access to Err() etc.).
type matchFunc func(lowerMsg string, resp ErrorResponse) bool

type classifier struct {
	match    matchFunc
	category ErrorCategory // "" means no category
	msgFmt   string        // if set, use fmt.Sprintf(msgFmt, resp.GetMessage())
	catchAll bool          // if true, matches when error code is still EcUnknownError
}

var classifierRegistry = map[ErrorCode]*classifier{}

func register(code ErrorCode, c *classifier) {
	classifierRegistry[code] = c
}

// containsAny returns a matchFunc that checks if lowerMsg contains any of the patterns.
func containsAny(patterns ...string) matchFunc {
	return func(lowerMsg string, _ ErrorResponse) bool {
		for _, p := range patterns {
			if strings.Contains(lowerMsg, p) {
				return true
			}
		}
		return false
	}
}

// matchRegexps returns a matchFunc that checks if lowerMsg matches any of the regexps.
func matchRegexps(res ...*regexp.Regexp) matchFunc {
	return func(lowerMsg string, _ ErrorResponse) bool {
		for _, re := range res {
			if re.FindString(lowerMsg) != "" {
				return true
			}
		}
		return false
	}
}

// matchErrCode returns a matchFunc that checks errResp.Err().Error() against
// the given codes (case-insensitive comparison after TrimSpace+ToUpper).
func matchErrCode(codes ...string) matchFunc {
	return func(_ string, resp ErrorResponse) bool {
		if resp.Err() == nil {
			return false
		}
		errStr := strings.ToUpper(strings.TrimSpace(resp.Err().Error()))
		for _, c := range codes {
			if errStr == c {
				return true
			}
		}
		return false
	}
}

// matchAnyOf combines multiple matchFuncs with OR logic.
func matchAnyOf(fns ...matchFunc) matchFunc {
	return func(lowerMsg string, resp ErrorResponse) bool {
		for _, fn := range fns {
			if fn(lowerMsg, resp) {
				return true
			}
		}
		return false
	}
}
