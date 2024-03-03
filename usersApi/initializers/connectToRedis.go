package initializers

import (
	"github.com/go-redis/redis"
	"os"
)

var RedisClient *redis.Client

func ConnectToRedis() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_URL"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0, // use default DB
	})

	_, err := RedisClient.Ping().Result()
	if err != nil {
		panic(err)
	}
}
