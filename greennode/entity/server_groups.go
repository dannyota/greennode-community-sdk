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

func (l *ListServerGroups) Add(item *ServerGroup) {
	l.Items = append(l.Items, item)
}

func (l *ListServerGroups) FindServerGroupByServerGroupID(serverGroupID string) (*ServerGroup, bool) {
	for _, item := range l.Items {
		if item.UUID == serverGroupID {
			return item, true
		}
	}

	return nil, false
}
