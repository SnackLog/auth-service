package userhandler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Delete Deletes an existing user
func (u *UserController) Delete(c *gin.Context) {
	c.AbortWithStatus(http.StatusNotImplemented)	
}
