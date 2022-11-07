package main

import (
	"fmt"

	"github.com/ario-team/glassnode-api/config"
	"github.com/ario-team/glassnode-api/database"
	"github.com/ario-team/glassnode-api/functions"
	"github.com/ario-team/glassnode-api/influxdb"
	job "github.com/ario-team/glassnode-api/jobs"
	"github.com/ario-team/glassnode-api/redis"
	"github.com/ario-team/glassnode-api/router"
	"github.com/ario-team/glassnode-api/schema"
	"github.com/ario-team/glassnode-api/web"
	"github.com/ario-team/glassnode-api/workers"
)

func main() {
	// Initilize the config
	config.Initialize()
	// Connect to the redis
	redis.Connect()
	defer redis.Connection.Close()
	// Connect to database
	database.Connect()
	database.Migrate()
	// Initialized the endpoints
	functions.GetEndPoints()
	// Start the charts initializing
	StartInitializing()
	// Register Jobs
	job.RegisterJobs()
	// Initialize Router and middleware
	web.InitailizeMiddlewares()
	router.InitializeRouter()
	// fmt.Println(time.Now().UTC().Unix())
	influxdb.Connect()
	defer influxdb.Client.Close()
	// Initialize workers
	workers.InitializeWorkers()
	// Start the webserver
	web.Start()
}

func StartInitializing() {
	var endpoints []schema.Endpoint
	database.Connection.Not("initialized = ?", true).Find(&endpoints)
	fmt.Println(len(endpoints))
	for _, endpoint := range endpoints {
		workers.ChartJobs <- workers.ChartInput{
			EndpointID: endpoint.ID,
		}
		fmt.Printf("Added %v\n", endpoint.ID)
	}
	fmt.Println("Jobs Passed")
}
