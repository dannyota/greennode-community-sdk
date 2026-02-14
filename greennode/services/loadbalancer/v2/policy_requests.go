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
func NewCreatePolicyRequest(lbID, lisID string) ICreatePolicyRequest {
	return &CreatePolicyRequest{
		LoadBalancerCommon: common.LoadBalancerCommon{LoadBalancerID: lbID},
		ListenerCommon:     common.ListenerCommon{ListenerID: lisID},
	}
}

var _ ICreatePolicyRequest = &CreatePolicyRequest{}

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

func (s *CreatePolicyRequest) ToRequestBody() interface{} {
	switch s.Action {
	case PolicyActionREJECT:
		return struct {
			Name   string          `json:"name"`
			Action string          `json:"action"`
			Rules  []L7RuleRequest `json:"rules"`
		}{
			Name:   s.Name,
			Action: string(s.Action),
			Rules:  s.Rules,
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
			Name:             s.Name,
			Action:           string(s.Action),
			Rules:            s.Rules,
			RedirectURL:      s.RedirectURL,
			RedirectHTTPCode: s.RedirectHTTPCode,
			KeepQueryString:  s.KeepQueryString,
		}
	case PolicyActionREDIRECTTOPOOL:
		return struct {
			Name           string          `json:"name"`
			Action         string          `json:"action"`
			Rules          []L7RuleRequest `json:"rules"`
			RedirectPoolID string          `json:"redirectPoolId"`
		}{
			Name:           s.Name,
			Action:         string(s.Action),
			Rules:          s.Rules,
			RedirectPoolID: s.RedirectPoolID,
		}
	}
	return nil
}

func (s *CreatePolicyRequest) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"name":             s.Name,
		"action":           string(s.Action),
		"rules":            s.Rules,
		"redirectPoolId":   s.RedirectPoolID,
		"redirectUrl":      s.RedirectURL,
		"redirectHttpCode": s.RedirectHTTPCode,
		"keepQueryString":  s.KeepQueryString,
	}
}

func (s *CreatePolicyRequest) WithName(name string) ICreatePolicyRequest {
	s.Name = name
	return s
}

func (s *CreatePolicyRequest) WithAction(action PolicyAction) ICreatePolicyRequest {
	s.Action = action
	return s
}

func (s *CreatePolicyRequest) WithRules(rules ...L7RuleRequest) ICreatePolicyRequest {
	s.Rules = rules
	return s
}

func (s *CreatePolicyRequest) WithRedirectPoolID(redirectPoolID string) ICreatePolicyRequest {
	s.RedirectPoolID = redirectPoolID
	return s
}

func (s *CreatePolicyRequest) WithRedirectURL(redirectURL string) ICreatePolicyRequest {
	s.RedirectURL = redirectURL
	return s
}

func (s *CreatePolicyRequest) WithRedirectHTTPCode(redirectHTTPCode int) ICreatePolicyRequest {
	s.RedirectHTTPCode = redirectHTTPCode
	return s
}

func (s *CreatePolicyRequest) WithKeepQueryString(keepQueryString bool) ICreatePolicyRequest {
	s.KeepQueryString = keepQueryString
	return s
}

func (s *CreatePolicyRequest) AddUserAgent(agent ...string) ICreatePolicyRequest {
	s.UserAgent.AddUserAgent(agent...)
	return s
}

// update policy request
func NewUpdatePolicyRequest(lbID, lisID, policyID string) IUpdatePolicyRequest {
	return &UpdatePolicyRequest{
		LoadBalancerCommon: common.LoadBalancerCommon{LoadBalancerID: lbID},
		ListenerCommon:     common.ListenerCommon{ListenerID: lisID},
		PolicyCommon:       common.PolicyCommon{PolicyID: policyID},
	}
}

var _ IUpdatePolicyRequest = &UpdatePolicyRequest{}

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

func (s *UpdatePolicyRequest) ToRequestBody() interface{} {
	switch s.Action {
	case PolicyActionREJECT:
		return struct {
			Action string          `json:"action"`
			Rules  []L7RuleRequest `json:"rules"`
		}{
			Action: string(s.Action),
			Rules:  s.Rules,
		}
	case PolicyActionREDIRECTTOURL:
		return struct {
			Action           string          `json:"action"`
			Rules            []L7RuleRequest `json:"rules"`
			RedirectURL      string          `json:"redirectUrl"`
			RedirectHTTPCode int             `json:"redirectHttpCode"`
			KeepQueryString  bool            `json:"keepQueryString"`
		}{
			Action:           string(s.Action),
			Rules:            s.Rules,
			RedirectURL:      s.RedirectURL,
			RedirectHTTPCode: s.RedirectHTTPCode,
			KeepQueryString:  s.KeepQueryString,
		}
	case PolicyActionREDIRECTTOPOOL:
		return struct {
			Action         string          `json:"action"`
			Rules          []L7RuleRequest `json:"rules"`
			RedirectPoolID string          `json:"redirectPoolId"`
		}{
			Action:         string(s.Action),
			Rules:          s.Rules,
			RedirectPoolID: s.RedirectPoolID,
		}
	}
	return nil
}

func (s *UpdatePolicyRequest) WithAction(action PolicyAction) IUpdatePolicyRequest {
	s.Action = action
	return s
}

func (s *UpdatePolicyRequest) WithRules(rules ...L7RuleRequest) IUpdatePolicyRequest {
	s.Rules = rules
	return s
}

func (s *UpdatePolicyRequest) WithRedirectPoolID(redirectPoolID string) IUpdatePolicyRequest {
	s.RedirectPoolID = redirectPoolID
	return s
}

func (s *UpdatePolicyRequest) WithRedirectURL(redirectURL string) IUpdatePolicyRequest {
	s.RedirectURL = redirectURL
	return s
}

func (s *UpdatePolicyRequest) WithRedirectHTTPCode(redirectHTTPCode int) IUpdatePolicyRequest {
	s.RedirectHTTPCode = redirectHTTPCode
	return s
}

func (s *UpdatePolicyRequest) WithKeepQueryString(keepQueryString bool) IUpdatePolicyRequest {
	s.KeepQueryString = keepQueryString
	return s
}

func (s *UpdatePolicyRequest) AddUserAgent(agent ...string) IUpdatePolicyRequest {
	s.UserAgent.AddUserAgent(agent...)
	return s
}

// get policy by id request
func NewGetPolicyByIDRequest(lbID, lisID, policyID string) IGetPolicyByIDRequest {
	return &GetPolicyByIDRequest{
		LoadBalancerCommon: common.LoadBalancerCommon{LoadBalancerID: lbID},
		ListenerCommon:     common.ListenerCommon{ListenerID: lisID},
		PolicyCommon:       common.PolicyCommon{PolicyID: policyID},
	}
}

func (s *GetPolicyByIDRequest) AddUserAgent(agent ...string) IGetPolicyByIDRequest {
	s.UserAgent.AddUserAgent(agent...)
	return s
}

var _ IGetPolicyByIDRequest = &GetPolicyByIDRequest{}

type GetPolicyByIDRequest struct {
	common.UserAgent
	common.LoadBalancerCommon
	common.ListenerCommon
	common.PolicyCommon
}

// delete policy by id request
func NewDeletePolicyByIDRequest(lbID, lisID, policyID string) IDeletePolicyByIDRequest {
	return &DeletePolicyByIDRequest{
		LoadBalancerCommon: common.LoadBalancerCommon{LoadBalancerID: lbID},
		ListenerCommon:     common.ListenerCommon{ListenerID: lisID},
		PolicyCommon:       common.PolicyCommon{PolicyID: policyID},
	}
}

func (s *DeletePolicyByIDRequest) AddUserAgent(agent ...string) IDeletePolicyByIDRequest {
	s.UserAgent.AddUserAgent(agent...)
	return s
}

var _ IDeletePolicyByIDRequest = &DeletePolicyByIDRequest{}

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

func NewReorderPoliciesRequest(lbID, lisID string) IReorderPoliciesRequest {
	return &ReorderPoliciesRequest{
		LoadBalancerCommon: common.LoadBalancerCommon{LoadBalancerID: lbID},
		ListenerCommon:     common.ListenerCommon{ListenerID: lisID},

		policyPositions: make([]policyPositionRequest, 0),
	}
}

var _ IReorderPoliciesRequest = &ReorderPoliciesRequest{}

type ReorderPoliciesRequest struct {
	common.UserAgent
	common.LoadBalancerCommon
	common.ListenerCommon

	policyPositions []policyPositionRequest
}

func (s *ReorderPoliciesRequest) AddUserAgent(agent ...string) IReorderPoliciesRequest {
	s.UserAgent.AddUserAgent(agent...)
	return s
}

func (s *ReorderPoliciesRequest) WithPoliciesOrder(policies []string) IReorderPoliciesRequest {
	s.policyPositions = make([]policyPositionRequest, len(policies))
	for i, policy := range policies {
		s.policyPositions[i] = policyPositionRequest{
			Position: i + 1,
			PolicyID: policy,
		}
	}
	return s
}

func (s *ReorderPoliciesRequest) ToRequestBody() interface{} {
	return map[string]interface{}{
		"policies": s.policyPositions,
	}
}

// --------------------------------------------------------

// list policies request
func NewListPoliciesRequest(lbID, lisID string) IListPoliciesRequest {
	return &ListPoliciesRequest{
		LoadBalancerCommon: common.LoadBalancerCommon{LoadBalancerID: lbID},
		ListenerCommon:     common.ListenerCommon{ListenerID: lisID},
	}
}

func (s *ListPoliciesRequest) AddUserAgent(agent ...string) IListPoliciesRequest {
	s.UserAgent.AddUserAgent(agent...)
	return s
}

var _ IListPoliciesRequest = &ListPoliciesRequest{}

type ListPoliciesRequest struct {
	common.UserAgent
	common.LoadBalancerCommon
	common.ListenerCommon
}
