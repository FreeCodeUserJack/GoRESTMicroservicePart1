package healthcheck

import (
	"net/http"

	"github.com/gin-gonic/gin"
)


const (
	health = "good"
)

func HandleHealthCheck(c *gin.Context) {
	c.String(http.StatusOK, health)
}