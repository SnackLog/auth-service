package userhandler

import (
	"crypto/rand"
	"log"
	"net/http"
	"syscall"
	"time"

	"github.com/SnackLog/auth-service/internal/crypto"
	"github.com/SnackLog/auth-service/internal/database/user"
	"github.com/SnackLog/auth-service/internal/handlers"
	"github.com/gin-gonic/gin"
)

type userBody struct {
	Username    string `json:"username" binding:"required"`
	Password    string `json:"password" binding:"required,min=8"`
	DisplayName string `json:"display_name" binding:"required"`

	Birthdate     time.Time `json:"birthdate" binding:"required"`
	Sex           string    `json:"sex" binding:"required,len=1"`
	Weight        float64   `json:"weight" binding:"required"`
	ActivityLevel float64   `json:"activity_level" binding:"required"`
}

// Post godoc
// @Summary Register
// @Description Creates a new user account
// @Tags user
// @Accept json
// @Produce json
// @Param body body userBody true "User registration details"
// @Success 201 "Created"
// @Failure 400 {object} handlers.Error
// @Failure 500 {object} handlers.Error
// @Router /auth/user [post]
func (u *UserController) Post(c *gin.Context) {
	var body userBody
	if err := c.ShouldBindJSON(&body); err != nil {
		c.Error(err)
		c.JSON(http.StatusBadRequest, handlers.Error{Error: "Invalid request body"})
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
		Username:      body.Username,
		DisplayName:   body.DisplayName,
		PasswordHash:  crypto.HashPassword(body.Password, salt),
		Birthdate:     body.Birthdate,
		Sex:           body.Sex,
		Weight:        body.Weight,
		ActivityLevel: body.ActivityLevel,
	}

	if err := user.CreateUser(u.DB, userStruct); err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, handlers.Error{Error: "Failed to create user"})
		return
	}

	c.Status(http.StatusCreated)
}
