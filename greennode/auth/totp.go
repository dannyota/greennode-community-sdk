package auth

import (
	"context"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base32"
	"encoding/binary"
	"fmt"
	"strings"
	"time"
)

// TOTPProvider returns a TOTP code for 2FA.
type TOTPProvider interface {
	GetCode(ctx context.Context) (string, error)
}

// TOTPFunc adapts a plain function into a TOTPProvider.
// Use this to integrate any external TOTP source (Vault, env vars, prompts, etc).
type TOTPFunc func(ctx context.Context) (string, error)

// GetCode delegates to the wrapped function.
func (f TOTPFunc) GetCode(ctx context.Context) (string, error) {
	return f(ctx)
}

// SecretTOTP computes TOTP codes from a base32-encoded shared secret (RFC 6238).
type SecretTOTP struct {
	Secret string // base32-encoded TOTP secret
}

// GetCode generates a 6-digit TOTP code for the current 30-second interval.
func (s *SecretTOTP) GetCode(_ context.Context) (string, error) {
	key, err := base32.StdEncoding.WithPadding(base32.NoPadding).DecodeString(
		strings.ToUpper(strings.TrimRight(s.Secret, "=")),
	)
	if err != nil {
		return "", fmt.Errorf("auth: decode TOTP secret: %w", err)
	}

	counter := uint64(time.Now().Unix() / 30)
	buf := make([]byte, 8)
	binary.BigEndian.PutUint64(buf, counter)

	mac := hmac.New(sha1.New, key)
	mac.Write(buf)
	sum := mac.Sum(nil)

	offset := sum[len(sum)-1] & 0x0f
	code := binary.BigEndian.Uint32(sum[offset:offset+4]) & 0x7fffffff

	return fmt.Sprintf("%06d", code%1_000_000), nil
}
