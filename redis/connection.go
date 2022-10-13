package redis

import (
	"context"

	"github.com/go-redis/redis/v9"
)

var Connection *redis.Client
var Ctx = context.Background()

func Connect() {
	Connection = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
}
