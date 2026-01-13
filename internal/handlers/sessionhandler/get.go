package sessionhandler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Get Retrieves session by it's token and returns relevant information
func (s *SessionController) Get(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"username": c.GetString("username")})
}
