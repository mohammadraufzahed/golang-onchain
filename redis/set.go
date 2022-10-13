package redis

import (
	"time"
)

func Set(key string, value any, duration time.Duration) {
	err := Connection.Set(Ctx, key, value, duration*time.Second).Err()
	if err != nil {
		panic(err)
	}
}
