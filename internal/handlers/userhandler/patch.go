package userhandler

import (
	"fmt"
	"net/http"

	"github.com/SnackLog/auth-service/internal/database/user"
	"github.com/gin-gonic/gin"
)

type userPatchBody struct {
	DisplayName *string `json:"displayName,omitempty"`
}

// Patch Updates an existing user
func (u *UserController) Patch(c *gin.Context) {
	username := c.GetString("username")

	var body userPatchBody
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	if err := u.updateUser(body, username); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
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
