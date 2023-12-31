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
	"github.com/ario-team/glassnode-api/logger"
	"github.com/ario-team/glassnode-api/redis"
	"github.com/ario-team/glassnode-api/schema"
	"github.com/ario-team/glassnode-api/sentry"
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
	apiCount := 0
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
								sentry.Sentry.CaptureException(err)
								logger.Logger.Panic(err)
							}
						}
						err := writeApi.Flush(context.Background())
						if err == nil {
							fmt.Printf("Writed %v data.\n", len(points))
						} else {

							fmt.Printf("Faild %v data.\n", len(points))
						}
						count = count + 1
						if count == 10 {
							count = 0
							time.Sleep(time.Minute * 1)
						}
					}
				} else {
					apiCount++
					if apiCount > 2 {
						sentry.Sentry.CaptureException(err)
						logger.Logger.Panic(err)
					}
				}
			}
		}
	}
	endpoint.Initialized = true
	database.Connection.Save(&endpoint)
}

func UpdateChart(endpoint schema.Endpoint, start time.Time) {
	var assets []types.Assets
	var resolutions []string
	json.Unmarshal([]byte(endpoint.Assets), &assets)
	json.Unmarshal([]byte(endpoint.Resolutions), &resolutions)
	apiKey := config.Viper.GetString("GLASSNODE_API_KEY")
	count := 0
	httpClient := &http.Client{}
	writeApi := influxdb.Client.WriteAPIBlocking("glassnode", "glassnode")
	fmt.Printf("Updating chart %v with %v assets with %v resolutions\n", endpoint.ID, len(assets), len(resolutions))
	for _, asset := range assets {
		baseUrl := fmt.Sprintf("%v%v?a=%v&f=JSON", config.Viper.GetString("GLASSNODE_BASE_URL"), endpoint.Path, asset.Symbol)
		for _, resolution := range resolutions {
			currectURL := fmt.Sprintf("%v&i=%v&s=%v", baseUrl, resolution, start.Unix())
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
								sentry.Sentry.CaptureException(err)
								logger.Logger.Panic(err)
							}
						}
						err := writeApi.Flush(context.Background())
						if err == nil {
							fmt.Printf("Updated %v data.\n", len(points))
						} else {

							fmt.Printf("Faild %v data.\n", len(points))
						}
						count = count + 1
						if count == 10 {
							count = 0
							time.Sleep(time.Minute * 1)
						}
					}
				} else {
					sentry.Sentry.CaptureException(err)
					logger.Logger.Panic(err)
				}
			}
		}
	}
	database.Connection.Model(&endpoint).Update("initialized", true)
}
func GetChart(path string, asset string, resolution string, start int, end int) (string, error) {
	redisKey := fmt.Sprintf("%v/%v/%v/%v/%v", path, asset, resolution, start, end)
	exists := redis.Exists(redisKey)
	if exists == 1 {
		return redis.Get(redisKey), nil
	} else {
		queryApi := influxdb.Client.QueryAPI("glassnode")
		result, err := queryApi.Query(context.Background(), fmt.Sprintf(`from (bucket: "glassnode") |> range(start: %v, stop: %v) |> filter(fn: (r) => r.asset == "%v" and r.resolution == "%v" and r.path == "%v")`, start, end, strings.ToUpper(asset), resolution, path))
		if err != nil {
			return "", err
		}
		var chart []types.ChartGetResponse
		for result.Next() {
			chart = append(chart, types.ChartGetResponse{
				T: result.Record().Time().Unix(),
				V: result.Record().Value(),
			})
		}
		chartJson, err := json.Marshal(chart)
		if err != nil {
			return "", err
		}

		redis.Set(redisKey, chartJson, time.Hour)
		return string(chartJson), nil
	}
}
