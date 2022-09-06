package main

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/gin-gonic/gin"
)

// Server 服务变量
type Server struct {
	AppMode  string
	HTTPPort string
}

// InitServer 初始化服务变量
func InitServer(server *Server) {
	if _, err := toml.DecodeFile("./conf/app.toml", &server); err != nil {
		// TODO 打日志
		fmt.Println(err)
	}
	gin.SetMode(server.AppMode) // 设置模式
}

// InitLogger 初始化日志模块
func InitLogger() {
	logFile, err := os.Create("./log/service.log")
	if err != nil {
		fmt.Println("Could not create log file")
	}
	gin.DefaultWriter = io.MultiWriter(logFile)
}

// InitRouter 初始化路由
func InitRouter(HTTPPort string) {
	r := gin.Default() // Default() 默认带有 Logger 和 Recovery 中间件
	router := r.Group("api/v1")
	{
		router.GET("hello", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"hello": "cookie",
			})
		})
	}
	r.Run(HTTPPort) // 启动服务
}
