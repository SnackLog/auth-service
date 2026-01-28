package sessionhandler

import (
	"fmt"
	"log"
	"net/http"

	argonhashutils "github.com/LightJack05/argon-hash-utils"
	"github.com/SnackLog/auth-service/internal/crypto"
	"github.com/SnackLog/auth-service/internal/database/user"
	"github.com/SnackLog/auth-service/internal/handlers"
	"github.com/gin-gonic/gin"
)

type loginRequestBody struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type loginResponseBody struct {
	Token string `json:"token"`
}

// Post godoc
// @Summary Login
// @Description Authenticates a user and returns a token
// @Tags session
// @Accept json
// @Produce json
// @Param body body loginRequestBody true "Login credentials"
// @Success 200 {object} loginResponseBody 
// @Failure 400 {object} handlers.Error
// @Failure 401 {object} handlers.Error
// @Failure 500 {object} handlers.Error
// @Router /auth/session [post]
func (s *SessionController) Post(c *gin.Context) {
	var body loginRequestBody
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, handlers.Error{Error: "Invalid request body"})
		return
	}

	authenticated, err := s.isCredentialsValid(body.Username, body.Password)
	if err != nil {
		log.Printf("Could not authenticate account: %v", err)
		c.AbortWithStatusJSON(http.StatusUnauthorized, handlers.Error{Error: "Authentication failed"})
		return
	}
	if !authenticated {
		c.AbortWithStatusJSON(http.StatusUnauthorized, handlers.Error{Error: "Authentication failed"})
		return
	}
	userToken, err := crypto.CreateAuthToken(body.Username)
	if err != nil {
		log.Println(fmt.Errorf("error signing token: %v", err))
		c.AbortWithStatusJSON(http.StatusInternalServerError, handlers.Error{Error: "failed to sign token for authentication"})
		return
	}
	c.JSON(http.StatusOK, loginResponseBody{
		Token: userToken,
	})
}

func (s *SessionController) isCredentialsValid(username, password string) (bool, error) {
	user, err := user.GetUserByUsername(s.DB, username)
	if err != nil {
		log.Println(fmt.Errorf("failed to get user: %v", err))
		return false, nil

	}
	if user == nil {
		return false, nil
	}

	hash, err := argonhashutils.ParseHash(user.PasswordHash)
	if err != nil {
		log.Println(fmt.Errorf("failed to parse password hash for user %v: %v", username, err))
		return false, nil
	}

	sentHash := argonhashutils.HashPassword(password, hash.Memory, hash.Time, hash.Parallelism, hash.Salt, uint32(len(hash.Hash)))
	if !argonhashutils.CompareHashes(hash, sentHash) {
		return false, nil
	}
	return true, nil
}
