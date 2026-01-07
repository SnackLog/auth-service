package sessionhandler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// DeleteID Deletes session by ID
func (s *SessionController) DeleteID(c *gin.Context) {
	c.AbortWithStatus(http.StatusNotImplemented)
}
