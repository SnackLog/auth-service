package userhandler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (u *UserController) Get(c *gin.Context) {
	c.AbortWithStatus(http.StatusNotImplemented)	
}
