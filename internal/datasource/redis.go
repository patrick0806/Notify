package datasource

import (
	"github.com/redis/go-redis/v9"
)

var redisClient *redis.Client

func GetRedisClient() *redis.Client {
	if redisClient == nil {

		client := redis.NewClient(&redis.Options{ //TODO: move to config
			Addr:     "localhost:6379",
			Password: "qwerty",
			DB:       0,
		})

		redisClient = client
	}

	return redisClient
}
