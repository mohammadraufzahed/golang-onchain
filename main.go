package main

import (
	"github.com/ario-team/glassnode-api/config"
	"github.com/ario-team/glassnode-api/database"
	"github.com/ario-team/glassnode-api/redis"
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
}
