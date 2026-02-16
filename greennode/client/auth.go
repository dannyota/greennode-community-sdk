package client

import "time"

type Token struct {
	AccessToken string
	ExpiresAt   int64
}

func (t *Token) NeedsReauth() bool {
	if t.AccessToken == "" {
		return true
	}

	ea := time.Unix(0, t.ExpiresAt)
	return time.Until(ea) < 5*time.Minute
}
