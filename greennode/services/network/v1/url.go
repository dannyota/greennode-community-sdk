package v1

import "github.com/dannyota/greennode-community-sdk/v2/greennode/client"

func getEndpointByIdUrl(sc client.ServiceClient, opts IGetEndpointByIdRequest) string {
	return sc.ServiceURL(
		sc.GetZoneId(),
		sc.GetProjectId(),
		"endpoints",
		opts.GetEndpointId())
}

func createEndpointUrl(sc client.ServiceClient) string {
	return sc.ServiceURL(
		sc.GetZoneId(),
		sc.GetProjectId(),
		"endpoints")
}

func deleteEndpointByIdUrl(sc client.ServiceClient, opts IDeleteEndpointByIdRequest) string {
	return sc.ServiceURL(
		sc.GetZoneId(),
		sc.GetProjectId(),
		"endpoints",
		opts.GetEndpointId())
}

func listEndpointsUrl(sc client.ServiceClient, opts IListEndpointsRequest) string {
	query, err := opts.ToListQuery()
	if err != nil {
		query = opts.GetDefaultQuery()
	}

	return sc.ServiceURL(sc.GetZoneId(), sc.GetProjectId(), "endpoints?") + query
}

func listTagsByEndpointIdUrl(sc client.ServiceClient, opts IListTagsByEndpointIdRequest) string {
	query, err := opts.ToListQuery()
	if err != nil {
		query = opts.GetDefaultQuery()
	}

	return sc.ServiceURL(
		opts.GetProjectId(),
		"tags") + query
}

func createTagsWithEndpointIdUrl(sc client.ServiceClient, opts ICreateTagsWithEndpointIdRequest) string {
	return sc.ServiceURL(
		opts.GetProjectId(),
		"tags")
}

func deleteTagOfEndpointUrl(sc client.ServiceClient, opts IDeleteTagOfEndpointRequest) string {
	return sc.ServiceURL(
		opts.GetProjectId(),
		"tags",
		opts.GetTagId())
}

func updateTagValueOfEndpointUrl(sc client.ServiceClient, opts IUpdateTagValueOfEndpointRequest) string {
	return sc.ServiceURL(
		opts.GetProjectId(),
		"tags",
		opts.GetTagId())
}
