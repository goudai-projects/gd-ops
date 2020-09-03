package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/goudai-projects/gd-ops/model"
	"net/http"
)

type ClusterController struct {
	Api
}

func (api *Api) InitCluster() {
	clusterController := ClusterController{
		*api,
	}
	api.router.GET("/v1/clr/list", clusterController.list)
	api.router.GET("/v1/clr/create", clusterController.create)
}

func (c *ClusterController) list(ctx *gin.Context) {
	clusterList, err := c.Srv.ClusterList()
	if err != nil {
		ctx.AbortWithError(500, err)
		return
	}
	ctx.JSON(http.StatusOK, clusterList)
}

func (c *ClusterController) create(ctx *gin.Context) {
	var dto model.ClusterCreateDto
	if err := ctx.BindJSON(&dto); err != nil {
		ctx.AbortWithError(500, err)
		return
	}
	userId := ctx.GetInt64("userId")
	if userId == 0 {

	}
	//_, err := c.Srv.ClusterCreate(userId, &dto)

}
