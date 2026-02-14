package v2

import (
	"strings"

	"github.com/dannyota/greennode-community-sdk/v2/greennode/entity"
)

type CreateLoadBalancerResponse struct {
	UUID string `json:"uuid"`
}

type ResizeLoadBalancerResponse struct {
	UUID string `json:"uuid"`
}

func (r *ResizeLoadBalancerResponse) ToEntityLoadBalancer() *entity.LoadBalancer {
	return &entity.LoadBalancer{
		UUID: r.UUID,
	}
}

type ScaleLoadBalancerResponse struct {
	UUID string `json:"uuid"`
}

func (r *ScaleLoadBalancerResponse) ToEntityLoadBalancer() *entity.LoadBalancer {
	return &entity.LoadBalancer{
		UUID: r.UUID,
	}
}

type ListLoadBalancerPackagesResponse struct {
	ListData []LoadBalancerPackageResponse `json:"listData"`
}

type LoadBalancerPackageResponse struct {
	UUID             string `json:"uuid"`
	Name             string `json:"name"`
	Type             string `json:"type"`
	ConnectionNumber int    `json:"connectionNumber"`
	DataTransfer     int    `json:"dataTransfer"`
	Mode             string `json:"mode"`
	LbType           string `json:"lbType"`
	DisplayLbType    string `json:"displayLbType"`
}

type GetLoadBalancerByIDResponse struct {
	Data LoadBalancer `json:"data"`
}

type ListLoadBalancersResponse struct {
	ListData  []LoadBalancer `json:"listData"`
	Page      int            `json:"page"`
	PageSize  int            `json:"pageSize"`
	TotalPage int            `json:"totalPage"`
	TotalItem int            `json:"totalItem"`
}

type (
	LoadBalancer struct {
		UUID               string `json:"uuid"`
		Name               string `json:"name"`
		DisplayStatus      string `json:"displayStatus"`
		Address            string `json:"address"`
		PrivateSubnetID    string `json:"privateSubnetId"`
		PrivateSubnetCidr  string `json:"privateSubnetCidr"`
		Type               string `json:"type"`
		DisplayType        string `json:"displayType"`
		LoadBalancerSchema string `json:"loadBalancerSchema"`
		PackageID          string `json:"packageId"`
		Description        string `json:"description"`
		Location           string `json:"location"`
		CreatedAt          string `json:"createdAt"`
		UpdatedAt          string `json:"updatedAt"`
		BackendSubnetID    string `json:"backendSubnetId"`
		PackageInfo        struct {
			PackageID        string `json:"packageId"`
			ConnectionNumber int    `json:"connectionNumber"`
			DataTransfer     int    `json:"dataTransfer"`
			Name             string `json:"name"`
		} `json:"packageInfo"`
		ProgressStatus string `json:"progressStatus"`
		AutoScalable   bool   `json:"autoScalable"`
		Zone           struct {
			UUID string `json:"uuid"`
			// Name        string `json:"name"`
			// ZoneType    string `json:"zoneType"`
			// IsDefault   bool   `json:"isDefault"`
			// Description string `json:"description"`
			// IsEnabled   bool   `json:"isEnabled"`
			// VolumeCount int    `json:"volumeCount"`
			// ServerCount int    `json:"serverCount"`
		} `json:"zone"`
		MinSize    int    `json:"minSize"`
		MaxSize    int    `json:"maxSize"`
		TotalNodes int    `json:"totalNodes"`
		Nodes      []Node `json:"nodes"`
	}

	Node struct {
		Status   string `json:"status"`
		ZoneID   string `json:"zoneId"`
		ZoneName string `json:"zoneName"`
		SubnetID string `json:"subnetId"`
	}
)

func (r *CreateLoadBalancerResponse) ToEntityLoadBalancer() *entity.LoadBalancer {
	return &entity.LoadBalancer{
		UUID: r.UUID,
	}
}

func (r *ListLoadBalancerPackagesResponse) ToEntityListLoadBalancerPackages() *entity.ListLoadBalancerPackages {
	if r == nil || r.ListData == nil || len(r.ListData) < 1 {
		return &entity.ListLoadBalancerPackages{
			Items: make([]*entity.LoadBalancerPackage, 0),
		}
	}

	result := &entity.ListLoadBalancerPackages{
		Items: make([]*entity.LoadBalancerPackage, 0),
	}

	for _, item := range r.ListData {
		result.Items = append(result.Items, &entity.LoadBalancerPackage{
			UUID:             item.UUID,
			Name:             item.Name,
			Type:             item.Type,
			ConnectionNumber: item.ConnectionNumber,
			DataTransfer:     item.DataTransfer,
			Mode:             item.Mode,
			LbType:           item.LbType,
			DisplayLbType:    item.DisplayLbType,
		})
	}

	return result
}

func (r *GetLoadBalancerByIDResponse) ToEntityLoadBalancer() *entity.LoadBalancer {
	if r == nil {
		return nil
	}

	return r.Data.toEntityLoadBalancer()
}

func (r *ListLoadBalancersResponse) ToEntityListLoadBalancers() *entity.ListLoadBalancers {
	if r == nil || r.ListData == nil || len(r.ListData) < 1 {
		return &entity.ListLoadBalancers{
			Page:      0,
			PageSize:  0,
			TotalPage: 0,
			TotalItem: 0,
			Items:     make([]*entity.LoadBalancer, 0),
		}
	}

	result := &entity.ListLoadBalancers{
		Page:      r.Page,
		PageSize:  r.PageSize,
		TotalPage: r.TotalPage,
		TotalItem: r.TotalItem,
	}

	for _, itemLb := range r.ListData {
		result.Add(itemLb.toEntityLoadBalancer())
	}

	return result
}

func (lb *LoadBalancer) toEntityLoadBalancer() *entity.LoadBalancer {
	internal := strings.TrimSpace(strings.ToUpper(lb.LoadBalancerSchema)) == "INTERNAL"

	// Convert nodes from response to entity
	nodes := make([]*entity.Node, 0, len(lb.Nodes))
	for _, node := range lb.Nodes {
		nodes = append(nodes, &entity.Node{
			Status:   node.Status,
			ZoneID:   node.ZoneID,
			ZoneName: node.ZoneName,
			SubnetID: node.SubnetID,
		})
	}

	return &entity.LoadBalancer{
		UUID:               lb.UUID,
		Name:               lb.Name,
		Address:            lb.Address,
		DisplayStatus:      lb.DisplayStatus,
		PrivateSubnetID:    lb.PrivateSubnetID,
		PrivateSubnetCidr:  lb.PrivateSubnetCidr,
		Type:               lb.Type,
		DisplayType:        lb.DisplayType,
		LoadBalancerSchema: lb.LoadBalancerSchema,
		PackageID:          lb.PackageID,
		Description:        lb.Description,
		Location:           lb.Location,
		CreatedAt:          lb.CreatedAt,
		UpdatedAt:          lb.UpdatedAt,
		ProgressStatus:     lb.ProgressStatus,
		AutoScalable:       lb.AutoScalable,
		ZoneID:             lb.Zone.UUID,
		MinSize:            lb.MinSize,
		MaxSize:            lb.MaxSize,
		TotalNodes:         lb.TotalNodes,
		Nodes:              nodes,
		BackendSubnetID: func() string {
			if lb.BackendSubnetID != "" {
				return lb.BackendSubnetID
			}
			return lb.PrivateSubnetID
		}(),

		// will be removed
		Status:   lb.DisplayStatus,
		Internal: internal,
	}
}
