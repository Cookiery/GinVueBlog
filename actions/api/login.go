package api

import (
	"main/commond/errmsg"
	"main/middleware"
	"main/model/service/mysql"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var (
		userData   mysql.User
		token      string
		statusCode int
	)
	c.ShouldBindJSON(&userData)
	statusCode = mysql.VerifyLogin(userData.UserName, userData.Password)
	if statusCode == errmsg.SUCCSE {
		token, statusCode = middleware.CreateToken(userData.UserName)
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  statusCode,
		"message": errmsg.ErrMsg(statusCode),
		"token":   token,
	})
}
