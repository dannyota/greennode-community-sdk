package v2

import "github.com/dannyota/greennode-community-sdk/v2/greennode/client"

func createBlockVolumeUrl(psc client.ServiceClient) string {
	return psc.ServiceURL(
		psc.GetProjectId(),
		"volumes")
}

func deleteBlockVolumeByIdUrl(psc client.ServiceClient, popts IDeleteBlockVolumeByIdRequest) string {
	return psc.ServiceURL(
		psc.GetProjectId(),
		"volumes",
		popts.GetBlockVolumeId())
}

func listBlockVolumesUrl(psc client.ServiceClient, popts IListBlockVolumesRequest) string {
	query, err := popts.ToQuery()
	if err != nil {
		query = popts.GetDefaultQuery()
	}

	return psc.ServiceURL(
		psc.GetProjectId(),
		"volumes") + query
}

func getBlockVolumeByIdUrl(psc client.ServiceClient, popts IGetBlockVolumeByIdRequest) string {
	return psc.ServiceURL(
		psc.GetProjectId(),
		"volumes",
		popts.GetBlockVolumeId())
}

func resizeBlockVolumeByIdUrl(psc client.ServiceClient, popts IResizeBlockVolumeByIdRequest) string {
	return psc.ServiceURL(
		psc.GetProjectId(),
		"volumes",
		popts.GetBlockVolumeId(),
		"resize")
}

func listSnapshotsByBlockVolumeIdUrl(psc client.ServiceClient, popts IListSnapshotsByBlockVolumeIdRequest) string {
	query, err := popts.ToQuery()
	if err != nil {
		query = popts.GetDefaultQuery()
	}

	return psc.ServiceURL(
		psc.GetProjectId(),
		"volumes",
		popts.GetBlockVolumeId(),
		"snapshots",
	) + query
}

func createSnapshotByBlockVolumeIdUrl(psc client.ServiceClient, popts ICreateSnapshotByBlockVolumeIdRequest) string {
	return psc.ServiceURL(
		psc.GetProjectId(),
		"volumes",
		popts.GetBlockVolumeId(),
		"snapshots")
}

func deleteSnapshotByIdUrl(psc client.ServiceClient, popts IDeleteSnapshotByIdRequest) string {
	return psc.ServiceURL(
		psc.GetProjectId(),
		"volumes",
		popts.GetBlockVolumeId(),
		"snapshots",
		popts.GetSnapshotId(),
	)
}

func getUnderBlockVolumeIdUrl(psc client.ServiceClient, popts IGetUnderBlockVolumeIdRequest) string {
	return psc.ServiceURL(
		psc.GetProjectId(),
		"volumes",
		popts.GetBlockVolumeId(),
		"mapping",
	)
}

func migrateBlockVolumeByIdUrl(psc client.ServiceClient, popts IMigrateBlockVolumeByIdRequest) string {
	return psc.ServiceURL(
		psc.GetProjectId(),
		"volumes",
		popts.GetBlockVolumeId(),
		"change-device-type",
	)
}
