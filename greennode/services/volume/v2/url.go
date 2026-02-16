package v2

import "github.com/dannyota/greennode-community-sdk/v2/greennode/client"

func createBlockVolumeURL(sc *client.ServiceClient) string {
	return sc.ServiceURL(
		sc.GetProjectID(),
		"volumes")
}

func deleteBlockVolumeByIDURL(sc *client.ServiceClient, opts *DeleteBlockVolumeByIDRequest) string {
	return sc.ServiceURL(
		sc.GetProjectID(),
		"volumes",
		opts.BlockVolumeID)
}

func listBlockVolumesURL(sc *client.ServiceClient, opts *ListBlockVolumesRequest) string {
	query, err := opts.ToQuery()
	if err != nil {
		query = opts.GetDefaultQuery()
	}

	return sc.ServiceURL(
		sc.GetProjectID(),
		"volumes") + query
}

func getBlockVolumeByIDURL(sc *client.ServiceClient, opts *GetBlockVolumeByIDRequest) string {
	return sc.ServiceURL(
		sc.GetProjectID(),
		"volumes",
		opts.BlockVolumeID)
}

func resizeBlockVolumeByIDURL(sc *client.ServiceClient, opts *ResizeBlockVolumeByIDRequest) string {
	return sc.ServiceURL(
		sc.GetProjectID(),
		"volumes",
		opts.BlockVolumeID,
		"resize")
}

func listSnapshotsByBlockVolumeIDURL(sc *client.ServiceClient, opts *ListSnapshotsByBlockVolumeIDRequest) string {
	query, err := opts.ToQuery()
	if err != nil {
		query = opts.GetDefaultQuery()
	}

	return sc.ServiceURL(
		sc.GetProjectID(),
		"volumes",
		opts.BlockVolumeID,
		"snapshots",
	) + query
}

func createSnapshotByBlockVolumeIDURL(sc *client.ServiceClient, opts *CreateSnapshotByBlockVolumeIDRequest) string {
	return sc.ServiceURL(
		sc.GetProjectID(),
		"volumes",
		opts.BlockVolumeID,
		"snapshots")
}

func deleteSnapshotByIDURL(sc *client.ServiceClient, opts *DeleteSnapshotByIDRequest) string {
	return sc.ServiceURL(
		sc.GetProjectID(),
		"volumes",
		opts.BlockVolumeID,
		"snapshots",
		opts.SnapshotID,
	)
}

func getUnderBlockVolumeIDURL(sc *client.ServiceClient, opts *GetUnderBlockVolumeIDRequest) string {
	return sc.ServiceURL(
		sc.GetProjectID(),
		"volumes",
		opts.BlockVolumeID,
		"mapping",
	)
}

func migrateBlockVolumeByIDURL(sc *client.ServiceClient, opts *MigrateBlockVolumeByIDRequest) string {
	return sc.ServiceURL(
		sc.GetProjectID(),
		"volumes",
		opts.BlockVolumeID,
		"change-device-type",
	)
}
