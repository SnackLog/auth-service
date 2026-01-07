package userhandler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Post Creates a new user
func (u *UserController) Post(c *gin.Context) {
	c.AbortWithStatus(http.StatusNotImplemented)	
}
