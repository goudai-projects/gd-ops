package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/goudai-projects/gd-ops/app"
)

type Api struct {
	Srv    *app.Server
	router *gin.Engine
}

func Init(srv *app.Server, router *gin.Engine) *Api {
	api := Api{
		Srv:    srv,
		router: router,
	}
	api.InitUser()
	return &api
}
