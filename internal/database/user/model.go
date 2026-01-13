package user

// User represents a user in the database
type User struct {
	ID           string
	Username     string
	DisplayName  string
	PasswordHash string
}
