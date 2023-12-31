package types

type ChildGroupCreate struct {
	Name          string `json:"name"`
	Description   string `json:"description"`
	MiddleGroupID uint   `json:"middlegroup_id"`
	EndpointID    uint   `json:"endpoint_id"`
}

type ChildGroupUpdate struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}
