package main

import "main/model/service/mysql"

// main 程序入口
// 程序入口主要是初始化各个模块
func main() {

	// 初始化数据库
	mysql.InitDB()
	// mysql.WriteUser("yxq", "123")
	// mysql.ReadUser("yxq")

	var server Server // gin服务配置
	// 初始化日志
	// InitLogger()
	// 初始化服务配置
	InitServer(&server)
	// 初始化路由和启动
	InitRouter(server.HTTPPort)
}
