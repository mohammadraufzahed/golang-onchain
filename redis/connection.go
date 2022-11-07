package redis

import (
	"context"
	"fmt"

	"github.com/ario-team/glassnode-api/config"
	"github.com/go-redis/redis/v9"
)

var Connection *redis.Client
var Ctx = context.Background()

func Connect() {
	Connection = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%v:6379", config.Viper.GetString("REDIS_HOST")),
		Password: "",
		DB:       0,
	})
}
