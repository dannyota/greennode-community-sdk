package inter

import "github.com/dannyota/greennode-community-sdk/v2/greennode/services/common"

const (
	InterVpcLoadBalancerScheme LoadBalancerScheme = "InterVPC"
	InternalLoadBalancerScheme LoadBalancerScheme = "Internal"
	InternetLoadBalancerScheme LoadBalancerScheme = "Internet"
)

const (
	CreateOptsTypeOptLayer4 LoadBalancerType = "Layer 4"
	CreateOptsTypeOptLayer7 LoadBalancerType = "Layer 7"
)

type (
	LoadBalancerScheme string
	LoadBalancerType   string
)

type CreateLoadBalancerRequest struct {
	Name            string                 `json:"name"`
	PackageID       string                 `json:"packageId"`
	Scheme          LoadBalancerScheme     `json:"scheme"`
	SubnetID        string                 `json:"subnetId"`
	BackEndSubnetId string                 `json:"backendSubnetId,omitempty"`
	ProjectId       string                 `json:"projectId,omitempty"`
	Type            LoadBalancerType       `json:"type"`
	Listener        ICreateListenerRequest `json:"listener,omitempty"`
	Pool            ICreatePoolRequest     `json:"pool,omitempty"`
	Tags            []common.Tag           `json:"tags,omitempty"`
	ZoneId          *common.Zone           `json:"zoneId,omitempty"`

	common.PortalUser
	common.UserAgent
}

func (s *CreateLoadBalancerRequest) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"name":            s.Name,
		"packageId":       s.PackageID,
		"scheme":          s.Scheme,
		"subnetId":        s.SubnetID,
		"backendSubnetId": s.BackEndSubnetId,
		"projectId":       s.ProjectId,
		"type":            s.Type,
		"tags":            s.Tags,
		"zoneId":          s.ZoneId,
	}
}

func (s *CreateLoadBalancerRequest) ToRequestBody() interface{} {
	if s.Pool != nil {
		s.Pool = s.Pool.ToRequestBody().(*CreatePoolRequest)
	}

	if s.Listener != nil {
		s.Listener = s.Listener.ToRequestBody().(*CreateListenerRequest)
	}

	return s
}

func (s *CreateLoadBalancerRequest) WithProjectId(projectId string) ICreateLoadBalancerRequest {
	s.ProjectId = projectId
	return s
}

func (s *CreateLoadBalancerRequest) AddUserAgent(agent ...string) ICreateLoadBalancerRequest {
	s.Agent = append(s.Agent, agent...)
	return s
}

func (s *CreateLoadBalancerRequest) GetMapHeaders() map[string]string {
	return s.PortalUser.GetMapHeaders()
}

func (s *CreateLoadBalancerRequest) WithListener(listener ICreateListenerRequest) ICreateLoadBalancerRequest {
	s.Listener = listener
	return s
}

func (s *CreateLoadBalancerRequest) WithPool(pool ICreatePoolRequest) ICreateLoadBalancerRequest {
	s.Pool = pool
	return s
}

func (s *CreateLoadBalancerRequest) WithTags(tags ...string) ICreateLoadBalancerRequest {
	if s.Tags == nil {
		s.Tags = make([]common.Tag, 0)
	}

	if len(tags)%2 != 0 {
		tags = append(tags, "none")
	}

	for i := 0; i < len(tags); i += 2 {
		s.Tags = append(s.Tags, common.Tag{Key: tags[i], Value: tags[i+1]})
	}

	return s
}

func (s *CreateLoadBalancerRequest) WithZoneId(zoneId common.Zone) ICreateLoadBalancerRequest {
	s.ZoneId = &zoneId
	return s
}
