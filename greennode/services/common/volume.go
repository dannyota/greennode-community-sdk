package common

type BlockVolumeCommon struct {
	BlockVolumeID string
}

func (s *BlockVolumeCommon) GetBlockVolumeID() string {
	return s.BlockVolumeID
}

type VolumeTypeCommon struct {
	VolumeTypeID string
}

func (s *VolumeTypeCommon) GetVolumeTypeID() string {
	return s.VolumeTypeID
}
