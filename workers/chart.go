package workers

import (
	"github.com/ario-team/glassnode-api/functions"
)

var ChartJobs = make(chan ChartInput, 500)
var ChartsJobsLen = 0

type ChartInput struct {
	EndpointID uint
}

func InitializeChartJobs() {
	for i := 0; i < 2; i++ {
		go worker(ChartJobs, i)
	}

}

func worker(endpoints <-chan ChartInput, id int) {
	for endpoint := range endpoints {
		ChartsJobsLen = ChartsJobsLen + 1
		functions.InitializeChart(endpoint.EndpointID)
		ChartsJobsLen = ChartsJobsLen - 1
	}
}
