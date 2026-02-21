package v2

func NewGetNetworkByIDRequest(networkID string) *GetNetworkByIDRequest {
	return &GetNetworkByIDRequest{
		NetworkID: networkID,
	}
}

type GetNetworkByIDRequest struct {
	NetworkID string
}
