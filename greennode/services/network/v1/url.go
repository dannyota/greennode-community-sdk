package v1

import "github.com/dannyota/greennode-community-sdk/v2/greennode/client"

func getEndpointByIdUrl(psc client.IServiceClient, popts IGetEndpointByIdRequest) string {
	return psc.ServiceURL(
		psc.GetZoneId(),
		psc.GetProjectId(),
		"endpoints",
		popts.GetEndpointId())
}

func createEndpointUrl(psc client.IServiceClient) string {
	return psc.ServiceURL(
		psc.GetZoneId(),
		psc.GetProjectId(),
		"endpoints")
}

func deleteEndpointByIdUrl(psc client.IServiceClient, popts IDeleteEndpointByIdRequest) string {
	return psc.ServiceURL(
		psc.GetZoneId(),
		psc.GetProjectId(),
		"endpoints",
		popts.GetEndpointId())
}

func listEndpointsUrl(psc client.IServiceClient, popts IListEndpointsRequest) string {
	query, err := popts.ToListQuery()
	if err != nil {
		query = popts.GetDefaultQuery()
	}

	return psc.ServiceURL(psc.GetZoneId(), psc.GetProjectId(), "endpoints?") + query
}

func listTagsByEndpointIdUrl(psc client.IServiceClient, popts IListTagsByEndpointIdRequest) string {
	query, err := popts.ToListQuery()
	if err != nil {
		query = popts.GetDefaultQuery()
	}

	return psc.ServiceURL(
		popts.GetProjectId(),
		"tags") + query
}

func createTagsWithEndpointIdUrl(psc client.IServiceClient, popts ICreateTagsWithEndpointIdRequest) string {
	return psc.ServiceURL(
		popts.GetProjectId(),
		"tags")
}

func deleteTagOfEndpointUrl(psc client.IServiceClient, popts IDeleteTagOfEndpointRequest) string {
	return psc.ServiceURL(
		popts.GetProjectId(),
		"tags",
		popts.GetTagId())
}

func updateTagValueOfEndpointUrl(psc client.IServiceClient, popts IUpdateTagValueOfEndpointRequest) string {
	return psc.ServiceURL(
		popts.GetProjectId(),
		"tags",
		popts.GetTagId())
}
