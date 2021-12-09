package database

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
)

var Cache *redis.Client
var CacheChannel chan string

func SetupRedis() {
	Cache = redis.NewClient(&redis.Options{
		// container名:port
		Addr: "redis:2379",
		// redisが自動的に作成するデータベースの最初の一つを使うことを以下に明示
		DB: 0,
	})
}

func SetupCacheChannel() {
	//initialize channel
	CacheChannel = make(chan string)

	go func(ch chan string) {
		for {
			time.Sleep(5 * time.Second)

			key := <-ch

			Cache.Del(context.Background(), key)

			fmt.Printf("Cache cleared %s", key)
		}
	}(CacheChannel)
}

func ClearCache(keys ...string) {
	for _, key := range keys {
		CacheChannel <- key
	}
}
