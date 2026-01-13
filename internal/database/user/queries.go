package user

import (
	"database/sql"
	"fmt"
)

// GetUserByUsername returns a matching user
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

// CreateUser creates a new user
func CreateUser(db *sql.DB, user *User) error {
	// insert into table users
	sqlStatement := `INSERT INTO users (username, display_name, password_hash) VALUES ($1, $2, $3)`
	_, err := db.Exec(sqlStatement, user.Username, user.DisplayName, user.PasswordHash)
	if err != nil {
		return fmt.Errorf("failed to create user: %v", err)
	}
	return nil
}

func UpdateDisplayName(db *sql.DB, username, displayName string) error {
	// update display_name in table users
	sqlStatement := `UPDATE users SET display_name=$1 WHERE username=$2`
	_, err := db.Exec(sqlStatement, displayName, username)
	if err != nil {
		return fmt.Errorf("failed to update display name in database: %v", err)
	}
	return nil
}

func DeleteUser(db *sql.DB, username string) error {
	// TODO: Make sure to notify other services about user deletion

	// delete the user object
	sqlStatement := `DELETE FROM users WHERE username=$1`
	_, err := db.Exec(sqlStatement, username)
	if err != nil {
		return fmt.Errorf("failed to delete user from database: %v", err)
	}
	return nil
}
