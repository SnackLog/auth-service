package sessionhandler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (sessionController *SessionController) DeleteID(c *gin.Context) {
	c.AbortWithStatus(http.StatusNotImplemented)	
}
