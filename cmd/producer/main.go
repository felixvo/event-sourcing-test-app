package main

import (
	"github.com/felixvo/lmax/cmd/producer/api"
	"github.com/felixvo/lmax/pkg/lhttp"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v7"
)

const (
	MaxUserIDRange = 10000
)

func main() {
	client, err := newRedisClient()
	if err != nil {
		panic(err)
	}
	r := gin.Default()
	r.Use(lhttp.CORSMiddleware())
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.POST("/event/publish", api.NewPublisherHandler(client))
	r.Run(":8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func newRedisClient() (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0, // use default DB
	})

	_, err := client.Ping().Result()
	return client, err

}
