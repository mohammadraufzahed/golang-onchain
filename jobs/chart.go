package job

import (
	"time"

	"github.com/ario-team/glassnode-api/database"
	"github.com/ario-team/glassnode-api/schema"
	"github.com/ario-team/glassnode-api/workers"
)

func InitializeChartJobs() {
	Scheduler.Cron("0 0 * * *").Do(ChartsUpdater)
}

func ChartsUpdater() {
	var endpoints []schema.Endpoint
	database.Connection.Where("initialized = ?", true).Find(&endpoints)
	now := time.Now().Add(-24 * time.Hour)
	for _, endpoint := range endpoints {
		workers.ChartUpdateJobs <- workers.ChartUpdateInput{
			Endpoint: endpoint,
			Time:     now,
		}
	}
}
