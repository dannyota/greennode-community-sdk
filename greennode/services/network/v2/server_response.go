package v2

import "github.com/dannyota/greennode-community-sdk/v2/greennode/entity"

type ListAllServersBySecgroupIdResponse struct {
	Data []struct {
		Name   string `json:"name"`
		UUID   string `json:"uuid"`
		Status string `json:"status"`
	} `json:"data"`
}

func (s *ListAllServersBySecgroupIdResponse) ToEntityListServers() *entity.ListServers {
	servers := make([]*entity.Server, 0, len(s.Data))
	for _, server := range s.Data {
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
