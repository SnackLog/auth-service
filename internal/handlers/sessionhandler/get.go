package sessionhandler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type SessionGetResponse struct {
	Username string `json:"username"`
}

// Get godoc
// @Summary Get session info
// @Description Retrieves session information by its token
// @Tags session
// @Produce json
// @Security BearerAuth
// @Success 200 {object} map[string]string
// @Failure 401 "Unauthorized"
// @Router /auth/session [get]
func (s *SessionController) Get(c *gin.Context) {
	c.JSON(http.StatusOK, SessionGetResponse{
		Username: c.GetString("username"),
	})
}
