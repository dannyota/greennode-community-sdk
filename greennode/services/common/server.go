package common

type ServerCommon struct {
	ServerID string
}

func (s *ServerCommon) GetServerID() string {
	return s.ServerID
}

type ServerGroupCommon struct {
	ServerGroupID string
}

func (s *ServerGroupCommon) GetServerGroupID() string {
	return s.ServerGroupID
}
