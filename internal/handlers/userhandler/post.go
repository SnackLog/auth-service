package userhandler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (user *UserController) Post(c *gin.Context) {
	c.AbortWithStatus(http.StatusNotImplemented)	
}
