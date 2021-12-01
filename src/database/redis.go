package database

import "github.com/go-redis/redis/v8"

var Cache *redis.Client

func SetupRedis() {
	Cache = redis.NewClient(&redis.Options{
		// container名:port
		Addr: "redis:2379",
		// redisが自動的に作成するデータベースの最初の一つを使うことを以下に明示
		DB: 0,
	})
}
