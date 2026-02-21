package v2

func NewListAllServersBySecgroupIDRequest(secgroupID string) *ListAllServersBySecgroupIDRequest {
	return &ListAllServersBySecgroupIDRequest{
		SecgroupID: secgroupID,
	}
}

type ListAllServersBySecgroupIDRequest struct {
	SecgroupID string
}
