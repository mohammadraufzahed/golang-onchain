package functions

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/ario-team/glassnode-api/config"
	"github.com/ario-team/glassnode-api/database"
	"github.com/ario-team/glassnode-api/influxdb"
	"github.com/ario-team/glassnode-api/redis"
	"github.com/ario-team/glassnode-api/schema"
	"github.com/ario-team/glassnode-api/types"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)

func InitializeChart(endpointID uint) {
	var endpoint schema.Endpoint
	database.Connection.Where("id = ?", endpointID).First(&endpoint)
	var assets []types.Assets
	var resolutions []string
	json.Unmarshal([]byte(endpoint.Assets), &assets)
	json.Unmarshal([]byte(endpoint.Resolutions), &resolutions)
	apiKey := config.Viper.GetString("GLASSNODE_API_KEY")
	count := 0
	httpClient := &http.Client{}
	writeApi := influxdb.Client.WriteAPIBlocking("glassnode", "glassnode")
	fmt.Printf("Collecting %v assets with %v resolutions\n", len(assets), len(resolutions))
	for _, asset := range assets {
		baseUrl := fmt.Sprintf("%v%v?a=%v&f=JSON", config.Viper.GetString("GLASSNODE_BASE_URL"), endpoint.Path, asset.Symbol)
		for _, resolution := range resolutions {
			currectURL := fmt.Sprintf("%v&i=%v&s=1167526800", baseUrl, resolution)
			req, err := http.NewRequest("GET", currectURL, nil)
			if err == nil {
				req.Header.Set("X-Api-Key", apiKey)
				res, err := httpClient.Do(req)

				if err == nil {
					if res.StatusCode == 200 {
						defer res.Body.Close()
						var points []types.ChartData
						json.NewDecoder(res.Body).Decode(&points)
						fmt.Printf("Collected %v points\n", len(points))
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
						err := writeApi.Flush(context.Background())
						if err == nil {
							fmt.Printf("Writed %v data.\n", len(points))
						} else {

							fmt.Printf("Faild %v data.\n", len(points))
						}
						count = count + 1
						if count == 40 {
							count = 0
							time.Sleep(time.Minute * 1)
						}
					}
				} else {
					fmt.Println(err.Error())
				}
			}
		}
	}
	endpoint.Initialized = true
	database.Connection.Save(&endpoint)
}

type chartReturn struct {
	T int64 `json:"time"`
	V any   `json:"value"`
}

func GetChart(id int, asset string, interval string, start uint, end uint) (chart string, err error) {
	var endpoint schema.Endpoint
	database.Connection.Where("id = ?", id).First(&endpoint)
	redisKey := fmt.Sprintf("%v/%v/%v/%v/%v", endpoint.Path, asset, interval, start, end)
	exists := redis.Exists(redisKey)
	var chartJson string
	if exists == 1 {
		return redis.Get(redisKey), nil
	} else {
		queryApi := influxdb.Client.QueryAPI("glassnode")
		result, err := queryApi.Query(context.Background(), fmt.Sprintf(`from (bucket: "glassnode") |> range(start: %v, stop: %v) |> filter(fn: (r) => r.asset == "%v" and r.resolution == "%v")`, start, end, strings.ToUpper(asset), interval))
		if err != nil {
			fmt.Println("Error")
			return "", err
		}
		var chart []chartReturn
		for result.Next() {
			chart = append(chart, chartReturn{
				T: result.Record().Time().Unix(),
				V: result.Record().Value(),
			})
		}
		chartJson, err := json.Marshal(chart)
		if err == nil {
			redis.Set(redisKey, chartJson, time.Hour)
		}
	}
	return chartJson, err
}
