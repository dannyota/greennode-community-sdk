package entity

type ServerGroupPolicy struct {
	Name         string
	UUID         string
	Status       string
	Descriptions map[string]string
}

type ListServerGroupPolicies struct {
	Items []*ServerGroupPolicy
}

func (l *ListServerGroupPolicies) Add(item *ServerGroupPolicy) {
	l.Items = append(l.Items, item)
}

func (l *ListServerGroupPolicies) At(idx int) *ServerGroupPolicy {
	if idx < 0 || idx >= l.Len() {
		return nil
	}

	return l.Items[idx]
}

func (l *ListServerGroupPolicies) Len() int {
	return len(l.Items)
}
