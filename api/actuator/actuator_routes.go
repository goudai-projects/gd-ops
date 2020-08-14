package actuator

import "github.com/gin-gonic/gin"

func Routes(router *gin.Engine) {
	health := new(HealthController)
	router.GET("/health", health.Status)
}
