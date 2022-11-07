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
	router.GET("users", api.GetUsers)

	// 分类模块的路由接口
	router.POST("category/add", api.AddCategory)
	router.GET("category/:id", api.CategoryExist)
	router.PUT("category/:id", api.EditUser)
	router.DELETE("category/:id", api.DeleteCategory)

	// 文章模块的路由接口
	router.POST("article/add", api.AddArticle)
	router.GET("article", api.GetArticles)
	router.GET("article/info/:id", api.GetArticleInfo)      // 查询单个文章
	router.GET("article/list/:id", api.GetCategoryArticles) // 查询分类下所有文章
	router.PUT("article/:id", api.EditArticle)              // 编辑文章
	router.DELETE("article/:id", api.DeleteArticle)         // 删除文章
}
