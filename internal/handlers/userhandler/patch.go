package userhandler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Patch Updates an existing user
func (u *UserController) Patch(c *gin.Context) {
	c.AbortWithStatus(http.StatusNotImplemented)	
}
