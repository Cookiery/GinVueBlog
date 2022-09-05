package main

import (
	"fmt"
	"main/model/service/MySQL"

	"github.com/BurntSushi/toml"
)

// Server 启动服务
type Server struct {
	AppMode  string
	HTTPPort string
}

// InitServer 初始化服务变量
func InitServer() *Server {
	var server Server
	if _, err := toml.DecodeFile("./conf/app.toml", &server); err != nil {
		// TODO 打日志
		fmt.Println(err)
	}
	return &server
}

func main() {
	MySQL.InitDB()
	// MySQL.WriteUser("yxq", "123")
	MySQL.ReadUser("yxq")
	// MySQL.InsertData(db)

	// 初始化服务
	// server := InitServer()
	// gin.SetMode(server.AppMode) // 设置模式
	// r := gin.Default()          // Default() 带有 Logger 和 Recovery 中间件
	// r.GET("/", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{"hello": "cookie"})
	// })
	// r.Run(server.HTTPPort)
}
