package userhandler

import (
	"net/http"

	"github.com/SnackLog/auth-service/internal/database/user"
	"github.com/gin-gonic/gin"
)

type userGetResponse struct {
	Username    string
	DisplayName string
}

// Get Retrieves an existing user
func (u *UserController) Get(c *gin.Context) {
	username := c.GetString("username")
	user, err := user.GetUserByUsername(u.DB, username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve user"})
		return
	}
	if user == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	response := userGetResponse{
		Username:    user.Username,
		DisplayName: user.DisplayName,
	}
	c.JSON(http.StatusOK, response)

}
