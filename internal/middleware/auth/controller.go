package auth

import (
	"database/sql"
)

type AuthController struct {
	DB *sql.DB
}
