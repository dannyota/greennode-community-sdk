package v2

type responseData struct {
	UUID             string   `json:"uuid"`
	Name             string   `json:"name"`
	Description      string   `json:"description"`
	RedirectPoolID   string   `json:"redirectPoolId"`
	RedirectPoolName string   `json:"redirectPoolName"`
	Action           string   `json:"action"`
	RedirectURL      string   `json:"redirectUrl"`
	RedirectHTTPCode int      `json:"redirectHttpCode"`
	KeepQueryString  bool     `json:"keepQueryString"`
	Position         int      `json:"position"`
	L7Rules          []l7Rule `json:"l7Rules"`
	DisplayStatus    string   `json:"displayStatus"`
	CreatedAt        string   `json:"createdAt"`
	UpdatedAt        string   `json:"updatedAt"`
	ProgressStatus   string   `json:"progressStatus"`
}

type l7Rule struct {
	UUID               string `json:"uuid"`
	CompareType        string `json:"compareType"`
	RuleValue          string `json:"ruleValue"`
	RuleType           string `json:"ruleType"`
	CreatedAt          string `json:"createdAt"`
	UpdatedAt          string `json:"updatedAt"`
	ProvisioningStatus string `json:"provisioningStatus"`
	OperatingStatus    string `json:"operatingStatus"`
}

func (l *l7Rule) toL7Rule() *L7Rule {
	return &L7Rule{
		UUID:               l.UUID,
		CompareType:        l.CompareType,
		RuleValue:          l.RuleValue,
		RuleType:           l.RuleType,
		ProvisioningStatus: l.ProvisioningStatus,
		OperatingStatus:    l.OperatingStatus,
	}
}

// ListPoliciesResponse

type ListPoliciesResponse struct {
	ListData  []responseData `json:"data"`
	Page      int            `json:"page"`
	PageSize  int            `json:"pageSize"`
	TotalPage int            `json:"totalPage"`
	TotalItem int            `json:"totalItem"`
}

func (r *ListPoliciesResponse) ToEntityListPolicies() *ListPolicies {
	if r == nil || r.ListData == nil {
		return nil
	}

	items := make([]*Policy, 0, len(r.ListData))
	for _, item := range r.ListData {
		l7Rule := make([]*L7Rule, 0)
		for _, rule := range item.L7Rules {
			l7Rule = append(l7Rule, rule.toL7Rule())
		}
		items = append(items, &Policy{
			UUID:             item.UUID,
			Name:             item.Name,
			Description:      item.Description,
			RedirectPoolID:   item.RedirectPoolID,
			RedirectPoolName: item.RedirectPoolName,
			Action:           item.Action,
			RedirectURL:      item.RedirectURL,
			RedirectHTTPCode: item.RedirectHTTPCode,
			KeepQueryString:  item.KeepQueryString,
			Position:         item.Position,
			DisplayStatus:    item.DisplayStatus,
			CreatedAt:        item.CreatedAt,
			UpdatedAt:        item.UpdatedAt,
			ProgressStatus:   item.ProgressStatus,
			L7Rules:          l7Rule,
		})
	}
	return &ListPolicies{
		Items: items,
	}
}

// CreatePolicyResponse

type CreatePolicyResponse struct {
	UUID string `json:"uuid"`
}

func (r *CreatePolicyResponse) ToEntityPolicy() *Policy {
	if r == nil {
		return nil
	}
	return &Policy{
		UUID: r.UUID,
	}
}

// GetPolicyResponse

type GetPolicyResponse struct {
	Data responseData `json:"data"`
}

func (r *GetPolicyResponse) ToEntityPolicy() *Policy {
	if r == nil {
		return nil
	}
	l7Rule := make([]*L7Rule, 0)
	for _, rule := range r.Data.L7Rules {
		l7Rule = append(l7Rule, rule.toL7Rule())
	}
	return &Policy{
		UUID:             r.Data.UUID,
		Name:             r.Data.Name,
		Description:      r.Data.Description,
		RedirectPoolID:   r.Data.RedirectPoolID,
		RedirectPoolName: r.Data.RedirectPoolName,
		Action:           r.Data.Action,
		RedirectURL:      r.Data.RedirectURL,
		RedirectHTTPCode: r.Data.RedirectHTTPCode,
		KeepQueryString:  r.Data.KeepQueryString,
		Position:         r.Data.Position,
		DisplayStatus:    r.Data.DisplayStatus,
		CreatedAt:        r.Data.CreatedAt,
		UpdatedAt:        r.Data.UpdatedAt,
		ProgressStatus:   r.Data.ProgressStatus,
		L7Rules:          l7Rule,
	}
}

// UpdatePolicyResponse and DeletePolicyResponse not have response data
