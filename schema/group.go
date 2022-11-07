package schema

import "gorm.io/gorm"

type TopGroup struct {
	gorm.Model
	Name         string `gorm:"not null;unique"`
	MiddleGroups []MiddleGroup
}

type MiddleGroup struct {
	gorm.Model
	Name       string `gorm:"not null;unique"`
	TopGroupID uint
	Endpoints  []Endpoint `gorm:"many2many:endpoints_middlegroup"`
}
