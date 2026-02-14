package common

type BlockVolumeCommon struct {
	BlockVolumeID string
}

func (b *BlockVolumeCommon) GetBlockVolumeID() string {
	return b.BlockVolumeID
}

type VolumeTypeCommon struct {
	VolumeTypeID string
}

func (v *VolumeTypeCommon) GetVolumeTypeID() string {
	return v.VolumeTypeID
}
