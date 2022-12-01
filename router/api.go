package router

import (
	"main/actions/api"
	"main/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterAPI(router *gin.RouterGroup, authAPI *gin.RouterGroup) {

	// 用户模块的路由接口
	router.GET("user/info", api.UserInfo)
	router.GET("users", api.GetUsers)

	// 分类模块的路由接口
	router.GET("category/:id", api.CategoryExist)

	// 文章模块的路由接口
	router.GET("article", api.GetArticles)
	router.GET("article/info/:id", api.GetArticleInfo)      // 查询单个文章
	router.GET("article/list/:id", api.GetCategoryArticles) // 查询分类下所有文章

	// 登录
	router.POST("login", api.Login)

	authAPI.Use(middleware.JWTToken()) // 注册JWT验证中间件

	// 用户模块的路由接口
	authAPI.POST("user/add", api.AddUser)
	authAPI.PUT("user/:id", api.EditUser)
	authAPI.DELETE("user/:id", api.DeleteUser)

	// 分类模块的路由接口
	authAPI.POST("category/add", api.AddCategory)
	authAPI.PUT("category/:id", api.EditUser)
	authAPI.DELETE("category/:id", api.DeleteCategory)

	// 文章模块的路由接口
	authAPI.POST("article/add", api.AddArticle)
	authAPI.PUT("article/:id", api.EditArticle)      // 编辑文章
	authAPI.DELETE("article/:id", api.DeleteArticle) // 删除文章
}
