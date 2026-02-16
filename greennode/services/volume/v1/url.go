package v1

import "github.com/dannyota/greennode-community-sdk/v2/greennode/client"

func getVolumeTypeByIDURL(sc *client.ServiceClient, opts *GetVolumeTypeByIDRequest) string {
	return sc.ServiceURL(
		sc.GetProjectID(),
		"volume_types",
		opts.VolumeTypeID)
}

func getDefaultVolumeTypeURL(sc *client.ServiceClient) string {
	return sc.ServiceURL(
		sc.GetProjectID(),
		"volume_default_id")
}

func getVolumeTypeZonesURL(sc *client.ServiceClient, opts *GetVolumeTypeZonesRequest) string {
	query, err := opts.ToQuery()
	if err != nil {
		query = opts.GetDefaultQuery()
	}
	return sc.ServiceURL(
		sc.GetProjectID(),
		"volume_type_zones",
	) + query
}

func getVolumeTypesURL(sc *client.ServiceClient, opts *GetListVolumeTypeRequest) string {
	return sc.ServiceURL(
		sc.GetProjectID(),
		opts.GetVolumeTypeZoneID(),
		"volume_types",
	)
}
