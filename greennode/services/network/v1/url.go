package v1

import "github.com/dannyota/greennode-community-sdk/v2/greennode/client"

func getEndpointByIDURL(sc client.ServiceClient, opts IGetEndpointByIDRequest) string {
	return sc.ServiceURL(
		sc.GetZoneID(),
		sc.GetProjectID(),
		"endpoints",
		opts.GetEndpointID())
}

func createEndpointURL(sc client.ServiceClient) string {
	return sc.ServiceURL(
		sc.GetZoneID(),
		sc.GetProjectID(),
		"endpoints")
}

func deleteEndpointByIDURL(sc client.ServiceClient, opts IDeleteEndpointByIDRequest) string {
	return sc.ServiceURL(
		sc.GetZoneID(),
		sc.GetProjectID(),
		"endpoints",
		opts.GetEndpointID())
}

func listEndpointsURL(sc client.ServiceClient, opts IListEndpointsRequest) string {
	query, err := opts.ToListQuery()
	if err != nil {
		query = opts.GetDefaultQuery()
	}

	return sc.ServiceURL(sc.GetZoneID(), sc.GetProjectID(), "endpoints?") + query
}

func listTagsByEndpointIDURL(sc client.ServiceClient, opts IListTagsByEndpointIDRequest) string {
	query, err := opts.ToListQuery()
	if err != nil {
		query = opts.GetDefaultQuery()
	}

	return sc.ServiceURL(
		opts.GetProjectID(),
		"tags") + query
}

func createTagsWithEndpointIDURL(sc client.ServiceClient, opts ICreateTagsWithEndpointIDRequest) string {
	return sc.ServiceURL(
		opts.GetProjectID(),
		"tags")
}

func deleteTagOfEndpointURL(sc client.ServiceClient, opts IDeleteTagOfEndpointRequest) string {
	return sc.ServiceURL(
		opts.GetProjectID(),
		"tags",
		opts.GetTagID())
}

func updateTagValueOfEndpointURL(sc client.ServiceClient, opts IUpdateTagValueOfEndpointRequest) string {
	return sc.ServiceURL(
		opts.GetProjectID(),
		"tags",
		opts.GetTagID())
}
