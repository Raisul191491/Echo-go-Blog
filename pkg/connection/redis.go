package connection

import "github.com/go-redis/redis"

var client *redis.Client

func RedisConnect() {
	client = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0, // use default DB
	})
}

func GetRedis() *redis.Client {
	if client == nil {
		RedisConnect()
	}
	return client
}
