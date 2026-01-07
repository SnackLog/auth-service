package sessionhandler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Delete Deletes all sessions
func (s *SessionController) Delete(c *gin.Context) {
	c.AbortWithStatus(http.StatusNotImplemented)	
}
