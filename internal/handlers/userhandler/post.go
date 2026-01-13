package userhandler

import (
	"crypto/rand"
	"net/http"

	"github.com/SnackLog/auth-service/internal/crypto"
	"github.com/SnackLog/auth-service/internal/database/user"
	"github.com/gin-gonic/gin"
)

type userBody struct {
	Username    string `json:"username" binding:"required"`
	Password    string `json:"password" binding:"required,min=8"`
	DisplayName string `json:"display_name" binding:"required"`
}

// Post Creates a new user
func (u *UserController) Post(c *gin.Context) {
	var body userBody
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	salt := make([]byte, 16)
	rand.Read(salt)

	userStruct := &user.User{
		Username:     body.Username,
		DisplayName:  body.DisplayName,
		PasswordHash: crypto.HashPassword(body.Password, salt),
	}

	if err := user.CreateUser(u.DB, userStruct); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}
