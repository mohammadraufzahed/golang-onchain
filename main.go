package main

import (
	"github.com/ario-team/glassnode-api/config"
	"github.com/ario-team/glassnode-api/redis"
)

func main() {
	// Connect to the redis
	redis.Connect()
	defer redis.Connection.Close()
	// Initilize the config
	config.Initialize()
}
