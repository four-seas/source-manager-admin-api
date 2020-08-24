package router

import (
	"github.com/gin-gonic/gin"
	v1 "source_manager_admin_api/api/v1"
)

func InitHouseRouter(Router *gin.RouterGroup) {
	UserRouter := Router.Group("house")
	{
		UserRouter.GET("/v1/list", v1.GetHouseFang)     // 修改密码
	}
}
