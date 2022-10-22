package types

import (
	"gorm.io/datatypes"
)

type EndpointGetAll struct {
	ID          uint           `json:"id"`
	Path        string         `json:"path"`
	Tier        int            `json:"tier"`
	Assets      datatypes.JSON `json:"assets"`
	Currencies  datatypes.JSON `json:"currencies"`
	Resolutions datatypes.JSON `json:"resolutions"`
	Formats     datatypes.JSON `json:"formats"`
}
