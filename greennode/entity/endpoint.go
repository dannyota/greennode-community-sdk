package entity

type Endpoint struct {
	IPv4Address string
	EndpointURL string
	Status      string
	VpcID       string
	ID          string
	Name        string
}

func (e *Endpoint) IsUsable() bool {
	return e.Status == "ACTIVE"
}

func (e *Endpoint) GetID() string {
	return e.ID
}

func (e *Endpoint) GetName() string {
	return e.Name
}

func (e *Endpoint) GetIPv4Address() string {
	return e.IPv4Address
}

func (e *Endpoint) GetEndpointURL() string {
	return e.EndpointURL
}

func (e *Endpoint) GetVpcID() string {
	return e.VpcID
}

func (e *Endpoint) GetStatus() string {
	return e.Status
}

func (e *Endpoint) IsError() bool {
	return e.Status == "ERROR"
}

type ListEndpoints struct {
	Items     []*Endpoint
	Page      int
	PageSize  int
	TotalPage int
	TotalItem int
}

func (l *ListEndpoints) Len() int {
	return len(l.Items)
}

func (l *ListEndpoints) At(idx int) *Endpoint {
	if idx < 0 || idx >= l.Len() {
		return nil
	}

	return l.Items[idx]
}
