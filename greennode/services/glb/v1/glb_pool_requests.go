package v1

import "github.com/dannyota/greennode-community-sdk/v2/greennode/services/common"

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
	common.LoadBalancerCommon
}

func (r *ListGlobalPoolsRequest) WithLoadBalancerID(lbID string) *ListGlobalPoolsRequest {
	r.LoadBalancerID = lbID
	return r
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

	common.LoadBalancerCommon
}

func (r *CreateGlobalPoolRequest) WithAlgorithm(algorithm GlobalPoolAlgorithm) *CreateGlobalPoolRequest {
	r.Algorithm = algorithm
	return r
}

func (r *CreateGlobalPoolRequest) WithDescription(description string) *CreateGlobalPoolRequest {
	r.Description = description
	return r
}

func (r *CreateGlobalPoolRequest) WithName(name string) *CreateGlobalPoolRequest {
	r.Name = name
	return r
}

func (r *CreateGlobalPoolRequest) WithProtocol(protocol GlobalPoolProtocol) *CreateGlobalPoolRequest {
	r.Protocol = protocol
	return r
}

func (r *CreateGlobalPoolRequest) WithHealthMonitor(health *GlobalHealthMonitorRequest) *CreateGlobalPoolRequest {
	r.HealthMonitor = health
	return r
}

func (r *CreateGlobalPoolRequest) WithMembers(members ...*GlobalPoolMemberRequest) *CreateGlobalPoolRequest {
	r.GlobalPoolMembers = append(r.GlobalPoolMembers, members...)
	return r
}

func (r *CreateGlobalPoolRequest) WithLoadBalancerID(lbID string) *CreateGlobalPoolRequest {
	r.LoadBalancerID = lbID
	return r
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

func (r *GlobalHealthMonitorRequest) WithHealthyThreshold(threshold int) *GlobalHealthMonitorRequest {
	r.HealthyThreshold = threshold
	return r
}

func (r *GlobalHealthMonitorRequest) WithUnhealthyThreshold(threshold int) *GlobalHealthMonitorRequest {
	r.UnhealthyThreshold = threshold
	return r
}

func (r *GlobalHealthMonitorRequest) WithProtocol(protocol GlobalPoolHealthCheckProtocol) *GlobalHealthMonitorRequest {
	r.HealthCheckProtocol = protocol
	return r
}

func (r *GlobalHealthMonitorRequest) WithInterval(interval int) *GlobalHealthMonitorRequest {
	r.Interval = interval
	return r
}

func (r *GlobalHealthMonitorRequest) WithTimeout(timeout int) *GlobalHealthMonitorRequest {
	r.Timeout = timeout
	return r
}

func (r *GlobalHealthMonitorRequest) WithHealthCheckMethod(method *GlobalPoolHealthCheckMethod) *GlobalHealthMonitorRequest {
	r.HTTPMethod = method
	return r
}

func (r *GlobalHealthMonitorRequest) WithHTTPVersion(version *GlobalPoolHealthCheckHTTPVersion) *GlobalHealthMonitorRequest {
	r.HTTPVersion = version
	return r
}

func (r *GlobalHealthMonitorRequest) WithDomainName(domain *string) *GlobalHealthMonitorRequest {
	r.DomainName = domain
	return r
}

func (r *GlobalHealthMonitorRequest) WithSuccessCode(code *string) *GlobalHealthMonitorRequest {
	r.SuccessCode = code
	return r
}

func (r *GlobalHealthMonitorRequest) WithPath(path *string) *GlobalHealthMonitorRequest {
	r.Path = path
	return r
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

	common.LoadBalancerCommon
	common.PoolCommon
}

func (r *GlobalPoolMemberRequest) WithName(name string) *GlobalPoolMemberRequest {
	r.Name = name
	return r
}

func (r *GlobalPoolMemberRequest) WithDescription(description string) *GlobalPoolMemberRequest {
	r.Description = description
	return r
}

func (r *GlobalPoolMemberRequest) WithRegion(region string) *GlobalPoolMemberRequest {
	r.Region = region
	return r
}

func (r *GlobalPoolMemberRequest) WithTrafficDial(dial int) *GlobalPoolMemberRequest {
	r.TrafficDial = dial
	return r
}

func (r *GlobalPoolMemberRequest) WithVPCID(vpcID string) *GlobalPoolMemberRequest {
	r.VPCID = vpcID
	return r
}

func (r *GlobalPoolMemberRequest) WithType(typeVal GlobalPoolMemberType) *GlobalPoolMemberRequest {
	r.Type = typeVal
	return r
}

func (r *GlobalPoolMemberRequest) WithMembers(members ...*GlobalMemberRequest) *GlobalPoolMemberRequest {
	r.Members = append(r.Members, members...)
	return r
}

func (r *GlobalPoolMemberRequest) WithLoadBalancerID(lbID string) *GlobalPoolMemberRequest {
	r.LoadBalancerID = lbID
	return r
}

func (r *GlobalPoolMemberRequest) WithPoolID(poolID string) *GlobalPoolMemberRequest {
	r.PoolID = poolID
	return r
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

func (r *GlobalMemberRequest) WithAddress(address string) *GlobalMemberRequest {
	r.Address = address
	return r
}

func (r *GlobalMemberRequest) WithBackupRole(backupRole bool) *GlobalMemberRequest {
	r.BackupRole = backupRole
	return r
}

func (r *GlobalMemberRequest) WithDescription(description string) *GlobalMemberRequest {
	r.Description = description
	return r
}

func (r *GlobalMemberRequest) WithMonitorPort(monitorPort int) *GlobalMemberRequest {
	r.MonitorPort = monitorPort
	return r
}

func (r *GlobalMemberRequest) WithName(name string) *GlobalMemberRequest {
	r.Name = name
	return r
}

func (r *GlobalMemberRequest) WithPort(port int) *GlobalMemberRequest {
	r.Port = port
	return r
}

func (r *GlobalMemberRequest) WithSubnetID(subnetID string) *GlobalMemberRequest {
	r.SubnetID = subnetID
	return r
}

func (r *GlobalMemberRequest) WithWeight(weight int) *GlobalMemberRequest {
	r.Weight = weight
	return r
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
	Algorithm     GlobalPoolAlgorithm         `json:"algorithm"`
	HealthMonitor *GlobalHealthMonitorRequest `json:"health"`
	common.LoadBalancerCommon
	common.PoolCommon
}

func (r *UpdateGlobalPoolRequest) WithAlgorithm(algorithm GlobalPoolAlgorithm) *UpdateGlobalPoolRequest {
	r.Algorithm = algorithm
	return r
}

func (r *UpdateGlobalPoolRequest) WithHealthMonitor(monitor *GlobalHealthMonitorRequest) *UpdateGlobalPoolRequest {
	r.HealthMonitor = monitor
	return r
}

func (r *UpdateGlobalPoolRequest) WithLoadBalancerID(lbID string) *UpdateGlobalPoolRequest {
	r.LoadBalancerID = lbID
	return r
}

func (r *UpdateGlobalPoolRequest) WithPoolID(poolID string) *UpdateGlobalPoolRequest {
	r.PoolID = poolID
	return r
}

func NewUpdateGlobalPoolRequest(lbID, poolID string) *UpdateGlobalPoolRequest {
	opts := &UpdateGlobalPoolRequest{
		Algorithm:     GlobalPoolAlgorithmRoundRobin,
		HealthMonitor: nil,
		LoadBalancerCommon: common.LoadBalancerCommon{
			LoadBalancerID: lbID,
		},
		PoolCommon: common.PoolCommon{
			PoolID: poolID,
		},
	}
	return opts
}

type DeleteGlobalPoolRequest struct {
	common.LoadBalancerCommon
	common.PoolCommon
}

func (r *DeleteGlobalPoolRequest) WithLoadBalancerID(lbID string) *DeleteGlobalPoolRequest {
	r.LoadBalancerID = lbID
	return r
}

func (r *DeleteGlobalPoolRequest) WithPoolID(poolID string) *DeleteGlobalPoolRequest {
	r.PoolID = poolID
	return r
}

func NewDeleteGlobalPoolRequest(lbID, poolID string) *DeleteGlobalPoolRequest {
	opts := &DeleteGlobalPoolRequest{
		LoadBalancerCommon: common.LoadBalancerCommon{
			LoadBalancerID: lbID,
		},
		PoolCommon: common.PoolCommon{
			PoolID: poolID,
		},
	}
	return opts
}

type ListGlobalPoolMembersRequest struct {
	common.LoadBalancerCommon
	common.PoolCommon
}

func (r *ListGlobalPoolMembersRequest) WithLoadBalancerID(lbID string) *ListGlobalPoolMembersRequest {
	r.LoadBalancerID = lbID
	return r
}

func (r *ListGlobalPoolMembersRequest) WithPoolID(poolID string) *ListGlobalPoolMembersRequest {
	r.PoolID = poolID
	return r
}

func NewListGlobalPoolMembersRequest(lbID, poolID string) *ListGlobalPoolMembersRequest {
	opts := &ListGlobalPoolMembersRequest{
		LoadBalancerCommon: common.LoadBalancerCommon{
			LoadBalancerID: lbID,
		},
		PoolCommon: common.PoolCommon{
			PoolID: poolID,
		},
	}
	return opts
}

type PatchGlobalPoolMembersRequest struct {
	BulkActions []any `json:"bulkActions"`
	common.LoadBalancerCommon
	common.PoolCommon
}

func (r *PatchGlobalPoolMembersRequest) WithBulkAction(action ...any) *PatchGlobalPoolMembersRequest {
	r.BulkActions = action
	return r
}

func (r *PatchGlobalPoolMembersRequest) WithLoadBalancerID(lbID string) *PatchGlobalPoolMembersRequest {
	r.LoadBalancerID = lbID
	return r
}

func (r *PatchGlobalPoolMembersRequest) WithPoolID(poolID string) *PatchGlobalPoolMembersRequest {
	r.PoolID = poolID
	return r
}

func NewPatchGlobalPoolMembersRequest(lbID, poolID string) *PatchGlobalPoolMembersRequest {
	opts := &PatchGlobalPoolMembersRequest{
		BulkActions: make([]any, 0),
		LoadBalancerCommon: common.LoadBalancerCommon{
			LoadBalancerID: lbID,
		},
		PoolCommon: common.PoolCommon{
			PoolID: poolID,
		},
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

	common.LoadBalancerCommon
	common.PoolCommon
	common.PoolMemberCommon
}

func (r *UpdateGlobalPoolMemberRequest) WithTrafficDial(dial int) *UpdateGlobalPoolMemberRequest {
	r.TrafficDial = dial
	return r
}

func (r *UpdateGlobalPoolMemberRequest) WithMembers(members ...*GlobalMemberRequest) *UpdateGlobalPoolMemberRequest {
	r.Members = append(r.Members, members...)
	return r
}

func (r *UpdateGlobalPoolMemberRequest) WithLoadBalancerID(lbID string) *UpdateGlobalPoolMemberRequest {
	r.LoadBalancerID = lbID
	return r
}

func (r *UpdateGlobalPoolMemberRequest) WithPoolID(poolID string) *UpdateGlobalPoolMemberRequest {
	r.PoolID = poolID
	return r
}

func (r *UpdateGlobalPoolMemberRequest) WithPoolMemberID(poolMemberID string) *UpdateGlobalPoolMemberRequest {
	r.PoolMemberID = poolMemberID
	return r
}

func NewUpdateGlobalPoolMemberRequest(lbID, poolID, poolMemberID string, dial int) *UpdateGlobalPoolMemberRequest {
	opts := &UpdateGlobalPoolMemberRequest{
		TrafficDial: dial,
		Members:     make([]*GlobalMemberRequest, 0),
		LoadBalancerCommon: common.LoadBalancerCommon{
			LoadBalancerID: lbID,
		},
		PoolCommon: common.PoolCommon{
			PoolID: poolID,
		},
		PoolMemberCommon: common.PoolMemberCommon{
			PoolMemberID: poolMemberID,
		},
	}
	return opts
}

type DeleteGlobalPoolMemberRequest struct {
	common.LoadBalancerCommon
	common.PoolCommon
	common.PoolMemberCommon
}

func (r *DeleteGlobalPoolMemberRequest) WithLoadBalancerID(lbID string) *DeleteGlobalPoolMemberRequest {
	r.LoadBalancerID = lbID
	return r
}

func (r *DeleteGlobalPoolMemberRequest) WithPoolID(poolID string) *DeleteGlobalPoolMemberRequest {
	r.PoolID = poolID
	return r
}

func (r *DeleteGlobalPoolMemberRequest) WithPoolMemberID(poolMemberID string) *DeleteGlobalPoolMemberRequest {
	r.PoolMemberID = poolMemberID
	return r
}

func NewDeleteGlobalPoolMemberRequest(lbID, poolID, poolMemberID string) *DeleteGlobalPoolMemberRequest {
	opts := &DeleteGlobalPoolMemberRequest{
		LoadBalancerCommon: common.LoadBalancerCommon{
			LoadBalancerID: lbID,
		},
		PoolCommon: common.PoolCommon{
			PoolID: poolID,
		},
		PoolMemberCommon: common.PoolMemberCommon{
			PoolMemberID: poolMemberID,
		},
	}
	return opts
}

type GetGlobalPoolMemberRequest struct {
	common.LoadBalancerCommon
	common.PoolCommon
	common.PoolMemberCommon
}

func (r *GetGlobalPoolMemberRequest) WithLoadBalancerID(lbID string) *GetGlobalPoolMemberRequest {
	r.LoadBalancerID = lbID
	return r
}

func (r *GetGlobalPoolMemberRequest) WithPoolID(poolID string) *GetGlobalPoolMemberRequest {
	r.PoolID = poolID
	return r
}

func (r *GetGlobalPoolMemberRequest) WithPoolMemberID(poolMemberID string) *GetGlobalPoolMemberRequest {
	r.PoolMemberID = poolMemberID
	return r
}

func NewGetGlobalPoolMemberRequest(lbID, poolID, poolMemberID string) *GetGlobalPoolMemberRequest {
	opts := &GetGlobalPoolMemberRequest{
		LoadBalancerCommon: common.LoadBalancerCommon{
			LoadBalancerID: lbID,
		},
		PoolCommon: common.PoolCommon{
			PoolID: poolID,
		},
		PoolMemberCommon: common.PoolMemberCommon{
			PoolMemberID: poolMemberID,
		},
	}
	return opts
}
