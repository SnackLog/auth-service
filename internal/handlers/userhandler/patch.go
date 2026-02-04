package userhandler

import (
	"fmt"
	"net/http"
	"time"

	"github.com/SnackLog/auth-service/internal/database/user"
	"github.com/SnackLog/auth-service/internal/handlers"
	"github.com/gin-gonic/gin"
)

type userPatchBody struct {
	DisplayName *string `json:"displayName,omitempty"`

	Birthdate     *time.Time `json:"birthdate,omitempty"`
	Sex           *string    `json:"sex,omitempty" binding:"omitempty,len=1"`
	Weight        *float64   `json:"weight,omitempty"`
	ActivityLevel *float64   `json:"activityLevel,omitempty"`
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
	tx, err := u.DB.Begin()
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %v", err)
	}
	defer tx.Rollback()
	if patchBody.DisplayName != nil {
		if err := user.UpdateDisplayName(tx, username, *patchBody.DisplayName); err != nil {
			return fmt.Errorf("failed to update display name: %v", err)
		}
	}
	if patchBody.Birthdate != nil {
		if err := user.UpdateBirthdate(tx, username, *patchBody.Birthdate); err != nil {
			return fmt.Errorf("failed to update birthdate: %v", err)
		}
	}
	if patchBody.Sex != nil {
		if err := user.UpdateSex(tx, username, *patchBody.Sex); err != nil {
			return fmt.Errorf("failed to update sex: %v", err)
		}
	}
	if patchBody.Weight != nil {
		if err := user.UpdateWeight(tx, username, *patchBody.Weight); err != nil {
			return fmt.Errorf("failed to update weight: %v", err)
		}
	}
	if patchBody.ActivityLevel != nil {
		if err := user.UpdateActivityLevel(tx, username, *patchBody.ActivityLevel); err != nil {
			return fmt.Errorf("failed to update activity level: %v", err)
		}
	}
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %v", err)
	}
	return nil
}
