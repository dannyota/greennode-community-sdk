package entity

type Policy struct {
	UUID             string    `json:"uuid"`
	Name             string    `json:"name"`
	Description      string    `json:"description"`
	RedirectPoolID   string    `json:"redirectPoolId"`
	RedirectPoolName string    `json:"redirectPoolName"`
	Action           string    `json:"action"`
	RedirectURL      string    `json:"redirectUrl"`
	RedirectHTTPCode int       `json:"redirectHttpCode"`
	KeepQueryString  bool      `json:"keepQueryString"`
	Position         int       `json:"position"`
	L7Rules          []*L7Rule `json:"l7Rules"`
	DisplayStatus    string    `json:"displayStatus"`
	CreatedAt        string    `json:"createdAt"`
	UpdatedAt        string    `json:"updatedAt"`
	ProgressStatus   string    `json:"progressStatus"`
}

type L7Rule struct {
	UUID               string `json:"uuid"`
	CompareType        string `json:"compareType"`
	RuleValue          string `json:"ruleValue"`
	RuleType           string `json:"ruleType"`
	ProvisioningStatus string `json:"provisioningStatus"`
	OperatingStatus    string `json:"operatingStatus"`
}

type ListPolicies struct {
	Items []*Policy
}
