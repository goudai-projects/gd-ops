package server

import (
	"github.com/gin-gonic/gin"
	"github.com/goudai-projects/gdops/api/actuator"
)

func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	// add routes
	actuator.Routes(router)

	return router
}
