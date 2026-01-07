package userhandler

import "database/sql"

type UserController struct {
	DB *sql.DB
}
