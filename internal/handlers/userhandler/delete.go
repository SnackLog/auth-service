package userhandler

import (
	"net/http"

	"github.com/SnackLog/auth-service/internal/database/user"
	"github.com/gin-gonic/gin"
)

// Delete Deletes an existing user
func (u *UserController) Delete(c *gin.Context) {
	username := c.GetString("username")

	err := user.DeleteUser(u.DB, username)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}
