package userhandler

import (
	"fmt"
	"net/http"

	"github.com/SnackLog/auth-service/internal/database/user"
	"github.com/SnackLog/auth-service/internal/handlers"
	"github.com/gin-gonic/gin"
)

type userPatchBody struct {
	DisplayName *string `json:"displayName,omitempty"`
}

// Patch godoc
// @Summary Update user profile
// @Description Updates the profile of the authenticated user
// @Tags user
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param body body userPatchBody true "Fields to update"
// @Success 204 "No Content"
// @Failure 400 {object} handlers.Error
// @Failure 401 "Unauthorized"
// @Failure 500 {object} handlers.Error
// @Router /auth/user [patch]
func (u *UserController) Patch(c *gin.Context) {
	username := c.GetString("username")

	var body userPatchBody
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, handlers.Error{Error: "Invalid request body"})
		return
	}

	if err := u.updateUser(body, username); err != nil {
		c.JSON(http.StatusInternalServerError, handlers.Error{Error: "Failed to update user"})
		return
	}

	c.Status(http.StatusNoContent)
}

func (u *UserController) updateUser(patchBody userPatchBody, username string) error {
	if patchBody.DisplayName != nil {
		if err := u.updateDisplayName(username, *patchBody.DisplayName); err != nil {
			return fmt.Errorf("failed to update display name: %v", err)
		}
	}
	return nil
}

func (u *UserController) updateDisplayName(username, displayName string) error {
	err := user.UpdateDisplayName(u.DB, username, displayName)
	if err != nil {
		return fmt.Errorf("request failed: %v", err)
	}
	return nil
}
