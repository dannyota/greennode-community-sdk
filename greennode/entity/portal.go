package entity

type Portal struct {
	ProjectID string `json:"projectId"`
	UserID    int    `json:"userId"`
}

type Quota struct {
	Description string `json:"description"`
	Name        string `json:"quotaName"`
	Type        string `json:"type"`
	Limit       int    `json:"limit"`
	Used        int    `json:"used"`
}

type ListQuotas struct {
	Items []*Quota
}

func (l ListQuotas) FindQuotaByName(name string) *Quota {
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

func (l ListPortals) At(index int) *Portal {
	if index < 0 || index >= len(l.Items) {
		return nil
	}

	return l.Items[index]
}
