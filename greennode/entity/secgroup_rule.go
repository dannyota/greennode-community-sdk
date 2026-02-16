package entity

type SecgroupRule struct {
	ID             string `json:"id"`
	SecgroupID     string `json:"secgroupId"`
	Direction      string `json:"direction"`
	EtherType      string `json:"etherType"`
	Protocol       string `json:"protocol"`
	Description    string `json:"description"`
	RemoteIPPrefix string `json:"remoteIpPrefix"`
	PortRangeMax   int    `json:"portRangeMax"`
	PortRangeMin   int    `json:"portRangeMin"`
}

type ListSecgroupRules struct {
	Items []*SecgroupRule
}

func (l ListSecgroupRules) Len() int {
	return len(l.Items)
}

func (l ListSecgroupRules) Get(i int) *SecgroupRule {
	if i < 0 || i >= len(l.Items) {
		return nil
	}
	return l.Items[i]
}
