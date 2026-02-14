package entity

type SecgroupRule struct {
	ID             string
	SecgroupID     string
	Direction      string
	EtherType      string
	Protocol       string
	Description    string
	RemoteIPPrefix string
	PortRangeMax   int
	PortRangeMin   int
}

type ListSecgroupRules struct {
	Items []*SecgroupRule
}

func (l *ListSecgroupRules) Len() int {
	return len(l.Items)
}

func (l *ListSecgroupRules) Get(i int) *SecgroupRule {
	if i < 0 || i >= len(l.Items) {
		return nil
	}
	return l.Items[i]
}
