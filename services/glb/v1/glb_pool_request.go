package v1

import "danny.vn/greennode/services/common"

type (
	GlobalPoolAlgorithm              string
	GlobalPoolProtocol               string
	GlobalPoolHealthCheckProtocol    string
	GlobalPoolHealthCheckMethod      string
	GlobalPoolHealthCheckHTTPVersion string

	GlobalPoolMemberType string
)

const (
	GlobalPoolAlgorithmRoundRobin GlobalPoolAlgorithm = "ROUND_ROBIN"
	GlobalPoolAlgorithmLeastConn  GlobalPoolAlgorithm = "LEAST_CONNECTIONS"
	GlobalPoolAlgorithmSourceIP   GlobalPoolAlgorithm = "SOURCE_IP"
)

const (
	GlobalPoolProtocolTCP   GlobalPoolProtocol = "TCP"
	GlobalPoolProtocolProxy GlobalPoolProtocol = "PROXY"
)

const (
	GlobalPoolHealthCheckProtocolTCP   GlobalPoolHealthCheckProtocol = "TCP"
	GlobalPoolHealthCheckProtocolHTTP  GlobalPoolHealthCheckProtocol = "HTTP"
	GlobalPoolHealthCheckProtocolHTTPS GlobalPoolHealthCheckProtocol = "HTTPS"
)

const (
	GlobalPoolHealthCheckMethodGET  GlobalPoolHealthCheckMethod = "GET"
	GlobalPoolHealthCheckMethodPUT  GlobalPoolHealthCheckMethod = "PUT"
	GlobalPoolHealthCheckMethodPOST GlobalPoolHealthCheckMethod = "POST"
)

const (
	GlobalPoolHealthCheckHTTPVersionHTTP1       GlobalPoolHealthCheckHTTPVersion = "1.0"
	GlobalPoolHealthCheckHTTPVersionHTTP1Minor1 GlobalPoolHealthCheckHTTPVersion = "1.1"
)

const (
	GlobalPoolMemberTypePublic  GlobalPoolMemberType = "PUBLIC"
	GlobalPoolMemberTypePrivate GlobalPoolMemberType = "PRIVATE"
)

type ListGlobalPoolsRequest struct {
	LoadBalancerID string
}

func NewListGlobalPoolsRequest(lbID string) *ListGlobalPoolsRequest {
	opts := &ListGlobalPoolsRequest{}
	opts.LoadBalancerID = lbID
	return opts
}

type CreateGlobalPoolRequest struct {
	Algorithm         GlobalPoolAlgorithm         `json:"algorithm"`
	Description       string                      `json:"description,omitempty"`
	Name              string                      `json:"name"`
	Protocol          GlobalPoolProtocol          `json:"protocol"`
	Stickiness        *bool                       `json:"stickiness,omitempty"`
	TLSEncryption     *bool                       `json:"tlsEncryption,omitempty"`
	HealthMonitor     *GlobalHealthMonitorRequest `json:"health"`
	GlobalPoolMembers []*GlobalPoolMemberRequest  `json:"globalPoolMembers"`

	LoadBalancerID string
}

func NewCreateGlobalPoolRequest(name string, protocol GlobalPoolProtocol) *CreateGlobalPoolRequest {
	opts := &CreateGlobalPoolRequest{
		Name:              name,
		Protocol:          protocol,
		Algorithm:         GlobalPoolAlgorithmRoundRobin,
		Description:       "",
		Stickiness:        nil,
		TLSEncryption:     nil,
		HealthMonitor:     nil,
		GlobalPoolMembers: make([]*GlobalPoolMemberRequest, 0),
	}

	return opts
}

type GlobalHealthMonitorRequest struct {
	HealthCheckProtocol GlobalPoolHealthCheckProtocol     `json:"protocol"`
	HealthyThreshold    int                               `json:"healthyThreshold"`
	UnhealthyThreshold  int                               `json:"unhealthyThreshold"`
	Interval            int                               `json:"interval"`
	Timeout             int                               `json:"timeout"`
	HTTPMethod          *GlobalPoolHealthCheckMethod      `json:"httpMethod,omitempty"`
	HTTPVersion         *GlobalPoolHealthCheckHTTPVersion `json:"httpVersion,omitempty"`
	Path                *string                           `json:"path,omitempty"`
	DomainName          *string                           `json:"domainName,omitempty"`
	SuccessCode         *string                           `json:"successCode,omitempty"`
}

func NewGlobalHealthMonitor(checkProtocol GlobalPoolHealthCheckProtocol) *GlobalHealthMonitorRequest {
	opts := &GlobalHealthMonitorRequest{
		HealthCheckProtocol: checkProtocol,
		HealthyThreshold:    3,
		UnhealthyThreshold:  3,
		Interval:            30,
		Timeout:             5,
		HTTPMethod:          nil,
		HTTPVersion:         nil,
		Path:                nil,
		DomainName:          nil,
		SuccessCode:         nil,
	}
	if checkProtocol == GlobalPoolHealthCheckProtocolHTTP || checkProtocol == GlobalPoolHealthCheckProtocolHTTPS {
		opts.HTTPMethod = common.Ptr(GlobalPoolHealthCheckMethodGET)
		opts.HTTPVersion = common.Ptr(GlobalPoolHealthCheckHTTPVersionHTTP1Minor1)
		opts.Path = common.Ptr("/")
		opts.DomainName = common.Ptr("")
		opts.SuccessCode = common.Ptr("200")
	}
	return opts
}

type GlobalPoolMemberRequest struct {
	Name        string                 `json:"name"`
	Description string                 `json:"description,omitempty"`
	Region      string                 `json:"region"`
	TrafficDial int                    `json:"trafficDial"`
	VPCID       string                 `json:"vpcId"`
	Type        GlobalPoolMemberType   `json:"type"`
	Members     []*GlobalMemberRequest `json:"members"`

	LoadBalancerID string
	PoolID         string
}

func NewGlobalPoolMemberRequest(name, region, vpcID string, dial int, typeVal GlobalPoolMemberType) *GlobalPoolMemberRequest {
	opts := &GlobalPoolMemberRequest{
		Name:        name,
		Description: "",
		Region:      region,
		TrafficDial: dial,
		Type:        typeVal,
		VPCID:       vpcID,
		Members:     make([]*GlobalMemberRequest, 0),
	}
	return opts
}

type GlobalMemberRequest struct {
	Address     string `json:"address"`
	BackupRole  bool   `json:"backupRole"`
	Description string `json:"description,omitempty"`
	MonitorPort int    `json:"monitorPort"`
	Name        string `json:"name"`
	Port        int    `json:"port"`
	SubnetID    string `json:"subnetId"`
	Weight      int    `json:"weight"`
}

func NewGlobalMemberRequest(name, address, subnetID string, port, monitorPort, weight int, backupRole bool) *GlobalMemberRequest {
	opts := &GlobalMemberRequest{
		Name:        name,
		Address:     address,
		BackupRole:  backupRole,
		Description: "",
		MonitorPort: monitorPort,
		Port:        port,
		SubnetID:    subnetID,
		Weight:      weight,
	}
	return opts
}

type UpdateGlobalPoolRequest struct {
	Algorithm      GlobalPoolAlgorithm        `json:"algorithm"`
	HealthMonitor  *GlobalHealthMonitorRequest `json:"health"`
	LoadBalancerID string
	PoolID         string
}

func NewUpdateGlobalPoolRequest(lbID, poolID string) *UpdateGlobalPoolRequest {
	opts := &UpdateGlobalPoolRequest{
		Algorithm:      GlobalPoolAlgorithmRoundRobin,
		HealthMonitor:  nil,
		LoadBalancerID: lbID,
		PoolID:         poolID,
	}
	return opts
}

type DeleteGlobalPoolRequest struct {
	LoadBalancerID string
	PoolID         string
}

func NewDeleteGlobalPoolRequest(lbID, poolID string) *DeleteGlobalPoolRequest {
	opts := &DeleteGlobalPoolRequest{
		LoadBalancerID: lbID,
		PoolID:         poolID,
	}
	return opts
}

type ListGlobalPoolMembersRequest struct {
	LoadBalancerID string
	PoolID         string
}

func NewListGlobalPoolMembersRequest(lbID, poolID string) *ListGlobalPoolMembersRequest {
	opts := &ListGlobalPoolMembersRequest{
		LoadBalancerID: lbID,
		PoolID:         poolID,
	}
	return opts
}

type PatchGlobalPoolMembersRequest struct {
	BulkActions    []any `json:"bulkActions"`
	LoadBalancerID string
	PoolID         string
}

func NewPatchGlobalPoolMembersRequest(lbID, poolID string) *PatchGlobalPoolMembersRequest {
	opts := &PatchGlobalPoolMembersRequest{
		BulkActions:    make([]any, 0),
		LoadBalancerID: lbID,
		PoolID:         poolID,
	}
	return opts
}

type PatchGlobalPoolCreateBulkActionRequest struct {
	Action           string                   `json:"action"`
	CreatePoolMember *GlobalPoolMemberRequest `json:"createPoolMember"`
}

func NewPatchGlobalPoolCreateBulkActionRequest(member *GlobalPoolMemberRequest) *PatchGlobalPoolCreateBulkActionRequest {
	opts := &PatchGlobalPoolCreateBulkActionRequest{
		Action:           "create",
		CreatePoolMember: member,
	}
	return opts
}

type PatchGlobalPoolDeleteBulkActionRequest struct {
	Action string `json:"action"`
	ID     string `json:"id"`
}

func NewPatchGlobalPoolDeleteBulkActionRequest(id string) *PatchGlobalPoolDeleteBulkActionRequest {
	opts := &PatchGlobalPoolDeleteBulkActionRequest{
		Action: "delete",
		ID:     id,
	}
	return opts
}

type PatchGlobalPoolUpdateBulkActionRequest struct {
	Action           string                         `json:"action"`
	ID               string                         `json:"id"`
	UpdatePoolMember *UpdateGlobalPoolMemberRequest `json:"updatePoolMember"`
}

func NewPatchGlobalPoolUpdateBulkActionRequest(id string, member *UpdateGlobalPoolMemberRequest) *PatchGlobalPoolUpdateBulkActionRequest {
	opts := &PatchGlobalPoolUpdateBulkActionRequest{
		Action:           "update",
		ID:               id,
		UpdatePoolMember: member,
	}
	return opts
}

type UpdateGlobalPoolMemberRequest struct {
	TrafficDial int                    `json:"trafficDial"`
	Members     []*GlobalMemberRequest `json:"members"`

	LoadBalancerID string
	PoolID         string
	PoolMemberID   string
}

func NewUpdateGlobalPoolMemberRequest(lbID, poolID, poolMemberID string, dial int) *UpdateGlobalPoolMemberRequest {
	opts := &UpdateGlobalPoolMemberRequest{
		TrafficDial:    dial,
		Members:        make([]*GlobalMemberRequest, 0),
		LoadBalancerID: lbID,
		PoolID:         poolID,
		PoolMemberID:   poolMemberID,
	}
	return opts
}

type DeleteGlobalPoolMemberRequest struct {
	LoadBalancerID string
	PoolID         string
	PoolMemberID   string
}

func NewDeleteGlobalPoolMemberRequest(lbID, poolID, poolMemberID string) *DeleteGlobalPoolMemberRequest {
	opts := &DeleteGlobalPoolMemberRequest{
		LoadBalancerID: lbID,
		PoolID:         poolID,
		PoolMemberID:   poolMemberID,
	}
	return opts
}

type GetGlobalPoolMemberRequest struct {
	LoadBalancerID string
	PoolID         string
	PoolMemberID   string
}

func NewGetGlobalPoolMemberRequest(lbID, poolID, poolMemberID string) *GetGlobalPoolMemberRequest {
	opts := &GetGlobalPoolMemberRequest{
		LoadBalancerID: lbID,
		PoolID:         poolID,
		PoolMemberID:   poolMemberID,
	}
	return opts
}
