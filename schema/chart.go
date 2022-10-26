package schema

import "gorm.io/gorm"

type Chart struct {
	gorm.Model
	Name         string `gorm:"unique;not null"`
	ChildGroupID uint
	Endpoint     Endpoint
	EndpointID   uint
}
