package schema

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Endpoint struct {
	gorm.Model
	Path          string         `gorm:"unique;not null" json:"path"`
	Tier          int            `gorm:"not null" json:"tier"`
	Assets        datatypes.JSON `gorm:"not null" json:"assets"`
	Currencies    datatypes.JSON `gorm:"not null" json:"currencies"`
	MiddleGroupID uint
	Name          string         `gorm:"unique" json:"name"`
	Description   string         `gorm:"unique" json:"description"`
	Resolutions   datatypes.JSON `gorm:"not null" json:"resolutions"`
	Formats       datatypes.JSON `gorm:"not null" json:"formats"`
	Initialized   bool           `gorm:"default:false" json:"initialized"`
}
