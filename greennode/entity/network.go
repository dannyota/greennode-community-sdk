package entity

type Network struct {
	Status     string   `json:"status"`
	ElasticIps []string `json:"elasticIps"`
	Name       string   `json:"displayName"`
	ID         string   `json:"id"`
	CreatedAt  string   `json:"createdAt"`
	Cidr       string   `json:"cidr"`
}
