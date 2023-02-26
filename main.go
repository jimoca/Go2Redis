package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
)

var redisDB *redis.Client

type User struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

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
	router.POST("/api/user", saveUser)
	router.GET("/api/user/:id", getUserById)

	router.Run("localhost:8080")
}

func hello(c *gin.Context) {
	c.JSON(http.StatusOK, "hello")
}

/*
	user {
			name: "Jim",
			age: 100
		}
*/
func saveUser(c *gin.Context) {
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		errorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	user.Id = uuid.New().String()
	json, err := json.Marshal(user)
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	redisDB.HSet(c, "user", user.Id, json)
	c.JSON(http.StatusOK, user)
}

func getUserById(c *gin.Context) {
	var u User
	user, err := redisDB.HGet(c, "user", c.Param("id")).Result()

	if err != nil && err != redis.Nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	err = json.Unmarshal([]byte(user), &u)
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": u,
	})
}

func errorResponse(c *gin.Context, statusCode int, message string) {
	c.JSON(statusCode, gin.H{
		"error": message,
	})
}
