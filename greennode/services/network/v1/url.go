package v1

import "github.com/dannyota/greennode-community-sdk/v2/greennode/client"

func getEndpointByIdUrl(psc client.ServiceClient, popts IGetEndpointByIdRequest) string {
	return psc.ServiceURL(
		psc.GetZoneId(),
		psc.GetProjectId(),
		"endpoints",
		popts.GetEndpointId())
}

func createEndpointUrl(psc client.ServiceClient) string {
	return psc.ServiceURL(
		psc.GetZoneId(),
		psc.GetProjectId(),
		"endpoints")
}

func deleteEndpointByIdUrl(psc client.ServiceClient, popts IDeleteEndpointByIdRequest) string {
	return psc.ServiceURL(
		psc.GetZoneId(),
		psc.GetProjectId(),
		"endpoints",
		popts.GetEndpointId())
}

func listEndpointsUrl(psc client.ServiceClient, popts IListEndpointsRequest) string {
	query, err := popts.ToListQuery()
	if err != nil {
		query = popts.GetDefaultQuery()
	}

	return psc.ServiceURL(psc.GetZoneId(), psc.GetProjectId(), "endpoints?") + query
}

func listTagsByEndpointIdUrl(psc client.ServiceClient, popts IListTagsByEndpointIdRequest) string {
	query, err := popts.ToListQuery()
	if err != nil {
		query = popts.GetDefaultQuery()
	}

	return psc.ServiceURL(
		popts.GetProjectId(),
		"tags") + query
}

func createTagsWithEndpointIdUrl(psc client.ServiceClient, popts ICreateTagsWithEndpointIdRequest) string {
	return psc.ServiceURL(
		popts.GetProjectId(),
		"tags")
}

func deleteTagOfEndpointUrl(psc client.ServiceClient, popts IDeleteTagOfEndpointRequest) string {
	return psc.ServiceURL(
		popts.GetProjectId(),
		"tags",
		popts.GetTagId())
}

func updateTagValueOfEndpointUrl(psc client.ServiceClient, popts IUpdateTagValueOfEndpointRequest) string {
	return psc.ServiceURL(
		popts.GetProjectId(),
		"tags",
		popts.GetTagId())
}
