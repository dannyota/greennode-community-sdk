package v2

import "github.com/dannyota/greennode-community-sdk/v2/greennode/services/common"

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
		LoadBalancerCommon: common.LoadBalancerCommon{LoadBalancerID: lbID},
		ListenerCommon:     common.ListenerCommon{ListenerID: lisID},
	}
}

type CreatePolicyRequest struct {
	common.UserAgent
	common.LoadBalancerCommon
	common.ListenerCommon

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

func (r *CreatePolicyRequest) ToRequestBody() any {
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

func (r *CreatePolicyRequest) ToMap() map[string]any {
	return map[string]any{
		"name":             r.Name,
		"action":           string(r.Action),
		"rules":            r.Rules,
		"redirectPoolId":   r.RedirectPoolID,
		"redirectUrl":      r.RedirectURL,
		"redirectHttpCode": r.RedirectHTTPCode,
		"keepQueryString":  r.KeepQueryString,
	}
}

func (r *CreatePolicyRequest) WithName(name string) *CreatePolicyRequest {
	r.Name = name
	return r
}

func (r *CreatePolicyRequest) WithAction(action PolicyAction) *CreatePolicyRequest {
	r.Action = action
	return r
}

func (r *CreatePolicyRequest) WithRules(rules ...L7RuleRequest) *CreatePolicyRequest {
	r.Rules = rules
	return r
}

func (r *CreatePolicyRequest) WithRedirectPoolID(redirectPoolID string) *CreatePolicyRequest {
	r.RedirectPoolID = redirectPoolID
	return r
}

func (r *CreatePolicyRequest) WithRedirectURL(redirectURL string) *CreatePolicyRequest {
	r.RedirectURL = redirectURL
	return r
}

func (r *CreatePolicyRequest) WithRedirectHTTPCode(redirectHTTPCode int) *CreatePolicyRequest {
	r.RedirectHTTPCode = redirectHTTPCode
	return r
}

func (r *CreatePolicyRequest) WithKeepQueryString(keepQueryString bool) *CreatePolicyRequest {
	r.KeepQueryString = keepQueryString
	return r
}

func (r *CreatePolicyRequest) AddUserAgent(agent ...string) *CreatePolicyRequest {
	r.UserAgent.AddUserAgent(agent...)
	return r
}

// update policy request
func NewUpdatePolicyRequest(lbID, lisID, policyID string) *UpdatePolicyRequest {
	return &UpdatePolicyRequest{
		LoadBalancerCommon: common.LoadBalancerCommon{LoadBalancerID: lbID},
		ListenerCommon:     common.ListenerCommon{ListenerID: lisID},
		PolicyCommon:       common.PolicyCommon{PolicyID: policyID},
	}
}

type UpdatePolicyRequest struct {
	common.UserAgent
	common.LoadBalancerCommon
	common.ListenerCommon
	common.PolicyCommon

	Action           PolicyAction    `json:"action"`
	Rules            []L7RuleRequest `json:"rules"`
	RedirectPoolID   string          `json:"redirectPoolId"`
	RedirectURL      string          `json:"redirectUrl"`
	RedirectHTTPCode int             `json:"redirectHttpCode"`
	KeepQueryString  bool            `json:"keepQueryString"`
}

func (r *UpdatePolicyRequest) ToRequestBody() any {
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

func (r *UpdatePolicyRequest) WithAction(action PolicyAction) *UpdatePolicyRequest {
	r.Action = action
	return r
}

func (r *UpdatePolicyRequest) WithRules(rules ...L7RuleRequest) *UpdatePolicyRequest {
	r.Rules = rules
	return r
}

func (r *UpdatePolicyRequest) WithRedirectPoolID(redirectPoolID string) *UpdatePolicyRequest {
	r.RedirectPoolID = redirectPoolID
	return r
}

func (r *UpdatePolicyRequest) WithRedirectURL(redirectURL string) *UpdatePolicyRequest {
	r.RedirectURL = redirectURL
	return r
}

func (r *UpdatePolicyRequest) WithRedirectHTTPCode(redirectHTTPCode int) *UpdatePolicyRequest {
	r.RedirectHTTPCode = redirectHTTPCode
	return r
}

func (r *UpdatePolicyRequest) WithKeepQueryString(keepQueryString bool) *UpdatePolicyRequest {
	r.KeepQueryString = keepQueryString
	return r
}

func (r *UpdatePolicyRequest) AddUserAgent(agent ...string) *UpdatePolicyRequest {
	r.UserAgent.AddUserAgent(agent...)
	return r
}

// get policy by id request
func NewGetPolicyByIDRequest(lbID, lisID, policyID string) *GetPolicyByIDRequest {
	return &GetPolicyByIDRequest{
		LoadBalancerCommon: common.LoadBalancerCommon{LoadBalancerID: lbID},
		ListenerCommon:     common.ListenerCommon{ListenerID: lisID},
		PolicyCommon:       common.PolicyCommon{PolicyID: policyID},
	}
}

func (r *GetPolicyByIDRequest) AddUserAgent(agent ...string) *GetPolicyByIDRequest {
	r.UserAgent.AddUserAgent(agent...)
	return r
}

type GetPolicyByIDRequest struct {
	common.UserAgent
	common.LoadBalancerCommon
	common.ListenerCommon
	common.PolicyCommon
}

// delete policy by id request
func NewDeletePolicyByIDRequest(lbID, lisID, policyID string) *DeletePolicyByIDRequest {
	return &DeletePolicyByIDRequest{
		LoadBalancerCommon: common.LoadBalancerCommon{LoadBalancerID: lbID},
		ListenerCommon:     common.ListenerCommon{ListenerID: lisID},
		PolicyCommon:       common.PolicyCommon{PolicyID: policyID},
	}
}

func (r *DeletePolicyByIDRequest) AddUserAgent(agent ...string) *DeletePolicyByIDRequest {
	r.UserAgent.AddUserAgent(agent...)
	return r
}

type DeletePolicyByIDRequest struct {
	common.UserAgent
	common.LoadBalancerCommon
	common.ListenerCommon
	common.PolicyCommon
}

type policyPositionRequest struct {
	Position int    `json:"position"`
	PolicyID string `json:"policyId"`
}

func NewReorderPoliciesRequest(lbID, lisID string) *ReorderPoliciesRequest {
	return &ReorderPoliciesRequest{
		LoadBalancerCommon: common.LoadBalancerCommon{LoadBalancerID: lbID},
		ListenerCommon:     common.ListenerCommon{ListenerID: lisID},

		policyPositions: make([]policyPositionRequest, 0),
	}
}

type ReorderPoliciesRequest struct {
	common.UserAgent
	common.LoadBalancerCommon
	common.ListenerCommon

	policyPositions []policyPositionRequest
}

func (r *ReorderPoliciesRequest) AddUserAgent(agent ...string) *ReorderPoliciesRequest {
	r.UserAgent.AddUserAgent(agent...)
	return r
}

func (r *ReorderPoliciesRequest) WithPoliciesOrder(policies []string) *ReorderPoliciesRequest {
	r.policyPositions = make([]policyPositionRequest, len(policies))
	for i, policy := range policies {
		r.policyPositions[i] = policyPositionRequest{
			Position: i + 1,
			PolicyID: policy,
		}
	}
	return r
}

func (r *ReorderPoliciesRequest) ToRequestBody() any {
	return map[string]any{
		"policies": r.policyPositions,
	}
}


// list policies request
func NewListPoliciesRequest(lbID, lisID string) *ListPoliciesRequest {
	return &ListPoliciesRequest{
		LoadBalancerCommon: common.LoadBalancerCommon{LoadBalancerID: lbID},
		ListenerCommon:     common.ListenerCommon{ListenerID: lisID},
	}
}

func (r *ListPoliciesRequest) AddUserAgent(agent ...string) *ListPoliciesRequest {
	r.UserAgent.AddUserAgent(agent...)
	return r
}

type ListPoliciesRequest struct {
	common.UserAgent
	common.LoadBalancerCommon
	common.ListenerCommon
}
