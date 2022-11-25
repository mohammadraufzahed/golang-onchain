package functions

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/ario-team/glassnode-api/config"
	"github.com/ario-team/glassnode-api/database"
	"github.com/ario-team/glassnode-api/logger"
	"github.com/ario-team/glassnode-api/schema"
)

func GetEndPoints() {
	baseURL := config.Viper.GetString("GLASSNODE_BASE_URL")
	apiKey := config.Viper.GetString("GLASSNODE_API_KEY")
	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/v2/metrics/endpoints", baseURL), nil)
	if err != nil {
		logger.Logger.Panic(err)
	}
	req.Header.Set("X-Api-Key", apiKey)
	res, err := httpClient.Do(req)
	if err != nil {
		logger.Logger.Panic(err)
	}
	if res.StatusCode != http.StatusOK {
		b, _ := io.ReadAll(res.Body)
		logger.Logger.Printf("Non-OK HTTP Status: %v. body: %v\n", res.StatusCode, string(b))
		panic(string(b))
	}
	defer res.Body.Close()
	var data []schema.Endpoint
	json.NewDecoder(res.Body).Decode(&data)
	for _, endpoint := range data {
		if strings.Contains(endpoint.Path, "indicators") {
			if !strings.Contains(endpoint.Path, "sopr") {
				continue
			}
		}
		var dbEndPoint schema.Endpoint
		result := database.Connection.Where("path = ?", endpoint.Path).Take(&dbEndPoint)
		if result.RowsAffected == 0 {
			database.Connection.Create(&endpoint)
		} else {
			dbEndPoint.Tier = endpoint.Tier
			dbEndPoint.Assets = endpoint.Assets
			dbEndPoint.Currencies = endpoint.Currencies
			dbEndPoint.Resolutions = endpoint.Resolutions
			dbEndPoint.Formats = endpoint.Formats
			database.Connection.Save(&dbEndPoint)
		}
	}
	fmt.Println("Job: Endpoints finished")
}
