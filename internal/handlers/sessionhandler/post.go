package sessionhandler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *SessionController) Post(c *gin.Context) {
	c.AbortWithStatus(http.StatusNotImplemented)	
}
