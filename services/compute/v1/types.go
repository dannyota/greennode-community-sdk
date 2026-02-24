package v1

type OSImage struct {
	ID            string       `json:"id"`
	ImageType     string       `json:"imageType"`
	ImageVersion  string       `json:"imageVersion"`
	Licence       *bool        `json:"licence"`
	FlavorZoneIDs []string     `json:"flavorZoneIds"`
	PackageLimit  PackageLimit `json:"packageLimit"`
	LicenseKey    *string      `json:"licenseKey"`
	DefaultTagIDs []string     `json:"defaultTagIds"`
	ZoneID        string       `json:"zoneId"`
	Description   string       `json:"description"`
}

type PackageLimit struct {
	Cpu      int64 `json:"cpu"`
	Memory   int64 `json:"memory"`
	DiskSize int64 `json:"diskSize"`
}

type ListOSImages struct {
	Items []*OSImage
}
