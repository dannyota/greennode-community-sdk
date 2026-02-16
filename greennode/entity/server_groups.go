package entity

type ServerGroup struct {
	UUID        string              `json:"uuid"`
	Name        string              `json:"name"`
	Description string              `json:"description"`
	PolicyID    string              `json:"policyId"`
	PolicyName  string              `json:"policyName"`
	Servers     []ServerGroupMember `json:"servers"`
}

type ServerGroupMember struct {
	Name string `json:"name"`
	UUID string `json:"uuid"`
}

type ListServerGroups struct {
	Items     []*ServerGroup `json:"listData"`
	Page      int            `json:"page"`
	PageSize  int            `json:"pageSize"`
	TotalPage int            `json:"totalPage"`
	TotalItem int            `json:"totalItem"`
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
