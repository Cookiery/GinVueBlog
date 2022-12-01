package api

import (
	"main/commond/errmsg"
	"main/model/service/mysql"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// 查询分类名是否存在
func CategoryExist(ctx *gin.Context) {
	ctx.Writer.Header().Set("SM-Encrypted", "true")
	ctx.JSON(http.StatusOK, gin.H{
		"status": 200,
	})
}

// 添加分类
func AddCategory(ctx *gin.Context) {
	var data mysql.Category
	ctx.ShouldBindJSON(&data)
	statusCode := mysql.CategoryExist(data.Name)
	if statusCode == errmsg.SUCCSE {
		mysql.AddCategory(&data)
	}
	if statusCode == errmsg.ERROR_CATEGORY_USED {
		statusCode = errmsg.ERROR_CATEGORY_USED
	}
	ctx.JSON(http.StatusOK, gin.H{
		"status":  statusCode,
		"data":    data,
		"message": errmsg.ErrMsg(statusCode),
	})
}

// 删除分类
func DeleteCategory(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	statusCode := mysql.DeleteCategory(id)
	ctx.JSON(http.StatusOK, gin.H{
		"status":  statusCode,
		"message": errmsg.ErrMsg(statusCode),
	})
}

// 查询单个分类下的文章

// 查询分类列表
func GetCategory(ctx *gin.Context) {
	pageSize, _ := strconv.Atoi(ctx.Query("pageSize"))
	pageNum, _ := strconv.Atoi(ctx.Query("pageNum"))
	if pageSize == 0 {
		pageSize = -1
	}
	if pageNum == 0 {
		pageNum = -1
	}
	data := mysql.GetCategory(pageSize, pageNum)
	statusCode := errmsg.SUCCSE
	ctx.JSON(http.StatusOK, gin.H{
		"status":  statusCode,
		"data":    data,
		"message": errmsg.ErrMsg(statusCode),
	})
}
