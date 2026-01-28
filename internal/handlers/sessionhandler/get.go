package sessionhandler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type sessionGetResponse struct {
	Username string `json:"username"`
}

// Get godoc
// @Summary Get session info
// @Description Retrieves session information by its token
// @Tags session
// @Produce json
// @Security BearerAuth
// @Success 200 {object} SessionGetResponse
// @Failure 401 "Unauthorized"
// @Router /auth/session [get]
func (s *SessionController) Get(c *gin.Context) {
	c.JSON(http.StatusOK, sessionGetResponse{
		Username: c.GetString("username"),
	})
}
