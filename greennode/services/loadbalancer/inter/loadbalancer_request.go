package inter

import "github.com/dannyota/greennode-community-sdk/v2/greennode/services/common"

type ICreateLoadBalancerRequest interface {
	ToRequestBody() any
	AddUserAgent(agent ...string) ICreateLoadBalancerRequest
	WithListener(listener ICreateListenerRequest) ICreateLoadBalancerRequest
	WithPool(pool ICreatePoolRequest) ICreateLoadBalancerRequest
	WithProjectID(projectID string) ICreateLoadBalancerRequest
	WithTags(tags ...string) ICreateLoadBalancerRequest
	WithZoneID(zoneID common.Zone) ICreateLoadBalancerRequest
	GetMapHeaders() map[string]string
	ParseUserAgent() string
	ToMap() map[string]any
}

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
	BackEndSubnetID string                 `json:"backendSubnetId,omitempty"`
	ProjectID       string                 `json:"projectId,omitempty"`
	Type            LoadBalancerType       `json:"type"`
	Listener        ICreateListenerRequest `json:"listener,omitempty"`
	Pool            ICreatePoolRequest     `json:"pool,omitempty"`
	Tags            []common.Tag           `json:"tags,omitempty"`
	ZoneID          *common.Zone           `json:"zoneId,omitempty"`

	common.PortalUser
	common.UserAgent
}

func (r *CreateLoadBalancerRequest) ToMap() map[string]any {
	return map[string]any{
		"name":            r.Name,
		"packageId":       r.PackageID,
		"scheme":          r.Scheme,
		"subnetId":        r.SubnetID,
		"backendSubnetId": r.BackEndSubnetID,
		"projectId":       r.ProjectID,
		"type":            r.Type,
		"tags":            r.Tags,
		"zoneId":          r.ZoneID,
	}
}

func (r *CreateLoadBalancerRequest) ToRequestBody() any {
	if r.Pool != nil {
		r.Pool = r.Pool.ToRequestBody().(*CreatePoolRequest)
	}

	if r.Listener != nil {
		r.Listener = r.Listener.ToRequestBody().(*CreateListenerRequest)
	}

	return r
}

func (r *CreateLoadBalancerRequest) WithProjectID(projectID string) ICreateLoadBalancerRequest {
	r.ProjectID = projectID
	return r
}

func (r *CreateLoadBalancerRequest) AddUserAgent(agent ...string) ICreateLoadBalancerRequest {
	r.Agent = append(r.Agent, agent...)
	return r
}

func (r *CreateLoadBalancerRequest) GetMapHeaders() map[string]string {
	return r.PortalUser.GetMapHeaders()
}

func (r *CreateLoadBalancerRequest) WithListener(listener ICreateListenerRequest) ICreateLoadBalancerRequest {
	r.Listener = listener
	return r
}

func (r *CreateLoadBalancerRequest) WithPool(pool ICreatePoolRequest) ICreateLoadBalancerRequest {
	r.Pool = pool
	return r
}

func (r *CreateLoadBalancerRequest) WithTags(tags ...string) ICreateLoadBalancerRequest {
	if r.Tags == nil {
		r.Tags = make([]common.Tag, 0)
	}

	if len(tags)%2 != 0 {
		tags = append(tags, "none")
	}

	for i := 0; i < len(tags); i += 2 {
		r.Tags = append(r.Tags, common.Tag{Key: tags[i], Value: tags[i+1]})
	}

	return r
}

func (r *CreateLoadBalancerRequest) WithZoneID(zoneID common.Zone) ICreateLoadBalancerRequest {
	r.ZoneID = &zoneID
	return r
}
