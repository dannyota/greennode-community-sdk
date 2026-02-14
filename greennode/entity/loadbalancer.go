package entity

type LoadBalancer struct {
	UUID               string
	Name               string
	DisplayStatus      string
	Address            string
	PrivateSubnetID    string // subnet that its IP belongs to
	PrivateSubnetCidr  string // cidr of subnet that its IP belongs to
	Type               string
	DisplayType        string
	LoadBalancerSchema string
	PackageID          string
	Description        string
	Location           string
	CreatedAt          string
	UpdatedAt          string
	ProgressStatus     string
	Status             string
	BackendSubnetID    string // subnet that load balancer connects to backend servers
	Internal           bool
	AutoScalable       bool
	ZoneID             string
	MinSize            int     // minimum number of nodes for HA configuration
	MaxSize            int     // maximum number of nodes for HA configuration
	TotalNodes         int     // total number of nodes currently running
	Nodes              []*Node // list of nodes in the load balancer
}

type Node struct {
	Status   string
	ZoneID   string
	ZoneName string
	SubnetID string
}

type ListLoadBalancers struct {
	Items     []*LoadBalancer
	Page      int
	PageSize  int
	TotalPage int
	TotalItem int
}

func (lb LoadBalancer) GetID() string {
	return lb.UUID
}

func (lb LoadBalancer) GetName() string {
	return lb.Name
}

func (lb LoadBalancer) GetAddress() string {
	return lb.Address
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
	UUID             string
	Name             string
	Type             string
	ConnectionNumber int
	DataTransfer     int
	Mode             string
	LbType           string
	DisplayLbType    string
}
