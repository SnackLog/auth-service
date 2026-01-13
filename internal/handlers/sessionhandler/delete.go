package sessionhandler

import (
	"net/http"

	"github.com/SnackLog/auth-service/internal/crypto"
	revokedtokens "github.com/SnackLog/auth-service/internal/database/revoked_tokens"
	"github.com/gin-gonic/gin"
)

type deleteSessionBody struct {
	Token string `json:"token" binding:"required"`
}

// Delete Revokes a session of a user
func (s *SessionController) Delete(c *gin.Context) {
	var body deleteSessionBody
	if err := c.ShouldBindJSON(&body); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}
	claims, err := crypto.ParseAndValidateToken(body.Token)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}
	if claims.Subject != c.GetString("username") {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
	}
	_, err = revokedtokens.RevokeToken(s.DB, claims.ID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to revoke token"})
		return
	}

	c.Status(http.StatusNoContent)
}
