package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func DummyHandler(c *gin.Context) {
	c.AbortWithStatus(http.StatusNotImplemented)
}
