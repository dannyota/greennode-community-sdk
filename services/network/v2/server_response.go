package v2

import computev2 "danny.vn/greennode/services/compute/v2"

type ListAllServersBySecgroupIDResponse struct {
	Data []struct {
		Name   string `json:"name"`
		UUID   string `json:"uuid"`
		Status string `json:"status"`
	} `json:"data"`
}

func (r *ListAllServersBySecgroupIDResponse) ToEntityListServers() *computev2.ListServers {
	servers := make([]*computev2.Server, 0, len(r.Data))
	for _, server := range r.Data {
		servers = append(servers, &computev2.Server{
			Name:   server.Name,
			Uuid:   server.UUID,
			Status: server.Status,
		})
	}
	return &computev2.ListServers{
		Items: servers,
	}
}
