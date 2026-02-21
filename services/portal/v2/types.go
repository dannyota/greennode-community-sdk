package v2

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
