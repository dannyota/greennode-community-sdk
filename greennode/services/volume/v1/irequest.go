package v1

type IGetVolumeTypeByIDRequest interface {
	GetVolumeTypeID() string
}

type IGetListVolumeTypeRequest interface {
	GetVolumeTypeZoneID() string
}

type IGetVolumeTypeZonesRequest interface {
	ToQuery() (string, error)
	GetDefaultQuery() string
}
