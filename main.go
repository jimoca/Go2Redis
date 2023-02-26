package main

import (
	"context"
	"log"

	"github.com/go-redis/redis/v8"
)

var redisDB *redis.Client

func init() {
	redisDB = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "!rvqii", // no password set
		DB:       0,        // use default DB
	})
}

func main() {
	pong, err := redisDB.Ping(context.Background()).Result()
	if err == nil {
		log.Println("Redis response: ", pong)
	} else {
		log.Fatal("Error: ", err)
	}
}
