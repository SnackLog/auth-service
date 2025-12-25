package main

import (
	"fmt"
	"log"

	serviceConfig "github.com/SnackLog/service-config-lib"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/SnackLog/auth-service/internal/handlers"
)

func main() {
	err := serviceConfig.LoadConfig()
	if err != nil {
		panic(fmt.Sprintf("Failed to load service configuration: %v", err))
	}

	log.Println(serviceConfig.GetConfig().ServiceName)

	router := gin.Default()
	router.Use(cors.Default())

	auth := router.Group("/auth")
	auth.GET("/user", handlers.DummyHandler)
	auth.POST("/user", handlers.DummyHandler)
	auth.DELETE("/user", handlers.DummyHandler)
	auth.PATCH("/user", handlers.DummyHandler)

	auth.POST("/session", handlers.DummyHandler)
	auth.DELETE("/session", handlers.DummyHandler)

	auth.GET("/session/:id", handlers.DummyHandler)
	auth.DELETE("/session/:id", handlers.DummyHandler)

	router.Run(":80")
}
