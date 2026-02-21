package v2

type CreateBlockVolumeResponse struct {
	Data BlockVolume `json:"data"`
}

type ListBlockVolumesResponse struct {
	Page      int64         `json:"page"`
	PageSize  int64         `json:"pageSize"`
	TotalPage int64         `json:"totalPage"`
	TotalItem int64         `json:"totalItem"`
	ListData  []BlockVolume `json:"listData"`
}

type GetBlockVolumeByIDResponse struct {
	Data BlockVolume `json:"data"`
}

type ResizeBlockVolumeByIDResponse struct {
	Data BlockVolume `json:"data"`
}

type GetUnderBlockVolumeIDResponse struct {
	Uuid string `json:"uuid"`
}

type Zone struct {
	Uuid string `json:"uuid"`
}

func (r *GetBlockVolumeByIDResponse) ToEntityVolume() *Volume {
	return r.Data.toEntityVolume()
}

type (
	BlockVolume struct {
		UUID               string   `json:"uuid"`
		ProjectID          string   `json:"projectId"`
		Name               string   `json:"name"`
		Size               uint64   `json:"size"`
		Status             string   `json:"status"`
		VolumeTypeID       string   `json:"volumeTypeId"`
		VolumeTypeZoneName string   `json:"volumeTypeZoneName"`
		IOPS               string   `json:"iops"`
		ServerID           string   `json:"serverId,omitempty"`
		CreatedAt          string   `json:"createdAt"`
		UpdatedAt          *string  `json:"updatedAt"`
		Bootable           bool     `json:"bootable"`
		EncryptionType     *string  `json:"encryptionType"`
		BootIndex          int      `json:"bootIndex"`
		MultiAttach        bool     `json:"multiAttach"`
		ServerIDList       []string `json:"serverIdList"`
		Location           *string  `json:"location"`
		Product            string   `json:"product"`
		PersistentVolume   bool     `json:"persistentVolume"`
		MigrateState       string   `json:"migrateState,omitempty"`
		Zone               Zone     `json:"zone"`
	}
)

func (r *CreateBlockVolumeResponse) ToEntityVolume() *Volume {
	return r.Data.toEntityVolume()
}

func (r *ListBlockVolumesResponse) ToEntityListVolumes() *ListVolumes {
	lstVolumes := new(ListVolumes)
	for _, vol := range r.ListData {
		lstVolumes.Items = append(lstVolumes.Items, vol.toEntityVolume())
	}

	return lstVolumes
}

func (b *BlockVolume) toEntityVolume() *Volume {
	return &Volume{
		ID:              b.UUID,
		Name:            b.Name,
		Size:            b.Size,
		Status:          b.Status,
		CreatedAt:       b.CreatedAt,
		UpdatedAt:       b.UpdatedAt,
		VmID:            b.ServerID,
		AttachedMachine: b.ServerIDList,
		VolumeTypeID:    b.VolumeTypeID,
		MigrateState:    b.MigrateState,
		MultiAttach:     b.MultiAttach,
		ZoneID:          b.Zone.Uuid,
	}
}

func (r *ResizeBlockVolumeByIDResponse) ToEntityVolume() *Volume {
	return r.Data.toEntityVolume()
}

func (r *GetUnderBlockVolumeIDResponse) ToEntityVolume() *Volume {
	return &Volume{
		UnderID: r.Uuid,
	}
}
