package workers

import (
	"github.com/ario-team/glassnode-api/functions"
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
		functions.InitializeChart(endpoint.EndpointID)
	}
}
