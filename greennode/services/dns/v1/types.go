package v1

import "time"

type VpcMapRegion struct {
	VpcID  string `json:"vpcId"`
	Region string `json:"region"`
}

type HostedZone struct {
	HostedZoneID      string         `json:"hostedZoneId"`
	DomainName        string         `json:"domainName"`
	Status            string         `json:"status"`
	Description       string         `json:"description"`
	Type              string         `json:"type"`
	CountRecords      int            `json:"countRecords"`
	AssocVpcIDs       []string       `json:"assocVpcIds"`
	AssocVpcMapRegion []VpcMapRegion `json:"assocVpcMapRegion"`
	PortalUserID      int            `json:"portalUserId"`
	CreatedAt         time.Time      `json:"createdAt"`
	DeletedAt         *time.Time     `json:"deletedAt"`
	UpdatedAt         time.Time      `json:"updatedAt"`
}

type ListHostedZones struct {
	ListData  []*HostedZone `json:"listData"`
	Page      int           `json:"page"`
	PageSize  int           `json:"pageSize"`
	TotalPage int           `json:"totalPage"`
	TotalItem int           `json:"totalItem"`
}

type RecordValue struct {
	Value    string  `json:"value"`
	Location *string `json:"location"`
	Weight   *int    `json:"weight"`
}

type DnsRecord struct {
	RecordID            string        `json:"recordId"`
	SubDomain           string        `json:"subDomain"`
	HostedZoneID        string        `json:"hostedZoneId"`
	Status              string        `json:"status"`
	Type                string        `json:"type"`
	RoutingPolicy       string        `json:"routingPolicy"`
	Value               []RecordValue `json:"value"`
	TTL                 int           `json:"ttl"`
	EnableStickySession *bool         `json:"enableStickySession"`
	CreatedAt           time.Time     `json:"createdAt"`
	DeletedAt           *time.Time    `json:"deletedAt"`
	UpdatedAt           time.Time     `json:"updatedAt"`
}

type ListDnsRecords struct {
	ListData  []*DnsRecord `json:"listData"`
	Page      int          `json:"page"`
	PageSize  int          `json:"pageSize"`
	TotalPage int          `json:"totalPage"`
	TotalItem int          `json:"totalItem"`
}
