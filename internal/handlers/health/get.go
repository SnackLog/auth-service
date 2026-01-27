package health

import (
	"github.com/SnackLog/auth-service/internal/handlers"
	"github.com/gin-gonic/gin"
)

// Get handles requests to the /health endpoint
func (hc *HealthController) Get(c *gin.Context) {
	err := hc.DB.Ping()
	if err != nil {
		c.AbortWithStatusJSON(500, handlers.Error{Error: "database unavail"})
		return
	}
	c.Status(200)
}
