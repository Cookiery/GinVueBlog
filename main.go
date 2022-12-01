package main

import "main/model/service/mysql"

// 第一个需求：记录自己读过的书！！！
// 1. Go Web 编程
//

// 第二个需求：刷题展示
// 参考：https://github.com/halfrost/LeetCode-Go

// 抓一些热榜：https://github.com/tophubs/TopList

// 修改图片尺寸大小

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
