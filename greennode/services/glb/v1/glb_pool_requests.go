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

// --------------------------------------------------------------------------

var _ IListGlobalPoolsRequest = &ListGlobalPoolsRequest{}

type ListGlobalPoolsRequest struct {
	common.UserAgent
	common.LoadBalancerCommon
}

func (s *ListGlobalPoolsRequest) WithLoadBalancerID(lbID string) IListGlobalPoolsRequest {
	s.LoadBalancerID = lbID
	return s
}

func (s *ListGlobalPoolsRequest) AddUserAgent(agent ...string) IListGlobalPoolsRequest {
	s.UserAgent.AddUserAgent(agent...)
	return s
}

func NewListGlobalPoolsRequest(lbID string) IListGlobalPoolsRequest {
	opts := &ListGlobalPoolsRequest{}
	opts.LoadBalancerID = lbID
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

func (s *CreateGlobalPoolRequest) WithLoadBalancerID(lbID string) ICreateGlobalPoolRequest {
	s.LoadBalancerID = lbID
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
	HTTPMethod          *GlobalPoolHealthCheckMethod      `json:"httpMethod,omitempty"`
	HTTPVersion         *GlobalPoolHealthCheckHTTPVersion `json:"httpVersion,omitempty"`
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
	s.HTTPMethod = method
	return s
}

func (s *GlobalHealthMonitorRequest) WithHTTPVersion(version *GlobalPoolHealthCheckHTTPVersion) IGlobalHealthMonitorRequest {
	s.HTTPVersion = version
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
		"httpMethod":         s.HTTPMethod,
		"httpVersion":        s.HTTPVersion,
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

func (s *GlobalPoolMemberRequest) WithVPCID(vpcID string) ICreateGlobalPoolMemberRequest {
	s.VPCID = vpcID
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

func (s *GlobalPoolMemberRequest) WithLoadBalancerID(lbID string) ICreateGlobalPoolMemberRequest {
	s.LoadBalancerID = lbID
	return s
}

func (s *GlobalPoolMemberRequest) WithPoolID(poolID string) ICreateGlobalPoolMemberRequest {
	s.PoolID = poolID
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

func NewGlobalPoolMemberRequest(name, region, vpcID string, dial int, typeVal GlobalPoolMemberType) ICreateGlobalPoolMemberRequest {
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

func (s *GlobalMemberRequest) WithSubnetID(subnetID string) IGlobalMemberRequest {
	s.SubnetID = subnetID
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

func NewGlobalMemberRequest(name, address, subnetID string, port, monitorPort, weight int, backupRole bool) IGlobalMemberRequest {
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

func (s *UpdateGlobalPoolRequest) WithLoadBalancerID(lbID string) IUpdateGlobalPoolRequest {
	s.LoadBalancerID = lbID
	return s
}

func (s *UpdateGlobalPoolRequest) WithPoolID(poolID string) IUpdateGlobalPoolRequest {
	s.PoolID = poolID
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

func NewUpdateGlobalPoolRequest(lbID, poolID string) IUpdateGlobalPoolRequest {
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

// --------------------------------------------------------------------------

var _ IDeleteGlobalPoolRequest = &DeleteGlobalPoolRequest{}

type DeleteGlobalPoolRequest struct {
	common.LoadBalancerCommon
	common.PoolCommon
	common.UserAgent
}

func (s *DeleteGlobalPoolRequest) WithLoadBalancerID(lbID string) IDeleteGlobalPoolRequest {
	s.LoadBalancerID = lbID
	return s
}

func (s *DeleteGlobalPoolRequest) WithPoolID(poolID string) IDeleteGlobalPoolRequest {
	s.PoolID = poolID
	return s
}

func (s *DeleteGlobalPoolRequest) AddUserAgent(agent ...string) IDeleteGlobalPoolRequest {
	s.UserAgent.AddUserAgent(agent...)
	return s
}

func NewDeleteGlobalPoolRequest(lbID, poolID string) IDeleteGlobalPoolRequest {
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

// --------------------------------------------------------------------------

var _ IListGlobalPoolMembersRequest = &ListGlobalPoolMembersRequest{}

type ListGlobalPoolMembersRequest struct {
	common.LoadBalancerCommon
	common.PoolCommon
	common.UserAgent
}

func (s *ListGlobalPoolMembersRequest) WithLoadBalancerID(lbID string) IListGlobalPoolMembersRequest {
	s.LoadBalancerID = lbID
	return s
}

func (s *ListGlobalPoolMembersRequest) WithPoolID(poolID string) IListGlobalPoolMembersRequest {
	s.PoolID = poolID
	return s
}

func (s *ListGlobalPoolMembersRequest) AddUserAgent(agent ...string) IListGlobalPoolMembersRequest {
	s.UserAgent.AddUserAgent(agent...)
	return s
}

func NewListGlobalPoolMembersRequest(lbID, poolID string) IListGlobalPoolMembersRequest {
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

func (s *PatchGlobalPoolMembersRequest) WithLoadBalancerID(lbID string) IPatchGlobalPoolMembersRequest {
	s.LoadBalancerID = lbID
	return s
}

func (s *PatchGlobalPoolMembersRequest) WithPoolID(poolID string) IPatchGlobalPoolMembersRequest {
	s.PoolID = poolID
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

func NewPatchGlobalPoolMembersRequest(lbID, poolID string) IPatchGlobalPoolMembersRequest {
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

func (s *UpdateGlobalPoolMemberRequest) WithLoadBalancerID(lbID string) IUpdateGlobalPoolMemberRequest {
	s.LoadBalancerID = lbID
	return s
}

func (s *UpdateGlobalPoolMemberRequest) WithPoolID(poolID string) IUpdateGlobalPoolMemberRequest {
	s.PoolID = poolID
	return s
}

func (s *UpdateGlobalPoolMemberRequest) WithPoolMemberID(poolMemberID string) IUpdateGlobalPoolMemberRequest {
	s.PoolMemberID = poolMemberID
	return s
}

func (s *UpdateGlobalPoolMemberRequest) AddUserAgent(agent ...string) IUpdateGlobalPoolMemberRequest {
	s.UserAgent.AddUserAgent(agent...)
	return s
}

func (s *UpdateGlobalPoolMemberRequest) ToRequestBody() interface{} {
	return s
}

func NewUpdateGlobalPoolMemberRequest(lbID, poolID, poolMemberID string, dial int) IUpdateGlobalPoolMemberRequest {
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

// --------------------------------------------------------------------------

var _ IDeleteGlobalPoolMemberRequest = &DeleteGlobalPoolMemberRequest{}

type DeleteGlobalPoolMemberRequest struct {
	common.UserAgent
	common.LoadBalancerCommon
	common.PoolCommon
	common.PoolMemberCommon
}

func (s *DeleteGlobalPoolMemberRequest) WithLoadBalancerID(lbID string) IDeleteGlobalPoolMemberRequest {
	s.LoadBalancerID = lbID
	return s
}

func (s *DeleteGlobalPoolMemberRequest) WithPoolID(poolID string) IDeleteGlobalPoolMemberRequest {
	s.PoolID = poolID
	return s
}

func (s *DeleteGlobalPoolMemberRequest) WithPoolMemberID(poolMemberID string) IDeleteGlobalPoolMemberRequest {
	s.PoolMemberID = poolMemberID
	return s
}

func (s *DeleteGlobalPoolMemberRequest) AddUserAgent(agent ...string) IDeleteGlobalPoolMemberRequest {
	s.UserAgent.AddUserAgent(agent...)
	return s
}

func NewDeleteGlobalPoolMemberRequest(lbID, poolID, poolMemberID string) IDeleteGlobalPoolMemberRequest {
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

// --------------------------------------------------------------------------

var _ IGetGlobalPoolMemberRequest = &GetGlobalPoolMemberRequest{}

type GetGlobalPoolMemberRequest struct {
	common.UserAgent
	common.LoadBalancerCommon
	common.PoolCommon
	common.PoolMemberCommon
}

func (s *GetGlobalPoolMemberRequest) WithLoadBalancerID(lbID string) IGetGlobalPoolMemberRequest {
	s.LoadBalancerID = lbID
	return s
}

func (s *GetGlobalPoolMemberRequest) WithPoolID(poolID string) IGetGlobalPoolMemberRequest {
	s.PoolID = poolID
	return s
}

func (s *GetGlobalPoolMemberRequest) WithPoolMemberID(poolMemberID string) IGetGlobalPoolMemberRequest {
	s.PoolMemberID = poolMemberID
	return s
}

func (s *GetGlobalPoolMemberRequest) AddUserAgent(agent ...string) IGetGlobalPoolMemberRequest {
	s.UserAgent.AddUserAgent(agent...)
	return s
}

func NewGetGlobalPoolMemberRequest(lbID, poolID, poolMemberID string) IGetGlobalPoolMemberRequest {
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
