package sessionhandler

import "database/sql"

// SessionController Controller for sessions, to be used by Gin
type SessionController struct {
	DB *sql.DB
}
