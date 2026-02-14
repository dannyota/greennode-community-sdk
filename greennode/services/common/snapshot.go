package common

type SnapshotCommon struct {
	SnapshotID string
}

func (s *SnapshotCommon) GetSnapshotID() string {
	return s.SnapshotID
}
