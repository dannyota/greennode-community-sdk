package v1

type SystemTag struct {
	Key       string `json:"key"`
	Value     string `json:"value"`
	CreatedAt string `json:"createdAt"`
	SystemTag bool   `json:"systemTag"`
}
