package workers

import (
	"time"

	"github.com/ario-team/glassnode-api/functions"
	"github.com/ario-team/glassnode-api/logger"
	"github.com/ario-team/glassnode-api/schema"
)

var ChartUpdateJobs = make(chan ChartUpdateInput, 500)

type ChartUpdateInput struct {
	Endpoint schema.Endpoint
	Time     time.Time
}

func InitializeChartUpdateJobs() {
	for i := 0; i < 6; i++ {
		go chartUpdateWorker(ChartUpdateJobs, i)
	}

}

func chartUpdateWorker(endpoints <-chan ChartUpdateInput, id int) {
	for endpoint := range endpoints {
		logger.Logger.Printf("Worker %v: Updating the chart with %v id", endpoint.Endpoint.ID, id)
		functions.UpdateChart(endpoint.Endpoint, endpoint.Time)
		logger.Logger.Printf("Worker %v: Updated the chart with %v id", endpoint.Endpoint.ID, id)
	}
}
