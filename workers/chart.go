package workers

import (
	"github.com/ario-team/glassnode-api/functions"
	"github.com/ario-team/glassnode-api/logger"
)

var ChartJobs = make(chan ChartInput, 500)

type ChartInput struct {
	EndpointID uint
}

func InitializeChartJobs() {
	for i := 0; i < 6; i++ {
		go worker(ChartJobs, i)
	}

}

func worker(endpoints <-chan ChartInput, id int) {
	for endpoint := range endpoints {
		logger.Logger.Printf("Worker %v: Collecting the chart with %v id", endpoint.EndpointID, id)
		functions.InitializeChart(endpoint.EndpointID)
		logger.Logger.Printf("Worker %v: Collected the chart with %v id", endpoint.EndpointID, id)
	}
}
