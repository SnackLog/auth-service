package sessionhandler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *SessionController) Delete(c *gin.Context) {
	c.AbortWithStatus(http.StatusNotImplemented)	
}
