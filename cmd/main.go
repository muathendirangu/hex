package main

import (
	"flag"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/muathendirangu/hex/internal/adapters/http"
	"github.com/muathendirangu/hex/internal/adapters/repository/postgres"
	"github.com/muathendirangu/hex/internal/adapters/repository/redis"
	"github.com/muathendirangu/hex/internal/core/services"
)

var (
	repo      = flag.String("db", "postgres", "Database for storing messages")
	redisHost = "localhost:6379"
	svc       *services.MessageService
)

func main() {
	flag.Parse()

	fmt.Printf("Application running using %s\n", *repo)
	switch *repo {
	case "redis":
		store := redis.NewMessageRedisRepository(redisHost)
		svc = services.NewMessageService(store)
	default:
		store := postgres.NewMessagePostgresRepository()
		svc = services.NewMessageService(store)
	}

	InitRoutes()
}

func InitRoutes() {
	router := gin.Default()
	handler := http.NewHTTPHandler(*svc)
	router.GET("/messages/:id", handler.ReadMessage)
	router.GET("/messages", handler.ReadMessages)
	router.POST("/messages", handler.SaveMessage)
	router.Run(":5000")
}
