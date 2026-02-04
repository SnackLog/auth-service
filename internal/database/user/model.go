package user

import "time"

// User represents a user in the database
type User struct {
	ID            string
	Username      string
	DisplayName   string
	PasswordHash  string
	Birthdate     *time.Time
	Sex           *string
	Weight        *float64
	ActivityLevel *float64
}
