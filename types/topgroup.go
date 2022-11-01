package types

import "github.com/ario-team/glassnode-api/schema"

// Top Group Request data
// @Description Required data to create the Top Group
type CreateTopGroupReq struct {
	// Name
	Name string `json:"name"`
}

// Top Group Response
// @Description Top Group response data
type CreateTopGroupRes struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type GetTopGroups struct {
	ID           uint          `json:"id"`
	Name         string        `json:"name"`
	MiddleGroups []MiddleGroup `json:"middle_groups"`
}

type MiddleGroup struct {
	ID          uint          `json:"id"`
	Name        string        `json:"name"`
	ChildGroups []ChildGroups `json:"child_groups"`
}

type ChildGroups struct {
	ID          uint            `json:"id"`
	Name        string          `json:"name"`
	Description string          `json:"description"`
	Endpoint    schema.Endpoint `json:"endpoint"`
	Initialized bool            `json:"initialized"`
}
