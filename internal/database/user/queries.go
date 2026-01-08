package user

import (
	"database/sql"
	"fmt"
)

func GetUserByUsername(db *sql.DB, username string) (*User, error) { 
	// query the table users to get user by username
	sqlStatement := `SELECT id, username, display_name, password_hash FROM users WHERE username=$1`
	var user User
	row := db.QueryRow(sqlStatement, username)
	err := row.Scan(&user.ID, &user.Username, &user.DisplayName, &user.PasswordHash)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // User not found
		}
		return nil, fmt.Errorf("failed to query user: %v", err)
	}
	return &user, nil

}
