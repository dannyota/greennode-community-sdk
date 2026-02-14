package v2

import "github.com/dannyota/greennode-community-sdk/v2/greennode/client"

func createBlockVolumeUrl(sc client.ServiceClient) string {
	return sc.ServiceURL(
		sc.GetProjectId(),
		"volumes")
}

func deleteBlockVolumeByIdUrl(sc client.ServiceClient, opts IDeleteBlockVolumeByIdRequest) string {
	return sc.ServiceURL(
		sc.GetProjectId(),
		"volumes",
		opts.GetBlockVolumeId())
}

func listBlockVolumesUrl(sc client.ServiceClient, opts IListBlockVolumesRequest) string {
	query, err := opts.ToQuery()
	if err != nil {
		query = opts.GetDefaultQuery()
	}

	return sc.ServiceURL(
		sc.GetProjectId(),
		"volumes") + query
}

func getBlockVolumeByIdUrl(sc client.ServiceClient, opts IGetBlockVolumeByIdRequest) string {
	return sc.ServiceURL(
		sc.GetProjectId(),
		"volumes",
		opts.GetBlockVolumeId())
}

func resizeBlockVolumeByIdUrl(sc client.ServiceClient, opts IResizeBlockVolumeByIdRequest) string {
	return sc.ServiceURL(
		sc.GetProjectId(),
		"volumes",
		opts.GetBlockVolumeId(),
		"resize")
}

func listSnapshotsByBlockVolumeIdUrl(sc client.ServiceClient, opts IListSnapshotsByBlockVolumeIdRequest) string {
	query, err := opts.ToQuery()
	if err != nil {
		query = opts.GetDefaultQuery()
	}

	return sc.ServiceURL(
		sc.GetProjectId(),
		"volumes",
		opts.GetBlockVolumeId(),
		"snapshots",
	) + query
}

func createSnapshotByBlockVolumeIdUrl(sc client.ServiceClient, opts ICreateSnapshotByBlockVolumeIdRequest) string {
	return sc.ServiceURL(
		sc.GetProjectId(),
		"volumes",
		opts.GetBlockVolumeId(),
		"snapshots")
}

func deleteSnapshotByIdUrl(sc client.ServiceClient, opts IDeleteSnapshotByIdRequest) string {
	return sc.ServiceURL(
		sc.GetProjectId(),
		"volumes",
		opts.GetBlockVolumeId(),
		"snapshots",
		opts.GetSnapshotId(),
	)
}

func getUnderBlockVolumeIdUrl(sc client.ServiceClient, opts IGetUnderBlockVolumeIdRequest) string {
	return sc.ServiceURL(
		sc.GetProjectId(),
		"volumes",
		opts.GetBlockVolumeId(),
		"mapping",
	)
}

func migrateBlockVolumeByIdUrl(sc client.ServiceClient, opts IMigrateBlockVolumeByIdRequest) string {
	return sc.ServiceURL(
		sc.GetProjectId(),
		"volumes",
		opts.GetBlockVolumeId(),
		"change-device-type",
	)
}
