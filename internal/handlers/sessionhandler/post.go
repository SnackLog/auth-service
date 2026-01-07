package sessionhandler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Post Creates a new session
func (s *SessionController) Post(c *gin.Context) {
	c.AbortWithStatus(http.StatusNotImplemented)	
}
