package entity

type Tag struct {
	Key        string
	Value      string
	SystemTag  bool
	ResourceID string
	TagID      string
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
