package sessionhandler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// DeleteAll Deletes all session of the user issuing the request
func (s *SessionController) DeleteAll(c *gin.Context) {
	c.AbortWithStatus(http.StatusNotImplemented)
}
