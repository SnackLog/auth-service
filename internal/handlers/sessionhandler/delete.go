package sessionhandler

import (
	"net/http"

	"github.com/SnackLog/auth-service/internal/crypto"
	revokedtokens "github.com/SnackLog/auth-service/internal/database/revoked_tokens"
	"github.com/gin-gonic/gin"
)

type DeleteSessionBody struct {
	Token string `json:"token" binding:"required"`
}

// Delete Revokes a session of a user
func (s *SessionController) Delete(c *gin.Context) {
	var body DeleteSessionBody
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}
	claims, err := crypto.ParseAndValidateToken(body.Token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}
	_, err = revokedtokens.RevokeToken(s.DB, claims.ID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to revoke token"})
		return
	}

	c.Status(http.StatusNoContent)
}
