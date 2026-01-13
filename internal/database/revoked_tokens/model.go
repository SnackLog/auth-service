package revokedtokens

import (
)

// RevokedToken represents a revoked token
type RevokedToken struct {
	TokenID string `db:"token_uuid"`
}

