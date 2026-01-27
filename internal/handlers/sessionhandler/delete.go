package sessionhandler

import (
	"net/http"

	"github.com/SnackLog/auth-service/internal/crypto"
	revokedtokens "github.com/SnackLog/auth-service/internal/database/revoked_tokens"
	"github.com/SnackLog/auth-service/internal/handlers"
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
// @Failure 400 {object} handlers.Error
// @Failure 401 {object} handlers.Error
// @Failure 500 {object} handlers.Error
// @Router /auth/session [delete]
func (s *SessionController) Delete(c *gin.Context) {
	var body deleteSessionBody
	if err := c.ShouldBindJSON(&body); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, handlers.Error{Error: "Invalid request body"})
		return
	}
	claims, err := crypto.ParseAndValidateToken(body.Token)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, handlers.Error{Error: "Invalid token"})
		return
	}
	if claims.Subject != c.GetString("username") {
		c.AbortWithStatusJSON(http.StatusUnauthorized, handlers.Error{Error: "Invalid token"})
		return
	}

	_, err = revokedtokens.RevokeToken(s.DB, claims.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, handlers.Error{Error: "Failed to revoke token"})
		return
	}

	c.Status(http.StatusNoContent)
}
