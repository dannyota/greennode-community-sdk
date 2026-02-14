package entity

type ServerGroup struct {
	UUID        string
	Name        string
	Description string
	PolicyID    string
	PolicyName  string
	Servers     []ServerGroupMember
}

type ServerGroupMember struct {
	Name string
	UUID string
}

type ListServerGroups struct {
	Items     []*ServerGroup
	Page      int
	PageSize  int
	TotalPage int
	TotalItem int
}

func (s *ListServerGroups) Add(item *ServerGroup) {
	s.Items = append(s.Items, item)
}

func (s *ListServerGroups) FindServerGroupByServerGroupID(serverGroupID string) (*ServerGroup, bool) {
	for _, item := range s.Items {
		if item.UUID == serverGroupID {
			return item, true
		}
	}

	return nil, false
}
