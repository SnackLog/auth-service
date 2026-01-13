package revokedtokens

import (
	"database/sql"
	"fmt"
	"log"
)

// IsTokenRevoked Checks wether the given token is revoked
func IsTokenRevoked(db *sql.DB, uuid string) (bool, error) {
	var exists bool
	err := db.QueryRow("SELECT EXISTS(SELECT 1 FROM revoked_tokens WHERE token_uuid = ?)", uuid).Scan(&exists)
	if err != nil {
		// In case of error, assume the token is revoked
		return true, err
	}
	return exists, nil
}

// RevokeToken revokes a token with the given id
func RevokeToken(db *sql.DB, tokenID string) (*RevokedToken, error) {
	_, err := db.Exec("INSERT INTO revoked_tokens (token_uuid) VALUES ($1)", tokenID)
	if err != nil {
		log.Println("Failed to revoke token:", err)
		return nil, fmt.Errorf("Failed to revoke token: %v", err)
	}
	return &RevokedToken{TokenID: tokenID}, nil
}
