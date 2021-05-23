package util

import (
	"fmt"
	"github.com/go-redis/redis"
)

const (
	redisPassword = `passwd`
	redisHost     = `127.0.0.1`
	redisPort     = `6379`
)

var RedisClient *redis.Client

func InitRedis() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", redisHost, redisPort),
		Password: redisPassword,
	})

	_, err := RedisClient.Ping().Result()
	if err != nil {
		panic(err)
	}

	fmt.Println("Connect to Redis!")
}
