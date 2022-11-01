package main

import (
	"github.com/ario-team/glassnode-api/config"
	"github.com/ario-team/glassnode-api/database"
	"github.com/ario-team/glassnode-api/influxdb"
	job "github.com/ario-team/glassnode-api/jobs"
	"github.com/ario-team/glassnode-api/redis"
	"github.com/ario-team/glassnode-api/router"
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
