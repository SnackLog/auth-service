package sessionhandler

import (
	"net/http"

	"github.com/SnackLog/auth-service/internal/crypto"
	revokedtokens "github.com/SnackLog/auth-service/internal/database/revoked_tokens"
	"github.com/gin-gonic/gin"
)

// GetID Retrieves session by ID
func (s *SessionController) GetID(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	token := authHeader[len("Bearer "):]
	claims, err := crypto.ParseAndValidateToken(token)
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	revoked, err := revokedtokens.IsTokenRevoked(s.DB, claims.ID)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	if revoked {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"username": claims.Subject,
	})
}
