package entity

type Endpoint struct {
	IPv4Address string
	EndpointURL string
	Status      string
	VpcID       string
	ID          string
	Name        string
}

func (e Endpoint) IsUsable() bool {
	return e.Status == "ACTIVE"
}

func (e Endpoint) IsError() bool {
	return e.Status == "ERROR"
}

type ListEndpoints struct {
	Items     []*Endpoint
	Page      int
	PageSize  int
	TotalPage int
	TotalItem int
}

func (l ListEndpoints) Len() int {
	return len(l.Items)
}

func (l ListEndpoints) At(idx int) *Endpoint {
	if idx < 0 || idx >= l.Len() {
		return nil
	}

	return l.Items[idx]
}
