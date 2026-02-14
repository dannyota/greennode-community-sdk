package entity

type Secgroup struct {
	ID          string
	Name        string
	Description string
	Status      string
}

type ListSecgroups struct {
	Items []*Secgroup
}
