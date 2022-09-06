package main

// main 程序入口
// 程序入口主要是初始化各个模块
func main() {
	var server Server // gin服务配置

	// 初始化数据库
	// MySQL.InitDB()
	// MySQL.WriteUser("yxq", "123")
	// MySQL.ReadUser("yxq")
	// MySQL.InsertData(db)

	// 初始化日志
	InitLogger()
	// 初始化服务配置
	InitServer(&server)
	// 初始化路由和启动
	InitRouter(server.HTTPPort)
}
