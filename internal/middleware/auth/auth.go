package auth

import (
	"log"
	"net/http"

	"github.com/SnackLog/auth-service/internal/crypto"
	revokedtokens "github.com/SnackLog/auth-service/internal/database/revoked_tokens"
	"github.com/gin-gonic/gin"
)

func (a *AuthController) Authenticate(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	if len(authHeader) < len("Bearer ")+1 || authHeader[:len("Bearer ")] != "Bearer " {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	token := authHeader[len("Bearer "):]
	claims, err := crypto.ParseAndValidateToken(token)
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	revoked, err := revokedtokens.IsTokenRevoked(a.DB, claims.ID)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	if revoked {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	c.Set("username", claims.Subject)
	log.Println(claims.Subject)
	log.Println(c.GetString("username"))
	c.Next()
}
