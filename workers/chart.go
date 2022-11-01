package workers

import (
	"fmt"

	"github.com/ario-team/glassnode-api/functions"
)

var ChartJobs = make(chan ChartInput, 2)
var ChartsJobsLen = 0

type ChartInput struct {
	EndpointID uint
	ChartID    uint
}

func InitializeChartJobs() {
	for i := 0; i < 2; i++ {
		go worker(ChartJobs, i)
	}
}

func worker(endpoints <-chan ChartInput, id int) {
	for endpoint := range endpoints {
		ChartsJobsLen = ChartsJobsLen + 1
		fmt.Printf("Worker %v started\n", id)
		functions.InitializeChart(endpoint.EndpointID, endpoint.ChartID)
		fmt.Printf("Worker %v finished\n", id)
		ChartsJobsLen = ChartsJobsLen - 1
	}
}
