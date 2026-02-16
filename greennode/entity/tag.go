package entity

type Tag struct {
	Key        string `json:"key"`
	Value      string `json:"value"`
	SystemTag  bool   `json:"systemTag"`
	ResourceID string `json:"resourceId"`
	TagID      string `json:"tagId"`
}

type ListTags struct {
	Items []*Tag
}

func (l *ListTags) Len() int {
	return len(l.Items)
}

func (l *ListTags) Empty() bool {
	return l.Len() < 1
}

func (l *ListTags) Add(tags ...*Tag) {
	l.Items = append(l.Items, tags...)
}
