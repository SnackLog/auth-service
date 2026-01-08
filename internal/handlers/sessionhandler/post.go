package sessionhandler

import (
	"fmt"
	"log"
	"net/http"

	argonhashutils "github.com/LightJack05/argon-hash-utils"
	"github.com/SnackLog/auth-service/internal/database/user"
	"github.com/gin-gonic/gin"
)

type loginRequestBody struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// Post Creates a new session
func (s *SessionController) Post(c *gin.Context) {
	var body loginRequestBody
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	message, err := s.authenticateUser(body.Username, body.Password)
	if err != nil {
		log.Printf("Could not authenticate account: %v", err)
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authentication failed"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": message})
}

func (s *SessionController) authenticateUser(username, password string) (string, error) {
	user, err := user.GetUserByUsername(s.DB, username)
	if err != nil {
		return "", fmt.Errorf("failed to get user: %v", err)
	}
	if user == nil {
		return "", fmt.Errorf("user not found")
	}

	hash, err := argonhashutils.ParseHash(user.PasswordHash)
	if err != nil {
		return "", fmt.Errorf("failed to parse password hash: %v", err)
	}

	sentHash := argonhashutils.HashPassword(password, hash.Memory, hash.Time, hash.Parallelism, hash.Salt, uint32(len(hash.Hash)))
	if !argonhashutils.CompareHashes(hash, sentHash) {
		return "", fmt.Errorf("invalid password")
	}
	return "success!", nil
}
