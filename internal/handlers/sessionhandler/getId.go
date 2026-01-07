package sessionhandler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetID Retrieves session by ID
func (s *SessionController) GetID(c *gin.Context) {
	c.AbortWithStatus(http.StatusNotImplemented)
}
