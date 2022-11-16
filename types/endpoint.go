package types

import (
	"gorm.io/datatypes"
)

type EndpointGetAll struct {
	ID          uint           `json:"id"`
	Path        string         `json:"path"`
	Tier        int            `json:"tier"`
	Name        string         `json:"name"`
	Description string         `json:"description"`
	Initialized bool           `json:"initialized"`
	Assets      datatypes.JSON `json:"assets"`
	Currencies  datatypes.JSON `json:"currencies"`
	Resolutions datatypes.JSON `json:"resolutions"`
	Formats     datatypes.JSON `json:"formats"`
}

type EndpointUpdate struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type Assets struct {
	Symbol    string   `json:"symbol"`
	Name      string   `json:"name"`
	Tags      []string `json:"tags"`
	Exchanges []string `json:"exchanges"`
}

type ChartData struct {
	Time  int     `json:"t"`
	Value float32 `json:"v"`
}
