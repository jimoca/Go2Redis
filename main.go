package main

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

var redisDB *redis.Client

func init() {
	redisDB = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "!rvqii",
		DB:       0,
	})

}

func main() {
	pong, err := redisDB.Ping(context.Background()).Result()
	if err == nil {
		log.Println("Redis response: ", pong)
	} else {
		log.Fatal("Error: ", err)
	}
	router := gin.Default()
	router.GET("/api/hello", hello)
	router.Run("localhost:8080")
func hello(c *gin.Context) {
	c.JSON(http.StatusOK, "hello")
}
}
