package redis

func Exists(key string) int64 {
	result, err := Connection.Exists(Ctx, key).Result()
	if err != nil {
		panic(err)
	}
	return result
}
