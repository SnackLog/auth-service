package userhandler

import "database/sql"

// UserController Controller for users, to be used by Gin
type UserController struct {
	DB *sql.DB
}
