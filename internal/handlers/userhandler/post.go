package userhandler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (u *UserController) Post(c *gin.Context) {
	c.AbortWithStatus(http.StatusNotImplemented)	
}
