package entity

type Portal struct {
	ProjectID string
	UserID    int
}

type Quota struct {
	Description string
	Name        string
	Type        string
	Limit       int
	Used        int
}

type ListQuotas struct {
	Items []*Quota
}

func (l *ListQuotas) FindQuotaByName(name string) *Quota {
	for _, q := range l.Items {
		if q.Name == name {
			return q
		}
	}

	return nil
}

type ListPortals struct {
	Items []*Portal
}

func NewListPortals() *ListPortals {
	return &ListPortals{
		Items: make([]*Portal, 0),
	}
}

func (l *ListPortals) At(index int) *Portal {
	if index < 0 || index >= len(l.Items) {
		return nil
	}

	return l.Items[index]
}
