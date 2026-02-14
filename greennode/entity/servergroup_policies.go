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

func (s *ListServerGroupPolicies) Add(item *ServerGroupPolicy) {
	s.Items = append(s.Items, item)
}

func (s *ListServerGroupPolicies) At(idx int) *ServerGroupPolicy {
	if idx < 0 || idx >= s.Len() {
		return nil
	}

	return s.Items[idx]
}

func (s *ListServerGroupPolicies) Len() int {
	return len(s.Items)
}
