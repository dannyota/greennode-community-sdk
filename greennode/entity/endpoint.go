package entity

type Endpoint struct {
	IPv4Address string
	EndpointURL string
	Status      string
	VpcID       string
	ID          string
	Name        string
}

func (s *Endpoint) IsUsable() bool {
	return s.Status == "ACTIVE"
}

func (s *Endpoint) GetID() string {
	return s.ID
}

func (s *Endpoint) GetName() string {
	return s.Name
}

func (s *Endpoint) GetIPv4Address() string {
	return s.IPv4Address
}

func (s *Endpoint) GetEndpointURL() string {
	return s.EndpointURL
}

func (s *Endpoint) GetVpcID() string {
	return s.VpcID
}

func (s *Endpoint) GetStatus() string {
	return s.Status
}

func (s *Endpoint) IsError() bool {
	return s.Status == "ERROR"
}

type ListEndpoints struct {
	Items     []*Endpoint
	Page      int
	PageSize  int
	TotalPage int
	TotalItem int
}

func (s *ListEndpoints) Len() int {
	return len(s.Items)
}

func (s *ListEndpoints) At(idx int) *Endpoint {
	if idx < 0 || idx >= s.Len() {
		return nil
	}

	return s.Items[idx]
}
