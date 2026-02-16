package v2

type GetAccessTokenRequest struct {
	ClientID     string
	ClientSecret string

	GrantType string `json:"grant_type"`
}

func NewGetAccessTokenRequest(clientID, clientSecret string) *GetAccessTokenRequest {
	return &GetAccessTokenRequest{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		GrantType:    "client_credentials",
	}
}


