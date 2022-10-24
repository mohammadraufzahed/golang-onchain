package schema

import "gorm.io/gorm"

type TopGroup struct {
	gorm.Model
	Name         string `gorm:"not null;unique"`
	MiddleGroups []MiddleGroup
}

type MiddleGroup struct {
	gorm.Model
	Name        string `gorm:"not null;unique"`
	TopGroupID  uint
	ChildGroups []ChildGroup
}

type ChildGroup struct {
	gorm.Model
	Name          string `gorm:"not null;unique"`
	MiddleGroupID uint
	Description   string
	Endpoints     []Endpoint `gorm:"many2many:relative_endpoints;"`
}
