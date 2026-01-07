package userhandler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Get Retrieves an existing user
func (u *UserController) Get(c *gin.Context) {
	c.AbortWithStatus(http.StatusNotImplemented)	
}
