package util

import (
	"fmt"
	"testing"
	"time"
)

func TestInitRedis(t *testing.T) {
	InitRedis()

	result, _ := RedisClient.Get("hello").Result()
	fmt.Println(result) // ""

	result, _ = RedisClient.Set("hello", 1, time.Second * 2).Result()
	fmt.Println(result) // OK

	result, _ = RedisClient.Get("hello").Result()
	fmt.Println(result) // 1

	time.Sleep(time.Second * 2)
	result, _ = RedisClient.Get("hello").Result()
	fmt.Println(result) // ""
}
