package sessionhandler

import "database/sql"

type SessionController struct {
	DB *sql.DB
}
