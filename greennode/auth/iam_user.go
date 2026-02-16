package auth

import (
	"bytes"
	"context"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"regexp"
	"strings"
	"time"
)

const (
	signinBaseURL     = "https://signin.vngcloud.vn"
	loginPath         = "/ap/auth/iam/login"
	tokenURL          = "https://dashboard.console.vngcloud.vn/accounts-api/v1/auth/token"
	dashboardURI      = "https://dashboard.console.vngcloud.vn/"
	dashboardClientID = "c9e78411-f2a2-41ba-a9e4-3c56263c181a"
	twoFAPathMatch    = "/ap/auth/iam/google"
)

// IAMUserAuth holds credentials for IAM user authentication.
type IAMUserAuth struct {
	RootEmail string
	Username  string
	Password  string
	TOTP      TOTPProvider // optional, for accounts with 2FA
}

// Authenticate performs the full OAuth2 PKCE + login + optional TOTP flow
// and returns an access token with its expiry (Unix nanoseconds).
func (a *IAMUserAuth) Authenticate(ctx context.Context) (accessToken string, expiresAt int64, err error) {
	jar, _ := cookiejar.New(nil)
	client := &http.Client{
		Jar: jar,
		CheckRedirect: func(_ *http.Request, _ []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	_, challenge := generatePKCE()

	// GET login page for CSRF token
	loginURL := fmt.Sprintf("%s%s?clientId=%s&responseType=code&codeChallenge=%s&codeChallengeMethod=S256&redirectUri=%s&rootEmail=%s",
		signinBaseURL, loginPath,
		dashboardClientID, challenge,
		url.QueryEscape(dashboardURI),
		url.QueryEscape(a.RootEmail),
	)

	pageBody, err := doGet(ctx, client, loginURL)
	if err != nil {
		return "", 0, fmt.Errorf("auth: GET login page: %w", err)
	}

	csrfToken := extractCSRFToken(string(pageBody))
	if csrfToken == "" {
		return "", 0, fmt.Errorf("auth: no CSRF token on login page")
	}

	// POST credentials
	formData := url.Values{
		"_csrf":     {csrfToken},
		"rootEmail": {a.RootEmail},
		"username":  {a.Username},
		"password":  {a.Password},
	}

	location, err := doPostForm(ctx, client, loginURL, formData)
	if err != nil {
		return "", 0, fmt.Errorf("auth: POST login: %w", err)
	}
	if location == "" {
		return "", 0, fmt.Errorf("auth: login failed (no redirect)")
	}

	// Handle 2FA if needed
	if strings.Contains(location, twoFAPathMatch) {
		if a.TOTP == nil {
			return "", 0, fmt.Errorf("auth: 2FA required but no TOTPProvider configured")
		}
		location, err = a.handle2FA(ctx, client, location)
		if err != nil {
			return "", 0, err
		}
	}

	// Extract auth code from redirect
	authCode := extractAuthCode(location)
	if authCode == "" {
		return "", 0, fmt.Errorf("auth: no authorization code in redirect: %s", location)
	}

	// Exchange code for token
	token, err := exchangeToken(ctx, client, authCode)
	if err != nil {
		return "", 0, err
	}

	// Compute expiresAt as nanoseconds (matching SdkAuthentication convention)
	expiresAt = time.Now().Add(time.Duration(token.ExpiresIn) * time.Second).UnixNano()

	return token.AccessToken, expiresAt, nil
}

func (a *IAMUserAuth) handle2FA(ctx context.Context, client *http.Client, redirectPath string) (string, error) {
	twoFAURL := redirectPath
	if !strings.HasPrefix(redirectPath, "http") {
		twoFAURL = signinBaseURL + redirectPath
	}

	pageBody, err := doGet(ctx, client, twoFAURL)
	if err != nil {
		return "", fmt.Errorf("auth: GET 2FA page: %w", err)
	}

	csrfToken := extractCSRFToken(string(pageBody))
	if csrfToken == "" {
		return "", fmt.Errorf("auth: no CSRF token on 2FA page")
	}

	totpCode, err := a.TOTP.GetCode(ctx)
	if err != nil {
		return "", fmt.Errorf("auth: get TOTP code: %w", err)
	}

	formData := url.Values{
		"_csrf": {csrfToken},
		"token": {totpCode},
	}

	location, err := doPostForm(ctx, client, twoFAURL, formData)
	if err != nil {
		return "", fmt.Errorf("auth: POST 2FA: %w", err)
	}
	if location == "" {
		return "", fmt.Errorf("auth: 2FA failed (no redirect)")
	}

	return location, nil
}

type tokenResponse struct {
	AccessToken string `json:"accessToken"`
	ExpiresIn   int64  `json:"expiresIn"`
}

func exchangeToken(ctx context.Context, client *http.Client, authCode string) (*tokenResponse, error) {
	body, _ := json.Marshal(map[string]string{
		"grantType":   "authorization_code",
		"code":        authCode,
		"redirectUri": dashboardURI,
		"scope":       "openid",
	})

	req, err := http.NewRequestWithContext(ctx, "POST", tokenURL, bytes.NewReader(body))
	if err != nil {
		return nil, fmt.Errorf("auth: create token request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Basic "+base64.StdEncoding.EncodeToString([]byte(dashboardClientID+":")))

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("auth: token exchange request: %w", err)
	}
	defer resp.Body.Close()

	respBody, _ := io.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("auth: token exchange failed (status %d): %s", resp.StatusCode, respBody)
	}

	var token tokenResponse
	if err := json.Unmarshal(respBody, &token); err != nil {
		return nil, fmt.Errorf("auth: parse token response: %w", err)
	}
	if token.AccessToken == "" {
		return nil, fmt.Errorf("auth: empty access token in response")
	}

	return &token, nil
}

// --- helpers ---

func generatePKCE() (verifier, challenge string) {
	b := make([]byte, 32)
	_, _ = rand.Read(b)
	verifier = base64.RawURLEncoding.EncodeToString(b)
	hash := sha256.Sum256([]byte(verifier))
	challenge = base64.RawURLEncoding.EncodeToString(hash[:])
	return
}

// doGet performs a GET that follows redirects (overriding the client's no-redirect policy)
// and returns the response body.
func doGet(ctx context.Context, client *http.Client, reqURL string) ([]byte, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", reqURL, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")

	// Use a temporary client that follows redirects (the main client blocks them).
	tempClient := &http.Client{Jar: client.Jar, Transport: client.Transport}
	resp, err := tempClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return io.ReadAll(resp.Body)
}

// doPostForm POSTs form data with no-redirect and returns the Location header.
func doPostForm(ctx context.Context, client *http.Client, postURL string, form url.Values) (string, error) {
	req, err := http.NewRequestWithContext(ctx, "POST", postURL, strings.NewReader(form.Encode()))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Origin", signinBaseURL)
	req.Header.Set("Referer", postURL)

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	return resp.Header.Get("Location"), nil
}

var (
	csrfRe1 = regexp.MustCompile(`content="([^"]+)"[^>]*name="csrf-token"`)
	csrfRe2 = regexp.MustCompile(`name="_csrf"[^>]*value="([^"]+)"`)
)

func extractCSRFToken(html string) string {
	if m := csrfRe1.FindStringSubmatch(html); len(m) > 1 {
		return m[1]
	}
	if m := csrfRe2.FindStringSubmatch(html); len(m) > 1 {
		return m[1]
	}
	return ""
}

func extractAuthCode(location string) string {
	u, err := url.Parse(location)
	if err != nil {
		return ""
	}
	return u.Query().Get("code")
}
