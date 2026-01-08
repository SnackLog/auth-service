package health

import "github.com/gin-gonic/gin"

// Get handles requests to the /health endpoint
func (hc *HealthController) Get(c *gin.Context) {
	err := hc.DB.Ping()
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"error": "database unavail"})
		return
	}
	c.Status(200)
}
