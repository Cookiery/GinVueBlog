package main

import (
	"fmt"
	"io"
	"main/router"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/gin-gonic/gin"
)

// logFile 日志存储路径
const logFile = "./log/service.log"

// Server Gin 服务配置
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
	f, err := os.Create(logFile)
	if err != nil {
		fmt.Println("Could not create log file")
	}
	gin.DefaultWriter = io.MultiWriter(f)
}

// InitRouter 初始化路由
func InitRouter(HTTPPort string) {
	r := gin.Default() // Default() 默认带有 Logger 和 Recovery 中间件

	routerAPI := r.Group("api/")
	router.RegisterAPI(routerAPI)

	r.Run(HTTPPort) // 启动服务
}
