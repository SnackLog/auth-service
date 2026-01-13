package userhandler

import (
	"crypto/rand"
	"log"
	"net/http"
	"syscall"

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
	_, err := rand.Read(salt)
	if err != nil {
		log.Println("rand.Read returned an error, this should not happen! (This means the OS is unable to provide proper crypto APIs. Do NOT run this program here.)")
		log.Println("Tearing down the application with SIGABRT...")
		syscall.Kill(syscall.Getpid(), syscall.SIGABRT)
	}


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
