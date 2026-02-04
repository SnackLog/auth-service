package user

import (
	"database/sql"
	"fmt"
	"time"
)

// GetUserByUsername returns a matching user
func GetUserByUsername(db *sql.DB, username string) (*User, error) {
	// query the table users to get user by username
	sqlStatement := `SELECT id, username, display_name, password_hash, birthdate, sex, weight, activity_level FROM users WHERE username=$1`
	var user User
	row := db.QueryRow(sqlStatement, username)
	err := row.Scan(&user.ID, &user.Username, &user.DisplayName, &user.PasswordHash, &user.Birthdate, &user.Sex, &user.Weight, &user.ActivityLevel)
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
	sqlStatement := `INSERT INTO users (username, display_name, password_hash, birthdate, sex, weight, activity_level) VALUES ($1, $2, $3)`
	_, err := db.Exec(sqlStatement, user.Username, user.DisplayName, user.PasswordHash, user.Birthdate, user.Sex, user.Weight, user.ActivityLevel)
	if err != nil {
		return fmt.Errorf("failed to create user: %v", err)
	}
	return nil
}

func UpdateDisplayName(db *sql.Tx, username, displayName string) error {
	// update display_name in table users
	sqlStatement := `UPDATE users SET display_name=$1 WHERE username=$2`
	_, err := db.Exec(sqlStatement, displayName, username)
	if err != nil {
		return fmt.Errorf("failed to update display name in database: %v", err)
	}
	return nil
}

func UpdateBirthdate(db *sql.Tx, username string, birthdate time.Time) error {
	// update birthdate in table users
	sqlStatement := `UPDATE users SET birthdate=$1 WHERE username=$2`
	_, err := db.Exec(sqlStatement, birthdate, username)
	if err != nil {
		return fmt.Errorf("failed to update birthdate in database: %v", err)
	}
	return nil
}

func UpdateSex(db *sql.Tx, username, sex string) error {
	// update sex in table users
	sqlStatement := `UPDATE users SET sex=$1 WHERE username=$2`
	_, err := db.Exec(sqlStatement, sex, username)
	if err != nil {
		return fmt.Errorf("failed to update sex in database: %v", err)
	}
	return nil
}

func UpdateWeight(db *sql.Tx, username string, weight float64) error {
	// update weight in table users
	sqlStatement := `UPDATE users SET weight=$1 WHERE username=$2`
	_, err := db.Exec(sqlStatement, weight, username)
	if err != nil {
		return fmt.Errorf("failed to update weight in database: %v", err)
	}
	return nil
}

func UpdateActivityLevel(db *sql.Tx, username string, activityLevel float64) error {
	// update activity_level in table users
	sqlStatement := `UPDATE users SET activity_level=$1 WHERE username=$2`
	_, err := db.Exec(sqlStatement, activityLevel, username)
	if err != nil {
		return fmt.Errorf("failed to update activity level in database: %v", err)
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
