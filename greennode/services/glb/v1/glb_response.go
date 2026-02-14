package v1

import (
	"github.com/dannyota/greennode-community-sdk/v2/greennode/entity"
)

type (
	GlobalLoadBalancerResponse struct {
		CreatedAt   string                             `json:"createdAt"`
		UpdatedAt   string                             `json:"updatedAt"`
		DeletedAt   string                             `json:"deletedAt"`
		ID          string                             `json:"id"`
		Name        string                             `json:"name"`
		Description string                             `json:"description"`
		Status      string                             `json:"status"`
		Package     string                             `json:"package"`
		Type        string                             `json:"type"`
		UserID      int                                `json:"userId"`
		Vips        []GlobalLoadBalancerVIPResponse    `json:"vips"`
		Domains     []GlobalLoadBalancerDomainResponse `json:"domains"`
	}

	GlobalLoadBalancerVIPResponse struct {
		ID                   int    `json:"id"`
		CreatedAt            string `json:"createdAt"`
		UpdatedAt            string `json:"updatedAt"`
		DeletedAt            string `json:"deletedAt"`
		Address              string `json:"address"`
		Status               string `json:"status"`
		Region               string `json:"region"`
		GlobalLoadBalancerID string `json:"globalLoadBalancerId"`
	}

	GlobalLoadBalancerDomainResponse struct {
		CreatedAt            string `json:"createdAt"`
		UpdatedAt            string `json:"updatedAt"`
		DeletedAt            string `json:"deletedAt"`
		ID                   int    `json:"id"`
		Hostname             string `json:"hostname"`
		Status               string `json:"status"`
		GlobalLoadBalancerID string `json:"globalLoadBalancerId"`
		DNSHostedZoneID      string `json:"dnsHostedZoneId"`
		DNSServerID          string `json:"dnsServerId"`
	}
)

func (r *GlobalLoadBalancerResponse) ToEntityGlobalLoadBalancer() *entity.GlobalLoadBalancer {
	vips := make([]*entity.GlobalLoadBalancerVIP, 0, len(r.Vips))
	for _, vip := range r.Vips {
		vips = append(vips, &entity.GlobalLoadBalancerVIP{
			ID:                   vip.ID,
			CreatedAt:            vip.CreatedAt,
			UpdatedAt:            vip.UpdatedAt,
			DeletedAt:            vip.DeletedAt,
			Address:              vip.Address,
			Status:               vip.Status,
			Region:               vip.Region,
			GlobalLoadBalancerID: vip.GlobalLoadBalancerID,
		})
	}

	domains := make([]*entity.GlobalLoadBalancerDomain, 0, len(r.Domains))
	for _, domain := range r.Domains {
		domains = append(domains, &entity.GlobalLoadBalancerDomain{
			CreatedAt:            domain.CreatedAt,
			UpdatedAt:            domain.UpdatedAt,
			DeletedAt:            domain.DeletedAt,
			ID:                   domain.ID,
			Hostname:             domain.Hostname,
			Status:               domain.Status,
			GlobalLoadBalancerID: domain.GlobalLoadBalancerID,
			DNSHostedZoneID:      domain.DNSHostedZoneID,
			DNSServerID:          domain.DNSServerID,
		})
	}

	return &entity.GlobalLoadBalancer{
		CreatedAt:   r.CreatedAt,
		UpdatedAt:   r.UpdatedAt,
		DeletedAt:   r.DeletedAt,
		ID:          r.ID,
		Name:        r.Name,
		Description: r.Description,
		Status:      r.Status,
		Package:     r.Package,
		Type:        r.Type,
		UserID:      r.UserID,
		Vips:        vips,
		Domains:     domains,
	}
}

// --------------------------------------------------
type ListGlobalLoadBalancersResponse struct {
	Items  []GlobalLoadBalancerResponse `json:"items"`
	Limit  int                          `json:"limit"`
	Total  int                          `json:"total"`
	Offset int                          `json:"offset"`
}

func (r *ListGlobalLoadBalancersResponse) ToEntityListGlobalLoadBalancers() *entity.ListGlobalLoadBalancers {
	result := &entity.ListGlobalLoadBalancers{
		Items:  make([]*entity.GlobalLoadBalancer, 0),
		Limit:  0,
		Total:  0,
		Offset: 0,
	}

	if r == nil || r.Items == nil || len(r.Items) < 1 {
		return result
	}

	result.Limit = r.Limit
	result.Total = r.Total
	result.Offset = r.Offset

	for _, itemLb := range r.Items {
		result.Items = append(result.Items, itemLb.ToEntityGlobalLoadBalancer())
	}

	return result
}

// --------------------------------------------------

type CreateGlobalLoadBalancerResponse struct {
	GlobalLoadBalancer GlobalLoadBalancerResponse `json:"globalLoadBalancer"`
	GlobalListener     GlobalListenerResponse     `json:"globalListener"`
	GlobalPool         GlobalPoolResponse         `json:"globalPool"`
}

func (r *CreateGlobalLoadBalancerResponse) ToEntityGlobalLoadBalancer() *entity.GlobalLoadBalancer {
	return r.GlobalLoadBalancer.ToEntityGlobalLoadBalancer()
}

// --------------------------------------------------

type GetGlobalLoadBalancerByIDResponse GlobalLoadBalancerResponse

func (r *GetGlobalLoadBalancerByIDResponse) ToEntityGlobalLoadBalancer() *entity.GlobalLoadBalancer {
	return (*GlobalLoadBalancerResponse)(r).ToEntityGlobalLoadBalancer()
}

// --------------------------------------------------

type ListGlobalPackagesResponse []GlobalPackageResponse

type GlobalPackageResponse struct {
	ID                          string                     `json:"id"`
	Name                        string                     `json:"name"`
	Description                 string                     `json:"description"`
	DescriptionEn               string                     `json:"descriptionEn"`
	Detail                      any                        `json:"detail"`
	Enabled                     bool                       `json:"enabled"`
	BaseSku                     string                     `json:"baseSku"`
	BaseConnectionRate          int                        `json:"baseConnectionRate"`
	BaseDomesticTrafficTotal    int                        `json:"baseDomesticTrafficTotal"`
	BaseNonDomesticTrafficTotal int                        `json:"baseNonDomesticTrafficTotal"`
	ConnectionSku               string                     `json:"connectionSku"`
	DomesticTrafficSku          string                     `json:"domesticTrafficSku"`
	NonDomesticTrafficSku       string                     `json:"nonDomesticTrafficSku"`
	CreatedAt                   string                     `json:"createdAt"`
	UpdatedAt                   string                     `json:"updatedAt"`
	VlbPackages                 []VlbGlobalPackageResponse `json:"vlbPackages"`
}

type VlbGlobalPackageResponse struct {
	ID           int    `json:"id"`
	GlbPackageID string `json:"glb_package_id"`
	Region       string `json:"region"`
	VlbPackageID string `json:"vlb_package_id"`
	CreatedAt    string `json:"created_at"`
}

func (r *ListGlobalPackagesResponse) ToEntityListGlobalPackages() *entity.ListGlobalPackages {
	packages := make([]entity.GlobalPackage, 0)
	if r != nil {
		for _, item := range *r {
			packages = append(packages, *item.ToEntityGlobalPackage())
		}
	}
	return &entity.ListGlobalPackages{Items: packages}
}

func (r *GlobalPackageResponse) ToEntityGlobalPackage() *entity.GlobalPackage {
	vlbPackages := make([]entity.VlbPackage, 0, len(r.VlbPackages))
	for _, vlb := range r.VlbPackages {
		vlbPackages = append(vlbPackages, entity.VlbPackage{
			ID:           vlb.ID,
			GlbPackageID: vlb.GlbPackageID,
			Region:       vlb.Region,
			VlbPackageID: vlb.VlbPackageID,
			CreatedAt:    vlb.CreatedAt,
		})
	}

	return &entity.GlobalPackage{
		ID:                          r.ID,
		Name:                        r.Name,
		Description:                 r.Description,
		DescriptionEn:               r.DescriptionEn,
		Detail:                      r.Detail,
		Enabled:                     r.Enabled,
		BaseSku:                     r.BaseSku,
		BaseConnectionRate:          r.BaseConnectionRate,
		BaseDomesticTrafficTotal:    r.BaseDomesticTrafficTotal,
		BaseNonDomesticTrafficTotal: r.BaseNonDomesticTrafficTotal,
		ConnectionSku:               r.ConnectionSku,
		DomesticTrafficSku:          r.DomesticTrafficSku,
		NonDomesticTrafficSku:       r.NonDomesticTrafficSku,
		CreatedAt:                   r.CreatedAt,
		UpdatedAt:                   r.UpdatedAt,
		VlbPackages:                 vlbPackages,
	}
}

// --------------------------------------------------

type ListGlobalRegionsResponse []GlobalRegionResponse

type GlobalRegionResponse struct {
	ID               string `json:"id"`
	Name             string `json:"name"`
	VServerEndpoint  string `json:"vserverEndpoint"`
	VlbEndpoint      string `json:"vlbEndpoint"`
	UIServerEndpoint string `json:"uiServerEndpoint"`
}

func (r *ListGlobalRegionsResponse) ToEntityListGlobalRegions() *entity.ListGlobalRegions {
	regions := make([]entity.GlobalRegion, 0)
	if r != nil {
		for _, item := range *r {
			regions = append(regions, *item.ToEntityGlobalRegion())
		}
	}
	return &entity.ListGlobalRegions{Items: regions}
}

func (r *GlobalRegionResponse) ToEntityGlobalRegion() *entity.GlobalRegion {
	return &entity.GlobalRegion{
		ID:               r.ID,
		Name:             r.Name,
		VServerEndpoint:  r.VServerEndpoint,
		VlbEndpoint:      r.VlbEndpoint,
		UIServerEndpoint: r.UIServerEndpoint,
	}
}

// --------------------------------------------------

type GetGlobalLoadBalancerUsageHistoriesResponse struct {
	Type  string                                   `json:"type"`
	Items []GlobalLoadBalancerUsageHistoryResponse `json:"items"`
	From  string                                   `json:"from"`
	To    string                                   `json:"to"`
}

type GlobalLoadBalancerUsageHistoryResponse struct {
	Timestamp string  `json:"timestamp"`
	Value     float64 `json:"value"`
	Type      string  `json:"type"`
}

func (r *GetGlobalLoadBalancerUsageHistoriesResponse) ToEntityGlobalLoadBalancerUsageHistories() *entity.ListGlobalLoadBalancerUsageHistories {
	histories := make([]entity.GlobalLoadBalancerUsageHistory, 0)
	if r != nil && r.Items != nil {
		for _, item := range r.Items {
			histories = append(histories, *item.ToEntityGlobalLoadBalancerUsageHistory())
		}
	}
	return &entity.ListGlobalLoadBalancerUsageHistories{
		Type:  r.Type,
		Items: histories,
		From:  r.From,
		To:    r.To,
	}
}

func (r *GlobalLoadBalancerUsageHistoryResponse) ToEntityGlobalLoadBalancerUsageHistory() *entity.GlobalLoadBalancerUsageHistory {
	return &entity.GlobalLoadBalancerUsageHistory{
		Timestamp: r.Timestamp,
		Value:     r.Value,
		Type:      r.Type,
	}
}
