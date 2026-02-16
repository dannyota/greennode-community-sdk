package v2

type (
	PolicyAction      string
	PolicyCompareType string
	PolicyRuleType    string
)

const (
	PolicyActionREJECT         PolicyAction = "REJECT"
	PolicyActionREDIRECTTOURL  PolicyAction = "REDIRECT_TO_URL"
	PolicyActionREDIRECTTOPOOL PolicyAction = "REDIRECT_TO_POOL"

	PolicyCompareTypeCONTAINS   PolicyCompareType = "CONTAINS"
	PolicyCompareTypeEQUALS     PolicyCompareType = "EQUAL_TO"
	PolicyCompareTypeREGEX      PolicyCompareType = "REGEX"
	PolicyCompareTypeSTARTSWITH PolicyCompareType = "STARTS_WITH"
	PolicyCompareTypeENDSWITH   PolicyCompareType = "ENDS_WITH"

	PolicyRuleTypeHOSTNAME PolicyRuleType = "HOST_NAME"
	PolicyRuleTypePATH     PolicyRuleType = "PATH"
)

// create policy request
func NewCreatePolicyRequest(lbID, lisID string) *CreatePolicyRequest {
	return &CreatePolicyRequest{
		LoadBalancerID: lbID,
		ListenerID:     lisID,
	}
}

type CreatePolicyRequest struct {
	LoadBalancerID string
	ListenerID     string

	Name             string          `json:"name"`
	Action           PolicyAction    `json:"action"`
	Rules            []L7RuleRequest `json:"rules"`
	RedirectPoolID   string          `json:"redirectPoolId"`
	RedirectURL      string          `json:"redirectUrl"`
	RedirectHTTPCode int             `json:"redirectHttpCode"`
	KeepQueryString  bool            `json:"keepQueryString"`
}

type L7RuleRequest struct {
	CompareType PolicyCompareType `json:"compareType"`
	RuleType    PolicyRuleType    `json:"ruleType"`
	RuleValue   string            `json:"ruleValue"`
}

func (r *CreatePolicyRequest) toRequestBody() any {
	switch r.Action {
	case PolicyActionREJECT:
		return struct {
			Name   string          `json:"name"`
			Action string          `json:"action"`
			Rules  []L7RuleRequest `json:"rules"`
		}{
			Name:   r.Name,
			Action: string(r.Action),
			Rules:  r.Rules,
		}
	case PolicyActionREDIRECTTOURL:
		return struct {
			Name             string          `json:"name"`
			Action           string          `json:"action"`
			Rules            []L7RuleRequest `json:"rules"`
			RedirectURL      string          `json:"redirectUrl"`
			RedirectHTTPCode int             `json:"redirectHttpCode"`
			KeepQueryString  bool            `json:"keepQueryString"`
		}{
			Name:             r.Name,
			Action:           string(r.Action),
			Rules:            r.Rules,
			RedirectURL:      r.RedirectURL,
			RedirectHTTPCode: r.RedirectHTTPCode,
			KeepQueryString:  r.KeepQueryString,
		}
	case PolicyActionREDIRECTTOPOOL:
		return struct {
			Name           string          `json:"name"`
			Action         string          `json:"action"`
			Rules          []L7RuleRequest `json:"rules"`
			RedirectPoolID string          `json:"redirectPoolId"`
		}{
			Name:           r.Name,
			Action:         string(r.Action),
			Rules:          r.Rules,
			RedirectPoolID: r.RedirectPoolID,
		}
	}
	return nil
}

// update policy request
func NewUpdatePolicyRequest(lbID, lisID, policyID string) *UpdatePolicyRequest {
	return &UpdatePolicyRequest{
		LoadBalancerID: lbID,
		ListenerID:     lisID,
		PolicyID:       policyID,
	}
}

type UpdatePolicyRequest struct {
	LoadBalancerID string
	ListenerID     string
	PolicyID       string

	Action           PolicyAction    `json:"action"`
	Rules            []L7RuleRequest `json:"rules"`
	RedirectPoolID   string          `json:"redirectPoolId"`
	RedirectURL      string          `json:"redirectUrl"`
	RedirectHTTPCode int             `json:"redirectHttpCode"`
	KeepQueryString  bool            `json:"keepQueryString"`
}

func (r *UpdatePolicyRequest) toRequestBody() any {
	switch r.Action {
	case PolicyActionREJECT:
		return struct {
			Action string          `json:"action"`
			Rules  []L7RuleRequest `json:"rules"`
		}{
			Action: string(r.Action),
			Rules:  r.Rules,
		}
	case PolicyActionREDIRECTTOURL:
		return struct {
			Action           string          `json:"action"`
			Rules            []L7RuleRequest `json:"rules"`
			RedirectURL      string          `json:"redirectUrl"`
			RedirectHTTPCode int             `json:"redirectHttpCode"`
			KeepQueryString  bool            `json:"keepQueryString"`
		}{
			Action:           string(r.Action),
			Rules:            r.Rules,
			RedirectURL:      r.RedirectURL,
			RedirectHTTPCode: r.RedirectHTTPCode,
			KeepQueryString:  r.KeepQueryString,
		}
	case PolicyActionREDIRECTTOPOOL:
		return struct {
			Action         string          `json:"action"`
			Rules          []L7RuleRequest `json:"rules"`
			RedirectPoolID string          `json:"redirectPoolId"`
		}{
			Action:         string(r.Action),
			Rules:          r.Rules,
			RedirectPoolID: r.RedirectPoolID,
		}
	}
	return nil
}

// get policy by id request
func NewGetPolicyByIDRequest(lbID, lisID, policyID string) *GetPolicyByIDRequest {
	return &GetPolicyByIDRequest{
		LoadBalancerID: lbID,
		ListenerID:     lisID,
		PolicyID:       policyID,
	}
}

type GetPolicyByIDRequest struct {
	LoadBalancerID string
	ListenerID     string
	PolicyID       string
}

// delete policy by id request
func NewDeletePolicyByIDRequest(lbID, lisID, policyID string) *DeletePolicyByIDRequest {
	return &DeletePolicyByIDRequest{
		LoadBalancerID: lbID,
		ListenerID:     lisID,
		PolicyID:       policyID,
	}
}

type DeletePolicyByIDRequest struct {
	LoadBalancerID string
	ListenerID     string
	PolicyID       string
}

type PolicyPosition struct {
	Position int    `json:"position"`
	PolicyID string `json:"policyId"`
}

func NewReorderPoliciesRequest(lbID, lisID string) *ReorderPoliciesRequest {
	return &ReorderPoliciesRequest{
		LoadBalancerID: lbID,
		ListenerID:     lisID,

		PolicyPositions: make([]PolicyPosition, 0),
	}
}

type ReorderPoliciesRequest struct {
	LoadBalancerID string
	ListenerID     string

	PolicyPositions []PolicyPosition
}

func (r *ReorderPoliciesRequest) toRequestBody() any {
	return map[string]any{
		"policies": r.PolicyPositions,
	}
}


// list policies request
func NewListPoliciesRequest(lbID, lisID string) *ListPoliciesRequest {
	return &ListPoliciesRequest{
		LoadBalancerID: lbID,
		ListenerID:     lisID,
	}
}

type ListPoliciesRequest struct {
	LoadBalancerID string
	ListenerID     string
}
