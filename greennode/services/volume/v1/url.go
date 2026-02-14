package v1

import "github.com/dannyota/greennode-community-sdk/v2/greennode/client"

func getVolumeTypeByIdUrl(sc client.ServiceClient, opts IGetVolumeTypeByIdRequest) string {
	return sc.ServiceURL(
		sc.GetProjectId(),
		"volume_types",
		opts.GetVolumeTypeId())
}

func getDefaultVolumeTypeUrl(sc client.ServiceClient) string {
	return sc.ServiceURL(
		sc.GetProjectId(),
		"volume_default_id")
}

func getVolumeTypeZonesUrl(sc client.ServiceClient, opts IGetVolumeTypeZonesRequest) string {
	query, err := opts.ToQuery()
	if err != nil {
		query = opts.GetDefaultQuery()
	}
	return sc.ServiceURL(
		sc.GetProjectId(),
		"volume_type_zones",
	) + query
}

func getVolumeTypesUrl(sc client.ServiceClient, opts IGetListVolumeTypeRequest) string {
	return sc.ServiceURL(
		sc.GetProjectId(),
		opts.GetVolumeTypeZoneId(),
		"volume_types",
	)
}
