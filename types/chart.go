package types

type ChartGetRequest struct {
	ID         int    `json:"id"`
	Asset      string `json:"asset"`
	Resolution string `json:"resolution"`
	Start      uint   `json:"start"`
	End        uint   `json:"end"`
}

type ChartGetResponse struct {
	T int64 `json:"time"`
	V any   `json:"value"`
}
