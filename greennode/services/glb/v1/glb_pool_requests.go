package v1

import "github.com/dannyota/greennode-community-sdk/v2/greennode/services/common"

type IListGlobalPoolsRequest interface {
	WithLoadBalancerID(lbID string) IListGlobalPoolsRequest
	GetLoadBalancerID() string

	AddUserAgent(agent ...string) IListGlobalPoolsRequest
	ParseUserAgent() string
}

type ICreateGlobalPoolRequest interface {
	WithAlgorithm(algorithm GlobalPoolAlgorithm) ICreateGlobalPoolRequest
	WithDescription(desc string) ICreateGlobalPoolRequest
	WithName(name string) ICreateGlobalPoolRequest
	WithProtocol(protocol GlobalPoolProtocol) ICreateGlobalPoolRequest
	WithHealthMonitor(monitor IGlobalHealthMonitorRequest) ICreateGlobalPoolRequest
	WithMembers(members ...ICreateGlobalPoolMemberRequest) ICreateGlobalPoolRequest

	WithLoadBalancerID(lbID string) ICreateGlobalPoolRequest
	GetLoadBalancerID() string

	AddUserAgent(agent ...string) ICreateGlobalPoolRequest
	ParseUserAgent() string
	ToRequestBody() any
	ToMap() map[string]any
}

type IGlobalHealthMonitorRequest interface {
	WithHealthyThreshold(ht int) IGlobalHealthMonitorRequest
	WithInterval(interval int) IGlobalHealthMonitorRequest
	WithProtocol(protocol GlobalPoolHealthCheckProtocol) IGlobalHealthMonitorRequest
	WithTimeout(to int) IGlobalHealthMonitorRequest
	WithUnhealthyThreshold(uht int) IGlobalHealthMonitorRequest

	WithHealthCheckMethod(method *GlobalPoolHealthCheckMethod) IGlobalHealthMonitorRequest
	WithHTTPVersion(version *GlobalPoolHealthCheckHTTPVersion) IGlobalHealthMonitorRequest
	WithPath(path *string) IGlobalHealthMonitorRequest
	WithSuccessCode(code *string) IGlobalHealthMonitorRequest
	WithDomainName(domain *string) IGlobalHealthMonitorRequest

	AddUserAgent(agent ...string) IGlobalHealthMonitorRequest
	ParseUserAgent() string
	ToRequestBody() any
	ToMap() map[string]any
}

type ICreateGlobalPoolMemberRequest interface {
	WithName(name string) ICreateGlobalPoolMemberRequest
	WithDescription(desc string) ICreateGlobalPoolMemberRequest
	WithRegion(region string) ICreateGlobalPoolMemberRequest
	WithVPCID(vpcID string) ICreateGlobalPoolMemberRequest
	WithTrafficDial(dial int) ICreateGlobalPoolMemberRequest
	WithMembers(members ...IGlobalMemberRequest) ICreateGlobalPoolMemberRequest
	WithType(typeVal GlobalPoolMemberType) ICreateGlobalPoolMemberRequest

	WithLoadBalancerID(lbID string) ICreateGlobalPoolMemberRequest
	WithPoolID(poolID string) ICreateGlobalPoolMemberRequest
	GetLoadBalancerID() string
	GetPoolID() string

	AddUserAgent(agent ...string) ICreateGlobalPoolMemberRequest
	ParseUserAgent() string
	ToRequestBody() any
	ToMap() map[string]any
}

type IGlobalMemberRequest interface {
	WithAddress(addr string) IGlobalMemberRequest
	WithBackupRole(backup bool) IGlobalMemberRequest
	WithDescription(desc string) IGlobalMemberRequest
	WithMonitorPort(port int) IGlobalMemberRequest
	WithName(name string) IGlobalMemberRequest
	WithPort(port int) IGlobalMemberRequest
	WithSubnetID(subnetID string) IGlobalMemberRequest
	WithWeight(weight int) IGlobalMemberRequest

	AddUserAgent(agent ...string) IGlobalMemberRequest
	ParseUserAgent() string
	ToRequestBody() any
	ToMap() map[string]any
}

type IUpdateGlobalPoolRequest interface {
	WithAlgorithm(algorithm GlobalPoolAlgorithm) IUpdateGlobalPoolRequest
	WithHealthMonitor(monitor IGlobalHealthMonitorRequest) IUpdateGlobalPoolRequest

	WithLoadBalancerID(lbID string) IUpdateGlobalPoolRequest
	WithPoolID(poolID string) IUpdateGlobalPoolRequest
	GetLoadBalancerID() string
	GetPoolID() string

	AddUserAgent(agent ...string) IUpdateGlobalPoolRequest
	ParseUserAgent() string
	ToRequestBody() any
	ToMap() map[string]any
}

type IDeleteGlobalPoolRequest interface {
	WithLoadBalancerID(lbID string) IDeleteGlobalPoolRequest
	WithPoolID(poolID string) IDeleteGlobalPoolRequest
	GetLoadBalancerID() string
	GetPoolID() string

	AddUserAgent(agent ...string) IDeleteGlobalPoolRequest
	ParseUserAgent() string
}

type IListGlobalPoolMembersRequest interface {
	WithLoadBalancerID(lbID string) IListGlobalPoolMembersRequest
	WithPoolID(poolID string) IListGlobalPoolMembersRequest
	GetLoadBalancerID() string
	GetPoolID() string

	AddUserAgent(agent ...string) IListGlobalPoolMembersRequest
	ParseUserAgent() string
}

type IGetGlobalPoolMemberRequest interface {
	WithLoadBalancerID(lbID string) IGetGlobalPoolMemberRequest
	WithPoolID(poolID string) IGetGlobalPoolMemberRequest
	WithPoolMemberID(poolMemberID string) IGetGlobalPoolMemberRequest
	GetLoadBalancerID() string
	GetPoolID() string
	GetPoolMemberID() string

	AddUserAgent(agent ...string) IGetGlobalPoolMemberRequest
	ParseUserAgent() string
}

type IDeleteGlobalPoolMemberRequest interface {
	WithLoadBalancerID(lbID string) IDeleteGlobalPoolMemberRequest
	WithPoolID(poolID string) IDeleteGlobalPoolMemberRequest
	WithPoolMemberID(poolMemberID string) IDeleteGlobalPoolMemberRequest
	GetLoadBalancerID() string
	GetPoolID() string
	GetPoolMemberID() string

	AddUserAgent(agent ...string) IDeleteGlobalPoolMemberRequest
	ParseUserAgent() string
}

type IPatchGlobalPoolMembersRequest interface {
	WithBulkAction(action ...IBulkActionRequest) IPatchGlobalPoolMembersRequest

	WithLoadBalancerID(lbID string) IPatchGlobalPoolMembersRequest
	WithPoolID(poolID string) IPatchGlobalPoolMembersRequest
	GetLoadBalancerID() string
	GetPoolID() string

	AddUserAgent(agent ...string) IPatchGlobalPoolMembersRequest
	ParseUserAgent() string
	ToRequestBody() any
	ToMap() map[string]any
}

type IBulkActionRequest interface {
	ToRequestBody() any
	ToMap() map[string]any
}

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
	GlobalPoolHealthCheckProtocolHTTPs GlobalPoolHealthCheckProtocol = "HTTPS"
)

const (
	GlobalPoolHealthCheckMethodGET  GlobalPoolHealthCheckMethod = "GET"
	GlobalPoolHealthCheckMethodPUT  GlobalPoolHealthCheckMethod = "PUT"
	GlobalPoolHealthCheckMethodPOST GlobalPoolHealthCheckMethod = "POST"
)

const (
	GlobalPoolHealthCheckHTTPVersionHttp1       GlobalPoolHealthCheckHTTPVersion = "1.0"
	GlobalPoolHealthCheckHTTPVersionHttp1Minor1 GlobalPoolHealthCheckHTTPVersion = "1.1"
)

const (
	GlobalPoolMemberTypePublic  GlobalPoolMemberType = "PUBLIC"
	GlobalPoolMemberTypePrivate GlobalPoolMemberType = "PRIVATE"
)


var _ IListGlobalPoolsRequest = &ListGlobalPoolsRequest{}

type ListGlobalPoolsRequest struct {
	common.UserAgent
	common.LoadBalancerCommon
}

func (r *ListGlobalPoolsRequest) WithLoadBalancerID(lbID string) IListGlobalPoolsRequest {
	r.LoadBalancerID = lbID
	return r
}

func (r *ListGlobalPoolsRequest) AddUserAgent(agent ...string) IListGlobalPoolsRequest {
	r.UserAgent.AddUserAgent(agent...)
	return r
}

func NewListGlobalPoolsRequest(lbID string) *ListGlobalPoolsRequest {
	opts := &ListGlobalPoolsRequest{}
	opts.LoadBalancerID = lbID
	return opts
}


var _ ICreateGlobalPoolRequest = &CreateGlobalPoolRequest{}

type CreateGlobalPoolRequest struct {
	Algorithm         GlobalPoolAlgorithm              `json:"algorithm"`
	Description       string                           `json:"description,omitempty"`
	Name              string                           `json:"name"`
	Protocol          GlobalPoolProtocol               `json:"protocol"`
	Stickiness        *bool                            `json:"stickiness,omitempty"`    // only for l7, l4 doesn't have this field => nil
	TLSEncryption     *bool                            `json:"tlsEncryption,omitempty"` // only for l7, l4 doesn't have this field => nil
	HealthMonitor     IGlobalHealthMonitorRequest      `json:"health"`
	GlobalPoolMembers []ICreateGlobalPoolMemberRequest `json:"globalPoolMembers"`

	common.LoadBalancerCommon
	common.UserAgent
}

func (r *CreateGlobalPoolRequest) WithAlgorithm(algorithm GlobalPoolAlgorithm) ICreateGlobalPoolRequest {
	r.Algorithm = algorithm
	return r
}

func (r *CreateGlobalPoolRequest) WithDescription(description string) ICreateGlobalPoolRequest {
	r.Description = description
	return r
}

func (r *CreateGlobalPoolRequest) WithName(name string) ICreateGlobalPoolRequest {
	r.Name = name
	return r
}

func (r *CreateGlobalPoolRequest) WithProtocol(protocol GlobalPoolProtocol) ICreateGlobalPoolRequest {
	r.Protocol = protocol
	return r
}

func (r *CreateGlobalPoolRequest) WithHealthMonitor(health IGlobalHealthMonitorRequest) ICreateGlobalPoolRequest {
	r.HealthMonitor = health
	return r
}

func (r *CreateGlobalPoolRequest) WithMembers(members ...ICreateGlobalPoolMemberRequest) ICreateGlobalPoolRequest {
	r.GlobalPoolMembers = append(r.GlobalPoolMembers, members...)
	return r
}

func (r *CreateGlobalPoolRequest) WithLoadBalancerID(lbID string) ICreateGlobalPoolRequest {
	r.LoadBalancerID = lbID
	return r
}

func (r *CreateGlobalPoolRequest) ToMap() map[string]any {
	err := map[string]any{
		"algorithm":         r.Algorithm,
		"description":       r.Description,
		"name":              r.Name,
		"protocol":          r.Protocol,
		"stickiness":        r.Stickiness,
		"tlsEncryption":     r.TLSEncryption,
		"health":            r.HealthMonitor.ToMap(),
		"globalPoolMembers": make([]map[string]any, 0),
	}

	for _, member := range r.GlobalPoolMembers {
		err["globalPoolMembers"] = append(err["globalPoolMembers"].([]map[string]any), member.ToMap())
	}

	return err
}

func (r *CreateGlobalPoolRequest) ToRequestBody() any {
	return r
}

func (r *CreateGlobalPoolRequest) AddUserAgent(agent ...string) ICreateGlobalPoolRequest {
	r.UserAgent.AddUserAgent(agent...)
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
		GlobalPoolMembers: make([]ICreateGlobalPoolMemberRequest, 0),
	}

	return opts
}


var _ IGlobalHealthMonitorRequest = &GlobalHealthMonitorRequest{}

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

	common.UserAgent
}

func (r *GlobalHealthMonitorRequest) WithHealthyThreshold(threshold int) IGlobalHealthMonitorRequest {
	r.HealthyThreshold = threshold
	return r
}

func (r *GlobalHealthMonitorRequest) WithUnhealthyThreshold(threshold int) IGlobalHealthMonitorRequest {
	r.UnhealthyThreshold = threshold
	return r
}

func (r *GlobalHealthMonitorRequest) WithProtocol(protocol GlobalPoolHealthCheckProtocol) IGlobalHealthMonitorRequest {
	r.HealthCheckProtocol = protocol
	return r
}

func (r *GlobalHealthMonitorRequest) WithInterval(interval int) IGlobalHealthMonitorRequest {
	r.Interval = interval
	return r
}

func (r *GlobalHealthMonitorRequest) WithTimeout(timeout int) IGlobalHealthMonitorRequest {
	r.Timeout = timeout
	return r
}

func (r *GlobalHealthMonitorRequest) WithHealthCheckMethod(method *GlobalPoolHealthCheckMethod) IGlobalHealthMonitorRequest {
	r.HTTPMethod = method
	return r
}

func (r *GlobalHealthMonitorRequest) WithHTTPVersion(version *GlobalPoolHealthCheckHTTPVersion) IGlobalHealthMonitorRequest {
	r.HTTPVersion = version
	return r
}

func (r *GlobalHealthMonitorRequest) WithDomainName(domain *string) IGlobalHealthMonitorRequest {
	r.DomainName = domain
	return r
}

func (r *GlobalHealthMonitorRequest) WithSuccessCode(code *string) IGlobalHealthMonitorRequest {
	r.SuccessCode = code
	return r
}

func (r *GlobalHealthMonitorRequest) WithPath(path *string) IGlobalHealthMonitorRequest {
	r.Path = path
	return r
}

func (r *GlobalHealthMonitorRequest) AddUserAgent(agent ...string) IGlobalHealthMonitorRequest {
	r.UserAgent.AddUserAgent(agent...)
	return r
}

func (r *GlobalHealthMonitorRequest) ToMap() map[string]any {
	err := map[string]any{
		"protocol":           r.HealthCheckProtocol,
		"healthyThreshold":   r.HealthyThreshold,
		"unhealthyThreshold": r.UnhealthyThreshold,
		"interval":           r.Interval,
		"timeout":            r.Timeout,
		"httpMethod":         r.HTTPMethod,
		"httpVersion":        r.HTTPVersion,
		"path":               r.Path,
		"domainName":         r.DomainName,
		"successCode":        r.SuccessCode,
	}

	return err
}

func (r *GlobalHealthMonitorRequest) ToRequestBody() any {
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
	if checkProtocol == GlobalPoolHealthCheckProtocolHTTP || checkProtocol == GlobalPoolHealthCheckProtocolHTTPs {
		opts.HTTPMethod = common.Ptr(GlobalPoolHealthCheckMethodGET)
		opts.HTTPVersion = common.Ptr(GlobalPoolHealthCheckHTTPVersionHttp1Minor1)
		opts.Path = common.Ptr("/")
		opts.DomainName = common.Ptr("")
		opts.SuccessCode = common.Ptr("200")
	}
	return opts
}


var _ ICreateGlobalPoolMemberRequest = &GlobalPoolMemberRequest{}

type GlobalPoolMemberRequest struct {
	Name        string                 `json:"name"`
	Description string                 `json:"description,omitempty"`
	Region      string                 `json:"region"`
	TrafficDial int                    `json:"trafficDial"`
	VPCID       string                 `json:"vpcId"` // only need for private type
	Type        GlobalPoolMemberType   `json:"type"`
	Members     []IGlobalMemberRequest `json:"members"`

	common.LoadBalancerCommon
	common.PoolCommon
	common.UserAgent
}

func (r *GlobalPoolMemberRequest) WithName(name string) ICreateGlobalPoolMemberRequest {
	r.Name = name
	return r
}

func (r *GlobalPoolMemberRequest) WithDescription(description string) ICreateGlobalPoolMemberRequest {
	r.Description = description
	return r
}

func (r *GlobalPoolMemberRequest) WithRegion(region string) ICreateGlobalPoolMemberRequest {
	r.Region = region
	return r
}

func (r *GlobalPoolMemberRequest) WithTrafficDial(dial int) ICreateGlobalPoolMemberRequest {
	r.TrafficDial = dial
	return r
}

func (r *GlobalPoolMemberRequest) WithVPCID(vpcID string) ICreateGlobalPoolMemberRequest {
	r.VPCID = vpcID
	return r
}

func (r *GlobalPoolMemberRequest) WithType(typeVal GlobalPoolMemberType) ICreateGlobalPoolMemberRequest {
	r.Type = typeVal
	return r
}

func (r *GlobalPoolMemberRequest) WithMembers(members ...IGlobalMemberRequest) ICreateGlobalPoolMemberRequest {
	r.Members = append(r.Members, members...)
	return r
}

func (r *GlobalPoolMemberRequest) WithLoadBalancerID(lbID string) ICreateGlobalPoolMemberRequest {
	r.LoadBalancerID = lbID
	return r
}

func (r *GlobalPoolMemberRequest) WithPoolID(poolID string) ICreateGlobalPoolMemberRequest {
	r.PoolID = poolID
	return r
}

func (r *GlobalPoolMemberRequest) AddUserAgent(agent ...string) ICreateGlobalPoolMemberRequest {
	r.UserAgent.AddUserAgent(agent...)
	return r
}

func (r *GlobalPoolMemberRequest) ToMap() map[string]any {
	err := map[string]any{
		"name":        r.Name,
		"description": r.Description,
		"region":      r.Region,
		"trafficDial": r.TrafficDial,
		"vpcId":       r.VPCID,
		"members":     make([]map[string]any, 0),
	}

	for _, member := range r.Members {
		err["members"] = append(err["members"].([]map[string]any), member.ToMap())
	}

	return err
}

func (r *GlobalPoolMemberRequest) ToRequestBody() any {
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
		Members:     make([]IGlobalMemberRequest, 0),
	}
	return opts
}


var _ IGlobalMemberRequest = &GlobalMemberRequest{}

type GlobalMemberRequest struct {
	Address     string `json:"address"`
	BackupRole  bool   `json:"backupRole"`
	Description string `json:"description,omitempty"`
	MonitorPort int    `json:"monitorPort"`
	Name        string `json:"name"`
	Port        int    `json:"port"`
	SubnetID    string `json:"subnetId"`
	Weight      int    `json:"weight"`

	common.UserAgent
}

func (r *GlobalMemberRequest) WithAddress(address string) IGlobalMemberRequest {
	r.Address = address
	return r
}

func (r *GlobalMemberRequest) WithBackupRole(backupRole bool) IGlobalMemberRequest {
	r.BackupRole = backupRole
	return r
}

func (r *GlobalMemberRequest) WithDescription(description string) IGlobalMemberRequest {
	r.Description = description
	return r
}

func (r *GlobalMemberRequest) WithMonitorPort(monitorPort int) IGlobalMemberRequest {
	r.MonitorPort = monitorPort
	return r
}

func (r *GlobalMemberRequest) WithName(name string) IGlobalMemberRequest {
	r.Name = name
	return r
}

func (r *GlobalMemberRequest) WithPort(port int) IGlobalMemberRequest {
	r.Port = port
	return r
}

func (r *GlobalMemberRequest) WithSubnetID(subnetID string) IGlobalMemberRequest {
	r.SubnetID = subnetID
	return r
}

func (r *GlobalMemberRequest) WithWeight(weight int) IGlobalMemberRequest {
	r.Weight = weight
	return r
}

func (r *GlobalMemberRequest) AddUserAgent(agent ...string) IGlobalMemberRequest {
	r.UserAgent.AddUserAgent(agent...)
	return r
}

func (r *GlobalMemberRequest) ToMap() map[string]any {
	err := map[string]any{
		"address":     r.Address,
		"backupRole":  r.BackupRole,
		"description": r.Description,
		"monitorPort": r.MonitorPort,
		"name":        r.Name,
		"port":        r.Port,
		"subnetId":    r.SubnetID,
		"weight":      r.Weight,
	}
	return err
}

func (r *GlobalMemberRequest) ToRequestBody() any {
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


var _ IUpdateGlobalPoolRequest = &UpdateGlobalPoolRequest{}

type UpdateGlobalPoolRequest struct {
	Algorithm     GlobalPoolAlgorithm         `json:"algorithm"`
	HealthMonitor IGlobalHealthMonitorRequest `json:"health"`
	common.LoadBalancerCommon
	common.PoolCommon
	common.UserAgent
}

func (r *UpdateGlobalPoolRequest) WithAlgorithm(algorithm GlobalPoolAlgorithm) IUpdateGlobalPoolRequest {
	r.Algorithm = algorithm
	return r
}

func (r *UpdateGlobalPoolRequest) WithHealthMonitor(monitor IGlobalHealthMonitorRequest) IUpdateGlobalPoolRequest {
	r.HealthMonitor = monitor
	return r
}

func (r *UpdateGlobalPoolRequest) WithLoadBalancerID(lbID string) IUpdateGlobalPoolRequest {
	r.LoadBalancerID = lbID
	return r
}

func (r *UpdateGlobalPoolRequest) WithPoolID(poolID string) IUpdateGlobalPoolRequest {
	r.PoolID = poolID
	return r
}

func (r *UpdateGlobalPoolRequest) AddUserAgent(agent ...string) IUpdateGlobalPoolRequest {
	r.UserAgent.AddUserAgent(agent...)
	return r
}

func (r *UpdateGlobalPoolRequest) ToMap() map[string]any {
	err := map[string]any{
		"algorithm": r.Algorithm,
		"health":    r.HealthMonitor.ToMap(),
	}
	return err
}

func (r *UpdateGlobalPoolRequest) ToRequestBody() any {
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


var _ IDeleteGlobalPoolRequest = &DeleteGlobalPoolRequest{}

type DeleteGlobalPoolRequest struct {
	common.LoadBalancerCommon
	common.PoolCommon
	common.UserAgent
}

func (r *DeleteGlobalPoolRequest) WithLoadBalancerID(lbID string) IDeleteGlobalPoolRequest {
	r.LoadBalancerID = lbID
	return r
}

func (r *DeleteGlobalPoolRequest) WithPoolID(poolID string) IDeleteGlobalPoolRequest {
	r.PoolID = poolID
	return r
}

func (r *DeleteGlobalPoolRequest) AddUserAgent(agent ...string) IDeleteGlobalPoolRequest {
	r.UserAgent.AddUserAgent(agent...)
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


var _ IListGlobalPoolMembersRequest = &ListGlobalPoolMembersRequest{}

type ListGlobalPoolMembersRequest struct {
	common.LoadBalancerCommon
	common.PoolCommon
	common.UserAgent
}

func (r *ListGlobalPoolMembersRequest) WithLoadBalancerID(lbID string) IListGlobalPoolMembersRequest {
	r.LoadBalancerID = lbID
	return r
}

func (r *ListGlobalPoolMembersRequest) WithPoolID(poolID string) IListGlobalPoolMembersRequest {
	r.PoolID = poolID
	return r
}

func (r *ListGlobalPoolMembersRequest) AddUserAgent(agent ...string) IListGlobalPoolMembersRequest {
	r.UserAgent.AddUserAgent(agent...)
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


var _ IPatchGlobalPoolMembersRequest = &PatchGlobalPoolMembersRequest{}

type PatchGlobalPoolMembersRequest struct {
	BulkActions []IBulkActionRequest `json:"bulkActions"`
	common.LoadBalancerCommon
	common.PoolCommon
	common.UserAgent
}

func (r *PatchGlobalPoolMembersRequest) WithBulkAction(action ...IBulkActionRequest) IPatchGlobalPoolMembersRequest {
	r.BulkActions = action
	return r
}

func (r *PatchGlobalPoolMembersRequest) WithLoadBalancerID(lbID string) IPatchGlobalPoolMembersRequest {
	r.LoadBalancerID = lbID
	return r
}

func (r *PatchGlobalPoolMembersRequest) WithPoolID(poolID string) IPatchGlobalPoolMembersRequest {
	r.PoolID = poolID
	return r
}

func (r *PatchGlobalPoolMembersRequest) AddUserAgent(agent ...string) IPatchGlobalPoolMembersRequest {
	r.UserAgent.AddUserAgent(agent...)
	return r
}

func (r *PatchGlobalPoolMembersRequest) ToMap() map[string]any {
	err := map[string]any{
		"bulkActions": make([]map[string]any, 0),
	}

	for _, action := range r.BulkActions {
		err["bulkActions"] = append(err["bulkActions"].([]map[string]any), action.ToMap())
	}

	return err
}

func (r *PatchGlobalPoolMembersRequest) ToRequestBody() any {
	return r
}

func NewPatchGlobalPoolMembersRequest(lbID, poolID string) *PatchGlobalPoolMembersRequest {
	opts := &PatchGlobalPoolMembersRequest{
		BulkActions: make([]IBulkActionRequest, 0),
		LoadBalancerCommon: common.LoadBalancerCommon{
			LoadBalancerID: lbID,
		},
		PoolCommon: common.PoolCommon{
			PoolID: poolID,
		},
	}
	return opts
}

var _ IBulkActionRequest = &PatchGlobalPoolCreateBulkActionRequest{}

type PatchGlobalPoolCreateBulkActionRequest struct {
	Action           string                         `json:"action"`
	CreatePoolMember ICreateGlobalPoolMemberRequest `json:"createPoolMember"`
}

func (r *PatchGlobalPoolCreateBulkActionRequest) ToMap() map[string]any {
	err := map[string]any{
		"action":           r.Action,
		"createPoolMember": r.CreatePoolMember.ToMap(),
	}
	return err
}

func (r *PatchGlobalPoolCreateBulkActionRequest) ToRequestBody() any {
	return r
}

func NewPatchGlobalPoolCreateBulkActionRequest(member ICreateGlobalPoolMemberRequest) *PatchGlobalPoolCreateBulkActionRequest {
	opts := &PatchGlobalPoolCreateBulkActionRequest{
		Action:           "create",
		CreatePoolMember: member,
	}
	return opts
}

var _ IBulkActionRequest = &PatchGlobalPoolDeleteBulkActionRequest{}

type PatchGlobalPoolDeleteBulkActionRequest struct {
	Action string `json:"action"`
	ID     string `json:"id"`
}

func (r *PatchGlobalPoolDeleteBulkActionRequest) ToMap() map[string]any {
	err := map[string]any{
		"action": r.Action,
		"id":     r.ID,
	}
	return err
}

func (r *PatchGlobalPoolDeleteBulkActionRequest) ToRequestBody() any {
	return r
}

func NewPatchGlobalPoolDeleteBulkActionRequest(id string) *PatchGlobalPoolDeleteBulkActionRequest {
	opts := &PatchGlobalPoolDeleteBulkActionRequest{
		Action: "delete",
		ID:     id,
	}
	return opts
}

var _ IBulkActionRequest = &PatchGlobalPoolUpdateBulkActionRequest{}

type PatchGlobalPoolUpdateBulkActionRequest struct {
	Action           string                         `json:"action"`
	ID               string                         `json:"id"`
	UpdatePoolMember IUpdateGlobalPoolMemberRequest `json:"updatePoolMember"`
}

func (r *PatchGlobalPoolUpdateBulkActionRequest) ToMap() map[string]any {
	err := map[string]any{
		"action":           r.Action,
		"id":               r.ID,
		"updatePoolMember": r.UpdatePoolMember.ToMap(),
	}
	return err
}

func (r *PatchGlobalPoolUpdateBulkActionRequest) ToRequestBody() any {
	return r
}

func NewPatchGlobalPoolUpdateBulkActionRequest(id string, member IUpdateGlobalPoolMemberRequest) *PatchGlobalPoolUpdateBulkActionRequest {
	opts := &PatchGlobalPoolUpdateBulkActionRequest{
		Action:           "update",
		ID:               id,
		UpdatePoolMember: member,
	}
	return opts
}

type IUpdateGlobalPoolMemberRequest interface {
	WithLoadBalancerID(lbID string) IUpdateGlobalPoolMemberRequest
	WithPoolID(poolID string) IUpdateGlobalPoolMemberRequest
	WithPoolMemberID(poolMemberID string) IUpdateGlobalPoolMemberRequest
	WithTrafficDial(dial int) IUpdateGlobalPoolMemberRequest
	WithMembers(members ...IGlobalMemberRequest) IUpdateGlobalPoolMemberRequest
	GetLoadBalancerID() string // to use in request url
	GetPoolID() string
	GetPoolMemberID() string

	AddUserAgent(agent ...string) IUpdateGlobalPoolMemberRequest
	ParseUserAgent() string
	ToRequestBody() any
	ToMap() map[string]any
}

var _ IUpdateGlobalPoolMemberRequest = &UpdateGlobalPoolMemberRequest{}

type UpdateGlobalPoolMemberRequest struct {
	TrafficDial int                    `json:"trafficDial"`
	Members     []IGlobalMemberRequest `json:"members"`

	common.UserAgent
	common.LoadBalancerCommon
	common.PoolCommon
	common.PoolMemberCommon
}

func (r *UpdateGlobalPoolMemberRequest) WithTrafficDial(dial int) IUpdateGlobalPoolMemberRequest {
	r.TrafficDial = dial
	return r
}

func (r *UpdateGlobalPoolMemberRequest) WithMembers(members ...IGlobalMemberRequest) IUpdateGlobalPoolMemberRequest {
	r.Members = append(r.Members, members...)
	return r
}

func (r *UpdateGlobalPoolMemberRequest) ToMap() map[string]any {
	err := map[string]any{
		"trafficDial": r.TrafficDial,
		"members":     make([]map[string]any, 0),
	}

	for _, member := range r.Members {
		err["members"] = append(err["members"].([]map[string]any), member.ToMap())
	}

	return err
}

func (r *UpdateGlobalPoolMemberRequest) WithLoadBalancerID(lbID string) IUpdateGlobalPoolMemberRequest {
	r.LoadBalancerID = lbID
	return r
}

func (r *UpdateGlobalPoolMemberRequest) WithPoolID(poolID string) IUpdateGlobalPoolMemberRequest {
	r.PoolID = poolID
	return r
}

func (r *UpdateGlobalPoolMemberRequest) WithPoolMemberID(poolMemberID string) IUpdateGlobalPoolMemberRequest {
	r.PoolMemberID = poolMemberID
	return r
}

func (r *UpdateGlobalPoolMemberRequest) AddUserAgent(agent ...string) IUpdateGlobalPoolMemberRequest {
	r.UserAgent.AddUserAgent(agent...)
	return r
}

func (r *UpdateGlobalPoolMemberRequest) ToRequestBody() any {
	return r
}

func NewUpdateGlobalPoolMemberRequest(lbID, poolID, poolMemberID string, dial int) *UpdateGlobalPoolMemberRequest {
	opts := &UpdateGlobalPoolMemberRequest{
		TrafficDial: dial,
		Members:     make([]IGlobalMemberRequest, 0),
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


var _ IDeleteGlobalPoolMemberRequest = &DeleteGlobalPoolMemberRequest{}

type DeleteGlobalPoolMemberRequest struct {
	common.UserAgent
	common.LoadBalancerCommon
	common.PoolCommon
	common.PoolMemberCommon
}

func (r *DeleteGlobalPoolMemberRequest) WithLoadBalancerID(lbID string) IDeleteGlobalPoolMemberRequest {
	r.LoadBalancerID = lbID
	return r
}

func (r *DeleteGlobalPoolMemberRequest) WithPoolID(poolID string) IDeleteGlobalPoolMemberRequest {
	r.PoolID = poolID
	return r
}

func (r *DeleteGlobalPoolMemberRequest) WithPoolMemberID(poolMemberID string) IDeleteGlobalPoolMemberRequest {
	r.PoolMemberID = poolMemberID
	return r
}

func (r *DeleteGlobalPoolMemberRequest) AddUserAgent(agent ...string) IDeleteGlobalPoolMemberRequest {
	r.UserAgent.AddUserAgent(agent...)
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


var _ IGetGlobalPoolMemberRequest = &GetGlobalPoolMemberRequest{}

type GetGlobalPoolMemberRequest struct {
	common.UserAgent
	common.LoadBalancerCommon
	common.PoolCommon
	common.PoolMemberCommon
}

func (r *GetGlobalPoolMemberRequest) WithLoadBalancerID(lbID string) IGetGlobalPoolMemberRequest {
	r.LoadBalancerID = lbID
	return r
}

func (r *GetGlobalPoolMemberRequest) WithPoolID(poolID string) IGetGlobalPoolMemberRequest {
	r.PoolID = poolID
	return r
}

func (r *GetGlobalPoolMemberRequest) WithPoolMemberID(poolMemberID string) IGetGlobalPoolMemberRequest {
	r.PoolMemberID = poolMemberID
	return r
}

func (r *GetGlobalPoolMemberRequest) AddUserAgent(agent ...string) IGetGlobalPoolMemberRequest {
	r.UserAgent.AddUserAgent(agent...)
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
