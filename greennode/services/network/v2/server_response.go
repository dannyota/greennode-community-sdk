package v2

import "github.com/dannyota/greennode-community-sdk/v2/greennode/entity"

type ListAllServersBySecgroupIDResponse struct {
	Data []struct {
		Name   string `json:"name"`
		UUID   string `json:"uuid"`
		Status string `json:"status"`
	} `json:"data"`
}

func (r *ListAllServersBySecgroupIDResponse) ToEntityListServers() *entity.ListServers {
	servers := make([]*entity.Server, 0, len(r.Data))
	for _, server := range r.Data {
		servers = append(servers, &entity.Server{
			Name:   server.Name,
			Uuid:   server.UUID,
			Status: server.Status,
		})
	}
	return &entity.ListServers{
		Items: servers,
	}
}
