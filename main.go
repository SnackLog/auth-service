package main

import (
	"fmt"
	"log"

	databaseConfig "github.com/SnackLog/database-config-lib"
	serviceConfig "github.com/SnackLog/service-config-lib"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/SnackLog/auth-service/internal/handlers"
)

func main() {
	loadConfig()
	doMigrations()

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

func doMigrations() {
	panic("unimplemented")
}

func loadConfig() {
	err := serviceConfig.LoadConfig()
	if err != nil {
		panic(fmt.Sprintf("Failed to load service configuration: %v", err))
	}

	err = databaseConfig.LoadConfig()
	if err != nil {
		panic(fmt.Sprintf("Failed to load database configuration: %v", err))
	}
}
