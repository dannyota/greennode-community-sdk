package v2

type CreateSecgroupRuleResponse struct {
	Data struct {
		ID             int     `json:"id"`
		UUID           string  `json:"uuid"`
		CreatedAt      string  `json:"createdAt"`
		DeletedAt      *string `json:"deletedAt,omitempty"`
		UpdatedAt      *string `json:"updatedAt,omitempty"`
		Status         string  `json:"status"`
		SecgroupUuid   string  `json:"secgroupUuid"`
		Direction      string  `json:"direction"`
		EtherType      string  `json:"etherType"`
		PortRangeMax   int     `json:"portRangeMax"`
		PortRangeMin   int     `json:"portRangeMin"`
		Protocol       string  `json:"protocol"`
		RuleID         int     `json:"ruleId"`
		Description    string  `json:"description"`
		RemoteIpPrefix string  `json:"remoteIpPrefix"`
		IsSystem       bool    `json:"isSystem"`
	} `json:"data"`
}

func (r *CreateSecgroupRuleResponse) ToEntitySecgroupRule() *SecgroupRule {
	return &SecgroupRule{
		ID:             r.Data.UUID,
		SecgroupID:     r.Data.SecgroupUuid,
		Direction:      r.Data.Direction,
		EtherType:      r.Data.EtherType,
		PortRangeMax:   r.Data.PortRangeMax,
		PortRangeMin:   r.Data.PortRangeMin,
		Protocol:       r.Data.Protocol,
		Description:    r.Data.Description,
		RemoteIPPrefix: r.Data.RemoteIpPrefix,
	}
}

type ListSecgroupRulesBySecgroupIDResponse struct {
	Data []struct {
		ID             string `json:"id"`
		Direction      string `json:"direction"`
		EtherType      string `json:"etherType"`
		Protocol       string `json:"protocol"`
		PortRangeMin   int    `json:"portRangeMin"`
		PortRangeMax   int    `json:"portRangeMax"`
		RemoteIpPrefix string `json:"remoteIpPrefix"`
		RemoteGroupID  string `json:"remoteGroupId"`
		Status         string `json:"status"`
		Description    string `json:"description"`
		CreatedAt      string `json:"createdAt"`
	} `json:"data"`
}

func (r *ListSecgroupRulesBySecgroupIDResponse) ToEntityListSecgroupRules() *ListSecgroupRules {
	lsr := &ListSecgroupRules{
		Items: make([]*SecgroupRule, 0),
	}

	for _, rule := range r.Data {
		lsr.Items = append(lsr.Items, &SecgroupRule{
			ID:             rule.ID,
			Direction:      rule.Direction,
			EtherType:      rule.EtherType,
			Protocol:       rule.Protocol,
			PortRangeMin:   rule.PortRangeMin,
			PortRangeMax:   rule.PortRangeMax,
			RemoteIPPrefix: rule.RemoteIpPrefix,
			Description:    rule.Description,
		})
	}

	return lsr
}
