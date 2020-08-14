package actuator

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type HealthController struct{}

type Health struct {
	Status string
}

func (h HealthController) Status(c *gin.Context) {
	c.JSON(http.StatusOK, &Health{Status: "ok"})
}
