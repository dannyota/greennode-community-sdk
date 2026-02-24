package v1

import "danny.vn/greennode/client"

func listOSImagesURL(sc *client.ServiceClient, opts *ListOSImagesRequest) string {
	return sc.ServiceURL(
		sc.ProjectID,
		"images",
		"os",
	) + opts.ToQuery()
}

func listGPUImagesURL(sc *client.ServiceClient) string {
	return sc.ServiceURL(
		sc.ProjectID,
		"images",
		"gpu",
	)
}
