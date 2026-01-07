package sessionhandler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *SessionController) DeleteID(c *gin.Context) {
	c.AbortWithStatus(http.StatusNotImplemented)	
}
