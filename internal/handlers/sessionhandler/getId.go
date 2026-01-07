package sessionhandler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *SessionController) GetID(c *gin.Context) {
	c.AbortWithStatus(http.StatusNotImplemented)	
}
