package v1

import "github.com/dannyota/greennode-community-sdk/v2/greennode/entity"

type GlobalListenerResponse struct {
	CreatedAt            string  `json:"createdAt"`
	UpdatedAt            string  `json:"updatedAt"`
	DeletedAt            *string `json:"deletedAt"`
	ID                   string  `json:"id"`
	Name                 string  `json:"name"`
	Description          string  `json:"description"`
	Protocol             string  `json:"protocol"`
	Port                 int     `json:"port"`
	GlobalLoadBalancerID string  `json:"globalLoadBalancerId"`
	GlobalPoolID         string  `json:"globalPoolId"`
	TimeoutClient        int     `json:"timeoutClient"`
	TimeoutMember        int     `json:"timeoutMember"`
	TimeoutConnection    int     `json:"timeoutConnection"`
	AllowedCidrs         string  `json:"allowedCidrs"`
	Headers              *string `json:"headers"`
	Status               string  `json:"status"`
}

func (r *GlobalListenerResponse) ToEntityGlobalListener() *entity.GlobalListener {
	return &entity.GlobalListener{
		CreatedAt:            r.CreatedAt,
		UpdatedAt:            r.UpdatedAt,
		DeletedAt:            r.DeletedAt,
		ID:                   r.ID,
		Name:                 r.Name,
		Description:          r.Description,
		Protocol:             r.Protocol,
		Port:                 r.Port,
		GlobalLoadBalancerID: r.GlobalLoadBalancerID,
		GlobalPoolID:         r.GlobalPoolID,
		TimeoutClient:        r.TimeoutClient,
		TimeoutMember:        r.TimeoutMember,
		TimeoutConnection:    r.TimeoutConnection,
		AllowedCidrs:         r.AllowedCidrs,
		Headers:              r.Headers,
		Status:               r.Status,
	}
}

type ListGlobalListenersResponse []GlobalListenerResponse

func (r ListGlobalListenersResponse) ToEntityListGlobalListeners() *entity.ListGlobalListeners {
	listeners := &entity.ListGlobalListeners{}
	for _, itemListener := range r {
		listeners.Items = append(listeners.Items, itemListener.ToEntityGlobalListener())
	}
	return listeners
}


type CreateGlobalListenerResponse GlobalListenerResponse

func (r *CreateGlobalListenerResponse) ToEntityGlobalListener() *entity.GlobalListener {
	return &entity.GlobalListener{
		CreatedAt:            r.CreatedAt,
		UpdatedAt:            r.UpdatedAt,
		DeletedAt:            r.DeletedAt,
		ID:                   r.ID,
		Name:                 r.Name,
		Description:          r.Description,
		Protocol:             r.Protocol,
		Port:                 r.Port,
		GlobalLoadBalancerID: r.GlobalLoadBalancerID,
		GlobalPoolID:         r.GlobalPoolID,
		TimeoutClient:        r.TimeoutClient,
		TimeoutMember:        r.TimeoutMember,
		TimeoutConnection:    r.TimeoutConnection,
		AllowedCidrs:         r.AllowedCidrs,
		Headers:              r.Headers,
		Status:               r.Status,
	}
}


type UpdateGlobalListenerResponse GlobalListenerResponse

func (r *UpdateGlobalListenerResponse) ToEntityGlobalListener() *entity.GlobalListener {
	return &entity.GlobalListener{
		CreatedAt:            r.CreatedAt,
		UpdatedAt:            r.UpdatedAt,
		DeletedAt:            r.DeletedAt,
		ID:                   r.ID,
		Name:                 r.Name,
		Description:          r.Description,
		Protocol:             r.Protocol,
		Port:                 r.Port,
		GlobalLoadBalancerID: r.GlobalLoadBalancerID,
		GlobalPoolID:         r.GlobalPoolID,
		TimeoutClient:        r.TimeoutClient,
		TimeoutMember:        r.TimeoutMember,
		TimeoutConnection:    r.TimeoutConnection,
		AllowedCidrs:         r.AllowedCidrs,
		Headers:              r.Headers,
		Status:               r.Status,
	}
}


type GetGlobalListenerResponse GlobalListenerResponse

func (r *GetGlobalListenerResponse) ToEntityGlobalListener() *entity.GlobalListener {
	return &entity.GlobalListener{
		CreatedAt:            r.CreatedAt,
		UpdatedAt:            r.UpdatedAt,
		DeletedAt:            r.DeletedAt,
		ID:                   r.ID,
		Name:                 r.Name,
		Description:          r.Description,
		Protocol:             r.Protocol,
		Port:                 r.Port,
		GlobalLoadBalancerID: r.GlobalLoadBalancerID,
		GlobalPoolID:         r.GlobalPoolID,
		TimeoutClient:        r.TimeoutClient,
		TimeoutMember:        r.TimeoutMember,
		TimeoutConnection:    r.TimeoutConnection,
		AllowedCidrs:         r.AllowedCidrs,
		Headers:              r.Headers,
		Status:               r.Status,
	}
}
