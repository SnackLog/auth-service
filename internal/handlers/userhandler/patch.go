package userhandler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (user *UserController) Patch(c *gin.Context) {
	c.AbortWithStatus(http.StatusNotImplemented)	
}
