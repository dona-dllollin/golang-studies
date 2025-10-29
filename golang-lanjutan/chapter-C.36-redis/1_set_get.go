package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)

func newRedisClient(host string, password string) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     host,
		Password: password, // no password set
		DB:       0,        // use default DB
	})
	return client
}

func main() {
	var redisHost = "localhost:6379"
	var redisPassword = ""

	rdb := newRedisClient(redisHost, redisPassword)
	fmt.Println("Redis client initialized")

	key := "key-1"
	data := "Hello Redis"
	ttl := time.Duration(3) * time.Second

	//store data using SET command
	op1 := rdb.Set(context.Background(), key, data, ttl)
	if err := op1.Err(); err != nil {
		fmt.Printf("Unable to SET data. error: %v", err)
		return
	}
	log.Println("Set operation success")

	time.Sleep(time.Duration(4) * time.Second)

	// get data
	op2 := rdb.Get(context.Background(), key)
	if err := op2.Err(); err != nil {
		fmt.Printf("Unable to GET data. error: %v", err)
		return
	}
	res, err := op2.Result()
	if err != nil {
		fmt.Printf("Unable to GET data. error: %v", err)
		return
	}
	log.Println("get operation success. result:", res)

}
