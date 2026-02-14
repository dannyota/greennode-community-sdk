package v1

import "github.com/dannyota/greennode-community-sdk/v2/greennode/services/common"

type (
	GlobalPoolAlgorithm              string
	GlobalPoolProtocol               string
	GlobalPoolHealthCheckProtocol    string
	GlobalPoolHealthCheckMethod      string
	GlobalPoolHealthCheckHttpVersion string

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
	GlobalPoolHealthCheckHttpVersionHttp1       GlobalPoolHealthCheckHttpVersion = "1.0"
	GlobalPoolHealthCheckHttpVersionHttp1Minor1 GlobalPoolHealthCheckHttpVersion = "1.1"
)

const (
	GlobalPoolMemberTypePublic  GlobalPoolMemberType = "PUBLIC"
	GlobalPoolMemberTypePrivate GlobalPoolMemberType = "PRIVATE"
)

// --------------------------------------------------------------------------

var _ IListGlobalPoolsRequest = &ListGlobalPoolsRequest{}

type ListGlobalPoolsRequest struct {
	common.UserAgent
	common.LoadBalancerCommon
}

func (s *ListGlobalPoolsRequest) WithLoadBalancerId(lbId string) IListGlobalPoolsRequest {
	s.LoadBalancerId = lbId
	return s
}

func (s *ListGlobalPoolsRequest) AddUserAgent(agent ...string) IListGlobalPoolsRequest {
	s.UserAgent.AddUserAgent(agent...)
	return s
}

func NewListGlobalPoolsRequest(lbId string) IListGlobalPoolsRequest {
	opts := &ListGlobalPoolsRequest{}
	opts.LoadBalancerId = lbId
	return opts
}

// --------------------------------------------------------------------------

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

func (s *CreateGlobalPoolRequest) WithAlgorithm(algorithm GlobalPoolAlgorithm) ICreateGlobalPoolRequest {
	s.Algorithm = algorithm
	return s
}

func (s *CreateGlobalPoolRequest) WithDescription(description string) ICreateGlobalPoolRequest {
	s.Description = description
	return s
}

func (s *CreateGlobalPoolRequest) WithName(name string) ICreateGlobalPoolRequest {
	s.Name = name
	return s
}

func (s *CreateGlobalPoolRequest) WithProtocol(protocol GlobalPoolProtocol) ICreateGlobalPoolRequest {
	s.Protocol = protocol
	return s
}

func (s *CreateGlobalPoolRequest) WithHealthMonitor(health IGlobalHealthMonitorRequest) ICreateGlobalPoolRequest {
	s.HealthMonitor = health
	return s
}

func (s *CreateGlobalPoolRequest) WithMembers(members ...ICreateGlobalPoolMemberRequest) ICreateGlobalPoolRequest {
	s.GlobalPoolMembers = append(s.GlobalPoolMembers, members...)
	return s
}

func (s *CreateGlobalPoolRequest) WithLoadBalancerId(lbId string) ICreateGlobalPoolRequest {
	s.LoadBalancerId = lbId
	return s
}

func (s *CreateGlobalPoolRequest) ToMap() map[string]interface{} {
	err := map[string]interface{}{
		"algorithm":         s.Algorithm,
		"description":       s.Description,
		"name":              s.Name,
		"protocol":          s.Protocol,
		"stickiness":        s.Stickiness,
		"tlsEncryption":     s.TLSEncryption,
		"health":            s.HealthMonitor.ToMap(),
		"globalPoolMembers": make([]map[string]interface{}, 0),
	}

	for _, member := range s.GlobalPoolMembers {
		err["globalPoolMembers"] = append(err["globalPoolMembers"].([]map[string]interface{}), member.ToMap())
	}

	return err
}

func (s *CreateGlobalPoolRequest) ToRequestBody() interface{} {
	return s
}

func (s *CreateGlobalPoolRequest) AddUserAgent(agent ...string) ICreateGlobalPoolRequest {
	s.UserAgent.AddUserAgent(agent...)
	return s
}

func NewCreateGlobalPoolRequest(name string, protocol GlobalPoolProtocol) ICreateGlobalPoolRequest {
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

// --------------------------------------------------------------------------

var _ IGlobalHealthMonitorRequest = &GlobalHealthMonitorRequest{}

type GlobalHealthMonitorRequest struct {
	HealthCheckProtocol GlobalPoolHealthCheckProtocol     `json:"protocol"`
	HealthyThreshold    int                               `json:"healthyThreshold"`
	UnhealthyThreshold  int                               `json:"unhealthyThreshold"`
	Interval            int                               `json:"interval"`
	Timeout             int                               `json:"timeout"`
	HttpMethod          *GlobalPoolHealthCheckMethod      `json:"httpMethod,omitempty"`
	HttpVersion         *GlobalPoolHealthCheckHttpVersion `json:"httpVersion,omitempty"`
	Path                *string                           `json:"path,omitempty"`
	DomainName          *string                           `json:"domainName,omitempty"`
	SuccessCode         *string                           `json:"successCode,omitempty"`

	common.UserAgent
}

func (s *GlobalHealthMonitorRequest) WithHealthyThreshold(threshold int) IGlobalHealthMonitorRequest {
	s.HealthyThreshold = threshold
	return s
}

func (s *GlobalHealthMonitorRequest) WithUnhealthyThreshold(threshold int) IGlobalHealthMonitorRequest {
	s.UnhealthyThreshold = threshold
	return s
}

func (s *GlobalHealthMonitorRequest) WithProtocol(protocol GlobalPoolHealthCheckProtocol) IGlobalHealthMonitorRequest {
	s.HealthCheckProtocol = protocol
	return s
}

func (s *GlobalHealthMonitorRequest) WithInterval(interval int) IGlobalHealthMonitorRequest {
	s.Interval = interval
	return s
}

func (s *GlobalHealthMonitorRequest) WithTimeout(timeout int) IGlobalHealthMonitorRequest {
	s.Timeout = timeout
	return s
}

func (s *GlobalHealthMonitorRequest) WithHealthCheckMethod(method *GlobalPoolHealthCheckMethod) IGlobalHealthMonitorRequest {
	s.HttpMethod = method
	return s
}

func (s *GlobalHealthMonitorRequest) WithHttpVersion(version *GlobalPoolHealthCheckHttpVersion) IGlobalHealthMonitorRequest {
	s.HttpVersion = version
	return s
}

func (s *GlobalHealthMonitorRequest) WithDomainName(domain *string) IGlobalHealthMonitorRequest {
	s.DomainName = domain
	return s
}

func (s *GlobalHealthMonitorRequest) WithSuccessCode(code *string) IGlobalHealthMonitorRequest {
	s.SuccessCode = code
	return s
}

func (s *GlobalHealthMonitorRequest) WithPath(path *string) IGlobalHealthMonitorRequest {
	s.Path = path
	return s
}

func (s *GlobalHealthMonitorRequest) AddUserAgent(agent ...string) IGlobalHealthMonitorRequest {
	s.UserAgent.AddUserAgent(agent...)
	return s
}

func (s *GlobalHealthMonitorRequest) ToMap() map[string]interface{} {
	err := map[string]interface{}{
		"protocol":           s.HealthCheckProtocol,
		"healthyThreshold":   s.HealthyThreshold,
		"unhealthyThreshold": s.UnhealthyThreshold,
		"interval":           s.Interval,
		"timeout":            s.Timeout,
		"httpMethod":         s.HttpMethod,
		"httpVersion":        s.HttpVersion,
		"path":               s.Path,
		"domainName":         s.DomainName,
		"successCode":        s.SuccessCode,
	}

	return err
}

func (s *GlobalHealthMonitorRequest) ToRequestBody() interface{} {
	return s
}

func NewGlobalHealthMonitor(checkProtocol GlobalPoolHealthCheckProtocol) IGlobalHealthMonitorRequest {
	opts := &GlobalHealthMonitorRequest{
		HealthCheckProtocol: checkProtocol,
		HealthyThreshold:    3,
		UnhealthyThreshold:  3,
		Interval:            30,
		Timeout:             5,
		HttpMethod:          nil,
		HttpVersion:         nil,
		Path:                nil,
		DomainName:          nil,
		SuccessCode:         nil,
	}
	if checkProtocol == GlobalPoolHealthCheckProtocolHTTP || checkProtocol == GlobalPoolHealthCheckProtocolHTTPs {
		opts.HttpMethod = common.Ptr(GlobalPoolHealthCheckMethodGET)
		opts.HttpVersion = common.Ptr(GlobalPoolHealthCheckHttpVersionHttp1Minor1)
		opts.Path = common.Ptr("/")
		opts.DomainName = common.Ptr("")
		opts.SuccessCode = common.Ptr("200")
	}
	return opts
}

// --------------------------------------------------------------------------

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

func (s *GlobalPoolMemberRequest) WithName(name string) ICreateGlobalPoolMemberRequest {
	s.Name = name
	return s
}

func (s *GlobalPoolMemberRequest) WithDescription(description string) ICreateGlobalPoolMemberRequest {
	s.Description = description
	return s
}

func (s *GlobalPoolMemberRequest) WithRegion(region string) ICreateGlobalPoolMemberRequest {
	s.Region = region
	return s
}

func (s *GlobalPoolMemberRequest) WithTrafficDial(dial int) ICreateGlobalPoolMemberRequest {
	s.TrafficDial = dial
	return s
}

func (s *GlobalPoolMemberRequest) WithVPCID(vpcId string) ICreateGlobalPoolMemberRequest {
	s.VPCID = vpcId
	return s
}

func (s *GlobalPoolMemberRequest) WithType(typeVal GlobalPoolMemberType) ICreateGlobalPoolMemberRequest {
	s.Type = typeVal
	return s
}

func (s *GlobalPoolMemberRequest) WithMembers(members ...IGlobalMemberRequest) ICreateGlobalPoolMemberRequest {
	s.Members = append(s.Members, members...)
	return s
}

func (s *GlobalPoolMemberRequest) WithLoadBalancerId(lbId string) ICreateGlobalPoolMemberRequest {
	s.LoadBalancerId = lbId
	return s
}

func (s *GlobalPoolMemberRequest) WithPoolId(poolId string) ICreateGlobalPoolMemberRequest {
	s.PoolId = poolId
	return s
}

func (s *GlobalPoolMemberRequest) AddUserAgent(agent ...string) ICreateGlobalPoolMemberRequest {
	s.UserAgent.AddUserAgent(agent...)
	return s
}

func (s *GlobalPoolMemberRequest) ToMap() map[string]interface{} {
	err := map[string]interface{}{
		"name":        s.Name,
		"description": s.Description,
		"region":      s.Region,
		"trafficDial": s.TrafficDial,
		"vpcId":       s.VPCID,
		"members":     make([]map[string]interface{}, 0),
	}

	for _, member := range s.Members {
		err["members"] = append(err["members"].([]map[string]interface{}), member.ToMap())
	}

	return err
}

func (s *GlobalPoolMemberRequest) ToRequestBody() interface{} {
	return s
}

func NewGlobalPoolMemberRequest(name, region, vpcId string, dial int, typeVal GlobalPoolMemberType) ICreateGlobalPoolMemberRequest {
	opts := &GlobalPoolMemberRequest{
		Name:        name,
		Description: "",
		Region:      region,
		TrafficDial: dial,
		Type:        typeVal,
		VPCID:       vpcId,
		Members:     make([]IGlobalMemberRequest, 0),
	}
	return opts
}

// --------------------------------------------------------------------------

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

func (s *GlobalMemberRequest) WithAddress(address string) IGlobalMemberRequest {
	s.Address = address
	return s
}

func (s *GlobalMemberRequest) WithBackupRole(backupRole bool) IGlobalMemberRequest {
	s.BackupRole = backupRole
	return s
}

func (s *GlobalMemberRequest) WithDescription(description string) IGlobalMemberRequest {
	s.Description = description
	return s
}

func (s *GlobalMemberRequest) WithMonitorPort(monitorPort int) IGlobalMemberRequest {
	s.MonitorPort = monitorPort
	return s
}

func (s *GlobalMemberRequest) WithName(name string) IGlobalMemberRequest {
	s.Name = name
	return s
}

func (s *GlobalMemberRequest) WithPort(port int) IGlobalMemberRequest {
	s.Port = port
	return s
}

func (s *GlobalMemberRequest) WithSubnetID(subnetId string) IGlobalMemberRequest {
	s.SubnetID = subnetId
	return s
}

func (s *GlobalMemberRequest) WithWeight(weight int) IGlobalMemberRequest {
	s.Weight = weight
	return s
}

func (s *GlobalMemberRequest) AddUserAgent(agent ...string) IGlobalMemberRequest {
	s.UserAgent.AddUserAgent(agent...)
	return s
}

func (s *GlobalMemberRequest) ToMap() map[string]interface{} {
	err := map[string]interface{}{
		"address":     s.Address,
		"backupRole":  s.BackupRole,
		"description": s.Description,
		"monitorPort": s.MonitorPort,
		"name":        s.Name,
		"port":        s.Port,
		"subnetId":    s.SubnetID,
		"weight":      s.Weight,
	}
	return err
}

func (s *GlobalMemberRequest) ToRequestBody() interface{} {
	return s
}

func NewGlobalMemberRequest(name, address, subnetId string, port, monitorPort, weight int, backupRole bool) IGlobalMemberRequest {
	opts := &GlobalMemberRequest{
		Name:        name,
		Address:     address,
		BackupRole:  backupRole,
		Description: "",
		MonitorPort: monitorPort,
		Port:        port,
		SubnetID:    subnetId,
		Weight:      weight,
	}
	return opts
}

// --------------------------------------------------------------------------

var _ IUpdateGlobalPoolRequest = &UpdateGlobalPoolRequest{}

type UpdateGlobalPoolRequest struct {
	Algorithm     GlobalPoolAlgorithm         `json:"algorithm"`
	HealthMonitor IGlobalHealthMonitorRequest `json:"health"`
	common.LoadBalancerCommon
	common.PoolCommon
	common.UserAgent
}

func (s *UpdateGlobalPoolRequest) WithAlgorithm(algorithm GlobalPoolAlgorithm) IUpdateGlobalPoolRequest {
	s.Algorithm = algorithm
	return s
}

func (s *UpdateGlobalPoolRequest) WithHealthMonitor(monitor IGlobalHealthMonitorRequest) IUpdateGlobalPoolRequest {
	s.HealthMonitor = monitor
	return s
}

func (s *UpdateGlobalPoolRequest) WithLoadBalancerId(lbId string) IUpdateGlobalPoolRequest {
	s.LoadBalancerId = lbId
	return s
}

func (s *UpdateGlobalPoolRequest) WithPoolId(poolId string) IUpdateGlobalPoolRequest {
	s.PoolId = poolId
	return s
}

func (s *UpdateGlobalPoolRequest) AddUserAgent(agent ...string) IUpdateGlobalPoolRequest {
	s.UserAgent.AddUserAgent(agent...)
	return s
}

func (s *UpdateGlobalPoolRequest) ToMap() map[string]interface{} {
	err := map[string]interface{}{
		"algorithm": s.Algorithm,
		"health":    s.HealthMonitor.ToMap(),
	}
	return err
}

func (s *UpdateGlobalPoolRequest) ToRequestBody() interface{} {
	return s
}

func NewUpdateGlobalPoolRequest(lbId, poolId string) IUpdateGlobalPoolRequest {
	opts := &UpdateGlobalPoolRequest{
		Algorithm:     GlobalPoolAlgorithmRoundRobin,
		HealthMonitor: nil,
		LoadBalancerCommon: common.LoadBalancerCommon{
			LoadBalancerId: lbId,
		},
		PoolCommon: common.PoolCommon{
			PoolId: poolId,
		},
	}
	return opts
}

// --------------------------------------------------------------------------

var _ IDeleteGlobalPoolRequest = &DeleteGlobalPoolRequest{}

type DeleteGlobalPoolRequest struct {
	common.LoadBalancerCommon
	common.PoolCommon
	common.UserAgent
}

func (s *DeleteGlobalPoolRequest) WithLoadBalancerId(lbId string) IDeleteGlobalPoolRequest {
	s.LoadBalancerId = lbId
	return s
}

func (s *DeleteGlobalPoolRequest) WithPoolId(poolId string) IDeleteGlobalPoolRequest {
	s.PoolId = poolId
	return s
}

func (s *DeleteGlobalPoolRequest) AddUserAgent(agent ...string) IDeleteGlobalPoolRequest {
	s.UserAgent.AddUserAgent(agent...)
	return s
}

func NewDeleteGlobalPoolRequest(lbId, poolId string) IDeleteGlobalPoolRequest {
	opts := &DeleteGlobalPoolRequest{
		LoadBalancerCommon: common.LoadBalancerCommon{
			LoadBalancerId: lbId,
		},
		PoolCommon: common.PoolCommon{
			PoolId: poolId,
		},
	}
	return opts
}

// --------------------------------------------------------------------------

var _ IListGlobalPoolMembersRequest = &ListGlobalPoolMembersRequest{}

type ListGlobalPoolMembersRequest struct {
	common.LoadBalancerCommon
	common.PoolCommon
	common.UserAgent
}

func (s *ListGlobalPoolMembersRequest) WithLoadBalancerId(lbId string) IListGlobalPoolMembersRequest {
	s.LoadBalancerId = lbId
	return s
}

func (s *ListGlobalPoolMembersRequest) WithPoolId(poolId string) IListGlobalPoolMembersRequest {
	s.PoolId = poolId
	return s
}

func (s *ListGlobalPoolMembersRequest) AddUserAgent(agent ...string) IListGlobalPoolMembersRequest {
	s.UserAgent.AddUserAgent(agent...)
	return s
}

func NewListGlobalPoolMembersRequest(lbId, poolId string) IListGlobalPoolMembersRequest {
	opts := &ListGlobalPoolMembersRequest{
		LoadBalancerCommon: common.LoadBalancerCommon{
			LoadBalancerId: lbId,
		},
		PoolCommon: common.PoolCommon{
			PoolId: poolId,
		},
	}
	return opts
}

// --------------------------------------------------------------------------

var _ IPatchGlobalPoolMembersRequest = &PatchGlobalPoolMembersRequest{}

type PatchGlobalPoolMembersRequest struct {
	BulkActions []IBulkActionRequest `json:"bulkActions"`
	common.LoadBalancerCommon
	common.PoolCommon
	common.UserAgent
}

func (s *PatchGlobalPoolMembersRequest) WithBulkAction(action ...IBulkActionRequest) IPatchGlobalPoolMembersRequest {
	s.BulkActions = action
	return s
}

func (s *PatchGlobalPoolMembersRequest) WithLoadBalancerId(lbId string) IPatchGlobalPoolMembersRequest {
	s.LoadBalancerId = lbId
	return s
}

func (s *PatchGlobalPoolMembersRequest) WithPoolId(poolId string) IPatchGlobalPoolMembersRequest {
	s.PoolId = poolId
	return s
}

func (s *PatchGlobalPoolMembersRequest) AddUserAgent(agent ...string) IPatchGlobalPoolMembersRequest {
	s.UserAgent.AddUserAgent(agent...)
	return s
}

func (s *PatchGlobalPoolMembersRequest) ToMap() map[string]interface{} {
	err := map[string]interface{}{
		"bulkActions": make([]map[string]interface{}, 0),
	}

	for _, action := range s.BulkActions {
		err["bulkActions"] = append(err["bulkActions"].([]map[string]interface{}), action.ToMap())
	}

	return err
}

func (s *PatchGlobalPoolMembersRequest) ToRequestBody() interface{} {
	return s
}

func NewPatchGlobalPoolMembersRequest(lbId, poolId string) IPatchGlobalPoolMembersRequest {
	opts := &PatchGlobalPoolMembersRequest{
		BulkActions: make([]IBulkActionRequest, 0),
		LoadBalancerCommon: common.LoadBalancerCommon{
			LoadBalancerId: lbId,
		},
		PoolCommon: common.PoolCommon{
			PoolId: poolId,
		},
	}
	return opts
}

var _ IBulkActionRequest = &PatchGlobalPoolCreateBulkActionRequest{}

type PatchGlobalPoolCreateBulkActionRequest struct {
	Action           string                         `json:"action"`
	CreatePoolMember ICreateGlobalPoolMemberRequest `json:"createPoolMember"`
}

func (s *PatchGlobalPoolCreateBulkActionRequest) ToMap() map[string]interface{} {
	err := map[string]interface{}{
		"action":           s.Action,
		"createPoolMember": s.CreatePoolMember.ToMap(),
	}
	return err
}

func (s *PatchGlobalPoolCreateBulkActionRequest) ToRequestBody() interface{} {
	return s
}

func NewPatchGlobalPoolCreateBulkActionRequest(member ICreateGlobalPoolMemberRequest) IBulkActionRequest {
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

func (s *PatchGlobalPoolDeleteBulkActionRequest) ToMap() map[string]interface{} {
	err := map[string]interface{}{
		"action": s.Action,
		"id":     s.ID,
	}
	return err
}

func (s *PatchGlobalPoolDeleteBulkActionRequest) ToRequestBody() interface{} {
	return s
}

func NewPatchGlobalPoolDeleteBulkActionRequest(id string) IBulkActionRequest {
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

func (s *PatchGlobalPoolUpdateBulkActionRequest) ToMap() map[string]interface{} {
	err := map[string]interface{}{
		"action":           s.Action,
		"id":               s.ID,
		"updatePoolMember": s.UpdatePoolMember.ToMap(),
	}
	return err
}

func (s *PatchGlobalPoolUpdateBulkActionRequest) ToRequestBody() interface{} {
	return s
}

func NewPatchGlobalPoolUpdateBulkActionRequest(id string, member IUpdateGlobalPoolMemberRequest) IBulkActionRequest {
	opts := &PatchGlobalPoolUpdateBulkActionRequest{
		Action:           "update",
		ID:               id,
		UpdatePoolMember: member,
	}
	return opts
}

type IUpdateGlobalPoolMemberRequest interface {
	WithLoadBalancerId(lbId string) IUpdateGlobalPoolMemberRequest
	WithPoolId(poolId string) IUpdateGlobalPoolMemberRequest
	WithPoolMemberId(poolMemberId string) IUpdateGlobalPoolMemberRequest
	WithTrafficDial(dial int) IUpdateGlobalPoolMemberRequest
	WithMembers(members ...IGlobalMemberRequest) IUpdateGlobalPoolMemberRequest
	GetLoadBalancerId() string // to use in request url
	GetPoolId() string
	GetPoolMemberId() string

	AddUserAgent(agent ...string) IUpdateGlobalPoolMemberRequest
	ParseUserAgent() string
	ToRequestBody() interface{}
	ToMap() map[string]interface{}
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

func (s *UpdateGlobalPoolMemberRequest) WithTrafficDial(dial int) IUpdateGlobalPoolMemberRequest {
	s.TrafficDial = dial
	return s
}

func (s *UpdateGlobalPoolMemberRequest) WithMembers(members ...IGlobalMemberRequest) IUpdateGlobalPoolMemberRequest {
	s.Members = append(s.Members, members...)
	return s
}

func (s *UpdateGlobalPoolMemberRequest) ToMap() map[string]interface{} {
	err := map[string]interface{}{
		"trafficDial": s.TrafficDial,
		"members":     make([]map[string]interface{}, 0),
	}

	for _, member := range s.Members {
		err["members"] = append(err["members"].([]map[string]interface{}), member.ToMap())
	}

	return err
}

func (s *UpdateGlobalPoolMemberRequest) WithLoadBalancerId(lbId string) IUpdateGlobalPoolMemberRequest {
	s.LoadBalancerId = lbId
	return s
}

func (s *UpdateGlobalPoolMemberRequest) WithPoolId(poolId string) IUpdateGlobalPoolMemberRequest {
	s.PoolId = poolId
	return s
}

func (s *UpdateGlobalPoolMemberRequest) WithPoolMemberId(poolMemberId string) IUpdateGlobalPoolMemberRequest {
	s.PoolMemberId = poolMemberId
	return s
}

func (s *UpdateGlobalPoolMemberRequest) AddUserAgent(agent ...string) IUpdateGlobalPoolMemberRequest {
	s.UserAgent.AddUserAgent(agent...)
	return s
}

func (s *UpdateGlobalPoolMemberRequest) ToRequestBody() interface{} {
	return s
}

func NewUpdateGlobalPoolMemberRequest(lbId, poolId, poolMemberId string, dial int) IUpdateGlobalPoolMemberRequest {
	opts := &UpdateGlobalPoolMemberRequest{
		TrafficDial: dial,
		Members:     make([]IGlobalMemberRequest, 0),
		LoadBalancerCommon: common.LoadBalancerCommon{
			LoadBalancerId: lbId,
		},
		PoolCommon: common.PoolCommon{
			PoolId: poolId,
		},
		PoolMemberCommon: common.PoolMemberCommon{
			PoolMemberId: poolMemberId,
		},
	}
	return opts
}

// --------------------------------------------------------------------------

var _ IDeleteGlobalPoolMemberRequest = &DeleteGlobalPoolMemberRequest{}

type DeleteGlobalPoolMemberRequest struct {
	common.UserAgent
	common.LoadBalancerCommon
	common.PoolCommon
	common.PoolMemberCommon
}

func (s *DeleteGlobalPoolMemberRequest) WithLoadBalancerId(lbId string) IDeleteGlobalPoolMemberRequest {
	s.LoadBalancerId = lbId
	return s
}

func (s *DeleteGlobalPoolMemberRequest) WithPoolId(poolId string) IDeleteGlobalPoolMemberRequest {
	s.PoolId = poolId
	return s
}

func (s *DeleteGlobalPoolMemberRequest) WithPoolMemberId(poolMemberId string) IDeleteGlobalPoolMemberRequest {
	s.PoolMemberId = poolMemberId
	return s
}

func (s *DeleteGlobalPoolMemberRequest) AddUserAgent(agent ...string) IDeleteGlobalPoolMemberRequest {
	s.UserAgent.AddUserAgent(agent...)
	return s
}

func NewDeleteGlobalPoolMemberRequest(lbId, poolId, poolMemberId string) IDeleteGlobalPoolMemberRequest {
	opts := &DeleteGlobalPoolMemberRequest{
		LoadBalancerCommon: common.LoadBalancerCommon{
			LoadBalancerId: lbId,
		},
		PoolCommon: common.PoolCommon{
			PoolId: poolId,
		},
		PoolMemberCommon: common.PoolMemberCommon{
			PoolMemberId: poolMemberId,
		},
	}
	return opts
}

// --------------------------------------------------------------------------

var _ IGetGlobalPoolMemberRequest = &GetGlobalPoolMemberRequest{}

type GetGlobalPoolMemberRequest struct {
	common.UserAgent
	common.LoadBalancerCommon
	common.PoolCommon
	common.PoolMemberCommon
}

func (s *GetGlobalPoolMemberRequest) WithLoadBalancerId(lbId string) IGetGlobalPoolMemberRequest {
	s.LoadBalancerId = lbId
	return s
}

func (s *GetGlobalPoolMemberRequest) WithPoolId(poolId string) IGetGlobalPoolMemberRequest {
	s.PoolId = poolId
	return s
}

func (s *GetGlobalPoolMemberRequest) WithPoolMemberId(poolMemberId string) IGetGlobalPoolMemberRequest {
	s.PoolMemberId = poolMemberId
	return s
}

func (s *GetGlobalPoolMemberRequest) AddUserAgent(agent ...string) IGetGlobalPoolMemberRequest {
	s.UserAgent.AddUserAgent(agent...)
	return s
}

func NewGetGlobalPoolMemberRequest(lbId, poolId, poolMemberId string) IGetGlobalPoolMemberRequest {
	opts := &GetGlobalPoolMemberRequest{
		LoadBalancerCommon: common.LoadBalancerCommon{
			LoadBalancerId: lbId,
		},
		PoolCommon: common.PoolCommon{
			PoolId: poolId,
		},
		PoolMemberCommon: common.PoolMemberCommon{
			PoolMemberId: poolMemberId,
		},
	}
	return opts
}
