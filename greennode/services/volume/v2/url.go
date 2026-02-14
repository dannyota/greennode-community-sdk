package v2

import "github.com/dannyota/greennode-community-sdk/v2/greennode/client"

func createBlockVolumeUrl(psc client.IServiceClient) string {
	return psc.ServiceURL(
		psc.GetProjectId(),
		"volumes")
}

func deleteBlockVolumeByIdUrl(psc client.IServiceClient, popts IDeleteBlockVolumeByIdRequest) string {
	return psc.ServiceURL(
		psc.GetProjectId(),
		"volumes",
		popts.GetBlockVolumeId())
}

func listBlockVolumesUrl(psc client.IServiceClient, popts IListBlockVolumesRequest) string {
	query, err := popts.ToQuery()
	if err != nil {
		query = popts.GetDefaultQuery()
	}

	return psc.ServiceURL(
		psc.GetProjectId(),
		"volumes") + query
}

func getBlockVolumeByIdUrl(psc client.IServiceClient, popts IGetBlockVolumeByIdRequest) string {
	return psc.ServiceURL(
		psc.GetProjectId(),
		"volumes",
		popts.GetBlockVolumeId())
}

func resizeBlockVolumeByIdUrl(psc client.IServiceClient, popts IResizeBlockVolumeByIdRequest) string {
	return psc.ServiceURL(
		psc.GetProjectId(),
		"volumes",
		popts.GetBlockVolumeId(),
		"resize")
}

func listSnapshotsByBlockVolumeIdUrl(psc client.IServiceClient, popts IListSnapshotsByBlockVolumeIdRequest) string {
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

func createSnapshotByBlockVolumeIdUrl(psc client.IServiceClient, popts ICreateSnapshotByBlockVolumeIdRequest) string {
	return psc.ServiceURL(
		psc.GetProjectId(),
		"volumes",
		popts.GetBlockVolumeId(),
		"snapshots")
}

func deleteSnapshotByIdUrl(psc client.IServiceClient, popts IDeleteSnapshotByIdRequest) string {
	return psc.ServiceURL(
		psc.GetProjectId(),
		"volumes",
		popts.GetBlockVolumeId(),
		"snapshots",
		popts.GetSnapshotId(),
	)
}

func getUnderBlockVolumeIdUrl(psc client.IServiceClient, popts IGetUnderBlockVolumeIdRequest) string {
	return psc.ServiceURL(
		psc.GetProjectId(),
		"volumes",
		popts.GetBlockVolumeId(),
		"mapping",
	)
}

func migrateBlockVolumeByIdUrl(psc client.IServiceClient, popts IMigrateBlockVolumeByIdRequest) string {
	return psc.ServiceURL(
		psc.GetProjectId(),
		"volumes",
		popts.GetBlockVolumeId(),
		"change-device-type",
	)
}
