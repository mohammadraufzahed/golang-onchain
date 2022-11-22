package main

import (
	"github.com/ario-team/glassnode-api/config"
	"github.com/ario-team/glassnode-api/database"
	"github.com/ario-team/glassnode-api/functions"
	"github.com/ario-team/glassnode-api/influxdb"
	job "github.com/ario-team/glassnode-api/jobs"
	"github.com/ario-team/glassnode-api/logger"
	"github.com/ario-team/glassnode-api/redis"
	"github.com/ario-team/glassnode-api/router"
	"github.com/ario-team/glassnode-api/schema"
	"github.com/ario-team/glassnode-api/sentry"
	"github.com/ario-team/glassnode-api/web"
	"github.com/ario-team/glassnode-api/workers"
)

func main() {
	// Initilize the config
	config.Initialize()
	// Initialize the logger
	logger.Init()
	// Initialize the sentry
	sentry.Init()
	// Connect to the redis
	redis.Connect()
	defer redis.Connection.Close()
	influxdb.Connect()
	defer influxdb.Client.Close()
	// Connect to database
	database.Connect()
	database.Migrate()
	// // Initialized the endpoints
	functions.GetEndPoints()
	// Start the charts initializing
	StartInitializing()
	// Register Jobs
	job.RegisterJobs()
	// Initialize Router and middleware
	web.InitailizeMiddlewares()
	router.InitializeRouter()
	// Initialize workers
	workers.InitializeWorkers()
	// Start the webserver
	web.Start()
}

func StartInitializing() {
	var endpoints []schema.Endpoint
	database.Connection.Not("initialized = ?", true).Find(&endpoints)
	for _, endpoint := range endpoints {
		workers.ChartJobs <- workers.ChartInput{
			EndpointID: endpoint.ID,
		}
	}
}
