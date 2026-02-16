package entity

type Secgroup struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

type ListSecgroups struct {
	Items []*Secgroup
}
