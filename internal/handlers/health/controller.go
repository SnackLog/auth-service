package health

import "database/sql"

// HealthController Controller for sessions, to be used by Gin
type HealthController struct {
	DB *sql.DB
}
