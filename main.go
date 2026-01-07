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

	"github.com/SnackLog/auth-service/internal/database"
	"github.com/SnackLog/auth-service/internal/handlers/sessionhandler"
	"github.com/SnackLog/auth-service/internal/handlers/userhandler"
)

func main() {
	initConfig()
	migrateDatabase()
	db := initDatabaseConnection()

	initApi(db)
}

// initApi initializes the API server and routes.
func initApi(db *sql.DB) {
	router := gin.Default()
	router.Use(cors.Default())

	auth := router.Group("/auth")
	setupAuthEndpoints(auth, db)

	router.Run(":80")
}

// setupAuthEndpoints sets up the authentication-related API endpoints.
func setupAuthEndpoints(auth *gin.RouterGroup, db *sql.DB) {
	setupUserEndpoints(auth, db)

	setupSessionEndpoints(auth, db)
}

func setupSessionEndpoints(auth *gin.RouterGroup, db *sql.DB) {
	sessionController := sessionhandler.SessionController{
		DB: db,
	}

	auth.GET("/session/:id", sessionController.GetID)
	auth.POST("/session", sessionController.Post)
	auth.DELETE("/session", sessionController.Delete)
	auth.DELETE("/session/:id", sessionController.DeleteID)
}

func setupUserEndpoints(auth *gin.RouterGroup, db *sql.DB) {
	userController := userhandler.UserController{
		DB: db,
	}

	auth.GET("/user", userController.Get)
	auth.POST("/user", userController.Post)
	auth.PATCH("/user", userController.Patch)
	auth.DELETE("/user", userController.Delete)
}

// initDatabaseConnection initializes the database connection.
func initDatabaseConnection() *sql.DB {
	db, err := database.Connect(databaseConfig.GetDatabaseConnectionString())
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
//
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
