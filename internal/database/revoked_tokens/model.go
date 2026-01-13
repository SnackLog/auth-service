package revokedtokens

import (
	"database/sql"
	"fmt"
	"log"
)

type RevokedToken struct {
	TokenID string `db:"token_uuid"`
}

func RevokeToken(db *sql.DB, tokenID string) (*RevokedToken, error) {
	_, err := db.Exec("INSERT INTO revoked_tokens (token_uuid) VALUES ($1)", tokenID)
	if err != nil {
		log.Println("Failed to revoke token:", err)
		return nil, fmt.Errorf("Failed to revoke token: %v", err)
	}
	return &RevokedToken{TokenID: tokenID}, nil
}
