package main

import (
	"database/sql"
	"embed"
	"fmt"

	databaseConfig "github.com/SnackLog/database-config-lib"
	serviceConfig "github.com/SnackLog/service-config-lib"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/golang-migrate/migrate/v4/source/iofs"

	"github.com/SnackLog/auth-service/internal/db"
	"github.com/SnackLog/auth-service/internal/handlers"
)

func main() {
	initConfig()
	migrateDatabase()
	_ = initDatabaseConnection()

	initApi()
}

// initApi initializes the API server and routes.
func initApi() {
	router := gin.Default()
	router.Use(cors.Default())

	auth := router.Group("/auth")
	setupAuthEndpoints(auth)

	router.Run(":80")
}

// setupAuthEndpoints sets up the authentication-related API endpoints.
func setupAuthEndpoints(auth *gin.RouterGroup) {
	auth.GET("/user", handlers.DummyHandler)
	auth.POST("/user", handlers.DummyHandler)
	auth.DELETE("/user", handlers.DummyHandler)
	auth.PATCH("/user", handlers.DummyHandler)

	auth.POST("/session", handlers.DummyHandler)
	auth.DELETE("/session", handlers.DummyHandler)

	auth.GET("/session/:id", handlers.DummyHandler)
	auth.DELETE("/session/:id", handlers.DummyHandler)
}

// initDatabaseConnection initializes the database connection.
func initDatabaseConnection() *sql.DB {
	db, err := db.Connect(databaseConfig.GetDatabaseConnectionString())
	if err != nil {
		panic(fmt.Errorf("Failed to connect to database: %v", err))
	}
	return db
}

// migrateDatabase runs database migrations.
func migrateDatabase() {
	err := doMigrations()
	if err != nil {
		panic(fmt.Sprintf("Database migration failed: %v", err))
	}
}

// migrationFiles embeds SQL migration files.
//go:embed db/migrations/*.sql
var migrationFiles embed.FS

// doMigrations performs database migrations using embedded SQL files.
func doMigrations() error {
	migrationDriver, err := iofs.New(migrationFiles, "db/migrations")
	if err != nil {
		return fmt.Errorf("Failed to create iofs driver: %v", err)
	}

	m, err := migrate.NewWithSourceInstance(
		"iofs",
		migrationDriver,
		databaseConfig.GetDatabaseConnectionString(),
	)

	if err != nil {
		return fmt.Errorf("Failed to create migrate instance: %v", err)
	}

	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("Failed to run migrations: %v", err)
	}

	return nil
}

// initConfig initializes service and database configurations.
func initConfig() {
	err := serviceConfig.LoadConfig()
	if err != nil {
		panic(fmt.Sprintf("Failed to load service configuration: %v", err))
	}

	err = databaseConfig.LoadConfig()
	if err != nil {
		panic(fmt.Sprintf("Failed to load database configuration: %v", err))
	}
}
