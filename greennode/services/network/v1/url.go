package v1

import "github.com/dannyota/greennode-community-sdk/v2/greennode/client"

func getEndpointByIDURL(sc *client.ServiceClient, opts *GetEndpointByIDRequest) string {
	return sc.ServiceURL(
		sc.ZoneID(),
		sc.ProjectID(),
		"endpoints",
		opts.EndpointID)
}

func createEndpointURL(sc *client.ServiceClient) string {
	return sc.ServiceURL(
		sc.ZoneID(),
		sc.ProjectID(),
		"endpoints")
}

func deleteEndpointByIDURL(sc *client.ServiceClient, opts *DeleteEndpointByIDRequest) string {
	return sc.ServiceURL(
		sc.ZoneID(),
		sc.ProjectID(),
		"endpoints",
		opts.EndpointID)
}

func listEndpointsURL(sc *client.ServiceClient, opts *ListEndpointsRequest) string {
	query, err := opts.ToListQuery()
	if err != nil {
		query = opts.getDefaultQuery()
	}

	return sc.ServiceURL(sc.ZoneID(), sc.ProjectID(), "endpoints?") + query
}

func listTagsByEndpointIDURL(sc *client.ServiceClient, opts *ListTagsByEndpointIDRequest) string {
	query, err := opts.ToListQuery()
	if err != nil {
		query = opts.getDefaultQuery()
	}

	return sc.ServiceURL(
		opts.GetProjectID(),
		"tags") + query
}

func createTagsWithEndpointIDURL(sc *client.ServiceClient, opts *CreateTagsWithEndpointIDRequest) string {
	return sc.ServiceURL(
		opts.GetProjectID(),
		"tags")
}

func deleteTagOfEndpointURL(sc *client.ServiceClient, opts *DeleteTagOfEndpointRequest) string {
	return sc.ServiceURL(
		opts.GetProjectID(),
		"tags",
		opts.GetTagID())
}

func updateTagValueOfEndpointURL(sc *client.ServiceClient, opts *UpdateTagValueOfEndpointRequest) string {
	return sc.ServiceURL(
		opts.GetProjectID(),
		"tags",
		opts.GetTagID())
}
