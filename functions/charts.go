package functions

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/ario-team/glassnode-api/config"
	"github.com/ario-team/glassnode-api/database"
	"github.com/ario-team/glassnode-api/influxdb"
	"github.com/ario-team/glassnode-api/schema"
	"github.com/ario-team/glassnode-api/types"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)

func InitializeChart(endpointID uint, chartID uint) {
	var endpoint schema.Endpoint
	var childgroup schema.ChildGroup
	database.Connection.Where("id = ?", endpointID).First(&endpoint)
	database.Connection.Where("id = ?", chartID).First(&childgroup)
	var assets []types.Assets
	var resolutions []string
	apiKey := config.Viper.GetString("GLASSNODE_API_KEY")
	json.Unmarshal([]byte(endpoint.Assets), &assets)
	json.Unmarshal([]byte(endpoint.Resolutions), &resolutions)
	count := 0
	httpClient := &http.Client{}
	writeApi := influxdb.Client.WriteAPIBlocking("glassnode", "glassnode")
	fmt.Println(resolutions)
	fmt.Printf("Collectig %v assets with resolutions: %v\n", len(assets), len(resolutions))
	for _, asset := range assets {
		baseUrl := fmt.Sprintf("%v%v?a=%v&f=JSON", config.Viper.GetString("GLASSNODE_BASE_URL"), endpoint.Path, asset.Symbol)
		for _, resolution := range resolutions {
			currectURL := fmt.Sprintf("%v&i=%v&s=1167526800", baseUrl, resolution)
			req, err := http.NewRequest("GET", currectURL, nil)
			if err == nil {
				req.Header.Set("X-Api-Key", apiKey)
				res, err := httpClient.Do(req)
				if err == nil {
					defer res.Body.Close()
					var points []types.ChartData
					json.NewDecoder(res.Body).Decode(&points)
					for _, point := range points {
						p := influxdb2.NewPointWithMeasurement("charts").
							AddTag("path", endpoint.Path).
							AddTag("resolution", resolution).
							AddTag("asset", asset.Symbol).
							AddField("value", point.Value).
							SetTime(time.Unix(int64(point.Time), 0))
						err := writeApi.WritePoint(context.Background(), p)
						if err != nil {
							fmt.Println(err.Error())
						}
					}
					writeApi.Flush(context.Background())
					count = count + 1
					if count == 25 {
						count = 0
						time.Sleep(time.Minute * 1)
					}
				} else {
					fmt.Println(err.Error())
				}
			}
		}
	}
	childgroup.Initialized = true
	database.Connection.Save(&childgroup)
}
