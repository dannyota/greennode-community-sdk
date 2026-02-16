package entity

type LoadBalancer struct {
	UUID               string  `json:"uuid"`
	Name               string  `json:"name"`
	DisplayStatus      string  `json:"displayStatus"`
	Address            string  `json:"address"`
	PrivateSubnetID    string  `json:"privateSubnetId"`    // subnet that its IP belongs to
	PrivateSubnetCidr  string  `json:"privateSubnetCidr"`  // cidr of subnet that its IP belongs to
	Type               string  `json:"type"`
	DisplayType        string  `json:"displayType"`
	LoadBalancerSchema string  `json:"loadBalancerSchema"`
	PackageID          string  `json:"packageId"`
	Description        string  `json:"description"`
	Location           string  `json:"location"`
	CreatedAt          string  `json:"createdAt"`
	UpdatedAt          string  `json:"updatedAt"`
	ProgressStatus     string  `json:"progressStatus"`
	Status             string  `json:"status"`
	BackendSubnetID    string  `json:"backendSubnetId"`    // subnet that load balancer connects to backend servers
	Internal           bool    `json:"internal"`
	AutoScalable       bool    `json:"autoScalable"`
	ZoneID             string  `json:"zoneId"`
	MinSize            int     `json:"minSize"`            // minimum number of nodes for HA configuration
	MaxSize            int     `json:"maxSize"`            // maximum number of nodes for HA configuration
	TotalNodes         int     `json:"totalNodes"`         // total number of nodes currently running
	Nodes              []*Node `json:"nodes"`              // list of nodes in the load balancer
}

type Node struct {
	Status   string `json:"status"`
	ZoneID   string `json:"zoneId"`
	ZoneName string `json:"zoneName"`
	SubnetID string `json:"subnetId"`
}

type ListLoadBalancers struct {
	Items     []*LoadBalancer
	Page      int
	PageSize  int
	TotalPage int
	TotalItem int
}

func (l *ListLoadBalancers) Len() int {
	return len(l.Items)
}

func (l *ListLoadBalancers) Empty() bool {
	return l.Len() < 1
}

func (l *ListLoadBalancers) Add(item *LoadBalancer) {
	l.Items = append(l.Items, item)
}

func (l *ListLoadBalancers) At(idx int) *LoadBalancer {
	if idx < 0 || idx >= l.Len() {
		return nil
	}

	return l.Items[idx]
}

type ListLoadBalancerPackages struct {
	Items []*LoadBalancerPackage
}

type LoadBalancerPackage struct {
	UUID             string `json:"uuid"`
	Name             string `json:"name"`
	Type             string `json:"type"`
	ConnectionNumber int    `json:"connectionNumber"`
	DataTransfer     int    `json:"dataTransfer"`
	Mode             string `json:"mode"`
	LbType           string `json:"lbType"`
	DisplayLbType    string `json:"displayLbType"`
}
