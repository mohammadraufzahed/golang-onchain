package types

type MiddleGroupCreate struct {
	Name       string `json:"name"`
	TopGroupID uint   `json:"topgroup_id"`
}

type MiddleGroupUpdate struct {
	Name string `json:"name"`
}
