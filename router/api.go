package router

import (
	"main/actions/api"

	"github.com/gin-gonic/gin"
)

func RegisterAPI(router *gin.RouterGroup) {
	// 用户模块的路由接口
	router.POST("user/add", api.AddUser)
	router.GET("user/info", api.UserInfo)
	router.PUT("user/:id", api.EditUser)
	router.DELETE("user/:id", api.DeleteUser)

	// 分类模块的路由接口

	// 文章模块的路由接口
}
