package api

import (
	"main/commond/errmsg"
	"main/model/service/mysql"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// 添加文章
func AddArticle(ctx *gin.Context) {
	var article mysql.Article
	_ = ctx.ShouldBindJSON(&article)
	statusCode := mysql.AddArticle(&article)
	ctx.JSON(http.StatusOK, gin.H{
		"status":  statusCode,
		"data":    article,
		"message": errmsg.ErrMsg(statusCode),
	})
}

// 查询分类下的所有文章

// 编辑文章
func EditArticle(ctx *gin.Context) {
	var article mysql.Article
	id, _ := strconv.Atoi(ctx.Param("id"))
	_ = ctx.ShouldBindJSON(&article)
	statusCode := mysql.EditArticle(id, &article)
	ctx.JSON(http.StatusOK, gin.H{
		"status":  statusCode,
		"message": errmsg.ErrMsg(statusCode),
	})
}

// 查询单个文章

// 查询文章列表

// 删除文章
func DeleteArticle(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	statusCode := mysql.DeleteArticle(id)
	ctx.JSON(http.StatusOK, gin.H{
		"status":  statusCode,
		"message": errmsg.ErrMsg(statusCode),
	})
}
