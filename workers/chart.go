package workers

import (
	"fmt"

	"github.com/ario-team/glassnode-api/functions"
)

var ChartJobs = make(chan ChartInput, 500)
var ChartsJobsLen = 0

type ChartInput struct {
	EndpointID uint
}

func InitializeChartJobs() {
	go worker(ChartJobs, 1)

}

func worker(endpoints <-chan ChartInput, id int) {
	for endpoint := range endpoints {
		ChartsJobsLen = ChartsJobsLen + 1
		fmt.Printf("Worker %v started\n", id)
		functions.InitializeChart(endpoint.EndpointID)
		fmt.Printf("Worker %v finished\n", id)
		ChartsJobsLen = ChartsJobsLen - 1
	}
}
