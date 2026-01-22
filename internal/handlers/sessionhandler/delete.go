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

// Delete godoc
// @Summary Logout
// @Description Revokes a user session
// @Tags session
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param body body deleteSessionBody true "Token to revoke"
// @Success 204 "No Content"
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /auth/session [delete]
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
		return
	}

	_, err = revokedtokens.RevokeToken(s.DB, claims.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to revoke token"})
		return
	}

	c.Status(http.StatusNoContent)
}
