package userhandler

import (
	"net/http"

	"github.com/SnackLog/auth-service/internal/database/user"
	"github.com/SnackLog/auth-service/internal/handlers"
	"github.com/gin-gonic/gin"
)

type userGetResponse struct {
	Username    string `json:"username"`
	DisplayName string `json:"display_name"`
}

// Get godoc
// @Summary Get user profile
// @Description Retrieves the profile of the authenticated user
// @Tags user
// @Produce json
// @Security BearerAuth
// @Success 200 {object} userGetResponse
// @Failure 401 "Unauthorized"
// @Failure 404 {object} handlers.Error
// @Failure 500 {object} handlers.Error
// @Router /auth/user [get]
func (u *UserController) Get(c *gin.Context) {
	username := c.GetString("username")
	user, err := user.GetUserByUsername(u.DB, username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, handlers.Error{Error: "Failed to retrieve user"})
		return
	}
	if user == nil {
		c.JSON(http.StatusNotFound, handlers.Error{Error: "User not found"})
		return
	}

	response := userGetResponse{
		Username:    user.Username,
		DisplayName: user.DisplayName,
	}
	c.JSON(http.StatusOK, response)

}
