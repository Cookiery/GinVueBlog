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

// 查询分类下的所有文章
func GetCategoryArticles(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	pageSize, _ := strconv.Atoi(ctx.Query("pageSize"))
	pageNum, _ := strconv.Atoi(ctx.Query("pageNum"))
	if pageSize == 0 {
		pageSize = -1
	}
	if pageNum == 0 {
		pageNum = -1
	}
	data, statusCode := mysql.GetCategoryArticles(id, pageSize, pageNum)
	ctx.JSON(http.StatusOK, gin.H{
		"status":  statusCode,
		"data":    data,
		"message": errmsg.ErrMsg(statusCode),
	})
}

// 查询单个文章
func GetArticleInfo(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	data, statusCode := mysql.GetArticleInfo(id)
	ctx.JSON(http.StatusOK, gin.H{
		"status":  statusCode,
		"data":    data,
		"message": errmsg.ErrMsg(statusCode),
	})
}

// 查询文章列表
func GetArticles(ctx *gin.Context) {
	pageSize, _ := strconv.Atoi(ctx.Query("pageSize"))
	pageNum, _ := strconv.Atoi(ctx.Query("pageNum"))
	if pageSize == 0 {
		pageSize = -1
	}
	if pageNum == 0 {
		pageSize = -1
	}
	data, statusCode := mysql.GetArticles(pageSize, pageNum)
	ctx.JSON(http.StatusOK, gin.H{
		"status":  statusCode,
		"data":    data,
		"message": errmsg.ErrMsg(statusCode),
	})
}

// 删除文章
func DeleteArticle(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	statusCode := mysql.DeleteArticle(id)
	ctx.JSON(http.StatusOK, gin.H{
		"status":  statusCode,
		"message": errmsg.ErrMsg(statusCode),
	})
}
