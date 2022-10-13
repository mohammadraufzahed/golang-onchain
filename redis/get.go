package redis

func Get(key string) string {
	result, err := Connection.Get(Ctx, key).Result()
	if err != nil {
		panic(err)
	}
	return result
}
