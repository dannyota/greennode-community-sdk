package v1

import "github.com/dannyota/greennode-community-sdk/v2/greennode/client"

func getVolumeTypeByIdUrl(psc client.IServiceClient, popts IGetVolumeTypeByIdRequest) string {
	return psc.ServiceURL(
		psc.GetProjectId(),
		"volume_types",
		popts.GetVolumeTypeId())
}

func getDefaultVolumeTypeUrl(psc client.IServiceClient) string {
	return psc.ServiceURL(
		psc.GetProjectId(),
		"volume_default_id")
}

func getVolumeTypeZonesUrl(psc client.IServiceClient, popts IGetVolumeTypeZonesRequest) string {
	query, err := popts.ToQuery()
	if err != nil {
		query = popts.GetDefaultQuery()
	}
	return psc.ServiceURL(
		psc.GetProjectId(),
		"volume_type_zones",
	) + query
}

func getVolumeTypesUrl(psc client.IServiceClient, popts IGetListVolumeTypeRequest) string {
	return psc.ServiceURL(
		psc.GetProjectId(),
		popts.GetVolumeTypeZoneId(),
		"volume_types",
	)
}
