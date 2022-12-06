package api

import (
	"main/commond/errmsg"
	"main/commond/validator"
	"main/model/service/mysql"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var errCode int

// AddUser 添加用户
func AddUser(ctx *gin.Context) {
	var user mysql.User
	ctx.ShouldBindJSON(&user)
	msg, statusCode := validator.Validate(&user)
	if statusCode != errmsg.SUCCSE {
		ctx.JSON(http.StatusOK, gin.H{
			"stats":   statusCode,
			"message": msg,
		})
		return
	}
	errCode = mysql.UserExist(user.UserName)
	if errCode == errmsg.ERROR_USERNAME_USED {
		errCode = errmsg.ERROR_USERNAME_USED
	}
	if errCode == errmsg.SUCCSE {
		errCode = mysql.AddUser(&user)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  errCode,
		"data":    user,
		"message": errmsg.ErrMsg(errCode),
	})
}

// UserExist 查询用户是否存在
func UserExist(ctx *gin.Context) {

}

// UserInfo 用户信息
func UserInfo(ctx *gin.Context) {

}

// EditUser 编辑用户
func EditUser(ctx *gin.Context) {
	var user mysql.User
	id, _ := strconv.Atoi(ctx.Param("id"))
	ctx.ShouldBindJSON(&user)
	statusCode := mysql.UserExist(user.UserName)
	if statusCode == errmsg.SUCCSE {
		mysql.EditUser(id, &user)
	}
	if statusCode == errmsg.ERROR_USERNAME_USED {
		ctx.Abort()
	}
	ctx.JSON(http.StatusOK, gin.H{
		"status":  statusCode,
		"message": errmsg.ErrMsg(statusCode),
	})
}

// 删除用户
func DeleteUser(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	statusCode := mysql.DeleteUser(id)
	ctx.JSON(http.StatusOK, gin.H{
		"status":  statusCode,
		"message": errmsg.ErrMsg(statusCode),
	})
}

// 获取用户列表
func GetUsers(ctx *gin.Context) {
	pageSize, _ := strconv.Atoi(ctx.Query("pageSize"))
	pageNum, _ := strconv.Atoi(ctx.Query("pageNum"))
	if pageSize == 0 {
		pageSize = -1
	}
	if pageNum == 0 {
		pageNum = -1
	}
	data, total := mysql.GetUsers(pageSize, pageNum)
	statusCode := errmsg.SUCCSE
	ctx.JSON(http.StatusOK, gin.H{
		"status":  statusCode,
		"total":   total,
		"data":    data,
		"message": errmsg.ErrMsg(statusCode),
	})
}
