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
	BackEndSubnetID string                 `json:"backendSubnetId,omitempty"`
	ProjectID       string                 `json:"projectId,omitempty"`
	Type            LoadBalancerType       `json:"type"`
	Listener        *CreateListenerRequest `json:"listener,omitempty"`
	Pool            *CreatePoolRequest     `json:"pool,omitempty"`
	Tags            []common.Tag           `json:"tags,omitempty"`
	ZoneID          *common.Zone           `json:"zoneId,omitempty"`

	common.PortalUser
	common.UserAgent
}

func (r *CreateLoadBalancerRequest) prepare() {
	if r.Pool != nil {
		r.Pool.prepare()
	}

	if r.Listener != nil {
		r.Listener.prepare()
	}
}

func (r *CreateLoadBalancerRequest) WithProjectID(projectID string) *CreateLoadBalancerRequest {
	r.ProjectID = projectID
	return r
}

func (r *CreateLoadBalancerRequest) AddUserAgent(agent ...string) *CreateLoadBalancerRequest {
	r.Agent = append(r.Agent, agent...)
	return r
}

func (r *CreateLoadBalancerRequest) GetMapHeaders() map[string]string {
	return r.PortalUser.GetMapHeaders()
}

func (r *CreateLoadBalancerRequest) WithListener(listener *CreateListenerRequest) *CreateLoadBalancerRequest {
	r.Listener = listener
	return r
}

func (r *CreateLoadBalancerRequest) WithPool(pool *CreatePoolRequest) *CreateLoadBalancerRequest {
	r.Pool = pool
	return r
}

func (r *CreateLoadBalancerRequest) WithTags(tags ...string) *CreateLoadBalancerRequest {
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

func (r *CreateLoadBalancerRequest) WithZoneID(zoneID common.Zone) *CreateLoadBalancerRequest {
	r.ZoneID = &zoneID
	return r
}
