package main

import "github.com/gin-gonic/gin"

func main() {

	// db := MySQL.InitDB()
	// MySQL.InsertData(db)

	server := InitServer()
	// 启动服务
	gin.SetMode(server.AppMode) // 设置模式
	r := gin.Default()          // Default() 带有 Logger 和 Recovery 中间件
	r.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{"msg": "hello"})
	})
	r.Run(server.HTTPPort)
}
