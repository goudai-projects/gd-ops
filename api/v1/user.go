package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/goudai-projects/gd-ops/log"
	"github.com/goudai-projects/gd-ops/model"
	"net/http"
)

type UserController struct {
	Api
}

func (api *Api) InitUser() {
	userController := UserController{
		*api,
	}
	api.router.GET("/user/search", userController.SearchAllPaged)
}

func (res *UserController) SearchAllPaged(c *gin.Context) {
	var userSearch model.UserSearch
	if c.ShouldBindQuery(&userSearch) == nil {
		log.Infof("Username: %s, Page: %d, Size: %d", userSearch.Username,
			userSearch.Page, userSearch.Size)
	}
	users, total, err := res.Srv.SearchUsersPage(c, &userSearch)
	if err != nil {

	}
	c.JSON(http.StatusOK, model.NewPage(&userSearch.Pageable, total, users))
}
