package userhandler

import (
	"net/http"

	"github.com/SnackLog/auth-service/internal/database/user"
	"github.com/SnackLog/auth-service/internal/handlers"
	"github.com/gin-gonic/gin"
)

// Delete godoc
// @Summary Delete user
// @Description Deletes the authenticated user account
// @Tags user
// @Produce json
// @Security BearerAuth
// @Success 204 "No Content"
// @Failure 401 "Unauthorized"
// @Failure 500 {object} handlers.Error
// @Router /auth/user [delete]
func (u *UserController) Delete(c *gin.Context) {
	username := c.GetString("username")

	err := user.DeleteUser(u.DB, username)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, handlers.Error{Error: "Failed to delete user"})
		return
	}

	c.Status(http.StatusNoContent)
}
